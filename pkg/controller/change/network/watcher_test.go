// Copyright 2019-present Open Networking Foundation.
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

package network

import (
	"github.com/golang/mock/gomock"
	"github.com/onosproject/onos-config/api/types"
	changetypes "github.com/onosproject/onos-config/api/types/change"
	devicechange "github.com/onosproject/onos-config/api/types/change/device"
	networkchange "github.com/onosproject/onos-config/api/types/change/network"
	devicechangestore "github.com/onosproject/onos-config/pkg/store/change/device"
	networkchangestore "github.com/onosproject/onos-config/pkg/store/change/network"
	devicestore "github.com/onosproject/onos-config/pkg/store/device"
	topodevice "github.com/onosproject/onos-topo/api/device"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

func TestNetworkWatcher(t *testing.T) {
	store, err := networkchangestore.NewLocalStore()
	assert.NoError(t, err)
	defer store.Close()

	watcher := &Watcher{
		Store: store,
	}

	ch := make(chan types.ID)
	err = watcher.Start(ch)
	assert.NoError(t, err)

	change1 := &networkchange.NetworkChange{
		ID: "change-1",
		Changes: []*devicechange.Change{
			{

				DeviceID:      "device-1",
				DeviceVersion: "1.0.0",
				Values: []*devicechange.ChangeValue{
					{
						Path: "foo",
						Value: &devicechange.TypedValue{
							Bytes: []byte("Hello world!"),
							Type:  devicechange.ValueType_STRING,
						},
					},
					{
						Path: "bar",
						Value: &devicechange.TypedValue{
							Bytes: []byte("Hello world again!"),
							Type:  devicechange.ValueType_STRING,
						},
					},
				},
			},
			{
				DeviceID:      "device-2",
				DeviceVersion: "1.0.0",
				Values: []*devicechange.ChangeValue{
					{
						Path: "baz",
						Value: &devicechange.TypedValue{
							Bytes: []byte("Goodbye world!"),
							Type:  devicechange.ValueType_STRING,
						},
					},
				},
			},
		},
	}

	err = store.Create(change1)
	assert.NoError(t, err)

	select {
	case id := <-ch:
		assert.Equal(t, change1.ID, networkchange.ID(id))
	case <-time.After(5 * time.Second):
		t.FailNow()
	}

	change2 := &networkchange.NetworkChange{
		ID: "change-2",
		Changes: []*devicechange.Change{
			{
				DeviceID:      "device-1",
				DeviceVersion: "1.0.0",
				Values: []*devicechange.ChangeValue{
					{
						Path:    "foo",
						Removed: true,
					},
				},
			},
		},
	}

	err = store.Create(change2)
	assert.NoError(t, err)

	select {
	case <-ch:
	case <-time.After(5 * time.Second):
		t.FailNow()
	}

	change1.Status.State = changetypes.State_RUNNING
	err = store.Update(change1)
	assert.NoError(t, err)

	select {
	case <-ch:
	case <-time.After(5 * time.Second):
		t.FailNow()
	}
}

func TestDeviceWatcher(t *testing.T) {
	ctrl := gomock.NewController(t)

	stream := NewMockDeviceService_ListClient(ctrl)
	stream.EXPECT().Recv().Return(&topodevice.ListResponse{Device: &topodevice.Device{ID: topodevice.ID("device-1"), Version: "1.0.0"}}, nil)
	stream.EXPECT().Recv().Return(&topodevice.ListResponse{Device: &topodevice.Device{ID: topodevice.ID("device-2"), Version: "1.0.0"}}, nil)
	stream.EXPECT().Recv().Return(nil, io.EOF)

	client := NewMockDeviceServiceClient(ctrl)
	client.EXPECT().List(gomock.Any(), gomock.Any()).Return(stream, nil).AnyTimes()

	deviceStore, err := devicestore.NewStore(client)
	assert.NoError(t, err)

	changeStore, err := devicechangestore.NewLocalStore()
	assert.NoError(t, err)
	defer changeStore.Close()

	watcher := &DeviceWatcher{
		DeviceStore: deviceStore,
		ChangeStore: changeStore,
	}

	ch := make(chan types.ID)
	err = watcher.Start(ch)
	assert.NoError(t, err)

	change1 := &devicechange.DeviceChange{
		NetworkChange: devicechange.NetworkChangeRef{
			ID:    "network-change-1",
			Index: 2,
		},
		Change: &devicechange.Change{
			DeviceID:      "device-1",
			DeviceVersion: "1.0.0",
			DeviceType:    "Stratum",
			Values: []*devicechange.ChangeValue{
				{
					Path: "foo",
					Value: &devicechange.TypedValue{
						Bytes: []byte("Hello world!"),
						Type:  devicechange.ValueType_STRING,
					},
				},
				{
					Path: "bar",
					Value: &devicechange.TypedValue{
						Bytes: []byte("Hello world again!"),
						Type:  devicechange.ValueType_STRING,
					},
				},
			},
		},
	}

	err = changeStore.Create(change1)
	assert.NoError(t, err)

	select {
	case id := <-ch:
		assert.Equal(t, change1.NetworkChange.ID, id)
	case <-time.After(5 * time.Second):
		t.FailNow()
	}

	change2 := &devicechange.DeviceChange{
		NetworkChange: devicechange.NetworkChangeRef{
			ID:    "network-change-2",
			Index: 2,
		},
		Change: &devicechange.Change{
			DeviceID:      "device-2",
			DeviceVersion: "1.0.0",
			DeviceType:    "Stratum",
			Values: []*devicechange.ChangeValue{
				{
					Path: "baz",
					Value: &devicechange.TypedValue{
						Bytes: []byte("Goodbye world!"),
						Type:  devicechange.ValueType_STRING,
					},
				},
			},
		},
	}

	err = changeStore.Create(change2)
	assert.NoError(t, err)

	select {
	case <-ch:
	case <-time.After(5 * time.Second):
		t.FailNow()
	}

	change1.Status.State = changetypes.State_RUNNING
	err = changeStore.Update(change1)
	assert.NoError(t, err)

	select {
	case <-ch:
	case <-time.After(5 * time.Second):
		t.FailNow()
	}
}
