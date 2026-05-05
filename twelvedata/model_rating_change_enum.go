// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
	"fmt"
)

// RatingChangeEnum the model 'RatingChangeEnum'
type RatingChangeEnum string

// List of RatingChangeEnum
const (
	RATINGCHANGEENUM_MAINTAINS  RatingChangeEnum = "Maintains"
	RATINGCHANGEENUM_UPGRADE    RatingChangeEnum = "Upgrade"
	RATINGCHANGEENUM_DOWNGRADE  RatingChangeEnum = "Downgrade"
	RATINGCHANGEENUM_INITIATES  RatingChangeEnum = "Initiates"
	RATINGCHANGEENUM_REITERATES RatingChangeEnum = "Reiterates"
)

// All allowed values of RatingChangeEnum enum
var AllowedRatingChangeEnumEnumValues = []RatingChangeEnum{
	"Maintains",
	"Upgrade",
	"Downgrade",
	"Initiates",
	"Reiterates",
}

func (v *RatingChangeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RatingChangeEnum(value)
	for _, existing := range AllowedRatingChangeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RatingChangeEnum", value)
}

// NewRatingChangeEnumFromValue returns a pointer to a valid RatingChangeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRatingChangeEnumFromValue(v string) (*RatingChangeEnum, error) {
	ev := RatingChangeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RatingChangeEnum: valid values are %v", v, AllowedRatingChangeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RatingChangeEnum) IsValid() bool {
	for _, existing := range AllowedRatingChangeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RatingChangeEnum value
func (v RatingChangeEnum) Ptr() *RatingChangeEnum {
	return &v
}

type NullableRatingChangeEnum struct {
	value *RatingChangeEnum
	isSet bool
}

func (v NullableRatingChangeEnum) Get() *RatingChangeEnum {
	return v.value
}

func (v *NullableRatingChangeEnum) Set(val *RatingChangeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRatingChangeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRatingChangeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatingChangeEnum(val *RatingChangeEnum) *NullableRatingChangeEnum {
	return &NullableRatingChangeEnum{value: val, isSet: true}
}

func (v NullableRatingChangeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatingChangeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
