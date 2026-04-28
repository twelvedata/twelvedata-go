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

// checks if the GetTimeSeriesMom200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMom200Response{}

// GetTimeSeriesMom200Response struct for GetTimeSeriesMom200Response
type GetTimeSeriesMom200Response struct {
	Meta GetTimeSeriesMom200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMom200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMom200Response GetTimeSeriesMom200Response

// NewGetTimeSeriesMom200Response instantiates a new GetTimeSeriesMom200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMom200Response(meta GetTimeSeriesMom200ResponseMeta, values []GetTimeSeriesMom200ResponseValuesInner, status string) *GetTimeSeriesMom200Response {
	this := GetTimeSeriesMom200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMom200ResponseWithDefaults instantiates a new GetTimeSeriesMom200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMom200ResponseWithDefaults() *GetTimeSeriesMom200Response {
	this := GetTimeSeriesMom200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMom200Response) GetMeta() GetTimeSeriesMom200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMom200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMom200Response) GetMetaOk() (*GetTimeSeriesMom200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMom200Response) SetMeta(v GetTimeSeriesMom200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMom200Response) GetValues() []GetTimeSeriesMom200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMom200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMom200Response) GetValuesOk() ([]GetTimeSeriesMom200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMom200Response) SetValues(v []GetTimeSeriesMom200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMom200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMom200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMom200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMom200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMom200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMom200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMom200Response := _GetTimeSeriesMom200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMom200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMom200Response(varGetTimeSeriesMom200Response)

	return err
}

type NullableGetTimeSeriesMom200Response struct {
	value *GetTimeSeriesMom200Response
	isSet bool
}

func (v NullableGetTimeSeriesMom200Response) Get() *GetTimeSeriesMom200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMom200Response) Set(val *GetTimeSeriesMom200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMom200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMom200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMom200Response(val *GetTimeSeriesMom200Response) *NullableGetTimeSeriesMom200Response {
	return &NullableGetTimeSeriesMom200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMom200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMom200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
