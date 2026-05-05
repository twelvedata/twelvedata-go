// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesLog10200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesLog10200Response{}

// GetTimeSeriesLog10200Response struct for GetTimeSeriesLog10200Response
type GetTimeSeriesLog10200Response struct {
	Meta GetTimeSeriesLog10200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesLog10200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesLog10200Response GetTimeSeriesLog10200Response

// NewGetTimeSeriesLog10200Response instantiates a new GetTimeSeriesLog10200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesLog10200Response(meta GetTimeSeriesLog10200ResponseMeta, values []GetTimeSeriesLog10200ResponseValuesInner, status string) *GetTimeSeriesLog10200Response {
	this := GetTimeSeriesLog10200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesLog10200ResponseWithDefaults instantiates a new GetTimeSeriesLog10200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesLog10200ResponseWithDefaults() *GetTimeSeriesLog10200Response {
	this := GetTimeSeriesLog10200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesLog10200Response) GetMeta() GetTimeSeriesLog10200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesLog10200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLog10200Response) GetMetaOk() (*GetTimeSeriesLog10200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesLog10200Response) SetMeta(v GetTimeSeriesLog10200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesLog10200Response) GetValues() []GetTimeSeriesLog10200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesLog10200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLog10200Response) GetValuesOk() ([]GetTimeSeriesLog10200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesLog10200Response) SetValues(v []GetTimeSeriesLog10200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesLog10200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLog10200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesLog10200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesLog10200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesLog10200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesLog10200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesLog10200Response := _GetTimeSeriesLog10200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesLog10200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesLog10200Response(varGetTimeSeriesLog10200Response)

	return err
}

type NullableGetTimeSeriesLog10200Response struct {
	value *GetTimeSeriesLog10200Response
	isSet bool
}

func (v NullableGetTimeSeriesLog10200Response) Get() *GetTimeSeriesLog10200Response {
	return v.value
}

func (v *NullableGetTimeSeriesLog10200Response) Set(val *GetTimeSeriesLog10200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesLog10200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesLog10200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesLog10200Response(val *GetTimeSeriesLog10200Response) *NullableGetTimeSeriesLog10200Response {
	return &NullableGetTimeSeriesLog10200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesLog10200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesLog10200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
