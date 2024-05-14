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

// ModelRecommendationType the model 'ModelRecommendationType'
type ModelRecommendationType string

// List of RecommendationType
const (
	MODELRECOMMENDATIONTYPE_SIMILAR_TO_RECENTLY_PLAYED ModelRecommendationType = "SimilarToRecentlyPlayed"
	MODELRECOMMENDATIONTYPE_SIMILAR_TO_LIKED_ITEM ModelRecommendationType = "SimilarToLikedItem"
	MODELRECOMMENDATIONTYPE_HAS_DIRECTOR_FROM_RECENTLY_PLAYED ModelRecommendationType = "HasDirectorFromRecentlyPlayed"
	MODELRECOMMENDATIONTYPE_HAS_ACTOR_FROM_RECENTLY_PLAYED ModelRecommendationType = "HasActorFromRecentlyPlayed"
	MODELRECOMMENDATIONTYPE_HAS_LIKED_DIRECTOR ModelRecommendationType = "HasLikedDirector"
	MODELRECOMMENDATIONTYPE_HAS_LIKED_ACTOR ModelRecommendationType = "HasLikedActor"
)

// All allowed values of ModelRecommendationType enum
var AllowedModelRecommendationTypeEnumValues = []ModelRecommendationType{
	"SimilarToRecentlyPlayed",
	"SimilarToLikedItem",
	"HasDirectorFromRecentlyPlayed",
	"HasActorFromRecentlyPlayed",
	"HasLikedDirector",
	"HasLikedActor",
}

func (v *ModelRecommendationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelRecommendationType(value)
	for _, existing := range AllowedModelRecommendationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelRecommendationType", value)
}

// NewModelRecommendationTypeFromValue returns a pointer to a valid ModelRecommendationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelRecommendationTypeFromValue(v string) (*ModelRecommendationType, error) {
	ev := ModelRecommendationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelRecommendationType: valid values are %v", v, AllowedModelRecommendationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelRecommendationType) IsValid() bool {
	for _, existing := range AllowedModelRecommendationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecommendationType value
func (v ModelRecommendationType) Ptr() *ModelRecommendationType {
	return &v
}

type NullableModelRecommendationType struct {
	value *ModelRecommendationType
	isSet bool
}

func (v NullableModelRecommendationType) Get() *ModelRecommendationType {
	return v.value
}

func (v *NullableModelRecommendationType) Set(val *ModelRecommendationType) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRecommendationType) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRecommendationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRecommendationType(val *ModelRecommendationType) *NullableModelRecommendationType {
	return &NullableModelRecommendationType{value: val, isSet: true}
}

func (v NullableModelRecommendationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRecommendationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
