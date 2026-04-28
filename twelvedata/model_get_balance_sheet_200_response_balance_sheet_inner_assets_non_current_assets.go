/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets{}

// GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets Non-current assets section
type GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets struct {
	// Represents property owned
	Properties *float64 `json:"properties,omitempty"`
	// Represents land and improvements owned
	LandAndImprovements *float64 `json:"land_and_improvements,omitempty"`
	// Represents office equipment, furniture, and vehicles owned
	MachineryFurnitureEquipment *float64 `json:"machinery_furniture_equipment,omitempty"`
	// Represents the cost of construction work, which is not yet completed
	ConstructionInProgress *float64 `json:"construction_in_progress,omitempty"`
	// Represents operating and financial leases
	Leases *float64 `json:"leases,omitempty"`
	// Represents the cumulative depreciation of an asset that has been recorded
	AccumulatedDepreciation *float64 `json:"accumulated_depreciation,omitempty"`
	// Represents the value of a brand name, solid customer base, good customer relations, good employee relations, and proprietary technology
	Goodwill *float64 `json:"goodwill,omitempty"`
	// Represents real estate property purchased with the intention of earning a return on the investment
	InvestmentProperties *float64 `json:"investment_properties,omitempty"`
	// Represents liquid asset that gets its value from a contractual right or ownership claim
	FinancialAssets *float64 `json:"financial_assets,omitempty"`
	// Represents the patents, trademarks, and other intellectual properties
	IntangibleAssets *float64 `json:"intangible_assets,omitempty"`
	// Represents available for sale financial securities
	InvestmentsAndAdvances *float64 `json:"investments_and_advances,omitempty"`
	// Represents other long-term assets
	OtherNonCurrentAssets *float64 `json:"other_non_current_assets,omitempty"`
	// All long-term assets values in total
	TotalNonCurrentAssets *float64 `json:"total_non_current_assets,omitempty"`
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets() *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets {
	this := GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets{}
	return &this
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssetsWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssetsWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets {
	this := GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetProperties() float64 {
	if o == nil || IsNil(o.Properties) {
		var ret float64
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetPropertiesOk() (*float64, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given float64 and assigns it to the Properties field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetProperties(v float64) {
	o.Properties = &v
}

// GetLandAndImprovements returns the LandAndImprovements field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetLandAndImprovements() float64 {
	if o == nil || IsNil(o.LandAndImprovements) {
		var ret float64
		return ret
	}
	return *o.LandAndImprovements
}

// GetLandAndImprovementsOk returns a tuple with the LandAndImprovements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetLandAndImprovementsOk() (*float64, bool) {
	if o == nil || IsNil(o.LandAndImprovements) {
		return nil, false
	}
	return o.LandAndImprovements, true
}

// HasLandAndImprovements returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasLandAndImprovements() bool {
	if o != nil && !IsNil(o.LandAndImprovements) {
		return true
	}

	return false
}

// SetLandAndImprovements gets a reference to the given float64 and assigns it to the LandAndImprovements field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetLandAndImprovements(v float64) {
	o.LandAndImprovements = &v
}

// GetMachineryFurnitureEquipment returns the MachineryFurnitureEquipment field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetMachineryFurnitureEquipment() float64 {
	if o == nil || IsNil(o.MachineryFurnitureEquipment) {
		var ret float64
		return ret
	}
	return *o.MachineryFurnitureEquipment
}

// GetMachineryFurnitureEquipmentOk returns a tuple with the MachineryFurnitureEquipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetMachineryFurnitureEquipmentOk() (*float64, bool) {
	if o == nil || IsNil(o.MachineryFurnitureEquipment) {
		return nil, false
	}
	return o.MachineryFurnitureEquipment, true
}

// HasMachineryFurnitureEquipment returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasMachineryFurnitureEquipment() bool {
	if o != nil && !IsNil(o.MachineryFurnitureEquipment) {
		return true
	}

	return false
}

// SetMachineryFurnitureEquipment gets a reference to the given float64 and assigns it to the MachineryFurnitureEquipment field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetMachineryFurnitureEquipment(v float64) {
	o.MachineryFurnitureEquipment = &v
}

// GetConstructionInProgress returns the ConstructionInProgress field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetConstructionInProgress() float64 {
	if o == nil || IsNil(o.ConstructionInProgress) {
		var ret float64
		return ret
	}
	return *o.ConstructionInProgress
}

// GetConstructionInProgressOk returns a tuple with the ConstructionInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetConstructionInProgressOk() (*float64, bool) {
	if o == nil || IsNil(o.ConstructionInProgress) {
		return nil, false
	}
	return o.ConstructionInProgress, true
}

// HasConstructionInProgress returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasConstructionInProgress() bool {
	if o != nil && !IsNil(o.ConstructionInProgress) {
		return true
	}

	return false
}

// SetConstructionInProgress gets a reference to the given float64 and assigns it to the ConstructionInProgress field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetConstructionInProgress(v float64) {
	o.ConstructionInProgress = &v
}

// GetLeases returns the Leases field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetLeases() float64 {
	if o == nil || IsNil(o.Leases) {
		var ret float64
		return ret
	}
	return *o.Leases
}

// GetLeasesOk returns a tuple with the Leases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetLeasesOk() (*float64, bool) {
	if o == nil || IsNil(o.Leases) {
		return nil, false
	}
	return o.Leases, true
}

// HasLeases returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasLeases() bool {
	if o != nil && !IsNil(o.Leases) {
		return true
	}

	return false
}

// SetLeases gets a reference to the given float64 and assigns it to the Leases field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetLeases(v float64) {
	o.Leases = &v
}

// GetAccumulatedDepreciation returns the AccumulatedDepreciation field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetAccumulatedDepreciation() float64 {
	if o == nil || IsNil(o.AccumulatedDepreciation) {
		var ret float64
		return ret
	}
	return *o.AccumulatedDepreciation
}

// GetAccumulatedDepreciationOk returns a tuple with the AccumulatedDepreciation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetAccumulatedDepreciationOk() (*float64, bool) {
	if o == nil || IsNil(o.AccumulatedDepreciation) {
		return nil, false
	}
	return o.AccumulatedDepreciation, true
}

// HasAccumulatedDepreciation returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasAccumulatedDepreciation() bool {
	if o != nil && !IsNil(o.AccumulatedDepreciation) {
		return true
	}

	return false
}

// SetAccumulatedDepreciation gets a reference to the given float64 and assigns it to the AccumulatedDepreciation field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetAccumulatedDepreciation(v float64) {
	o.AccumulatedDepreciation = &v
}

// GetGoodwill returns the Goodwill field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetGoodwill() float64 {
	if o == nil || IsNil(o.Goodwill) {
		var ret float64
		return ret
	}
	return *o.Goodwill
}

// GetGoodwillOk returns a tuple with the Goodwill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetGoodwillOk() (*float64, bool) {
	if o == nil || IsNil(o.Goodwill) {
		return nil, false
	}
	return o.Goodwill, true
}

// HasGoodwill returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasGoodwill() bool {
	if o != nil && !IsNil(o.Goodwill) {
		return true
	}

	return false
}

// SetGoodwill gets a reference to the given float64 and assigns it to the Goodwill field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetGoodwill(v float64) {
	o.Goodwill = &v
}

// GetInvestmentProperties returns the InvestmentProperties field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetInvestmentProperties() float64 {
	if o == nil || IsNil(o.InvestmentProperties) {
		var ret float64
		return ret
	}
	return *o.InvestmentProperties
}

// GetInvestmentPropertiesOk returns a tuple with the InvestmentProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetInvestmentPropertiesOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentProperties) {
		return nil, false
	}
	return o.InvestmentProperties, true
}

// HasInvestmentProperties returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasInvestmentProperties() bool {
	if o != nil && !IsNil(o.InvestmentProperties) {
		return true
	}

	return false
}

// SetInvestmentProperties gets a reference to the given float64 and assigns it to the InvestmentProperties field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetInvestmentProperties(v float64) {
	o.InvestmentProperties = &v
}

// GetFinancialAssets returns the FinancialAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetFinancialAssets() float64 {
	if o == nil || IsNil(o.FinancialAssets) {
		var ret float64
		return ret
	}
	return *o.FinancialAssets
}

// GetFinancialAssetsOk returns a tuple with the FinancialAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetFinancialAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.FinancialAssets) {
		return nil, false
	}
	return o.FinancialAssets, true
}

// HasFinancialAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasFinancialAssets() bool {
	if o != nil && !IsNil(o.FinancialAssets) {
		return true
	}

	return false
}

// SetFinancialAssets gets a reference to the given float64 and assigns it to the FinancialAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetFinancialAssets(v float64) {
	o.FinancialAssets = &v
}

// GetIntangibleAssets returns the IntangibleAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetIntangibleAssets() float64 {
	if o == nil || IsNil(o.IntangibleAssets) {
		var ret float64
		return ret
	}
	return *o.IntangibleAssets
}

// GetIntangibleAssetsOk returns a tuple with the IntangibleAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetIntangibleAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.IntangibleAssets) {
		return nil, false
	}
	return o.IntangibleAssets, true
}

// HasIntangibleAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasIntangibleAssets() bool {
	if o != nil && !IsNil(o.IntangibleAssets) {
		return true
	}

	return false
}

// SetIntangibleAssets gets a reference to the given float64 and assigns it to the IntangibleAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetIntangibleAssets(v float64) {
	o.IntangibleAssets = &v
}

// GetInvestmentsAndAdvances returns the InvestmentsAndAdvances field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetInvestmentsAndAdvances() float64 {
	if o == nil || IsNil(o.InvestmentsAndAdvances) {
		var ret float64
		return ret
	}
	return *o.InvestmentsAndAdvances
}

// GetInvestmentsAndAdvancesOk returns a tuple with the InvestmentsAndAdvances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetInvestmentsAndAdvancesOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentsAndAdvances) {
		return nil, false
	}
	return o.InvestmentsAndAdvances, true
}

// HasInvestmentsAndAdvances returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasInvestmentsAndAdvances() bool {
	if o != nil && !IsNil(o.InvestmentsAndAdvances) {
		return true
	}

	return false
}

// SetInvestmentsAndAdvances gets a reference to the given float64 and assigns it to the InvestmentsAndAdvances field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetInvestmentsAndAdvances(v float64) {
	o.InvestmentsAndAdvances = &v
}

// GetOtherNonCurrentAssets returns the OtherNonCurrentAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetOtherNonCurrentAssets() float64 {
	if o == nil || IsNil(o.OtherNonCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.OtherNonCurrentAssets
}

// GetOtherNonCurrentAssetsOk returns a tuple with the OtherNonCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetOtherNonCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherNonCurrentAssets) {
		return nil, false
	}
	return o.OtherNonCurrentAssets, true
}

// HasOtherNonCurrentAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasOtherNonCurrentAssets() bool {
	if o != nil && !IsNil(o.OtherNonCurrentAssets) {
		return true
	}

	return false
}

// SetOtherNonCurrentAssets gets a reference to the given float64 and assigns it to the OtherNonCurrentAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetOtherNonCurrentAssets(v float64) {
	o.OtherNonCurrentAssets = &v
}

// GetTotalNonCurrentAssets returns the TotalNonCurrentAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetTotalNonCurrentAssets() float64 {
	if o == nil || IsNil(o.TotalNonCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.TotalNonCurrentAssets
}

// GetTotalNonCurrentAssetsOk returns a tuple with the TotalNonCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetTotalNonCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalNonCurrentAssets) {
		return nil, false
	}
	return o.TotalNonCurrentAssets, true
}

// HasTotalNonCurrentAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasTotalNonCurrentAssets() bool {
	if o != nil && !IsNil(o.TotalNonCurrentAssets) {
		return true
	}

	return false
}

// SetTotalNonCurrentAssets gets a reference to the given float64 and assigns it to the TotalNonCurrentAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetTotalNonCurrentAssets(v float64) {
	o.TotalNonCurrentAssets = &v
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.LandAndImprovements) {
		toSerialize["land_and_improvements"] = o.LandAndImprovements
	}
	if !IsNil(o.MachineryFurnitureEquipment) {
		toSerialize["machinery_furniture_equipment"] = o.MachineryFurnitureEquipment
	}
	if !IsNil(o.ConstructionInProgress) {
		toSerialize["construction_in_progress"] = o.ConstructionInProgress
	}
	if !IsNil(o.Leases) {
		toSerialize["leases"] = o.Leases
	}
	if !IsNil(o.AccumulatedDepreciation) {
		toSerialize["accumulated_depreciation"] = o.AccumulatedDepreciation
	}
	if !IsNil(o.Goodwill) {
		toSerialize["goodwill"] = o.Goodwill
	}
	if !IsNil(o.InvestmentProperties) {
		toSerialize["investment_properties"] = o.InvestmentProperties
	}
	if !IsNil(o.FinancialAssets) {
		toSerialize["financial_assets"] = o.FinancialAssets
	}
	if !IsNil(o.IntangibleAssets) {
		toSerialize["intangible_assets"] = o.IntangibleAssets
	}
	if !IsNil(o.InvestmentsAndAdvances) {
		toSerialize["investments_and_advances"] = o.InvestmentsAndAdvances
	}
	if !IsNil(o.OtherNonCurrentAssets) {
		toSerialize["other_non_current_assets"] = o.OtherNonCurrentAssets
	}
	if !IsNil(o.TotalNonCurrentAssets) {
		toSerialize["total_non_current_assets"] = o.TotalNonCurrentAssets
	}
	return toSerialize, nil
}

type NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets struct {
	value *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets
	isSet bool
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) Get() *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets {
	return v.value
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) Set(val *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets(val *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets {
	return &NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets{value: val, isSet: true}
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
