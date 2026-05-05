// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesSma200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSma200Response{}

// GetTimeSeriesSma200Response struct for GetTimeSeriesSma200Response
type GetTimeSeriesSma200Response struct {
	Meta GetTimeSeriesSma200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesSma200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesSma200Response GetTimeSeriesSma200Response

// NewGetTimeSeriesSma200Response instantiates a new GetTimeSeriesSma200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSma200Response(meta GetTimeSeriesSma200ResponseMeta, values []GetTimeSeriesSma200ResponseValuesInner, status string) *GetTimeSeriesSma200Response {
	this := GetTimeSeriesSma200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesSma200ResponseWithDefaults instantiates a new GetTimeSeriesSma200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSma200ResponseWithDefaults() *GetTimeSeriesSma200Response {
	this := GetTimeSeriesSma200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesSma200Response) GetMeta() GetTimeSeriesSma200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesSma200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSma200Response) GetMetaOk() (*GetTimeSeriesSma200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesSma200Response) SetMeta(v GetTimeSeriesSma200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesSma200Response) GetValues() []GetTimeSeriesSma200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesSma200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSma200Response) GetValuesOk() ([]GetTimeSeriesSma200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesSma200Response) SetValues(v []GetTimeSeriesSma200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesSma200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSma200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesSma200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesSma200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSma200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesSma200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesSma200Response := _GetTimeSeriesSma200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSma200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSma200Response(varGetTimeSeriesSma200Response)

	return err
}

type NullableGetTimeSeriesSma200Response struct {
	value *GetTimeSeriesSma200Response
	isSet bool
}

func (v NullableGetTimeSeriesSma200Response) Get() *GetTimeSeriesSma200Response {
	return v.value
}

func (v *NullableGetTimeSeriesSma200Response) Set(val *GetTimeSeriesSma200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSma200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSma200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSma200Response(val *GetTimeSeriesSma200Response) *NullableGetTimeSeriesSma200Response {
	return &NullableGetTimeSeriesSma200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSma200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSma200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
