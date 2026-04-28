/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner{}

// GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner struct for GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner
type GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner struct {
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

// NewGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner instantiates a new GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner() *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner {
	this := GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInnerWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInnerWithDefaults() *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner {
	this := GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) SetName(v string) {
	o.Name = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) SetExchange(v string) {
	o.Exchange = &v
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) SetMicCode(v string) {
	o.MicCode = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) SetWeight(v float64) {
	o.Weight = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) ToMap() (map[string]interface{}, error) {
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

type NullableGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner struct {
	value *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) Get() *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) Set(val *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner(val *GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) *NullableGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner {
	return &NullableGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
