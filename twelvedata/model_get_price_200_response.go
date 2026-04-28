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

// checks if the GetPrice200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPrice200Response{}

// GetPrice200Response struct for GetPrice200Response
type GetPrice200Response struct {
	// Real-time or the latest available price
	Price string `json:"price"`
}

type _GetPrice200Response GetPrice200Response

// NewGetPrice200Response instantiates a new GetPrice200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPrice200Response(price string) *GetPrice200Response {
	this := GetPrice200Response{}
	this.Price = price
	return &this
}

// NewGetPrice200ResponseWithDefaults instantiates a new GetPrice200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPrice200ResponseWithDefaults() *GetPrice200Response {
	this := GetPrice200Response{}
	return &this
}

// GetPrice returns the Price field value
func (o *GetPrice200Response) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *GetPrice200Response) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *GetPrice200Response) SetPrice(v string) {
	o.Price = v
}

func (o GetPrice200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPrice200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["price"] = o.Price
	return toSerialize, nil
}

func (o *GetPrice200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"price",
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

	varGetPrice200Response := _GetPrice200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPrice200Response)

	if err != nil {
		return err
	}

	*o = GetPrice200Response(varGetPrice200Response)

	return err
}

type NullableGetPrice200Response struct {
	value *GetPrice200Response
	isSet bool
}

func (v NullableGetPrice200Response) Get() *GetPrice200Response {
	return v.value
}

func (v *NullableGetPrice200Response) Set(val *GetPrice200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPrice200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPrice200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPrice200Response(val *GetPrice200Response) *NullableGetPrice200Response {
	return &NullableGetPrice200Response{value: val, isSet: true}
}

func (v NullableGetPrice200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPrice200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
