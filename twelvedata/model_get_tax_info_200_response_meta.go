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

// checks if the GetTaxInfo200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTaxInfo200ResponseMeta{}

// GetTaxInfo200ResponseMeta Metadata about the requested instrument
type GetTaxInfo200ResponseMeta struct {
	// The ticker symbol of an instrument for which data was requested
	Symbol string `json:"symbol"`
	// The instrument name
	Name string `json:"name"`
	// The exchange name where the instrument is traded
	Exchange string `json:"exchange"`
	// The Market Identifier Code (MIC) of the exchange where the instrument is traded
	MicCode string `json:"mic_code"`
	// The instrument country name
	Country string `json:"country"`
}

type _GetTaxInfo200ResponseMeta GetTaxInfo200ResponseMeta

// NewGetTaxInfo200ResponseMeta instantiates a new GetTaxInfo200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTaxInfo200ResponseMeta(symbol string, name string, exchange string, micCode string, country string) *GetTaxInfo200ResponseMeta {
	this := GetTaxInfo200ResponseMeta{}
	this.Symbol = symbol
	this.Name = name
	this.Exchange = exchange
	this.MicCode = micCode
	this.Country = country
	return &this
}

// NewGetTaxInfo200ResponseMetaWithDefaults instantiates a new GetTaxInfo200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTaxInfo200ResponseMetaWithDefaults() *GetTaxInfo200ResponseMeta {
	this := GetTaxInfo200ResponseMeta{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetTaxInfo200ResponseMeta) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetTaxInfo200ResponseMeta) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetTaxInfo200ResponseMeta) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *GetTaxInfo200ResponseMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTaxInfo200ResponseMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTaxInfo200ResponseMeta) SetName(v string) {
	o.Name = v
}

// GetExchange returns the Exchange field value
func (o *GetTaxInfo200ResponseMeta) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetTaxInfo200ResponseMeta) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetTaxInfo200ResponseMeta) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *GetTaxInfo200ResponseMeta) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *GetTaxInfo200ResponseMeta) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *GetTaxInfo200ResponseMeta) SetMicCode(v string) {
	o.MicCode = v
}

// GetCountry returns the Country field value
func (o *GetTaxInfo200ResponseMeta) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *GetTaxInfo200ResponseMeta) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *GetTaxInfo200ResponseMeta) SetCountry(v string) {
	o.Country = v
}

func (o GetTaxInfo200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTaxInfo200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["country"] = o.Country
	return toSerialize, nil
}

func (o *GetTaxInfo200ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"exchange",
		"mic_code",
		"country",
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

	varGetTaxInfo200ResponseMeta := _GetTaxInfo200ResponseMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTaxInfo200ResponseMeta)

	if err != nil {
		return err
	}

	*o = GetTaxInfo200ResponseMeta(varGetTaxInfo200ResponseMeta)

	return err
}

type NullableGetTaxInfo200ResponseMeta struct {
	value *GetTaxInfo200ResponseMeta
	isSet bool
}

func (v NullableGetTaxInfo200ResponseMeta) Get() *GetTaxInfo200ResponseMeta {
	return v.value
}

func (v *NullableGetTaxInfo200ResponseMeta) Set(val *GetTaxInfo200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTaxInfo200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTaxInfo200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTaxInfo200ResponseMeta(val *GetTaxInfo200ResponseMeta) *NullableGetTaxInfo200ResponseMeta {
	return &NullableGetTaxInfo200ResponseMeta{value: val, isSet: true}
}

func (v NullableGetTaxInfo200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTaxInfo200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
