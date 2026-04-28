/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorldRisk200ResponseEtf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorldRisk200ResponseEtf{}

// GetETFsWorldRisk200ResponseEtf Etf information
type GetETFsWorldRisk200ResponseEtf struct {
	Risk *GetETFsWorldRisk200ResponseEtfRisk `json:"risk,omitempty"`
}

// NewGetETFsWorldRisk200ResponseEtf instantiates a new GetETFsWorldRisk200ResponseEtf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorldRisk200ResponseEtf() *GetETFsWorldRisk200ResponseEtf {
	this := GetETFsWorldRisk200ResponseEtf{}
	return &this
}

// NewGetETFsWorldRisk200ResponseEtfWithDefaults instantiates a new GetETFsWorldRisk200ResponseEtf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorldRisk200ResponseEtfWithDefaults() *GetETFsWorldRisk200ResponseEtf {
	this := GetETFsWorldRisk200ResponseEtf{}
	return &this
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *GetETFsWorldRisk200ResponseEtf) GetRisk() GetETFsWorldRisk200ResponseEtfRisk {
	if o == nil || IsNil(o.Risk) {
		var ret GetETFsWorldRisk200ResponseEtfRisk
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorldRisk200ResponseEtf) GetRiskOk() (*GetETFsWorldRisk200ResponseEtfRisk, bool) {
	if o == nil || IsNil(o.Risk) {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *GetETFsWorldRisk200ResponseEtf) HasRisk() bool {
	if o != nil && !IsNil(o.Risk) {
		return true
	}

	return false
}

// SetRisk gets a reference to the given GetETFsWorldRisk200ResponseEtfRisk and assigns it to the Risk field.
func (o *GetETFsWorldRisk200ResponseEtf) SetRisk(v GetETFsWorldRisk200ResponseEtfRisk) {
	o.Risk = &v
}

func (o GetETFsWorldRisk200ResponseEtf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorldRisk200ResponseEtf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Risk) {
		toSerialize["risk"] = o.Risk
	}
	return toSerialize, nil
}

type NullableGetETFsWorldRisk200ResponseEtf struct {
	value *GetETFsWorldRisk200ResponseEtf
	isSet bool
}

func (v NullableGetETFsWorldRisk200ResponseEtf) Get() *GetETFsWorldRisk200ResponseEtf {
	return v.value
}

func (v *NullableGetETFsWorldRisk200ResponseEtf) Set(val *GetETFsWorldRisk200ResponseEtf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorldRisk200ResponseEtf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorldRisk200ResponseEtf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorldRisk200ResponseEtf(val *GetETFsWorldRisk200ResponseEtf) *NullableGetETFsWorldRisk200ResponseEtf {
	return &NullableGetETFsWorldRisk200ResponseEtf{value: val, isSet: true}
}

func (v NullableGetETFsWorldRisk200ResponseEtf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorldRisk200ResponseEtf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
