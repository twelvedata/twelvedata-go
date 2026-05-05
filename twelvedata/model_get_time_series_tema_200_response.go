// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesTema200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesTema200Response{}

// GetTimeSeriesTema200Response struct for GetTimeSeriesTema200Response
type GetTimeSeriesTema200Response struct {
	Meta GetTimeSeriesTema200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesTema200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesTema200Response GetTimeSeriesTema200Response

// NewGetTimeSeriesTema200Response instantiates a new GetTimeSeriesTema200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesTema200Response(meta GetTimeSeriesTema200ResponseMeta, values []GetTimeSeriesTema200ResponseValuesInner, status string) *GetTimeSeriesTema200Response {
	this := GetTimeSeriesTema200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesTema200ResponseWithDefaults instantiates a new GetTimeSeriesTema200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesTema200ResponseWithDefaults() *GetTimeSeriesTema200Response {
	this := GetTimeSeriesTema200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesTema200Response) GetMeta() GetTimeSeriesTema200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesTema200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTema200Response) GetMetaOk() (*GetTimeSeriesTema200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesTema200Response) SetMeta(v GetTimeSeriesTema200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesTema200Response) GetValues() []GetTimeSeriesTema200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesTema200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTema200Response) GetValuesOk() ([]GetTimeSeriesTema200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesTema200Response) SetValues(v []GetTimeSeriesTema200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesTema200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTema200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesTema200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesTema200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesTema200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesTema200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesTema200Response := _GetTimeSeriesTema200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesTema200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesTema200Response(varGetTimeSeriesTema200Response)

	return err
}

type NullableGetTimeSeriesTema200Response struct {
	value *GetTimeSeriesTema200Response
	isSet bool
}

func (v NullableGetTimeSeriesTema200Response) Get() *GetTimeSeriesTema200Response {
	return v.value
}

func (v *NullableGetTimeSeriesTema200Response) Set(val *GetTimeSeriesTema200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesTema200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesTema200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesTema200Response(val *GetTimeSeriesTema200Response) *NullableGetTimeSeriesTema200Response {
	return &NullableGetTimeSeriesTema200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesTema200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesTema200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
