// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMult200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMult200Response{}

// GetTimeSeriesMult200Response struct for GetTimeSeriesMult200Response
type GetTimeSeriesMult200Response struct {
	Meta GetTimeSeriesMult200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMult200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMult200Response GetTimeSeriesMult200Response

// NewGetTimeSeriesMult200Response instantiates a new GetTimeSeriesMult200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMult200Response(meta GetTimeSeriesMult200ResponseMeta, values []GetTimeSeriesMult200ResponseValuesInner, status string) *GetTimeSeriesMult200Response {
	this := GetTimeSeriesMult200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMult200ResponseWithDefaults instantiates a new GetTimeSeriesMult200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMult200ResponseWithDefaults() *GetTimeSeriesMult200Response {
	this := GetTimeSeriesMult200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMult200Response) GetMeta() GetTimeSeriesMult200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMult200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMult200Response) GetMetaOk() (*GetTimeSeriesMult200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMult200Response) SetMeta(v GetTimeSeriesMult200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMult200Response) GetValues() []GetTimeSeriesMult200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMult200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMult200Response) GetValuesOk() ([]GetTimeSeriesMult200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMult200Response) SetValues(v []GetTimeSeriesMult200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMult200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMult200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMult200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMult200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMult200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMult200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMult200Response := _GetTimeSeriesMult200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMult200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMult200Response(varGetTimeSeriesMult200Response)

	return err
}

type NullableGetTimeSeriesMult200Response struct {
	value *GetTimeSeriesMult200Response
	isSet bool
}

func (v NullableGetTimeSeriesMult200Response) Get() *GetTimeSeriesMult200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMult200Response) Set(val *GetTimeSeriesMult200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMult200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMult200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMult200Response(val *GetTimeSeriesMult200Response) *NullableGetTimeSeriesMult200Response {
	return &NullableGetTimeSeriesMult200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMult200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMult200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
