// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetDividends200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDividends200Response{}

// GetDividends200Response struct for GetDividends200Response
type GetDividends200Response struct {
	Meta GetDividends200ResponseMeta `json:"meta"`
	// List of dividends
	Dividends []GetDividends200ResponseDividendsInner `json:"dividends"`
}

type _GetDividends200Response GetDividends200Response

// NewGetDividends200Response instantiates a new GetDividends200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDividends200Response(meta GetDividends200ResponseMeta, dividends []GetDividends200ResponseDividendsInner) *GetDividends200Response {
	this := GetDividends200Response{}
	this.Meta = meta
	this.Dividends = dividends
	return &this
}

// NewGetDividends200ResponseWithDefaults instantiates a new GetDividends200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDividends200ResponseWithDefaults() *GetDividends200Response {
	this := GetDividends200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetDividends200Response) GetMeta() GetDividends200ResponseMeta {
	if o == nil {
		var ret GetDividends200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetDividends200Response) GetMetaOk() (*GetDividends200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetDividends200Response) SetMeta(v GetDividends200ResponseMeta) {
	o.Meta = v
}

// GetDividends returns the Dividends field value
func (o *GetDividends200Response) GetDividends() []GetDividends200ResponseDividendsInner {
	if o == nil {
		var ret []GetDividends200ResponseDividendsInner
		return ret
	}

	return o.Dividends
}

// GetDividendsOk returns a tuple with the Dividends field value
// and a boolean to check if the value has been set.
func (o *GetDividends200Response) GetDividendsOk() ([]GetDividends200ResponseDividendsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dividends, true
}

// SetDividends sets field value
func (o *GetDividends200Response) SetDividends(v []GetDividends200ResponseDividendsInner) {
	o.Dividends = v
}

func (o GetDividends200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDividends200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["dividends"] = o.Dividends
	return toSerialize, nil
}

func (o *GetDividends200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"dividends",
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

	varGetDividends200Response := _GetDividends200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetDividends200Response)

	if err != nil {
		return err
	}

	*o = GetDividends200Response(varGetDividends200Response)

	return err
}

type NullableGetDividends200Response struct {
	value *GetDividends200Response
	isSet bool
}

func (v NullableGetDividends200Response) Get() *GetDividends200Response {
	return v.value
}

func (v *NullableGetDividends200Response) Set(val *GetDividends200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDividends200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDividends200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDividends200Response(val *GetDividends200Response) *NullableGetDividends200Response {
	return &NullableGetDividends200Response{value: val, isSet: true}
}

func (v NullableGetDividends200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDividends200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
