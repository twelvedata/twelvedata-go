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

// checks if the StocksResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StocksResponseItem{}

// StocksResponseItem StocksResponseItem represents details of a stock instrument
type StocksResponseItem struct {
	// Instrument symbol (ticker)
	Symbol string `json:"symbol"`
	// Full name of instrument
	Name string `json:"name"`
	// Currency of the instrument according to the ISO 4217 standard
	Currency string `json:"currency"`
	// Exchange where instrument is traded
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Country where exchange is located
	Country string `json:"country"`
	// Common issue type
	Type string `json:"type"`
	// Financial instrument global identifier (FIGI)
	FigiCode string `json:"figi_code"`
	// Classification of Financial Instruments (CFI)
	CfiCode string `json:"cfi_code"`
	// International securities identification number (ISIN), available by individual request to support
	Isin string `json:"isin"`
	// A unique nine-character alphanumeric code used to identify financial securities, ensuring accurate data retrieval for the specified asset
	Cusip  string                 `json:"cusip"`
	Access *EtfResponseItemAccess `json:"access,omitempty"`
}

type _StocksResponseItem StocksResponseItem

// NewStocksResponseItem instantiates a new StocksResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStocksResponseItem(symbol string, name string, currency string, exchange string, micCode string, country string, type_ string, figiCode string, cfiCode string, isin string, cusip string) *StocksResponseItem {
	this := StocksResponseItem{}
	this.Symbol = symbol
	this.Name = name
	this.Currency = currency
	this.Exchange = exchange
	this.MicCode = micCode
	this.Country = country
	this.Type = type_
	this.FigiCode = figiCode
	this.CfiCode = cfiCode
	this.Isin = isin
	this.Cusip = cusip
	return &this
}

// NewStocksResponseItemWithDefaults instantiates a new StocksResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStocksResponseItemWithDefaults() *StocksResponseItem {
	this := StocksResponseItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *StocksResponseItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *StocksResponseItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *StocksResponseItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *StocksResponseItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StocksResponseItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StocksResponseItem) SetName(v string) {
	o.Name = v
}

// GetCurrency returns the Currency field value
func (o *StocksResponseItem) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *StocksResponseItem) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *StocksResponseItem) SetCurrency(v string) {
	o.Currency = v
}

// GetExchange returns the Exchange field value
func (o *StocksResponseItem) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *StocksResponseItem) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *StocksResponseItem) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *StocksResponseItem) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *StocksResponseItem) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *StocksResponseItem) SetMicCode(v string) {
	o.MicCode = v
}

// GetCountry returns the Country field value
func (o *StocksResponseItem) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *StocksResponseItem) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *StocksResponseItem) SetCountry(v string) {
	o.Country = v
}

// GetType returns the Type field value
func (o *StocksResponseItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StocksResponseItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StocksResponseItem) SetType(v string) {
	o.Type = v
}

// GetFigiCode returns the FigiCode field value
func (o *StocksResponseItem) GetFigiCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FigiCode
}

// GetFigiCodeOk returns a tuple with the FigiCode field value
// and a boolean to check if the value has been set.
func (o *StocksResponseItem) GetFigiCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FigiCode, true
}

// SetFigiCode sets field value
func (o *StocksResponseItem) SetFigiCode(v string) {
	o.FigiCode = v
}

// GetCfiCode returns the CfiCode field value
func (o *StocksResponseItem) GetCfiCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CfiCode
}

// GetCfiCodeOk returns a tuple with the CfiCode field value
// and a boolean to check if the value has been set.
func (o *StocksResponseItem) GetCfiCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CfiCode, true
}

// SetCfiCode sets field value
func (o *StocksResponseItem) SetCfiCode(v string) {
	o.CfiCode = v
}

// GetIsin returns the Isin field value
func (o *StocksResponseItem) GetIsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Isin
}

// GetIsinOk returns a tuple with the Isin field value
// and a boolean to check if the value has been set.
func (o *StocksResponseItem) GetIsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isin, true
}

// SetIsin sets field value
func (o *StocksResponseItem) SetIsin(v string) {
	o.Isin = v
}

// GetCusip returns the Cusip field value
func (o *StocksResponseItem) GetCusip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value
// and a boolean to check if the value has been set.
func (o *StocksResponseItem) GetCusipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cusip, true
}

// SetCusip sets field value
func (o *StocksResponseItem) SetCusip(v string) {
	o.Cusip = v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *StocksResponseItem) GetAccess() EtfResponseItemAccess {
	if o == nil || IsNil(o.Access) {
		var ret EtfResponseItemAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StocksResponseItem) GetAccessOk() (*EtfResponseItemAccess, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *StocksResponseItem) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EtfResponseItemAccess and assigns it to the Access field.
func (o *StocksResponseItem) SetAccess(v EtfResponseItemAccess) {
	o.Access = &v
}

func (o StocksResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StocksResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["currency"] = o.Currency
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["country"] = o.Country
	toSerialize["type"] = o.Type
	toSerialize["figi_code"] = o.FigiCode
	toSerialize["cfi_code"] = o.CfiCode
	toSerialize["isin"] = o.Isin
	toSerialize["cusip"] = o.Cusip
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}

func (o *StocksResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"currency",
		"exchange",
		"mic_code",
		"country",
		"type",
		"figi_code",
		"cfi_code",
		"isin",
		"cusip",
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

	varStocksResponseItem := _StocksResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStocksResponseItem)

	if err != nil {
		return err
	}

	*o = StocksResponseItem(varStocksResponseItem)

	return err
}

type NullableStocksResponseItem struct {
	value *StocksResponseItem
	isSet bool
}

func (v NullableStocksResponseItem) Get() *StocksResponseItem {
	return v.value
}

func (v *NullableStocksResponseItem) Set(val *StocksResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStocksResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStocksResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStocksResponseItem(val *StocksResponseItem) *NullableStocksResponseItem {
	return &NullableStocksResponseItem{value: val, isSet: true}
}

func (v NullableStocksResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStocksResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
