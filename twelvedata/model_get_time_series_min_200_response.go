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

// checks if the GetTimeSeriesMin200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMin200Response{}

// GetTimeSeriesMin200Response struct for GetTimeSeriesMin200Response
type GetTimeSeriesMin200Response struct {
	Meta GetTimeSeriesMin200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMin200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMin200Response GetTimeSeriesMin200Response

// NewGetTimeSeriesMin200Response instantiates a new GetTimeSeriesMin200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMin200Response(meta GetTimeSeriesMin200ResponseMeta, values []GetTimeSeriesMin200ResponseValuesInner, status string) *GetTimeSeriesMin200Response {
	this := GetTimeSeriesMin200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMin200ResponseWithDefaults instantiates a new GetTimeSeriesMin200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMin200ResponseWithDefaults() *GetTimeSeriesMin200Response {
	this := GetTimeSeriesMin200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMin200Response) GetMeta() GetTimeSeriesMin200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMin200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMin200Response) GetMetaOk() (*GetTimeSeriesMin200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMin200Response) SetMeta(v GetTimeSeriesMin200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMin200Response) GetValues() []GetTimeSeriesMin200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMin200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMin200Response) GetValuesOk() ([]GetTimeSeriesMin200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMin200Response) SetValues(v []GetTimeSeriesMin200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMin200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMin200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMin200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMin200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMin200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMin200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMin200Response := _GetTimeSeriesMin200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMin200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMin200Response(varGetTimeSeriesMin200Response)

	return err
}

type NullableGetTimeSeriesMin200Response struct {
	value *GetTimeSeriesMin200Response
	isSet bool
}

func (v NullableGetTimeSeriesMin200Response) Get() *GetTimeSeriesMin200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMin200Response) Set(val *GetTimeSeriesMin200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMin200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMin200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMin200Response(val *GetTimeSeriesMin200Response) *NullableGetTimeSeriesMin200Response {
	return &NullableGetTimeSeriesMin200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMin200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMin200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
