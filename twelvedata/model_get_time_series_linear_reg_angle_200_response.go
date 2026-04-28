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

// checks if the GetTimeSeriesLinearRegAngle200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesLinearRegAngle200Response{}

// GetTimeSeriesLinearRegAngle200Response struct for GetTimeSeriesLinearRegAngle200Response
type GetTimeSeriesLinearRegAngle200Response struct {
	Meta GetTimeSeriesLinearRegAngle200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesLinearRegAngle200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesLinearRegAngle200Response GetTimeSeriesLinearRegAngle200Response

// NewGetTimeSeriesLinearRegAngle200Response instantiates a new GetTimeSeriesLinearRegAngle200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesLinearRegAngle200Response(meta GetTimeSeriesLinearRegAngle200ResponseMeta, values []GetTimeSeriesLinearRegAngle200ResponseValuesInner, status string) *GetTimeSeriesLinearRegAngle200Response {
	this := GetTimeSeriesLinearRegAngle200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesLinearRegAngle200ResponseWithDefaults instantiates a new GetTimeSeriesLinearRegAngle200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesLinearRegAngle200ResponseWithDefaults() *GetTimeSeriesLinearRegAngle200Response {
	this := GetTimeSeriesLinearRegAngle200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesLinearRegAngle200Response) GetMeta() GetTimeSeriesLinearRegAngle200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesLinearRegAngle200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegAngle200Response) GetMetaOk() (*GetTimeSeriesLinearRegAngle200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesLinearRegAngle200Response) SetMeta(v GetTimeSeriesLinearRegAngle200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesLinearRegAngle200Response) GetValues() []GetTimeSeriesLinearRegAngle200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesLinearRegAngle200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegAngle200Response) GetValuesOk() ([]GetTimeSeriesLinearRegAngle200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesLinearRegAngle200Response) SetValues(v []GetTimeSeriesLinearRegAngle200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesLinearRegAngle200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegAngle200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesLinearRegAngle200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesLinearRegAngle200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesLinearRegAngle200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesLinearRegAngle200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesLinearRegAngle200Response := _GetTimeSeriesLinearRegAngle200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesLinearRegAngle200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesLinearRegAngle200Response(varGetTimeSeriesLinearRegAngle200Response)

	return err
}

type NullableGetTimeSeriesLinearRegAngle200Response struct {
	value *GetTimeSeriesLinearRegAngle200Response
	isSet bool
}

func (v NullableGetTimeSeriesLinearRegAngle200Response) Get() *GetTimeSeriesLinearRegAngle200Response {
	return v.value
}

func (v *NullableGetTimeSeriesLinearRegAngle200Response) Set(val *GetTimeSeriesLinearRegAngle200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesLinearRegAngle200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesLinearRegAngle200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesLinearRegAngle200Response(val *GetTimeSeriesLinearRegAngle200Response) *NullableGetTimeSeriesLinearRegAngle200Response {
	return &NullableGetTimeSeriesLinearRegAngle200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesLinearRegAngle200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesLinearRegAngle200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
