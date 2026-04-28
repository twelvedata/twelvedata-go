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

// checks if the GetTimeSeriesHtPhasor200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHtPhasor200Response{}

// GetTimeSeriesHtPhasor200Response struct for GetTimeSeriesHtPhasor200Response
type GetTimeSeriesHtPhasor200Response struct {
	Meta GetTimeSeriesHtPhasor200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesHtPhasor200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesHtPhasor200Response GetTimeSeriesHtPhasor200Response

// NewGetTimeSeriesHtPhasor200Response instantiates a new GetTimeSeriesHtPhasor200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHtPhasor200Response(meta GetTimeSeriesHtPhasor200ResponseMeta, values []GetTimeSeriesHtPhasor200ResponseValuesInner, status string) *GetTimeSeriesHtPhasor200Response {
	this := GetTimeSeriesHtPhasor200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesHtPhasor200ResponseWithDefaults instantiates a new GetTimeSeriesHtPhasor200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHtPhasor200ResponseWithDefaults() *GetTimeSeriesHtPhasor200Response {
	this := GetTimeSeriesHtPhasor200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesHtPhasor200Response) GetMeta() GetTimeSeriesHtPhasor200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesHtPhasor200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtPhasor200Response) GetMetaOk() (*GetTimeSeriesHtPhasor200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesHtPhasor200Response) SetMeta(v GetTimeSeriesHtPhasor200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesHtPhasor200Response) GetValues() []GetTimeSeriesHtPhasor200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesHtPhasor200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtPhasor200Response) GetValuesOk() ([]GetTimeSeriesHtPhasor200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesHtPhasor200Response) SetValues(v []GetTimeSeriesHtPhasor200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesHtPhasor200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtPhasor200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesHtPhasor200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesHtPhasor200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHtPhasor200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesHtPhasor200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesHtPhasor200Response := _GetTimeSeriesHtPhasor200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesHtPhasor200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHtPhasor200Response(varGetTimeSeriesHtPhasor200Response)

	return err
}

type NullableGetTimeSeriesHtPhasor200Response struct {
	value *GetTimeSeriesHtPhasor200Response
	isSet bool
}

func (v NullableGetTimeSeriesHtPhasor200Response) Get() *GetTimeSeriesHtPhasor200Response {
	return v.value
}

func (v *NullableGetTimeSeriesHtPhasor200Response) Set(val *GetTimeSeriesHtPhasor200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHtPhasor200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHtPhasor200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHtPhasor200Response(val *GetTimeSeriesHtPhasor200Response) *NullableGetTimeSeriesHtPhasor200Response {
	return &NullableGetTimeSeriesHtPhasor200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHtPhasor200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHtPhasor200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
