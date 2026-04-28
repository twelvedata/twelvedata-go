/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFund{}

// GetMutualFundsWorld200ResponseMutualFund Mutual fund information
type GetMutualFundsWorld200ResponseMutualFund struct {
	Summary        *GetMutualFundsWorld200ResponseMutualFundSummary        `json:"summary,omitempty"`
	Performance    *GetMutualFundsWorld200ResponseMutualFundPerformance    `json:"performance,omitempty"`
	Risk           *GetMutualFundsWorld200ResponseMutualFundRisk           `json:"risk,omitempty"`
	Ratings        *GetMutualFundsWorld200ResponseMutualFundRatings        `json:"ratings,omitempty"`
	Composition    *GetMutualFundsWorld200ResponseMutualFundComposition    `json:"composition,omitempty"`
	PurchaseInfo   *GetMutualFundsWorld200ResponseMutualFundPurchaseInfo   `json:"purchase_info,omitempty"`
	Sustainability *GetMutualFundsWorld200ResponseMutualFundSustainability `json:"sustainability,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFund instantiates a new GetMutualFundsWorld200ResponseMutualFund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFund() *GetMutualFundsWorld200ResponseMutualFund {
	this := GetMutualFundsWorld200ResponseMutualFund{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundWithDefaults() *GetMutualFundsWorld200ResponseMutualFund {
	this := GetMutualFundsWorld200ResponseMutualFund{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetSummary() GetMutualFundsWorld200ResponseMutualFundSummary {
	if o == nil || IsNil(o.Summary) {
		var ret GetMutualFundsWorld200ResponseMutualFundSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetSummaryOk() (*GetMutualFundsWorld200ResponseMutualFundSummary, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given GetMutualFundsWorld200ResponseMutualFundSummary and assigns it to the Summary field.
func (o *GetMutualFundsWorld200ResponseMutualFund) SetSummary(v GetMutualFundsWorld200ResponseMutualFundSummary) {
	o.Summary = &v
}

// GetPerformance returns the Performance field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetPerformance() GetMutualFundsWorld200ResponseMutualFundPerformance {
	if o == nil || IsNil(o.Performance) {
		var ret GetMutualFundsWorld200ResponseMutualFundPerformance
		return ret
	}
	return *o.Performance
}

// GetPerformanceOk returns a tuple with the Performance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetPerformanceOk() (*GetMutualFundsWorld200ResponseMutualFundPerformance, bool) {
	if o == nil || IsNil(o.Performance) {
		return nil, false
	}
	return o.Performance, true
}

// HasPerformance returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) HasPerformance() bool {
	if o != nil && !IsNil(o.Performance) {
		return true
	}

	return false
}

// SetPerformance gets a reference to the given GetMutualFundsWorld200ResponseMutualFundPerformance and assigns it to the Performance field.
func (o *GetMutualFundsWorld200ResponseMutualFund) SetPerformance(v GetMutualFundsWorld200ResponseMutualFundPerformance) {
	o.Performance = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetRisk() GetMutualFundsWorld200ResponseMutualFundRisk {
	if o == nil || IsNil(o.Risk) {
		var ret GetMutualFundsWorld200ResponseMutualFundRisk
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetRiskOk() (*GetMutualFundsWorld200ResponseMutualFundRisk, bool) {
	if o == nil || IsNil(o.Risk) {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) HasRisk() bool {
	if o != nil && !IsNil(o.Risk) {
		return true
	}

	return false
}

// SetRisk gets a reference to the given GetMutualFundsWorld200ResponseMutualFundRisk and assigns it to the Risk field.
func (o *GetMutualFundsWorld200ResponseMutualFund) SetRisk(v GetMutualFundsWorld200ResponseMutualFundRisk) {
	o.Risk = &v
}

// GetRatings returns the Ratings field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetRatings() GetMutualFundsWorld200ResponseMutualFundRatings {
	if o == nil || IsNil(o.Ratings) {
		var ret GetMutualFundsWorld200ResponseMutualFundRatings
		return ret
	}
	return *o.Ratings
}

// GetRatingsOk returns a tuple with the Ratings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetRatingsOk() (*GetMutualFundsWorld200ResponseMutualFundRatings, bool) {
	if o == nil || IsNil(o.Ratings) {
		return nil, false
	}
	return o.Ratings, true
}

// HasRatings returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) HasRatings() bool {
	if o != nil && !IsNil(o.Ratings) {
		return true
	}

	return false
}

// SetRatings gets a reference to the given GetMutualFundsWorld200ResponseMutualFundRatings and assigns it to the Ratings field.
func (o *GetMutualFundsWorld200ResponseMutualFund) SetRatings(v GetMutualFundsWorld200ResponseMutualFundRatings) {
	o.Ratings = &v
}

// GetComposition returns the Composition field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetComposition() GetMutualFundsWorld200ResponseMutualFundComposition {
	if o == nil || IsNil(o.Composition) {
		var ret GetMutualFundsWorld200ResponseMutualFundComposition
		return ret
	}
	return *o.Composition
}

// GetCompositionOk returns a tuple with the Composition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetCompositionOk() (*GetMutualFundsWorld200ResponseMutualFundComposition, bool) {
	if o == nil || IsNil(o.Composition) {
		return nil, false
	}
	return o.Composition, true
}

// HasComposition returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) HasComposition() bool {
	if o != nil && !IsNil(o.Composition) {
		return true
	}

	return false
}

// SetComposition gets a reference to the given GetMutualFundsWorld200ResponseMutualFundComposition and assigns it to the Composition field.
func (o *GetMutualFundsWorld200ResponseMutualFund) SetComposition(v GetMutualFundsWorld200ResponseMutualFundComposition) {
	o.Composition = &v
}

// GetPurchaseInfo returns the PurchaseInfo field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetPurchaseInfo() GetMutualFundsWorld200ResponseMutualFundPurchaseInfo {
	if o == nil || IsNil(o.PurchaseInfo) {
		var ret GetMutualFundsWorld200ResponseMutualFundPurchaseInfo
		return ret
	}
	return *o.PurchaseInfo
}

// GetPurchaseInfoOk returns a tuple with the PurchaseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetPurchaseInfoOk() (*GetMutualFundsWorld200ResponseMutualFundPurchaseInfo, bool) {
	if o == nil || IsNil(o.PurchaseInfo) {
		return nil, false
	}
	return o.PurchaseInfo, true
}

// HasPurchaseInfo returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) HasPurchaseInfo() bool {
	if o != nil && !IsNil(o.PurchaseInfo) {
		return true
	}

	return false
}

// SetPurchaseInfo gets a reference to the given GetMutualFundsWorld200ResponseMutualFundPurchaseInfo and assigns it to the PurchaseInfo field.
func (o *GetMutualFundsWorld200ResponseMutualFund) SetPurchaseInfo(v GetMutualFundsWorld200ResponseMutualFundPurchaseInfo) {
	o.PurchaseInfo = &v
}

// GetSustainability returns the Sustainability field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetSustainability() GetMutualFundsWorld200ResponseMutualFundSustainability {
	if o == nil || IsNil(o.Sustainability) {
		var ret GetMutualFundsWorld200ResponseMutualFundSustainability
		return ret
	}
	return *o.Sustainability
}

// GetSustainabilityOk returns a tuple with the Sustainability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) GetSustainabilityOk() (*GetMutualFundsWorld200ResponseMutualFundSustainability, bool) {
	if o == nil || IsNil(o.Sustainability) {
		return nil, false
	}
	return o.Sustainability, true
}

// HasSustainability returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFund) HasSustainability() bool {
	if o != nil && !IsNil(o.Sustainability) {
		return true
	}

	return false
}

// SetSustainability gets a reference to the given GetMutualFundsWorld200ResponseMutualFundSustainability and assigns it to the Sustainability field.
func (o *GetMutualFundsWorld200ResponseMutualFund) SetSustainability(v GetMutualFundsWorld200ResponseMutualFundSustainability) {
	o.Sustainability = &v
}

func (o GetMutualFundsWorld200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Performance) {
		toSerialize["performance"] = o.Performance
	}
	if !IsNil(o.Risk) {
		toSerialize["risk"] = o.Risk
	}
	if !IsNil(o.Ratings) {
		toSerialize["ratings"] = o.Ratings
	}
	if !IsNil(o.Composition) {
		toSerialize["composition"] = o.Composition
	}
	if !IsNil(o.PurchaseInfo) {
		toSerialize["purchase_info"] = o.PurchaseInfo
	}
	if !IsNil(o.Sustainability) {
		toSerialize["sustainability"] = o.Sustainability
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFund struct {
	value *GetMutualFundsWorld200ResponseMutualFund
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFund) Get() *GetMutualFundsWorld200ResponseMutualFund {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFund) Set(val *GetMutualFundsWorld200ResponseMutualFund) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFund) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFund(val *GetMutualFundsWorld200ResponseMutualFund) *NullableGetMutualFundsWorld200ResponseMutualFund {
	return &NullableGetMutualFundsWorld200ResponseMutualFund{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
