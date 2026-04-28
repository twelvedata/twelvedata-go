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

// checks if the GetTimeSeriesMidPoint200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMidPoint200Response{}

// GetTimeSeriesMidPoint200Response struct for GetTimeSeriesMidPoint200Response
type GetTimeSeriesMidPoint200Response struct {
	Meta GetTimeSeriesMidPoint200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMidPoint200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMidPoint200Response GetTimeSeriesMidPoint200Response

// NewGetTimeSeriesMidPoint200Response instantiates a new GetTimeSeriesMidPoint200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMidPoint200Response(meta GetTimeSeriesMidPoint200ResponseMeta, values []GetTimeSeriesMidPoint200ResponseValuesInner, status string) *GetTimeSeriesMidPoint200Response {
	this := GetTimeSeriesMidPoint200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMidPoint200ResponseWithDefaults instantiates a new GetTimeSeriesMidPoint200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMidPoint200ResponseWithDefaults() *GetTimeSeriesMidPoint200Response {
	this := GetTimeSeriesMidPoint200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMidPoint200Response) GetMeta() GetTimeSeriesMidPoint200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMidPoint200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMidPoint200Response) GetMetaOk() (*GetTimeSeriesMidPoint200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMidPoint200Response) SetMeta(v GetTimeSeriesMidPoint200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMidPoint200Response) GetValues() []GetTimeSeriesMidPoint200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMidPoint200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMidPoint200Response) GetValuesOk() ([]GetTimeSeriesMidPoint200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMidPoint200Response) SetValues(v []GetTimeSeriesMidPoint200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMidPoint200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMidPoint200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMidPoint200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMidPoint200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMidPoint200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMidPoint200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMidPoint200Response := _GetTimeSeriesMidPoint200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMidPoint200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMidPoint200Response(varGetTimeSeriesMidPoint200Response)

	return err
}

type NullableGetTimeSeriesMidPoint200Response struct {
	value *GetTimeSeriesMidPoint200Response
	isSet bool
}

func (v NullableGetTimeSeriesMidPoint200Response) Get() *GetTimeSeriesMidPoint200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMidPoint200Response) Set(val *GetTimeSeriesMidPoint200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMidPoint200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMidPoint200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMidPoint200Response(val *GetTimeSeriesMidPoint200Response) *NullableGetTimeSeriesMidPoint200Response {
	return &NullableGetTimeSeriesMidPoint200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMidPoint200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMidPoint200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
