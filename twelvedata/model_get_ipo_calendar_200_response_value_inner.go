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

// checks if the GetIpoCalendar200ResponseValueInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIpoCalendar200ResponseValueInner{}

// GetIpoCalendar200ResponseValueInner struct for GetIpoCalendar200ResponseValueInner
type GetIpoCalendar200ResponseValueInner struct {
	// Ticker of the company
	Symbol string `json:"symbol"`
	// Name of the company
	Name string `json:"name"`
	// Exchange name where the company is listed
	Exchange string `json:"exchange"`
	// Market Identifier Code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// The lower bound of stock price range if available
	PriceRangeLow *float64 `json:"price_range_low,omitempty"`
	// The upper bound of stock price range if available
	PriceRangeHigh *float64 `json:"price_range_high,omitempty"`
	// Initial offer price if available
	OfferPrice *float64 `json:"offer_price,omitempty"`
	// Currency of the stock
	Currency string `json:"currency"`
	// Number of shares, if available
	Shares *int64 `json:"shares,omitempty"`
}

type _GetIpoCalendar200ResponseValueInner GetIpoCalendar200ResponseValueInner

// NewGetIpoCalendar200ResponseValueInner instantiates a new GetIpoCalendar200ResponseValueInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIpoCalendar200ResponseValueInner(symbol string, name string, exchange string, micCode string, currency string) *GetIpoCalendar200ResponseValueInner {
	this := GetIpoCalendar200ResponseValueInner{}
	this.Symbol = symbol
	this.Name = name
	this.Exchange = exchange
	this.MicCode = micCode
	this.Currency = currency
	return &this
}

// NewGetIpoCalendar200ResponseValueInnerWithDefaults instantiates a new GetIpoCalendar200ResponseValueInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIpoCalendar200ResponseValueInnerWithDefaults() *GetIpoCalendar200ResponseValueInner {
	this := GetIpoCalendar200ResponseValueInner{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetIpoCalendar200ResponseValueInner) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetIpoCalendar200ResponseValueInner) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetIpoCalendar200ResponseValueInner) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *GetIpoCalendar200ResponseValueInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetIpoCalendar200ResponseValueInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetIpoCalendar200ResponseValueInner) SetName(v string) {
	o.Name = v
}

// GetExchange returns the Exchange field value
func (o *GetIpoCalendar200ResponseValueInner) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetIpoCalendar200ResponseValueInner) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetIpoCalendar200ResponseValueInner) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *GetIpoCalendar200ResponseValueInner) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *GetIpoCalendar200ResponseValueInner) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *GetIpoCalendar200ResponseValueInner) SetMicCode(v string) {
	o.MicCode = v
}

// GetPriceRangeLow returns the PriceRangeLow field value if set, zero value otherwise.
func (o *GetIpoCalendar200ResponseValueInner) GetPriceRangeLow() float64 {
	if o == nil || IsNil(o.PriceRangeLow) {
		var ret float64
		return ret
	}
	return *o.PriceRangeLow
}

// GetPriceRangeLowOk returns a tuple with the PriceRangeLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpoCalendar200ResponseValueInner) GetPriceRangeLowOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceRangeLow) {
		return nil, false
	}
	return o.PriceRangeLow, true
}

// HasPriceRangeLow returns a boolean if a field has been set.
func (o *GetIpoCalendar200ResponseValueInner) HasPriceRangeLow() bool {
	if o != nil && !IsNil(o.PriceRangeLow) {
		return true
	}

	return false
}

// SetPriceRangeLow gets a reference to the given float64 and assigns it to the PriceRangeLow field.
func (o *GetIpoCalendar200ResponseValueInner) SetPriceRangeLow(v float64) {
	o.PriceRangeLow = &v
}

// GetPriceRangeHigh returns the PriceRangeHigh field value if set, zero value otherwise.
func (o *GetIpoCalendar200ResponseValueInner) GetPriceRangeHigh() float64 {
	if o == nil || IsNil(o.PriceRangeHigh) {
		var ret float64
		return ret
	}
	return *o.PriceRangeHigh
}

// GetPriceRangeHighOk returns a tuple with the PriceRangeHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpoCalendar200ResponseValueInner) GetPriceRangeHighOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceRangeHigh) {
		return nil, false
	}
	return o.PriceRangeHigh, true
}

// HasPriceRangeHigh returns a boolean if a field has been set.
func (o *GetIpoCalendar200ResponseValueInner) HasPriceRangeHigh() bool {
	if o != nil && !IsNil(o.PriceRangeHigh) {
		return true
	}

	return false
}

// SetPriceRangeHigh gets a reference to the given float64 and assigns it to the PriceRangeHigh field.
func (o *GetIpoCalendar200ResponseValueInner) SetPriceRangeHigh(v float64) {
	o.PriceRangeHigh = &v
}

// GetOfferPrice returns the OfferPrice field value if set, zero value otherwise.
func (o *GetIpoCalendar200ResponseValueInner) GetOfferPrice() float64 {
	if o == nil || IsNil(o.OfferPrice) {
		var ret float64
		return ret
	}
	return *o.OfferPrice
}

// GetOfferPriceOk returns a tuple with the OfferPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpoCalendar200ResponseValueInner) GetOfferPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.OfferPrice) {
		return nil, false
	}
	return o.OfferPrice, true
}

// HasOfferPrice returns a boolean if a field has been set.
func (o *GetIpoCalendar200ResponseValueInner) HasOfferPrice() bool {
	if o != nil && !IsNil(o.OfferPrice) {
		return true
	}

	return false
}

// SetOfferPrice gets a reference to the given float64 and assigns it to the OfferPrice field.
func (o *GetIpoCalendar200ResponseValueInner) SetOfferPrice(v float64) {
	o.OfferPrice = &v
}

// GetCurrency returns the Currency field value
func (o *GetIpoCalendar200ResponseValueInner) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GetIpoCalendar200ResponseValueInner) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GetIpoCalendar200ResponseValueInner) SetCurrency(v string) {
	o.Currency = v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *GetIpoCalendar200ResponseValueInner) GetShares() int64 {
	if o == nil || IsNil(o.Shares) {
		var ret int64
		return ret
	}
	return *o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpoCalendar200ResponseValueInner) GetSharesOk() (*int64, bool) {
	if o == nil || IsNil(o.Shares) {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *GetIpoCalendar200ResponseValueInner) HasShares() bool {
	if o != nil && !IsNil(o.Shares) {
		return true
	}

	return false
}

// SetShares gets a reference to the given int64 and assigns it to the Shares field.
func (o *GetIpoCalendar200ResponseValueInner) SetShares(v int64) {
	o.Shares = &v
}

func (o GetIpoCalendar200ResponseValueInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIpoCalendar200ResponseValueInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	if !IsNil(o.PriceRangeLow) {
		toSerialize["price_range_low"] = o.PriceRangeLow
	}
	if !IsNil(o.PriceRangeHigh) {
		toSerialize["price_range_high"] = o.PriceRangeHigh
	}
	if !IsNil(o.OfferPrice) {
		toSerialize["offer_price"] = o.OfferPrice
	}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Shares) {
		toSerialize["shares"] = o.Shares
	}
	return toSerialize, nil
}

func (o *GetIpoCalendar200ResponseValueInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"exchange",
		"mic_code",
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

	varGetIpoCalendar200ResponseValueInner := _GetIpoCalendar200ResponseValueInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetIpoCalendar200ResponseValueInner)

	if err != nil {
		return err
	}

	*o = GetIpoCalendar200ResponseValueInner(varGetIpoCalendar200ResponseValueInner)

	return err
}

type NullableGetIpoCalendar200ResponseValueInner struct {
	value *GetIpoCalendar200ResponseValueInner
	isSet bool
}

func (v NullableGetIpoCalendar200ResponseValueInner) Get() *GetIpoCalendar200ResponseValueInner {
	return v.value
}

func (v *NullableGetIpoCalendar200ResponseValueInner) Set(val *GetIpoCalendar200ResponseValueInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIpoCalendar200ResponseValueInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIpoCalendar200ResponseValueInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIpoCalendar200ResponseValueInner(val *GetIpoCalendar200ResponseValueInner) *NullableGetIpoCalendar200ResponseValueInner {
	return &NullableGetIpoCalendar200ResponseValueInner{value: val, isSet: true}
}

func (v NullableGetIpoCalendar200ResponseValueInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIpoCalendar200ResponseValueInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
