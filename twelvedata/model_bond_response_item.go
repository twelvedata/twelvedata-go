// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BondResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BondResponseItem{}

// BondResponseItem BondResponseItem represents details of a bond
type BondResponseItem struct {
	// Bond symbol
	Symbol string `json:"symbol"`
	// Full name of the bond
	Name string `json:"name"`
	// Country where the bond is located
	Country string `json:"country"`
	// Currency of the bond according to the ISO 4217 standard
	Currency string `json:"currency"`
	// Exchange where the bond is traded
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Type of the bond
	Type   string                   `json:"type"`
	Access *BondsResponseItemAccess `json:"access,omitempty"`
}

type _BondResponseItem BondResponseItem

// NewBondResponseItem instantiates a new BondResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBondResponseItem(symbol string, name string, country string, currency string, exchange string, micCode string, type_ string) *BondResponseItem {
	this := BondResponseItem{}
	this.Symbol = symbol
	this.Name = name
	this.Country = country
	this.Currency = currency
	this.Exchange = exchange
	this.MicCode = micCode
	this.Type = type_
	return &this
}

// NewBondResponseItemWithDefaults instantiates a new BondResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBondResponseItemWithDefaults() *BondResponseItem {
	this := BondResponseItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *BondResponseItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *BondResponseItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *BondResponseItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *BondResponseItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BondResponseItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BondResponseItem) SetName(v string) {
	o.Name = v
}

// GetCountry returns the Country field value
func (o *BondResponseItem) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *BondResponseItem) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *BondResponseItem) SetCountry(v string) {
	o.Country = v
}

// GetCurrency returns the Currency field value
func (o *BondResponseItem) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *BondResponseItem) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *BondResponseItem) SetCurrency(v string) {
	o.Currency = v
}

// GetExchange returns the Exchange field value
func (o *BondResponseItem) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *BondResponseItem) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *BondResponseItem) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *BondResponseItem) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *BondResponseItem) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *BondResponseItem) SetMicCode(v string) {
	o.MicCode = v
}

// GetType returns the Type field value
func (o *BondResponseItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BondResponseItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BondResponseItem) SetType(v string) {
	o.Type = v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *BondResponseItem) GetAccess() BondsResponseItemAccess {
	if o == nil || IsNil(o.Access) {
		var ret BondsResponseItemAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BondResponseItem) GetAccessOk() (*BondsResponseItemAccess, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *BondResponseItem) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given BondsResponseItemAccess and assigns it to the Access field.
func (o *BondResponseItem) SetAccess(v BondsResponseItemAccess) {
	o.Access = &v
}

func (o BondResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BondResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["country"] = o.Country
	toSerialize["currency"] = o.Currency
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["type"] = o.Type
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}

func (o *BondResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"country",
		"currency",
		"exchange",
		"mic_code",
		"type",
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

	varBondResponseItem := _BondResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varBondResponseItem)

	if err != nil {
		return err
	}

	*o = BondResponseItem(varBondResponseItem)

	return err
}

type NullableBondResponseItem struct {
	value *BondResponseItem
	isSet bool
}

func (v NullableBondResponseItem) Get() *BondResponseItem {
	return v.value
}

func (v *NullableBondResponseItem) Set(val *BondResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBondResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBondResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBondResponseItem(val *BondResponseItem) *NullableBondResponseItem {
	return &NullableBondResponseItem{value: val, isSet: true}
}

func (v NullableBondResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBondResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
