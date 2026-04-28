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

// checks if the GetTimeSeriesHlc3200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHlc3200Response{}

// GetTimeSeriesHlc3200Response struct for GetTimeSeriesHlc3200Response
type GetTimeSeriesHlc3200Response struct {
	Meta GetTimeSeriesHlc3200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesHlc3200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesHlc3200Response GetTimeSeriesHlc3200Response

// NewGetTimeSeriesHlc3200Response instantiates a new GetTimeSeriesHlc3200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHlc3200Response(meta GetTimeSeriesHlc3200ResponseMeta, values []GetTimeSeriesHlc3200ResponseValuesInner, status string) *GetTimeSeriesHlc3200Response {
	this := GetTimeSeriesHlc3200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesHlc3200ResponseWithDefaults instantiates a new GetTimeSeriesHlc3200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHlc3200ResponseWithDefaults() *GetTimeSeriesHlc3200Response {
	this := GetTimeSeriesHlc3200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesHlc3200Response) GetMeta() GetTimeSeriesHlc3200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesHlc3200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHlc3200Response) GetMetaOk() (*GetTimeSeriesHlc3200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesHlc3200Response) SetMeta(v GetTimeSeriesHlc3200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesHlc3200Response) GetValues() []GetTimeSeriesHlc3200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesHlc3200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHlc3200Response) GetValuesOk() ([]GetTimeSeriesHlc3200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesHlc3200Response) SetValues(v []GetTimeSeriesHlc3200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesHlc3200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHlc3200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesHlc3200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesHlc3200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHlc3200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesHlc3200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesHlc3200Response := _GetTimeSeriesHlc3200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesHlc3200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHlc3200Response(varGetTimeSeriesHlc3200Response)

	return err
}

type NullableGetTimeSeriesHlc3200Response struct {
	value *GetTimeSeriesHlc3200Response
	isSet bool
}

func (v NullableGetTimeSeriesHlc3200Response) Get() *GetTimeSeriesHlc3200Response {
	return v.value
}

func (v *NullableGetTimeSeriesHlc3200Response) Set(val *GetTimeSeriesHlc3200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHlc3200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHlc3200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHlc3200Response(val *GetTimeSeriesHlc3200Response) *NullableGetTimeSeriesHlc3200Response {
	return &NullableGetTimeSeriesHlc3200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHlc3200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHlc3200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
