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

// checks if the GetTimeSeriesSum200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSum200Response{}

// GetTimeSeriesSum200Response struct for GetTimeSeriesSum200Response
type GetTimeSeriesSum200Response struct {
	Meta GetTimeSeriesSum200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesSum200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesSum200Response GetTimeSeriesSum200Response

// NewGetTimeSeriesSum200Response instantiates a new GetTimeSeriesSum200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSum200Response(meta GetTimeSeriesSum200ResponseMeta, values []GetTimeSeriesSum200ResponseValuesInner, status string) *GetTimeSeriesSum200Response {
	this := GetTimeSeriesSum200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesSum200ResponseWithDefaults instantiates a new GetTimeSeriesSum200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSum200ResponseWithDefaults() *GetTimeSeriesSum200Response {
	this := GetTimeSeriesSum200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesSum200Response) GetMeta() GetTimeSeriesSum200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesSum200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSum200Response) GetMetaOk() (*GetTimeSeriesSum200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesSum200Response) SetMeta(v GetTimeSeriesSum200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesSum200Response) GetValues() []GetTimeSeriesSum200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesSum200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSum200Response) GetValuesOk() ([]GetTimeSeriesSum200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesSum200Response) SetValues(v []GetTimeSeriesSum200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesSum200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSum200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesSum200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesSum200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSum200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesSum200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesSum200Response := _GetTimeSeriesSum200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesSum200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSum200Response(varGetTimeSeriesSum200Response)

	return err
}

type NullableGetTimeSeriesSum200Response struct {
	value *GetTimeSeriesSum200Response
	isSet bool
}

func (v NullableGetTimeSeriesSum200Response) Get() *GetTimeSeriesSum200Response {
	return v.value
}

func (v *NullableGetTimeSeriesSum200Response) Set(val *GetTimeSeriesSum200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSum200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSum200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSum200Response(val *GetTimeSeriesSum200Response) *NullableGetTimeSeriesSum200Response {
	return &NullableGetTimeSeriesSum200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSum200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSum200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
