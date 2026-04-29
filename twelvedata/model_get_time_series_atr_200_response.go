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

// checks if the GetTimeSeriesAtr200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAtr200Response{}

// GetTimeSeriesAtr200Response struct for GetTimeSeriesAtr200Response
type GetTimeSeriesAtr200Response struct {
	Meta GetTimeSeriesAtr200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesAtr200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesAtr200Response GetTimeSeriesAtr200Response

// NewGetTimeSeriesAtr200Response instantiates a new GetTimeSeriesAtr200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAtr200Response(meta GetTimeSeriesAtr200ResponseMeta, values []GetTimeSeriesAtr200ResponseValuesInner, status string) *GetTimeSeriesAtr200Response {
	this := GetTimeSeriesAtr200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesAtr200ResponseWithDefaults instantiates a new GetTimeSeriesAtr200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAtr200ResponseWithDefaults() *GetTimeSeriesAtr200Response {
	this := GetTimeSeriesAtr200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesAtr200Response) GetMeta() GetTimeSeriesAtr200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesAtr200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAtr200Response) GetMetaOk() (*GetTimeSeriesAtr200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesAtr200Response) SetMeta(v GetTimeSeriesAtr200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesAtr200Response) GetValues() []GetTimeSeriesAtr200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesAtr200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAtr200Response) GetValuesOk() ([]GetTimeSeriesAtr200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesAtr200Response) SetValues(v []GetTimeSeriesAtr200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesAtr200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAtr200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesAtr200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesAtr200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAtr200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesAtr200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAtr200Response := _GetTimeSeriesAtr200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAtr200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAtr200Response(varGetTimeSeriesAtr200Response)

	return err
}

type NullableGetTimeSeriesAtr200Response struct {
	value *GetTimeSeriesAtr200Response
	isSet bool
}

func (v NullableGetTimeSeriesAtr200Response) Get() *GetTimeSeriesAtr200Response {
	return v.value
}

func (v *NullableGetTimeSeriesAtr200Response) Set(val *GetTimeSeriesAtr200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAtr200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAtr200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAtr200Response(val *GetTimeSeriesAtr200Response) *NullableGetTimeSeriesAtr200Response {
	return &NullableGetTimeSeriesAtr200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAtr200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAtr200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
