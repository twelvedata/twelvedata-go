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

// checks if the FundResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FundResponseItem{}

// FundResponseItem FundResponseItem represents details of a fund
type FundResponseItem struct {
	// Instrument symbol (ticker)
	Symbol string `json:"symbol"`
	// Full name of the fund
	Name string `json:"name"`
	// Country where the fund is located
	Country string `json:"country"`
	// Currency of the fund according to the ISO 4217 standard
	Currency string `json:"currency"`
	// Exchange where the fund is traded
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Type of the fund
	Type string `json:"type"`
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

type _FundResponseItem FundResponseItem

// NewFundResponseItem instantiates a new FundResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundResponseItem(symbol string, name string, country string, currency string, exchange string, micCode string, type_ string, figiCode string, cfiCode string, isin string, cusip string) *FundResponseItem {
	this := FundResponseItem{}
	this.Symbol = symbol
	this.Name = name
	this.Country = country
	this.Currency = currency
	this.Exchange = exchange
	this.MicCode = micCode
	this.Type = type_
	this.FigiCode = figiCode
	this.CfiCode = cfiCode
	this.Isin = isin
	this.Cusip = cusip
	return &this
}

// NewFundResponseItemWithDefaults instantiates a new FundResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundResponseItemWithDefaults() *FundResponseItem {
	this := FundResponseItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *FundResponseItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *FundResponseItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *FundResponseItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *FundResponseItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FundResponseItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FundResponseItem) SetName(v string) {
	o.Name = v
}

// GetCountry returns the Country field value
func (o *FundResponseItem) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *FundResponseItem) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *FundResponseItem) SetCountry(v string) {
	o.Country = v
}

// GetCurrency returns the Currency field value
func (o *FundResponseItem) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *FundResponseItem) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *FundResponseItem) SetCurrency(v string) {
	o.Currency = v
}

// GetExchange returns the Exchange field value
func (o *FundResponseItem) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *FundResponseItem) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *FundResponseItem) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *FundResponseItem) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *FundResponseItem) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *FundResponseItem) SetMicCode(v string) {
	o.MicCode = v
}

// GetType returns the Type field value
func (o *FundResponseItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FundResponseItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FundResponseItem) SetType(v string) {
	o.Type = v
}

// GetFigiCode returns the FigiCode field value
func (o *FundResponseItem) GetFigiCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FigiCode
}

// GetFigiCodeOk returns a tuple with the FigiCode field value
// and a boolean to check if the value has been set.
func (o *FundResponseItem) GetFigiCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FigiCode, true
}

// SetFigiCode sets field value
func (o *FundResponseItem) SetFigiCode(v string) {
	o.FigiCode = v
}

// GetCfiCode returns the CfiCode field value
func (o *FundResponseItem) GetCfiCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CfiCode
}

// GetCfiCodeOk returns a tuple with the CfiCode field value
// and a boolean to check if the value has been set.
func (o *FundResponseItem) GetCfiCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CfiCode, true
}

// SetCfiCode sets field value
func (o *FundResponseItem) SetCfiCode(v string) {
	o.CfiCode = v
}

// GetIsin returns the Isin field value
func (o *FundResponseItem) GetIsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Isin
}

// GetIsinOk returns a tuple with the Isin field value
// and a boolean to check if the value has been set.
func (o *FundResponseItem) GetIsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isin, true
}

// SetIsin sets field value
func (o *FundResponseItem) SetIsin(v string) {
	o.Isin = v
}

// GetCusip returns the Cusip field value
func (o *FundResponseItem) GetCusip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value
// and a boolean to check if the value has been set.
func (o *FundResponseItem) GetCusipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cusip, true
}

// SetCusip sets field value
func (o *FundResponseItem) SetCusip(v string) {
	o.Cusip = v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *FundResponseItem) GetAccess() EtfResponseItemAccess {
	if o == nil || IsNil(o.Access) {
		var ret EtfResponseItemAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundResponseItem) GetAccessOk() (*EtfResponseItemAccess, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *FundResponseItem) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EtfResponseItemAccess and assigns it to the Access field.
func (o *FundResponseItem) SetAccess(v EtfResponseItemAccess) {
	o.Access = &v
}

func (o FundResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["country"] = o.Country
	toSerialize["currency"] = o.Currency
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
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

func (o *FundResponseItem) UnmarshalJSON(data []byte) (err error) {
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

	varFundResponseItem := _FundResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFundResponseItem)

	if err != nil {
		return err
	}

	*o = FundResponseItem(varFundResponseItem)

	return err
}

type NullableFundResponseItem struct {
	value *FundResponseItem
	isSet bool
}

func (v NullableFundResponseItem) Get() *FundResponseItem {
	return v.value
}

func (v *NullableFundResponseItem) Set(val *FundResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFundResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFundResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundResponseItem(val *FundResponseItem) *NullableFundResponseItem {
	return &NullableFundResponseItem{value: val, isSet: true}
}

func (v NullableFundResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
