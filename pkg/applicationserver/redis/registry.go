// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis"
	"go.thethings.network/lorawan-stack/pkg/errors"
	ttnredis "go.thethings.network/lorawan-stack/pkg/redis"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/unique"
)

func applyDeviceFieldMask(dst, src *ttnpb.EndDevice, paths ...string) (*ttnpb.EndDevice, error) {
	if dst == nil {
		dst = &ttnpb.EndDevice{}
	}
	if err := dst.SetFields(src, append(paths, "ids")...); err != nil {
		return nil, err
	}
	if err := dst.EndDeviceIdentifiers.Validate(); err != nil {
		return nil, err
	}
	return dst, nil
}

// DeviceRegistry is a Redis device registry.
type DeviceRegistry struct {
	Redis *ttnredis.Client
}

// Get returns the end device by its identifiers.
func (r *DeviceRegistry) Get(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, paths []string) (*ttnpb.EndDevice, error) {
	k := r.Redis.Key(unique.ID(ctx, ids))
	pb := &ttnpb.EndDevice{}
	if err := ttnredis.GetProto(r.Redis, k).ScanProto(pb); err != nil {
		return nil, err
	}
	return applyDeviceFieldMask(nil, pb, paths...)
}

// Set creates, updates or deletes the end device by its identifiers.
func (r *DeviceRegistry) Set(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, gets []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error) {
	k := r.Redis.Key(unique.ID(ctx, ids))

	var pb *ttnpb.EndDevice
	err := r.Redis.Watch(func(tx *redis.Tx) error {
		var create bool
		cmd := ttnredis.GetProto(tx, k)
		stored := &ttnpb.EndDevice{}
		if err := cmd.ScanProto(stored); errors.IsNotFound(err) {
			create = true
			stored = nil
		} else if err != nil {
			return err
		}

		var err error
		if stored != nil {
			pb, err = applyDeviceFieldMask(nil, stored, gets...)
			if err != nil {
				return err
			}
		}

		var sets []string
		pb, sets, err = f(pb)
		if err != nil {
			return err
		}
		if stored == nil && pb == nil {
			return nil
		}

		var f func(redis.Pipeliner) error
		if pb == nil {
			f = func(p redis.Pipeliner) error {
				p.Del(k)
				return nil
			}
		} else {
			pb.EndDeviceIdentifiers = ids
			pb.UpdatedAt = time.Now().UTC()
			sets = append(sets, "updated_at")
			if create {
				pb.CreatedAt = pb.UpdatedAt
				sets = append(sets, "created_at")
			}
			stored = &ttnpb.EndDevice{}
			if err := cmd.ScanProto(stored); err != nil && !errors.IsNotFound(err) {
				return err
			}
			stored, err = applyDeviceFieldMask(stored, pb, sets...)
			if err != nil {
				return err
			}
			pb, err = applyDeviceFieldMask(nil, stored, gets...)
			if err != nil {
				return err
			}
			f = func(p redis.Pipeliner) error {
				_, err := ttnredis.SetProto(p, k, stored, 0)
				return err
			}
		}

		cmds, err := tx.Pipelined(f)
		if err != nil {
			return err
		}
		for _, cmd := range cmds {
			if err := cmd.Err(); err != nil {
				return err
			}
		}
		return nil
	}, k)
	if err != nil {
		return nil, err
	}
	return pb, nil
}

func applyLinkFieldMask(dst, src *ttnpb.ApplicationLink, paths ...string) (*ttnpb.ApplicationLink, error) {
	if dst == nil {
		dst = &ttnpb.ApplicationLink{}
	}
	return dst, dst.SetFields(src, paths...)
}

// LinkRegistry is a store for application links.
type LinkRegistry struct {
	Redis *ttnredis.Client
}

const (
	allKey  = "all"
	linkKey = "link"
)

// Get returns the link by the application identifiers.
func (r *LinkRegistry) Get(ctx context.Context, ids ttnpb.ApplicationIdentifiers, paths []string) (*ttnpb.ApplicationLink, error) {
	k := r.Redis.Key(linkKey, unique.ID(ctx, ids))
	pb := &ttnpb.ApplicationLink{}
	if err := ttnredis.GetProto(r.Redis, k).ScanProto(pb); err != nil {
		return nil, err
	}
	return applyLinkFieldMask(nil, pb, paths...)
}

var errApplicationUID = errors.DefineCorruption("application_uid", "invalid application UID `{application_uid}`")

// Range ranges the links and calls the callback function, until false is returned.
func (r *LinkRegistry) Range(ctx context.Context, paths []string, f func(context.Context, ttnpb.ApplicationIdentifiers, *ttnpb.ApplicationLink) bool) error {
	uids, err := r.Redis.SMembers(r.Redis.Key(allKey)).Result()
	if err != nil {
		return err
	}
	for _, uid := range uids {
		ctx, err := unique.WithContext(ctx, uid)
		if err != nil {
			return errApplicationUID.WithCause(err).WithAttributes("application_uid", uid)
		}
		ids, err := unique.ToApplicationID(uid)
		if err != nil {
			return errApplicationUID.WithCause(err).WithAttributes("application_uid", uid)
		}
		pb := &ttnpb.ApplicationLink{}
		if err := ttnredis.GetProto(r.Redis, r.Redis.Key(linkKey, uid)).ScanProto(pb); err != nil {
			return err
		}
		pb, err = applyLinkFieldMask(nil, pb, paths...)
		if err != nil {
			return err
		}
		if !f(ctx, ids, pb) {
			break
		}
	}
	return nil
}

// Set creates, updates or deletes the link by the application identifiers.
func (r *LinkRegistry) Set(ctx context.Context, ids ttnpb.ApplicationIdentifiers, gets []string, f func(*ttnpb.ApplicationLink) (*ttnpb.ApplicationLink, []string, error)) (*ttnpb.ApplicationLink, error) {
	uid := unique.ID(ctx, ids)
	k := r.Redis.Key(linkKey, uid)
	var pb *ttnpb.ApplicationLink
	err := r.Redis.Watch(func(tx *redis.Tx) error {
		cmd := ttnredis.GetProto(tx, k)
		stored := &ttnpb.ApplicationLink{}
		if err := cmd.ScanProto(stored); errors.IsNotFound(err) {
			stored = nil
		} else if err != nil {
			return err
		}

		var err error
		if pb != nil {
			pb, err = applyLinkFieldMask(nil, stored, gets...)
			if err != nil {
				return err
			}
		}

		var sets []string
		pb, sets, err = f(pb)
		if err != nil {
			return err
		}

		var f func(redis.Pipeliner) error
		if pb == nil {
			f = func(p redis.Pipeliner) error {
				p.Del(k)
				p.SRem(r.Redis.Key(allKey), uid)
				return nil
			}
		} else {
			stored = &ttnpb.ApplicationLink{}
			if err := cmd.ScanProto(stored); err != nil && !errors.IsNotFound(err) {
				return err
			}
			stored, err = applyLinkFieldMask(stored, pb, sets...)
			if err != nil {
				return err
			}
			pb, err = applyLinkFieldMask(nil, stored, gets...)
			if err != nil {
				return err
			}
			f = func(p redis.Pipeliner) error {
				_, err := ttnredis.SetProto(p, k, stored, 0)
				if err != nil {
					return err
				}
				p.SAdd(r.Redis.Key(allKey), uid)
				return nil
			}
		}
		_, err = tx.Pipelined(f)
		return err
	}, k)
	if err != nil {
		return nil, err
	}
	return pb, nil
}
