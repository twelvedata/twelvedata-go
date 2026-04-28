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

// checks if the GetTimeSeriesHtTrendMode200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHtTrendMode200Response{}

// GetTimeSeriesHtTrendMode200Response struct for GetTimeSeriesHtTrendMode200Response
type GetTimeSeriesHtTrendMode200Response struct {
	Meta GetTimeSeriesHtTrendMode200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesHtTrendMode200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesHtTrendMode200Response GetTimeSeriesHtTrendMode200Response

// NewGetTimeSeriesHtTrendMode200Response instantiates a new GetTimeSeriesHtTrendMode200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHtTrendMode200Response(meta GetTimeSeriesHtTrendMode200ResponseMeta, values []GetTimeSeriesHtTrendMode200ResponseValuesInner, status string) *GetTimeSeriesHtTrendMode200Response {
	this := GetTimeSeriesHtTrendMode200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesHtTrendMode200ResponseWithDefaults instantiates a new GetTimeSeriesHtTrendMode200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHtTrendMode200ResponseWithDefaults() *GetTimeSeriesHtTrendMode200Response {
	this := GetTimeSeriesHtTrendMode200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesHtTrendMode200Response) GetMeta() GetTimeSeriesHtTrendMode200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesHtTrendMode200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtTrendMode200Response) GetMetaOk() (*GetTimeSeriesHtTrendMode200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesHtTrendMode200Response) SetMeta(v GetTimeSeriesHtTrendMode200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesHtTrendMode200Response) GetValues() []GetTimeSeriesHtTrendMode200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesHtTrendMode200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtTrendMode200Response) GetValuesOk() ([]GetTimeSeriesHtTrendMode200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesHtTrendMode200Response) SetValues(v []GetTimeSeriesHtTrendMode200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesHtTrendMode200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtTrendMode200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesHtTrendMode200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesHtTrendMode200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHtTrendMode200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesHtTrendMode200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesHtTrendMode200Response := _GetTimeSeriesHtTrendMode200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesHtTrendMode200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHtTrendMode200Response(varGetTimeSeriesHtTrendMode200Response)

	return err
}

type NullableGetTimeSeriesHtTrendMode200Response struct {
	value *GetTimeSeriesHtTrendMode200Response
	isSet bool
}

func (v NullableGetTimeSeriesHtTrendMode200Response) Get() *GetTimeSeriesHtTrendMode200Response {
	return v.value
}

func (v *NullableGetTimeSeriesHtTrendMode200Response) Set(val *GetTimeSeriesHtTrendMode200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHtTrendMode200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHtTrendMode200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHtTrendMode200Response(val *GetTimeSeriesHtTrendMode200Response) *NullableGetTimeSeriesHtTrendMode200Response {
	return &NullableGetTimeSeriesHtTrendMode200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHtTrendMode200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHtTrendMode200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
