//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.

package notifications

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

const (
	KDeviceConnected       = "DEVICE_CONNECTED"
	KDeviceDisconnected    = "DEVICE_DISCONNECTED"
	KDeviceMessageReceived = "DEVICE_MESSAGE_RECEIVED"
)

// DeviceNotification DeviceNotification
type DeviceNotification struct {
	UserID     string `json:"userID"`
	DeviceID   string `json:"deviceID"`
	DeviceName string `json:"deviceName"`
	Type       string `json:"type"`
	Payload    []byte `json:"payload"`
}

// Validate validates this deployment
func (m *DeviceNotification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceNotification) UnmarshalBinary(b []byte) error {
	var res DeviceNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
