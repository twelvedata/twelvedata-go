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

// checks if the GetTimeSeriesCross200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCross200Response{}

// GetTimeSeriesCross200Response struct for GetTimeSeriesCross200Response
type GetTimeSeriesCross200Response struct {
	Meta CrossMeta `json:"meta"`
	// Array of time series data points
	Values []TimeSeriesCrossItem `json:"values"`
}

type _GetTimeSeriesCross200Response GetTimeSeriesCross200Response

// NewGetTimeSeriesCross200Response instantiates a new GetTimeSeriesCross200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCross200Response(meta CrossMeta, values []TimeSeriesCrossItem) *GetTimeSeriesCross200Response {
	this := GetTimeSeriesCross200Response{}
	this.Meta = meta
	this.Values = values
	return &this
}

// NewGetTimeSeriesCross200ResponseWithDefaults instantiates a new GetTimeSeriesCross200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCross200ResponseWithDefaults() *GetTimeSeriesCross200Response {
	this := GetTimeSeriesCross200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesCross200Response) GetMeta() CrossMeta {
	if o == nil {
		var ret CrossMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCross200Response) GetMetaOk() (*CrossMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesCross200Response) SetMeta(v CrossMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesCross200Response) GetValues() []TimeSeriesCrossItem {
	if o == nil {
		var ret []TimeSeriesCrossItem
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCross200Response) GetValuesOk() ([]TimeSeriesCrossItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesCross200Response) SetValues(v []TimeSeriesCrossItem) {
	o.Values = v
}

func (o GetTimeSeriesCross200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCross200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *GetTimeSeriesCross200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesCross200Response := _GetTimeSeriesCross200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesCross200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCross200Response(varGetTimeSeriesCross200Response)

	return err
}

type NullableGetTimeSeriesCross200Response struct {
	value *GetTimeSeriesCross200Response
	isSet bool
}

func (v NullableGetTimeSeriesCross200Response) Get() *GetTimeSeriesCross200Response {
	return v.value
}

func (v *NullableGetTimeSeriesCross200Response) Set(val *GetTimeSeriesCross200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCross200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCross200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCross200Response(val *GetTimeSeriesCross200Response) *NullableGetTimeSeriesCross200Response {
	return &NullableGetTimeSeriesCross200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCross200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCross200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
