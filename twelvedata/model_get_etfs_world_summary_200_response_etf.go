/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorldSummary200ResponseEtf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorldSummary200ResponseEtf{}

// GetETFsWorldSummary200ResponseEtf Etf information
type GetETFsWorldSummary200ResponseEtf struct {
	Summary *GetETFsWorldSummary200ResponseEtfSummary `json:"summary,omitempty"`
}

// NewGetETFsWorldSummary200ResponseEtf instantiates a new GetETFsWorldSummary200ResponseEtf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorldSummary200ResponseEtf() *GetETFsWorldSummary200ResponseEtf {
	this := GetETFsWorldSummary200ResponseEtf{}
	return &this
}

// NewGetETFsWorldSummary200ResponseEtfWithDefaults instantiates a new GetETFsWorldSummary200ResponseEtf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorldSummary200ResponseEtfWithDefaults() *GetETFsWorldSummary200ResponseEtf {
	this := GetETFsWorldSummary200ResponseEtf{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *GetETFsWorldSummary200ResponseEtf) GetSummary() GetETFsWorldSummary200ResponseEtfSummary {
	if o == nil || IsNil(o.Summary) {
		var ret GetETFsWorldSummary200ResponseEtfSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorldSummary200ResponseEtf) GetSummaryOk() (*GetETFsWorldSummary200ResponseEtfSummary, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *GetETFsWorldSummary200ResponseEtf) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given GetETFsWorldSummary200ResponseEtfSummary and assigns it to the Summary field.
func (o *GetETFsWorldSummary200ResponseEtf) SetSummary(v GetETFsWorldSummary200ResponseEtfSummary) {
	o.Summary = &v
}

func (o GetETFsWorldSummary200ResponseEtf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorldSummary200ResponseEtf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	return toSerialize, nil
}

type NullableGetETFsWorldSummary200ResponseEtf struct {
	value *GetETFsWorldSummary200ResponseEtf
	isSet bool
}

func (v NullableGetETFsWorldSummary200ResponseEtf) Get() *GetETFsWorldSummary200ResponseEtf {
	return v.value
}

func (v *NullableGetETFsWorldSummary200ResponseEtf) Set(val *GetETFsWorldSummary200ResponseEtf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorldSummary200ResponseEtf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorldSummary200ResponseEtf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorldSummary200ResponseEtf(val *GetETFsWorldSummary200ResponseEtf) *NullableGetETFsWorldSummary200ResponseEtf {
	return &NullableGetETFsWorldSummary200ResponseEtf{value: val, isSet: true}
}

func (v NullableGetETFsWorldSummary200ResponseEtf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorldSummary200ResponseEtf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
