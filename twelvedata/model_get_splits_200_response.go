// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetSplits200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSplits200Response{}

// GetSplits200Response struct for GetSplits200Response
type GetSplits200Response struct {
	Meta GetSplits200ResponseMeta `json:"meta"`
	// List of stock splits
	Splits []GetSplits200ResponseSplitsInner `json:"splits"`
}

type _GetSplits200Response GetSplits200Response

// NewGetSplits200Response instantiates a new GetSplits200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSplits200Response(meta GetSplits200ResponseMeta, splits []GetSplits200ResponseSplitsInner) *GetSplits200Response {
	this := GetSplits200Response{}
	this.Meta = meta
	this.Splits = splits
	return &this
}

// NewGetSplits200ResponseWithDefaults instantiates a new GetSplits200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSplits200ResponseWithDefaults() *GetSplits200Response {
	this := GetSplits200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetSplits200Response) GetMeta() GetSplits200ResponseMeta {
	if o == nil {
		var ret GetSplits200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetSplits200Response) GetMetaOk() (*GetSplits200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetSplits200Response) SetMeta(v GetSplits200ResponseMeta) {
	o.Meta = v
}

// GetSplits returns the Splits field value
func (o *GetSplits200Response) GetSplits() []GetSplits200ResponseSplitsInner {
	if o == nil {
		var ret []GetSplits200ResponseSplitsInner
		return ret
	}

	return o.Splits
}

// GetSplitsOk returns a tuple with the Splits field value
// and a boolean to check if the value has been set.
func (o *GetSplits200Response) GetSplitsOk() ([]GetSplits200ResponseSplitsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Splits, true
}

// SetSplits sets field value
func (o *GetSplits200Response) SetSplits(v []GetSplits200ResponseSplitsInner) {
	o.Splits = v
}

func (o GetSplits200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSplits200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["splits"] = o.Splits
	return toSerialize, nil
}

func (o *GetSplits200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"splits",
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

	varGetSplits200Response := _GetSplits200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetSplits200Response)

	if err != nil {
		return err
	}

	*o = GetSplits200Response(varGetSplits200Response)

	return err
}

type NullableGetSplits200Response struct {
	value *GetSplits200Response
	isSet bool
}

func (v NullableGetSplits200Response) Get() *GetSplits200Response {
	return v.value
}

func (v *NullableGetSplits200Response) Set(val *GetSplits200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSplits200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSplits200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSplits200Response(val *GetSplits200Response) *NullableGetSplits200Response {
	return &NullableGetSplits200Response{value: val, isSet: true}
}

func (v NullableGetSplits200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSplits200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
