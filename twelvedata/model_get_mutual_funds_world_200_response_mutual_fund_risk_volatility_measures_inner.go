// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner{}

// GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner struct for GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner
type GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner struct {
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

// NewGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner instantiates a new GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner() *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner {
	this := GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInnerWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInnerWithDefaults() *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner {
	this := GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetPeriod(v string) {
	o.Period = &v
}

// GetAlpha returns the Alpha field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetAlpha() float64 {
	if o == nil || IsNil(o.Alpha) {
		var ret float64
		return ret
	}
	return *o.Alpha
}

// GetAlphaOk returns a tuple with the Alpha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetAlphaOk() (*float64, bool) {
	if o == nil || IsNil(o.Alpha) {
		return nil, false
	}
	return o.Alpha, true
}

// HasAlpha returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasAlpha() bool {
	if o != nil && !IsNil(o.Alpha) {
		return true
	}

	return false
}

// SetAlpha gets a reference to the given float64 and assigns it to the Alpha field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetAlpha(v float64) {
	o.Alpha = &v
}

// GetAlphaCategory returns the AlphaCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetAlphaCategory() float64 {
	if o == nil || IsNil(o.AlphaCategory) {
		var ret float64
		return ret
	}
	return *o.AlphaCategory
}

// GetAlphaCategoryOk returns a tuple with the AlphaCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetAlphaCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.AlphaCategory) {
		return nil, false
	}
	return o.AlphaCategory, true
}

// HasAlphaCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasAlphaCategory() bool {
	if o != nil && !IsNil(o.AlphaCategory) {
		return true
	}

	return false
}

// SetAlphaCategory gets a reference to the given float64 and assigns it to the AlphaCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetAlphaCategory(v float64) {
	o.AlphaCategory = &v
}

// GetBeta returns the Beta field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetBeta() float64 {
	if o == nil || IsNil(o.Beta) {
		var ret float64
		return ret
	}
	return *o.Beta
}

// GetBetaOk returns a tuple with the Beta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetBetaOk() (*float64, bool) {
	if o == nil || IsNil(o.Beta) {
		return nil, false
	}
	return o.Beta, true
}

// HasBeta returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasBeta() bool {
	if o != nil && !IsNil(o.Beta) {
		return true
	}

	return false
}

// SetBeta gets a reference to the given float64 and assigns it to the Beta field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetBeta(v float64) {
	o.Beta = &v
}

// GetBetaCategory returns the BetaCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetBetaCategory() float64 {
	if o == nil || IsNil(o.BetaCategory) {
		var ret float64
		return ret
	}
	return *o.BetaCategory
}

// GetBetaCategoryOk returns a tuple with the BetaCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetBetaCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.BetaCategory) {
		return nil, false
	}
	return o.BetaCategory, true
}

// HasBetaCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasBetaCategory() bool {
	if o != nil && !IsNil(o.BetaCategory) {
		return true
	}

	return false
}

// SetBetaCategory gets a reference to the given float64 and assigns it to the BetaCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetBetaCategory(v float64) {
	o.BetaCategory = &v
}

// GetMeanAnnualReturn returns the MeanAnnualReturn field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetMeanAnnualReturn() float64 {
	if o == nil || IsNil(o.MeanAnnualReturn) {
		var ret float64
		return ret
	}
	return *o.MeanAnnualReturn
}

// GetMeanAnnualReturnOk returns a tuple with the MeanAnnualReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetMeanAnnualReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.MeanAnnualReturn) {
		return nil, false
	}
	return o.MeanAnnualReturn, true
}

// HasMeanAnnualReturn returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasMeanAnnualReturn() bool {
	if o != nil && !IsNil(o.MeanAnnualReturn) {
		return true
	}

	return false
}

// SetMeanAnnualReturn gets a reference to the given float64 and assigns it to the MeanAnnualReturn field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetMeanAnnualReturn(v float64) {
	o.MeanAnnualReturn = &v
}

// GetMeanAnnualReturnCategory returns the MeanAnnualReturnCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetMeanAnnualReturnCategory() float64 {
	if o == nil || IsNil(o.MeanAnnualReturnCategory) {
		var ret float64
		return ret
	}
	return *o.MeanAnnualReturnCategory
}

// GetMeanAnnualReturnCategoryOk returns a tuple with the MeanAnnualReturnCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetMeanAnnualReturnCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.MeanAnnualReturnCategory) {
		return nil, false
	}
	return o.MeanAnnualReturnCategory, true
}

// HasMeanAnnualReturnCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasMeanAnnualReturnCategory() bool {
	if o != nil && !IsNil(o.MeanAnnualReturnCategory) {
		return true
	}

	return false
}

// SetMeanAnnualReturnCategory gets a reference to the given float64 and assigns it to the MeanAnnualReturnCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetMeanAnnualReturnCategory(v float64) {
	o.MeanAnnualReturnCategory = &v
}

// GetRSquared returns the RSquared field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetRSquared() float64 {
	if o == nil || IsNil(o.RSquared) {
		var ret float64
		return ret
	}
	return *o.RSquared
}

// GetRSquaredOk returns a tuple with the RSquared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetRSquaredOk() (*float64, bool) {
	if o == nil || IsNil(o.RSquared) {
		return nil, false
	}
	return o.RSquared, true
}

// HasRSquared returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasRSquared() bool {
	if o != nil && !IsNil(o.RSquared) {
		return true
	}

	return false
}

// SetRSquared gets a reference to the given float64 and assigns it to the RSquared field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetRSquared(v float64) {
	o.RSquared = &v
}

// GetRSquaredCategory returns the RSquaredCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetRSquaredCategory() float64 {
	if o == nil || IsNil(o.RSquaredCategory) {
		var ret float64
		return ret
	}
	return *o.RSquaredCategory
}

// GetRSquaredCategoryOk returns a tuple with the RSquaredCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetRSquaredCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.RSquaredCategory) {
		return nil, false
	}
	return o.RSquaredCategory, true
}

// HasRSquaredCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasRSquaredCategory() bool {
	if o != nil && !IsNil(o.RSquaredCategory) {
		return true
	}

	return false
}

// SetRSquaredCategory gets a reference to the given float64 and assigns it to the RSquaredCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetRSquaredCategory(v float64) {
	o.RSquaredCategory = &v
}

// GetStd returns the Std field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetStd() float64 {
	if o == nil || IsNil(o.Std) {
		var ret float64
		return ret
	}
	return *o.Std
}

// GetStdOk returns a tuple with the Std field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetStdOk() (*float64, bool) {
	if o == nil || IsNil(o.Std) {
		return nil, false
	}
	return o.Std, true
}

// HasStd returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasStd() bool {
	if o != nil && !IsNil(o.Std) {
		return true
	}

	return false
}

// SetStd gets a reference to the given float64 and assigns it to the Std field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetStd(v float64) {
	o.Std = &v
}

// GetStdCategory returns the StdCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetStdCategory() float64 {
	if o == nil || IsNil(o.StdCategory) {
		var ret float64
		return ret
	}
	return *o.StdCategory
}

// GetStdCategoryOk returns a tuple with the StdCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetStdCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.StdCategory) {
		return nil, false
	}
	return o.StdCategory, true
}

// HasStdCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasStdCategory() bool {
	if o != nil && !IsNil(o.StdCategory) {
		return true
	}

	return false
}

// SetStdCategory gets a reference to the given float64 and assigns it to the StdCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetStdCategory(v float64) {
	o.StdCategory = &v
}

// GetSharpeRatio returns the SharpeRatio field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetSharpeRatio() float64 {
	if o == nil || IsNil(o.SharpeRatio) {
		var ret float64
		return ret
	}
	return *o.SharpeRatio
}

// GetSharpeRatioOk returns a tuple with the SharpeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetSharpeRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.SharpeRatio) {
		return nil, false
	}
	return o.SharpeRatio, true
}

// HasSharpeRatio returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasSharpeRatio() bool {
	if o != nil && !IsNil(o.SharpeRatio) {
		return true
	}

	return false
}

// SetSharpeRatio gets a reference to the given float64 and assigns it to the SharpeRatio field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetSharpeRatio(v float64) {
	o.SharpeRatio = &v
}

// GetSharpeRatioCategory returns the SharpeRatioCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetSharpeRatioCategory() float64 {
	if o == nil || IsNil(o.SharpeRatioCategory) {
		var ret float64
		return ret
	}
	return *o.SharpeRatioCategory
}

// GetSharpeRatioCategoryOk returns a tuple with the SharpeRatioCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetSharpeRatioCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.SharpeRatioCategory) {
		return nil, false
	}
	return o.SharpeRatioCategory, true
}

// HasSharpeRatioCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasSharpeRatioCategory() bool {
	if o != nil && !IsNil(o.SharpeRatioCategory) {
		return true
	}

	return false
}

// SetSharpeRatioCategory gets a reference to the given float64 and assigns it to the SharpeRatioCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetSharpeRatioCategory(v float64) {
	o.SharpeRatioCategory = &v
}

// GetTreynorRatio returns the TreynorRatio field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetTreynorRatio() float64 {
	if o == nil || IsNil(o.TreynorRatio) {
		var ret float64
		return ret
	}
	return *o.TreynorRatio
}

// GetTreynorRatioOk returns a tuple with the TreynorRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetTreynorRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.TreynorRatio) {
		return nil, false
	}
	return o.TreynorRatio, true
}

// HasTreynorRatio returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasTreynorRatio() bool {
	if o != nil && !IsNil(o.TreynorRatio) {
		return true
	}

	return false
}

// SetTreynorRatio gets a reference to the given float64 and assigns it to the TreynorRatio field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetTreynorRatio(v float64) {
	o.TreynorRatio = &v
}

// GetTreynorRatioCategory returns the TreynorRatioCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetTreynorRatioCategory() float64 {
	if o == nil || IsNil(o.TreynorRatioCategory) {
		var ret float64
		return ret
	}
	return *o.TreynorRatioCategory
}

// GetTreynorRatioCategoryOk returns a tuple with the TreynorRatioCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetTreynorRatioCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.TreynorRatioCategory) {
		return nil, false
	}
	return o.TreynorRatioCategory, true
}

// HasTreynorRatioCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasTreynorRatioCategory() bool {
	if o != nil && !IsNil(o.TreynorRatioCategory) {
		return true
	}

	return false
}

// SetTreynorRatioCategory gets a reference to the given float64 and assigns it to the TreynorRatioCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetTreynorRatioCategory(v float64) {
	o.TreynorRatioCategory = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) ToMap() (map[string]interface{}, error) {
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

type NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner struct {
	value *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) Get() *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) Set(val *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner(val *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) *NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner {
	return &NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
