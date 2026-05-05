// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetDirectHolders200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDirectHolders200Response{}

// GetDirectHolders200Response struct for GetDirectHolders200Response
type GetDirectHolders200Response struct {
	Meta GetDirectHolders200ResponseMeta `json:"meta"`
	// List of direct holders for the financial instrument
	DirectHolders []DirectHolderItem `json:"direct_holders"`
}

type _GetDirectHolders200Response GetDirectHolders200Response

// NewGetDirectHolders200Response instantiates a new GetDirectHolders200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDirectHolders200Response(meta GetDirectHolders200ResponseMeta, directHolders []DirectHolderItem) *GetDirectHolders200Response {
	this := GetDirectHolders200Response{}
	this.Meta = meta
	this.DirectHolders = directHolders
	return &this
}

// NewGetDirectHolders200ResponseWithDefaults instantiates a new GetDirectHolders200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDirectHolders200ResponseWithDefaults() *GetDirectHolders200Response {
	this := GetDirectHolders200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetDirectHolders200Response) GetMeta() GetDirectHolders200ResponseMeta {
	if o == nil {
		var ret GetDirectHolders200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetDirectHolders200Response) GetMetaOk() (*GetDirectHolders200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetDirectHolders200Response) SetMeta(v GetDirectHolders200ResponseMeta) {
	o.Meta = v
}

// GetDirectHolders returns the DirectHolders field value
func (o *GetDirectHolders200Response) GetDirectHolders() []DirectHolderItem {
	if o == nil {
		var ret []DirectHolderItem
		return ret
	}

	return o.DirectHolders
}

// GetDirectHoldersOk returns a tuple with the DirectHolders field value
// and a boolean to check if the value has been set.
func (o *GetDirectHolders200Response) GetDirectHoldersOk() ([]DirectHolderItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectHolders, true
}

// SetDirectHolders sets field value
func (o *GetDirectHolders200Response) SetDirectHolders(v []DirectHolderItem) {
	o.DirectHolders = v
}

func (o GetDirectHolders200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDirectHolders200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["direct_holders"] = o.DirectHolders
	return toSerialize, nil
}

func (o *GetDirectHolders200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"direct_holders",
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

	varGetDirectHolders200Response := _GetDirectHolders200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetDirectHolders200Response)

	if err != nil {
		return err
	}

	*o = GetDirectHolders200Response(varGetDirectHolders200Response)

	return err
}

type NullableGetDirectHolders200Response struct {
	value *GetDirectHolders200Response
	isSet bool
}

func (v NullableGetDirectHolders200Response) Get() *GetDirectHolders200Response {
	return v.value
}

func (v *NullableGetDirectHolders200Response) Set(val *GetDirectHolders200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDirectHolders200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDirectHolders200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDirectHolders200Response(val *GetDirectHolders200Response) *NullableGetDirectHolders200Response {
	return &NullableGetDirectHolders200Response{value: val, isSet: true}
}

func (v NullableGetDirectHolders200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDirectHolders200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
