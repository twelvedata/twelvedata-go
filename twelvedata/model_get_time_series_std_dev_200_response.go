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

// checks if the GetTimeSeriesStdDev200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesStdDev200Response{}

// GetTimeSeriesStdDev200Response struct for GetTimeSeriesStdDev200Response
type GetTimeSeriesStdDev200Response struct {
	Meta GetTimeSeriesStdDev200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesStdDev200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesStdDev200Response GetTimeSeriesStdDev200Response

// NewGetTimeSeriesStdDev200Response instantiates a new GetTimeSeriesStdDev200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesStdDev200Response(meta GetTimeSeriesStdDev200ResponseMeta, values []GetTimeSeriesStdDev200ResponseValuesInner, status string) *GetTimeSeriesStdDev200Response {
	this := GetTimeSeriesStdDev200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesStdDev200ResponseWithDefaults instantiates a new GetTimeSeriesStdDev200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesStdDev200ResponseWithDefaults() *GetTimeSeriesStdDev200Response {
	this := GetTimeSeriesStdDev200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesStdDev200Response) GetMeta() GetTimeSeriesStdDev200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesStdDev200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStdDev200Response) GetMetaOk() (*GetTimeSeriesStdDev200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesStdDev200Response) SetMeta(v GetTimeSeriesStdDev200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesStdDev200Response) GetValues() []GetTimeSeriesStdDev200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesStdDev200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStdDev200Response) GetValuesOk() ([]GetTimeSeriesStdDev200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesStdDev200Response) SetValues(v []GetTimeSeriesStdDev200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesStdDev200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStdDev200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesStdDev200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesStdDev200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesStdDev200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesStdDev200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesStdDev200Response := _GetTimeSeriesStdDev200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesStdDev200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesStdDev200Response(varGetTimeSeriesStdDev200Response)

	return err
}

type NullableGetTimeSeriesStdDev200Response struct {
	value *GetTimeSeriesStdDev200Response
	isSet bool
}

func (v NullableGetTimeSeriesStdDev200Response) Get() *GetTimeSeriesStdDev200Response {
	return v.value
}

func (v *NullableGetTimeSeriesStdDev200Response) Set(val *GetTimeSeriesStdDev200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesStdDev200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesStdDev200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesStdDev200Response(val *GetTimeSeriesStdDev200Response) *NullableGetTimeSeriesStdDev200Response {
	return &NullableGetTimeSeriesStdDev200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesStdDev200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesStdDev200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
