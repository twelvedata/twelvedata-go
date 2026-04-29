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

// checks if the SymbolSearchResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SymbolSearchResponseItem{}

// SymbolSearchResponseItem SymbolSearchResponseItem represents details of a symbol search result
type SymbolSearchResponseItem struct {
	// Ticker symbol of instrument
	Symbol string `json:"symbol"`
	// Name of exchange
	InstrumentName string `json:"instrument_name"`
	// Exchange where instrument is traded
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Time zone where exchange is located
	ExchangeTimezone string `json:"exchange_timezone"`
	// Type of instrument
	InstrumentType string `json:"instrument_type"`
	// Country to which stock exchange belongs to
	Country string `json:"country"`
	// Currency in which the instrument is traded
	Currency string                          `json:"currency"`
	Access   *SymbolSearchResponseItemAccess `json:"access,omitempty"`
}

type _SymbolSearchResponseItem SymbolSearchResponseItem

// NewSymbolSearchResponseItem instantiates a new SymbolSearchResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolSearchResponseItem(symbol string, instrumentName string, exchange string, micCode string, exchangeTimezone string, instrumentType string, country string, currency string) *SymbolSearchResponseItem {
	this := SymbolSearchResponseItem{}
	this.Symbol = symbol
	this.InstrumentName = instrumentName
	this.Exchange = exchange
	this.MicCode = micCode
	this.ExchangeTimezone = exchangeTimezone
	this.InstrumentType = instrumentType
	this.Country = country
	this.Currency = currency
	return &this
}

// NewSymbolSearchResponseItemWithDefaults instantiates a new SymbolSearchResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolSearchResponseItemWithDefaults() *SymbolSearchResponseItem {
	this := SymbolSearchResponseItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *SymbolSearchResponseItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *SymbolSearchResponseItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *SymbolSearchResponseItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetInstrumentName returns the InstrumentName field value
func (o *SymbolSearchResponseItem) GetInstrumentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstrumentName
}

// GetInstrumentNameOk returns a tuple with the InstrumentName field value
// and a boolean to check if the value has been set.
func (o *SymbolSearchResponseItem) GetInstrumentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstrumentName, true
}

// SetInstrumentName sets field value
func (o *SymbolSearchResponseItem) SetInstrumentName(v string) {
	o.InstrumentName = v
}

// GetExchange returns the Exchange field value
func (o *SymbolSearchResponseItem) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *SymbolSearchResponseItem) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *SymbolSearchResponseItem) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *SymbolSearchResponseItem) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *SymbolSearchResponseItem) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *SymbolSearchResponseItem) SetMicCode(v string) {
	o.MicCode = v
}

// GetExchangeTimezone returns the ExchangeTimezone field value
func (o *SymbolSearchResponseItem) GetExchangeTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeTimezone
}

// GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field value
// and a boolean to check if the value has been set.
func (o *SymbolSearchResponseItem) GetExchangeTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeTimezone, true
}

// SetExchangeTimezone sets field value
func (o *SymbolSearchResponseItem) SetExchangeTimezone(v string) {
	o.ExchangeTimezone = v
}

// GetInstrumentType returns the InstrumentType field value
func (o *SymbolSearchResponseItem) GetInstrumentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstrumentType
}

// GetInstrumentTypeOk returns a tuple with the InstrumentType field value
// and a boolean to check if the value has been set.
func (o *SymbolSearchResponseItem) GetInstrumentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstrumentType, true
}

// SetInstrumentType sets field value
func (o *SymbolSearchResponseItem) SetInstrumentType(v string) {
	o.InstrumentType = v
}

// GetCountry returns the Country field value
func (o *SymbolSearchResponseItem) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *SymbolSearchResponseItem) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *SymbolSearchResponseItem) SetCountry(v string) {
	o.Country = v
}

// GetCurrency returns the Currency field value
func (o *SymbolSearchResponseItem) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *SymbolSearchResponseItem) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *SymbolSearchResponseItem) SetCurrency(v string) {
	o.Currency = v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *SymbolSearchResponseItem) GetAccess() SymbolSearchResponseItemAccess {
	if o == nil || IsNil(o.Access) {
		var ret SymbolSearchResponseItemAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolSearchResponseItem) GetAccessOk() (*SymbolSearchResponseItemAccess, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *SymbolSearchResponseItem) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given SymbolSearchResponseItemAccess and assigns it to the Access field.
func (o *SymbolSearchResponseItem) SetAccess(v SymbolSearchResponseItemAccess) {
	o.Access = &v
}

func (o SymbolSearchResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolSearchResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["instrument_name"] = o.InstrumentName
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["exchange_timezone"] = o.ExchangeTimezone
	toSerialize["instrument_type"] = o.InstrumentType
	toSerialize["country"] = o.Country
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}

func (o *SymbolSearchResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"instrument_name",
		"exchange",
		"mic_code",
		"exchange_timezone",
		"instrument_type",
		"country",
		"currency",
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

	varSymbolSearchResponseItem := _SymbolSearchResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSymbolSearchResponseItem)

	if err != nil {
		return err
	}

	*o = SymbolSearchResponseItem(varSymbolSearchResponseItem)

	return err
}

type NullableSymbolSearchResponseItem struct {
	value *SymbolSearchResponseItem
	isSet bool
}

func (v NullableSymbolSearchResponseItem) Get() *SymbolSearchResponseItem {
	return v.value
}

func (v *NullableSymbolSearchResponseItem) Set(val *SymbolSearchResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolSearchResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolSearchResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolSearchResponseItem(val *SymbolSearchResponseItem) *NullableSymbolSearchResponseItem {
	return &NullableSymbolSearchResponseItem{value: val, isSet: true}
}

func (v NullableSymbolSearchResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolSearchResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
