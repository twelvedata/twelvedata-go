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

// checks if the GetTimeSeriesLinearReg200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesLinearReg200Response{}

// GetTimeSeriesLinearReg200Response struct for GetTimeSeriesLinearReg200Response
type GetTimeSeriesLinearReg200Response struct {
	Meta GetTimeSeriesLinearReg200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesLinearReg200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesLinearReg200Response GetTimeSeriesLinearReg200Response

// NewGetTimeSeriesLinearReg200Response instantiates a new GetTimeSeriesLinearReg200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesLinearReg200Response(meta GetTimeSeriesLinearReg200ResponseMeta, values []GetTimeSeriesLinearReg200ResponseValuesInner, status string) *GetTimeSeriesLinearReg200Response {
	this := GetTimeSeriesLinearReg200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesLinearReg200ResponseWithDefaults instantiates a new GetTimeSeriesLinearReg200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesLinearReg200ResponseWithDefaults() *GetTimeSeriesLinearReg200Response {
	this := GetTimeSeriesLinearReg200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesLinearReg200Response) GetMeta() GetTimeSeriesLinearReg200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesLinearReg200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearReg200Response) GetMetaOk() (*GetTimeSeriesLinearReg200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesLinearReg200Response) SetMeta(v GetTimeSeriesLinearReg200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesLinearReg200Response) GetValues() []GetTimeSeriesLinearReg200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesLinearReg200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearReg200Response) GetValuesOk() ([]GetTimeSeriesLinearReg200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesLinearReg200Response) SetValues(v []GetTimeSeriesLinearReg200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesLinearReg200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearReg200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesLinearReg200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesLinearReg200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesLinearReg200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesLinearReg200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesLinearReg200Response := _GetTimeSeriesLinearReg200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesLinearReg200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesLinearReg200Response(varGetTimeSeriesLinearReg200Response)

	return err
}

type NullableGetTimeSeriesLinearReg200Response struct {
	value *GetTimeSeriesLinearReg200Response
	isSet bool
}

func (v NullableGetTimeSeriesLinearReg200Response) Get() *GetTimeSeriesLinearReg200Response {
	return v.value
}

func (v *NullableGetTimeSeriesLinearReg200Response) Set(val *GetTimeSeriesLinearReg200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesLinearReg200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesLinearReg200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesLinearReg200Response(val *GetTimeSeriesLinearReg200Response) *NullableGetTimeSeriesLinearReg200Response {
	return &NullableGetTimeSeriesLinearReg200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesLinearReg200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesLinearReg200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
