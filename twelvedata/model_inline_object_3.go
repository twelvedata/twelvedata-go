/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the InlineObject3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject3{}

// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	PerformanceRating *int64 `json:"performance_rating,omitempty"`
	ReturnRating      *int64 `json:"return_rating,omitempty"`
	RiskRating        *int64 `json:"risk_rating,omitempty"`
}

// NewInlineObject3 instantiates a new InlineObject3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject3() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// NewInlineObject3WithDefaults instantiates a new InlineObject3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject3WithDefaults() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// GetPerformanceRating returns the PerformanceRating field value if set, zero value otherwise.
func (o *InlineObject3) GetPerformanceRating() int64 {
	if o == nil || IsNil(o.PerformanceRating) {
		var ret int64
		return ret
	}
	return *o.PerformanceRating
}

// GetPerformanceRatingOk returns a tuple with the PerformanceRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetPerformanceRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.PerformanceRating) {
		return nil, false
	}
	return o.PerformanceRating, true
}

// HasPerformanceRating returns a boolean if a field has been set.
func (o *InlineObject3) HasPerformanceRating() bool {
	if o != nil && !IsNil(o.PerformanceRating) {
		return true
	}

	return false
}

// SetPerformanceRating gets a reference to the given int64 and assigns it to the PerformanceRating field.
func (o *InlineObject3) SetPerformanceRating(v int64) {
	o.PerformanceRating = &v
}

// GetReturnRating returns the ReturnRating field value if set, zero value otherwise.
func (o *InlineObject3) GetReturnRating() int64 {
	if o == nil || IsNil(o.ReturnRating) {
		var ret int64
		return ret
	}
	return *o.ReturnRating
}

// GetReturnRatingOk returns a tuple with the ReturnRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetReturnRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.ReturnRating) {
		return nil, false
	}
	return o.ReturnRating, true
}

// HasReturnRating returns a boolean if a field has been set.
func (o *InlineObject3) HasReturnRating() bool {
	if o != nil && !IsNil(o.ReturnRating) {
		return true
	}

	return false
}

// SetReturnRating gets a reference to the given int64 and assigns it to the ReturnRating field.
func (o *InlineObject3) SetReturnRating(v int64) {
	o.ReturnRating = &v
}

// GetRiskRating returns the RiskRating field value if set, zero value otherwise.
func (o *InlineObject3) GetRiskRating() int64 {
	if o == nil || IsNil(o.RiskRating) {
		var ret int64
		return ret
	}
	return *o.RiskRating
}

// GetRiskRatingOk returns a tuple with the RiskRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetRiskRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.RiskRating) {
		return nil, false
	}
	return o.RiskRating, true
}

// HasRiskRating returns a boolean if a field has been set.
func (o *InlineObject3) HasRiskRating() bool {
	if o != nil && !IsNil(o.RiskRating) {
		return true
	}

	return false
}

// SetRiskRating gets a reference to the given int64 and assigns it to the RiskRating field.
func (o *InlineObject3) SetRiskRating(v int64) {
	o.RiskRating = &v
}

func (o InlineObject3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PerformanceRating) {
		toSerialize["performance_rating"] = o.PerformanceRating
	}
	if !IsNil(o.ReturnRating) {
		toSerialize["return_rating"] = o.ReturnRating
	}
	if !IsNil(o.RiskRating) {
		toSerialize["risk_rating"] = o.RiskRating
	}
	return toSerialize, nil
}

type NullableInlineObject3 struct {
	value *InlineObject3
	isSet bool
}

func (v NullableInlineObject3) Get() *InlineObject3 {
	return v.value
}

func (v *NullableInlineObject3) Set(val *InlineObject3) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject3) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject3(val *InlineObject3) *NullableInlineObject3 {
	return &NullableInlineObject3{value: val, isSet: true}
}

func (v NullableInlineObject3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
