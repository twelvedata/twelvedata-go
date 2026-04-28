/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the InlineObject2Minimums type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject2Minimums{}

// InlineObject2Minimums struct for InlineObject2Minimums
type InlineObject2Minimums struct {
	AdditionalInvestment    *int64                 `json:"additional_investment,omitempty"`
	AdditionalIraInvestment map[string]interface{} `json:"additional_ira_investment,omitempty"`
	InitialInvestment       *int64                 `json:"initial_investment,omitempty"`
	InitialIraInvestment    map[string]interface{} `json:"initial_ira_investment,omitempty"`
}

// NewInlineObject2Minimums instantiates a new InlineObject2Minimums object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject2Minimums() *InlineObject2Minimums {
	this := InlineObject2Minimums{}
	return &this
}

// NewInlineObject2MinimumsWithDefaults instantiates a new InlineObject2Minimums object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject2MinimumsWithDefaults() *InlineObject2Minimums {
	this := InlineObject2Minimums{}
	return &this
}

// GetAdditionalInvestment returns the AdditionalInvestment field value if set, zero value otherwise.
func (o *InlineObject2Minimums) GetAdditionalInvestment() int64 {
	if o == nil || IsNil(o.AdditionalInvestment) {
		var ret int64
		return ret
	}
	return *o.AdditionalInvestment
}

// GetAdditionalInvestmentOk returns a tuple with the AdditionalInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2Minimums) GetAdditionalInvestmentOk() (*int64, bool) {
	if o == nil || IsNil(o.AdditionalInvestment) {
		return nil, false
	}
	return o.AdditionalInvestment, true
}

// HasAdditionalInvestment returns a boolean if a field has been set.
func (o *InlineObject2Minimums) HasAdditionalInvestment() bool {
	if o != nil && !IsNil(o.AdditionalInvestment) {
		return true
	}

	return false
}

// SetAdditionalInvestment gets a reference to the given int64 and assigns it to the AdditionalInvestment field.
func (o *InlineObject2Minimums) SetAdditionalInvestment(v int64) {
	o.AdditionalInvestment = &v
}

// GetAdditionalIraInvestment returns the AdditionalIraInvestment field value if set, zero value otherwise.
func (o *InlineObject2Minimums) GetAdditionalIraInvestment() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalIraInvestment) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalIraInvestment
}

// GetAdditionalIraInvestmentOk returns a tuple with the AdditionalIraInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2Minimums) GetAdditionalIraInvestmentOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalIraInvestment) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalIraInvestment, true
}

// HasAdditionalIraInvestment returns a boolean if a field has been set.
func (o *InlineObject2Minimums) HasAdditionalIraInvestment() bool {
	if o != nil && !IsNil(o.AdditionalIraInvestment) {
		return true
	}

	return false
}

// SetAdditionalIraInvestment gets a reference to the given map[string]interface{} and assigns it to the AdditionalIraInvestment field.
func (o *InlineObject2Minimums) SetAdditionalIraInvestment(v map[string]interface{}) {
	o.AdditionalIraInvestment = v
}

// GetInitialInvestment returns the InitialInvestment field value if set, zero value otherwise.
func (o *InlineObject2Minimums) GetInitialInvestment() int64 {
	if o == nil || IsNil(o.InitialInvestment) {
		var ret int64
		return ret
	}
	return *o.InitialInvestment
}

// GetInitialInvestmentOk returns a tuple with the InitialInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2Minimums) GetInitialInvestmentOk() (*int64, bool) {
	if o == nil || IsNil(o.InitialInvestment) {
		return nil, false
	}
	return o.InitialInvestment, true
}

// HasInitialInvestment returns a boolean if a field has been set.
func (o *InlineObject2Minimums) HasInitialInvestment() bool {
	if o != nil && !IsNil(o.InitialInvestment) {
		return true
	}

	return false
}

// SetInitialInvestment gets a reference to the given int64 and assigns it to the InitialInvestment field.
func (o *InlineObject2Minimums) SetInitialInvestment(v int64) {
	o.InitialInvestment = &v
}

// GetInitialIraInvestment returns the InitialIraInvestment field value if set, zero value otherwise.
func (o *InlineObject2Minimums) GetInitialIraInvestment() map[string]interface{} {
	if o == nil || IsNil(o.InitialIraInvestment) {
		var ret map[string]interface{}
		return ret
	}
	return o.InitialIraInvestment
}

// GetInitialIraInvestmentOk returns a tuple with the InitialIraInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2Minimums) GetInitialIraInvestmentOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.InitialIraInvestment) {
		return map[string]interface{}{}, false
	}
	return o.InitialIraInvestment, true
}

// HasInitialIraInvestment returns a boolean if a field has been set.
func (o *InlineObject2Minimums) HasInitialIraInvestment() bool {
	if o != nil && !IsNil(o.InitialIraInvestment) {
		return true
	}

	return false
}

// SetInitialIraInvestment gets a reference to the given map[string]interface{} and assigns it to the InitialIraInvestment field.
func (o *InlineObject2Minimums) SetInitialIraInvestment(v map[string]interface{}) {
	o.InitialIraInvestment = v
}

func (o InlineObject2Minimums) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject2Minimums) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalInvestment) {
		toSerialize["additional_investment"] = o.AdditionalInvestment
	}
	if !IsNil(o.AdditionalIraInvestment) {
		toSerialize["additional_ira_investment"] = o.AdditionalIraInvestment
	}
	if !IsNil(o.InitialInvestment) {
		toSerialize["initial_investment"] = o.InitialInvestment
	}
	if !IsNil(o.InitialIraInvestment) {
		toSerialize["initial_ira_investment"] = o.InitialIraInvestment
	}
	return toSerialize, nil
}

type NullableInlineObject2Minimums struct {
	value *InlineObject2Minimums
	isSet bool
}

func (v NullableInlineObject2Minimums) Get() *InlineObject2Minimums {
	return v.value
}

func (v *NullableInlineObject2Minimums) Set(val *InlineObject2Minimums) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject2Minimums) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject2Minimums) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject2Minimums(val *InlineObject2Minimums) *NullableInlineObject2Minimums {
	return &NullableInlineObject2Minimums{value: val, isSet: true}
}

func (v NullableInlineObject2Minimums) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject2Minimums) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
