// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMinIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMinIndex200Response{}

// GetTimeSeriesMinIndex200Response struct for GetTimeSeriesMinIndex200Response
type GetTimeSeriesMinIndex200Response struct {
	Meta GetTimeSeriesMinIndex200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMinIndex200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMinIndex200Response GetTimeSeriesMinIndex200Response

// NewGetTimeSeriesMinIndex200Response instantiates a new GetTimeSeriesMinIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMinIndex200Response(meta GetTimeSeriesMinIndex200ResponseMeta, values []GetTimeSeriesMinIndex200ResponseValuesInner, status string) *GetTimeSeriesMinIndex200Response {
	this := GetTimeSeriesMinIndex200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMinIndex200ResponseWithDefaults instantiates a new GetTimeSeriesMinIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMinIndex200ResponseWithDefaults() *GetTimeSeriesMinIndex200Response {
	this := GetTimeSeriesMinIndex200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMinIndex200Response) GetMeta() GetTimeSeriesMinIndex200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMinIndex200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinIndex200Response) GetMetaOk() (*GetTimeSeriesMinIndex200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMinIndex200Response) SetMeta(v GetTimeSeriesMinIndex200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMinIndex200Response) GetValues() []GetTimeSeriesMinIndex200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMinIndex200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinIndex200Response) GetValuesOk() ([]GetTimeSeriesMinIndex200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMinIndex200Response) SetValues(v []GetTimeSeriesMinIndex200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMinIndex200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinIndex200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMinIndex200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMinIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMinIndex200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMinIndex200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMinIndex200Response := _GetTimeSeriesMinIndex200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMinIndex200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMinIndex200Response(varGetTimeSeriesMinIndex200Response)

	return err
}

type NullableGetTimeSeriesMinIndex200Response struct {
	value *GetTimeSeriesMinIndex200Response
	isSet bool
}

func (v NullableGetTimeSeriesMinIndex200Response) Get() *GetTimeSeriesMinIndex200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMinIndex200Response) Set(val *GetTimeSeriesMinIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMinIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMinIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMinIndex200Response(val *GetTimeSeriesMinIndex200Response) *NullableGetTimeSeriesMinIndex200Response {
	return &NullableGetTimeSeriesMinIndex200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMinIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMinIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
