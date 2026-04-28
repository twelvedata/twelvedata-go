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

// checks if the GetTimeSeriesAvg200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAvg200Response{}

// GetTimeSeriesAvg200Response struct for GetTimeSeriesAvg200Response
type GetTimeSeriesAvg200Response struct {
	Meta GetTimeSeriesAvg200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesAvg200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesAvg200Response GetTimeSeriesAvg200Response

// NewGetTimeSeriesAvg200Response instantiates a new GetTimeSeriesAvg200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAvg200Response(meta GetTimeSeriesAvg200ResponseMeta, values []GetTimeSeriesAvg200ResponseValuesInner, status string) *GetTimeSeriesAvg200Response {
	this := GetTimeSeriesAvg200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesAvg200ResponseWithDefaults instantiates a new GetTimeSeriesAvg200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAvg200ResponseWithDefaults() *GetTimeSeriesAvg200Response {
	this := GetTimeSeriesAvg200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesAvg200Response) GetMeta() GetTimeSeriesAvg200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesAvg200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAvg200Response) GetMetaOk() (*GetTimeSeriesAvg200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesAvg200Response) SetMeta(v GetTimeSeriesAvg200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesAvg200Response) GetValues() []GetTimeSeriesAvg200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesAvg200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAvg200Response) GetValuesOk() ([]GetTimeSeriesAvg200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesAvg200Response) SetValues(v []GetTimeSeriesAvg200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesAvg200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAvg200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesAvg200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesAvg200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAvg200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesAvg200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAvg200Response := _GetTimeSeriesAvg200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesAvg200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAvg200Response(varGetTimeSeriesAvg200Response)

	return err
}

type NullableGetTimeSeriesAvg200Response struct {
	value *GetTimeSeriesAvg200Response
	isSet bool
}

func (v NullableGetTimeSeriesAvg200Response) Get() *GetTimeSeriesAvg200Response {
	return v.value
}

func (v *NullableGetTimeSeriesAvg200Response) Set(val *GetTimeSeriesAvg200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAvg200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAvg200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAvg200Response(val *GetTimeSeriesAvg200Response) *NullableGetTimeSeriesAvg200Response {
	return &NullableGetTimeSeriesAvg200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAvg200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAvg200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
