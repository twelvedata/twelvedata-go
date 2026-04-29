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

// checks if the GetTimeSeriesMinusDI200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMinusDI200Response{}

// GetTimeSeriesMinusDI200Response struct for GetTimeSeriesMinusDI200Response
type GetTimeSeriesMinusDI200Response struct {
	Meta GetTimeSeriesMinusDI200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMinusDI200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMinusDI200Response GetTimeSeriesMinusDI200Response

// NewGetTimeSeriesMinusDI200Response instantiates a new GetTimeSeriesMinusDI200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMinusDI200Response(meta GetTimeSeriesMinusDI200ResponseMeta, values []GetTimeSeriesMinusDI200ResponseValuesInner, status string) *GetTimeSeriesMinusDI200Response {
	this := GetTimeSeriesMinusDI200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMinusDI200ResponseWithDefaults instantiates a new GetTimeSeriesMinusDI200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMinusDI200ResponseWithDefaults() *GetTimeSeriesMinusDI200Response {
	this := GetTimeSeriesMinusDI200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMinusDI200Response) GetMeta() GetTimeSeriesMinusDI200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMinusDI200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDI200Response) GetMetaOk() (*GetTimeSeriesMinusDI200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMinusDI200Response) SetMeta(v GetTimeSeriesMinusDI200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMinusDI200Response) GetValues() []GetTimeSeriesMinusDI200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMinusDI200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDI200Response) GetValuesOk() ([]GetTimeSeriesMinusDI200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMinusDI200Response) SetValues(v []GetTimeSeriesMinusDI200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMinusDI200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDI200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMinusDI200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMinusDI200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMinusDI200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMinusDI200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMinusDI200Response := _GetTimeSeriesMinusDI200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMinusDI200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMinusDI200Response(varGetTimeSeriesMinusDI200Response)

	return err
}

type NullableGetTimeSeriesMinusDI200Response struct {
	value *GetTimeSeriesMinusDI200Response
	isSet bool
}

func (v NullableGetTimeSeriesMinusDI200Response) Get() *GetTimeSeriesMinusDI200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMinusDI200Response) Set(val *GetTimeSeriesMinusDI200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMinusDI200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMinusDI200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMinusDI200Response(val *GetTimeSeriesMinusDI200Response) *NullableGetTimeSeriesMinusDI200Response {
	return &NullableGetTimeSeriesMinusDI200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMinusDI200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMinusDI200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
