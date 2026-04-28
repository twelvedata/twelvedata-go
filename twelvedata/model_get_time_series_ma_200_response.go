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

// checks if the GetTimeSeriesMa200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMa200Response{}

// GetTimeSeriesMa200Response struct for GetTimeSeriesMa200Response
type GetTimeSeriesMa200Response struct {
	Meta GetTimeSeriesMa200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMa200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMa200Response GetTimeSeriesMa200Response

// NewGetTimeSeriesMa200Response instantiates a new GetTimeSeriesMa200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMa200Response(meta GetTimeSeriesMa200ResponseMeta, values []GetTimeSeriesMa200ResponseValuesInner, status string) *GetTimeSeriesMa200Response {
	this := GetTimeSeriesMa200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMa200ResponseWithDefaults instantiates a new GetTimeSeriesMa200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMa200ResponseWithDefaults() *GetTimeSeriesMa200Response {
	this := GetTimeSeriesMa200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMa200Response) GetMeta() GetTimeSeriesMa200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMa200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMa200Response) GetMetaOk() (*GetTimeSeriesMa200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMa200Response) SetMeta(v GetTimeSeriesMa200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMa200Response) GetValues() []GetTimeSeriesMa200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMa200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMa200Response) GetValuesOk() ([]GetTimeSeriesMa200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMa200Response) SetValues(v []GetTimeSeriesMa200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMa200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMa200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMa200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMa200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMa200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMa200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMa200Response := _GetTimeSeriesMa200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMa200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMa200Response(varGetTimeSeriesMa200Response)

	return err
}

type NullableGetTimeSeriesMa200Response struct {
	value *GetTimeSeriesMa200Response
	isSet bool
}

func (v NullableGetTimeSeriesMa200Response) Get() *GetTimeSeriesMa200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMa200Response) Set(val *GetTimeSeriesMa200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMa200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMa200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMa200Response(val *GetTimeSeriesMa200Response) *NullableGetTimeSeriesMa200Response {
	return &NullableGetTimeSeriesMa200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMa200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMa200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
