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

// checks if the GetTimeSeriesPlusDM200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesPlusDM200Response{}

// GetTimeSeriesPlusDM200Response struct for GetTimeSeriesPlusDM200Response
type GetTimeSeriesPlusDM200Response struct {
	Meta GetTimeSeriesPlusDM200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesPlusDM200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesPlusDM200Response GetTimeSeriesPlusDM200Response

// NewGetTimeSeriesPlusDM200Response instantiates a new GetTimeSeriesPlusDM200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesPlusDM200Response(meta GetTimeSeriesPlusDM200ResponseMeta, values []GetTimeSeriesPlusDM200ResponseValuesInner, status string) *GetTimeSeriesPlusDM200Response {
	this := GetTimeSeriesPlusDM200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesPlusDM200ResponseWithDefaults instantiates a new GetTimeSeriesPlusDM200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesPlusDM200ResponseWithDefaults() *GetTimeSeriesPlusDM200Response {
	this := GetTimeSeriesPlusDM200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesPlusDM200Response) GetMeta() GetTimeSeriesPlusDM200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesPlusDM200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPlusDM200Response) GetMetaOk() (*GetTimeSeriesPlusDM200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesPlusDM200Response) SetMeta(v GetTimeSeriesPlusDM200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesPlusDM200Response) GetValues() []GetTimeSeriesPlusDM200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesPlusDM200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPlusDM200Response) GetValuesOk() ([]GetTimeSeriesPlusDM200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesPlusDM200Response) SetValues(v []GetTimeSeriesPlusDM200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesPlusDM200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPlusDM200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesPlusDM200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesPlusDM200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesPlusDM200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesPlusDM200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesPlusDM200Response := _GetTimeSeriesPlusDM200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesPlusDM200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesPlusDM200Response(varGetTimeSeriesPlusDM200Response)

	return err
}

type NullableGetTimeSeriesPlusDM200Response struct {
	value *GetTimeSeriesPlusDM200Response
	isSet bool
}

func (v NullableGetTimeSeriesPlusDM200Response) Get() *GetTimeSeriesPlusDM200Response {
	return v.value
}

func (v *NullableGetTimeSeriesPlusDM200Response) Set(val *GetTimeSeriesPlusDM200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesPlusDM200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesPlusDM200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesPlusDM200Response(val *GetTimeSeriesPlusDM200Response) *NullableGetTimeSeriesPlusDM200Response {
	return &NullableGetTimeSeriesPlusDM200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesPlusDM200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesPlusDM200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
