// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
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

// Package joinserver provides a LoRaWAN 1.1-compliant Join Server implementation.
package joinserver

import (
	"context"

	"go.thethings.network/lorawan-stack/pkg/ttnpb"
)

// GetAppSKey returns the AppSKey associated with session keys identified by the supplied request.
func (js *JoinServer) GetAppSKey(ctx context.Context, req *ttnpb.SessionKeyRequest) (*ttnpb.AppSKeyResponse, error) {
	if req.DevEUI.IsZero() {
		return nil, errNoDevEUI
	}
	if req.SessionKeyID == "" {
		return nil, errNoSessionKeyID
	}

	ks, err := js.keys.GetByID(ctx, req.DevEUI, req.SessionKeyID)
	if err != nil {
		return nil, err
	}
	if ks.AppSKey == nil {
		return nil, errNoAppSKey
	}
	return &ttnpb.AppSKeyResponse{
		AppSKey: *ks.AppSKey,
	}, nil
}
