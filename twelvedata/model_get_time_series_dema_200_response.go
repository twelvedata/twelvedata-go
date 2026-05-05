// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesDema200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesDema200Response{}

// GetTimeSeriesDema200Response struct for GetTimeSeriesDema200Response
type GetTimeSeriesDema200Response struct {
	Meta GetTimeSeriesDema200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesDema200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesDema200Response GetTimeSeriesDema200Response

// NewGetTimeSeriesDema200Response instantiates a new GetTimeSeriesDema200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesDema200Response(meta GetTimeSeriesDema200ResponseMeta, values []GetTimeSeriesDema200ResponseValuesInner, status string) *GetTimeSeriesDema200Response {
	this := GetTimeSeriesDema200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesDema200ResponseWithDefaults instantiates a new GetTimeSeriesDema200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesDema200ResponseWithDefaults() *GetTimeSeriesDema200Response {
	this := GetTimeSeriesDema200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesDema200Response) GetMeta() GetTimeSeriesDema200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesDema200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDema200Response) GetMetaOk() (*GetTimeSeriesDema200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesDema200Response) SetMeta(v GetTimeSeriesDema200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesDema200Response) GetValues() []GetTimeSeriesDema200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesDema200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDema200Response) GetValuesOk() ([]GetTimeSeriesDema200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesDema200Response) SetValues(v []GetTimeSeriesDema200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesDema200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDema200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesDema200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesDema200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesDema200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesDema200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesDema200Response := _GetTimeSeriesDema200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesDema200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesDema200Response(varGetTimeSeriesDema200Response)

	return err
}

type NullableGetTimeSeriesDema200Response struct {
	value *GetTimeSeriesDema200Response
	isSet bool
}

func (v NullableGetTimeSeriesDema200Response) Get() *GetTimeSeriesDema200Response {
	return v.value
}

func (v *NullableGetTimeSeriesDema200Response) Set(val *GetTimeSeriesDema200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesDema200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesDema200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesDema200Response(val *GetTimeSeriesDema200Response) *NullableGetTimeSeriesDema200Response {
	return &NullableGetTimeSeriesDema200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesDema200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesDema200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
