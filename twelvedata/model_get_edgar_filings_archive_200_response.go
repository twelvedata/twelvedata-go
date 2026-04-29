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

// checks if the GetEdgarFilingsArchive200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEdgarFilingsArchive200Response{}

// GetEdgarFilingsArchive200Response struct for GetEdgarFilingsArchive200Response
type GetEdgarFilingsArchive200Response struct {
	Meta GetEdgarFilingsArchive200ResponseMeta `json:"meta"`
	// List of filings
	Values []EdgarFilingValue `json:"values"`
}

type _GetEdgarFilingsArchive200Response GetEdgarFilingsArchive200Response

// NewGetEdgarFilingsArchive200Response instantiates a new GetEdgarFilingsArchive200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEdgarFilingsArchive200Response(meta GetEdgarFilingsArchive200ResponseMeta, values []EdgarFilingValue) *GetEdgarFilingsArchive200Response {
	this := GetEdgarFilingsArchive200Response{}
	this.Meta = meta
	this.Values = values
	return &this
}

// NewGetEdgarFilingsArchive200ResponseWithDefaults instantiates a new GetEdgarFilingsArchive200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEdgarFilingsArchive200ResponseWithDefaults() *GetEdgarFilingsArchive200Response {
	this := GetEdgarFilingsArchive200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetEdgarFilingsArchive200Response) GetMeta() GetEdgarFilingsArchive200ResponseMeta {
	if o == nil {
		var ret GetEdgarFilingsArchive200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetEdgarFilingsArchive200Response) GetMetaOk() (*GetEdgarFilingsArchive200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetEdgarFilingsArchive200Response) SetMeta(v GetEdgarFilingsArchive200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetEdgarFilingsArchive200Response) GetValues() []EdgarFilingValue {
	if o == nil {
		var ret []EdgarFilingValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetEdgarFilingsArchive200Response) GetValuesOk() ([]EdgarFilingValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetEdgarFilingsArchive200Response) SetValues(v []EdgarFilingValue) {
	o.Values = v
}

func (o GetEdgarFilingsArchive200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEdgarFilingsArchive200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *GetEdgarFilingsArchive200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"values",
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

	varGetEdgarFilingsArchive200Response := _GetEdgarFilingsArchive200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetEdgarFilingsArchive200Response)

	if err != nil {
		return err
	}

	*o = GetEdgarFilingsArchive200Response(varGetEdgarFilingsArchive200Response)

	return err
}

type NullableGetEdgarFilingsArchive200Response struct {
	value *GetEdgarFilingsArchive200Response
	isSet bool
}

func (v NullableGetEdgarFilingsArchive200Response) Get() *GetEdgarFilingsArchive200Response {
	return v.value
}

func (v *NullableGetEdgarFilingsArchive200Response) Set(val *GetEdgarFilingsArchive200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEdgarFilingsArchive200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEdgarFilingsArchive200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEdgarFilingsArchive200Response(val *GetEdgarFilingsArchive200Response) *NullableGetEdgarFilingsArchive200Response {
	return &NullableGetEdgarFilingsArchive200Response{value: val, isSet: true}
}

func (v NullableGetEdgarFilingsArchive200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEdgarFilingsArchive200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
