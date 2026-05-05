// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration{}

// GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration Average duration of bond holding of a fund
type GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration struct {
	// Average duration of bond holding of a fund
	Fund *float64 `json:"fund,omitempty"`
	// Average duration of bond holding of funds in the same category
	Category *float64 `json:"category,omitempty"`
}

// NewGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration instantiates a new GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration() *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration {
	this := GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration{}
	return &this
}

// NewGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDurationWithDefaults instantiates a new GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDurationWithDefaults() *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration {
	this := GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration{}
	return &this
}

// GetFund returns the Fund field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) GetFund() float64 {
	if o == nil || IsNil(o.Fund) {
		var ret float64
		return ret
	}
	return *o.Fund
}

// GetFundOk returns a tuple with the Fund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) GetFundOk() (*float64, bool) {
	if o == nil || IsNil(o.Fund) {
		return nil, false
	}
	return o.Fund, true
}

// HasFund returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) HasFund() bool {
	if o != nil && !IsNil(o.Fund) {
		return true
	}

	return false
}

// SetFund gets a reference to the given float64 and assigns it to the Fund field.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) SetFund(v float64) {
	o.Fund = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) GetCategory() float64 {
	if o == nil || IsNil(o.Category) {
		var ret float64
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) GetCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given float64 and assigns it to the Category field.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) SetCategory(v float64) {
	o.Category = &v
}

func (o GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fund) {
		toSerialize["fund"] = o.Fund
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration struct {
	value *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) Get() *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) Set(val *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration(val *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) *NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration {
	return &NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
