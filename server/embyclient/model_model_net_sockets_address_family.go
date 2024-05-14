/*
Emby Server REST API (BETA)

Explore the Emby Server API

API version: 4.8.0.61
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package embyclient

import (
	"encoding/json"
	"fmt"
)

// ModelNetSocketsAddressFamily the model 'ModelNetSocketsAddressFamily'
type ModelNetSocketsAddressFamily string

// List of Net.Sockets.AddressFamily
const (
	MODELNETSOCKETSADDRESSFAMILY_UNSPECIFIED ModelNetSocketsAddressFamily = "Unspecified"
	MODELNETSOCKETSADDRESSFAMILY_UNIX ModelNetSocketsAddressFamily = "Unix"
	MODELNETSOCKETSADDRESSFAMILY_INTER_NETWORK ModelNetSocketsAddressFamily = "InterNetwork"
	MODELNETSOCKETSADDRESSFAMILY_IMP_LINK ModelNetSocketsAddressFamily = "ImpLink"
	MODELNETSOCKETSADDRESSFAMILY_PUP ModelNetSocketsAddressFamily = "Pup"
	MODELNETSOCKETSADDRESSFAMILY_CHAOS ModelNetSocketsAddressFamily = "Chaos"
	MODELNETSOCKETSADDRESSFAMILY_NS ModelNetSocketsAddressFamily = "NS"
	MODELNETSOCKETSADDRESSFAMILY_IPX ModelNetSocketsAddressFamily = "Ipx"
	MODELNETSOCKETSADDRESSFAMILY_OSI ModelNetSocketsAddressFamily = "Osi"
	MODELNETSOCKETSADDRESSFAMILY_ISO ModelNetSocketsAddressFamily = "Iso"
	MODELNETSOCKETSADDRESSFAMILY_ECMA ModelNetSocketsAddressFamily = "Ecma"
	MODELNETSOCKETSADDRESSFAMILY_DATA_KIT ModelNetSocketsAddressFamily = "DataKit"
	MODELNETSOCKETSADDRESSFAMILY_CCITT ModelNetSocketsAddressFamily = "Ccitt"
	MODELNETSOCKETSADDRESSFAMILY_SNA ModelNetSocketsAddressFamily = "Sna"
	MODELNETSOCKETSADDRESSFAMILY_DEC_NET ModelNetSocketsAddressFamily = "DecNet"
	MODELNETSOCKETSADDRESSFAMILY_DATA_LINK ModelNetSocketsAddressFamily = "DataLink"
	MODELNETSOCKETSADDRESSFAMILY_LAT ModelNetSocketsAddressFamily = "Lat"
	MODELNETSOCKETSADDRESSFAMILY_HYPER_CHANNEL ModelNetSocketsAddressFamily = "HyperChannel"
	MODELNETSOCKETSADDRESSFAMILY_APPLE_TALK ModelNetSocketsAddressFamily = "AppleTalk"
	MODELNETSOCKETSADDRESSFAMILY_NET_BIOS ModelNetSocketsAddressFamily = "NetBios"
	MODELNETSOCKETSADDRESSFAMILY_VOICE_VIEW ModelNetSocketsAddressFamily = "VoiceView"
	MODELNETSOCKETSADDRESSFAMILY_FIRE_FOX ModelNetSocketsAddressFamily = "FireFox"
	MODELNETSOCKETSADDRESSFAMILY_BANYAN ModelNetSocketsAddressFamily = "Banyan"
	MODELNETSOCKETSADDRESSFAMILY_ATM ModelNetSocketsAddressFamily = "Atm"
	MODELNETSOCKETSADDRESSFAMILY_INTER_NETWORK_V6 ModelNetSocketsAddressFamily = "InterNetworkV6"
	MODELNETSOCKETSADDRESSFAMILY_CLUSTER ModelNetSocketsAddressFamily = "Cluster"
	MODELNETSOCKETSADDRESSFAMILY_IEEE12844 ModelNetSocketsAddressFamily = "Ieee12844"
	MODELNETSOCKETSADDRESSFAMILY_IRDA ModelNetSocketsAddressFamily = "Irda"
	MODELNETSOCKETSADDRESSFAMILY_NETWORK_DESIGNERS ModelNetSocketsAddressFamily = "NetworkDesigners"
	MODELNETSOCKETSADDRESSFAMILY_MAX ModelNetSocketsAddressFamily = "Max"
	MODELNETSOCKETSADDRESSFAMILY_PACKET ModelNetSocketsAddressFamily = "Packet"
	MODELNETSOCKETSADDRESSFAMILY_CONTROLLER_AREA_NETWORK ModelNetSocketsAddressFamily = "ControllerAreaNetwork"
	MODELNETSOCKETSADDRESSFAMILY_UNKNOWN ModelNetSocketsAddressFamily = "Unknown"
)

// All allowed values of ModelNetSocketsAddressFamily enum
var AllowedModelNetSocketsAddressFamilyEnumValues = []ModelNetSocketsAddressFamily{
	"Unspecified",
	"Unix",
	"InterNetwork",
	"ImpLink",
	"Pup",
	"Chaos",
	"NS",
	"Ipx",
	"Osi",
	"Iso",
	"Ecma",
	"DataKit",
	"Ccitt",
	"Sna",
	"DecNet",
	"DataLink",
	"Lat",
	"HyperChannel",
	"AppleTalk",
	"NetBios",
	"VoiceView",
	"FireFox",
	"Banyan",
	"Atm",
	"InterNetworkV6",
	"Cluster",
	"Ieee12844",
	"Irda",
	"NetworkDesigners",
	"Max",
	"Packet",
	"ControllerAreaNetwork",
	"Unknown",
}

func (v *ModelNetSocketsAddressFamily) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelNetSocketsAddressFamily(value)
	for _, existing := range AllowedModelNetSocketsAddressFamilyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelNetSocketsAddressFamily", value)
}

// NewModelNetSocketsAddressFamilyFromValue returns a pointer to a valid ModelNetSocketsAddressFamily
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelNetSocketsAddressFamilyFromValue(v string) (*ModelNetSocketsAddressFamily, error) {
	ev := ModelNetSocketsAddressFamily(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelNetSocketsAddressFamily: valid values are %v", v, AllowedModelNetSocketsAddressFamilyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelNetSocketsAddressFamily) IsValid() bool {
	for _, existing := range AllowedModelNetSocketsAddressFamilyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Net.Sockets.AddressFamily value
func (v ModelNetSocketsAddressFamily) Ptr() *ModelNetSocketsAddressFamily {
	return &v
}

type NullableModelNetSocketsAddressFamily struct {
	value *ModelNetSocketsAddressFamily
	isSet bool
}

func (v NullableModelNetSocketsAddressFamily) Get() *ModelNetSocketsAddressFamily {
	return v.value
}

func (v *NullableModelNetSocketsAddressFamily) Set(val *ModelNetSocketsAddressFamily) {
	v.value = val
	v.isSet = true
}

func (v NullableModelNetSocketsAddressFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableModelNetSocketsAddressFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelNetSocketsAddressFamily(val *ModelNetSocketsAddressFamily) *NullableModelNetSocketsAddressFamily {
	return &NullableModelNetSocketsAddressFamily{value: val, isSet: true}
}

func (v NullableModelNetSocketsAddressFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelNetSocketsAddressFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
