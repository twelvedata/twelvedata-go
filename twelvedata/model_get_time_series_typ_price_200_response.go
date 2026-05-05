// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesTypPrice200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesTypPrice200Response{}

// GetTimeSeriesTypPrice200Response struct for GetTimeSeriesTypPrice200Response
type GetTimeSeriesTypPrice200Response struct {
	Meta GetTimeSeriesTypPrice200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesTypPrice200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesTypPrice200Response GetTimeSeriesTypPrice200Response

// NewGetTimeSeriesTypPrice200Response instantiates a new GetTimeSeriesTypPrice200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesTypPrice200Response(meta GetTimeSeriesTypPrice200ResponseMeta, values []GetTimeSeriesTypPrice200ResponseValuesInner, status string) *GetTimeSeriesTypPrice200Response {
	this := GetTimeSeriesTypPrice200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesTypPrice200ResponseWithDefaults instantiates a new GetTimeSeriesTypPrice200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesTypPrice200ResponseWithDefaults() *GetTimeSeriesTypPrice200Response {
	this := GetTimeSeriesTypPrice200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesTypPrice200Response) GetMeta() GetTimeSeriesTypPrice200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesTypPrice200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTypPrice200Response) GetMetaOk() (*GetTimeSeriesTypPrice200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesTypPrice200Response) SetMeta(v GetTimeSeriesTypPrice200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesTypPrice200Response) GetValues() []GetTimeSeriesTypPrice200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesTypPrice200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTypPrice200Response) GetValuesOk() ([]GetTimeSeriesTypPrice200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesTypPrice200Response) SetValues(v []GetTimeSeriesTypPrice200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesTypPrice200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTypPrice200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesTypPrice200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesTypPrice200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesTypPrice200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesTypPrice200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesTypPrice200Response := _GetTimeSeriesTypPrice200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesTypPrice200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesTypPrice200Response(varGetTimeSeriesTypPrice200Response)

	return err
}

type NullableGetTimeSeriesTypPrice200Response struct {
	value *GetTimeSeriesTypPrice200Response
	isSet bool
}

func (v NullableGetTimeSeriesTypPrice200Response) Get() *GetTimeSeriesTypPrice200Response {
	return v.value
}

func (v *NullableGetTimeSeriesTypPrice200Response) Set(val *GetTimeSeriesTypPrice200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesTypPrice200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesTypPrice200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesTypPrice200Response(val *GetTimeSeriesTypPrice200Response) *NullableGetTimeSeriesTypPrice200Response {
	return &NullableGetTimeSeriesTypPrice200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesTypPrice200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesTypPrice200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
