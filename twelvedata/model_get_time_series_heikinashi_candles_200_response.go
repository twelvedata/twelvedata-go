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

// checks if the GetTimeSeriesHeikinashiCandles200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHeikinashiCandles200Response{}

// GetTimeSeriesHeikinashiCandles200Response struct for GetTimeSeriesHeikinashiCandles200Response
type GetTimeSeriesHeikinashiCandles200Response struct {
	Meta GetTimeSeriesHeikinashiCandles200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesHeikinashiCandles200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesHeikinashiCandles200Response GetTimeSeriesHeikinashiCandles200Response

// NewGetTimeSeriesHeikinashiCandles200Response instantiates a new GetTimeSeriesHeikinashiCandles200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHeikinashiCandles200Response(meta GetTimeSeriesHeikinashiCandles200ResponseMeta, values []GetTimeSeriesHeikinashiCandles200ResponseValuesInner, status string) *GetTimeSeriesHeikinashiCandles200Response {
	this := GetTimeSeriesHeikinashiCandles200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesHeikinashiCandles200ResponseWithDefaults instantiates a new GetTimeSeriesHeikinashiCandles200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHeikinashiCandles200ResponseWithDefaults() *GetTimeSeriesHeikinashiCandles200Response {
	this := GetTimeSeriesHeikinashiCandles200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesHeikinashiCandles200Response) GetMeta() GetTimeSeriesHeikinashiCandles200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesHeikinashiCandles200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHeikinashiCandles200Response) GetMetaOk() (*GetTimeSeriesHeikinashiCandles200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesHeikinashiCandles200Response) SetMeta(v GetTimeSeriesHeikinashiCandles200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesHeikinashiCandles200Response) GetValues() []GetTimeSeriesHeikinashiCandles200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesHeikinashiCandles200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHeikinashiCandles200Response) GetValuesOk() ([]GetTimeSeriesHeikinashiCandles200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesHeikinashiCandles200Response) SetValues(v []GetTimeSeriesHeikinashiCandles200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesHeikinashiCandles200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHeikinashiCandles200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesHeikinashiCandles200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesHeikinashiCandles200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHeikinashiCandles200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesHeikinashiCandles200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesHeikinashiCandles200Response := _GetTimeSeriesHeikinashiCandles200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesHeikinashiCandles200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHeikinashiCandles200Response(varGetTimeSeriesHeikinashiCandles200Response)

	return err
}

type NullableGetTimeSeriesHeikinashiCandles200Response struct {
	value *GetTimeSeriesHeikinashiCandles200Response
	isSet bool
}

func (v NullableGetTimeSeriesHeikinashiCandles200Response) Get() *GetTimeSeriesHeikinashiCandles200Response {
	return v.value
}

func (v *NullableGetTimeSeriesHeikinashiCandles200Response) Set(val *GetTimeSeriesHeikinashiCandles200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHeikinashiCandles200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHeikinashiCandles200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHeikinashiCandles200Response(val *GetTimeSeriesHeikinashiCandles200Response) *NullableGetTimeSeriesHeikinashiCandles200Response {
	return &NullableGetTimeSeriesHeikinashiCandles200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHeikinashiCandles200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHeikinashiCandles200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
