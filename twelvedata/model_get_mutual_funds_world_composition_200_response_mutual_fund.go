// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorldComposition200ResponseMutualFund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldComposition200ResponseMutualFund{}

// GetMutualFundsWorldComposition200ResponseMutualFund Mutual fund information
type GetMutualFundsWorldComposition200ResponseMutualFund struct {
	Composition *ResponseMutualFundWorldComposition `json:"composition,omitempty"`
}

// NewGetMutualFundsWorldComposition200ResponseMutualFund instantiates a new GetMutualFundsWorldComposition200ResponseMutualFund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldComposition200ResponseMutualFund() *GetMutualFundsWorldComposition200ResponseMutualFund {
	this := GetMutualFundsWorldComposition200ResponseMutualFund{}
	return &this
}

// NewGetMutualFundsWorldComposition200ResponseMutualFundWithDefaults instantiates a new GetMutualFundsWorldComposition200ResponseMutualFund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldComposition200ResponseMutualFundWithDefaults() *GetMutualFundsWorldComposition200ResponseMutualFund {
	this := GetMutualFundsWorldComposition200ResponseMutualFund{}
	return &this
}

// GetComposition returns the Composition field value if set, zero value otherwise.
func (o *GetMutualFundsWorldComposition200ResponseMutualFund) GetComposition() ResponseMutualFundWorldComposition {
	if o == nil || IsNil(o.Composition) {
		var ret ResponseMutualFundWorldComposition
		return ret
	}
	return *o.Composition
}

// GetCompositionOk returns a tuple with the Composition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldComposition200ResponseMutualFund) GetCompositionOk() (*ResponseMutualFundWorldComposition, bool) {
	if o == nil || IsNil(o.Composition) {
		return nil, false
	}
	return o.Composition, true
}

// HasComposition returns a boolean if a field has been set.
func (o *GetMutualFundsWorldComposition200ResponseMutualFund) HasComposition() bool {
	if o != nil && !IsNil(o.Composition) {
		return true
	}

	return false
}

// SetComposition gets a reference to the given ResponseMutualFundWorldComposition and assigns it to the Composition field.
func (o *GetMutualFundsWorldComposition200ResponseMutualFund) SetComposition(v ResponseMutualFundWorldComposition) {
	o.Composition = &v
}

func (o GetMutualFundsWorldComposition200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldComposition200ResponseMutualFund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Composition) {
		toSerialize["composition"] = o.Composition
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorldComposition200ResponseMutualFund struct {
	value *GetMutualFundsWorldComposition200ResponseMutualFund
	isSet bool
}

func (v NullableGetMutualFundsWorldComposition200ResponseMutualFund) Get() *GetMutualFundsWorldComposition200ResponseMutualFund {
	return v.value
}

func (v *NullableGetMutualFundsWorldComposition200ResponseMutualFund) Set(val *GetMutualFundsWorldComposition200ResponseMutualFund) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldComposition200ResponseMutualFund) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldComposition200ResponseMutualFund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldComposition200ResponseMutualFund(val *GetMutualFundsWorldComposition200ResponseMutualFund) *NullableGetMutualFundsWorldComposition200ResponseMutualFund {
	return &NullableGetMutualFundsWorldComposition200ResponseMutualFund{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldComposition200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldComposition200ResponseMutualFund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
