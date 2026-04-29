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

// checks if the CrossListingsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CrossListingsItem{}

// CrossListingsItem Represents details of a cross listing
type CrossListingsItem struct {
	// Ticker symbol of instrument
	Symbol string `json:"symbol"`
	// Name of symbol
	Name string `json:"name"`
	// Exchange where instrument is traded
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
}

type _CrossListingsItem CrossListingsItem

// NewCrossListingsItem instantiates a new CrossListingsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrossListingsItem(symbol string, name string, exchange string, micCode string) *CrossListingsItem {
	this := CrossListingsItem{}
	this.Symbol = symbol
	this.Name = name
	this.Exchange = exchange
	this.MicCode = micCode
	return &this
}

// NewCrossListingsItemWithDefaults instantiates a new CrossListingsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrossListingsItemWithDefaults() *CrossListingsItem {
	this := CrossListingsItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *CrossListingsItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *CrossListingsItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *CrossListingsItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *CrossListingsItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CrossListingsItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CrossListingsItem) SetName(v string) {
	o.Name = v
}

// GetExchange returns the Exchange field value
func (o *CrossListingsItem) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *CrossListingsItem) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *CrossListingsItem) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *CrossListingsItem) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *CrossListingsItem) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *CrossListingsItem) SetMicCode(v string) {
	o.MicCode = v
}

func (o CrossListingsItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CrossListingsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	return toSerialize, nil
}

func (o *CrossListingsItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"exchange",
		"mic_code",
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

	varCrossListingsItem := _CrossListingsItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCrossListingsItem)

	if err != nil {
		return err
	}

	*o = CrossListingsItem(varCrossListingsItem)

	return err
}

type NullableCrossListingsItem struct {
	value *CrossListingsItem
	isSet bool
}

func (v NullableCrossListingsItem) Get() *CrossListingsItem {
	return v.value
}

func (v *NullableCrossListingsItem) Set(val *CrossListingsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCrossListingsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCrossListingsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrossListingsItem(val *CrossListingsItem) *NullableCrossListingsItem {
	return &NullableCrossListingsItem{value: val, isSet: true}
}

func (v NullableCrossListingsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrossListingsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
