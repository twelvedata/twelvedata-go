/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorldSustainability200ResponseMutualFund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldSustainability200ResponseMutualFund{}

// GetMutualFundsWorldSustainability200ResponseMutualFund Mutual fund information
type GetMutualFundsWorldSustainability200ResponseMutualFund struct {
	Sustainability *ResponseMutualFundWorldSustainability `json:"sustainability,omitempty"`
}

// NewGetMutualFundsWorldSustainability200ResponseMutualFund instantiates a new GetMutualFundsWorldSustainability200ResponseMutualFund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldSustainability200ResponseMutualFund() *GetMutualFundsWorldSustainability200ResponseMutualFund {
	this := GetMutualFundsWorldSustainability200ResponseMutualFund{}
	return &this
}

// NewGetMutualFundsWorldSustainability200ResponseMutualFundWithDefaults instantiates a new GetMutualFundsWorldSustainability200ResponseMutualFund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldSustainability200ResponseMutualFundWithDefaults() *GetMutualFundsWorldSustainability200ResponseMutualFund {
	this := GetMutualFundsWorldSustainability200ResponseMutualFund{}
	return &this
}

// GetSustainability returns the Sustainability field value if set, zero value otherwise.
func (o *GetMutualFundsWorldSustainability200ResponseMutualFund) GetSustainability() ResponseMutualFundWorldSustainability {
	if o == nil || IsNil(o.Sustainability) {
		var ret ResponseMutualFundWorldSustainability
		return ret
	}
	return *o.Sustainability
}

// GetSustainabilityOk returns a tuple with the Sustainability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldSustainability200ResponseMutualFund) GetSustainabilityOk() (*ResponseMutualFundWorldSustainability, bool) {
	if o == nil || IsNil(o.Sustainability) {
		return nil, false
	}
	return o.Sustainability, true
}

// HasSustainability returns a boolean if a field has been set.
func (o *GetMutualFundsWorldSustainability200ResponseMutualFund) HasSustainability() bool {
	if o != nil && !IsNil(o.Sustainability) {
		return true
	}

	return false
}

// SetSustainability gets a reference to the given ResponseMutualFundWorldSustainability and assigns it to the Sustainability field.
func (o *GetMutualFundsWorldSustainability200ResponseMutualFund) SetSustainability(v ResponseMutualFundWorldSustainability) {
	o.Sustainability = &v
}

func (o GetMutualFundsWorldSustainability200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldSustainability200ResponseMutualFund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sustainability) {
		toSerialize["sustainability"] = o.Sustainability
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorldSustainability200ResponseMutualFund struct {
	value *GetMutualFundsWorldSustainability200ResponseMutualFund
	isSet bool
}

func (v NullableGetMutualFundsWorldSustainability200ResponseMutualFund) Get() *GetMutualFundsWorldSustainability200ResponseMutualFund {
	return v.value
}

func (v *NullableGetMutualFundsWorldSustainability200ResponseMutualFund) Set(val *GetMutualFundsWorldSustainability200ResponseMutualFund) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldSustainability200ResponseMutualFund) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldSustainability200ResponseMutualFund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldSustainability200ResponseMutualFund(val *GetMutualFundsWorldSustainability200ResponseMutualFund) *NullableGetMutualFundsWorldSustainability200ResponseMutualFund {
	return &NullableGetMutualFundsWorldSustainability200ResponseMutualFund{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldSustainability200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldSustainability200ResponseMutualFund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
