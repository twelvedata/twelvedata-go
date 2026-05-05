// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesBeta200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesBeta200Response{}

// GetTimeSeriesBeta200Response struct for GetTimeSeriesBeta200Response
type GetTimeSeriesBeta200Response struct {
	Meta GetTimeSeriesBeta200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesBeta200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesBeta200Response GetTimeSeriesBeta200Response

// NewGetTimeSeriesBeta200Response instantiates a new GetTimeSeriesBeta200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesBeta200Response(meta GetTimeSeriesBeta200ResponseMeta, values []GetTimeSeriesBeta200ResponseValuesInner, status string) *GetTimeSeriesBeta200Response {
	this := GetTimeSeriesBeta200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesBeta200ResponseWithDefaults instantiates a new GetTimeSeriesBeta200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesBeta200ResponseWithDefaults() *GetTimeSeriesBeta200Response {
	this := GetTimeSeriesBeta200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesBeta200Response) GetMeta() GetTimeSeriesBeta200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesBeta200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBeta200Response) GetMetaOk() (*GetTimeSeriesBeta200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesBeta200Response) SetMeta(v GetTimeSeriesBeta200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesBeta200Response) GetValues() []GetTimeSeriesBeta200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesBeta200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBeta200Response) GetValuesOk() ([]GetTimeSeriesBeta200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesBeta200Response) SetValues(v []GetTimeSeriesBeta200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesBeta200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBeta200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesBeta200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesBeta200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesBeta200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesBeta200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesBeta200Response := _GetTimeSeriesBeta200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesBeta200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesBeta200Response(varGetTimeSeriesBeta200Response)

	return err
}

type NullableGetTimeSeriesBeta200Response struct {
	value *GetTimeSeriesBeta200Response
	isSet bool
}

func (v NullableGetTimeSeriesBeta200Response) Get() *GetTimeSeriesBeta200Response {
	return v.value
}

func (v *NullableGetTimeSeriesBeta200Response) Set(val *GetTimeSeriesBeta200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesBeta200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesBeta200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesBeta200Response(val *GetTimeSeriesBeta200Response) *NullableGetTimeSeriesBeta200Response {
	return &NullableGetTimeSeriesBeta200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesBeta200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesBeta200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
