// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesSarExt200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSarExt200Response{}

// GetTimeSeriesSarExt200Response struct for GetTimeSeriesSarExt200Response
type GetTimeSeriesSarExt200Response struct {
	Meta GetTimeSeriesSarExt200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesSarExt200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesSarExt200Response GetTimeSeriesSarExt200Response

// NewGetTimeSeriesSarExt200Response instantiates a new GetTimeSeriesSarExt200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSarExt200Response(meta GetTimeSeriesSarExt200ResponseMeta, values []GetTimeSeriesSarExt200ResponseValuesInner, status string) *GetTimeSeriesSarExt200Response {
	this := GetTimeSeriesSarExt200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesSarExt200ResponseWithDefaults instantiates a new GetTimeSeriesSarExt200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSarExt200ResponseWithDefaults() *GetTimeSeriesSarExt200Response {
	this := GetTimeSeriesSarExt200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesSarExt200Response) GetMeta() GetTimeSeriesSarExt200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesSarExt200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200Response) GetMetaOk() (*GetTimeSeriesSarExt200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesSarExt200Response) SetMeta(v GetTimeSeriesSarExt200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesSarExt200Response) GetValues() []GetTimeSeriesSarExt200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesSarExt200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200Response) GetValuesOk() ([]GetTimeSeriesSarExt200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesSarExt200Response) SetValues(v []GetTimeSeriesSarExt200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesSarExt200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesSarExt200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesSarExt200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSarExt200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesSarExt200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesSarExt200Response := _GetTimeSeriesSarExt200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSarExt200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSarExt200Response(varGetTimeSeriesSarExt200Response)

	return err
}

type NullableGetTimeSeriesSarExt200Response struct {
	value *GetTimeSeriesSarExt200Response
	isSet bool
}

func (v NullableGetTimeSeriesSarExt200Response) Get() *GetTimeSeriesSarExt200Response {
	return v.value
}

func (v *NullableGetTimeSeriesSarExt200Response) Set(val *GetTimeSeriesSarExt200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSarExt200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSarExt200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSarExt200Response(val *GetTimeSeriesSarExt200Response) *NullableGetTimeSeriesSarExt200Response {
	return &NullableGetTimeSeriesSarExt200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSarExt200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSarExt200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
