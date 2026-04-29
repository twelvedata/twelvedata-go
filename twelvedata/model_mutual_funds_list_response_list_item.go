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

// checks if the MutualFundsListResponseListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MutualFundsListResponseListItem{}

// MutualFundsListResponseListItem struct for MutualFundsListResponseListItem
type MutualFundsListResponseListItem struct {
	// Fund symbol ticker
	Symbol string `json:"symbol"`
	// Fund name
	Name string `json:"name"`
	// Country of fund incorporation
	Country string `json:"country"`
	// Investment company that manages the fund
	FundFamily string `json:"fund_family"`
	// Type of fund
	FundType string `json:"fund_type"`
	// Performance rating from `0` to `5`
	PerformanceRating *int64 `json:"performance_rating,omitempty"`
	// Risk rating from `0` to `5`
	RiskRating *int64 `json:"risk_rating,omitempty"`
	// Currency code in which the fund is denominated
	Currency string `json:"currency"`
	// Exchange name where the fund is listed
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
}

type _MutualFundsListResponseListItem MutualFundsListResponseListItem

// NewMutualFundsListResponseListItem instantiates a new MutualFundsListResponseListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutualFundsListResponseListItem(symbol string, name string, country string, fundFamily string, fundType string, currency string, exchange string, micCode string) *MutualFundsListResponseListItem {
	this := MutualFundsListResponseListItem{}
	this.Symbol = symbol
	this.Name = name
	this.Country = country
	this.FundFamily = fundFamily
	this.FundType = fundType
	this.Currency = currency
	this.Exchange = exchange
	this.MicCode = micCode
	return &this
}

// NewMutualFundsListResponseListItemWithDefaults instantiates a new MutualFundsListResponseListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutualFundsListResponseListItemWithDefaults() *MutualFundsListResponseListItem {
	this := MutualFundsListResponseListItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *MutualFundsListResponseListItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *MutualFundsListResponseListItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *MutualFundsListResponseListItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MutualFundsListResponseListItem) SetName(v string) {
	o.Name = v
}

// GetCountry returns the Country field value
func (o *MutualFundsListResponseListItem) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *MutualFundsListResponseListItem) SetCountry(v string) {
	o.Country = v
}

// GetFundFamily returns the FundFamily field value
func (o *MutualFundsListResponseListItem) GetFundFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundFamily
}

// GetFundFamilyOk returns a tuple with the FundFamily field value
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetFundFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundFamily, true
}

// SetFundFamily sets field value
func (o *MutualFundsListResponseListItem) SetFundFamily(v string) {
	o.FundFamily = v
}

// GetFundType returns the FundType field value
func (o *MutualFundsListResponseListItem) GetFundType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundType
}

// GetFundTypeOk returns a tuple with the FundType field value
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetFundTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundType, true
}

// SetFundType sets field value
func (o *MutualFundsListResponseListItem) SetFundType(v string) {
	o.FundType = v
}

// GetPerformanceRating returns the PerformanceRating field value if set, zero value otherwise.
func (o *MutualFundsListResponseListItem) GetPerformanceRating() int64 {
	if o == nil || IsNil(o.PerformanceRating) {
		var ret int64
		return ret
	}
	return *o.PerformanceRating
}

// GetPerformanceRatingOk returns a tuple with the PerformanceRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetPerformanceRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.PerformanceRating) {
		return nil, false
	}
	return o.PerformanceRating, true
}

// HasPerformanceRating returns a boolean if a field has been set.
func (o *MutualFundsListResponseListItem) HasPerformanceRating() bool {
	if o != nil && !IsNil(o.PerformanceRating) {
		return true
	}

	return false
}

// SetPerformanceRating gets a reference to the given int64 and assigns it to the PerformanceRating field.
func (o *MutualFundsListResponseListItem) SetPerformanceRating(v int64) {
	o.PerformanceRating = &v
}

// GetRiskRating returns the RiskRating field value if set, zero value otherwise.
func (o *MutualFundsListResponseListItem) GetRiskRating() int64 {
	if o == nil || IsNil(o.RiskRating) {
		var ret int64
		return ret
	}
	return *o.RiskRating
}

// GetRiskRatingOk returns a tuple with the RiskRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetRiskRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.RiskRating) {
		return nil, false
	}
	return o.RiskRating, true
}

// HasRiskRating returns a boolean if a field has been set.
func (o *MutualFundsListResponseListItem) HasRiskRating() bool {
	if o != nil && !IsNil(o.RiskRating) {
		return true
	}

	return false
}

// SetRiskRating gets a reference to the given int64 and assigns it to the RiskRating field.
func (o *MutualFundsListResponseListItem) SetRiskRating(v int64) {
	o.RiskRating = &v
}

// GetCurrency returns the Currency field value
func (o *MutualFundsListResponseListItem) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *MutualFundsListResponseListItem) SetCurrency(v string) {
	o.Currency = v
}

// GetExchange returns the Exchange field value
func (o *MutualFundsListResponseListItem) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *MutualFundsListResponseListItem) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *MutualFundsListResponseListItem) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *MutualFundsListResponseListItem) SetMicCode(v string) {
	o.MicCode = v
}

func (o MutualFundsListResponseListItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutualFundsListResponseListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["country"] = o.Country
	toSerialize["fund_family"] = o.FundFamily
	toSerialize["fund_type"] = o.FundType
	if !IsNil(o.PerformanceRating) {
		toSerialize["performance_rating"] = o.PerformanceRating
	}
	if !IsNil(o.RiskRating) {
		toSerialize["risk_rating"] = o.RiskRating
	}
	toSerialize["currency"] = o.Currency
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	return toSerialize, nil
}

func (o *MutualFundsListResponseListItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"country",
		"fund_family",
		"fund_type",
		"currency",
		"exchange",
		"mic_code",
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

	varMutualFundsListResponseListItem := _MutualFundsListResponseListItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varMutualFundsListResponseListItem)

	if err != nil {
		return err
	}

	*o = MutualFundsListResponseListItem(varMutualFundsListResponseListItem)

	return err
}

type NullableMutualFundsListResponseListItem struct {
	value *MutualFundsListResponseListItem
	isSet bool
}

func (v NullableMutualFundsListResponseListItem) Get() *MutualFundsListResponseListItem {
	return v.value
}

func (v *NullableMutualFundsListResponseListItem) Set(val *MutualFundsListResponseListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMutualFundsListResponseListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMutualFundsListResponseListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutualFundsListResponseListItem(val *MutualFundsListResponseListItem) *NullableMutualFundsListResponseListItem {
	return &NullableMutualFundsListResponseListItem{value: val, isSet: true}
}

func (v NullableMutualFundsListResponseListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutualFundsListResponseListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
