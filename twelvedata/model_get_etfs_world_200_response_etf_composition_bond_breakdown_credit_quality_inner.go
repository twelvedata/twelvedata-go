/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner{}

// GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner struct for GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner
type GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner struct {
	// Rating of bond holding of a fund from AAA to below B
	Grade *string `json:"grade,omitempty"`
	// Weight of bond holding in fund portfolio
	Weight *float64 `json:"weight,omitempty"`
}

// NewGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner instantiates a new GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner() *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner {
	this := GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner{}
	return &this
}

// NewGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInnerWithDefaults instantiates a new GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInnerWithDefaults() *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner {
	this := GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner{}
	return &this
}

// GetGrade returns the Grade field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) GetGrade() string {
	if o == nil || IsNil(o.Grade) {
		var ret string
		return ret
	}
	return *o.Grade
}

// GetGradeOk returns a tuple with the Grade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) GetGradeOk() (*string, bool) {
	if o == nil || IsNil(o.Grade) {
		return nil, false
	}
	return o.Grade, true
}

// HasGrade returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) HasGrade() bool {
	if o != nil && !IsNil(o.Grade) {
		return true
	}

	return false
}

// SetGrade gets a reference to the given string and assigns it to the Grade field.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) SetGrade(v string) {
	o.Grade = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) SetWeight(v float64) {
	o.Weight = &v
}

func (o GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Grade) {
		toSerialize["grade"] = o.Grade
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner struct {
	value *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) Get() *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) Set(val *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner(val *GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) *NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner {
	return &NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
