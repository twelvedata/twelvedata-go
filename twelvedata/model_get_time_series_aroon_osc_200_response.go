// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesAroonOsc200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAroonOsc200Response{}

// GetTimeSeriesAroonOsc200Response struct for GetTimeSeriesAroonOsc200Response
type GetTimeSeriesAroonOsc200Response struct {
	Meta GetTimeSeriesAroonOsc200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesAroonOsc200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesAroonOsc200Response GetTimeSeriesAroonOsc200Response

// NewGetTimeSeriesAroonOsc200Response instantiates a new GetTimeSeriesAroonOsc200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAroonOsc200Response(meta GetTimeSeriesAroonOsc200ResponseMeta, values []GetTimeSeriesAroonOsc200ResponseValuesInner, status string) *GetTimeSeriesAroonOsc200Response {
	this := GetTimeSeriesAroonOsc200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesAroonOsc200ResponseWithDefaults instantiates a new GetTimeSeriesAroonOsc200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAroonOsc200ResponseWithDefaults() *GetTimeSeriesAroonOsc200Response {
	this := GetTimeSeriesAroonOsc200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesAroonOsc200Response) GetMeta() GetTimeSeriesAroonOsc200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesAroonOsc200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroonOsc200Response) GetMetaOk() (*GetTimeSeriesAroonOsc200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesAroonOsc200Response) SetMeta(v GetTimeSeriesAroonOsc200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesAroonOsc200Response) GetValues() []GetTimeSeriesAroonOsc200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesAroonOsc200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroonOsc200Response) GetValuesOk() ([]GetTimeSeriesAroonOsc200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesAroonOsc200Response) SetValues(v []GetTimeSeriesAroonOsc200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesAroonOsc200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroonOsc200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesAroonOsc200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesAroonOsc200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAroonOsc200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesAroonOsc200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAroonOsc200Response := _GetTimeSeriesAroonOsc200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAroonOsc200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAroonOsc200Response(varGetTimeSeriesAroonOsc200Response)

	return err
}

type NullableGetTimeSeriesAroonOsc200Response struct {
	value *GetTimeSeriesAroonOsc200Response
	isSet bool
}

func (v NullableGetTimeSeriesAroonOsc200Response) Get() *GetTimeSeriesAroonOsc200Response {
	return v.value
}

func (v *NullableGetTimeSeriesAroonOsc200Response) Set(val *GetTimeSeriesAroonOsc200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAroonOsc200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAroonOsc200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAroonOsc200Response(val *GetTimeSeriesAroonOsc200Response) *NullableGetTimeSeriesAroonOsc200Response {
	return &NullableGetTimeSeriesAroonOsc200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAroonOsc200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAroonOsc200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
