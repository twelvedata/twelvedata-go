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

// checks if the GetTimeSeriesExp200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesExp200Response{}

// GetTimeSeriesExp200Response struct for GetTimeSeriesExp200Response
type GetTimeSeriesExp200Response struct {
	Meta GetTimeSeriesExp200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesExp200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesExp200Response GetTimeSeriesExp200Response

// NewGetTimeSeriesExp200Response instantiates a new GetTimeSeriesExp200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesExp200Response(meta GetTimeSeriesExp200ResponseMeta, values []GetTimeSeriesExp200ResponseValuesInner, status string) *GetTimeSeriesExp200Response {
	this := GetTimeSeriesExp200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesExp200ResponseWithDefaults instantiates a new GetTimeSeriesExp200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesExp200ResponseWithDefaults() *GetTimeSeriesExp200Response {
	this := GetTimeSeriesExp200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesExp200Response) GetMeta() GetTimeSeriesExp200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesExp200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesExp200Response) GetMetaOk() (*GetTimeSeriesExp200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesExp200Response) SetMeta(v GetTimeSeriesExp200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesExp200Response) GetValues() []GetTimeSeriesExp200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesExp200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesExp200Response) GetValuesOk() ([]GetTimeSeriesExp200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesExp200Response) SetValues(v []GetTimeSeriesExp200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesExp200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesExp200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesExp200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesExp200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesExp200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesExp200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesExp200Response := _GetTimeSeriesExp200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesExp200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesExp200Response(varGetTimeSeriesExp200Response)

	return err
}

type NullableGetTimeSeriesExp200Response struct {
	value *GetTimeSeriesExp200Response
	isSet bool
}

func (v NullableGetTimeSeriesExp200Response) Get() *GetTimeSeriesExp200Response {
	return v.value
}

func (v *NullableGetTimeSeriesExp200Response) Set(val *GetTimeSeriesExp200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesExp200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesExp200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesExp200Response(val *GetTimeSeriesExp200Response) *NullableGetTimeSeriesExp200Response {
	return &NullableGetTimeSeriesExp200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesExp200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesExp200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
