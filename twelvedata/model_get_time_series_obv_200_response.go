// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesObv200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesObv200Response{}

// GetTimeSeriesObv200Response struct for GetTimeSeriesObv200Response
type GetTimeSeriesObv200Response struct {
	Meta GetTimeSeriesObv200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesObv200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesObv200Response GetTimeSeriesObv200Response

// NewGetTimeSeriesObv200Response instantiates a new GetTimeSeriesObv200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesObv200Response(meta GetTimeSeriesObv200ResponseMeta, values []GetTimeSeriesObv200ResponseValuesInner, status string) *GetTimeSeriesObv200Response {
	this := GetTimeSeriesObv200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesObv200ResponseWithDefaults instantiates a new GetTimeSeriesObv200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesObv200ResponseWithDefaults() *GetTimeSeriesObv200Response {
	this := GetTimeSeriesObv200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesObv200Response) GetMeta() GetTimeSeriesObv200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesObv200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesObv200Response) GetMetaOk() (*GetTimeSeriesObv200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesObv200Response) SetMeta(v GetTimeSeriesObv200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesObv200Response) GetValues() []GetTimeSeriesObv200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesObv200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesObv200Response) GetValuesOk() ([]GetTimeSeriesObv200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesObv200Response) SetValues(v []GetTimeSeriesObv200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesObv200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesObv200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesObv200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesObv200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesObv200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesObv200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesObv200Response := _GetTimeSeriesObv200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesObv200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesObv200Response(varGetTimeSeriesObv200Response)

	return err
}

type NullableGetTimeSeriesObv200Response struct {
	value *GetTimeSeriesObv200Response
	isSet bool
}

func (v NullableGetTimeSeriesObv200Response) Get() *GetTimeSeriesObv200Response {
	return v.value
}

func (v *NullableGetTimeSeriesObv200Response) Set(val *GetTimeSeriesObv200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesObv200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesObv200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesObv200Response(val *GetTimeSeriesObv200Response) *NullableGetTimeSeriesObv200Response {
	return &NullableGetTimeSeriesObv200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesObv200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesObv200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
