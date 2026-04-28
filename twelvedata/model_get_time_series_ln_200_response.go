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

// checks if the GetTimeSeriesLn200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesLn200Response{}

// GetTimeSeriesLn200Response struct for GetTimeSeriesLn200Response
type GetTimeSeriesLn200Response struct {
	Meta GetTimeSeriesLn200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesLn200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesLn200Response GetTimeSeriesLn200Response

// NewGetTimeSeriesLn200Response instantiates a new GetTimeSeriesLn200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesLn200Response(meta GetTimeSeriesLn200ResponseMeta, values []GetTimeSeriesLn200ResponseValuesInner, status string) *GetTimeSeriesLn200Response {
	this := GetTimeSeriesLn200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesLn200ResponseWithDefaults instantiates a new GetTimeSeriesLn200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesLn200ResponseWithDefaults() *GetTimeSeriesLn200Response {
	this := GetTimeSeriesLn200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesLn200Response) GetMeta() GetTimeSeriesLn200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesLn200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLn200Response) GetMetaOk() (*GetTimeSeriesLn200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesLn200Response) SetMeta(v GetTimeSeriesLn200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesLn200Response) GetValues() []GetTimeSeriesLn200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesLn200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLn200Response) GetValuesOk() ([]GetTimeSeriesLn200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesLn200Response) SetValues(v []GetTimeSeriesLn200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesLn200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLn200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesLn200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesLn200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesLn200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesLn200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesLn200Response := _GetTimeSeriesLn200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesLn200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesLn200Response(varGetTimeSeriesLn200Response)

	return err
}

type NullableGetTimeSeriesLn200Response struct {
	value *GetTimeSeriesLn200Response
	isSet bool
}

func (v NullableGetTimeSeriesLn200Response) Get() *GetTimeSeriesLn200Response {
	return v.value
}

func (v *NullableGetTimeSeriesLn200Response) Set(val *GetTimeSeriesLn200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesLn200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesLn200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesLn200Response(val *GetTimeSeriesLn200Response) *NullableGetTimeSeriesLn200Response {
	return &NullableGetTimeSeriesLn200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesLn200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesLn200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
