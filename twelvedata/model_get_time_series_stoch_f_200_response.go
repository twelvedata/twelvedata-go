// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesStochF200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesStochF200Response{}

// GetTimeSeriesStochF200Response struct for GetTimeSeriesStochF200Response
type GetTimeSeriesStochF200Response struct {
	Meta GetTimeSeriesStochF200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesStochF200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesStochF200Response GetTimeSeriesStochF200Response

// NewGetTimeSeriesStochF200Response instantiates a new GetTimeSeriesStochF200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesStochF200Response(meta GetTimeSeriesStochF200ResponseMeta, values []GetTimeSeriesStochF200ResponseValuesInner, status string) *GetTimeSeriesStochF200Response {
	this := GetTimeSeriesStochF200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesStochF200ResponseWithDefaults instantiates a new GetTimeSeriesStochF200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesStochF200ResponseWithDefaults() *GetTimeSeriesStochF200Response {
	this := GetTimeSeriesStochF200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesStochF200Response) GetMeta() GetTimeSeriesStochF200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesStochF200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochF200Response) GetMetaOk() (*GetTimeSeriesStochF200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesStochF200Response) SetMeta(v GetTimeSeriesStochF200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesStochF200Response) GetValues() []GetTimeSeriesStochF200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesStochF200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochF200Response) GetValuesOk() ([]GetTimeSeriesStochF200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesStochF200Response) SetValues(v []GetTimeSeriesStochF200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesStochF200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochF200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesStochF200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesStochF200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesStochF200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesStochF200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesStochF200Response := _GetTimeSeriesStochF200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesStochF200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesStochF200Response(varGetTimeSeriesStochF200Response)

	return err
}

type NullableGetTimeSeriesStochF200Response struct {
	value *GetTimeSeriesStochF200Response
	isSet bool
}

func (v NullableGetTimeSeriesStochF200Response) Get() *GetTimeSeriesStochF200Response {
	return v.value
}

func (v *NullableGetTimeSeriesStochF200Response) Set(val *GetTimeSeriesStochF200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesStochF200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesStochF200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesStochF200Response(val *GetTimeSeriesStochF200Response) *NullableGetTimeSeriesStochF200Response {
	return &NullableGetTimeSeriesStochF200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesStochF200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesStochF200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
