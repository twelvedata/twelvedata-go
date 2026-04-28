/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorldPurchaseInfo200ResponseMutualFund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldPurchaseInfo200ResponseMutualFund{}

// GetMutualFundsWorldPurchaseInfo200ResponseMutualFund Mutual fund information
type GetMutualFundsWorldPurchaseInfo200ResponseMutualFund struct {
	PurchaseInfo *ResponseMutualFundWorldPurchaseInfo `json:"purchase_info,omitempty"`
}

// NewGetMutualFundsWorldPurchaseInfo200ResponseMutualFund instantiates a new GetMutualFundsWorldPurchaseInfo200ResponseMutualFund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldPurchaseInfo200ResponseMutualFund() *GetMutualFundsWorldPurchaseInfo200ResponseMutualFund {
	this := GetMutualFundsWorldPurchaseInfo200ResponseMutualFund{}
	return &this
}

// NewGetMutualFundsWorldPurchaseInfo200ResponseMutualFundWithDefaults instantiates a new GetMutualFundsWorldPurchaseInfo200ResponseMutualFund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldPurchaseInfo200ResponseMutualFundWithDefaults() *GetMutualFundsWorldPurchaseInfo200ResponseMutualFund {
	this := GetMutualFundsWorldPurchaseInfo200ResponseMutualFund{}
	return &this
}

// GetPurchaseInfo returns the PurchaseInfo field value if set, zero value otherwise.
func (o *GetMutualFundsWorldPurchaseInfo200ResponseMutualFund) GetPurchaseInfo() ResponseMutualFundWorldPurchaseInfo {
	if o == nil || IsNil(o.PurchaseInfo) {
		var ret ResponseMutualFundWorldPurchaseInfo
		return ret
	}
	return *o.PurchaseInfo
}

// GetPurchaseInfoOk returns a tuple with the PurchaseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldPurchaseInfo200ResponseMutualFund) GetPurchaseInfoOk() (*ResponseMutualFundWorldPurchaseInfo, bool) {
	if o == nil || IsNil(o.PurchaseInfo) {
		return nil, false
	}
	return o.PurchaseInfo, true
}

// HasPurchaseInfo returns a boolean if a field has been set.
func (o *GetMutualFundsWorldPurchaseInfo200ResponseMutualFund) HasPurchaseInfo() bool {
	if o != nil && !IsNil(o.PurchaseInfo) {
		return true
	}

	return false
}

// SetPurchaseInfo gets a reference to the given ResponseMutualFundWorldPurchaseInfo and assigns it to the PurchaseInfo field.
func (o *GetMutualFundsWorldPurchaseInfo200ResponseMutualFund) SetPurchaseInfo(v ResponseMutualFundWorldPurchaseInfo) {
	o.PurchaseInfo = &v
}

func (o GetMutualFundsWorldPurchaseInfo200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldPurchaseInfo200ResponseMutualFund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PurchaseInfo) {
		toSerialize["purchase_info"] = o.PurchaseInfo
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorldPurchaseInfo200ResponseMutualFund struct {
	value *GetMutualFundsWorldPurchaseInfo200ResponseMutualFund
	isSet bool
}

func (v NullableGetMutualFundsWorldPurchaseInfo200ResponseMutualFund) Get() *GetMutualFundsWorldPurchaseInfo200ResponseMutualFund {
	return v.value
}

func (v *NullableGetMutualFundsWorldPurchaseInfo200ResponseMutualFund) Set(val *GetMutualFundsWorldPurchaseInfo200ResponseMutualFund) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldPurchaseInfo200ResponseMutualFund) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldPurchaseInfo200ResponseMutualFund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldPurchaseInfo200ResponseMutualFund(val *GetMutualFundsWorldPurchaseInfo200ResponseMutualFund) *NullableGetMutualFundsWorldPurchaseInfo200ResponseMutualFund {
	return &NullableGetMutualFundsWorldPurchaseInfo200ResponseMutualFund{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldPurchaseInfo200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldPurchaseInfo200ResponseMutualFund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
