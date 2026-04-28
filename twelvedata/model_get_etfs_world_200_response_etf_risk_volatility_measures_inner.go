/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner{}

// GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner struct for GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner
type GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner struct {
	// Period of a measure
	Period *string `json:"period,omitempty"`
	// Alpha score of a fund
	Alpha *float64 `json:"alpha,omitempty"`
	// Average alpha score of a fund's category
	AlphaCategory *float64 `json:"alpha_category,omitempty"`
	// Beta score of a fund
	Beta *float64 `json:"beta,omitempty"`
	// Average beta score of a fund's category
	BetaCategory *float64 `json:"beta_category,omitempty"`
	// Mean annual return of a fund
	MeanAnnualReturn *float64 `json:"mean_annual_return,omitempty"`
	// Average mean annual return of a fund's category
	MeanAnnualReturnCategory *float64 `json:"mean_annual_return_category,omitempty"`
	// R-squared metric of a fund
	RSquared *float64 `json:"r_squared,omitempty"`
	// Average r-squared metric of a fund's category
	RSquaredCategory *float64 `json:"r_squared_category,omitempty"`
	// Standard deviation of a fund
	Std *float64 `json:"std,omitempty"`
	// Average standard deviation of a fund's category
	StdCategory *float64 `json:"std_category,omitempty"`
	// Sharpe ratio of a fund
	SharpeRatio *float64 `json:"sharpe_ratio,omitempty"`
	// Average sharpe ratio of a fund's category
	SharpeRatioCategory *float64 `json:"sharpe_ratio_category,omitempty"`
	// Treynor ratio of a fund
	TreynorRatio *float64 `json:"treynor_ratio,omitempty"`
	// Average treynor ratio of a fund's category
	TreynorRatioCategory *float64 `json:"treynor_ratio_category,omitempty"`
}

// NewGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner instantiates a new GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner() *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner {
	this := GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner{}
	return &this
}

// NewGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInnerWithDefaults instantiates a new GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInnerWithDefaults() *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner {
	this := GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetPeriod(v string) {
	o.Period = &v
}

// GetAlpha returns the Alpha field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetAlpha() float64 {
	if o == nil || IsNil(o.Alpha) {
		var ret float64
		return ret
	}
	return *o.Alpha
}

// GetAlphaOk returns a tuple with the Alpha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetAlphaOk() (*float64, bool) {
	if o == nil || IsNil(o.Alpha) {
		return nil, false
	}
	return o.Alpha, true
}

// HasAlpha returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasAlpha() bool {
	if o != nil && !IsNil(o.Alpha) {
		return true
	}

	return false
}

// SetAlpha gets a reference to the given float64 and assigns it to the Alpha field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetAlpha(v float64) {
	o.Alpha = &v
}

// GetAlphaCategory returns the AlphaCategory field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetAlphaCategory() float64 {
	if o == nil || IsNil(o.AlphaCategory) {
		var ret float64
		return ret
	}
	return *o.AlphaCategory
}

// GetAlphaCategoryOk returns a tuple with the AlphaCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetAlphaCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.AlphaCategory) {
		return nil, false
	}
	return o.AlphaCategory, true
}

// HasAlphaCategory returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasAlphaCategory() bool {
	if o != nil && !IsNil(o.AlphaCategory) {
		return true
	}

	return false
}

// SetAlphaCategory gets a reference to the given float64 and assigns it to the AlphaCategory field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetAlphaCategory(v float64) {
	o.AlphaCategory = &v
}

// GetBeta returns the Beta field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetBeta() float64 {
	if o == nil || IsNil(o.Beta) {
		var ret float64
		return ret
	}
	return *o.Beta
}

// GetBetaOk returns a tuple with the Beta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetBetaOk() (*float64, bool) {
	if o == nil || IsNil(o.Beta) {
		return nil, false
	}
	return o.Beta, true
}

// HasBeta returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasBeta() bool {
	if o != nil && !IsNil(o.Beta) {
		return true
	}

	return false
}

// SetBeta gets a reference to the given float64 and assigns it to the Beta field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetBeta(v float64) {
	o.Beta = &v
}

// GetBetaCategory returns the BetaCategory field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetBetaCategory() float64 {
	if o == nil || IsNil(o.BetaCategory) {
		var ret float64
		return ret
	}
	return *o.BetaCategory
}

// GetBetaCategoryOk returns a tuple with the BetaCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetBetaCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.BetaCategory) {
		return nil, false
	}
	return o.BetaCategory, true
}

// HasBetaCategory returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasBetaCategory() bool {
	if o != nil && !IsNil(o.BetaCategory) {
		return true
	}

	return false
}

// SetBetaCategory gets a reference to the given float64 and assigns it to the BetaCategory field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetBetaCategory(v float64) {
	o.BetaCategory = &v
}

// GetMeanAnnualReturn returns the MeanAnnualReturn field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetMeanAnnualReturn() float64 {
	if o == nil || IsNil(o.MeanAnnualReturn) {
		var ret float64
		return ret
	}
	return *o.MeanAnnualReturn
}

// GetMeanAnnualReturnOk returns a tuple with the MeanAnnualReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetMeanAnnualReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.MeanAnnualReturn) {
		return nil, false
	}
	return o.MeanAnnualReturn, true
}

// HasMeanAnnualReturn returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasMeanAnnualReturn() bool {
	if o != nil && !IsNil(o.MeanAnnualReturn) {
		return true
	}

	return false
}

// SetMeanAnnualReturn gets a reference to the given float64 and assigns it to the MeanAnnualReturn field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetMeanAnnualReturn(v float64) {
	o.MeanAnnualReturn = &v
}

// GetMeanAnnualReturnCategory returns the MeanAnnualReturnCategory field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetMeanAnnualReturnCategory() float64 {
	if o == nil || IsNil(o.MeanAnnualReturnCategory) {
		var ret float64
		return ret
	}
	return *o.MeanAnnualReturnCategory
}

// GetMeanAnnualReturnCategoryOk returns a tuple with the MeanAnnualReturnCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetMeanAnnualReturnCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.MeanAnnualReturnCategory) {
		return nil, false
	}
	return o.MeanAnnualReturnCategory, true
}

// HasMeanAnnualReturnCategory returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasMeanAnnualReturnCategory() bool {
	if o != nil && !IsNil(o.MeanAnnualReturnCategory) {
		return true
	}

	return false
}

// SetMeanAnnualReturnCategory gets a reference to the given float64 and assigns it to the MeanAnnualReturnCategory field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetMeanAnnualReturnCategory(v float64) {
	o.MeanAnnualReturnCategory = &v
}

// GetRSquared returns the RSquared field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetRSquared() float64 {
	if o == nil || IsNil(o.RSquared) {
		var ret float64
		return ret
	}
	return *o.RSquared
}

// GetRSquaredOk returns a tuple with the RSquared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetRSquaredOk() (*float64, bool) {
	if o == nil || IsNil(o.RSquared) {
		return nil, false
	}
	return o.RSquared, true
}

// HasRSquared returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasRSquared() bool {
	if o != nil && !IsNil(o.RSquared) {
		return true
	}

	return false
}

// SetRSquared gets a reference to the given float64 and assigns it to the RSquared field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetRSquared(v float64) {
	o.RSquared = &v
}

// GetRSquaredCategory returns the RSquaredCategory field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetRSquaredCategory() float64 {
	if o == nil || IsNil(o.RSquaredCategory) {
		var ret float64
		return ret
	}
	return *o.RSquaredCategory
}

// GetRSquaredCategoryOk returns a tuple with the RSquaredCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetRSquaredCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.RSquaredCategory) {
		return nil, false
	}
	return o.RSquaredCategory, true
}

// HasRSquaredCategory returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasRSquaredCategory() bool {
	if o != nil && !IsNil(o.RSquaredCategory) {
		return true
	}

	return false
}

// SetRSquaredCategory gets a reference to the given float64 and assigns it to the RSquaredCategory field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetRSquaredCategory(v float64) {
	o.RSquaredCategory = &v
}

// GetStd returns the Std field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetStd() float64 {
	if o == nil || IsNil(o.Std) {
		var ret float64
		return ret
	}
	return *o.Std
}

// GetStdOk returns a tuple with the Std field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetStdOk() (*float64, bool) {
	if o == nil || IsNil(o.Std) {
		return nil, false
	}
	return o.Std, true
}

// HasStd returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasStd() bool {
	if o != nil && !IsNil(o.Std) {
		return true
	}

	return false
}

// SetStd gets a reference to the given float64 and assigns it to the Std field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetStd(v float64) {
	o.Std = &v
}

// GetStdCategory returns the StdCategory field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetStdCategory() float64 {
	if o == nil || IsNil(o.StdCategory) {
		var ret float64
		return ret
	}
	return *o.StdCategory
}

// GetStdCategoryOk returns a tuple with the StdCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetStdCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.StdCategory) {
		return nil, false
	}
	return o.StdCategory, true
}

// HasStdCategory returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasStdCategory() bool {
	if o != nil && !IsNil(o.StdCategory) {
		return true
	}

	return false
}

// SetStdCategory gets a reference to the given float64 and assigns it to the StdCategory field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetStdCategory(v float64) {
	o.StdCategory = &v
}

// GetSharpeRatio returns the SharpeRatio field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetSharpeRatio() float64 {
	if o == nil || IsNil(o.SharpeRatio) {
		var ret float64
		return ret
	}
	return *o.SharpeRatio
}

// GetSharpeRatioOk returns a tuple with the SharpeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetSharpeRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.SharpeRatio) {
		return nil, false
	}
	return o.SharpeRatio, true
}

// HasSharpeRatio returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasSharpeRatio() bool {
	if o != nil && !IsNil(o.SharpeRatio) {
		return true
	}

	return false
}

// SetSharpeRatio gets a reference to the given float64 and assigns it to the SharpeRatio field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetSharpeRatio(v float64) {
	o.SharpeRatio = &v
}

// GetSharpeRatioCategory returns the SharpeRatioCategory field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetSharpeRatioCategory() float64 {
	if o == nil || IsNil(o.SharpeRatioCategory) {
		var ret float64
		return ret
	}
	return *o.SharpeRatioCategory
}

// GetSharpeRatioCategoryOk returns a tuple with the SharpeRatioCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetSharpeRatioCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.SharpeRatioCategory) {
		return nil, false
	}
	return o.SharpeRatioCategory, true
}

// HasSharpeRatioCategory returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasSharpeRatioCategory() bool {
	if o != nil && !IsNil(o.SharpeRatioCategory) {
		return true
	}

	return false
}

// SetSharpeRatioCategory gets a reference to the given float64 and assigns it to the SharpeRatioCategory field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetSharpeRatioCategory(v float64) {
	o.SharpeRatioCategory = &v
}

// GetTreynorRatio returns the TreynorRatio field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetTreynorRatio() float64 {
	if o == nil || IsNil(o.TreynorRatio) {
		var ret float64
		return ret
	}
	return *o.TreynorRatio
}

// GetTreynorRatioOk returns a tuple with the TreynorRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetTreynorRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.TreynorRatio) {
		return nil, false
	}
	return o.TreynorRatio, true
}

// HasTreynorRatio returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasTreynorRatio() bool {
	if o != nil && !IsNil(o.TreynorRatio) {
		return true
	}

	return false
}

// SetTreynorRatio gets a reference to the given float64 and assigns it to the TreynorRatio field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetTreynorRatio(v float64) {
	o.TreynorRatio = &v
}

// GetTreynorRatioCategory returns the TreynorRatioCategory field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetTreynorRatioCategory() float64 {
	if o == nil || IsNil(o.TreynorRatioCategory) {
		var ret float64
		return ret
	}
	return *o.TreynorRatioCategory
}

// GetTreynorRatioCategoryOk returns a tuple with the TreynorRatioCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) GetTreynorRatioCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.TreynorRatioCategory) {
		return nil, false
	}
	return o.TreynorRatioCategory, true
}

// HasTreynorRatioCategory returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) HasTreynorRatioCategory() bool {
	if o != nil && !IsNil(o.TreynorRatioCategory) {
		return true
	}

	return false
}

// SetTreynorRatioCategory gets a reference to the given float64 and assigns it to the TreynorRatioCategory field.
func (o *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) SetTreynorRatioCategory(v float64) {
	o.TreynorRatioCategory = &v
}

func (o GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Alpha) {
		toSerialize["alpha"] = o.Alpha
	}
	if !IsNil(o.AlphaCategory) {
		toSerialize["alpha_category"] = o.AlphaCategory
	}
	if !IsNil(o.Beta) {
		toSerialize["beta"] = o.Beta
	}
	if !IsNil(o.BetaCategory) {
		toSerialize["beta_category"] = o.BetaCategory
	}
	if !IsNil(o.MeanAnnualReturn) {
		toSerialize["mean_annual_return"] = o.MeanAnnualReturn
	}
	if !IsNil(o.MeanAnnualReturnCategory) {
		toSerialize["mean_annual_return_category"] = o.MeanAnnualReturnCategory
	}
	if !IsNil(o.RSquared) {
		toSerialize["r_squared"] = o.RSquared
	}
	if !IsNil(o.RSquaredCategory) {
		toSerialize["r_squared_category"] = o.RSquaredCategory
	}
	if !IsNil(o.Std) {
		toSerialize["std"] = o.Std
	}
	if !IsNil(o.StdCategory) {
		toSerialize["std_category"] = o.StdCategory
	}
	if !IsNil(o.SharpeRatio) {
		toSerialize["sharpe_ratio"] = o.SharpeRatio
	}
	if !IsNil(o.SharpeRatioCategory) {
		toSerialize["sharpe_ratio_category"] = o.SharpeRatioCategory
	}
	if !IsNil(o.TreynorRatio) {
		toSerialize["treynor_ratio"] = o.TreynorRatio
	}
	if !IsNil(o.TreynorRatioCategory) {
		toSerialize["treynor_ratio_category"] = o.TreynorRatioCategory
	}
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner struct {
	value *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) Get() *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) Set(val *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner(val *GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) *NullableGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner {
	return &NullableGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
