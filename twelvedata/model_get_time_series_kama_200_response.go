// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesKama200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesKama200Response{}

// GetTimeSeriesKama200Response struct for GetTimeSeriesKama200Response
type GetTimeSeriesKama200Response struct {
	Meta GetTimeSeriesKama200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesKama200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesKama200Response GetTimeSeriesKama200Response

// NewGetTimeSeriesKama200Response instantiates a new GetTimeSeriesKama200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesKama200Response(meta GetTimeSeriesKama200ResponseMeta, values []GetTimeSeriesKama200ResponseValuesInner, status string) *GetTimeSeriesKama200Response {
	this := GetTimeSeriesKama200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesKama200ResponseWithDefaults instantiates a new GetTimeSeriesKama200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesKama200ResponseWithDefaults() *GetTimeSeriesKama200Response {
	this := GetTimeSeriesKama200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesKama200Response) GetMeta() GetTimeSeriesKama200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesKama200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200Response) GetMetaOk() (*GetTimeSeriesKama200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesKama200Response) SetMeta(v GetTimeSeriesKama200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesKama200Response) GetValues() []GetTimeSeriesKama200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesKama200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200Response) GetValuesOk() ([]GetTimeSeriesKama200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesKama200Response) SetValues(v []GetTimeSeriesKama200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesKama200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesKama200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesKama200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesKama200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesKama200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesKama200Response := _GetTimeSeriesKama200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesKama200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesKama200Response(varGetTimeSeriesKama200Response)

	return err
}

type NullableGetTimeSeriesKama200Response struct {
	value *GetTimeSeriesKama200Response
	isSet bool
}

func (v NullableGetTimeSeriesKama200Response) Get() *GetTimeSeriesKama200Response {
	return v.value
}

func (v *NullableGetTimeSeriesKama200Response) Set(val *GetTimeSeriesKama200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesKama200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesKama200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesKama200Response(val *GetTimeSeriesKama200Response) *NullableGetTimeSeriesKama200Response {
	return &NullableGetTimeSeriesKama200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesKama200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesKama200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
