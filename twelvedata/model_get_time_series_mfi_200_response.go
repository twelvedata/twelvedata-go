// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMfi200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMfi200Response{}

// GetTimeSeriesMfi200Response struct for GetTimeSeriesMfi200Response
type GetTimeSeriesMfi200Response struct {
	Meta GetTimeSeriesMfi200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMfi200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMfi200Response GetTimeSeriesMfi200Response

// NewGetTimeSeriesMfi200Response instantiates a new GetTimeSeriesMfi200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMfi200Response(meta GetTimeSeriesMfi200ResponseMeta, values []GetTimeSeriesMfi200ResponseValuesInner, status string) *GetTimeSeriesMfi200Response {
	this := GetTimeSeriesMfi200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMfi200ResponseWithDefaults instantiates a new GetTimeSeriesMfi200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMfi200ResponseWithDefaults() *GetTimeSeriesMfi200Response {
	this := GetTimeSeriesMfi200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMfi200Response) GetMeta() GetTimeSeriesMfi200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMfi200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMfi200Response) GetMetaOk() (*GetTimeSeriesMfi200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMfi200Response) SetMeta(v GetTimeSeriesMfi200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMfi200Response) GetValues() []GetTimeSeriesMfi200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMfi200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMfi200Response) GetValuesOk() ([]GetTimeSeriesMfi200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMfi200Response) SetValues(v []GetTimeSeriesMfi200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMfi200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMfi200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMfi200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMfi200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMfi200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMfi200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMfi200Response := _GetTimeSeriesMfi200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMfi200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMfi200Response(varGetTimeSeriesMfi200Response)

	return err
}

type NullableGetTimeSeriesMfi200Response struct {
	value *GetTimeSeriesMfi200Response
	isSet bool
}

func (v NullableGetTimeSeriesMfi200Response) Get() *GetTimeSeriesMfi200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMfi200Response) Set(val *GetTimeSeriesMfi200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMfi200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMfi200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMfi200Response(val *GetTimeSeriesMfi200Response) *NullableGetTimeSeriesMfi200Response {
	return &NullableGetTimeSeriesMfi200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMfi200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMfi200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
