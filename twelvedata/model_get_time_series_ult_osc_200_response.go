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

// checks if the GetTimeSeriesUltOsc200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesUltOsc200Response{}

// GetTimeSeriesUltOsc200Response struct for GetTimeSeriesUltOsc200Response
type GetTimeSeriesUltOsc200Response struct {
	Meta GetTimeSeriesUltOsc200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesUltOsc200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesUltOsc200Response GetTimeSeriesUltOsc200Response

// NewGetTimeSeriesUltOsc200Response instantiates a new GetTimeSeriesUltOsc200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesUltOsc200Response(meta GetTimeSeriesUltOsc200ResponseMeta, values []GetTimeSeriesUltOsc200ResponseValuesInner, status string) *GetTimeSeriesUltOsc200Response {
	this := GetTimeSeriesUltOsc200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesUltOsc200ResponseWithDefaults instantiates a new GetTimeSeriesUltOsc200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesUltOsc200ResponseWithDefaults() *GetTimeSeriesUltOsc200Response {
	this := GetTimeSeriesUltOsc200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesUltOsc200Response) GetMeta() GetTimeSeriesUltOsc200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesUltOsc200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesUltOsc200Response) GetMetaOk() (*GetTimeSeriesUltOsc200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesUltOsc200Response) SetMeta(v GetTimeSeriesUltOsc200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesUltOsc200Response) GetValues() []GetTimeSeriesUltOsc200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesUltOsc200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesUltOsc200Response) GetValuesOk() ([]GetTimeSeriesUltOsc200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesUltOsc200Response) SetValues(v []GetTimeSeriesUltOsc200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesUltOsc200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesUltOsc200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesUltOsc200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesUltOsc200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesUltOsc200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesUltOsc200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesUltOsc200Response := _GetTimeSeriesUltOsc200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesUltOsc200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesUltOsc200Response(varGetTimeSeriesUltOsc200Response)

	return err
}

type NullableGetTimeSeriesUltOsc200Response struct {
	value *GetTimeSeriesUltOsc200Response
	isSet bool
}

func (v NullableGetTimeSeriesUltOsc200Response) Get() *GetTimeSeriesUltOsc200Response {
	return v.value
}

func (v *NullableGetTimeSeriesUltOsc200Response) Set(val *GetTimeSeriesUltOsc200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesUltOsc200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesUltOsc200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesUltOsc200Response(val *GetTimeSeriesUltOsc200Response) *NullableGetTimeSeriesUltOsc200Response {
	return &NullableGetTimeSeriesUltOsc200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesUltOsc200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesUltOsc200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
