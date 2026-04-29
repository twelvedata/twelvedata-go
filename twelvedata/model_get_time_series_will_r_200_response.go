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

// checks if the GetTimeSeriesWillR200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesWillR200Response{}

// GetTimeSeriesWillR200Response struct for GetTimeSeriesWillR200Response
type GetTimeSeriesWillR200Response struct {
	Meta GetTimeSeriesWillR200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesWillR200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesWillR200Response GetTimeSeriesWillR200Response

// NewGetTimeSeriesWillR200Response instantiates a new GetTimeSeriesWillR200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesWillR200Response(meta GetTimeSeriesWillR200ResponseMeta, values []GetTimeSeriesWillR200ResponseValuesInner, status string) *GetTimeSeriesWillR200Response {
	this := GetTimeSeriesWillR200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesWillR200ResponseWithDefaults instantiates a new GetTimeSeriesWillR200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesWillR200ResponseWithDefaults() *GetTimeSeriesWillR200Response {
	this := GetTimeSeriesWillR200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesWillR200Response) GetMeta() GetTimeSeriesWillR200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesWillR200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWillR200Response) GetMetaOk() (*GetTimeSeriesWillR200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesWillR200Response) SetMeta(v GetTimeSeriesWillR200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesWillR200Response) GetValues() []GetTimeSeriesWillR200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesWillR200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWillR200Response) GetValuesOk() ([]GetTimeSeriesWillR200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesWillR200Response) SetValues(v []GetTimeSeriesWillR200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesWillR200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWillR200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesWillR200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesWillR200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesWillR200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesWillR200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesWillR200Response := _GetTimeSeriesWillR200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesWillR200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesWillR200Response(varGetTimeSeriesWillR200Response)

	return err
}

type NullableGetTimeSeriesWillR200Response struct {
	value *GetTimeSeriesWillR200Response
	isSet bool
}

func (v NullableGetTimeSeriesWillR200Response) Get() *GetTimeSeriesWillR200Response {
	return v.value
}

func (v *NullableGetTimeSeriesWillR200Response) Set(val *GetTimeSeriesWillR200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesWillR200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesWillR200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesWillR200Response(val *GetTimeSeriesWillR200Response) *NullableGetTimeSeriesWillR200Response {
	return &NullableGetTimeSeriesWillR200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesWillR200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesWillR200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
