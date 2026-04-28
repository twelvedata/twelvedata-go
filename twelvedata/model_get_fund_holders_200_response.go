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

// checks if the GetFundHolders200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFundHolders200Response{}

// GetFundHolders200Response struct for GetFundHolders200Response
type GetFundHolders200Response struct {
	Meta GetFundHolders200ResponseMeta `json:"meta"`
	// List of fund holders for the financial instrument
	FundHolders []FundHolderItem `json:"fund_holders"`
}

type _GetFundHolders200Response GetFundHolders200Response

// NewGetFundHolders200Response instantiates a new GetFundHolders200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFundHolders200Response(meta GetFundHolders200ResponseMeta, fundHolders []FundHolderItem) *GetFundHolders200Response {
	this := GetFundHolders200Response{}
	this.Meta = meta
	this.FundHolders = fundHolders
	return &this
}

// NewGetFundHolders200ResponseWithDefaults instantiates a new GetFundHolders200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFundHolders200ResponseWithDefaults() *GetFundHolders200Response {
	this := GetFundHolders200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetFundHolders200Response) GetMeta() GetFundHolders200ResponseMeta {
	if o == nil {
		var ret GetFundHolders200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetFundHolders200Response) GetMetaOk() (*GetFundHolders200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetFundHolders200Response) SetMeta(v GetFundHolders200ResponseMeta) {
	o.Meta = v
}

// GetFundHolders returns the FundHolders field value
func (o *GetFundHolders200Response) GetFundHolders() []FundHolderItem {
	if o == nil {
		var ret []FundHolderItem
		return ret
	}

	return o.FundHolders
}

// GetFundHoldersOk returns a tuple with the FundHolders field value
// and a boolean to check if the value has been set.
func (o *GetFundHolders200Response) GetFundHoldersOk() ([]FundHolderItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.FundHolders, true
}

// SetFundHolders sets field value
func (o *GetFundHolders200Response) SetFundHolders(v []FundHolderItem) {
	o.FundHolders = v
}

func (o GetFundHolders200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFundHolders200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["fund_holders"] = o.FundHolders
	return toSerialize, nil
}

func (o *GetFundHolders200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"fund_holders",
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

	varGetFundHolders200Response := _GetFundHolders200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFundHolders200Response)

	if err != nil {
		return err
	}

	*o = GetFundHolders200Response(varGetFundHolders200Response)

	return err
}

type NullableGetFundHolders200Response struct {
	value *GetFundHolders200Response
	isSet bool
}

func (v NullableGetFundHolders200Response) Get() *GetFundHolders200Response {
	return v.value
}

func (v *NullableGetFundHolders200Response) Set(val *GetFundHolders200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFundHolders200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFundHolders200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFundHolders200Response(val *GetFundHolders200Response) *NullableGetFundHolders200Response {
	return &NullableGetFundHolders200Response{value: val, isSet: true}
}

func (v NullableGetFundHolders200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFundHolders200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
