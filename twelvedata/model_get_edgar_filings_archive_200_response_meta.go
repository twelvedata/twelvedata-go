// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetEdgarFilingsArchive200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEdgarFilingsArchive200ResponseMeta{}

// GetEdgarFilingsArchive200ResponseMeta Meta information about the company and listing
type GetEdgarFilingsArchive200ResponseMeta struct {
	// Ticker of the company
	Symbol string `json:"symbol"`
	// Exchange name where the company is listed
	Exchange string `json:"exchange"`
	// Market Identifier Code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Issue type of the stock
	Type string `json:"type"`
}

type _GetEdgarFilingsArchive200ResponseMeta GetEdgarFilingsArchive200ResponseMeta

// NewGetEdgarFilingsArchive200ResponseMeta instantiates a new GetEdgarFilingsArchive200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEdgarFilingsArchive200ResponseMeta(symbol string, exchange string, micCode string, type_ string) *GetEdgarFilingsArchive200ResponseMeta {
	this := GetEdgarFilingsArchive200ResponseMeta{}
	this.Symbol = symbol
	this.Exchange = exchange
	this.MicCode = micCode
	this.Type = type_
	return &this
}

// NewGetEdgarFilingsArchive200ResponseMetaWithDefaults instantiates a new GetEdgarFilingsArchive200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEdgarFilingsArchive200ResponseMetaWithDefaults() *GetEdgarFilingsArchive200ResponseMeta {
	this := GetEdgarFilingsArchive200ResponseMeta{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetEdgarFilingsArchive200ResponseMeta) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetEdgarFilingsArchive200ResponseMeta) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetEdgarFilingsArchive200ResponseMeta) SetSymbol(v string) {
	o.Symbol = v
}

// GetExchange returns the Exchange field value
func (o *GetEdgarFilingsArchive200ResponseMeta) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetEdgarFilingsArchive200ResponseMeta) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetEdgarFilingsArchive200ResponseMeta) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *GetEdgarFilingsArchive200ResponseMeta) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *GetEdgarFilingsArchive200ResponseMeta) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *GetEdgarFilingsArchive200ResponseMeta) SetMicCode(v string) {
	o.MicCode = v
}

// GetType returns the Type field value
func (o *GetEdgarFilingsArchive200ResponseMeta) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetEdgarFilingsArchive200ResponseMeta) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetEdgarFilingsArchive200ResponseMeta) SetType(v string) {
	o.Type = v
}

func (o GetEdgarFilingsArchive200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEdgarFilingsArchive200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *GetEdgarFilingsArchive200ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"exchange",
		"mic_code",
		"type",
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

	varGetEdgarFilingsArchive200ResponseMeta := _GetEdgarFilingsArchive200ResponseMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetEdgarFilingsArchive200ResponseMeta)

	if err != nil {
		return err
	}

	*o = GetEdgarFilingsArchive200ResponseMeta(varGetEdgarFilingsArchive200ResponseMeta)

	return err
}

type NullableGetEdgarFilingsArchive200ResponseMeta struct {
	value *GetEdgarFilingsArchive200ResponseMeta
	isSet bool
}

func (v NullableGetEdgarFilingsArchive200ResponseMeta) Get() *GetEdgarFilingsArchive200ResponseMeta {
	return v.value
}

func (v *NullableGetEdgarFilingsArchive200ResponseMeta) Set(val *GetEdgarFilingsArchive200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEdgarFilingsArchive200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEdgarFilingsArchive200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEdgarFilingsArchive200ResponseMeta(val *GetEdgarFilingsArchive200ResponseMeta) *NullableGetEdgarFilingsArchive200ResponseMeta {
	return &NullableGetEdgarFilingsArchive200ResponseMeta{value: val, isSet: true}
}

func (v NullableGetEdgarFilingsArchive200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEdgarFilingsArchive200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
