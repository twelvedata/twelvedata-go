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

// checks if the GetTimeSeriesNatr200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesNatr200Response{}

// GetTimeSeriesNatr200Response struct for GetTimeSeriesNatr200Response
type GetTimeSeriesNatr200Response struct {
	Meta GetTimeSeriesNatr200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesNatr200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesNatr200Response GetTimeSeriesNatr200Response

// NewGetTimeSeriesNatr200Response instantiates a new GetTimeSeriesNatr200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesNatr200Response(meta GetTimeSeriesNatr200ResponseMeta, values []GetTimeSeriesNatr200ResponseValuesInner, status string) *GetTimeSeriesNatr200Response {
	this := GetTimeSeriesNatr200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesNatr200ResponseWithDefaults instantiates a new GetTimeSeriesNatr200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesNatr200ResponseWithDefaults() *GetTimeSeriesNatr200Response {
	this := GetTimeSeriesNatr200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesNatr200Response) GetMeta() GetTimeSeriesNatr200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesNatr200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesNatr200Response) GetMetaOk() (*GetTimeSeriesNatr200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesNatr200Response) SetMeta(v GetTimeSeriesNatr200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesNatr200Response) GetValues() []GetTimeSeriesNatr200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesNatr200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesNatr200Response) GetValuesOk() ([]GetTimeSeriesNatr200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesNatr200Response) SetValues(v []GetTimeSeriesNatr200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesNatr200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesNatr200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesNatr200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesNatr200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesNatr200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesNatr200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesNatr200Response := _GetTimeSeriesNatr200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesNatr200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesNatr200Response(varGetTimeSeriesNatr200Response)

	return err
}

type NullableGetTimeSeriesNatr200Response struct {
	value *GetTimeSeriesNatr200Response
	isSet bool
}

func (v NullableGetTimeSeriesNatr200Response) Get() *GetTimeSeriesNatr200Response {
	return v.value
}

func (v *NullableGetTimeSeriesNatr200Response) Set(val *GetTimeSeriesNatr200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesNatr200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesNatr200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesNatr200Response(val *GetTimeSeriesNatr200Response) *NullableGetTimeSeriesNatr200Response {
	return &NullableGetTimeSeriesNatr200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesNatr200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesNatr200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
