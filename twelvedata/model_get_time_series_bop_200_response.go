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

// checks if the GetTimeSeriesBop200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesBop200Response{}

// GetTimeSeriesBop200Response struct for GetTimeSeriesBop200Response
type GetTimeSeriesBop200Response struct {
	Meta GetTimeSeriesBop200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesBop200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesBop200Response GetTimeSeriesBop200Response

// NewGetTimeSeriesBop200Response instantiates a new GetTimeSeriesBop200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesBop200Response(meta GetTimeSeriesBop200ResponseMeta, values []GetTimeSeriesBop200ResponseValuesInner, status string) *GetTimeSeriesBop200Response {
	this := GetTimeSeriesBop200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesBop200ResponseWithDefaults instantiates a new GetTimeSeriesBop200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesBop200ResponseWithDefaults() *GetTimeSeriesBop200Response {
	this := GetTimeSeriesBop200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesBop200Response) GetMeta() GetTimeSeriesBop200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesBop200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBop200Response) GetMetaOk() (*GetTimeSeriesBop200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesBop200Response) SetMeta(v GetTimeSeriesBop200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesBop200Response) GetValues() []GetTimeSeriesBop200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesBop200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBop200Response) GetValuesOk() ([]GetTimeSeriesBop200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesBop200Response) SetValues(v []GetTimeSeriesBop200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesBop200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBop200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesBop200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesBop200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesBop200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesBop200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesBop200Response := _GetTimeSeriesBop200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesBop200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesBop200Response(varGetTimeSeriesBop200Response)

	return err
}

type NullableGetTimeSeriesBop200Response struct {
	value *GetTimeSeriesBop200Response
	isSet bool
}

func (v NullableGetTimeSeriesBop200Response) Get() *GetTimeSeriesBop200Response {
	return v.value
}

func (v *NullableGetTimeSeriesBop200Response) Set(val *GetTimeSeriesBop200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesBop200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesBop200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesBop200Response(val *GetTimeSeriesBop200Response) *NullableGetTimeSeriesBop200Response {
	return &NullableGetTimeSeriesBop200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesBop200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesBop200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
