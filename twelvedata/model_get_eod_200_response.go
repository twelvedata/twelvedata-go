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

// checks if the GetEod200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEod200Response{}

// GetEod200Response struct for GetEod200Response
type GetEod200Response struct {
	// Symbol passed
	Symbol string `json:"symbol"`
	// Exchange where instrument is traded
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode *string `json:"mic_code,omitempty"`
	// Currency in which instrument is denominated
	Currency *string `json:"currency,omitempty"`
	// Datetime in defined timezone referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// The most recent end of day close price
	Close string `json:"close"`
}

type _GetEod200Response GetEod200Response

// NewGetEod200Response instantiates a new GetEod200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEod200Response(symbol string, exchange string, datetime string, close string) *GetEod200Response {
	this := GetEod200Response{}
	this.Symbol = symbol
	this.Exchange = exchange
	this.Datetime = datetime
	this.Close = close
	return &this
}

// NewGetEod200ResponseWithDefaults instantiates a new GetEod200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEod200ResponseWithDefaults() *GetEod200Response {
	this := GetEod200Response{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetEod200Response) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetEod200Response) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetEod200Response) SetSymbol(v string) {
	o.Symbol = v
}

// GetExchange returns the Exchange field value
func (o *GetEod200Response) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetEod200Response) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetEod200Response) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *GetEod200Response) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEod200Response) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *GetEod200Response) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *GetEod200Response) SetMicCode(v string) {
	o.MicCode = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetEod200Response) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEod200Response) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetEod200Response) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetEod200Response) SetCurrency(v string) {
	o.Currency = &v
}

// GetDatetime returns the Datetime field value
func (o *GetEod200Response) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetEod200Response) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetEod200Response) SetDatetime(v string) {
	o.Datetime = v
}

// GetClose returns the Close field value
func (o *GetEod200Response) GetClose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Close
}

// GetCloseOk returns a tuple with the Close field value
// and a boolean to check if the value has been set.
func (o *GetEod200Response) GetCloseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Close, true
}

// SetClose sets field value
func (o *GetEod200Response) SetClose(v string) {
	o.Close = v
}

func (o GetEod200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEod200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["exchange"] = o.Exchange
	if !IsNil(o.MicCode) {
		toSerialize["mic_code"] = o.MicCode
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	toSerialize["datetime"] = o.Datetime
	toSerialize["close"] = o.Close
	return toSerialize, nil
}

func (o *GetEod200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"exchange",
		"datetime",
		"close",
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

	varGetEod200Response := _GetEod200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetEod200Response)

	if err != nil {
		return err
	}

	*o = GetEod200Response(varGetEod200Response)

	return err
}

type NullableGetEod200Response struct {
	value *GetEod200Response
	isSet bool
}

func (v NullableGetEod200Response) Get() *GetEod200Response {
	return v.value
}

func (v *NullableGetEod200Response) Set(val *GetEod200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEod200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEod200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEod200Response(val *GetEod200Response) *NullableGetEod200Response {
	return &NullableGetEod200Response{value: val, isSet: true}
}

func (v NullableGetEod200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEod200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
