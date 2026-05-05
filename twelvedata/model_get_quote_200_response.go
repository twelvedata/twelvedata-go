// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetQuote200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetQuote200Response{}

// GetQuote200Response struct for GetQuote200Response
type GetQuote200Response struct {
	// Symbol passed
	Symbol string `json:"symbol"`
	// Name of the instrument
	Name string `json:"name"`
	// Exchange where instrument is traded
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard. Available for stocks, ETFs, mutual funds, bonds
	MicCode *string `json:"mic_code,omitempty"`
	// Currency in which the equity is denominated. Available for stocks, ETFs, mutual funds, bonds
	Currency *string `json:"currency,omitempty"`
	// Datetime in defined timezone referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Unix timestamp representing the opening candle of the specified interval
	Timestamp int64 `json:"timestamp"`
	// Unix timestamp of last minute candle
	LastQuoteAt *int64 `json:"last_quote_at,omitempty"`
	// Price at the opening of current bar
	Open string `json:"open"`
	// Highest price which occurred during the current bar
	High string `json:"high"`
	// Lowest price which occurred during the current bar
	Low string `json:"low"`
	// Close price at the end of the bar
	Close string `json:"close"`
	// Trading volume during the bar. Available not for all instrument types
	Volume *string `json:"volume,omitempty"`
	// Close price at the end of the previous bar
	PreviousClose string `json:"previous_close"`
	// Close - previous_close
	Change string `json:"change"`
	// (Close - previous_close) / previous_close * 100
	PercentChange string `json:"percent_change"`
	// Average volume of the specified period. Available not for all instrument types
	AverageVolume *string `json:"average_volume,omitempty"`
	// Percent change in price between the current and the backward one, where period is 1 day. Available for crypto
	Rolling1dChange *string `json:"rolling_1d_change,omitempty"`
	// Percent change in price between the current and the backward one, where period is 7 days. Available for crypto
	Rolling7dChange *string `json:"rolling_7d_change,omitempty"`
	// Percent change in price between the current and the backward one, where period specified in request param rolling_period. Available for crypto
	RollingChange *string `json:"rolling_change,omitempty"`
	// True if market is open; false if closed
	IsMarketOpen bool                            `json:"is_market_open"`
	FiftyTwoWeek GetQuote200ResponseFiftyTwoWeek `json:"fifty_two_week"`
	// Diff between the regular close price and the latest extended price. Displayed only if prepost is true
	ExtendedChange *string `json:"extended_change,omitempty"`
	// Percent change in price between the regular close price and the latest extended price. Displayed only if prepost is true
	ExtendedPercentChange *string `json:"extended_percent_change,omitempty"`
	// Latest extended price. Displayed only if prepost is true
	ExtendedPrice *string `json:"extended_price,omitempty"`
	// Unix timestamp of the last extended price. Displayed only if prepost is true
	ExtendedTimestamp *int64 `json:"extended_timestamp,omitempty"`
}

type _GetQuote200Response GetQuote200Response

// NewGetQuote200Response instantiates a new GetQuote200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetQuote200Response(symbol string, name string, exchange string, datetime string, timestamp int64, open string, high string, low string, close string, previousClose string, change string, percentChange string, isMarketOpen bool, fiftyTwoWeek GetQuote200ResponseFiftyTwoWeek) *GetQuote200Response {
	this := GetQuote200Response{}
	this.Symbol = symbol
	this.Name = name
	this.Exchange = exchange
	this.Datetime = datetime
	this.Timestamp = timestamp
	this.Open = open
	this.High = high
	this.Low = low
	this.Close = close
	this.PreviousClose = previousClose
	this.Change = change
	this.PercentChange = percentChange
	this.IsMarketOpen = isMarketOpen
	this.FiftyTwoWeek = fiftyTwoWeek
	return &this
}

// NewGetQuote200ResponseWithDefaults instantiates a new GetQuote200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetQuote200ResponseWithDefaults() *GetQuote200Response {
	this := GetQuote200Response{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetQuote200Response) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetQuote200Response) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *GetQuote200Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetQuote200Response) SetName(v string) {
	o.Name = v
}

// GetExchange returns the Exchange field value
func (o *GetQuote200Response) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetQuote200Response) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *GetQuote200Response) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *GetQuote200Response) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *GetQuote200Response) SetMicCode(v string) {
	o.MicCode = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetQuote200Response) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetQuote200Response) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetQuote200Response) SetCurrency(v string) {
	o.Currency = &v
}

// GetDatetime returns the Datetime field value
func (o *GetQuote200Response) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetQuote200Response) SetDatetime(v string) {
	o.Datetime = v
}

// GetTimestamp returns the Timestamp field value
func (o *GetQuote200Response) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *GetQuote200Response) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetLastQuoteAt returns the LastQuoteAt field value if set, zero value otherwise.
func (o *GetQuote200Response) GetLastQuoteAt() int64 {
	if o == nil || IsNil(o.LastQuoteAt) {
		var ret int64
		return ret
	}
	return *o.LastQuoteAt
}

// GetLastQuoteAtOk returns a tuple with the LastQuoteAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetLastQuoteAtOk() (*int64, bool) {
	if o == nil || IsNil(o.LastQuoteAt) {
		return nil, false
	}
	return o.LastQuoteAt, true
}

// HasLastQuoteAt returns a boolean if a field has been set.
func (o *GetQuote200Response) HasLastQuoteAt() bool {
	if o != nil && !IsNil(o.LastQuoteAt) {
		return true
	}

	return false
}

// SetLastQuoteAt gets a reference to the given int64 and assigns it to the LastQuoteAt field.
func (o *GetQuote200Response) SetLastQuoteAt(v int64) {
	o.LastQuoteAt = &v
}

// GetOpen returns the Open field value
func (o *GetQuote200Response) GetOpen() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Open
}

// GetOpenOk returns a tuple with the Open field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetOpenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Open, true
}

// SetOpen sets field value
func (o *GetQuote200Response) SetOpen(v string) {
	o.Open = v
}

// GetHigh returns the High field value
func (o *GetQuote200Response) GetHigh() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.High
}

// GetHighOk returns a tuple with the High field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetHighOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.High, true
}

// SetHigh sets field value
func (o *GetQuote200Response) SetHigh(v string) {
	o.High = v
}

// GetLow returns the Low field value
func (o *GetQuote200Response) GetLow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Low
}

// GetLowOk returns a tuple with the Low field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetLowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Low, true
}

// SetLow sets field value
func (o *GetQuote200Response) SetLow(v string) {
	o.Low = v
}

// GetClose returns the Close field value
func (o *GetQuote200Response) GetClose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Close
}

// GetCloseOk returns a tuple with the Close field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetCloseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Close, true
}

// SetClose sets field value
func (o *GetQuote200Response) SetClose(v string) {
	o.Close = v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *GetQuote200Response) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *GetQuote200Response) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *GetQuote200Response) SetVolume(v string) {
	o.Volume = &v
}

// GetPreviousClose returns the PreviousClose field value
func (o *GetQuote200Response) GetPreviousClose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousClose
}

// GetPreviousCloseOk returns a tuple with the PreviousClose field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetPreviousCloseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousClose, true
}

// SetPreviousClose sets field value
func (o *GetQuote200Response) SetPreviousClose(v string) {
	o.PreviousClose = v
}

// GetChange returns the Change field value
func (o *GetQuote200Response) GetChange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Change
}

// GetChangeOk returns a tuple with the Change field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetChangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Change, true
}

// SetChange sets field value
func (o *GetQuote200Response) SetChange(v string) {
	o.Change = v
}

// GetPercentChange returns the PercentChange field value
func (o *GetQuote200Response) GetPercentChange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PercentChange
}

// GetPercentChangeOk returns a tuple with the PercentChange field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetPercentChangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentChange, true
}

// SetPercentChange sets field value
func (o *GetQuote200Response) SetPercentChange(v string) {
	o.PercentChange = v
}

// GetAverageVolume returns the AverageVolume field value if set, zero value otherwise.
func (o *GetQuote200Response) GetAverageVolume() string {
	if o == nil || IsNil(o.AverageVolume) {
		var ret string
		return ret
	}
	return *o.AverageVolume
}

// GetAverageVolumeOk returns a tuple with the AverageVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetAverageVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.AverageVolume) {
		return nil, false
	}
	return o.AverageVolume, true
}

// HasAverageVolume returns a boolean if a field has been set.
func (o *GetQuote200Response) HasAverageVolume() bool {
	if o != nil && !IsNil(o.AverageVolume) {
		return true
	}

	return false
}

// SetAverageVolume gets a reference to the given string and assigns it to the AverageVolume field.
func (o *GetQuote200Response) SetAverageVolume(v string) {
	o.AverageVolume = &v
}

// GetRolling1dChange returns the Rolling1dChange field value if set, zero value otherwise.
func (o *GetQuote200Response) GetRolling1dChange() string {
	if o == nil || IsNil(o.Rolling1dChange) {
		var ret string
		return ret
	}
	return *o.Rolling1dChange
}

// GetRolling1dChangeOk returns a tuple with the Rolling1dChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetRolling1dChangeOk() (*string, bool) {
	if o == nil || IsNil(o.Rolling1dChange) {
		return nil, false
	}
	return o.Rolling1dChange, true
}

// HasRolling1dChange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasRolling1dChange() bool {
	if o != nil && !IsNil(o.Rolling1dChange) {
		return true
	}

	return false
}

// SetRolling1dChange gets a reference to the given string and assigns it to the Rolling1dChange field.
func (o *GetQuote200Response) SetRolling1dChange(v string) {
	o.Rolling1dChange = &v
}

// GetRolling7dChange returns the Rolling7dChange field value if set, zero value otherwise.
func (o *GetQuote200Response) GetRolling7dChange() string {
	if o == nil || IsNil(o.Rolling7dChange) {
		var ret string
		return ret
	}
	return *o.Rolling7dChange
}

// GetRolling7dChangeOk returns a tuple with the Rolling7dChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetRolling7dChangeOk() (*string, bool) {
	if o == nil || IsNil(o.Rolling7dChange) {
		return nil, false
	}
	return o.Rolling7dChange, true
}

// HasRolling7dChange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasRolling7dChange() bool {
	if o != nil && !IsNil(o.Rolling7dChange) {
		return true
	}

	return false
}

// SetRolling7dChange gets a reference to the given string and assigns it to the Rolling7dChange field.
func (o *GetQuote200Response) SetRolling7dChange(v string) {
	o.Rolling7dChange = &v
}

// GetRollingChange returns the RollingChange field value if set, zero value otherwise.
func (o *GetQuote200Response) GetRollingChange() string {
	if o == nil || IsNil(o.RollingChange) {
		var ret string
		return ret
	}
	return *o.RollingChange
}

// GetRollingChangeOk returns a tuple with the RollingChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetRollingChangeOk() (*string, bool) {
	if o == nil || IsNil(o.RollingChange) {
		return nil, false
	}
	return o.RollingChange, true
}

// HasRollingChange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasRollingChange() bool {
	if o != nil && !IsNil(o.RollingChange) {
		return true
	}

	return false
}

// SetRollingChange gets a reference to the given string and assigns it to the RollingChange field.
func (o *GetQuote200Response) SetRollingChange(v string) {
	o.RollingChange = &v
}

// GetIsMarketOpen returns the IsMarketOpen field value
func (o *GetQuote200Response) GetIsMarketOpen() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMarketOpen
}

// GetIsMarketOpenOk returns a tuple with the IsMarketOpen field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetIsMarketOpenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMarketOpen, true
}

// SetIsMarketOpen sets field value
func (o *GetQuote200Response) SetIsMarketOpen(v bool) {
	o.IsMarketOpen = v
}

// GetFiftyTwoWeek returns the FiftyTwoWeek field value
func (o *GetQuote200Response) GetFiftyTwoWeek() GetQuote200ResponseFiftyTwoWeek {
	if o == nil {
		var ret GetQuote200ResponseFiftyTwoWeek
		return ret
	}

	return o.FiftyTwoWeek
}

// GetFiftyTwoWeekOk returns a tuple with the FiftyTwoWeek field value
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetFiftyTwoWeekOk() (*GetQuote200ResponseFiftyTwoWeek, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiftyTwoWeek, true
}

// SetFiftyTwoWeek sets field value
func (o *GetQuote200Response) SetFiftyTwoWeek(v GetQuote200ResponseFiftyTwoWeek) {
	o.FiftyTwoWeek = v
}

// GetExtendedChange returns the ExtendedChange field value if set, zero value otherwise.
func (o *GetQuote200Response) GetExtendedChange() string {
	if o == nil || IsNil(o.ExtendedChange) {
		var ret string
		return ret
	}
	return *o.ExtendedChange
}

// GetExtendedChangeOk returns a tuple with the ExtendedChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetExtendedChangeOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendedChange) {
		return nil, false
	}
	return o.ExtendedChange, true
}

// HasExtendedChange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasExtendedChange() bool {
	if o != nil && !IsNil(o.ExtendedChange) {
		return true
	}

	return false
}

// SetExtendedChange gets a reference to the given string and assigns it to the ExtendedChange field.
func (o *GetQuote200Response) SetExtendedChange(v string) {
	o.ExtendedChange = &v
}

// GetExtendedPercentChange returns the ExtendedPercentChange field value if set, zero value otherwise.
func (o *GetQuote200Response) GetExtendedPercentChange() string {
	if o == nil || IsNil(o.ExtendedPercentChange) {
		var ret string
		return ret
	}
	return *o.ExtendedPercentChange
}

// GetExtendedPercentChangeOk returns a tuple with the ExtendedPercentChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetExtendedPercentChangeOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendedPercentChange) {
		return nil, false
	}
	return o.ExtendedPercentChange, true
}

// HasExtendedPercentChange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasExtendedPercentChange() bool {
	if o != nil && !IsNil(o.ExtendedPercentChange) {
		return true
	}

	return false
}

// SetExtendedPercentChange gets a reference to the given string and assigns it to the ExtendedPercentChange field.
func (o *GetQuote200Response) SetExtendedPercentChange(v string) {
	o.ExtendedPercentChange = &v
}

// GetExtendedPrice returns the ExtendedPrice field value if set, zero value otherwise.
func (o *GetQuote200Response) GetExtendedPrice() string {
	if o == nil || IsNil(o.ExtendedPrice) {
		var ret string
		return ret
	}
	return *o.ExtendedPrice
}

// GetExtendedPriceOk returns a tuple with the ExtendedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetExtendedPriceOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendedPrice) {
		return nil, false
	}
	return o.ExtendedPrice, true
}

// HasExtendedPrice returns a boolean if a field has been set.
func (o *GetQuote200Response) HasExtendedPrice() bool {
	if o != nil && !IsNil(o.ExtendedPrice) {
		return true
	}

	return false
}

// SetExtendedPrice gets a reference to the given string and assigns it to the ExtendedPrice field.
func (o *GetQuote200Response) SetExtendedPrice(v string) {
	o.ExtendedPrice = &v
}

// GetExtendedTimestamp returns the ExtendedTimestamp field value if set, zero value otherwise.
func (o *GetQuote200Response) GetExtendedTimestamp() int64 {
	if o == nil || IsNil(o.ExtendedTimestamp) {
		var ret int64
		return ret
	}
	return *o.ExtendedTimestamp
}

// GetExtendedTimestampOk returns a tuple with the ExtendedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetExtendedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.ExtendedTimestamp) {
		return nil, false
	}
	return o.ExtendedTimestamp, true
}

// HasExtendedTimestamp returns a boolean if a field has been set.
func (o *GetQuote200Response) HasExtendedTimestamp() bool {
	if o != nil && !IsNil(o.ExtendedTimestamp) {
		return true
	}

	return false
}

// SetExtendedTimestamp gets a reference to the given int64 and assigns it to the ExtendedTimestamp field.
func (o *GetQuote200Response) SetExtendedTimestamp(v int64) {
	o.ExtendedTimestamp = &v
}

func (o GetQuote200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetQuote200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["exchange"] = o.Exchange
	if !IsNil(o.MicCode) {
		toSerialize["mic_code"] = o.MicCode
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	toSerialize["datetime"] = o.Datetime
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.LastQuoteAt) {
		toSerialize["last_quote_at"] = o.LastQuoteAt
	}
	toSerialize["open"] = o.Open
	toSerialize["high"] = o.High
	toSerialize["low"] = o.Low
	toSerialize["close"] = o.Close
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	toSerialize["previous_close"] = o.PreviousClose
	toSerialize["change"] = o.Change
	toSerialize["percent_change"] = o.PercentChange
	if !IsNil(o.AverageVolume) {
		toSerialize["average_volume"] = o.AverageVolume
	}
	if !IsNil(o.Rolling1dChange) {
		toSerialize["rolling_1d_change"] = o.Rolling1dChange
	}
	if !IsNil(o.Rolling7dChange) {
		toSerialize["rolling_7d_change"] = o.Rolling7dChange
	}
	if !IsNil(o.RollingChange) {
		toSerialize["rolling_change"] = o.RollingChange
	}
	toSerialize["is_market_open"] = o.IsMarketOpen
	toSerialize["fifty_two_week"] = o.FiftyTwoWeek
	if !IsNil(o.ExtendedChange) {
		toSerialize["extended_change"] = o.ExtendedChange
	}
	if !IsNil(o.ExtendedPercentChange) {
		toSerialize["extended_percent_change"] = o.ExtendedPercentChange
	}
	if !IsNil(o.ExtendedPrice) {
		toSerialize["extended_price"] = o.ExtendedPrice
	}
	if !IsNil(o.ExtendedTimestamp) {
		toSerialize["extended_timestamp"] = o.ExtendedTimestamp
	}
	return toSerialize, nil
}

func (o *GetQuote200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"exchange",
		"datetime",
		"timestamp",
		"open",
		"high",
		"low",
		"close",
		"previous_close",
		"change",
		"percent_change",
		"is_market_open",
		"fifty_two_week",
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

	varGetQuote200Response := _GetQuote200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetQuote200Response)

	if err != nil {
		return err
	}

	*o = GetQuote200Response(varGetQuote200Response)

	return err
}

type NullableGetQuote200Response struct {
	value *GetQuote200Response
	isSet bool
}

func (v NullableGetQuote200Response) Get() *GetQuote200Response {
	return v.value
}

func (v *NullableGetQuote200Response) Set(val *GetQuote200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuote200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuote200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuote200Response(val *GetQuote200Response) *NullableGetQuote200Response {
	return &NullableGetQuote200Response{value: val, isSet: true}
}

func (v NullableGetQuote200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuote200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
