/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorldComposition200ResponseEtf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorldComposition200ResponseEtf{}

// GetETFsWorldComposition200ResponseEtf Etf information
type GetETFsWorldComposition200ResponseEtf struct {
	Composition *GetETFsWorldComposition200ResponseEtfComposition `json:"composition,omitempty"`
}

// NewGetETFsWorldComposition200ResponseEtf instantiates a new GetETFsWorldComposition200ResponseEtf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorldComposition200ResponseEtf() *GetETFsWorldComposition200ResponseEtf {
	this := GetETFsWorldComposition200ResponseEtf{}
	return &this
}

// NewGetETFsWorldComposition200ResponseEtfWithDefaults instantiates a new GetETFsWorldComposition200ResponseEtf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorldComposition200ResponseEtfWithDefaults() *GetETFsWorldComposition200ResponseEtf {
	this := GetETFsWorldComposition200ResponseEtf{}
	return &this
}

// GetComposition returns the Composition field value if set, zero value otherwise.
func (o *GetETFsWorldComposition200ResponseEtf) GetComposition() GetETFsWorldComposition200ResponseEtfComposition {
	if o == nil || IsNil(o.Composition) {
		var ret GetETFsWorldComposition200ResponseEtfComposition
		return ret
	}
	return *o.Composition
}

// GetCompositionOk returns a tuple with the Composition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorldComposition200ResponseEtf) GetCompositionOk() (*GetETFsWorldComposition200ResponseEtfComposition, bool) {
	if o == nil || IsNil(o.Composition) {
		return nil, false
	}
	return o.Composition, true
}

// HasComposition returns a boolean if a field has been set.
func (o *GetETFsWorldComposition200ResponseEtf) HasComposition() bool {
	if o != nil && !IsNil(o.Composition) {
		return true
	}

	return false
}

// SetComposition gets a reference to the given GetETFsWorldComposition200ResponseEtfComposition and assigns it to the Composition field.
func (o *GetETFsWorldComposition200ResponseEtf) SetComposition(v GetETFsWorldComposition200ResponseEtfComposition) {
	o.Composition = &v
}

func (o GetETFsWorldComposition200ResponseEtf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorldComposition200ResponseEtf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Composition) {
		toSerialize["composition"] = o.Composition
	}
	return toSerialize, nil
}

type NullableGetETFsWorldComposition200ResponseEtf struct {
	value *GetETFsWorldComposition200ResponseEtf
	isSet bool
}

func (v NullableGetETFsWorldComposition200ResponseEtf) Get() *GetETFsWorldComposition200ResponseEtf {
	return v.value
}

func (v *NullableGetETFsWorldComposition200ResponseEtf) Set(val *GetETFsWorldComposition200ResponseEtf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorldComposition200ResponseEtf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorldComposition200ResponseEtf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorldComposition200ResponseEtf(val *GetETFsWorldComposition200ResponseEtf) *NullableGetETFsWorldComposition200ResponseEtf {
	return &NullableGetETFsWorldComposition200ResponseEtf{value: val, isSet: true}
}

func (v NullableGetETFsWorldComposition200ResponseEtf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorldComposition200ResponseEtf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
