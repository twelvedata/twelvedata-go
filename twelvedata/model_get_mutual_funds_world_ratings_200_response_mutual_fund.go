// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorldRatings200ResponseMutualFund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldRatings200ResponseMutualFund{}

// GetMutualFundsWorldRatings200ResponseMutualFund Mutual fund information
type GetMutualFundsWorldRatings200ResponseMutualFund struct {
	Ratings *ResponseMutualFundWorldRatings `json:"ratings,omitempty"`
}

// NewGetMutualFundsWorldRatings200ResponseMutualFund instantiates a new GetMutualFundsWorldRatings200ResponseMutualFund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldRatings200ResponseMutualFund() *GetMutualFundsWorldRatings200ResponseMutualFund {
	this := GetMutualFundsWorldRatings200ResponseMutualFund{}
	return &this
}

// NewGetMutualFundsWorldRatings200ResponseMutualFundWithDefaults instantiates a new GetMutualFundsWorldRatings200ResponseMutualFund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldRatings200ResponseMutualFundWithDefaults() *GetMutualFundsWorldRatings200ResponseMutualFund {
	this := GetMutualFundsWorldRatings200ResponseMutualFund{}
	return &this
}

// GetRatings returns the Ratings field value if set, zero value otherwise.
func (o *GetMutualFundsWorldRatings200ResponseMutualFund) GetRatings() ResponseMutualFundWorldRatings {
	if o == nil || IsNil(o.Ratings) {
		var ret ResponseMutualFundWorldRatings
		return ret
	}
	return *o.Ratings
}

// GetRatingsOk returns a tuple with the Ratings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldRatings200ResponseMutualFund) GetRatingsOk() (*ResponseMutualFundWorldRatings, bool) {
	if o == nil || IsNil(o.Ratings) {
		return nil, false
	}
	return o.Ratings, true
}

// HasRatings returns a boolean if a field has been set.
func (o *GetMutualFundsWorldRatings200ResponseMutualFund) HasRatings() bool {
	if o != nil && !IsNil(o.Ratings) {
		return true
	}

	return false
}

// SetRatings gets a reference to the given ResponseMutualFundWorldRatings and assigns it to the Ratings field.
func (o *GetMutualFundsWorldRatings200ResponseMutualFund) SetRatings(v ResponseMutualFundWorldRatings) {
	o.Ratings = &v
}

func (o GetMutualFundsWorldRatings200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldRatings200ResponseMutualFund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ratings) {
		toSerialize["ratings"] = o.Ratings
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorldRatings200ResponseMutualFund struct {
	value *GetMutualFundsWorldRatings200ResponseMutualFund
	isSet bool
}

func (v NullableGetMutualFundsWorldRatings200ResponseMutualFund) Get() *GetMutualFundsWorldRatings200ResponseMutualFund {
	return v.value
}

func (v *NullableGetMutualFundsWorldRatings200ResponseMutualFund) Set(val *GetMutualFundsWorldRatings200ResponseMutualFund) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldRatings200ResponseMutualFund) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldRatings200ResponseMutualFund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldRatings200ResponseMutualFund(val *GetMutualFundsWorldRatings200ResponseMutualFund) *NullableGetMutualFundsWorldRatings200ResponseMutualFund {
	return &NullableGetMutualFundsWorldRatings200ResponseMutualFund{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldRatings200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldRatings200ResponseMutualFund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
