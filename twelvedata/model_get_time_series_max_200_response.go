// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMax200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMax200Response{}

// GetTimeSeriesMax200Response struct for GetTimeSeriesMax200Response
type GetTimeSeriesMax200Response struct {
	Meta GetTimeSeriesMax200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMax200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMax200Response GetTimeSeriesMax200Response

// NewGetTimeSeriesMax200Response instantiates a new GetTimeSeriesMax200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMax200Response(meta GetTimeSeriesMax200ResponseMeta, values []GetTimeSeriesMax200ResponseValuesInner, status string) *GetTimeSeriesMax200Response {
	this := GetTimeSeriesMax200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMax200ResponseWithDefaults instantiates a new GetTimeSeriesMax200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMax200ResponseWithDefaults() *GetTimeSeriesMax200Response {
	this := GetTimeSeriesMax200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMax200Response) GetMeta() GetTimeSeriesMax200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMax200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMax200Response) GetMetaOk() (*GetTimeSeriesMax200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMax200Response) SetMeta(v GetTimeSeriesMax200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMax200Response) GetValues() []GetTimeSeriesMax200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMax200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMax200Response) GetValuesOk() ([]GetTimeSeriesMax200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMax200Response) SetValues(v []GetTimeSeriesMax200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMax200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMax200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMax200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMax200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMax200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMax200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMax200Response := _GetTimeSeriesMax200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMax200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMax200Response(varGetTimeSeriesMax200Response)

	return err
}

type NullableGetTimeSeriesMax200Response struct {
	value *GetTimeSeriesMax200Response
	isSet bool
}

func (v NullableGetTimeSeriesMax200Response) Get() *GetTimeSeriesMax200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMax200Response) Set(val *GetTimeSeriesMax200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMax200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMax200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMax200Response(val *GetTimeSeriesMax200Response) *NullableGetTimeSeriesMax200Response {
	return &NullableGetTimeSeriesMax200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMax200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMax200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
