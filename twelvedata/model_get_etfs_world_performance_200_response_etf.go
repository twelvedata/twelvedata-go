// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorldPerformance200ResponseEtf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorldPerformance200ResponseEtf{}

// GetETFsWorldPerformance200ResponseEtf Etf information
type GetETFsWorldPerformance200ResponseEtf struct {
	Performance *GetETFsWorldPerformance200ResponseEtfPerformance `json:"performance,omitempty"`
}

// NewGetETFsWorldPerformance200ResponseEtf instantiates a new GetETFsWorldPerformance200ResponseEtf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorldPerformance200ResponseEtf() *GetETFsWorldPerformance200ResponseEtf {
	this := GetETFsWorldPerformance200ResponseEtf{}
	return &this
}

// NewGetETFsWorldPerformance200ResponseEtfWithDefaults instantiates a new GetETFsWorldPerformance200ResponseEtf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorldPerformance200ResponseEtfWithDefaults() *GetETFsWorldPerformance200ResponseEtf {
	this := GetETFsWorldPerformance200ResponseEtf{}
	return &this
}

// GetPerformance returns the Performance field value if set, zero value otherwise.
func (o *GetETFsWorldPerformance200ResponseEtf) GetPerformance() GetETFsWorldPerformance200ResponseEtfPerformance {
	if o == nil || IsNil(o.Performance) {
		var ret GetETFsWorldPerformance200ResponseEtfPerformance
		return ret
	}
	return *o.Performance
}

// GetPerformanceOk returns a tuple with the Performance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorldPerformance200ResponseEtf) GetPerformanceOk() (*GetETFsWorldPerformance200ResponseEtfPerformance, bool) {
	if o == nil || IsNil(o.Performance) {
		return nil, false
	}
	return o.Performance, true
}

// HasPerformance returns a boolean if a field has been set.
func (o *GetETFsWorldPerformance200ResponseEtf) HasPerformance() bool {
	if o != nil && !IsNil(o.Performance) {
		return true
	}

	return false
}

// SetPerformance gets a reference to the given GetETFsWorldPerformance200ResponseEtfPerformance and assigns it to the Performance field.
func (o *GetETFsWorldPerformance200ResponseEtf) SetPerformance(v GetETFsWorldPerformance200ResponseEtfPerformance) {
	o.Performance = &v
}

func (o GetETFsWorldPerformance200ResponseEtf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorldPerformance200ResponseEtf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Performance) {
		toSerialize["performance"] = o.Performance
	}
	return toSerialize, nil
}

type NullableGetETFsWorldPerformance200ResponseEtf struct {
	value *GetETFsWorldPerformance200ResponseEtf
	isSet bool
}

func (v NullableGetETFsWorldPerformance200ResponseEtf) Get() *GetETFsWorldPerformance200ResponseEtf {
	return v.value
}

func (v *NullableGetETFsWorldPerformance200ResponseEtf) Set(val *GetETFsWorldPerformance200ResponseEtf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorldPerformance200ResponseEtf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorldPerformance200ResponseEtf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorldPerformance200ResponseEtf(val *GetETFsWorldPerformance200ResponseEtf) *NullableGetETFsWorldPerformance200ResponseEtf {
	return &NullableGetETFsWorldPerformance200ResponseEtf{value: val, isSet: true}
}

func (v NullableGetETFsWorldPerformance200ResponseEtf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorldPerformance200ResponseEtf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
