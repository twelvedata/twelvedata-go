/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner{}

// GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner struct for GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner
type GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner struct {
	// Rating of bond holding of a fund from AAA to below B
	Grade *string `json:"grade,omitempty"`
	// Weight of bond holding in fund portfolio
	Weight *float64 `json:"weight,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner instantiates a new GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner() *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner {
	this := GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInnerWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInnerWithDefaults() *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner {
	this := GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner{}
	return &this
}

// GetGrade returns the Grade field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) GetGrade() string {
	if o == nil || IsNil(o.Grade) {
		var ret string
		return ret
	}
	return *o.Grade
}

// GetGradeOk returns a tuple with the Grade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) GetGradeOk() (*string, bool) {
	if o == nil || IsNil(o.Grade) {
		return nil, false
	}
	return o.Grade, true
}

// HasGrade returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) HasGrade() bool {
	if o != nil && !IsNil(o.Grade) {
		return true
	}

	return false
}

// SetGrade gets a reference to the given string and assigns it to the Grade field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) SetGrade(v string) {
	o.Grade = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) SetWeight(v float64) {
	o.Weight = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Grade) {
		toSerialize["grade"] = o.Grade
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner struct {
	value *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) Get() *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) Set(val *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner(val *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) *NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner {
	return &NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
