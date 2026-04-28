/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the InlineObject4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject4{}

// InlineObject4 struct for InlineObject4
type InlineObject4 struct {
	CorporateAum          *float64           `json:"corporate_aum,omitempty"`
	CorporateEsgPillars   map[string]float64 `json:"corporate_esg_pillars,omitempty"`
	Score                 *int64             `json:"score,omitempty"`
	SustainableInvestment *bool              `json:"sustainable_investment,omitempty"`
}

// NewInlineObject4 instantiates a new InlineObject4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject4() *InlineObject4 {
	this := InlineObject4{}
	return &this
}

// NewInlineObject4WithDefaults instantiates a new InlineObject4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject4WithDefaults() *InlineObject4 {
	this := InlineObject4{}
	return &this
}

// GetCorporateAum returns the CorporateAum field value if set, zero value otherwise.
func (o *InlineObject4) GetCorporateAum() float64 {
	if o == nil || IsNil(o.CorporateAum) {
		var ret float64
		return ret
	}
	return *o.CorporateAum
}

// GetCorporateAumOk returns a tuple with the CorporateAum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetCorporateAumOk() (*float64, bool) {
	if o == nil || IsNil(o.CorporateAum) {
		return nil, false
	}
	return o.CorporateAum, true
}

// HasCorporateAum returns a boolean if a field has been set.
func (o *InlineObject4) HasCorporateAum() bool {
	if o != nil && !IsNil(o.CorporateAum) {
		return true
	}

	return false
}

// SetCorporateAum gets a reference to the given float64 and assigns it to the CorporateAum field.
func (o *InlineObject4) SetCorporateAum(v float64) {
	o.CorporateAum = &v
}

// GetCorporateEsgPillars returns the CorporateEsgPillars field value if set, zero value otherwise.
func (o *InlineObject4) GetCorporateEsgPillars() map[string]float64 {
	if o == nil || IsNil(o.CorporateEsgPillars) {
		var ret map[string]float64
		return ret
	}
	return o.CorporateEsgPillars
}

// GetCorporateEsgPillarsOk returns a tuple with the CorporateEsgPillars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetCorporateEsgPillarsOk() (map[string]float64, bool) {
	if o == nil || IsNil(o.CorporateEsgPillars) {
		return map[string]float64{}, false
	}
	return o.CorporateEsgPillars, true
}

// HasCorporateEsgPillars returns a boolean if a field has been set.
func (o *InlineObject4) HasCorporateEsgPillars() bool {
	if o != nil && !IsNil(o.CorporateEsgPillars) {
		return true
	}

	return false
}

// SetCorporateEsgPillars gets a reference to the given map[string]float64 and assigns it to the CorporateEsgPillars field.
func (o *InlineObject4) SetCorporateEsgPillars(v map[string]float64) {
	o.CorporateEsgPillars = v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *InlineObject4) GetScore() int64 {
	if o == nil || IsNil(o.Score) {
		var ret int64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetScoreOk() (*int64, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *InlineObject4) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int64 and assigns it to the Score field.
func (o *InlineObject4) SetScore(v int64) {
	o.Score = &v
}

// GetSustainableInvestment returns the SustainableInvestment field value if set, zero value otherwise.
func (o *InlineObject4) GetSustainableInvestment() bool {
	if o == nil || IsNil(o.SustainableInvestment) {
		var ret bool
		return ret
	}
	return *o.SustainableInvestment
}

// GetSustainableInvestmentOk returns a tuple with the SustainableInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetSustainableInvestmentOk() (*bool, bool) {
	if o == nil || IsNil(o.SustainableInvestment) {
		return nil, false
	}
	return o.SustainableInvestment, true
}

// HasSustainableInvestment returns a boolean if a field has been set.
func (o *InlineObject4) HasSustainableInvestment() bool {
	if o != nil && !IsNil(o.SustainableInvestment) {
		return true
	}

	return false
}

// SetSustainableInvestment gets a reference to the given bool and assigns it to the SustainableInvestment field.
func (o *InlineObject4) SetSustainableInvestment(v bool) {
	o.SustainableInvestment = &v
}

func (o InlineObject4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorporateAum) {
		toSerialize["corporate_aum"] = o.CorporateAum
	}
	if !IsNil(o.CorporateEsgPillars) {
		toSerialize["corporate_esg_pillars"] = o.CorporateEsgPillars
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.SustainableInvestment) {
		toSerialize["sustainable_investment"] = o.SustainableInvestment
	}
	return toSerialize, nil
}

type NullableInlineObject4 struct {
	value *InlineObject4
	isSet bool
}

func (v NullableInlineObject4) Get() *InlineObject4 {
	return v.value
}

func (v *NullableInlineObject4) Set(val *InlineObject4) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject4) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject4(val *InlineObject4) *NullableInlineObject4 {
	return &NullableInlineObject4{value: val, isSet: true}
}

func (v NullableInlineObject4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
