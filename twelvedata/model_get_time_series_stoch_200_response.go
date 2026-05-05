// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesStoch200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesStoch200Response{}

// GetTimeSeriesStoch200Response struct for GetTimeSeriesStoch200Response
type GetTimeSeriesStoch200Response struct {
	Meta GetTimeSeriesStoch200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesStoch200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesStoch200Response GetTimeSeriesStoch200Response

// NewGetTimeSeriesStoch200Response instantiates a new GetTimeSeriesStoch200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesStoch200Response(meta GetTimeSeriesStoch200ResponseMeta, values []GetTimeSeriesStoch200ResponseValuesInner, status string) *GetTimeSeriesStoch200Response {
	this := GetTimeSeriesStoch200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesStoch200ResponseWithDefaults instantiates a new GetTimeSeriesStoch200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesStoch200ResponseWithDefaults() *GetTimeSeriesStoch200Response {
	this := GetTimeSeriesStoch200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesStoch200Response) GetMeta() GetTimeSeriesStoch200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesStoch200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStoch200Response) GetMetaOk() (*GetTimeSeriesStoch200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesStoch200Response) SetMeta(v GetTimeSeriesStoch200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesStoch200Response) GetValues() []GetTimeSeriesStoch200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesStoch200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStoch200Response) GetValuesOk() ([]GetTimeSeriesStoch200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesStoch200Response) SetValues(v []GetTimeSeriesStoch200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesStoch200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStoch200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesStoch200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesStoch200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesStoch200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesStoch200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesStoch200Response := _GetTimeSeriesStoch200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesStoch200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesStoch200Response(varGetTimeSeriesStoch200Response)

	return err
}

type NullableGetTimeSeriesStoch200Response struct {
	value *GetTimeSeriesStoch200Response
	isSet bool
}

func (v NullableGetTimeSeriesStoch200Response) Get() *GetTimeSeriesStoch200Response {
	return v.value
}

func (v *NullableGetTimeSeriesStoch200Response) Set(val *GetTimeSeriesStoch200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesStoch200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesStoch200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesStoch200Response(val *GetTimeSeriesStoch200Response) *NullableGetTimeSeriesStoch200Response {
	return &NullableGetTimeSeriesStoch200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesStoch200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesStoch200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
