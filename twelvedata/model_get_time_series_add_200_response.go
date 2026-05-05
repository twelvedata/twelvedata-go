// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesAdd200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAdd200Response{}

// GetTimeSeriesAdd200Response struct for GetTimeSeriesAdd200Response
type GetTimeSeriesAdd200Response struct {
	Meta GetTimeSeriesAdd200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesAdd200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesAdd200Response GetTimeSeriesAdd200Response

// NewGetTimeSeriesAdd200Response instantiates a new GetTimeSeriesAdd200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAdd200Response(meta GetTimeSeriesAdd200ResponseMeta, values []GetTimeSeriesAdd200ResponseValuesInner, status string) *GetTimeSeriesAdd200Response {
	this := GetTimeSeriesAdd200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesAdd200ResponseWithDefaults instantiates a new GetTimeSeriesAdd200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAdd200ResponseWithDefaults() *GetTimeSeriesAdd200Response {
	this := GetTimeSeriesAdd200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesAdd200Response) GetMeta() GetTimeSeriesAdd200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesAdd200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdd200Response) GetMetaOk() (*GetTimeSeriesAdd200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesAdd200Response) SetMeta(v GetTimeSeriesAdd200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesAdd200Response) GetValues() []GetTimeSeriesAdd200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesAdd200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdd200Response) GetValuesOk() ([]GetTimeSeriesAdd200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesAdd200Response) SetValues(v []GetTimeSeriesAdd200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesAdd200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdd200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesAdd200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesAdd200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAdd200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesAdd200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAdd200Response := _GetTimeSeriesAdd200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAdd200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAdd200Response(varGetTimeSeriesAdd200Response)

	return err
}

type NullableGetTimeSeriesAdd200Response struct {
	value *GetTimeSeriesAdd200Response
	isSet bool
}

func (v NullableGetTimeSeriesAdd200Response) Get() *GetTimeSeriesAdd200Response {
	return v.value
}

func (v *NullableGetTimeSeriesAdd200Response) Set(val *GetTimeSeriesAdd200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAdd200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAdd200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAdd200Response(val *GetTimeSeriesAdd200Response) *NullableGetTimeSeriesAdd200Response {
	return &NullableGetTimeSeriesAdd200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAdd200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAdd200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
