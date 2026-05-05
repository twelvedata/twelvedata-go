// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfCompositionCountryAllocationInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfCompositionCountryAllocationInner{}

// GetETFsWorld200ResponseEtfCompositionCountryAllocationInner struct for GetETFsWorld200ResponseEtfCompositionCountryAllocationInner
type GetETFsWorld200ResponseEtfCompositionCountryAllocationInner struct {
	// Country name
	Country *string `json:"country,omitempty"`
	// Percentages of a fund's net assets distributed to securities of the country
	Allocation *float64 `json:"allocation,omitempty"`
}

// NewGetETFsWorld200ResponseEtfCompositionCountryAllocationInner instantiates a new GetETFsWorld200ResponseEtfCompositionCountryAllocationInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfCompositionCountryAllocationInner() *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner {
	this := GetETFsWorld200ResponseEtfCompositionCountryAllocationInner{}
	return &this
}

// NewGetETFsWorld200ResponseEtfCompositionCountryAllocationInnerWithDefaults instantiates a new GetETFsWorld200ResponseEtfCompositionCountryAllocationInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfCompositionCountryAllocationInnerWithDefaults() *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner {
	this := GetETFsWorld200ResponseEtfCompositionCountryAllocationInner{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) SetCountry(v string) {
	o.Country = &v
}

// GetAllocation returns the Allocation field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) GetAllocation() float64 {
	if o == nil || IsNil(o.Allocation) {
		var ret float64
		return ret
	}
	return *o.Allocation
}

// GetAllocationOk returns a tuple with the Allocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) GetAllocationOk() (*float64, bool) {
	if o == nil || IsNil(o.Allocation) {
		return nil, false
	}
	return o.Allocation, true
}

// HasAllocation returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) HasAllocation() bool {
	if o != nil && !IsNil(o.Allocation) {
		return true
	}

	return false
}

// SetAllocation gets a reference to the given float64 and assigns it to the Allocation field.
func (o *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) SetAllocation(v float64) {
	o.Allocation = &v
}

func (o GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Allocation) {
		toSerialize["allocation"] = o.Allocation
	}
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtfCompositionCountryAllocationInner struct {
	value *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfCompositionCountryAllocationInner) Get() *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionCountryAllocationInner) Set(val *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfCompositionCountryAllocationInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionCountryAllocationInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfCompositionCountryAllocationInner(val *GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) *NullableGetETFsWorld200ResponseEtfCompositionCountryAllocationInner {
	return &NullableGetETFsWorld200ResponseEtfCompositionCountryAllocationInner{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfCompositionCountryAllocationInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionCountryAllocationInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
