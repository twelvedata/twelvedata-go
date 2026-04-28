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

// checks if the GetEarliestTimestamp200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEarliestTimestamp200Response{}

// GetEarliestTimestamp200Response struct for GetEarliestTimestamp200Response
type GetEarliestTimestamp200Response struct {
	// Earliest datetime, the format depends on interval
	Datetime string `json:"datetime"`
	// Datetime converted to UNIX timestamp
	UnixTime int64 `json:"unix_time"`
}

type _GetEarliestTimestamp200Response GetEarliestTimestamp200Response

// NewGetEarliestTimestamp200Response instantiates a new GetEarliestTimestamp200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEarliestTimestamp200Response(datetime string, unixTime int64) *GetEarliestTimestamp200Response {
	this := GetEarliestTimestamp200Response{}
	this.Datetime = datetime
	this.UnixTime = unixTime
	return &this
}

// NewGetEarliestTimestamp200ResponseWithDefaults instantiates a new GetEarliestTimestamp200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEarliestTimestamp200ResponseWithDefaults() *GetEarliestTimestamp200Response {
	this := GetEarliestTimestamp200Response{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetEarliestTimestamp200Response) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetEarliestTimestamp200Response) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetEarliestTimestamp200Response) SetDatetime(v string) {
	o.Datetime = v
}

// GetUnixTime returns the UnixTime field value
func (o *GetEarliestTimestamp200Response) GetUnixTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnixTime
}

// GetUnixTimeOk returns a tuple with the UnixTime field value
// and a boolean to check if the value has been set.
func (o *GetEarliestTimestamp200Response) GetUnixTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnixTime, true
}

// SetUnixTime sets field value
func (o *GetEarliestTimestamp200Response) SetUnixTime(v int64) {
	o.UnixTime = v
}

func (o GetEarliestTimestamp200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEarliestTimestamp200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["unix_time"] = o.UnixTime
	return toSerialize, nil
}

func (o *GetEarliestTimestamp200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"unix_time",
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

	varGetEarliestTimestamp200Response := _GetEarliestTimestamp200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEarliestTimestamp200Response)

	if err != nil {
		return err
	}

	*o = GetEarliestTimestamp200Response(varGetEarliestTimestamp200Response)

	return err
}

type NullableGetEarliestTimestamp200Response struct {
	value *GetEarliestTimestamp200Response
	isSet bool
}

func (v NullableGetEarliestTimestamp200Response) Get() *GetEarliestTimestamp200Response {
	return v.value
}

func (v *NullableGetEarliestTimestamp200Response) Set(val *GetEarliestTimestamp200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEarliestTimestamp200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEarliestTimestamp200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEarliestTimestamp200Response(val *GetEarliestTimestamp200Response) *NullableGetEarliestTimestamp200Response {
	return &NullableGetEarliestTimestamp200Response{value: val, isSet: true}
}

func (v NullableGetEarliestTimestamp200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEarliestTimestamp200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
