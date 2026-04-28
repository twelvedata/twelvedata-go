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

// checks if the CountryResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryResponseItem{}

// CountryResponseItem CountryResponseItem represents details of a country
type CountryResponseItem struct {
	// Two-letter country code defined in ISO 3166
	Iso2 string `json:"iso2"`
	// Three-letter country code defined in ISO 3166
	Iso3 string `json:"iso3"`
	// Numeric country code defined in ISO 3166
	Numeric string `json:"numeric"`
	// The full name of country
	Name string `json:"name"`
	// Official name of country
	OfficialName string `json:"official_name"`
	// Capital of country
	Capital string `json:"capital"`
	// Currency of country
	Currency string `json:"currency"`
}

type _CountryResponseItem CountryResponseItem

// NewCountryResponseItem instantiates a new CountryResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryResponseItem(iso2 string, iso3 string, numeric string, name string, officialName string, capital string, currency string) *CountryResponseItem {
	this := CountryResponseItem{}
	this.Iso2 = iso2
	this.Iso3 = iso3
	this.Numeric = numeric
	this.Name = name
	this.OfficialName = officialName
	this.Capital = capital
	this.Currency = currency
	return &this
}

// NewCountryResponseItemWithDefaults instantiates a new CountryResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryResponseItemWithDefaults() *CountryResponseItem {
	this := CountryResponseItem{}
	return &this
}

// GetIso2 returns the Iso2 field value
func (o *CountryResponseItem) GetIso2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iso2
}

// GetIso2Ok returns a tuple with the Iso2 field value
// and a boolean to check if the value has been set.
func (o *CountryResponseItem) GetIso2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iso2, true
}

// SetIso2 sets field value
func (o *CountryResponseItem) SetIso2(v string) {
	o.Iso2 = v
}

// GetIso3 returns the Iso3 field value
func (o *CountryResponseItem) GetIso3() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iso3
}

// GetIso3Ok returns a tuple with the Iso3 field value
// and a boolean to check if the value has been set.
func (o *CountryResponseItem) GetIso3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iso3, true
}

// SetIso3 sets field value
func (o *CountryResponseItem) SetIso3(v string) {
	o.Iso3 = v
}

// GetNumeric returns the Numeric field value
func (o *CountryResponseItem) GetNumeric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Numeric
}

// GetNumericOk returns a tuple with the Numeric field value
// and a boolean to check if the value has been set.
func (o *CountryResponseItem) GetNumericOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Numeric, true
}

// SetNumeric sets field value
func (o *CountryResponseItem) SetNumeric(v string) {
	o.Numeric = v
}

// GetName returns the Name field value
func (o *CountryResponseItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CountryResponseItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CountryResponseItem) SetName(v string) {
	o.Name = v
}

// GetOfficialName returns the OfficialName field value
func (o *CountryResponseItem) GetOfficialName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfficialName
}

// GetOfficialNameOk returns a tuple with the OfficialName field value
// and a boolean to check if the value has been set.
func (o *CountryResponseItem) GetOfficialNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfficialName, true
}

// SetOfficialName sets field value
func (o *CountryResponseItem) SetOfficialName(v string) {
	o.OfficialName = v
}

// GetCapital returns the Capital field value
func (o *CountryResponseItem) GetCapital() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Capital
}

// GetCapitalOk returns a tuple with the Capital field value
// and a boolean to check if the value has been set.
func (o *CountryResponseItem) GetCapitalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capital, true
}

// SetCapital sets field value
func (o *CountryResponseItem) SetCapital(v string) {
	o.Capital = v
}

// GetCurrency returns the Currency field value
func (o *CountryResponseItem) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CountryResponseItem) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CountryResponseItem) SetCurrency(v string) {
	o.Currency = v
}

func (o CountryResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iso2"] = o.Iso2
	toSerialize["iso3"] = o.Iso3
	toSerialize["numeric"] = o.Numeric
	toSerialize["name"] = o.Name
	toSerialize["official_name"] = o.OfficialName
	toSerialize["capital"] = o.Capital
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

func (o *CountryResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iso2",
		"iso3",
		"numeric",
		"name",
		"official_name",
		"capital",
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

	varCountryResponseItem := _CountryResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCountryResponseItem)

	if err != nil {
		return err
	}

	*o = CountryResponseItem(varCountryResponseItem)

	return err
}

type NullableCountryResponseItem struct {
	value *CountryResponseItem
	isSet bool
}

func (v NullableCountryResponseItem) Get() *CountryResponseItem {
	return v.value
}

func (v *NullableCountryResponseItem) Set(val *CountryResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryResponseItem(val *CountryResponseItem) *NullableCountryResponseItem {
	return &NullableCountryResponseItem{value: val, isSet: true}
}

func (v NullableCountryResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
