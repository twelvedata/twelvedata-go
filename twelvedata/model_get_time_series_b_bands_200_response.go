// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesBBands200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesBBands200Response{}

// GetTimeSeriesBBands200Response struct for GetTimeSeriesBBands200Response
type GetTimeSeriesBBands200Response struct {
	Meta GetTimeSeriesBBands200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesBBands200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesBBands200Response GetTimeSeriesBBands200Response

// NewGetTimeSeriesBBands200Response instantiates a new GetTimeSeriesBBands200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesBBands200Response(meta GetTimeSeriesBBands200ResponseMeta, values []GetTimeSeriesBBands200ResponseValuesInner, status string) *GetTimeSeriesBBands200Response {
	this := GetTimeSeriesBBands200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesBBands200ResponseWithDefaults instantiates a new GetTimeSeriesBBands200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesBBands200ResponseWithDefaults() *GetTimeSeriesBBands200Response {
	this := GetTimeSeriesBBands200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesBBands200Response) GetMeta() GetTimeSeriesBBands200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesBBands200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBBands200Response) GetMetaOk() (*GetTimeSeriesBBands200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesBBands200Response) SetMeta(v GetTimeSeriesBBands200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesBBands200Response) GetValues() []GetTimeSeriesBBands200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesBBands200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBBands200Response) GetValuesOk() ([]GetTimeSeriesBBands200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesBBands200Response) SetValues(v []GetTimeSeriesBBands200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesBBands200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBBands200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesBBands200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesBBands200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesBBands200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesBBands200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesBBands200Response := _GetTimeSeriesBBands200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesBBands200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesBBands200Response(varGetTimeSeriesBBands200Response)

	return err
}

type NullableGetTimeSeriesBBands200Response struct {
	value *GetTimeSeriesBBands200Response
	isSet bool
}

func (v NullableGetTimeSeriesBBands200Response) Get() *GetTimeSeriesBBands200Response {
	return v.value
}

func (v *NullableGetTimeSeriesBBands200Response) Set(val *GetTimeSeriesBBands200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesBBands200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesBBands200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesBBands200Response(val *GetTimeSeriesBBands200Response) *NullableGetTimeSeriesBBands200Response {
	return &NullableGetTimeSeriesBBands200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesBBands200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesBBands200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
