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

// checks if the GetTimeSeriesMidPrice200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMidPrice200Response{}

// GetTimeSeriesMidPrice200Response struct for GetTimeSeriesMidPrice200Response
type GetTimeSeriesMidPrice200Response struct {
	Meta GetTimeSeriesMidPrice200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMidPrice200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMidPrice200Response GetTimeSeriesMidPrice200Response

// NewGetTimeSeriesMidPrice200Response instantiates a new GetTimeSeriesMidPrice200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMidPrice200Response(meta GetTimeSeriesMidPrice200ResponseMeta, values []GetTimeSeriesMidPrice200ResponseValuesInner, status string) *GetTimeSeriesMidPrice200Response {
	this := GetTimeSeriesMidPrice200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMidPrice200ResponseWithDefaults instantiates a new GetTimeSeriesMidPrice200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMidPrice200ResponseWithDefaults() *GetTimeSeriesMidPrice200Response {
	this := GetTimeSeriesMidPrice200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMidPrice200Response) GetMeta() GetTimeSeriesMidPrice200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMidPrice200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMidPrice200Response) GetMetaOk() (*GetTimeSeriesMidPrice200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMidPrice200Response) SetMeta(v GetTimeSeriesMidPrice200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMidPrice200Response) GetValues() []GetTimeSeriesMidPrice200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMidPrice200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMidPrice200Response) GetValuesOk() ([]GetTimeSeriesMidPrice200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMidPrice200Response) SetValues(v []GetTimeSeriesMidPrice200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMidPrice200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMidPrice200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMidPrice200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMidPrice200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMidPrice200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMidPrice200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMidPrice200Response := _GetTimeSeriesMidPrice200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMidPrice200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMidPrice200Response(varGetTimeSeriesMidPrice200Response)

	return err
}

type NullableGetTimeSeriesMidPrice200Response struct {
	value *GetTimeSeriesMidPrice200Response
	isSet bool
}

func (v NullableGetTimeSeriesMidPrice200Response) Get() *GetTimeSeriesMidPrice200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMidPrice200Response) Set(val *GetTimeSeriesMidPrice200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMidPrice200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMidPrice200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMidPrice200Response(val *GetTimeSeriesMidPrice200Response) *NullableGetTimeSeriesMidPrice200Response {
	return &NullableGetTimeSeriesMidPrice200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMidPrice200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMidPrice200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
