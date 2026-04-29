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

// checks if the GetTimeSeriesT3ma200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesT3ma200Response{}

// GetTimeSeriesT3ma200Response struct for GetTimeSeriesT3ma200Response
type GetTimeSeriesT3ma200Response struct {
	Meta GetTimeSeriesT3ma200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesT3ma200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesT3ma200Response GetTimeSeriesT3ma200Response

// NewGetTimeSeriesT3ma200Response instantiates a new GetTimeSeriesT3ma200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesT3ma200Response(meta GetTimeSeriesT3ma200ResponseMeta, values []GetTimeSeriesT3ma200ResponseValuesInner, status string) *GetTimeSeriesT3ma200Response {
	this := GetTimeSeriesT3ma200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesT3ma200ResponseWithDefaults instantiates a new GetTimeSeriesT3ma200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesT3ma200ResponseWithDefaults() *GetTimeSeriesT3ma200Response {
	this := GetTimeSeriesT3ma200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesT3ma200Response) GetMeta() GetTimeSeriesT3ma200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesT3ma200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesT3ma200Response) GetMetaOk() (*GetTimeSeriesT3ma200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesT3ma200Response) SetMeta(v GetTimeSeriesT3ma200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesT3ma200Response) GetValues() []GetTimeSeriesT3ma200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesT3ma200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesT3ma200Response) GetValuesOk() ([]GetTimeSeriesT3ma200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesT3ma200Response) SetValues(v []GetTimeSeriesT3ma200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesT3ma200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesT3ma200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesT3ma200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesT3ma200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesT3ma200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesT3ma200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesT3ma200Response := _GetTimeSeriesT3ma200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesT3ma200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesT3ma200Response(varGetTimeSeriesT3ma200Response)

	return err
}

type NullableGetTimeSeriesT3ma200Response struct {
	value *GetTimeSeriesT3ma200Response
	isSet bool
}

func (v NullableGetTimeSeriesT3ma200Response) Get() *GetTimeSeriesT3ma200Response {
	return v.value
}

func (v *NullableGetTimeSeriesT3ma200Response) Set(val *GetTimeSeriesT3ma200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesT3ma200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesT3ma200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesT3ma200Response(val *GetTimeSeriesT3ma200Response) *NullableGetTimeSeriesT3ma200Response {
	return &NullableGetTimeSeriesT3ma200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesT3ma200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesT3ma200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
