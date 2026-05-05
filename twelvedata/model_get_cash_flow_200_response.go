// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetCashFlow200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCashFlow200Response{}

// GetCashFlow200Response struct for GetCashFlow200Response
type GetCashFlow200Response struct {
	Meta GetCashFlow200ResponseMeta `json:"meta"`
	// Cash flow data
	CashFlow []CashFlowStruct `json:"cash_flow"`
}

type _GetCashFlow200Response GetCashFlow200Response

// NewGetCashFlow200Response instantiates a new GetCashFlow200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCashFlow200Response(meta GetCashFlow200ResponseMeta, cashFlow []CashFlowStruct) *GetCashFlow200Response {
	this := GetCashFlow200Response{}
	this.Meta = meta
	this.CashFlow = cashFlow
	return &this
}

// NewGetCashFlow200ResponseWithDefaults instantiates a new GetCashFlow200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCashFlow200ResponseWithDefaults() *GetCashFlow200Response {
	this := GetCashFlow200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetCashFlow200Response) GetMeta() GetCashFlow200ResponseMeta {
	if o == nil {
		var ret GetCashFlow200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetCashFlow200Response) GetMetaOk() (*GetCashFlow200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetCashFlow200Response) SetMeta(v GetCashFlow200ResponseMeta) {
	o.Meta = v
}

// GetCashFlow returns the CashFlow field value
func (o *GetCashFlow200Response) GetCashFlow() []CashFlowStruct {
	if o == nil {
		var ret []CashFlowStruct
		return ret
	}

	return o.CashFlow
}

// GetCashFlowOk returns a tuple with the CashFlow field value
// and a boolean to check if the value has been set.
func (o *GetCashFlow200Response) GetCashFlowOk() ([]CashFlowStruct, bool) {
	if o == nil {
		return nil, false
	}
	return o.CashFlow, true
}

// SetCashFlow sets field value
func (o *GetCashFlow200Response) SetCashFlow(v []CashFlowStruct) {
	o.CashFlow = v
}

func (o GetCashFlow200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCashFlow200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["cash_flow"] = o.CashFlow
	return toSerialize, nil
}

func (o *GetCashFlow200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"cash_flow",
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

	varGetCashFlow200Response := _GetCashFlow200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetCashFlow200Response)

	if err != nil {
		return err
	}

	*o = GetCashFlow200Response(varGetCashFlow200Response)

	return err
}

type NullableGetCashFlow200Response struct {
	value *GetCashFlow200Response
	isSet bool
}

func (v NullableGetCashFlow200Response) Get() *GetCashFlow200Response {
	return v.value
}

func (v *NullableGetCashFlow200Response) Set(val *GetCashFlow200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCashFlow200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCashFlow200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCashFlow200Response(val *GetCashFlow200Response) *NullableGetCashFlow200Response {
	return &NullableGetCashFlow200Response{value: val, isSet: true}
}

func (v NullableGetCashFlow200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCashFlow200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
