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

// checks if the MarketMoversResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketMoversResponseBody{}

// MarketMoversResponseBody struct for MarketMoversResponseBody
type MarketMoversResponseBody struct {
	// Market movers list
	Values []MarketMoversResponseValue `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _MarketMoversResponseBody MarketMoversResponseBody

// NewMarketMoversResponseBody instantiates a new MarketMoversResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketMoversResponseBody(values []MarketMoversResponseValue, status string) *MarketMoversResponseBody {
	this := MarketMoversResponseBody{}
	this.Values = values
	this.Status = status
	return &this
}

// NewMarketMoversResponseBodyWithDefaults instantiates a new MarketMoversResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketMoversResponseBodyWithDefaults() *MarketMoversResponseBody {
	this := MarketMoversResponseBody{}
	return &this
}

// GetValues returns the Values field value
func (o *MarketMoversResponseBody) GetValues() []MarketMoversResponseValue {
	if o == nil {
		var ret []MarketMoversResponseValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseBody) GetValuesOk() ([]MarketMoversResponseValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *MarketMoversResponseBody) SetValues(v []MarketMoversResponseValue) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *MarketMoversResponseBody) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseBody) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MarketMoversResponseBody) SetStatus(v string) {
	o.Status = v
}

func (o MarketMoversResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketMoversResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *MarketMoversResponseBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varMarketMoversResponseBody := _MarketMoversResponseBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varMarketMoversResponseBody)

	if err != nil {
		return err
	}

	*o = MarketMoversResponseBody(varMarketMoversResponseBody)

	return err
}

type NullableMarketMoversResponseBody struct {
	value *MarketMoversResponseBody
	isSet bool
}

func (v NullableMarketMoversResponseBody) Get() *MarketMoversResponseBody {
	return v.value
}

func (v *NullableMarketMoversResponseBody) Set(val *MarketMoversResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketMoversResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketMoversResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketMoversResponseBody(val *MarketMoversResponseBody) *NullableMarketMoversResponseBody {
	return &NullableMarketMoversResponseBody{value: val, isSet: true}
}

func (v NullableMarketMoversResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketMoversResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
