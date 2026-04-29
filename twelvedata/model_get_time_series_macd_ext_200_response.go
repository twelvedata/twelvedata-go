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

// checks if the GetTimeSeriesMacdExt200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMacdExt200Response{}

// GetTimeSeriesMacdExt200Response struct for GetTimeSeriesMacdExt200Response
type GetTimeSeriesMacdExt200Response struct {
	Meta GetTimeSeriesMacdExt200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMacdExt200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMacdExt200Response GetTimeSeriesMacdExt200Response

// NewGetTimeSeriesMacdExt200Response instantiates a new GetTimeSeriesMacdExt200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMacdExt200Response(meta GetTimeSeriesMacdExt200ResponseMeta, values []GetTimeSeriesMacdExt200ResponseValuesInner, status string) *GetTimeSeriesMacdExt200Response {
	this := GetTimeSeriesMacdExt200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMacdExt200ResponseWithDefaults instantiates a new GetTimeSeriesMacdExt200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMacdExt200ResponseWithDefaults() *GetTimeSeriesMacdExt200Response {
	this := GetTimeSeriesMacdExt200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMacdExt200Response) GetMeta() GetTimeSeriesMacdExt200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMacdExt200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200Response) GetMetaOk() (*GetTimeSeriesMacdExt200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMacdExt200Response) SetMeta(v GetTimeSeriesMacdExt200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMacdExt200Response) GetValues() []GetTimeSeriesMacdExt200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMacdExt200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200Response) GetValuesOk() ([]GetTimeSeriesMacdExt200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMacdExt200Response) SetValues(v []GetTimeSeriesMacdExt200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMacdExt200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMacdExt200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMacdExt200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMacdExt200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMacdExt200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"values",
		"status",
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

	varGetTimeSeriesMacdExt200Response := _GetTimeSeriesMacdExt200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMacdExt200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMacdExt200Response(varGetTimeSeriesMacdExt200Response)

	return err
}

type NullableGetTimeSeriesMacdExt200Response struct {
	value *GetTimeSeriesMacdExt200Response
	isSet bool
}

func (v NullableGetTimeSeriesMacdExt200Response) Get() *GetTimeSeriesMacdExt200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMacdExt200Response) Set(val *GetTimeSeriesMacdExt200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMacdExt200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMacdExt200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMacdExt200Response(val *GetTimeSeriesMacdExt200Response) *NullableGetTimeSeriesMacdExt200Response {
	return &NullableGetTimeSeriesMacdExt200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMacdExt200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMacdExt200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
