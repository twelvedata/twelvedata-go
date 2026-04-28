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

// checks if the SplitsCalendarResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SplitsCalendarResponseItem{}

// SplitsCalendarResponseItem List of stock splits
type SplitsCalendarResponseItem struct {
	// Stands for the split date
	Date string `json:"date"`
	// Ticker of the company
	Symbol *string `json:"symbol,omitempty"`
	// Market Identifier Code (MIC) under ISO 10383 standard
	MicCode *string `json:"mic_code,omitempty"`
	// Exchange name where the company is listed
	Exchange *string `json:"exchange,omitempty"`
	// Specification of the split event
	Description string `json:"description"`
	// The ratio by which the number of a company's outstanding shares of stock are increased following a stock split. For example, a `4-for-1 split` results in four times as many outstanding shares, with each share selling at one forth of its pre-split price
	Ratio float64 `json:"ratio"`
	// From factor of the split
	FromFactor float64 `json:"from_factor"`
	// To factor of the split
	ToFactor float64 `json:"to_factor"`
}

type _SplitsCalendarResponseItem SplitsCalendarResponseItem

// NewSplitsCalendarResponseItem instantiates a new SplitsCalendarResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitsCalendarResponseItem(date string, description string, ratio float64, fromFactor float64, toFactor float64) *SplitsCalendarResponseItem {
	this := SplitsCalendarResponseItem{}
	this.Date = date
	this.Description = description
	this.Ratio = ratio
	this.FromFactor = fromFactor
	this.ToFactor = toFactor
	return &this
}

// NewSplitsCalendarResponseItemWithDefaults instantiates a new SplitsCalendarResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitsCalendarResponseItemWithDefaults() *SplitsCalendarResponseItem {
	this := SplitsCalendarResponseItem{}
	return &this
}

// GetDate returns the Date field value
func (o *SplitsCalendarResponseItem) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *SplitsCalendarResponseItem) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *SplitsCalendarResponseItem) SetDate(v string) {
	o.Date = v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SplitsCalendarResponseItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitsCalendarResponseItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SplitsCalendarResponseItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SplitsCalendarResponseItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *SplitsCalendarResponseItem) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitsCalendarResponseItem) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *SplitsCalendarResponseItem) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *SplitsCalendarResponseItem) SetMicCode(v string) {
	o.MicCode = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *SplitsCalendarResponseItem) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitsCalendarResponseItem) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *SplitsCalendarResponseItem) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *SplitsCalendarResponseItem) SetExchange(v string) {
	o.Exchange = &v
}

// GetDescription returns the Description field value
func (o *SplitsCalendarResponseItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SplitsCalendarResponseItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SplitsCalendarResponseItem) SetDescription(v string) {
	o.Description = v
}

// GetRatio returns the Ratio field value
func (o *SplitsCalendarResponseItem) GetRatio() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value
// and a boolean to check if the value has been set.
func (o *SplitsCalendarResponseItem) GetRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ratio, true
}

// SetRatio sets field value
func (o *SplitsCalendarResponseItem) SetRatio(v float64) {
	o.Ratio = v
}

// GetFromFactor returns the FromFactor field value
func (o *SplitsCalendarResponseItem) GetFromFactor() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.FromFactor
}

// GetFromFactorOk returns a tuple with the FromFactor field value
// and a boolean to check if the value has been set.
func (o *SplitsCalendarResponseItem) GetFromFactorOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromFactor, true
}

// SetFromFactor sets field value
func (o *SplitsCalendarResponseItem) SetFromFactor(v float64) {
	o.FromFactor = v
}

// GetToFactor returns the ToFactor field value
func (o *SplitsCalendarResponseItem) GetToFactor() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ToFactor
}

// GetToFactorOk returns a tuple with the ToFactor field value
// and a boolean to check if the value has been set.
func (o *SplitsCalendarResponseItem) GetToFactorOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToFactor, true
}

// SetToFactor sets field value
func (o *SplitsCalendarResponseItem) SetToFactor(v float64) {
	o.ToFactor = v
}

func (o SplitsCalendarResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitsCalendarResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.MicCode) {
		toSerialize["mic_code"] = o.MicCode
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	toSerialize["description"] = o.Description
	toSerialize["ratio"] = o.Ratio
	toSerialize["from_factor"] = o.FromFactor
	toSerialize["to_factor"] = o.ToFactor
	return toSerialize, nil
}

func (o *SplitsCalendarResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"description",
		"ratio",
		"from_factor",
		"to_factor",
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

	varSplitsCalendarResponseItem := _SplitsCalendarResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSplitsCalendarResponseItem)

	if err != nil {
		return err
	}

	*o = SplitsCalendarResponseItem(varSplitsCalendarResponseItem)

	return err
}

type NullableSplitsCalendarResponseItem struct {
	value *SplitsCalendarResponseItem
	isSet bool
}

func (v NullableSplitsCalendarResponseItem) Get() *SplitsCalendarResponseItem {
	return v.value
}

func (v *NullableSplitsCalendarResponseItem) Set(val *SplitsCalendarResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitsCalendarResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitsCalendarResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitsCalendarResponseItem(val *SplitsCalendarResponseItem) *NullableSplitsCalendarResponseItem {
	return &NullableSplitsCalendarResponseItem{value: val, isSet: true}
}

func (v NullableSplitsCalendarResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitsCalendarResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
