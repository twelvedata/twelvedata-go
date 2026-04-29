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

// checks if the GetInstitutionalHolders200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInstitutionalHolders200Response{}

// GetInstitutionalHolders200Response struct for GetInstitutionalHolders200Response
type GetInstitutionalHolders200Response struct {
	Meta GetFundHolders200ResponseMeta `json:"meta"`
	// List of institutional holders for the financial instrument
	InstitutionalHolders []InstitutionalHolderItem `json:"institutional_holders"`
}

type _GetInstitutionalHolders200Response GetInstitutionalHolders200Response

// NewGetInstitutionalHolders200Response instantiates a new GetInstitutionalHolders200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInstitutionalHolders200Response(meta GetFundHolders200ResponseMeta, institutionalHolders []InstitutionalHolderItem) *GetInstitutionalHolders200Response {
	this := GetInstitutionalHolders200Response{}
	this.Meta = meta
	this.InstitutionalHolders = institutionalHolders
	return &this
}

// NewGetInstitutionalHolders200ResponseWithDefaults instantiates a new GetInstitutionalHolders200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInstitutionalHolders200ResponseWithDefaults() *GetInstitutionalHolders200Response {
	this := GetInstitutionalHolders200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetInstitutionalHolders200Response) GetMeta() GetFundHolders200ResponseMeta {
	if o == nil {
		var ret GetFundHolders200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetInstitutionalHolders200Response) GetMetaOk() (*GetFundHolders200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetInstitutionalHolders200Response) SetMeta(v GetFundHolders200ResponseMeta) {
	o.Meta = v
}

// GetInstitutionalHolders returns the InstitutionalHolders field value
func (o *GetInstitutionalHolders200Response) GetInstitutionalHolders() []InstitutionalHolderItem {
	if o == nil {
		var ret []InstitutionalHolderItem
		return ret
	}

	return o.InstitutionalHolders
}

// GetInstitutionalHoldersOk returns a tuple with the InstitutionalHolders field value
// and a boolean to check if the value has been set.
func (o *GetInstitutionalHolders200Response) GetInstitutionalHoldersOk() ([]InstitutionalHolderItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstitutionalHolders, true
}

// SetInstitutionalHolders sets field value
func (o *GetInstitutionalHolders200Response) SetInstitutionalHolders(v []InstitutionalHolderItem) {
	o.InstitutionalHolders = v
}

func (o GetInstitutionalHolders200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInstitutionalHolders200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["institutional_holders"] = o.InstitutionalHolders
	return toSerialize, nil
}

func (o *GetInstitutionalHolders200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"institutional_holders",
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

	varGetInstitutionalHolders200Response := _GetInstitutionalHolders200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetInstitutionalHolders200Response)

	if err != nil {
		return err
	}

	*o = GetInstitutionalHolders200Response(varGetInstitutionalHolders200Response)

	return err
}

type NullableGetInstitutionalHolders200Response struct {
	value *GetInstitutionalHolders200Response
	isSet bool
}

func (v NullableGetInstitutionalHolders200Response) Get() *GetInstitutionalHolders200Response {
	return v.value
}

func (v *NullableGetInstitutionalHolders200Response) Set(val *GetInstitutionalHolders200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInstitutionalHolders200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInstitutionalHolders200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInstitutionalHolders200Response(val *GetInstitutionalHolders200Response) *NullableGetInstitutionalHolders200Response {
	return &NullableGetInstitutionalHolders200Response{value: val, isSet: true}
}

func (v NullableGetInstitutionalHolders200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInstitutionalHolders200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
