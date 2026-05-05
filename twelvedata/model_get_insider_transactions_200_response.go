// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetInsiderTransactions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsiderTransactions200Response{}

// GetInsiderTransactions200Response struct for GetInsiderTransactions200Response
type GetInsiderTransactions200Response struct {
	Meta GetInsiderTransactions200ResponseMeta `json:"meta"`
	// List of insider transactions
	InsiderTransactions []GetInsiderTransactions200ResponseInsiderTransactionsInner `json:"insider_transactions"`
}

type _GetInsiderTransactions200Response GetInsiderTransactions200Response

// NewGetInsiderTransactions200Response instantiates a new GetInsiderTransactions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsiderTransactions200Response(meta GetInsiderTransactions200ResponseMeta, insiderTransactions []GetInsiderTransactions200ResponseInsiderTransactionsInner) *GetInsiderTransactions200Response {
	this := GetInsiderTransactions200Response{}
	this.Meta = meta
	this.InsiderTransactions = insiderTransactions
	return &this
}

// NewGetInsiderTransactions200ResponseWithDefaults instantiates a new GetInsiderTransactions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsiderTransactions200ResponseWithDefaults() *GetInsiderTransactions200Response {
	this := GetInsiderTransactions200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetInsiderTransactions200Response) GetMeta() GetInsiderTransactions200ResponseMeta {
	if o == nil {
		var ret GetInsiderTransactions200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200Response) GetMetaOk() (*GetInsiderTransactions200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetInsiderTransactions200Response) SetMeta(v GetInsiderTransactions200ResponseMeta) {
	o.Meta = v
}

// GetInsiderTransactions returns the InsiderTransactions field value
func (o *GetInsiderTransactions200Response) GetInsiderTransactions() []GetInsiderTransactions200ResponseInsiderTransactionsInner {
	if o == nil {
		var ret []GetInsiderTransactions200ResponseInsiderTransactionsInner
		return ret
	}

	return o.InsiderTransactions
}

// GetInsiderTransactionsOk returns a tuple with the InsiderTransactions field value
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200Response) GetInsiderTransactionsOk() ([]GetInsiderTransactions200ResponseInsiderTransactionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.InsiderTransactions, true
}

// SetInsiderTransactions sets field value
func (o *GetInsiderTransactions200Response) SetInsiderTransactions(v []GetInsiderTransactions200ResponseInsiderTransactionsInner) {
	o.InsiderTransactions = v
}

func (o GetInsiderTransactions200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsiderTransactions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["insider_transactions"] = o.InsiderTransactions
	return toSerialize, nil
}

func (o *GetInsiderTransactions200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"insider_transactions",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetInsiderTransactions200Response := _GetInsiderTransactions200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetInsiderTransactions200Response)

	if err != nil {
		return err
	}

	*o = GetInsiderTransactions200Response(varGetInsiderTransactions200Response)

	return err
}

type NullableGetInsiderTransactions200Response struct {
	value *GetInsiderTransactions200Response
	isSet bool
}

func (v NullableGetInsiderTransactions200Response) Get() *GetInsiderTransactions200Response {
	return v.value
}

func (v *NullableGetInsiderTransactions200Response) Set(val *GetInsiderTransactions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsiderTransactions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsiderTransactions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsiderTransactions200Response(val *GetInsiderTransactions200Response) *NullableGetInsiderTransactions200Response {
	return &NullableGetInsiderTransactions200Response{value: val, isSet: true}
}

func (v NullableGetInsiderTransactions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsiderTransactions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
