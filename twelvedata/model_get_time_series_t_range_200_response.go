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

// checks if the GetTimeSeriesTRange200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesTRange200Response{}

// GetTimeSeriesTRange200Response struct for GetTimeSeriesTRange200Response
type GetTimeSeriesTRange200Response struct {
	Meta GetTimeSeriesTRange200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesTRange200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesTRange200Response GetTimeSeriesTRange200Response

// NewGetTimeSeriesTRange200Response instantiates a new GetTimeSeriesTRange200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesTRange200Response(meta GetTimeSeriesTRange200ResponseMeta, values []GetTimeSeriesTRange200ResponseValuesInner, status string) *GetTimeSeriesTRange200Response {
	this := GetTimeSeriesTRange200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesTRange200ResponseWithDefaults instantiates a new GetTimeSeriesTRange200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesTRange200ResponseWithDefaults() *GetTimeSeriesTRange200Response {
	this := GetTimeSeriesTRange200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesTRange200Response) GetMeta() GetTimeSeriesTRange200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesTRange200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTRange200Response) GetMetaOk() (*GetTimeSeriesTRange200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesTRange200Response) SetMeta(v GetTimeSeriesTRange200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesTRange200Response) GetValues() []GetTimeSeriesTRange200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesTRange200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTRange200Response) GetValuesOk() ([]GetTimeSeriesTRange200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesTRange200Response) SetValues(v []GetTimeSeriesTRange200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesTRange200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTRange200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesTRange200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesTRange200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesTRange200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesTRange200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesTRange200Response := _GetTimeSeriesTRange200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesTRange200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesTRange200Response(varGetTimeSeriesTRange200Response)

	return err
}

type NullableGetTimeSeriesTRange200Response struct {
	value *GetTimeSeriesTRange200Response
	isSet bool
}

func (v NullableGetTimeSeriesTRange200Response) Get() *GetTimeSeriesTRange200Response {
	return v.value
}

func (v *NullableGetTimeSeriesTRange200Response) Set(val *GetTimeSeriesTRange200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesTRange200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesTRange200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesTRange200Response(val *GetTimeSeriesTRange200Response) *NullableGetTimeSeriesTRange200Response {
	return &NullableGetTimeSeriesTRange200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesTRange200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesTRange200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
