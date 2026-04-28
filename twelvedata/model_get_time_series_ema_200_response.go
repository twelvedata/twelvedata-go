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

// checks if the GetTimeSeriesEma200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesEma200Response{}

// GetTimeSeriesEma200Response struct for GetTimeSeriesEma200Response
type GetTimeSeriesEma200Response struct {
	Meta GetTimeSeriesEma200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesEma200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesEma200Response GetTimeSeriesEma200Response

// NewGetTimeSeriesEma200Response instantiates a new GetTimeSeriesEma200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesEma200Response(meta GetTimeSeriesEma200ResponseMeta, values []GetTimeSeriesEma200ResponseValuesInner, status string) *GetTimeSeriesEma200Response {
	this := GetTimeSeriesEma200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesEma200ResponseWithDefaults instantiates a new GetTimeSeriesEma200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesEma200ResponseWithDefaults() *GetTimeSeriesEma200Response {
	this := GetTimeSeriesEma200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesEma200Response) GetMeta() GetTimeSeriesEma200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesEma200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesEma200Response) GetMetaOk() (*GetTimeSeriesEma200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesEma200Response) SetMeta(v GetTimeSeriesEma200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesEma200Response) GetValues() []GetTimeSeriesEma200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesEma200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesEma200Response) GetValuesOk() ([]GetTimeSeriesEma200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesEma200Response) SetValues(v []GetTimeSeriesEma200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesEma200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesEma200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesEma200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesEma200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesEma200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesEma200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesEma200Response := _GetTimeSeriesEma200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesEma200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesEma200Response(varGetTimeSeriesEma200Response)

	return err
}

type NullableGetTimeSeriesEma200Response struct {
	value *GetTimeSeriesEma200Response
	isSet bool
}

func (v NullableGetTimeSeriesEma200Response) Get() *GetTimeSeriesEma200Response {
	return v.value
}

func (v *NullableGetTimeSeriesEma200Response) Set(val *GetTimeSeriesEma200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesEma200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesEma200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesEma200Response(val *GetTimeSeriesEma200Response) *NullableGetTimeSeriesEma200Response {
	return &NullableGetTimeSeriesEma200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesEma200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesEma200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
