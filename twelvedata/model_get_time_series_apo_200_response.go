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

// checks if the GetTimeSeriesApo200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesApo200Response{}

// GetTimeSeriesApo200Response struct for GetTimeSeriesApo200Response
type GetTimeSeriesApo200Response struct {
	Meta GetTimeSeriesApo200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesApo200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesApo200Response GetTimeSeriesApo200Response

// NewGetTimeSeriesApo200Response instantiates a new GetTimeSeriesApo200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesApo200Response(meta GetTimeSeriesApo200ResponseMeta, values []GetTimeSeriesApo200ResponseValuesInner, status string) *GetTimeSeriesApo200Response {
	this := GetTimeSeriesApo200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesApo200ResponseWithDefaults instantiates a new GetTimeSeriesApo200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesApo200ResponseWithDefaults() *GetTimeSeriesApo200Response {
	this := GetTimeSeriesApo200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesApo200Response) GetMeta() GetTimeSeriesApo200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesApo200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesApo200Response) GetMetaOk() (*GetTimeSeriesApo200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesApo200Response) SetMeta(v GetTimeSeriesApo200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesApo200Response) GetValues() []GetTimeSeriesApo200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesApo200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesApo200Response) GetValuesOk() ([]GetTimeSeriesApo200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesApo200Response) SetValues(v []GetTimeSeriesApo200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesApo200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesApo200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesApo200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesApo200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesApo200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesApo200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesApo200Response := _GetTimeSeriesApo200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesApo200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesApo200Response(varGetTimeSeriesApo200Response)

	return err
}

type NullableGetTimeSeriesApo200Response struct {
	value *GetTimeSeriesApo200Response
	isSet bool
}

func (v NullableGetTimeSeriesApo200Response) Get() *GetTimeSeriesApo200Response {
	return v.value
}

func (v *NullableGetTimeSeriesApo200Response) Set(val *GetTimeSeriesApo200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesApo200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesApo200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesApo200Response(val *GetTimeSeriesApo200Response) *NullableGetTimeSeriesApo200Response {
	return &NullableGetTimeSeriesApo200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesApo200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesApo200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
