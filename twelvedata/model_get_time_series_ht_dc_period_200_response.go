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

// checks if the GetTimeSeriesHtDcPeriod200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHtDcPeriod200Response{}

// GetTimeSeriesHtDcPeriod200Response struct for GetTimeSeriesHtDcPeriod200Response
type GetTimeSeriesHtDcPeriod200Response struct {
	Meta GetTimeSeriesHtDcPeriod200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesHtDcPeriod200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesHtDcPeriod200Response GetTimeSeriesHtDcPeriod200Response

// NewGetTimeSeriesHtDcPeriod200Response instantiates a new GetTimeSeriesHtDcPeriod200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHtDcPeriod200Response(meta GetTimeSeriesHtDcPeriod200ResponseMeta, values []GetTimeSeriesHtDcPeriod200ResponseValuesInner, status string) *GetTimeSeriesHtDcPeriod200Response {
	this := GetTimeSeriesHtDcPeriod200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesHtDcPeriod200ResponseWithDefaults instantiates a new GetTimeSeriesHtDcPeriod200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHtDcPeriod200ResponseWithDefaults() *GetTimeSeriesHtDcPeriod200Response {
	this := GetTimeSeriesHtDcPeriod200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesHtDcPeriod200Response) GetMeta() GetTimeSeriesHtDcPeriod200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesHtDcPeriod200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtDcPeriod200Response) GetMetaOk() (*GetTimeSeriesHtDcPeriod200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesHtDcPeriod200Response) SetMeta(v GetTimeSeriesHtDcPeriod200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesHtDcPeriod200Response) GetValues() []GetTimeSeriesHtDcPeriod200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesHtDcPeriod200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtDcPeriod200Response) GetValuesOk() ([]GetTimeSeriesHtDcPeriod200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesHtDcPeriod200Response) SetValues(v []GetTimeSeriesHtDcPeriod200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesHtDcPeriod200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtDcPeriod200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesHtDcPeriod200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesHtDcPeriod200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHtDcPeriod200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesHtDcPeriod200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesHtDcPeriod200Response := _GetTimeSeriesHtDcPeriod200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesHtDcPeriod200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHtDcPeriod200Response(varGetTimeSeriesHtDcPeriod200Response)

	return err
}

type NullableGetTimeSeriesHtDcPeriod200Response struct {
	value *GetTimeSeriesHtDcPeriod200Response
	isSet bool
}

func (v NullableGetTimeSeriesHtDcPeriod200Response) Get() *GetTimeSeriesHtDcPeriod200Response {
	return v.value
}

func (v *NullableGetTimeSeriesHtDcPeriod200Response) Set(val *GetTimeSeriesHtDcPeriod200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHtDcPeriod200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHtDcPeriod200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHtDcPeriod200Response(val *GetTimeSeriesHtDcPeriod200Response) *NullableGetTimeSeriesHtDcPeriod200Response {
	return &NullableGetTimeSeriesHtDcPeriod200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHtDcPeriod200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHtDcPeriod200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
