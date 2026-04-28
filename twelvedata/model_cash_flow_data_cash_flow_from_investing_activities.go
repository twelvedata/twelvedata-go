/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the CashFlowDataCashFlowFromInvestingActivities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashFlowDataCashFlowFromInvestingActivities{}

// CashFlowDataCashFlowFromInvestingActivities Cash flow from investing activities
type CashFlowDataCashFlowFromInvestingActivities struct {
	// Investing cash flow
	InvestingCashFlow *float64 `json:"investing_cash_flow,omitempty"`
	// Cash flow from continuing investing activities
	CashFlowFromContinuingInvestingActivities *float64 `json:"cash_flow_from_continuing_investing_activities,omitempty"`
	// Cash from discontinued investing activities
	CashFromDiscontinuedInvestingActivities *float64 `json:"cash_from_discontinued_investing_activities,omitempty"`
	// Net other investing changes
	NetOtherInvestingChanges *float64 `json:"net_other_investing_changes,omitempty"`
	// Interest received cfi
	InterestReceivedCfi *float64 `json:"interest_received_cfi,omitempty"`
	// Dividends received cfi
	DividendsReceivedCfi *float64 `json:"dividends_received_cfi,omitempty"`
	// Net investment purchase and sale
	NetInvestmentPurchaseAndSale *float64 `json:"net_investment_purchase_and_sale,omitempty"`
	// Sale of investment
	SaleOfInvestment *float64 `json:"sale_of_investment,omitempty"`
	// Purchase of investment
	PurchaseOfInvestment *float64 `json:"purchase_of_investment,omitempty"`
	// Net investment properties purchase and sale
	NetInvestmentPropertiesPurchaseAndSale *float64 `json:"net_investment_properties_purchase_and_sale,omitempty"`
	// Sale of investment properties
	SaleOfInvestmentProperties *float64 `json:"sale_of_investment_properties,omitempty"`
	// Purchase of investment properties
	PurchaseOfInvestmentProperties *float64 `json:"purchase_of_investment_properties,omitempty"`
	// Net business purchase and sale
	NetBusinessPurchaseAndSale *float64 `json:"net_business_purchase_and_sale,omitempty"`
	// Sale of business
	SaleOfBusiness *float64 `json:"sale_of_business,omitempty"`
	// Purchase of business
	PurchaseOfBusiness *float64 `json:"purchase_of_business,omitempty"`
	// Net intangibles purchase and sale
	NetIntangiblesPurchaseAndSale *float64 `json:"net_intangibles_purchase_and_sale,omitempty"`
	// Sale of intangibles
	SaleOfIntangibles *float64 `json:"sale_of_intangibles,omitempty"`
	// Purchase of intangibles
	PurchaseOfIntangibles *float64 `json:"purchase_of_intangibles,omitempty"`
	// Net ppe purchase and sale
	NetPpePurchaseAndSale *float64 `json:"net_ppe_purchase_and_sale,omitempty"`
	// Sale of ppe
	SaleOfPpe *float64 `json:"sale_of_ppe,omitempty"`
	// Purchase of ppe
	PurchaseOfPpe *float64 `json:"purchase_of_ppe,omitempty"`
	// Capital expenditure reported
	CapitalExpenditureReported *float64 `json:"capital_expenditure_reported,omitempty"`
	// Capital expenditure
	CapitalExpenditure *float64 `json:"capital_expenditure,omitempty"`
}

// NewCashFlowDataCashFlowFromInvestingActivities instantiates a new CashFlowDataCashFlowFromInvestingActivities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashFlowDataCashFlowFromInvestingActivities() *CashFlowDataCashFlowFromInvestingActivities {
	this := CashFlowDataCashFlowFromInvestingActivities{}
	return &this
}

// NewCashFlowDataCashFlowFromInvestingActivitiesWithDefaults instantiates a new CashFlowDataCashFlowFromInvestingActivities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashFlowDataCashFlowFromInvestingActivitiesWithDefaults() *CashFlowDataCashFlowFromInvestingActivities {
	this := CashFlowDataCashFlowFromInvestingActivities{}
	return &this
}

// GetInvestingCashFlow returns the InvestingCashFlow field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetInvestingCashFlow() float64 {
	if o == nil || IsNil(o.InvestingCashFlow) {
		var ret float64
		return ret
	}
	return *o.InvestingCashFlow
}

// GetInvestingCashFlowOk returns a tuple with the InvestingCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetInvestingCashFlowOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestingCashFlow) {
		return nil, false
	}
	return o.InvestingCashFlow, true
}

// HasInvestingCashFlow returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasInvestingCashFlow() bool {
	if o != nil && !IsNil(o.InvestingCashFlow) {
		return true
	}

	return false
}

// SetInvestingCashFlow gets a reference to the given float64 and assigns it to the InvestingCashFlow field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetInvestingCashFlow(v float64) {
	o.InvestingCashFlow = &v
}

// GetCashFlowFromContinuingInvestingActivities returns the CashFlowFromContinuingInvestingActivities field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetCashFlowFromContinuingInvestingActivities() float64 {
	if o == nil || IsNil(o.CashFlowFromContinuingInvestingActivities) {
		var ret float64
		return ret
	}
	return *o.CashFlowFromContinuingInvestingActivities
}

// GetCashFlowFromContinuingInvestingActivitiesOk returns a tuple with the CashFlowFromContinuingInvestingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetCashFlowFromContinuingInvestingActivitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.CashFlowFromContinuingInvestingActivities) {
		return nil, false
	}
	return o.CashFlowFromContinuingInvestingActivities, true
}

// HasCashFlowFromContinuingInvestingActivities returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasCashFlowFromContinuingInvestingActivities() bool {
	if o != nil && !IsNil(o.CashFlowFromContinuingInvestingActivities) {
		return true
	}

	return false
}

// SetCashFlowFromContinuingInvestingActivities gets a reference to the given float64 and assigns it to the CashFlowFromContinuingInvestingActivities field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetCashFlowFromContinuingInvestingActivities(v float64) {
	o.CashFlowFromContinuingInvestingActivities = &v
}

// GetCashFromDiscontinuedInvestingActivities returns the CashFromDiscontinuedInvestingActivities field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetCashFromDiscontinuedInvestingActivities() float64 {
	if o == nil || IsNil(o.CashFromDiscontinuedInvestingActivities) {
		var ret float64
		return ret
	}
	return *o.CashFromDiscontinuedInvestingActivities
}

// GetCashFromDiscontinuedInvestingActivitiesOk returns a tuple with the CashFromDiscontinuedInvestingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetCashFromDiscontinuedInvestingActivitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.CashFromDiscontinuedInvestingActivities) {
		return nil, false
	}
	return o.CashFromDiscontinuedInvestingActivities, true
}

// HasCashFromDiscontinuedInvestingActivities returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasCashFromDiscontinuedInvestingActivities() bool {
	if o != nil && !IsNil(o.CashFromDiscontinuedInvestingActivities) {
		return true
	}

	return false
}

// SetCashFromDiscontinuedInvestingActivities gets a reference to the given float64 and assigns it to the CashFromDiscontinuedInvestingActivities field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetCashFromDiscontinuedInvestingActivities(v float64) {
	o.CashFromDiscontinuedInvestingActivities = &v
}

// GetNetOtherInvestingChanges returns the NetOtherInvestingChanges field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetNetOtherInvestingChanges() float64 {
	if o == nil || IsNil(o.NetOtherInvestingChanges) {
		var ret float64
		return ret
	}
	return *o.NetOtherInvestingChanges
}

// GetNetOtherInvestingChangesOk returns a tuple with the NetOtherInvestingChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetNetOtherInvestingChangesOk() (*float64, bool) {
	if o == nil || IsNil(o.NetOtherInvestingChanges) {
		return nil, false
	}
	return o.NetOtherInvestingChanges, true
}

// HasNetOtherInvestingChanges returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasNetOtherInvestingChanges() bool {
	if o != nil && !IsNil(o.NetOtherInvestingChanges) {
		return true
	}

	return false
}

// SetNetOtherInvestingChanges gets a reference to the given float64 and assigns it to the NetOtherInvestingChanges field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetNetOtherInvestingChanges(v float64) {
	o.NetOtherInvestingChanges = &v
}

// GetInterestReceivedCfi returns the InterestReceivedCfi field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetInterestReceivedCfi() float64 {
	if o == nil || IsNil(o.InterestReceivedCfi) {
		var ret float64
		return ret
	}
	return *o.InterestReceivedCfi
}

// GetInterestReceivedCfiOk returns a tuple with the InterestReceivedCfi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetInterestReceivedCfiOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestReceivedCfi) {
		return nil, false
	}
	return o.InterestReceivedCfi, true
}

// HasInterestReceivedCfi returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasInterestReceivedCfi() bool {
	if o != nil && !IsNil(o.InterestReceivedCfi) {
		return true
	}

	return false
}

// SetInterestReceivedCfi gets a reference to the given float64 and assigns it to the InterestReceivedCfi field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetInterestReceivedCfi(v float64) {
	o.InterestReceivedCfi = &v
}

// GetDividendsReceivedCfi returns the DividendsReceivedCfi field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetDividendsReceivedCfi() float64 {
	if o == nil || IsNil(o.DividendsReceivedCfi) {
		var ret float64
		return ret
	}
	return *o.DividendsReceivedCfi
}

// GetDividendsReceivedCfiOk returns a tuple with the DividendsReceivedCfi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetDividendsReceivedCfiOk() (*float64, bool) {
	if o == nil || IsNil(o.DividendsReceivedCfi) {
		return nil, false
	}
	return o.DividendsReceivedCfi, true
}

// HasDividendsReceivedCfi returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasDividendsReceivedCfi() bool {
	if o != nil && !IsNil(o.DividendsReceivedCfi) {
		return true
	}

	return false
}

// SetDividendsReceivedCfi gets a reference to the given float64 and assigns it to the DividendsReceivedCfi field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetDividendsReceivedCfi(v float64) {
	o.DividendsReceivedCfi = &v
}

// GetNetInvestmentPurchaseAndSale returns the NetInvestmentPurchaseAndSale field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetNetInvestmentPurchaseAndSale() float64 {
	if o == nil || IsNil(o.NetInvestmentPurchaseAndSale) {
		var ret float64
		return ret
	}
	return *o.NetInvestmentPurchaseAndSale
}

// GetNetInvestmentPurchaseAndSaleOk returns a tuple with the NetInvestmentPurchaseAndSale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetNetInvestmentPurchaseAndSaleOk() (*float64, bool) {
	if o == nil || IsNil(o.NetInvestmentPurchaseAndSale) {
		return nil, false
	}
	return o.NetInvestmentPurchaseAndSale, true
}

// HasNetInvestmentPurchaseAndSale returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasNetInvestmentPurchaseAndSale() bool {
	if o != nil && !IsNil(o.NetInvestmentPurchaseAndSale) {
		return true
	}

	return false
}

// SetNetInvestmentPurchaseAndSale gets a reference to the given float64 and assigns it to the NetInvestmentPurchaseAndSale field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetNetInvestmentPurchaseAndSale(v float64) {
	o.NetInvestmentPurchaseAndSale = &v
}

// GetSaleOfInvestment returns the SaleOfInvestment field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetSaleOfInvestment() float64 {
	if o == nil || IsNil(o.SaleOfInvestment) {
		var ret float64
		return ret
	}
	return *o.SaleOfInvestment
}

// GetSaleOfInvestmentOk returns a tuple with the SaleOfInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetSaleOfInvestmentOk() (*float64, bool) {
	if o == nil || IsNil(o.SaleOfInvestment) {
		return nil, false
	}
	return o.SaleOfInvestment, true
}

// HasSaleOfInvestment returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasSaleOfInvestment() bool {
	if o != nil && !IsNil(o.SaleOfInvestment) {
		return true
	}

	return false
}

// SetSaleOfInvestment gets a reference to the given float64 and assigns it to the SaleOfInvestment field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetSaleOfInvestment(v float64) {
	o.SaleOfInvestment = &v
}

// GetPurchaseOfInvestment returns the PurchaseOfInvestment field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetPurchaseOfInvestment() float64 {
	if o == nil || IsNil(o.PurchaseOfInvestment) {
		var ret float64
		return ret
	}
	return *o.PurchaseOfInvestment
}

// GetPurchaseOfInvestmentOk returns a tuple with the PurchaseOfInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetPurchaseOfInvestmentOk() (*float64, bool) {
	if o == nil || IsNil(o.PurchaseOfInvestment) {
		return nil, false
	}
	return o.PurchaseOfInvestment, true
}

// HasPurchaseOfInvestment returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasPurchaseOfInvestment() bool {
	if o != nil && !IsNil(o.PurchaseOfInvestment) {
		return true
	}

	return false
}

// SetPurchaseOfInvestment gets a reference to the given float64 and assigns it to the PurchaseOfInvestment field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetPurchaseOfInvestment(v float64) {
	o.PurchaseOfInvestment = &v
}

// GetNetInvestmentPropertiesPurchaseAndSale returns the NetInvestmentPropertiesPurchaseAndSale field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetNetInvestmentPropertiesPurchaseAndSale() float64 {
	if o == nil || IsNil(o.NetInvestmentPropertiesPurchaseAndSale) {
		var ret float64
		return ret
	}
	return *o.NetInvestmentPropertiesPurchaseAndSale
}

// GetNetInvestmentPropertiesPurchaseAndSaleOk returns a tuple with the NetInvestmentPropertiesPurchaseAndSale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetNetInvestmentPropertiesPurchaseAndSaleOk() (*float64, bool) {
	if o == nil || IsNil(o.NetInvestmentPropertiesPurchaseAndSale) {
		return nil, false
	}
	return o.NetInvestmentPropertiesPurchaseAndSale, true
}

// HasNetInvestmentPropertiesPurchaseAndSale returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasNetInvestmentPropertiesPurchaseAndSale() bool {
	if o != nil && !IsNil(o.NetInvestmentPropertiesPurchaseAndSale) {
		return true
	}

	return false
}

// SetNetInvestmentPropertiesPurchaseAndSale gets a reference to the given float64 and assigns it to the NetInvestmentPropertiesPurchaseAndSale field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetNetInvestmentPropertiesPurchaseAndSale(v float64) {
	o.NetInvestmentPropertiesPurchaseAndSale = &v
}

// GetSaleOfInvestmentProperties returns the SaleOfInvestmentProperties field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetSaleOfInvestmentProperties() float64 {
	if o == nil || IsNil(o.SaleOfInvestmentProperties) {
		var ret float64
		return ret
	}
	return *o.SaleOfInvestmentProperties
}

// GetSaleOfInvestmentPropertiesOk returns a tuple with the SaleOfInvestmentProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetSaleOfInvestmentPropertiesOk() (*float64, bool) {
	if o == nil || IsNil(o.SaleOfInvestmentProperties) {
		return nil, false
	}
	return o.SaleOfInvestmentProperties, true
}

// HasSaleOfInvestmentProperties returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasSaleOfInvestmentProperties() bool {
	if o != nil && !IsNil(o.SaleOfInvestmentProperties) {
		return true
	}

	return false
}

// SetSaleOfInvestmentProperties gets a reference to the given float64 and assigns it to the SaleOfInvestmentProperties field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetSaleOfInvestmentProperties(v float64) {
	o.SaleOfInvestmentProperties = &v
}

// GetPurchaseOfInvestmentProperties returns the PurchaseOfInvestmentProperties field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetPurchaseOfInvestmentProperties() float64 {
	if o == nil || IsNil(o.PurchaseOfInvestmentProperties) {
		var ret float64
		return ret
	}
	return *o.PurchaseOfInvestmentProperties
}

// GetPurchaseOfInvestmentPropertiesOk returns a tuple with the PurchaseOfInvestmentProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetPurchaseOfInvestmentPropertiesOk() (*float64, bool) {
	if o == nil || IsNil(o.PurchaseOfInvestmentProperties) {
		return nil, false
	}
	return o.PurchaseOfInvestmentProperties, true
}

// HasPurchaseOfInvestmentProperties returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasPurchaseOfInvestmentProperties() bool {
	if o != nil && !IsNil(o.PurchaseOfInvestmentProperties) {
		return true
	}

	return false
}

// SetPurchaseOfInvestmentProperties gets a reference to the given float64 and assigns it to the PurchaseOfInvestmentProperties field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetPurchaseOfInvestmentProperties(v float64) {
	o.PurchaseOfInvestmentProperties = &v
}

// GetNetBusinessPurchaseAndSale returns the NetBusinessPurchaseAndSale field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetNetBusinessPurchaseAndSale() float64 {
	if o == nil || IsNil(o.NetBusinessPurchaseAndSale) {
		var ret float64
		return ret
	}
	return *o.NetBusinessPurchaseAndSale
}

// GetNetBusinessPurchaseAndSaleOk returns a tuple with the NetBusinessPurchaseAndSale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetNetBusinessPurchaseAndSaleOk() (*float64, bool) {
	if o == nil || IsNil(o.NetBusinessPurchaseAndSale) {
		return nil, false
	}
	return o.NetBusinessPurchaseAndSale, true
}

// HasNetBusinessPurchaseAndSale returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasNetBusinessPurchaseAndSale() bool {
	if o != nil && !IsNil(o.NetBusinessPurchaseAndSale) {
		return true
	}

	return false
}

// SetNetBusinessPurchaseAndSale gets a reference to the given float64 and assigns it to the NetBusinessPurchaseAndSale field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetNetBusinessPurchaseAndSale(v float64) {
	o.NetBusinessPurchaseAndSale = &v
}

// GetSaleOfBusiness returns the SaleOfBusiness field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetSaleOfBusiness() float64 {
	if o == nil || IsNil(o.SaleOfBusiness) {
		var ret float64
		return ret
	}
	return *o.SaleOfBusiness
}

// GetSaleOfBusinessOk returns a tuple with the SaleOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetSaleOfBusinessOk() (*float64, bool) {
	if o == nil || IsNil(o.SaleOfBusiness) {
		return nil, false
	}
	return o.SaleOfBusiness, true
}

// HasSaleOfBusiness returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasSaleOfBusiness() bool {
	if o != nil && !IsNil(o.SaleOfBusiness) {
		return true
	}

	return false
}

// SetSaleOfBusiness gets a reference to the given float64 and assigns it to the SaleOfBusiness field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetSaleOfBusiness(v float64) {
	o.SaleOfBusiness = &v
}

// GetPurchaseOfBusiness returns the PurchaseOfBusiness field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetPurchaseOfBusiness() float64 {
	if o == nil || IsNil(o.PurchaseOfBusiness) {
		var ret float64
		return ret
	}
	return *o.PurchaseOfBusiness
}

// GetPurchaseOfBusinessOk returns a tuple with the PurchaseOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetPurchaseOfBusinessOk() (*float64, bool) {
	if o == nil || IsNil(o.PurchaseOfBusiness) {
		return nil, false
	}
	return o.PurchaseOfBusiness, true
}

// HasPurchaseOfBusiness returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasPurchaseOfBusiness() bool {
	if o != nil && !IsNil(o.PurchaseOfBusiness) {
		return true
	}

	return false
}

// SetPurchaseOfBusiness gets a reference to the given float64 and assigns it to the PurchaseOfBusiness field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetPurchaseOfBusiness(v float64) {
	o.PurchaseOfBusiness = &v
}

// GetNetIntangiblesPurchaseAndSale returns the NetIntangiblesPurchaseAndSale field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetNetIntangiblesPurchaseAndSale() float64 {
	if o == nil || IsNil(o.NetIntangiblesPurchaseAndSale) {
		var ret float64
		return ret
	}
	return *o.NetIntangiblesPurchaseAndSale
}

// GetNetIntangiblesPurchaseAndSaleOk returns a tuple with the NetIntangiblesPurchaseAndSale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetNetIntangiblesPurchaseAndSaleOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIntangiblesPurchaseAndSale) {
		return nil, false
	}
	return o.NetIntangiblesPurchaseAndSale, true
}

// HasNetIntangiblesPurchaseAndSale returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasNetIntangiblesPurchaseAndSale() bool {
	if o != nil && !IsNil(o.NetIntangiblesPurchaseAndSale) {
		return true
	}

	return false
}

// SetNetIntangiblesPurchaseAndSale gets a reference to the given float64 and assigns it to the NetIntangiblesPurchaseAndSale field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetNetIntangiblesPurchaseAndSale(v float64) {
	o.NetIntangiblesPurchaseAndSale = &v
}

// GetSaleOfIntangibles returns the SaleOfIntangibles field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetSaleOfIntangibles() float64 {
	if o == nil || IsNil(o.SaleOfIntangibles) {
		var ret float64
		return ret
	}
	return *o.SaleOfIntangibles
}

// GetSaleOfIntangiblesOk returns a tuple with the SaleOfIntangibles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetSaleOfIntangiblesOk() (*float64, bool) {
	if o == nil || IsNil(o.SaleOfIntangibles) {
		return nil, false
	}
	return o.SaleOfIntangibles, true
}

// HasSaleOfIntangibles returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasSaleOfIntangibles() bool {
	if o != nil && !IsNil(o.SaleOfIntangibles) {
		return true
	}

	return false
}

// SetSaleOfIntangibles gets a reference to the given float64 and assigns it to the SaleOfIntangibles field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetSaleOfIntangibles(v float64) {
	o.SaleOfIntangibles = &v
}

// GetPurchaseOfIntangibles returns the PurchaseOfIntangibles field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetPurchaseOfIntangibles() float64 {
	if o == nil || IsNil(o.PurchaseOfIntangibles) {
		var ret float64
		return ret
	}
	return *o.PurchaseOfIntangibles
}

// GetPurchaseOfIntangiblesOk returns a tuple with the PurchaseOfIntangibles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetPurchaseOfIntangiblesOk() (*float64, bool) {
	if o == nil || IsNil(o.PurchaseOfIntangibles) {
		return nil, false
	}
	return o.PurchaseOfIntangibles, true
}

// HasPurchaseOfIntangibles returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasPurchaseOfIntangibles() bool {
	if o != nil && !IsNil(o.PurchaseOfIntangibles) {
		return true
	}

	return false
}

// SetPurchaseOfIntangibles gets a reference to the given float64 and assigns it to the PurchaseOfIntangibles field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetPurchaseOfIntangibles(v float64) {
	o.PurchaseOfIntangibles = &v
}

// GetNetPpePurchaseAndSale returns the NetPpePurchaseAndSale field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetNetPpePurchaseAndSale() float64 {
	if o == nil || IsNil(o.NetPpePurchaseAndSale) {
		var ret float64
		return ret
	}
	return *o.NetPpePurchaseAndSale
}

// GetNetPpePurchaseAndSaleOk returns a tuple with the NetPpePurchaseAndSale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetNetPpePurchaseAndSaleOk() (*float64, bool) {
	if o == nil || IsNil(o.NetPpePurchaseAndSale) {
		return nil, false
	}
	return o.NetPpePurchaseAndSale, true
}

// HasNetPpePurchaseAndSale returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasNetPpePurchaseAndSale() bool {
	if o != nil && !IsNil(o.NetPpePurchaseAndSale) {
		return true
	}

	return false
}

// SetNetPpePurchaseAndSale gets a reference to the given float64 and assigns it to the NetPpePurchaseAndSale field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetNetPpePurchaseAndSale(v float64) {
	o.NetPpePurchaseAndSale = &v
}

// GetSaleOfPpe returns the SaleOfPpe field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetSaleOfPpe() float64 {
	if o == nil || IsNil(o.SaleOfPpe) {
		var ret float64
		return ret
	}
	return *o.SaleOfPpe
}

// GetSaleOfPpeOk returns a tuple with the SaleOfPpe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetSaleOfPpeOk() (*float64, bool) {
	if o == nil || IsNil(o.SaleOfPpe) {
		return nil, false
	}
	return o.SaleOfPpe, true
}

// HasSaleOfPpe returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasSaleOfPpe() bool {
	if o != nil && !IsNil(o.SaleOfPpe) {
		return true
	}

	return false
}

// SetSaleOfPpe gets a reference to the given float64 and assigns it to the SaleOfPpe field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetSaleOfPpe(v float64) {
	o.SaleOfPpe = &v
}

// GetPurchaseOfPpe returns the PurchaseOfPpe field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetPurchaseOfPpe() float64 {
	if o == nil || IsNil(o.PurchaseOfPpe) {
		var ret float64
		return ret
	}
	return *o.PurchaseOfPpe
}

// GetPurchaseOfPpeOk returns a tuple with the PurchaseOfPpe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetPurchaseOfPpeOk() (*float64, bool) {
	if o == nil || IsNil(o.PurchaseOfPpe) {
		return nil, false
	}
	return o.PurchaseOfPpe, true
}

// HasPurchaseOfPpe returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasPurchaseOfPpe() bool {
	if o != nil && !IsNil(o.PurchaseOfPpe) {
		return true
	}

	return false
}

// SetPurchaseOfPpe gets a reference to the given float64 and assigns it to the PurchaseOfPpe field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetPurchaseOfPpe(v float64) {
	o.PurchaseOfPpe = &v
}

// GetCapitalExpenditureReported returns the CapitalExpenditureReported field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetCapitalExpenditureReported() float64 {
	if o == nil || IsNil(o.CapitalExpenditureReported) {
		var ret float64
		return ret
	}
	return *o.CapitalExpenditureReported
}

// GetCapitalExpenditureReportedOk returns a tuple with the CapitalExpenditureReported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetCapitalExpenditureReportedOk() (*float64, bool) {
	if o == nil || IsNil(o.CapitalExpenditureReported) {
		return nil, false
	}
	return o.CapitalExpenditureReported, true
}

// HasCapitalExpenditureReported returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasCapitalExpenditureReported() bool {
	if o != nil && !IsNil(o.CapitalExpenditureReported) {
		return true
	}

	return false
}

// SetCapitalExpenditureReported gets a reference to the given float64 and assigns it to the CapitalExpenditureReported field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetCapitalExpenditureReported(v float64) {
	o.CapitalExpenditureReported = &v
}

// GetCapitalExpenditure returns the CapitalExpenditure field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetCapitalExpenditure() float64 {
	if o == nil || IsNil(o.CapitalExpenditure) {
		var ret float64
		return ret
	}
	return *o.CapitalExpenditure
}

// GetCapitalExpenditureOk returns a tuple with the CapitalExpenditure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) GetCapitalExpenditureOk() (*float64, bool) {
	if o == nil || IsNil(o.CapitalExpenditure) {
		return nil, false
	}
	return o.CapitalExpenditure, true
}

// HasCapitalExpenditure returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromInvestingActivities) HasCapitalExpenditure() bool {
	if o != nil && !IsNil(o.CapitalExpenditure) {
		return true
	}

	return false
}

// SetCapitalExpenditure gets a reference to the given float64 and assigns it to the CapitalExpenditure field.
func (o *CashFlowDataCashFlowFromInvestingActivities) SetCapitalExpenditure(v float64) {
	o.CapitalExpenditure = &v
}

func (o CashFlowDataCashFlowFromInvestingActivities) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashFlowDataCashFlowFromInvestingActivities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvestingCashFlow) {
		toSerialize["investing_cash_flow"] = o.InvestingCashFlow
	}
	if !IsNil(o.CashFlowFromContinuingInvestingActivities) {
		toSerialize["cash_flow_from_continuing_investing_activities"] = o.CashFlowFromContinuingInvestingActivities
	}
	if !IsNil(o.CashFromDiscontinuedInvestingActivities) {
		toSerialize["cash_from_discontinued_investing_activities"] = o.CashFromDiscontinuedInvestingActivities
	}
	if !IsNil(o.NetOtherInvestingChanges) {
		toSerialize["net_other_investing_changes"] = o.NetOtherInvestingChanges
	}
	if !IsNil(o.InterestReceivedCfi) {
		toSerialize["interest_received_cfi"] = o.InterestReceivedCfi
	}
	if !IsNil(o.DividendsReceivedCfi) {
		toSerialize["dividends_received_cfi"] = o.DividendsReceivedCfi
	}
	if !IsNil(o.NetInvestmentPurchaseAndSale) {
		toSerialize["net_investment_purchase_and_sale"] = o.NetInvestmentPurchaseAndSale
	}
	if !IsNil(o.SaleOfInvestment) {
		toSerialize["sale_of_investment"] = o.SaleOfInvestment
	}
	if !IsNil(o.PurchaseOfInvestment) {
		toSerialize["purchase_of_investment"] = o.PurchaseOfInvestment
	}
	if !IsNil(o.NetInvestmentPropertiesPurchaseAndSale) {
		toSerialize["net_investment_properties_purchase_and_sale"] = o.NetInvestmentPropertiesPurchaseAndSale
	}
	if !IsNil(o.SaleOfInvestmentProperties) {
		toSerialize["sale_of_investment_properties"] = o.SaleOfInvestmentProperties
	}
	if !IsNil(o.PurchaseOfInvestmentProperties) {
		toSerialize["purchase_of_investment_properties"] = o.PurchaseOfInvestmentProperties
	}
	if !IsNil(o.NetBusinessPurchaseAndSale) {
		toSerialize["net_business_purchase_and_sale"] = o.NetBusinessPurchaseAndSale
	}
	if !IsNil(o.SaleOfBusiness) {
		toSerialize["sale_of_business"] = o.SaleOfBusiness
	}
	if !IsNil(o.PurchaseOfBusiness) {
		toSerialize["purchase_of_business"] = o.PurchaseOfBusiness
	}
	if !IsNil(o.NetIntangiblesPurchaseAndSale) {
		toSerialize["net_intangibles_purchase_and_sale"] = o.NetIntangiblesPurchaseAndSale
	}
	if !IsNil(o.SaleOfIntangibles) {
		toSerialize["sale_of_intangibles"] = o.SaleOfIntangibles
	}
	if !IsNil(o.PurchaseOfIntangibles) {
		toSerialize["purchase_of_intangibles"] = o.PurchaseOfIntangibles
	}
	if !IsNil(o.NetPpePurchaseAndSale) {
		toSerialize["net_ppe_purchase_and_sale"] = o.NetPpePurchaseAndSale
	}
	if !IsNil(o.SaleOfPpe) {
		toSerialize["sale_of_ppe"] = o.SaleOfPpe
	}
	if !IsNil(o.PurchaseOfPpe) {
		toSerialize["purchase_of_ppe"] = o.PurchaseOfPpe
	}
	if !IsNil(o.CapitalExpenditureReported) {
		toSerialize["capital_expenditure_reported"] = o.CapitalExpenditureReported
	}
	if !IsNil(o.CapitalExpenditure) {
		toSerialize["capital_expenditure"] = o.CapitalExpenditure
	}
	return toSerialize, nil
}

type NullableCashFlowDataCashFlowFromInvestingActivities struct {
	value *CashFlowDataCashFlowFromInvestingActivities
	isSet bool
}

func (v NullableCashFlowDataCashFlowFromInvestingActivities) Get() *CashFlowDataCashFlowFromInvestingActivities {
	return v.value
}

func (v *NullableCashFlowDataCashFlowFromInvestingActivities) Set(val *CashFlowDataCashFlowFromInvestingActivities) {
	v.value = val
	v.isSet = true
}

func (v NullableCashFlowDataCashFlowFromInvestingActivities) IsSet() bool {
	return v.isSet
}

func (v *NullableCashFlowDataCashFlowFromInvestingActivities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashFlowDataCashFlowFromInvestingActivities(val *CashFlowDataCashFlowFromInvestingActivities) *NullableCashFlowDataCashFlowFromInvestingActivities {
	return &NullableCashFlowDataCashFlowFromInvestingActivities{value: val, isSet: true}
}

func (v NullableCashFlowDataCashFlowFromInvestingActivities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashFlowDataCashFlowFromInvestingActivities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
