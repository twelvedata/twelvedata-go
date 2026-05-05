// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemEarningsPerShare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemEarningsPerShare{}

// IncomeStatementItemEarningsPerShare Earnings per share information
type IncomeStatementItemEarningsPerShare struct {
	// Diluted EPS
	DilutedEps *float64 `json:"diluted_eps,omitempty"`
	// Basic EPS
	BasicEps *float64 `json:"basic_eps,omitempty"`
	// Continuing and discontinued diluted EPS
	ContinuingAndDiscontinuedDilutedEps *float64 `json:"continuing_and_discontinued_diluted_eps,omitempty"`
	// Continuing and discontinued basic EPS
	ContinuingAndDiscontinuedBasicEps *float64 `json:"continuing_and_discontinued_basic_eps,omitempty"`
	// Normalized diluted EPS
	NormalizedDilutedEps *float64 `json:"normalized_diluted_eps,omitempty"`
	// Normalized basic EPS
	NormalizedBasicEps *float64 `json:"normalized_basic_eps,omitempty"`
	// Reported normalized diluted EPS
	ReportedNormalizedDilutedEps *float64 `json:"reported_normalized_diluted_eps,omitempty"`
	// Reported normalized basic EPS
	ReportedNormalizedBasicEps *float64 `json:"reported_normalized_basic_eps,omitempty"`
	// Diluted EPS other gains losses
	DilutedEpsOtherGainsLosses *float64 `json:"diluted_eps_other_gains_losses,omitempty"`
	// Tax loss carryforward diluted EPS
	TaxLossCarryforwardDilutedEps *float64 `json:"tax_loss_carryforward_diluted_eps,omitempty"`
	// Diluted accounting change
	DilutedAccountingChange *float64 `json:"diluted_accounting_change,omitempty"`
	// Diluted extraordinary
	DilutedExtraordinary *float64 `json:"diluted_extraordinary,omitempty"`
	// Diluted discontinuous operations
	DilutedDiscontinuousOperations *float64 `json:"diluted_discontinuous_operations,omitempty"`
	// Diluted continuous operations
	DilutedContinuousOperations *float64 `json:"diluted_continuous_operations,omitempty"`
	// Basic EPS other gains losses
	BasicEpsOtherGainsLosses *float64 `json:"basic_eps_other_gains_losses,omitempty"`
	// Tax loss carryforward basic EPS
	TaxLossCarryforwardBasicEps *float64 `json:"tax_loss_carryforward_basic_eps,omitempty"`
	// Basic accounting change
	BasicAccountingChange *float64 `json:"basic_accounting_change,omitempty"`
	// Basic extraordinary
	BasicExtraordinary *float64 `json:"basic_extraordinary,omitempty"`
	// Basic discontinuous operations
	BasicDiscontinuousOperations *float64 `json:"basic_discontinuous_operations,omitempty"`
	// Basic continuous operations
	BasicContinuousOperations *float64 `json:"basic_continuous_operations,omitempty"`
	// Diluted NI available to common stockholders
	DilutedNiAvailToCommonStockholders *float64 `json:"diluted_ni_avail_to_common_stockholders,omitempty"`
	// Average dilution earnings
	AverageDilutionEarnings *float64 `json:"average_dilution_earnings,omitempty"`
}

// NewIncomeStatementItemEarningsPerShare instantiates a new IncomeStatementItemEarningsPerShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemEarningsPerShare() *IncomeStatementItemEarningsPerShare {
	this := IncomeStatementItemEarningsPerShare{}
	return &this
}

// NewIncomeStatementItemEarningsPerShareWithDefaults instantiates a new IncomeStatementItemEarningsPerShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemEarningsPerShareWithDefaults() *IncomeStatementItemEarningsPerShare {
	this := IncomeStatementItemEarningsPerShare{}
	return &this
}

// GetDilutedEps returns the DilutedEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedEps() float64 {
	if o == nil || IsNil(o.DilutedEps) {
		var ret float64
		return ret
	}
	return *o.DilutedEps
}

// GetDilutedEpsOk returns a tuple with the DilutedEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedEps) {
		return nil, false
	}
	return o.DilutedEps, true
}

// HasDilutedEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedEps() bool {
	if o != nil && !IsNil(o.DilutedEps) {
		return true
	}

	return false
}

// SetDilutedEps gets a reference to the given float64 and assigns it to the DilutedEps field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedEps(v float64) {
	o.DilutedEps = &v
}

// GetBasicEps returns the BasicEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetBasicEps() float64 {
	if o == nil || IsNil(o.BasicEps) {
		var ret float64
		return ret
	}
	return *o.BasicEps
}

// GetBasicEpsOk returns a tuple with the BasicEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetBasicEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicEps) {
		return nil, false
	}
	return o.BasicEps, true
}

// HasBasicEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasBasicEps() bool {
	if o != nil && !IsNil(o.BasicEps) {
		return true
	}

	return false
}

// SetBasicEps gets a reference to the given float64 and assigns it to the BasicEps field.
func (o *IncomeStatementItemEarningsPerShare) SetBasicEps(v float64) {
	o.BasicEps = &v
}

// GetContinuingAndDiscontinuedDilutedEps returns the ContinuingAndDiscontinuedDilutedEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetContinuingAndDiscontinuedDilutedEps() float64 {
	if o == nil || IsNil(o.ContinuingAndDiscontinuedDilutedEps) {
		var ret float64
		return ret
	}
	return *o.ContinuingAndDiscontinuedDilutedEps
}

// GetContinuingAndDiscontinuedDilutedEpsOk returns a tuple with the ContinuingAndDiscontinuedDilutedEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetContinuingAndDiscontinuedDilutedEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ContinuingAndDiscontinuedDilutedEps) {
		return nil, false
	}
	return o.ContinuingAndDiscontinuedDilutedEps, true
}

// HasContinuingAndDiscontinuedDilutedEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasContinuingAndDiscontinuedDilutedEps() bool {
	if o != nil && !IsNil(o.ContinuingAndDiscontinuedDilutedEps) {
		return true
	}

	return false
}

// SetContinuingAndDiscontinuedDilutedEps gets a reference to the given float64 and assigns it to the ContinuingAndDiscontinuedDilutedEps field.
func (o *IncomeStatementItemEarningsPerShare) SetContinuingAndDiscontinuedDilutedEps(v float64) {
	o.ContinuingAndDiscontinuedDilutedEps = &v
}

// GetContinuingAndDiscontinuedBasicEps returns the ContinuingAndDiscontinuedBasicEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetContinuingAndDiscontinuedBasicEps() float64 {
	if o == nil || IsNil(o.ContinuingAndDiscontinuedBasicEps) {
		var ret float64
		return ret
	}
	return *o.ContinuingAndDiscontinuedBasicEps
}

// GetContinuingAndDiscontinuedBasicEpsOk returns a tuple with the ContinuingAndDiscontinuedBasicEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetContinuingAndDiscontinuedBasicEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ContinuingAndDiscontinuedBasicEps) {
		return nil, false
	}
	return o.ContinuingAndDiscontinuedBasicEps, true
}

// HasContinuingAndDiscontinuedBasicEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasContinuingAndDiscontinuedBasicEps() bool {
	if o != nil && !IsNil(o.ContinuingAndDiscontinuedBasicEps) {
		return true
	}

	return false
}

// SetContinuingAndDiscontinuedBasicEps gets a reference to the given float64 and assigns it to the ContinuingAndDiscontinuedBasicEps field.
func (o *IncomeStatementItemEarningsPerShare) SetContinuingAndDiscontinuedBasicEps(v float64) {
	o.ContinuingAndDiscontinuedBasicEps = &v
}

// GetNormalizedDilutedEps returns the NormalizedDilutedEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetNormalizedDilutedEps() float64 {
	if o == nil || IsNil(o.NormalizedDilutedEps) {
		var ret float64
		return ret
	}
	return *o.NormalizedDilutedEps
}

// GetNormalizedDilutedEpsOk returns a tuple with the NormalizedDilutedEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetNormalizedDilutedEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.NormalizedDilutedEps) {
		return nil, false
	}
	return o.NormalizedDilutedEps, true
}

// HasNormalizedDilutedEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasNormalizedDilutedEps() bool {
	if o != nil && !IsNil(o.NormalizedDilutedEps) {
		return true
	}

	return false
}

// SetNormalizedDilutedEps gets a reference to the given float64 and assigns it to the NormalizedDilutedEps field.
func (o *IncomeStatementItemEarningsPerShare) SetNormalizedDilutedEps(v float64) {
	o.NormalizedDilutedEps = &v
}

// GetNormalizedBasicEps returns the NormalizedBasicEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetNormalizedBasicEps() float64 {
	if o == nil || IsNil(o.NormalizedBasicEps) {
		var ret float64
		return ret
	}
	return *o.NormalizedBasicEps
}

// GetNormalizedBasicEpsOk returns a tuple with the NormalizedBasicEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetNormalizedBasicEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.NormalizedBasicEps) {
		return nil, false
	}
	return o.NormalizedBasicEps, true
}

// HasNormalizedBasicEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasNormalizedBasicEps() bool {
	if o != nil && !IsNil(o.NormalizedBasicEps) {
		return true
	}

	return false
}

// SetNormalizedBasicEps gets a reference to the given float64 and assigns it to the NormalizedBasicEps field.
func (o *IncomeStatementItemEarningsPerShare) SetNormalizedBasicEps(v float64) {
	o.NormalizedBasicEps = &v
}

// GetReportedNormalizedDilutedEps returns the ReportedNormalizedDilutedEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetReportedNormalizedDilutedEps() float64 {
	if o == nil || IsNil(o.ReportedNormalizedDilutedEps) {
		var ret float64
		return ret
	}
	return *o.ReportedNormalizedDilutedEps
}

// GetReportedNormalizedDilutedEpsOk returns a tuple with the ReportedNormalizedDilutedEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetReportedNormalizedDilutedEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ReportedNormalizedDilutedEps) {
		return nil, false
	}
	return o.ReportedNormalizedDilutedEps, true
}

// HasReportedNormalizedDilutedEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasReportedNormalizedDilutedEps() bool {
	if o != nil && !IsNil(o.ReportedNormalizedDilutedEps) {
		return true
	}

	return false
}

// SetReportedNormalizedDilutedEps gets a reference to the given float64 and assigns it to the ReportedNormalizedDilutedEps field.
func (o *IncomeStatementItemEarningsPerShare) SetReportedNormalizedDilutedEps(v float64) {
	o.ReportedNormalizedDilutedEps = &v
}

// GetReportedNormalizedBasicEps returns the ReportedNormalizedBasicEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetReportedNormalizedBasicEps() float64 {
	if o == nil || IsNil(o.ReportedNormalizedBasicEps) {
		var ret float64
		return ret
	}
	return *o.ReportedNormalizedBasicEps
}

// GetReportedNormalizedBasicEpsOk returns a tuple with the ReportedNormalizedBasicEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetReportedNormalizedBasicEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ReportedNormalizedBasicEps) {
		return nil, false
	}
	return o.ReportedNormalizedBasicEps, true
}

// HasReportedNormalizedBasicEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasReportedNormalizedBasicEps() bool {
	if o != nil && !IsNil(o.ReportedNormalizedBasicEps) {
		return true
	}

	return false
}

// SetReportedNormalizedBasicEps gets a reference to the given float64 and assigns it to the ReportedNormalizedBasicEps field.
func (o *IncomeStatementItemEarningsPerShare) SetReportedNormalizedBasicEps(v float64) {
	o.ReportedNormalizedBasicEps = &v
}

// GetDilutedEpsOtherGainsLosses returns the DilutedEpsOtherGainsLosses field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedEpsOtherGainsLosses() float64 {
	if o == nil || IsNil(o.DilutedEpsOtherGainsLosses) {
		var ret float64
		return ret
	}
	return *o.DilutedEpsOtherGainsLosses
}

// GetDilutedEpsOtherGainsLossesOk returns a tuple with the DilutedEpsOtherGainsLosses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedEpsOtherGainsLossesOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedEpsOtherGainsLosses) {
		return nil, false
	}
	return o.DilutedEpsOtherGainsLosses, true
}

// HasDilutedEpsOtherGainsLosses returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedEpsOtherGainsLosses() bool {
	if o != nil && !IsNil(o.DilutedEpsOtherGainsLosses) {
		return true
	}

	return false
}

// SetDilutedEpsOtherGainsLosses gets a reference to the given float64 and assigns it to the DilutedEpsOtherGainsLosses field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedEpsOtherGainsLosses(v float64) {
	o.DilutedEpsOtherGainsLosses = &v
}

// GetTaxLossCarryforwardDilutedEps returns the TaxLossCarryforwardDilutedEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetTaxLossCarryforwardDilutedEps() float64 {
	if o == nil || IsNil(o.TaxLossCarryforwardDilutedEps) {
		var ret float64
		return ret
	}
	return *o.TaxLossCarryforwardDilutedEps
}

// GetTaxLossCarryforwardDilutedEpsOk returns a tuple with the TaxLossCarryforwardDilutedEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetTaxLossCarryforwardDilutedEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxLossCarryforwardDilutedEps) {
		return nil, false
	}
	return o.TaxLossCarryforwardDilutedEps, true
}

// HasTaxLossCarryforwardDilutedEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasTaxLossCarryforwardDilutedEps() bool {
	if o != nil && !IsNil(o.TaxLossCarryforwardDilutedEps) {
		return true
	}

	return false
}

// SetTaxLossCarryforwardDilutedEps gets a reference to the given float64 and assigns it to the TaxLossCarryforwardDilutedEps field.
func (o *IncomeStatementItemEarningsPerShare) SetTaxLossCarryforwardDilutedEps(v float64) {
	o.TaxLossCarryforwardDilutedEps = &v
}

// GetDilutedAccountingChange returns the DilutedAccountingChange field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedAccountingChange() float64 {
	if o == nil || IsNil(o.DilutedAccountingChange) {
		var ret float64
		return ret
	}
	return *o.DilutedAccountingChange
}

// GetDilutedAccountingChangeOk returns a tuple with the DilutedAccountingChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedAccountingChangeOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedAccountingChange) {
		return nil, false
	}
	return o.DilutedAccountingChange, true
}

// HasDilutedAccountingChange returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedAccountingChange() bool {
	if o != nil && !IsNil(o.DilutedAccountingChange) {
		return true
	}

	return false
}

// SetDilutedAccountingChange gets a reference to the given float64 and assigns it to the DilutedAccountingChange field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedAccountingChange(v float64) {
	o.DilutedAccountingChange = &v
}

// GetDilutedExtraordinary returns the DilutedExtraordinary field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedExtraordinary() float64 {
	if o == nil || IsNil(o.DilutedExtraordinary) {
		var ret float64
		return ret
	}
	return *o.DilutedExtraordinary
}

// GetDilutedExtraordinaryOk returns a tuple with the DilutedExtraordinary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedExtraordinaryOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedExtraordinary) {
		return nil, false
	}
	return o.DilutedExtraordinary, true
}

// HasDilutedExtraordinary returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedExtraordinary() bool {
	if o != nil && !IsNil(o.DilutedExtraordinary) {
		return true
	}

	return false
}

// SetDilutedExtraordinary gets a reference to the given float64 and assigns it to the DilutedExtraordinary field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedExtraordinary(v float64) {
	o.DilutedExtraordinary = &v
}

// GetDilutedDiscontinuousOperations returns the DilutedDiscontinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedDiscontinuousOperations() float64 {
	if o == nil || IsNil(o.DilutedDiscontinuousOperations) {
		var ret float64
		return ret
	}
	return *o.DilutedDiscontinuousOperations
}

// GetDilutedDiscontinuousOperationsOk returns a tuple with the DilutedDiscontinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedDiscontinuousOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedDiscontinuousOperations) {
		return nil, false
	}
	return o.DilutedDiscontinuousOperations, true
}

// HasDilutedDiscontinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedDiscontinuousOperations() bool {
	if o != nil && !IsNil(o.DilutedDiscontinuousOperations) {
		return true
	}

	return false
}

// SetDilutedDiscontinuousOperations gets a reference to the given float64 and assigns it to the DilutedDiscontinuousOperations field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedDiscontinuousOperations(v float64) {
	o.DilutedDiscontinuousOperations = &v
}

// GetDilutedContinuousOperations returns the DilutedContinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedContinuousOperations() float64 {
	if o == nil || IsNil(o.DilutedContinuousOperations) {
		var ret float64
		return ret
	}
	return *o.DilutedContinuousOperations
}

// GetDilutedContinuousOperationsOk returns a tuple with the DilutedContinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedContinuousOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedContinuousOperations) {
		return nil, false
	}
	return o.DilutedContinuousOperations, true
}

// HasDilutedContinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedContinuousOperations() bool {
	if o != nil && !IsNil(o.DilutedContinuousOperations) {
		return true
	}

	return false
}

// SetDilutedContinuousOperations gets a reference to the given float64 and assigns it to the DilutedContinuousOperations field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedContinuousOperations(v float64) {
	o.DilutedContinuousOperations = &v
}

// GetBasicEpsOtherGainsLosses returns the BasicEpsOtherGainsLosses field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetBasicEpsOtherGainsLosses() float64 {
	if o == nil || IsNil(o.BasicEpsOtherGainsLosses) {
		var ret float64
		return ret
	}
	return *o.BasicEpsOtherGainsLosses
}

// GetBasicEpsOtherGainsLossesOk returns a tuple with the BasicEpsOtherGainsLosses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetBasicEpsOtherGainsLossesOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicEpsOtherGainsLosses) {
		return nil, false
	}
	return o.BasicEpsOtherGainsLosses, true
}

// HasBasicEpsOtherGainsLosses returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasBasicEpsOtherGainsLosses() bool {
	if o != nil && !IsNil(o.BasicEpsOtherGainsLosses) {
		return true
	}

	return false
}

// SetBasicEpsOtherGainsLosses gets a reference to the given float64 and assigns it to the BasicEpsOtherGainsLosses field.
func (o *IncomeStatementItemEarningsPerShare) SetBasicEpsOtherGainsLosses(v float64) {
	o.BasicEpsOtherGainsLosses = &v
}

// GetTaxLossCarryforwardBasicEps returns the TaxLossCarryforwardBasicEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetTaxLossCarryforwardBasicEps() float64 {
	if o == nil || IsNil(o.TaxLossCarryforwardBasicEps) {
		var ret float64
		return ret
	}
	return *o.TaxLossCarryforwardBasicEps
}

// GetTaxLossCarryforwardBasicEpsOk returns a tuple with the TaxLossCarryforwardBasicEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetTaxLossCarryforwardBasicEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxLossCarryforwardBasicEps) {
		return nil, false
	}
	return o.TaxLossCarryforwardBasicEps, true
}

// HasTaxLossCarryforwardBasicEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasTaxLossCarryforwardBasicEps() bool {
	if o != nil && !IsNil(o.TaxLossCarryforwardBasicEps) {
		return true
	}

	return false
}

// SetTaxLossCarryforwardBasicEps gets a reference to the given float64 and assigns it to the TaxLossCarryforwardBasicEps field.
func (o *IncomeStatementItemEarningsPerShare) SetTaxLossCarryforwardBasicEps(v float64) {
	o.TaxLossCarryforwardBasicEps = &v
}

// GetBasicAccountingChange returns the BasicAccountingChange field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetBasicAccountingChange() float64 {
	if o == nil || IsNil(o.BasicAccountingChange) {
		var ret float64
		return ret
	}
	return *o.BasicAccountingChange
}

// GetBasicAccountingChangeOk returns a tuple with the BasicAccountingChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetBasicAccountingChangeOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicAccountingChange) {
		return nil, false
	}
	return o.BasicAccountingChange, true
}

// HasBasicAccountingChange returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasBasicAccountingChange() bool {
	if o != nil && !IsNil(o.BasicAccountingChange) {
		return true
	}

	return false
}

// SetBasicAccountingChange gets a reference to the given float64 and assigns it to the BasicAccountingChange field.
func (o *IncomeStatementItemEarningsPerShare) SetBasicAccountingChange(v float64) {
	o.BasicAccountingChange = &v
}

// GetBasicExtraordinary returns the BasicExtraordinary field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetBasicExtraordinary() float64 {
	if o == nil || IsNil(o.BasicExtraordinary) {
		var ret float64
		return ret
	}
	return *o.BasicExtraordinary
}

// GetBasicExtraordinaryOk returns a tuple with the BasicExtraordinary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetBasicExtraordinaryOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicExtraordinary) {
		return nil, false
	}
	return o.BasicExtraordinary, true
}

// HasBasicExtraordinary returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasBasicExtraordinary() bool {
	if o != nil && !IsNil(o.BasicExtraordinary) {
		return true
	}

	return false
}

// SetBasicExtraordinary gets a reference to the given float64 and assigns it to the BasicExtraordinary field.
func (o *IncomeStatementItemEarningsPerShare) SetBasicExtraordinary(v float64) {
	o.BasicExtraordinary = &v
}

// GetBasicDiscontinuousOperations returns the BasicDiscontinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetBasicDiscontinuousOperations() float64 {
	if o == nil || IsNil(o.BasicDiscontinuousOperations) {
		var ret float64
		return ret
	}
	return *o.BasicDiscontinuousOperations
}

// GetBasicDiscontinuousOperationsOk returns a tuple with the BasicDiscontinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetBasicDiscontinuousOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicDiscontinuousOperations) {
		return nil, false
	}
	return o.BasicDiscontinuousOperations, true
}

// HasBasicDiscontinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasBasicDiscontinuousOperations() bool {
	if o != nil && !IsNil(o.BasicDiscontinuousOperations) {
		return true
	}

	return false
}

// SetBasicDiscontinuousOperations gets a reference to the given float64 and assigns it to the BasicDiscontinuousOperations field.
func (o *IncomeStatementItemEarningsPerShare) SetBasicDiscontinuousOperations(v float64) {
	o.BasicDiscontinuousOperations = &v
}

// GetBasicContinuousOperations returns the BasicContinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetBasicContinuousOperations() float64 {
	if o == nil || IsNil(o.BasicContinuousOperations) {
		var ret float64
		return ret
	}
	return *o.BasicContinuousOperations
}

// GetBasicContinuousOperationsOk returns a tuple with the BasicContinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetBasicContinuousOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicContinuousOperations) {
		return nil, false
	}
	return o.BasicContinuousOperations, true
}

// HasBasicContinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasBasicContinuousOperations() bool {
	if o != nil && !IsNil(o.BasicContinuousOperations) {
		return true
	}

	return false
}

// SetBasicContinuousOperations gets a reference to the given float64 and assigns it to the BasicContinuousOperations field.
func (o *IncomeStatementItemEarningsPerShare) SetBasicContinuousOperations(v float64) {
	o.BasicContinuousOperations = &v
}

// GetDilutedNiAvailToCommonStockholders returns the DilutedNiAvailToCommonStockholders field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedNiAvailToCommonStockholders() float64 {
	if o == nil || IsNil(o.DilutedNiAvailToCommonStockholders) {
		var ret float64
		return ret
	}
	return *o.DilutedNiAvailToCommonStockholders
}

// GetDilutedNiAvailToCommonStockholdersOk returns a tuple with the DilutedNiAvailToCommonStockholders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedNiAvailToCommonStockholdersOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedNiAvailToCommonStockholders) {
		return nil, false
	}
	return o.DilutedNiAvailToCommonStockholders, true
}

// HasDilutedNiAvailToCommonStockholders returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedNiAvailToCommonStockholders() bool {
	if o != nil && !IsNil(o.DilutedNiAvailToCommonStockholders) {
		return true
	}

	return false
}

// SetDilutedNiAvailToCommonStockholders gets a reference to the given float64 and assigns it to the DilutedNiAvailToCommonStockholders field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedNiAvailToCommonStockholders(v float64) {
	o.DilutedNiAvailToCommonStockholders = &v
}

// GetAverageDilutionEarnings returns the AverageDilutionEarnings field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetAverageDilutionEarnings() float64 {
	if o == nil || IsNil(o.AverageDilutionEarnings) {
		var ret float64
		return ret
	}
	return *o.AverageDilutionEarnings
}

// GetAverageDilutionEarningsOk returns a tuple with the AverageDilutionEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetAverageDilutionEarningsOk() (*float64, bool) {
	if o == nil || IsNil(o.AverageDilutionEarnings) {
		return nil, false
	}
	return o.AverageDilutionEarnings, true
}

// HasAverageDilutionEarnings returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasAverageDilutionEarnings() bool {
	if o != nil && !IsNil(o.AverageDilutionEarnings) {
		return true
	}

	return false
}

// SetAverageDilutionEarnings gets a reference to the given float64 and assigns it to the AverageDilutionEarnings field.
func (o *IncomeStatementItemEarningsPerShare) SetAverageDilutionEarnings(v float64) {
	o.AverageDilutionEarnings = &v
}

func (o IncomeStatementItemEarningsPerShare) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemEarningsPerShare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DilutedEps) {
		toSerialize["diluted_eps"] = o.DilutedEps
	}
	if !IsNil(o.BasicEps) {
		toSerialize["basic_eps"] = o.BasicEps
	}
	if !IsNil(o.ContinuingAndDiscontinuedDilutedEps) {
		toSerialize["continuing_and_discontinued_diluted_eps"] = o.ContinuingAndDiscontinuedDilutedEps
	}
	if !IsNil(o.ContinuingAndDiscontinuedBasicEps) {
		toSerialize["continuing_and_discontinued_basic_eps"] = o.ContinuingAndDiscontinuedBasicEps
	}
	if !IsNil(o.NormalizedDilutedEps) {
		toSerialize["normalized_diluted_eps"] = o.NormalizedDilutedEps
	}
	if !IsNil(o.NormalizedBasicEps) {
		toSerialize["normalized_basic_eps"] = o.NormalizedBasicEps
	}
	if !IsNil(o.ReportedNormalizedDilutedEps) {
		toSerialize["reported_normalized_diluted_eps"] = o.ReportedNormalizedDilutedEps
	}
	if !IsNil(o.ReportedNormalizedBasicEps) {
		toSerialize["reported_normalized_basic_eps"] = o.ReportedNormalizedBasicEps
	}
	if !IsNil(o.DilutedEpsOtherGainsLosses) {
		toSerialize["diluted_eps_other_gains_losses"] = o.DilutedEpsOtherGainsLosses
	}
	if !IsNil(o.TaxLossCarryforwardDilutedEps) {
		toSerialize["tax_loss_carryforward_diluted_eps"] = o.TaxLossCarryforwardDilutedEps
	}
	if !IsNil(o.DilutedAccountingChange) {
		toSerialize["diluted_accounting_change"] = o.DilutedAccountingChange
	}
	if !IsNil(o.DilutedExtraordinary) {
		toSerialize["diluted_extraordinary"] = o.DilutedExtraordinary
	}
	if !IsNil(o.DilutedDiscontinuousOperations) {
		toSerialize["diluted_discontinuous_operations"] = o.DilutedDiscontinuousOperations
	}
	if !IsNil(o.DilutedContinuousOperations) {
		toSerialize["diluted_continuous_operations"] = o.DilutedContinuousOperations
	}
	if !IsNil(o.BasicEpsOtherGainsLosses) {
		toSerialize["basic_eps_other_gains_losses"] = o.BasicEpsOtherGainsLosses
	}
	if !IsNil(o.TaxLossCarryforwardBasicEps) {
		toSerialize["tax_loss_carryforward_basic_eps"] = o.TaxLossCarryforwardBasicEps
	}
	if !IsNil(o.BasicAccountingChange) {
		toSerialize["basic_accounting_change"] = o.BasicAccountingChange
	}
	if !IsNil(o.BasicExtraordinary) {
		toSerialize["basic_extraordinary"] = o.BasicExtraordinary
	}
	if !IsNil(o.BasicDiscontinuousOperations) {
		toSerialize["basic_discontinuous_operations"] = o.BasicDiscontinuousOperations
	}
	if !IsNil(o.BasicContinuousOperations) {
		toSerialize["basic_continuous_operations"] = o.BasicContinuousOperations
	}
	if !IsNil(o.DilutedNiAvailToCommonStockholders) {
		toSerialize["diluted_ni_avail_to_common_stockholders"] = o.DilutedNiAvailToCommonStockholders
	}
	if !IsNil(o.AverageDilutionEarnings) {
		toSerialize["average_dilution_earnings"] = o.AverageDilutionEarnings
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemEarningsPerShare struct {
	value *IncomeStatementItemEarningsPerShare
	isSet bool
}

func (v NullableIncomeStatementItemEarningsPerShare) Get() *IncomeStatementItemEarningsPerShare {
	return v.value
}

func (v *NullableIncomeStatementItemEarningsPerShare) Set(val *IncomeStatementItemEarningsPerShare) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemEarningsPerShare) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemEarningsPerShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemEarningsPerShare(val *IncomeStatementItemEarningsPerShare) *NullableIncomeStatementItemEarningsPerShare {
	return &NullableIncomeStatementItemEarningsPerShare{value: val, isSet: true}
}

func (v NullableIncomeStatementItemEarningsPerShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemEarningsPerShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
