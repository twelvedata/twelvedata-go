/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorldPerformance200ResponseMutualFund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldPerformance200ResponseMutualFund{}

// GetMutualFundsWorldPerformance200ResponseMutualFund Mutual fund information
type GetMutualFundsWorldPerformance200ResponseMutualFund struct {
	Performance *ResponseMutualFundWorldPerformance `json:"performance,omitempty"`
}

// NewGetMutualFundsWorldPerformance200ResponseMutualFund instantiates a new GetMutualFundsWorldPerformance200ResponseMutualFund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldPerformance200ResponseMutualFund() *GetMutualFundsWorldPerformance200ResponseMutualFund {
	this := GetMutualFundsWorldPerformance200ResponseMutualFund{}
	return &this
}

// NewGetMutualFundsWorldPerformance200ResponseMutualFundWithDefaults instantiates a new GetMutualFundsWorldPerformance200ResponseMutualFund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldPerformance200ResponseMutualFundWithDefaults() *GetMutualFundsWorldPerformance200ResponseMutualFund {
	this := GetMutualFundsWorldPerformance200ResponseMutualFund{}
	return &this
}

// GetPerformance returns the Performance field value if set, zero value otherwise.
func (o *GetMutualFundsWorldPerformance200ResponseMutualFund) GetPerformance() ResponseMutualFundWorldPerformance {
	if o == nil || IsNil(o.Performance) {
		var ret ResponseMutualFundWorldPerformance
		return ret
	}
	return *o.Performance
}

// GetPerformanceOk returns a tuple with the Performance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldPerformance200ResponseMutualFund) GetPerformanceOk() (*ResponseMutualFundWorldPerformance, bool) {
	if o == nil || IsNil(o.Performance) {
		return nil, false
	}
	return o.Performance, true
}

// HasPerformance returns a boolean if a field has been set.
func (o *GetMutualFundsWorldPerformance200ResponseMutualFund) HasPerformance() bool {
	if o != nil && !IsNil(o.Performance) {
		return true
	}

	return false
}

// SetPerformance gets a reference to the given ResponseMutualFundWorldPerformance and assigns it to the Performance field.
func (o *GetMutualFundsWorldPerformance200ResponseMutualFund) SetPerformance(v ResponseMutualFundWorldPerformance) {
	o.Performance = &v
}

func (o GetMutualFundsWorldPerformance200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldPerformance200ResponseMutualFund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Performance) {
		toSerialize["performance"] = o.Performance
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorldPerformance200ResponseMutualFund struct {
	value *GetMutualFundsWorldPerformance200ResponseMutualFund
	isSet bool
}

func (v NullableGetMutualFundsWorldPerformance200ResponseMutualFund) Get() *GetMutualFundsWorldPerformance200ResponseMutualFund {
	return v.value
}

func (v *NullableGetMutualFundsWorldPerformance200ResponseMutualFund) Set(val *GetMutualFundsWorldPerformance200ResponseMutualFund) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldPerformance200ResponseMutualFund) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldPerformance200ResponseMutualFund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldPerformance200ResponseMutualFund(val *GetMutualFundsWorldPerformance200ResponseMutualFund) *NullableGetMutualFundsWorldPerformance200ResponseMutualFund {
	return &NullableGetMutualFundsWorldPerformance200ResponseMutualFund{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldPerformance200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldPerformance200ResponseMutualFund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
