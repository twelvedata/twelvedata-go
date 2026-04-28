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

// checks if the GetTimeSeriesIchimoku200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesIchimoku200Response{}

// GetTimeSeriesIchimoku200Response struct for GetTimeSeriesIchimoku200Response
type GetTimeSeriesIchimoku200Response struct {
	Meta GetTimeSeriesIchimoku200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesIchimoku200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesIchimoku200Response GetTimeSeriesIchimoku200Response

// NewGetTimeSeriesIchimoku200Response instantiates a new GetTimeSeriesIchimoku200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesIchimoku200Response(meta GetTimeSeriesIchimoku200ResponseMeta, values []GetTimeSeriesIchimoku200ResponseValuesInner, status string) *GetTimeSeriesIchimoku200Response {
	this := GetTimeSeriesIchimoku200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesIchimoku200ResponseWithDefaults instantiates a new GetTimeSeriesIchimoku200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesIchimoku200ResponseWithDefaults() *GetTimeSeriesIchimoku200Response {
	this := GetTimeSeriesIchimoku200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesIchimoku200Response) GetMeta() GetTimeSeriesIchimoku200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesIchimoku200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200Response) GetMetaOk() (*GetTimeSeriesIchimoku200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesIchimoku200Response) SetMeta(v GetTimeSeriesIchimoku200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesIchimoku200Response) GetValues() []GetTimeSeriesIchimoku200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesIchimoku200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200Response) GetValuesOk() ([]GetTimeSeriesIchimoku200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesIchimoku200Response) SetValues(v []GetTimeSeriesIchimoku200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesIchimoku200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesIchimoku200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesIchimoku200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesIchimoku200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesIchimoku200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesIchimoku200Response := _GetTimeSeriesIchimoku200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesIchimoku200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesIchimoku200Response(varGetTimeSeriesIchimoku200Response)

	return err
}

type NullableGetTimeSeriesIchimoku200Response struct {
	value *GetTimeSeriesIchimoku200Response
	isSet bool
}

func (v NullableGetTimeSeriesIchimoku200Response) Get() *GetTimeSeriesIchimoku200Response {
	return v.value
}

func (v *NullableGetTimeSeriesIchimoku200Response) Set(val *GetTimeSeriesIchimoku200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesIchimoku200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesIchimoku200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesIchimoku200Response(val *GetTimeSeriesIchimoku200Response) *NullableGetTimeSeriesIchimoku200Response {
	return &NullableGetTimeSeriesIchimoku200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesIchimoku200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesIchimoku200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
