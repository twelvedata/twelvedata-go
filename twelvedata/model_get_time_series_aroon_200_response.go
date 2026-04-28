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

// checks if the GetTimeSeriesAroon200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAroon200Response{}

// GetTimeSeriesAroon200Response struct for GetTimeSeriesAroon200Response
type GetTimeSeriesAroon200Response struct {
	Meta GetTimeSeriesAroon200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesAroon200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesAroon200Response GetTimeSeriesAroon200Response

// NewGetTimeSeriesAroon200Response instantiates a new GetTimeSeriesAroon200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAroon200Response(meta GetTimeSeriesAroon200ResponseMeta, values []GetTimeSeriesAroon200ResponseValuesInner, status string) *GetTimeSeriesAroon200Response {
	this := GetTimeSeriesAroon200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesAroon200ResponseWithDefaults instantiates a new GetTimeSeriesAroon200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAroon200ResponseWithDefaults() *GetTimeSeriesAroon200Response {
	this := GetTimeSeriesAroon200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesAroon200Response) GetMeta() GetTimeSeriesAroon200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesAroon200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroon200Response) GetMetaOk() (*GetTimeSeriesAroon200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesAroon200Response) SetMeta(v GetTimeSeriesAroon200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesAroon200Response) GetValues() []GetTimeSeriesAroon200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesAroon200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroon200Response) GetValuesOk() ([]GetTimeSeriesAroon200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesAroon200Response) SetValues(v []GetTimeSeriesAroon200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesAroon200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroon200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesAroon200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesAroon200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAroon200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesAroon200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAroon200Response := _GetTimeSeriesAroon200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesAroon200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAroon200Response(varGetTimeSeriesAroon200Response)

	return err
}

type NullableGetTimeSeriesAroon200Response struct {
	value *GetTimeSeriesAroon200Response
	isSet bool
}

func (v NullableGetTimeSeriesAroon200Response) Get() *GetTimeSeriesAroon200Response {
	return v.value
}

func (v *NullableGetTimeSeriesAroon200Response) Set(val *GetTimeSeriesAroon200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAroon200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAroon200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAroon200Response(val *GetTimeSeriesAroon200Response) *NullableGetTimeSeriesAroon200Response {
	return &NullableGetTimeSeriesAroon200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAroon200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAroon200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
