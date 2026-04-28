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

// checks if the GetTimeSeriesVwap200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesVwap200Response{}

// GetTimeSeriesVwap200Response struct for GetTimeSeriesVwap200Response
type GetTimeSeriesVwap200Response struct {
	Meta GetTimeSeriesVwap200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesVwap200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesVwap200Response GetTimeSeriesVwap200Response

// NewGetTimeSeriesVwap200Response instantiates a new GetTimeSeriesVwap200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesVwap200Response(meta GetTimeSeriesVwap200ResponseMeta, values []GetTimeSeriesVwap200ResponseValuesInner, status string) *GetTimeSeriesVwap200Response {
	this := GetTimeSeriesVwap200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesVwap200ResponseWithDefaults instantiates a new GetTimeSeriesVwap200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesVwap200ResponseWithDefaults() *GetTimeSeriesVwap200Response {
	this := GetTimeSeriesVwap200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesVwap200Response) GetMeta() GetTimeSeriesVwap200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesVwap200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVwap200Response) GetMetaOk() (*GetTimeSeriesVwap200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesVwap200Response) SetMeta(v GetTimeSeriesVwap200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesVwap200Response) GetValues() []GetTimeSeriesVwap200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesVwap200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVwap200Response) GetValuesOk() ([]GetTimeSeriesVwap200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesVwap200Response) SetValues(v []GetTimeSeriesVwap200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesVwap200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVwap200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesVwap200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesVwap200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesVwap200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesVwap200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesVwap200Response := _GetTimeSeriesVwap200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesVwap200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesVwap200Response(varGetTimeSeriesVwap200Response)

	return err
}

type NullableGetTimeSeriesVwap200Response struct {
	value *GetTimeSeriesVwap200Response
	isSet bool
}

func (v NullableGetTimeSeriesVwap200Response) Get() *GetTimeSeriesVwap200Response {
	return v.value
}

func (v *NullableGetTimeSeriesVwap200Response) Set(val *GetTimeSeriesVwap200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesVwap200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesVwap200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesVwap200Response(val *GetTimeSeriesVwap200Response) *NullableGetTimeSeriesVwap200Response {
	return &NullableGetTimeSeriesVwap200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesVwap200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesVwap200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
