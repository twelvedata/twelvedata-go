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

// checks if the GetExchangeSchedule200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExchangeSchedule200Response{}

// GetExchangeSchedule200Response struct for GetExchangeSchedule200Response
type GetExchangeSchedule200Response struct {
	Data []ExchangeScheduleResponseItem `json:"data"`
}

type _GetExchangeSchedule200Response GetExchangeSchedule200Response

// NewGetExchangeSchedule200Response instantiates a new GetExchangeSchedule200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeSchedule200Response(data []ExchangeScheduleResponseItem) *GetExchangeSchedule200Response {
	this := GetExchangeSchedule200Response{}
	this.Data = data
	return &this
}

// NewGetExchangeSchedule200ResponseWithDefaults instantiates a new GetExchangeSchedule200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeSchedule200ResponseWithDefaults() *GetExchangeSchedule200Response {
	this := GetExchangeSchedule200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetExchangeSchedule200Response) GetData() []ExchangeScheduleResponseItem {
	if o == nil {
		var ret []ExchangeScheduleResponseItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetExchangeSchedule200Response) GetDataOk() ([]ExchangeScheduleResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetExchangeSchedule200Response) SetData(v []ExchangeScheduleResponseItem) {
	o.Data = v
}

func (o GetExchangeSchedule200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExchangeSchedule200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetExchangeSchedule200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGetExchangeSchedule200Response := _GetExchangeSchedule200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetExchangeSchedule200Response)

	if err != nil {
		return err
	}

	*o = GetExchangeSchedule200Response(varGetExchangeSchedule200Response)

	return err
}

type NullableGetExchangeSchedule200Response struct {
	value *GetExchangeSchedule200Response
	isSet bool
}

func (v NullableGetExchangeSchedule200Response) Get() *GetExchangeSchedule200Response {
	return v.value
}

func (v *NullableGetExchangeSchedule200Response) Set(val *GetExchangeSchedule200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeSchedule200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeSchedule200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeSchedule200Response(val *GetExchangeSchedule200Response) *NullableGetExchangeSchedule200Response {
	return &NullableGetExchangeSchedule200Response{value: val, isSet: true}
}

func (v NullableGetExchangeSchedule200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeSchedule200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
