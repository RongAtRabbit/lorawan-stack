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

import { createLogic } from 'redux-logic'

import api from '../../api'
import * as devices from '../actions/devices'

const getDevicesListLogic = createLogic({
  type: [
    devices.GET_DEVICES_LIST,
    devices.SEARCH_DEVICES_LIST,
  ],
  async process ({ getState, action }, dispatch, done) {
    const { appId, filters } = action

    try {
      const data = filters.query
        ? await api.devices.search(appId, filters)
        : await api.devices.list(appId, filters)

      dispatch(devices.getDevicesListSuccess(data.devices, data.totalCount))
    } catch (error) {
      dispatch(devices.getDevicesListFailure(error))
    }

    done()

  },
})

export default [
  getDevicesListLogic,
]
