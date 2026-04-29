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

// checks if the GetTimeSeriesKst200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesKst200Response{}

// GetTimeSeriesKst200Response struct for GetTimeSeriesKst200Response
type GetTimeSeriesKst200Response struct {
	Meta GetTimeSeriesKst200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesKst200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesKst200Response GetTimeSeriesKst200Response

// NewGetTimeSeriesKst200Response instantiates a new GetTimeSeriesKst200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesKst200Response(meta GetTimeSeriesKst200ResponseMeta, values []GetTimeSeriesKst200ResponseValuesInner, status string) *GetTimeSeriesKst200Response {
	this := GetTimeSeriesKst200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesKst200ResponseWithDefaults instantiates a new GetTimeSeriesKst200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesKst200ResponseWithDefaults() *GetTimeSeriesKst200Response {
	this := GetTimeSeriesKst200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesKst200Response) GetMeta() GetTimeSeriesKst200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesKst200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200Response) GetMetaOk() (*GetTimeSeriesKst200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesKst200Response) SetMeta(v GetTimeSeriesKst200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesKst200Response) GetValues() []GetTimeSeriesKst200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesKst200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200Response) GetValuesOk() ([]GetTimeSeriesKst200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesKst200Response) SetValues(v []GetTimeSeriesKst200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesKst200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesKst200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesKst200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesKst200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesKst200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesKst200Response := _GetTimeSeriesKst200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesKst200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesKst200Response(varGetTimeSeriesKst200Response)

	return err
}

type NullableGetTimeSeriesKst200Response struct {
	value *GetTimeSeriesKst200Response
	isSet bool
}

func (v NullableGetTimeSeriesKst200Response) Get() *GetTimeSeriesKst200Response {
	return v.value
}

func (v *NullableGetTimeSeriesKst200Response) Set(val *GetTimeSeriesKst200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesKst200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesKst200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesKst200Response(val *GetTimeSeriesKst200Response) *NullableGetTimeSeriesKst200Response {
	return &NullableGetTimeSeriesKst200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesKst200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesKst200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
