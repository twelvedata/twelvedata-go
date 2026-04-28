/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the ResponseMutualFundWorldSustainability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMutualFundWorldSustainability{}

// ResponseMutualFundWorldSustainability Sustainability score and ESG (Environmental, Social, Governance) metrics for the fund
type ResponseMutualFundWorldSustainability struct {
	// Sustainability score: asset-weighted average of normalized company-level ESG Scores for the covered holdings in the portfolio from `0` to `100`
	Score               *int64                                                                     `json:"score,omitempty"`
	CorporateEsgPillars *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars `json:"corporate_esg_pillars,omitempty"`
	// Indication that the fund discloses in their prospectus that they employ socially responsible or ESG principles in their investment selection processes
	SustainableInvestment *bool `json:"sustainable_investment,omitempty"`
	// Percentage of AUM used to calculate sustainability score
	CorporateAum *float64 `json:"corporate_aum,omitempty"`
}

// NewResponseMutualFundWorldSustainability instantiates a new ResponseMutualFundWorldSustainability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMutualFundWorldSustainability() *ResponseMutualFundWorldSustainability {
	this := ResponseMutualFundWorldSustainability{}
	return &this
}

// NewResponseMutualFundWorldSustainabilityWithDefaults instantiates a new ResponseMutualFundWorldSustainability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMutualFundWorldSustainabilityWithDefaults() *ResponseMutualFundWorldSustainability {
	this := ResponseMutualFundWorldSustainability{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSustainability) GetScore() int64 {
	if o == nil || IsNil(o.Score) {
		var ret int64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSustainability) GetScoreOk() (*int64, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSustainability) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int64 and assigns it to the Score field.
func (o *ResponseMutualFundWorldSustainability) SetScore(v int64) {
	o.Score = &v
}

// GetCorporateEsgPillars returns the CorporateEsgPillars field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSustainability) GetCorporateEsgPillars() GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars {
	if o == nil || IsNil(o.CorporateEsgPillars) {
		var ret GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars
		return ret
	}
	return *o.CorporateEsgPillars
}

// GetCorporateEsgPillarsOk returns a tuple with the CorporateEsgPillars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSustainability) GetCorporateEsgPillarsOk() (*GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars, bool) {
	if o == nil || IsNil(o.CorporateEsgPillars) {
		return nil, false
	}
	return o.CorporateEsgPillars, true
}

// HasCorporateEsgPillars returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSustainability) HasCorporateEsgPillars() bool {
	if o != nil && !IsNil(o.CorporateEsgPillars) {
		return true
	}

	return false
}

// SetCorporateEsgPillars gets a reference to the given GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars and assigns it to the CorporateEsgPillars field.
func (o *ResponseMutualFundWorldSustainability) SetCorporateEsgPillars(v GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) {
	o.CorporateEsgPillars = &v
}

// GetSustainableInvestment returns the SustainableInvestment field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSustainability) GetSustainableInvestment() bool {
	if o == nil || IsNil(o.SustainableInvestment) {
		var ret bool
		return ret
	}
	return *o.SustainableInvestment
}

// GetSustainableInvestmentOk returns a tuple with the SustainableInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSustainability) GetSustainableInvestmentOk() (*bool, bool) {
	if o == nil || IsNil(o.SustainableInvestment) {
		return nil, false
	}
	return o.SustainableInvestment, true
}

// HasSustainableInvestment returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSustainability) HasSustainableInvestment() bool {
	if o != nil && !IsNil(o.SustainableInvestment) {
		return true
	}

	return false
}

// SetSustainableInvestment gets a reference to the given bool and assigns it to the SustainableInvestment field.
func (o *ResponseMutualFundWorldSustainability) SetSustainableInvestment(v bool) {
	o.SustainableInvestment = &v
}

// GetCorporateAum returns the CorporateAum field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSustainability) GetCorporateAum() float64 {
	if o == nil || IsNil(o.CorporateAum) {
		var ret float64
		return ret
	}
	return *o.CorporateAum
}

// GetCorporateAumOk returns a tuple with the CorporateAum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSustainability) GetCorporateAumOk() (*float64, bool) {
	if o == nil || IsNil(o.CorporateAum) {
		return nil, false
	}
	return o.CorporateAum, true
}

// HasCorporateAum returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSustainability) HasCorporateAum() bool {
	if o != nil && !IsNil(o.CorporateAum) {
		return true
	}

	return false
}

// SetCorporateAum gets a reference to the given float64 and assigns it to the CorporateAum field.
func (o *ResponseMutualFundWorldSustainability) SetCorporateAum(v float64) {
	o.CorporateAum = &v
}

func (o ResponseMutualFundWorldSustainability) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMutualFundWorldSustainability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.CorporateEsgPillars) {
		toSerialize["corporate_esg_pillars"] = o.CorporateEsgPillars
	}
	if !IsNil(o.SustainableInvestment) {
		toSerialize["sustainable_investment"] = o.SustainableInvestment
	}
	if !IsNil(o.CorporateAum) {
		toSerialize["corporate_aum"] = o.CorporateAum
	}
	return toSerialize, nil
}

type NullableResponseMutualFundWorldSustainability struct {
	value *ResponseMutualFundWorldSustainability
	isSet bool
}

func (v NullableResponseMutualFundWorldSustainability) Get() *ResponseMutualFundWorldSustainability {
	return v.value
}

func (v *NullableResponseMutualFundWorldSustainability) Set(val *ResponseMutualFundWorldSustainability) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMutualFundWorldSustainability) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMutualFundWorldSustainability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMutualFundWorldSustainability(val *ResponseMutualFundWorldSustainability) *NullableResponseMutualFundWorldSustainability {
	return &NullableResponseMutualFundWorldSustainability{value: val, isSet: true}
}

func (v NullableResponseMutualFundWorldSustainability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMutualFundWorldSustainability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
