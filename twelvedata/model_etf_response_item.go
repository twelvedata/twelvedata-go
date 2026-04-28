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

// checks if the EtfResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EtfResponseItem{}

// EtfResponseItem EtfResponseItem represents details of an ETF
type EtfResponseItem struct {
	// Instrument symbol (ticker)
	Symbol string `json:"symbol"`
	// Full name of the ETF
	Name string `json:"name"`
	// Currency of the ETF according to the ISO 4217 standard
	Currency string `json:"currency"`
	// Exchange where the ETF is traded
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Country where the ETF is located
	Country string `json:"country"`
	// Financial instrument global identifier (FIGI)
	FigiCode string `json:"figi_code"`
	// Classification of Financial Instruments (CFI)
	CfiCode string `json:"cfi_code"`
	// International securities identification number (ISIN)
	Isin string `json:"isin"`
	// A unique nine-character alphanumeric code used to identify financial securities, ensuring accurate data retrieval for the specified asset
	Cusip  string                 `json:"cusip"`
	Access *EtfResponseItemAccess `json:"access,omitempty"`
}

type _EtfResponseItem EtfResponseItem

// NewEtfResponseItem instantiates a new EtfResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEtfResponseItem(symbol string, name string, currency string, exchange string, micCode string, country string, figiCode string, cfiCode string, isin string, cusip string) *EtfResponseItem {
	this := EtfResponseItem{}
	this.Symbol = symbol
	this.Name = name
	this.Currency = currency
	this.Exchange = exchange
	this.MicCode = micCode
	this.Country = country
	this.FigiCode = figiCode
	this.CfiCode = cfiCode
	this.Isin = isin
	this.Cusip = cusip
	return &this
}

// NewEtfResponseItemWithDefaults instantiates a new EtfResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEtfResponseItemWithDefaults() *EtfResponseItem {
	this := EtfResponseItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *EtfResponseItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *EtfResponseItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *EtfResponseItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EtfResponseItem) SetName(v string) {
	o.Name = v
}

// GetCurrency returns the Currency field value
func (o *EtfResponseItem) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItem) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *EtfResponseItem) SetCurrency(v string) {
	o.Currency = v
}

// GetExchange returns the Exchange field value
func (o *EtfResponseItem) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItem) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *EtfResponseItem) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *EtfResponseItem) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItem) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *EtfResponseItem) SetMicCode(v string) {
	o.MicCode = v
}

// GetCountry returns the Country field value
func (o *EtfResponseItem) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItem) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *EtfResponseItem) SetCountry(v string) {
	o.Country = v
}

// GetFigiCode returns the FigiCode field value
func (o *EtfResponseItem) GetFigiCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FigiCode
}

// GetFigiCodeOk returns a tuple with the FigiCode field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItem) GetFigiCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FigiCode, true
}

// SetFigiCode sets field value
func (o *EtfResponseItem) SetFigiCode(v string) {
	o.FigiCode = v
}

// GetCfiCode returns the CfiCode field value
func (o *EtfResponseItem) GetCfiCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CfiCode
}

// GetCfiCodeOk returns a tuple with the CfiCode field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItem) GetCfiCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CfiCode, true
}

// SetCfiCode sets field value
func (o *EtfResponseItem) SetCfiCode(v string) {
	o.CfiCode = v
}

// GetIsin returns the Isin field value
func (o *EtfResponseItem) GetIsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Isin
}

// GetIsinOk returns a tuple with the Isin field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItem) GetIsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isin, true
}

// SetIsin sets field value
func (o *EtfResponseItem) SetIsin(v string) {
	o.Isin = v
}

// GetCusip returns the Cusip field value
func (o *EtfResponseItem) GetCusip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItem) GetCusipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cusip, true
}

// SetCusip sets field value
func (o *EtfResponseItem) SetCusip(v string) {
	o.Cusip = v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *EtfResponseItem) GetAccess() EtfResponseItemAccess {
	if o == nil || IsNil(o.Access) {
		var ret EtfResponseItemAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtfResponseItem) GetAccessOk() (*EtfResponseItemAccess, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *EtfResponseItem) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EtfResponseItemAccess and assigns it to the Access field.
func (o *EtfResponseItem) SetAccess(v EtfResponseItemAccess) {
	o.Access = &v
}

func (o EtfResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EtfResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["currency"] = o.Currency
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["country"] = o.Country
	toSerialize["figi_code"] = o.FigiCode
	toSerialize["cfi_code"] = o.CfiCode
	toSerialize["isin"] = o.Isin
	toSerialize["cusip"] = o.Cusip
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}

func (o *EtfResponseItem) UnmarshalJSON(data []byte) (err error) {
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

	varEtfResponseItem := _EtfResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEtfResponseItem)

	if err != nil {
		return err
	}

	*o = EtfResponseItem(varEtfResponseItem)

	return err
}

type NullableEtfResponseItem struct {
	value *EtfResponseItem
	isSet bool
}

func (v NullableEtfResponseItem) Get() *EtfResponseItem {
	return v.value
}

func (v *NullableEtfResponseItem) Set(val *EtfResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEtfResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEtfResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtfResponseItem(val *EtfResponseItem) *NullableEtfResponseItem {
	return &NullableEtfResponseItem{value: val, isSet: true}
}

func (v NullableEtfResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtfResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
