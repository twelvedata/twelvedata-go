// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the ResponseMutualFundWorldRatings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMutualFundWorldRatings{}

// ResponseMutualFundWorldRatings Ratings of a mutual fund
type ResponseMutualFundWorldRatings struct {
	// Performance rating from 0 to 5
	PerformanceRating *int64 `json:"performance_rating,omitempty"`
	// Risk rating from 0 to 5
	RiskRating *int64 `json:"risk_rating,omitempty"`
	// Return rating from 0 to 5
	ReturnRating *int64 `json:"return_rating,omitempty"`
}

// NewResponseMutualFundWorldRatings instantiates a new ResponseMutualFundWorldRatings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMutualFundWorldRatings() *ResponseMutualFundWorldRatings {
	this := ResponseMutualFundWorldRatings{}
	return &this
}

// NewResponseMutualFundWorldRatingsWithDefaults instantiates a new ResponseMutualFundWorldRatings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMutualFundWorldRatingsWithDefaults() *ResponseMutualFundWorldRatings {
	this := ResponseMutualFundWorldRatings{}
	return &this
}

// GetPerformanceRating returns the PerformanceRating field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldRatings) GetPerformanceRating() int64 {
	if o == nil || IsNil(o.PerformanceRating) {
		var ret int64
		return ret
	}
	return *o.PerformanceRating
}

// GetPerformanceRatingOk returns a tuple with the PerformanceRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldRatings) GetPerformanceRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.PerformanceRating) {
		return nil, false
	}
	return o.PerformanceRating, true
}

// HasPerformanceRating returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldRatings) HasPerformanceRating() bool {
	if o != nil && !IsNil(o.PerformanceRating) {
		return true
	}

	return false
}

// SetPerformanceRating gets a reference to the given int64 and assigns it to the PerformanceRating field.
func (o *ResponseMutualFundWorldRatings) SetPerformanceRating(v int64) {
	o.PerformanceRating = &v
}

// GetRiskRating returns the RiskRating field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldRatings) GetRiskRating() int64 {
	if o == nil || IsNil(o.RiskRating) {
		var ret int64
		return ret
	}
	return *o.RiskRating
}

// GetRiskRatingOk returns a tuple with the RiskRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldRatings) GetRiskRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.RiskRating) {
		return nil, false
	}
	return o.RiskRating, true
}

// HasRiskRating returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldRatings) HasRiskRating() bool {
	if o != nil && !IsNil(o.RiskRating) {
		return true
	}

	return false
}

// SetRiskRating gets a reference to the given int64 and assigns it to the RiskRating field.
func (o *ResponseMutualFundWorldRatings) SetRiskRating(v int64) {
	o.RiskRating = &v
}

// GetReturnRating returns the ReturnRating field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldRatings) GetReturnRating() int64 {
	if o == nil || IsNil(o.ReturnRating) {
		var ret int64
		return ret
	}
	return *o.ReturnRating
}

// GetReturnRatingOk returns a tuple with the ReturnRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldRatings) GetReturnRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.ReturnRating) {
		return nil, false
	}
	return o.ReturnRating, true
}

// HasReturnRating returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldRatings) HasReturnRating() bool {
	if o != nil && !IsNil(o.ReturnRating) {
		return true
	}

	return false
}

// SetReturnRating gets a reference to the given int64 and assigns it to the ReturnRating field.
func (o *ResponseMutualFundWorldRatings) SetReturnRating(v int64) {
	o.ReturnRating = &v
}

func (o ResponseMutualFundWorldRatings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMutualFundWorldRatings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PerformanceRating) {
		toSerialize["performance_rating"] = o.PerformanceRating
	}
	if !IsNil(o.RiskRating) {
		toSerialize["risk_rating"] = o.RiskRating
	}
	if !IsNil(o.ReturnRating) {
		toSerialize["return_rating"] = o.ReturnRating
	}
	return toSerialize, nil
}

type NullableResponseMutualFundWorldRatings struct {
	value *ResponseMutualFundWorldRatings
	isSet bool
}

func (v NullableResponseMutualFundWorldRatings) Get() *ResponseMutualFundWorldRatings {
	return v.value
}

func (v *NullableResponseMutualFundWorldRatings) Set(val *ResponseMutualFundWorldRatings) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMutualFundWorldRatings) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMutualFundWorldRatings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMutualFundWorldRatings(val *ResponseMutualFundWorldRatings) *NullableResponseMutualFundWorldRatings {
	return &NullableResponseMutualFundWorldRatings{value: val, isSet: true}
}

func (v NullableResponseMutualFundWorldRatings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMutualFundWorldRatings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
