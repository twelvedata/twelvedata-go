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

// checks if the CryptocurrencyExchangesResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CryptocurrencyExchangesResponseItem{}

// CryptocurrencyExchangesResponseItem struct for CryptocurrencyExchangesResponseItem
type CryptocurrencyExchangesResponseItem struct {
	// Name of cryptocurrency exchange
	Name string `json:"name"`
}

type _CryptocurrencyExchangesResponseItem CryptocurrencyExchangesResponseItem

// NewCryptocurrencyExchangesResponseItem instantiates a new CryptocurrencyExchangesResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptocurrencyExchangesResponseItem(name string) *CryptocurrencyExchangesResponseItem {
	this := CryptocurrencyExchangesResponseItem{}
	this.Name = name
	return &this
}

// NewCryptocurrencyExchangesResponseItemWithDefaults instantiates a new CryptocurrencyExchangesResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptocurrencyExchangesResponseItemWithDefaults() *CryptocurrencyExchangesResponseItem {
	this := CryptocurrencyExchangesResponseItem{}
	return &this
}

// GetName returns the Name field value
func (o *CryptocurrencyExchangesResponseItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CryptocurrencyExchangesResponseItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CryptocurrencyExchangesResponseItem) SetName(v string) {
	o.Name = v
}

func (o CryptocurrencyExchangesResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CryptocurrencyExchangesResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *CryptocurrencyExchangesResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCryptocurrencyExchangesResponseItem := _CryptocurrencyExchangesResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCryptocurrencyExchangesResponseItem)

	if err != nil {
		return err
	}

	*o = CryptocurrencyExchangesResponseItem(varCryptocurrencyExchangesResponseItem)

	return err
}

type NullableCryptocurrencyExchangesResponseItem struct {
	value *CryptocurrencyExchangesResponseItem
	isSet bool
}

func (v NullableCryptocurrencyExchangesResponseItem) Get() *CryptocurrencyExchangesResponseItem {
	return v.value
}

func (v *NullableCryptocurrencyExchangesResponseItem) Set(val *CryptocurrencyExchangesResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptocurrencyExchangesResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptocurrencyExchangesResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptocurrencyExchangesResponseItem(val *CryptocurrencyExchangesResponseItem) *NullableCryptocurrencyExchangesResponseItem {
	return &NullableCryptocurrencyExchangesResponseItem{value: val, isSet: true}
}

func (v NullableCryptocurrencyExchangesResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptocurrencyExchangesResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
