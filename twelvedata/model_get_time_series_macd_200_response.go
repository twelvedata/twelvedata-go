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

// checks if the GetTimeSeriesMacd200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMacd200Response{}

// GetTimeSeriesMacd200Response struct for GetTimeSeriesMacd200Response
type GetTimeSeriesMacd200Response struct {
	Meta GetTimeSeriesMacd200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMacd200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMacd200Response GetTimeSeriesMacd200Response

// NewGetTimeSeriesMacd200Response instantiates a new GetTimeSeriesMacd200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMacd200Response(meta GetTimeSeriesMacd200ResponseMeta, values []GetTimeSeriesMacd200ResponseValuesInner, status string) *GetTimeSeriesMacd200Response {
	this := GetTimeSeriesMacd200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMacd200ResponseWithDefaults instantiates a new GetTimeSeriesMacd200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMacd200ResponseWithDefaults() *GetTimeSeriesMacd200Response {
	this := GetTimeSeriesMacd200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMacd200Response) GetMeta() GetTimeSeriesMacd200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMacd200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacd200Response) GetMetaOk() (*GetTimeSeriesMacd200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMacd200Response) SetMeta(v GetTimeSeriesMacd200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMacd200Response) GetValues() []GetTimeSeriesMacd200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMacd200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacd200Response) GetValuesOk() ([]GetTimeSeriesMacd200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMacd200Response) SetValues(v []GetTimeSeriesMacd200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMacd200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacd200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMacd200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMacd200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMacd200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMacd200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMacd200Response := _GetTimeSeriesMacd200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMacd200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMacd200Response(varGetTimeSeriesMacd200Response)

	return err
}

type NullableGetTimeSeriesMacd200Response struct {
	value *GetTimeSeriesMacd200Response
	isSet bool
}

func (v NullableGetTimeSeriesMacd200Response) Get() *GetTimeSeriesMacd200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMacd200Response) Set(val *GetTimeSeriesMacd200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMacd200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMacd200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMacd200Response(val *GetTimeSeriesMacd200Response) *NullableGetTimeSeriesMacd200Response {
	return &NullableGetTimeSeriesMacd200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMacd200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMacd200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
