// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesHtSine200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHtSine200Response{}

// GetTimeSeriesHtSine200Response struct for GetTimeSeriesHtSine200Response
type GetTimeSeriesHtSine200Response struct {
	Meta GetTimeSeriesHtSine200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesHtSine200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesHtSine200Response GetTimeSeriesHtSine200Response

// NewGetTimeSeriesHtSine200Response instantiates a new GetTimeSeriesHtSine200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHtSine200Response(meta GetTimeSeriesHtSine200ResponseMeta, values []GetTimeSeriesHtSine200ResponseValuesInner, status string) *GetTimeSeriesHtSine200Response {
	this := GetTimeSeriesHtSine200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesHtSine200ResponseWithDefaults instantiates a new GetTimeSeriesHtSine200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHtSine200ResponseWithDefaults() *GetTimeSeriesHtSine200Response {
	this := GetTimeSeriesHtSine200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesHtSine200Response) GetMeta() GetTimeSeriesHtSine200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesHtSine200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtSine200Response) GetMetaOk() (*GetTimeSeriesHtSine200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesHtSine200Response) SetMeta(v GetTimeSeriesHtSine200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesHtSine200Response) GetValues() []GetTimeSeriesHtSine200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesHtSine200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtSine200Response) GetValuesOk() ([]GetTimeSeriesHtSine200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesHtSine200Response) SetValues(v []GetTimeSeriesHtSine200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesHtSine200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtSine200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesHtSine200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesHtSine200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHtSine200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesHtSine200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesHtSine200Response := _GetTimeSeriesHtSine200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesHtSine200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHtSine200Response(varGetTimeSeriesHtSine200Response)

	return err
}

type NullableGetTimeSeriesHtSine200Response struct {
	value *GetTimeSeriesHtSine200Response
	isSet bool
}

func (v NullableGetTimeSeriesHtSine200Response) Get() *GetTimeSeriesHtSine200Response {
	return v.value
}

func (v *NullableGetTimeSeriesHtSine200Response) Set(val *GetTimeSeriesHtSine200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHtSine200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHtSine200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHtSine200Response(val *GetTimeSeriesHtSine200Response) *NullableGetTimeSeriesHtSine200Response {
	return &NullableGetTimeSeriesHtSine200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHtSine200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHtSine200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
