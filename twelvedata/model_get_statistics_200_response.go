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

// checks if the GetStatistics200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200Response{}

// GetStatistics200Response struct for GetStatistics200Response
type GetStatistics200Response struct {
	Meta       GetStatistics200ResponseMeta       `json:"meta"`
	Statistics GetStatistics200ResponseStatistics `json:"statistics"`
}

type _GetStatistics200Response GetStatistics200Response

// NewGetStatistics200Response instantiates a new GetStatistics200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200Response(meta GetStatistics200ResponseMeta, statistics GetStatistics200ResponseStatistics) *GetStatistics200Response {
	this := GetStatistics200Response{}
	this.Meta = meta
	this.Statistics = statistics
	return &this
}

// NewGetStatistics200ResponseWithDefaults instantiates a new GetStatistics200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseWithDefaults() *GetStatistics200Response {
	this := GetStatistics200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetStatistics200Response) GetMeta() GetStatistics200ResponseMeta {
	if o == nil {
		var ret GetStatistics200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetStatistics200Response) GetMetaOk() (*GetStatistics200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetStatistics200Response) SetMeta(v GetStatistics200ResponseMeta) {
	o.Meta = v
}

// GetStatistics returns the Statistics field value
func (o *GetStatistics200Response) GetStatistics() GetStatistics200ResponseStatistics {
	if o == nil {
		var ret GetStatistics200ResponseStatistics
		return ret
	}

	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value
// and a boolean to check if the value has been set.
func (o *GetStatistics200Response) GetStatisticsOk() (*GetStatistics200ResponseStatistics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statistics, true
}

// SetStatistics sets field value
func (o *GetStatistics200Response) SetStatistics(v GetStatistics200ResponseStatistics) {
	o.Statistics = v
}

func (o GetStatistics200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["statistics"] = o.Statistics
	return toSerialize, nil
}

func (o *GetStatistics200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"statistics",
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

	varGetStatistics200Response := _GetStatistics200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetStatistics200Response)

	if err != nil {
		return err
	}

	*o = GetStatistics200Response(varGetStatistics200Response)

	return err
}

type NullableGetStatistics200Response struct {
	value *GetStatistics200Response
	isSet bool
}

func (v NullableGetStatistics200Response) Get() *GetStatistics200Response {
	return v.value
}

func (v *NullableGetStatistics200Response) Set(val *GetStatistics200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200Response(val *GetStatistics200Response) *NullableGetStatistics200Response {
	return &NullableGetStatistics200Response{value: val, isSet: true}
}

func (v NullableGetStatistics200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
