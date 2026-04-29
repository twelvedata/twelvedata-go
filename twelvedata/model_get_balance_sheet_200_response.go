/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetBalanceSheet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheet200Response{}

// GetBalanceSheet200Response struct for GetBalanceSheet200Response
type GetBalanceSheet200Response struct {
	Meta GetBalanceSheet200ResponseMeta `json:"meta"`
	// Array of balance sheet records
	BalanceSheet []GetBalanceSheet200ResponseBalanceSheetInner `json:"balance_sheet"`
}

type _GetBalanceSheet200Response GetBalanceSheet200Response

// NewGetBalanceSheet200Response instantiates a new GetBalanceSheet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheet200Response(meta GetBalanceSheet200ResponseMeta, balanceSheet []GetBalanceSheet200ResponseBalanceSheetInner) *GetBalanceSheet200Response {
	this := GetBalanceSheet200Response{}
	this.Meta = meta
	this.BalanceSheet = balanceSheet
	return &this
}

// NewGetBalanceSheet200ResponseWithDefaults instantiates a new GetBalanceSheet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheet200ResponseWithDefaults() *GetBalanceSheet200Response {
	this := GetBalanceSheet200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetBalanceSheet200Response) GetMeta() GetBalanceSheet200ResponseMeta {
	if o == nil {
		var ret GetBalanceSheet200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200Response) GetMetaOk() (*GetBalanceSheet200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetBalanceSheet200Response) SetMeta(v GetBalanceSheet200ResponseMeta) {
	o.Meta = v
}

// GetBalanceSheet returns the BalanceSheet field value
func (o *GetBalanceSheet200Response) GetBalanceSheet() []GetBalanceSheet200ResponseBalanceSheetInner {
	if o == nil {
		var ret []GetBalanceSheet200ResponseBalanceSheetInner
		return ret
	}

	return o.BalanceSheet
}

// GetBalanceSheetOk returns a tuple with the BalanceSheet field value
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200Response) GetBalanceSheetOk() ([]GetBalanceSheet200ResponseBalanceSheetInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.BalanceSheet, true
}

// SetBalanceSheet sets field value
func (o *GetBalanceSheet200Response) SetBalanceSheet(v []GetBalanceSheet200ResponseBalanceSheetInner) {
	o.BalanceSheet = v
}

func (o GetBalanceSheet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["balance_sheet"] = o.BalanceSheet
	return toSerialize, nil
}

func (o *GetBalanceSheet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"balance_sheet",
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

	varGetBalanceSheet200Response := _GetBalanceSheet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetBalanceSheet200Response)

	if err != nil {
		return err
	}

	*o = GetBalanceSheet200Response(varGetBalanceSheet200Response)

	return err
}

type NullableGetBalanceSheet200Response struct {
	value *GetBalanceSheet200Response
	isSet bool
}

func (v NullableGetBalanceSheet200Response) Get() *GetBalanceSheet200Response {
	return v.value
}

func (v *NullableGetBalanceSheet200Response) Set(val *GetBalanceSheet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheet200Response(val *GetBalanceSheet200Response) *NullableGetBalanceSheet200Response {
	return &NullableGetBalanceSheet200Response{value: val, isSet: true}
}

func (v NullableGetBalanceSheet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
