// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundRatings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundRatings{}

// GetMutualFundsWorld200ResponseMutualFundRatings Ratings of a mutual fund
type GetMutualFundsWorld200ResponseMutualFundRatings struct {
	// Performance rating from 0 to 5
	PerformanceRating *int64 `json:"performance_rating,omitempty"`
	// Risk rating from 0 to 5
	RiskRating *int64 `json:"risk_rating,omitempty"`
	// Return rating from 0 to 5
	ReturnRating *int64 `json:"return_rating,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundRatings instantiates a new GetMutualFundsWorld200ResponseMutualFundRatings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundRatings() *GetMutualFundsWorld200ResponseMutualFundRatings {
	this := GetMutualFundsWorld200ResponseMutualFundRatings{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundRatingsWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundRatings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundRatingsWithDefaults() *GetMutualFundsWorld200ResponseMutualFundRatings {
	this := GetMutualFundsWorld200ResponseMutualFundRatings{}
	return &this
}

// GetPerformanceRating returns the PerformanceRating field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRatings) GetPerformanceRating() int64 {
	if o == nil || IsNil(o.PerformanceRating) {
		var ret int64
		return ret
	}
	return *o.PerformanceRating
}

// GetPerformanceRatingOk returns a tuple with the PerformanceRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRatings) GetPerformanceRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.PerformanceRating) {
		return nil, false
	}
	return o.PerformanceRating, true
}

// HasPerformanceRating returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRatings) HasPerformanceRating() bool {
	if o != nil && !IsNil(o.PerformanceRating) {
		return true
	}

	return false
}

// SetPerformanceRating gets a reference to the given int64 and assigns it to the PerformanceRating field.
func (o *GetMutualFundsWorld200ResponseMutualFundRatings) SetPerformanceRating(v int64) {
	o.PerformanceRating = &v
}

// GetRiskRating returns the RiskRating field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRatings) GetRiskRating() int64 {
	if o == nil || IsNil(o.RiskRating) {
		var ret int64
		return ret
	}
	return *o.RiskRating
}

// GetRiskRatingOk returns a tuple with the RiskRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRatings) GetRiskRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.RiskRating) {
		return nil, false
	}
	return o.RiskRating, true
}

// HasRiskRating returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRatings) HasRiskRating() bool {
	if o != nil && !IsNil(o.RiskRating) {
		return true
	}

	return false
}

// SetRiskRating gets a reference to the given int64 and assigns it to the RiskRating field.
func (o *GetMutualFundsWorld200ResponseMutualFundRatings) SetRiskRating(v int64) {
	o.RiskRating = &v
}

// GetReturnRating returns the ReturnRating field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRatings) GetReturnRating() int64 {
	if o == nil || IsNil(o.ReturnRating) {
		var ret int64
		return ret
	}
	return *o.ReturnRating
}

// GetReturnRatingOk returns a tuple with the ReturnRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRatings) GetReturnRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.ReturnRating) {
		return nil, false
	}
	return o.ReturnRating, true
}

// HasReturnRating returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRatings) HasReturnRating() bool {
	if o != nil && !IsNil(o.ReturnRating) {
		return true
	}

	return false
}

// SetReturnRating gets a reference to the given int64 and assigns it to the ReturnRating field.
func (o *GetMutualFundsWorld200ResponseMutualFundRatings) SetReturnRating(v int64) {
	o.ReturnRating = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundRatings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundRatings) ToMap() (map[string]interface{}, error) {
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

type NullableGetMutualFundsWorld200ResponseMutualFundRatings struct {
	value *GetMutualFundsWorld200ResponseMutualFundRatings
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRatings) Get() *GetMutualFundsWorld200ResponseMutualFundRatings {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRatings) Set(val *GetMutualFundsWorld200ResponseMutualFundRatings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRatings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRatings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundRatings(val *GetMutualFundsWorld200ResponseMutualFundRatings) *NullableGetMutualFundsWorld200ResponseMutualFundRatings {
	return &NullableGetMutualFundsWorld200ResponseMutualFundRatings{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRatings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRatings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
