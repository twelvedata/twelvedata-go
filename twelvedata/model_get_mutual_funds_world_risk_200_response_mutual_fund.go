// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorldRisk200ResponseMutualFund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldRisk200ResponseMutualFund{}

// GetMutualFundsWorldRisk200ResponseMutualFund Mutual fund information
type GetMutualFundsWorldRisk200ResponseMutualFund struct {
	Risk *ResponseMutualFundWorldRisk `json:"risk,omitempty"`
}

// NewGetMutualFundsWorldRisk200ResponseMutualFund instantiates a new GetMutualFundsWorldRisk200ResponseMutualFund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldRisk200ResponseMutualFund() *GetMutualFundsWorldRisk200ResponseMutualFund {
	this := GetMutualFundsWorldRisk200ResponseMutualFund{}
	return &this
}

// NewGetMutualFundsWorldRisk200ResponseMutualFundWithDefaults instantiates a new GetMutualFundsWorldRisk200ResponseMutualFund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldRisk200ResponseMutualFundWithDefaults() *GetMutualFundsWorldRisk200ResponseMutualFund {
	this := GetMutualFundsWorldRisk200ResponseMutualFund{}
	return &this
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *GetMutualFundsWorldRisk200ResponseMutualFund) GetRisk() ResponseMutualFundWorldRisk {
	if o == nil || IsNil(o.Risk) {
		var ret ResponseMutualFundWorldRisk
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldRisk200ResponseMutualFund) GetRiskOk() (*ResponseMutualFundWorldRisk, bool) {
	if o == nil || IsNil(o.Risk) {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *GetMutualFundsWorldRisk200ResponseMutualFund) HasRisk() bool {
	if o != nil && !IsNil(o.Risk) {
		return true
	}

	return false
}

// SetRisk gets a reference to the given ResponseMutualFundWorldRisk and assigns it to the Risk field.
func (o *GetMutualFundsWorldRisk200ResponseMutualFund) SetRisk(v ResponseMutualFundWorldRisk) {
	o.Risk = &v
}

func (o GetMutualFundsWorldRisk200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldRisk200ResponseMutualFund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Risk) {
		toSerialize["risk"] = o.Risk
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorldRisk200ResponseMutualFund struct {
	value *GetMutualFundsWorldRisk200ResponseMutualFund
	isSet bool
}

func (v NullableGetMutualFundsWorldRisk200ResponseMutualFund) Get() *GetMutualFundsWorldRisk200ResponseMutualFund {
	return v.value
}

func (v *NullableGetMutualFundsWorldRisk200ResponseMutualFund) Set(val *GetMutualFundsWorldRisk200ResponseMutualFund) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldRisk200ResponseMutualFund) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldRisk200ResponseMutualFund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldRisk200ResponseMutualFund(val *GetMutualFundsWorldRisk200ResponseMutualFund) *NullableGetMutualFundsWorldRisk200ResponseMutualFund {
	return &NullableGetMutualFundsWorldRisk200ResponseMutualFund{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldRisk200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldRisk200ResponseMutualFund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
