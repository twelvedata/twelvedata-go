// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the StockExchange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockExchange{}

// StockExchange Stock info
type StockExchange struct {
	Country *string `json:"country,omitempty"`
	Name    *string `json:"name,omitempty"`
}

// NewStockExchange instantiates a new StockExchange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockExchange() *StockExchange {
	this := StockExchange{}
	return &this
}

// NewStockExchangeWithDefaults instantiates a new StockExchange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockExchangeWithDefaults() *StockExchange {
	this := StockExchange{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *StockExchange) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockExchange) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *StockExchange) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *StockExchange) SetCountry(v string) {
	o.Country = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StockExchange) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockExchange) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StockExchange) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StockExchange) SetName(v string) {
	o.Name = &v
}

func (o StockExchange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockExchange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableStockExchange struct {
	value *StockExchange
	isSet bool
}

func (v NullableStockExchange) Get() *StockExchange {
	return v.value
}

func (v *NullableStockExchange) Set(val *StockExchange) {
	v.value = val
	v.isSet = true
}

func (v NullableStockExchange) IsSet() bool {
	return v.isSet
}

func (v *NullableStockExchange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockExchange(val *StockExchange) *NullableStockExchange {
	return &NullableStockExchange{value: val, isSet: true}
}

func (v NullableStockExchange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockExchange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
