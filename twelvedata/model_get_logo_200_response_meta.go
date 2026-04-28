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

// checks if the GetLogo200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLogo200ResponseMeta{}

// GetLogo200ResponseMeta Meta information about the instrument
type GetLogo200ResponseMeta struct {
	// The ticker symbol of an instrument
	Symbol string `json:"symbol"`
	// The exchange where the instrument is traded (for `crypto` only)
	Exchange *string `json:"exchange,omitempty"`
}

type _GetLogo200ResponseMeta GetLogo200ResponseMeta

// NewGetLogo200ResponseMeta instantiates a new GetLogo200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLogo200ResponseMeta(symbol string) *GetLogo200ResponseMeta {
	this := GetLogo200ResponseMeta{}
	this.Symbol = symbol
	return &this
}

// NewGetLogo200ResponseMetaWithDefaults instantiates a new GetLogo200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLogo200ResponseMetaWithDefaults() *GetLogo200ResponseMeta {
	this := GetLogo200ResponseMeta{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetLogo200ResponseMeta) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetLogo200ResponseMeta) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetLogo200ResponseMeta) SetSymbol(v string) {
	o.Symbol = v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *GetLogo200ResponseMeta) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogo200ResponseMeta) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *GetLogo200ResponseMeta) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *GetLogo200ResponseMeta) SetExchange(v string) {
	o.Exchange = &v
}

func (o GetLogo200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogo200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	return toSerialize, nil
}

func (o *GetLogo200ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
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

	varGetLogo200ResponseMeta := _GetLogo200ResponseMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetLogo200ResponseMeta)

	if err != nil {
		return err
	}

	*o = GetLogo200ResponseMeta(varGetLogo200ResponseMeta)

	return err
}

type NullableGetLogo200ResponseMeta struct {
	value *GetLogo200ResponseMeta
	isSet bool
}

func (v NullableGetLogo200ResponseMeta) Get() *GetLogo200ResponseMeta {
	return v.value
}

func (v *NullableGetLogo200ResponseMeta) Set(val *GetLogo200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogo200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogo200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogo200ResponseMeta(val *GetLogo200ResponseMeta) *NullableGetLogo200ResponseMeta {
	return &NullableGetLogo200ResponseMeta{value: val, isSet: true}
}

func (v NullableGetLogo200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogo200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
