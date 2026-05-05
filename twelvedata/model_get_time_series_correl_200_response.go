// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesCorrel200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCorrel200Response{}

// GetTimeSeriesCorrel200Response struct for GetTimeSeriesCorrel200Response
type GetTimeSeriesCorrel200Response struct {
	Meta GetTimeSeriesCorrel200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesCorrel200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesCorrel200Response GetTimeSeriesCorrel200Response

// NewGetTimeSeriesCorrel200Response instantiates a new GetTimeSeriesCorrel200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCorrel200Response(meta GetTimeSeriesCorrel200ResponseMeta, values []GetTimeSeriesCorrel200ResponseValuesInner, status string) *GetTimeSeriesCorrel200Response {
	this := GetTimeSeriesCorrel200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesCorrel200ResponseWithDefaults instantiates a new GetTimeSeriesCorrel200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCorrel200ResponseWithDefaults() *GetTimeSeriesCorrel200Response {
	this := GetTimeSeriesCorrel200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesCorrel200Response) GetMeta() GetTimeSeriesCorrel200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesCorrel200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCorrel200Response) GetMetaOk() (*GetTimeSeriesCorrel200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesCorrel200Response) SetMeta(v GetTimeSeriesCorrel200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesCorrel200Response) GetValues() []GetTimeSeriesCorrel200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesCorrel200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCorrel200Response) GetValuesOk() ([]GetTimeSeriesCorrel200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesCorrel200Response) SetValues(v []GetTimeSeriesCorrel200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesCorrel200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCorrel200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesCorrel200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesCorrel200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCorrel200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesCorrel200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesCorrel200Response := _GetTimeSeriesCorrel200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesCorrel200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCorrel200Response(varGetTimeSeriesCorrel200Response)

	return err
}

type NullableGetTimeSeriesCorrel200Response struct {
	value *GetTimeSeriesCorrel200Response
	isSet bool
}

func (v NullableGetTimeSeriesCorrel200Response) Get() *GetTimeSeriesCorrel200Response {
	return v.value
}

func (v *NullableGetTimeSeriesCorrel200Response) Set(val *GetTimeSeriesCorrel200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCorrel200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCorrel200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCorrel200Response(val *GetTimeSeriesCorrel200Response) *NullableGetTimeSeriesCorrel200Response {
	return &NullableGetTimeSeriesCorrel200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCorrel200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCorrel200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
