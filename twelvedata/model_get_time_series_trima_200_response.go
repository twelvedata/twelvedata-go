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

// checks if the GetTimeSeriesTrima200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesTrima200Response{}

// GetTimeSeriesTrima200Response struct for GetTimeSeriesTrima200Response
type GetTimeSeriesTrima200Response struct {
	Meta GetTimeSeriesTrima200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesTrima200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesTrima200Response GetTimeSeriesTrima200Response

// NewGetTimeSeriesTrima200Response instantiates a new GetTimeSeriesTrima200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesTrima200Response(meta GetTimeSeriesTrima200ResponseMeta, values []GetTimeSeriesTrima200ResponseValuesInner, status string) *GetTimeSeriesTrima200Response {
	this := GetTimeSeriesTrima200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesTrima200ResponseWithDefaults instantiates a new GetTimeSeriesTrima200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesTrima200ResponseWithDefaults() *GetTimeSeriesTrima200Response {
	this := GetTimeSeriesTrima200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesTrima200Response) GetMeta() GetTimeSeriesTrima200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesTrima200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTrima200Response) GetMetaOk() (*GetTimeSeriesTrima200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesTrima200Response) SetMeta(v GetTimeSeriesTrima200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesTrima200Response) GetValues() []GetTimeSeriesTrima200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesTrima200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTrima200Response) GetValuesOk() ([]GetTimeSeriesTrima200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesTrima200Response) SetValues(v []GetTimeSeriesTrima200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesTrima200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTrima200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesTrima200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesTrima200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesTrima200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesTrima200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesTrima200Response := _GetTimeSeriesTrima200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesTrima200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesTrima200Response(varGetTimeSeriesTrima200Response)

	return err
}

type NullableGetTimeSeriesTrima200Response struct {
	value *GetTimeSeriesTrima200Response
	isSet bool
}

func (v NullableGetTimeSeriesTrima200Response) Get() *GetTimeSeriesTrima200Response {
	return v.value
}

func (v *NullableGetTimeSeriesTrima200Response) Set(val *GetTimeSeriesTrima200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesTrima200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesTrima200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesTrima200Response(val *GetTimeSeriesTrima200Response) *NullableGetTimeSeriesTrima200Response {
	return &NullableGetTimeSeriesTrima200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesTrima200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesTrima200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
