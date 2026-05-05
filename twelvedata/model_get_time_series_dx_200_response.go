// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesDx200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesDx200Response{}

// GetTimeSeriesDx200Response struct for GetTimeSeriesDx200Response
type GetTimeSeriesDx200Response struct {
	Meta GetTimeSeriesDx200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesDx200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesDx200Response GetTimeSeriesDx200Response

// NewGetTimeSeriesDx200Response instantiates a new GetTimeSeriesDx200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesDx200Response(meta GetTimeSeriesDx200ResponseMeta, values []GetTimeSeriesDx200ResponseValuesInner, status string) *GetTimeSeriesDx200Response {
	this := GetTimeSeriesDx200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesDx200ResponseWithDefaults instantiates a new GetTimeSeriesDx200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesDx200ResponseWithDefaults() *GetTimeSeriesDx200Response {
	this := GetTimeSeriesDx200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesDx200Response) GetMeta() GetTimeSeriesDx200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesDx200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDx200Response) GetMetaOk() (*GetTimeSeriesDx200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesDx200Response) SetMeta(v GetTimeSeriesDx200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesDx200Response) GetValues() []GetTimeSeriesDx200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesDx200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDx200Response) GetValuesOk() ([]GetTimeSeriesDx200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesDx200Response) SetValues(v []GetTimeSeriesDx200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesDx200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDx200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesDx200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesDx200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesDx200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesDx200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesDx200Response := _GetTimeSeriesDx200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesDx200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesDx200Response(varGetTimeSeriesDx200Response)

	return err
}

type NullableGetTimeSeriesDx200Response struct {
	value *GetTimeSeriesDx200Response
	isSet bool
}

func (v NullableGetTimeSeriesDx200Response) Get() *GetTimeSeriesDx200Response {
	return v.value
}

func (v *NullableGetTimeSeriesDx200Response) Set(val *GetTimeSeriesDx200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesDx200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesDx200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesDx200Response(val *GetTimeSeriesDx200Response) *NullableGetTimeSeriesDx200Response {
	return &NullableGetTimeSeriesDx200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesDx200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesDx200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
