// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesCmo200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCmo200Response{}

// GetTimeSeriesCmo200Response struct for GetTimeSeriesCmo200Response
type GetTimeSeriesCmo200Response struct {
	Meta GetTimeSeriesCmo200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesCmo200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesCmo200Response GetTimeSeriesCmo200Response

// NewGetTimeSeriesCmo200Response instantiates a new GetTimeSeriesCmo200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCmo200Response(meta GetTimeSeriesCmo200ResponseMeta, values []GetTimeSeriesCmo200ResponseValuesInner, status string) *GetTimeSeriesCmo200Response {
	this := GetTimeSeriesCmo200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesCmo200ResponseWithDefaults instantiates a new GetTimeSeriesCmo200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCmo200ResponseWithDefaults() *GetTimeSeriesCmo200Response {
	this := GetTimeSeriesCmo200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesCmo200Response) GetMeta() GetTimeSeriesCmo200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesCmo200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCmo200Response) GetMetaOk() (*GetTimeSeriesCmo200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesCmo200Response) SetMeta(v GetTimeSeriesCmo200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesCmo200Response) GetValues() []GetTimeSeriesCmo200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesCmo200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCmo200Response) GetValuesOk() ([]GetTimeSeriesCmo200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesCmo200Response) SetValues(v []GetTimeSeriesCmo200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesCmo200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCmo200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesCmo200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesCmo200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCmo200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesCmo200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesCmo200Response := _GetTimeSeriesCmo200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesCmo200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCmo200Response(varGetTimeSeriesCmo200Response)

	return err
}

type NullableGetTimeSeriesCmo200Response struct {
	value *GetTimeSeriesCmo200Response
	isSet bool
}

func (v NullableGetTimeSeriesCmo200Response) Get() *GetTimeSeriesCmo200Response {
	return v.value
}

func (v *NullableGetTimeSeriesCmo200Response) Set(val *GetTimeSeriesCmo200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCmo200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCmo200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCmo200Response(val *GetTimeSeriesCmo200Response) *NullableGetTimeSeriesCmo200Response {
	return &NullableGetTimeSeriesCmo200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCmo200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCmo200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
