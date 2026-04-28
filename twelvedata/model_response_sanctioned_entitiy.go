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

// checks if the ResponseSanctionedEntitiy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSanctionedEntitiy{}

// ResponseSanctionedEntitiy struct for ResponseSanctionedEntitiy
type ResponseSanctionedEntitiy struct {
	// The instrument symbol ticker
	Symbol string `json:"symbol"`
	// The instrument name
	Name string `json:"name"`
	// Market Identifier Code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Country name
	Country  string               `json:"country"`
	Sanction ResponseSanctionItem `json:"sanction"`
}

type _ResponseSanctionedEntitiy ResponseSanctionedEntitiy

// NewResponseSanctionedEntitiy instantiates a new ResponseSanctionedEntitiy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSanctionedEntitiy(symbol string, name string, micCode string, country string, sanction ResponseSanctionItem) *ResponseSanctionedEntitiy {
	this := ResponseSanctionedEntitiy{}
	this.Symbol = symbol
	this.Name = name
	this.MicCode = micCode
	this.Country = country
	this.Sanction = sanction
	return &this
}

// NewResponseSanctionedEntitiyWithDefaults instantiates a new ResponseSanctionedEntitiy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSanctionedEntitiyWithDefaults() *ResponseSanctionedEntitiy {
	this := ResponseSanctionedEntitiy{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ResponseSanctionedEntitiy) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ResponseSanctionedEntitiy) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ResponseSanctionedEntitiy) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *ResponseSanctionedEntitiy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseSanctionedEntitiy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseSanctionedEntitiy) SetName(v string) {
	o.Name = v
}

// GetMicCode returns the MicCode field value
func (o *ResponseSanctionedEntitiy) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *ResponseSanctionedEntitiy) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *ResponseSanctionedEntitiy) SetMicCode(v string) {
	o.MicCode = v
}

// GetCountry returns the Country field value
func (o *ResponseSanctionedEntitiy) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *ResponseSanctionedEntitiy) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *ResponseSanctionedEntitiy) SetCountry(v string) {
	o.Country = v
}

// GetSanction returns the Sanction field value
func (o *ResponseSanctionedEntitiy) GetSanction() ResponseSanctionItem {
	if o == nil {
		var ret ResponseSanctionItem
		return ret
	}

	return o.Sanction
}

// GetSanctionOk returns a tuple with the Sanction field value
// and a boolean to check if the value has been set.
func (o *ResponseSanctionedEntitiy) GetSanctionOk() (*ResponseSanctionItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sanction, true
}

// SetSanction sets field value
func (o *ResponseSanctionedEntitiy) SetSanction(v ResponseSanctionItem) {
	o.Sanction = v
}

func (o ResponseSanctionedEntitiy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSanctionedEntitiy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["mic_code"] = o.MicCode
	toSerialize["country"] = o.Country
	toSerialize["sanction"] = o.Sanction
	return toSerialize, nil
}

func (o *ResponseSanctionedEntitiy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"mic_code",
		"country",
		"sanction",
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

	varResponseSanctionedEntitiy := _ResponseSanctionedEntitiy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseSanctionedEntitiy)

	if err != nil {
		return err
	}

	*o = ResponseSanctionedEntitiy(varResponseSanctionedEntitiy)

	return err
}

type NullableResponseSanctionedEntitiy struct {
	value *ResponseSanctionedEntitiy
	isSet bool
}

func (v NullableResponseSanctionedEntitiy) Get() *ResponseSanctionedEntitiy {
	return v.value
}

func (v *NullableResponseSanctionedEntitiy) Set(val *ResponseSanctionedEntitiy) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSanctionedEntitiy) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSanctionedEntitiy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSanctionedEntitiy(val *ResponseSanctionedEntitiy) *NullableResponseSanctionedEntitiy {
	return &NullableResponseSanctionedEntitiy{value: val, isSet: true}
}

func (v NullableResponseSanctionedEntitiy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSanctionedEntitiy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
