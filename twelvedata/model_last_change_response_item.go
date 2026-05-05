// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the LastChangeResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LastChangeResponseItem{}

// LastChangeResponseItem struct for LastChangeResponseItem
type LastChangeResponseItem struct {
	// Ticker of the company
	Symbol string `json:"symbol"`
	// Market Identifier Code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// The date and time of last changes, in `2006-01-02 15:04:05` format
	LastChange string `json:"last_change"`
}

type _LastChangeResponseItem LastChangeResponseItem

// NewLastChangeResponseItem instantiates a new LastChangeResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastChangeResponseItem(symbol string, micCode string, lastChange string) *LastChangeResponseItem {
	this := LastChangeResponseItem{}
	this.Symbol = symbol
	this.MicCode = micCode
	this.LastChange = lastChange
	return &this
}

// NewLastChangeResponseItemWithDefaults instantiates a new LastChangeResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastChangeResponseItemWithDefaults() *LastChangeResponseItem {
	this := LastChangeResponseItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *LastChangeResponseItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *LastChangeResponseItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *LastChangeResponseItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetMicCode returns the MicCode field value
func (o *LastChangeResponseItem) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *LastChangeResponseItem) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *LastChangeResponseItem) SetMicCode(v string) {
	o.MicCode = v
}

// GetLastChange returns the LastChange field value
func (o *LastChangeResponseItem) GetLastChange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastChange
}

// GetLastChangeOk returns a tuple with the LastChange field value
// and a boolean to check if the value has been set.
func (o *LastChangeResponseItem) GetLastChangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastChange, true
}

// SetLastChange sets field value
func (o *LastChangeResponseItem) SetLastChange(v string) {
	o.LastChange = v
}

func (o LastChangeResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LastChangeResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["mic_code"] = o.MicCode
	toSerialize["last_change"] = o.LastChange
	return toSerialize, nil
}

func (o *LastChangeResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"mic_code",
		"last_change",
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

	varLastChangeResponseItem := _LastChangeResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varLastChangeResponseItem)

	if err != nil {
		return err
	}

	*o = LastChangeResponseItem(varLastChangeResponseItem)

	return err
}

type NullableLastChangeResponseItem struct {
	value *LastChangeResponseItem
	isSet bool
}

func (v NullableLastChangeResponseItem) Get() *LastChangeResponseItem {
	return v.value
}

func (v *NullableLastChangeResponseItem) Set(val *LastChangeResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLastChangeResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLastChangeResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastChangeResponseItem(val *LastChangeResponseItem) *NullableLastChangeResponseItem {
	return &NullableLastChangeResponseItem{value: val, isSet: true}
}

func (v NullableLastChangeResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastChangeResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
