/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfCompositionTopHoldingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfCompositionTopHoldingsInner{}

// GetETFsWorld200ResponseEtfCompositionTopHoldingsInner struct for GetETFsWorld200ResponseEtfCompositionTopHoldingsInner
type GetETFsWorld200ResponseEtfCompositionTopHoldingsInner struct {
	// Symbol ticker of a holding instrument
	Symbol *string `json:"symbol,omitempty"`
	// Name of a holding instrument
	Name *string `json:"name,omitempty"`
	// Exchange where instrument is traded
	Exchange *string `json:"exchange,omitempty"`
	// Market Identifier Code (MIC) under ISO 10383 standard
	MicCode *string `json:"mic_code,omitempty"`
	// Weight of a holding instrument in overall portfolio composition
	Weight *float64 `json:"weight,omitempty"`
}

// NewGetETFsWorld200ResponseEtfCompositionTopHoldingsInner instantiates a new GetETFsWorld200ResponseEtfCompositionTopHoldingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfCompositionTopHoldingsInner() *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner {
	this := GetETFsWorld200ResponseEtfCompositionTopHoldingsInner{}
	return &this
}

// NewGetETFsWorld200ResponseEtfCompositionTopHoldingsInnerWithDefaults instantiates a new GetETFsWorld200ResponseEtfCompositionTopHoldingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfCompositionTopHoldingsInnerWithDefaults() *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner {
	this := GetETFsWorld200ResponseEtfCompositionTopHoldingsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) SetName(v string) {
	o.Name = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) SetExchange(v string) {
	o.Exchange = &v
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) SetMicCode(v string) {
	o.MicCode = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) SetWeight(v float64) {
	o.Weight = &v
}

func (o GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.MicCode) {
		toSerialize["mic_code"] = o.MicCode
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtfCompositionTopHoldingsInner struct {
	value *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfCompositionTopHoldingsInner) Get() *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionTopHoldingsInner) Set(val *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfCompositionTopHoldingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionTopHoldingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfCompositionTopHoldingsInner(val *GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) *NullableGetETFsWorld200ResponseEtfCompositionTopHoldingsInner {
	return &NullableGetETFsWorld200ResponseEtfCompositionTopHoldingsInner{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfCompositionTopHoldingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionTopHoldingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
