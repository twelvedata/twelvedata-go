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

// checks if the GetTimeSeriesMinMax200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMinMax200Response{}

// GetTimeSeriesMinMax200Response struct for GetTimeSeriesMinMax200Response
type GetTimeSeriesMinMax200Response struct {
	Meta GetTimeSeriesMinMax200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMinMax200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMinMax200Response GetTimeSeriesMinMax200Response

// NewGetTimeSeriesMinMax200Response instantiates a new GetTimeSeriesMinMax200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMinMax200Response(meta GetTimeSeriesMinMax200ResponseMeta, values []GetTimeSeriesMinMax200ResponseValuesInner, status string) *GetTimeSeriesMinMax200Response {
	this := GetTimeSeriesMinMax200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMinMax200ResponseWithDefaults instantiates a new GetTimeSeriesMinMax200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMinMax200ResponseWithDefaults() *GetTimeSeriesMinMax200Response {
	this := GetTimeSeriesMinMax200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMinMax200Response) GetMeta() GetTimeSeriesMinMax200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMinMax200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinMax200Response) GetMetaOk() (*GetTimeSeriesMinMax200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMinMax200Response) SetMeta(v GetTimeSeriesMinMax200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMinMax200Response) GetValues() []GetTimeSeriesMinMax200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMinMax200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinMax200Response) GetValuesOk() ([]GetTimeSeriesMinMax200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMinMax200Response) SetValues(v []GetTimeSeriesMinMax200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMinMax200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinMax200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMinMax200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMinMax200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMinMax200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMinMax200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMinMax200Response := _GetTimeSeriesMinMax200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMinMax200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMinMax200Response(varGetTimeSeriesMinMax200Response)

	return err
}

type NullableGetTimeSeriesMinMax200Response struct {
	value *GetTimeSeriesMinMax200Response
	isSet bool
}

func (v NullableGetTimeSeriesMinMax200Response) Get() *GetTimeSeriesMinMax200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMinMax200Response) Set(val *GetTimeSeriesMinMax200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMinMax200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMinMax200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMinMax200Response(val *GetTimeSeriesMinMax200Response) *NullableGetTimeSeriesMinMax200Response {
	return &NullableGetTimeSeriesMinMax200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMinMax200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMinMax200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
