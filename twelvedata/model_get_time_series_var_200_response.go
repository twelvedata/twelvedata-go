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

// checks if the GetTimeSeriesVar200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesVar200Response{}

// GetTimeSeriesVar200Response struct for GetTimeSeriesVar200Response
type GetTimeSeriesVar200Response struct {
	Meta GetTimeSeriesVar200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesVar200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesVar200Response GetTimeSeriesVar200Response

// NewGetTimeSeriesVar200Response instantiates a new GetTimeSeriesVar200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesVar200Response(meta GetTimeSeriesVar200ResponseMeta, values []GetTimeSeriesVar200ResponseValuesInner, status string) *GetTimeSeriesVar200Response {
	this := GetTimeSeriesVar200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesVar200ResponseWithDefaults instantiates a new GetTimeSeriesVar200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesVar200ResponseWithDefaults() *GetTimeSeriesVar200Response {
	this := GetTimeSeriesVar200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesVar200Response) GetMeta() GetTimeSeriesVar200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesVar200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVar200Response) GetMetaOk() (*GetTimeSeriesVar200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesVar200Response) SetMeta(v GetTimeSeriesVar200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesVar200Response) GetValues() []GetTimeSeriesVar200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesVar200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVar200Response) GetValuesOk() ([]GetTimeSeriesVar200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesVar200Response) SetValues(v []GetTimeSeriesVar200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesVar200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVar200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesVar200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesVar200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesVar200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesVar200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesVar200Response := _GetTimeSeriesVar200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesVar200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesVar200Response(varGetTimeSeriesVar200Response)

	return err
}

type NullableGetTimeSeriesVar200Response struct {
	value *GetTimeSeriesVar200Response
	isSet bool
}

func (v NullableGetTimeSeriesVar200Response) Get() *GetTimeSeriesVar200Response {
	return v.value
}

func (v *NullableGetTimeSeriesVar200Response) Set(val *GetTimeSeriesVar200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesVar200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesVar200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesVar200Response(val *GetTimeSeriesVar200Response) *NullableGetTimeSeriesVar200Response {
	return &NullableGetTimeSeriesVar200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesVar200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesVar200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
