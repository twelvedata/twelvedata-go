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

// checks if the GetKeyExecutives200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetKeyExecutives200Response{}

// GetKeyExecutives200Response struct for GetKeyExecutives200Response
type GetKeyExecutives200Response struct {
	Meta GetKeyExecutives200ResponseMeta `json:"meta"`
	// List of key executives
	KeyExecutives []GetKeyExecutives200ResponseKeyExecutivesInner `json:"key_executives,omitempty"`
}

type _GetKeyExecutives200Response GetKeyExecutives200Response

// NewGetKeyExecutives200Response instantiates a new GetKeyExecutives200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetKeyExecutives200Response(meta GetKeyExecutives200ResponseMeta) *GetKeyExecutives200Response {
	this := GetKeyExecutives200Response{}
	this.Meta = meta
	return &this
}

// NewGetKeyExecutives200ResponseWithDefaults instantiates a new GetKeyExecutives200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetKeyExecutives200ResponseWithDefaults() *GetKeyExecutives200Response {
	this := GetKeyExecutives200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetKeyExecutives200Response) GetMeta() GetKeyExecutives200ResponseMeta {
	if o == nil {
		var ret GetKeyExecutives200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetKeyExecutives200Response) GetMetaOk() (*GetKeyExecutives200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetKeyExecutives200Response) SetMeta(v GetKeyExecutives200ResponseMeta) {
	o.Meta = v
}

// GetKeyExecutives returns the KeyExecutives field value if set, zero value otherwise.
func (o *GetKeyExecutives200Response) GetKeyExecutives() []GetKeyExecutives200ResponseKeyExecutivesInner {
	if o == nil || IsNil(o.KeyExecutives) {
		var ret []GetKeyExecutives200ResponseKeyExecutivesInner
		return ret
	}
	return o.KeyExecutives
}

// GetKeyExecutivesOk returns a tuple with the KeyExecutives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKeyExecutives200Response) GetKeyExecutivesOk() ([]GetKeyExecutives200ResponseKeyExecutivesInner, bool) {
	if o == nil || IsNil(o.KeyExecutives) {
		return nil, false
	}
	return o.KeyExecutives, true
}

// HasKeyExecutives returns a boolean if a field has been set.
func (o *GetKeyExecutives200Response) HasKeyExecutives() bool {
	if o != nil && !IsNil(o.KeyExecutives) {
		return true
	}

	return false
}

// SetKeyExecutives gets a reference to the given []GetKeyExecutives200ResponseKeyExecutivesInner and assigns it to the KeyExecutives field.
func (o *GetKeyExecutives200Response) SetKeyExecutives(v []GetKeyExecutives200ResponseKeyExecutivesInner) {
	o.KeyExecutives = v
}

func (o GetKeyExecutives200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetKeyExecutives200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	if !IsNil(o.KeyExecutives) {
		toSerialize["key_executives"] = o.KeyExecutives
	}
	return toSerialize, nil
}

func (o *GetKeyExecutives200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
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

	varGetKeyExecutives200Response := _GetKeyExecutives200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetKeyExecutives200Response)

	if err != nil {
		return err
	}

	*o = GetKeyExecutives200Response(varGetKeyExecutives200Response)

	return err
}

type NullableGetKeyExecutives200Response struct {
	value *GetKeyExecutives200Response
	isSet bool
}

func (v NullableGetKeyExecutives200Response) Get() *GetKeyExecutives200Response {
	return v.value
}

func (v *NullableGetKeyExecutives200Response) Set(val *GetKeyExecutives200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetKeyExecutives200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetKeyExecutives200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetKeyExecutives200Response(val *GetKeyExecutives200Response) *NullableGetKeyExecutives200Response {
	return &NullableGetKeyExecutives200Response{value: val, isSet: true}
}

func (v NullableGetKeyExecutives200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetKeyExecutives200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
