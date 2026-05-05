// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesSuperTrendHeikinAshiCandles200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSuperTrendHeikinAshiCandles200Response{}

// GetTimeSeriesSuperTrendHeikinAshiCandles200Response struct for GetTimeSeriesSuperTrendHeikinAshiCandles200Response
type GetTimeSeriesSuperTrendHeikinAshiCandles200Response struct {
	Meta GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesSuperTrendHeikinAshiCandles200Response GetTimeSeriesSuperTrendHeikinAshiCandles200Response

// NewGetTimeSeriesSuperTrendHeikinAshiCandles200Response instantiates a new GetTimeSeriesSuperTrendHeikinAshiCandles200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSuperTrendHeikinAshiCandles200Response(meta GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMeta, values []GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner, status string) *GetTimeSeriesSuperTrendHeikinAshiCandles200Response {
	this := GetTimeSeriesSuperTrendHeikinAshiCandles200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseWithDefaults instantiates a new GetTimeSeriesSuperTrendHeikinAshiCandles200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseWithDefaults() *GetTimeSeriesSuperTrendHeikinAshiCandles200Response {
	this := GetTimeSeriesSuperTrendHeikinAshiCandles200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) GetMeta() GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) GetMetaOk() (*GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) SetMeta(v GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) GetValues() []GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) GetValuesOk() ([]GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) SetValues(v []GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesSuperTrendHeikinAshiCandles200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSuperTrendHeikinAshiCandles200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesSuperTrendHeikinAshiCandles200Response := _GetTimeSeriesSuperTrendHeikinAshiCandles200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSuperTrendHeikinAshiCandles200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSuperTrendHeikinAshiCandles200Response(varGetTimeSeriesSuperTrendHeikinAshiCandles200Response)

	return err
}

type NullableGetTimeSeriesSuperTrendHeikinAshiCandles200Response struct {
	value *GetTimeSeriesSuperTrendHeikinAshiCandles200Response
	isSet bool
}

func (v NullableGetTimeSeriesSuperTrendHeikinAshiCandles200Response) Get() *GetTimeSeriesSuperTrendHeikinAshiCandles200Response {
	return v.value
}

func (v *NullableGetTimeSeriesSuperTrendHeikinAshiCandles200Response) Set(val *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSuperTrendHeikinAshiCandles200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSuperTrendHeikinAshiCandles200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSuperTrendHeikinAshiCandles200Response(val *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) *NullableGetTimeSeriesSuperTrendHeikinAshiCandles200Response {
	return &NullableGetTimeSeriesSuperTrendHeikinAshiCandles200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSuperTrendHeikinAshiCandles200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSuperTrendHeikinAshiCandles200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
