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

// checks if the GetTimeSeriesRoc200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesRoc200Response{}

// GetTimeSeriesRoc200Response struct for GetTimeSeriesRoc200Response
type GetTimeSeriesRoc200Response struct {
	Meta GetTimeSeriesRoc200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesRoc200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesRoc200Response GetTimeSeriesRoc200Response

// NewGetTimeSeriesRoc200Response instantiates a new GetTimeSeriesRoc200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesRoc200Response(meta GetTimeSeriesRoc200ResponseMeta, values []GetTimeSeriesRoc200ResponseValuesInner, status string) *GetTimeSeriesRoc200Response {
	this := GetTimeSeriesRoc200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesRoc200ResponseWithDefaults instantiates a new GetTimeSeriesRoc200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesRoc200ResponseWithDefaults() *GetTimeSeriesRoc200Response {
	this := GetTimeSeriesRoc200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesRoc200Response) GetMeta() GetTimeSeriesRoc200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesRoc200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRoc200Response) GetMetaOk() (*GetTimeSeriesRoc200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesRoc200Response) SetMeta(v GetTimeSeriesRoc200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesRoc200Response) GetValues() []GetTimeSeriesRoc200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesRoc200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRoc200Response) GetValuesOk() ([]GetTimeSeriesRoc200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesRoc200Response) SetValues(v []GetTimeSeriesRoc200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesRoc200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRoc200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesRoc200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesRoc200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesRoc200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesRoc200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesRoc200Response := _GetTimeSeriesRoc200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesRoc200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesRoc200Response(varGetTimeSeriesRoc200Response)

	return err
}

type NullableGetTimeSeriesRoc200Response struct {
	value *GetTimeSeriesRoc200Response
	isSet bool
}

func (v NullableGetTimeSeriesRoc200Response) Get() *GetTimeSeriesRoc200Response {
	return v.value
}

func (v *NullableGetTimeSeriesRoc200Response) Set(val *GetTimeSeriesRoc200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesRoc200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesRoc200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesRoc200Response(val *GetTimeSeriesRoc200Response) *NullableGetTimeSeriesRoc200Response {
	return &NullableGetTimeSeriesRoc200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesRoc200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesRoc200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
