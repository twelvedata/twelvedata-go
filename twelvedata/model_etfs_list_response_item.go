// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ETFsListResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ETFsListResponseItem{}

// ETFsListResponseItem struct for ETFsListResponseItem
type ETFsListResponseItem struct {
	// Instrument symbol (ticker)
	Symbol string `json:"symbol"`
	// Full name of the fund
	Name string `json:"name"`
	// Country of fund incorporation
	Country string `json:"country"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Investment company that manages the fund
	FundFamily string `json:"fund_family"`
	// Type of fund
	FundType string `json:"fund_type"`
}

type _ETFsListResponseItem ETFsListResponseItem

// NewETFsListResponseItem instantiates a new ETFsListResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewETFsListResponseItem(symbol string, name string, country string, micCode string, fundFamily string, fundType string) *ETFsListResponseItem {
	this := ETFsListResponseItem{}
	this.Symbol = symbol
	this.Name = name
	this.Country = country
	this.MicCode = micCode
	this.FundFamily = fundFamily
	this.FundType = fundType
	return &this
}

// NewETFsListResponseItemWithDefaults instantiates a new ETFsListResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewETFsListResponseItemWithDefaults() *ETFsListResponseItem {
	this := ETFsListResponseItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ETFsListResponseItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ETFsListResponseItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ETFsListResponseItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *ETFsListResponseItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ETFsListResponseItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ETFsListResponseItem) SetName(v string) {
	o.Name = v
}

// GetCountry returns the Country field value
func (o *ETFsListResponseItem) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *ETFsListResponseItem) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *ETFsListResponseItem) SetCountry(v string) {
	o.Country = v
}

// GetMicCode returns the MicCode field value
func (o *ETFsListResponseItem) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *ETFsListResponseItem) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *ETFsListResponseItem) SetMicCode(v string) {
	o.MicCode = v
}

// GetFundFamily returns the FundFamily field value
func (o *ETFsListResponseItem) GetFundFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundFamily
}

// GetFundFamilyOk returns a tuple with the FundFamily field value
// and a boolean to check if the value has been set.
func (o *ETFsListResponseItem) GetFundFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundFamily, true
}

// SetFundFamily sets field value
func (o *ETFsListResponseItem) SetFundFamily(v string) {
	o.FundFamily = v
}

// GetFundType returns the FundType field value
func (o *ETFsListResponseItem) GetFundType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundType
}

// GetFundTypeOk returns a tuple with the FundType field value
// and a boolean to check if the value has been set.
func (o *ETFsListResponseItem) GetFundTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundType, true
}

// SetFundType sets field value
func (o *ETFsListResponseItem) SetFundType(v string) {
	o.FundType = v
}

func (o ETFsListResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ETFsListResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["country"] = o.Country
	toSerialize["mic_code"] = o.MicCode
	toSerialize["fund_family"] = o.FundFamily
	toSerialize["fund_type"] = o.FundType
	return toSerialize, nil
}

func (o *ETFsListResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"country",
		"mic_code",
		"fund_family",
		"fund_type",
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

	varETFsListResponseItem := _ETFsListResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varETFsListResponseItem)

	if err != nil {
		return err
	}

	*o = ETFsListResponseItem(varETFsListResponseItem)

	return err
}

type NullableETFsListResponseItem struct {
	value *ETFsListResponseItem
	isSet bool
}

func (v NullableETFsListResponseItem) Get() *ETFsListResponseItem {
	return v.value
}

func (v *NullableETFsListResponseItem) Set(val *ETFsListResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableETFsListResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableETFsListResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableETFsListResponseItem(val *ETFsListResponseItem) *NullableETFsListResponseItem {
	return &NullableETFsListResponseItem{value: val, isSet: true}
}

func (v NullableETFsListResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableETFsListResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
