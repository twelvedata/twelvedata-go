// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the DividendsCalendarItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DividendsCalendarItem{}

// DividendsCalendarItem DividendsCalendarItem represents details of a dividend
type DividendsCalendarItem struct {
	// Ticker symbol of instrument
	Symbol *string `json:"symbol,omitempty"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode *string `json:"mic_code,omitempty"`
	// Exchange where instrument is traded
	Exchange *string `json:"exchange,omitempty"`
	// Stands for the ex date
	ExDate string `json:"ex_date"`
	// Dividend payment amount
	Amount float64 `json:"amount"`
}

type _DividendsCalendarItem DividendsCalendarItem

// NewDividendsCalendarItem instantiates a new DividendsCalendarItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDividendsCalendarItem(exDate string, amount float64) *DividendsCalendarItem {
	this := DividendsCalendarItem{}
	this.ExDate = exDate
	this.Amount = amount
	return &this
}

// NewDividendsCalendarItemWithDefaults instantiates a new DividendsCalendarItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDividendsCalendarItemWithDefaults() *DividendsCalendarItem {
	this := DividendsCalendarItem{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *DividendsCalendarItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DividendsCalendarItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *DividendsCalendarItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *DividendsCalendarItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *DividendsCalendarItem) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DividendsCalendarItem) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *DividendsCalendarItem) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *DividendsCalendarItem) SetMicCode(v string) {
	o.MicCode = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *DividendsCalendarItem) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DividendsCalendarItem) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *DividendsCalendarItem) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *DividendsCalendarItem) SetExchange(v string) {
	o.Exchange = &v
}

// GetExDate returns the ExDate field value
func (o *DividendsCalendarItem) GetExDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExDate
}

// GetExDateOk returns a tuple with the ExDate field value
// and a boolean to check if the value has been set.
func (o *DividendsCalendarItem) GetExDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExDate, true
}

// SetExDate sets field value
func (o *DividendsCalendarItem) SetExDate(v string) {
	o.ExDate = v
}

// GetAmount returns the Amount field value
func (o *DividendsCalendarItem) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *DividendsCalendarItem) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *DividendsCalendarItem) SetAmount(v float64) {
	o.Amount = v
}

func (o DividendsCalendarItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DividendsCalendarItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.MicCode) {
		toSerialize["mic_code"] = o.MicCode
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	toSerialize["ex_date"] = o.ExDate
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *DividendsCalendarItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ex_date",
		"amount",
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

	varDividendsCalendarItem := _DividendsCalendarItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDividendsCalendarItem)

	if err != nil {
		return err
	}

	*o = DividendsCalendarItem(varDividendsCalendarItem)

	return err
}

type NullableDividendsCalendarItem struct {
	value *DividendsCalendarItem
	isSet bool
}

func (v NullableDividendsCalendarItem) Get() *DividendsCalendarItem {
	return v.value
}

func (v *NullableDividendsCalendarItem) Set(val *DividendsCalendarItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDividendsCalendarItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDividendsCalendarItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDividendsCalendarItem(val *DividendsCalendarItem) *NullableDividendsCalendarItem {
	return &NullableDividendsCalendarItem{value: val, isSet: true}
}

func (v NullableDividendsCalendarItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDividendsCalendarItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
