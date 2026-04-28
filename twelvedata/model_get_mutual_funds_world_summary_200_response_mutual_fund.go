/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorldSummary200ResponseMutualFund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldSummary200ResponseMutualFund{}

// GetMutualFundsWorldSummary200ResponseMutualFund Mutual fund information
type GetMutualFundsWorldSummary200ResponseMutualFund struct {
	Summary *ResponseMutualFundWorldSummary `json:"summary,omitempty"`
}

// NewGetMutualFundsWorldSummary200ResponseMutualFund instantiates a new GetMutualFundsWorldSummary200ResponseMutualFund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldSummary200ResponseMutualFund() *GetMutualFundsWorldSummary200ResponseMutualFund {
	this := GetMutualFundsWorldSummary200ResponseMutualFund{}
	return &this
}

// NewGetMutualFundsWorldSummary200ResponseMutualFundWithDefaults instantiates a new GetMutualFundsWorldSummary200ResponseMutualFund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldSummary200ResponseMutualFundWithDefaults() *GetMutualFundsWorldSummary200ResponseMutualFund {
	this := GetMutualFundsWorldSummary200ResponseMutualFund{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *GetMutualFundsWorldSummary200ResponseMutualFund) GetSummary() ResponseMutualFundWorldSummary {
	if o == nil || IsNil(o.Summary) {
		var ret ResponseMutualFundWorldSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldSummary200ResponseMutualFund) GetSummaryOk() (*ResponseMutualFundWorldSummary, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *GetMutualFundsWorldSummary200ResponseMutualFund) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given ResponseMutualFundWorldSummary and assigns it to the Summary field.
func (o *GetMutualFundsWorldSummary200ResponseMutualFund) SetSummary(v ResponseMutualFundWorldSummary) {
	o.Summary = &v
}

func (o GetMutualFundsWorldSummary200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldSummary200ResponseMutualFund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorldSummary200ResponseMutualFund struct {
	value *GetMutualFundsWorldSummary200ResponseMutualFund
	isSet bool
}

func (v NullableGetMutualFundsWorldSummary200ResponseMutualFund) Get() *GetMutualFundsWorldSummary200ResponseMutualFund {
	return v.value
}

func (v *NullableGetMutualFundsWorldSummary200ResponseMutualFund) Set(val *GetMutualFundsWorldSummary200ResponseMutualFund) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldSummary200ResponseMutualFund) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldSummary200ResponseMutualFund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldSummary200ResponseMutualFund(val *GetMutualFundsWorldSummary200ResponseMutualFund) *NullableGetMutualFundsWorldSummary200ResponseMutualFund {
	return &NullableGetMutualFundsWorldSummary200ResponseMutualFund{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldSummary200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldSummary200ResponseMutualFund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
