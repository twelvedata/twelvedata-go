// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesAdOsc200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAdOsc200Response{}

// GetTimeSeriesAdOsc200Response struct for GetTimeSeriesAdOsc200Response
type GetTimeSeriesAdOsc200Response struct {
	Meta GetTimeSeriesAdOsc200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesAdOsc200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesAdOsc200Response GetTimeSeriesAdOsc200Response

// NewGetTimeSeriesAdOsc200Response instantiates a new GetTimeSeriesAdOsc200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAdOsc200Response(meta GetTimeSeriesAdOsc200ResponseMeta, values []GetTimeSeriesAdOsc200ResponseValuesInner, status string) *GetTimeSeriesAdOsc200Response {
	this := GetTimeSeriesAdOsc200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesAdOsc200ResponseWithDefaults instantiates a new GetTimeSeriesAdOsc200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAdOsc200ResponseWithDefaults() *GetTimeSeriesAdOsc200Response {
	this := GetTimeSeriesAdOsc200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesAdOsc200Response) GetMeta() GetTimeSeriesAdOsc200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesAdOsc200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200Response) GetMetaOk() (*GetTimeSeriesAdOsc200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesAdOsc200Response) SetMeta(v GetTimeSeriesAdOsc200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesAdOsc200Response) GetValues() []GetTimeSeriesAdOsc200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesAdOsc200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200Response) GetValuesOk() ([]GetTimeSeriesAdOsc200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesAdOsc200Response) SetValues(v []GetTimeSeriesAdOsc200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesAdOsc200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesAdOsc200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesAdOsc200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAdOsc200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesAdOsc200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAdOsc200Response := _GetTimeSeriesAdOsc200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAdOsc200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAdOsc200Response(varGetTimeSeriesAdOsc200Response)

	return err
}

type NullableGetTimeSeriesAdOsc200Response struct {
	value *GetTimeSeriesAdOsc200Response
	isSet bool
}

func (v NullableGetTimeSeriesAdOsc200Response) Get() *GetTimeSeriesAdOsc200Response {
	return v.value
}

func (v *NullableGetTimeSeriesAdOsc200Response) Set(val *GetTimeSeriesAdOsc200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAdOsc200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAdOsc200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAdOsc200Response(val *GetTimeSeriesAdOsc200Response) *NullableGetTimeSeriesAdOsc200Response {
	return &NullableGetTimeSeriesAdOsc200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAdOsc200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAdOsc200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
