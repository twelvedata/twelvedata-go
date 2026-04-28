# GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonStock** | Pointer to **float64** | Represents net worth of investors shares, which is equal to the total_assets - total_liabilities | [optional] 
**RetainedEarnings** | Pointer to **float64** | Refers to the profits earned minus dividends paid | [optional] 
**OtherShareholdersEquity** | Pointer to **float64** | Represents other not affecting retained earnings gains and looses | [optional] 
**TotalShareholdersEquity** | Pointer to **float64** | Represents the net worth of a company, which is the amount that would be returned to shareholders if a company&#39;s total assets were liquidated, and all of its debts were repaid | [optional] 
**AdditionalPaidInCapital** | Pointer to **float64** | Represents the additional paid-in capital, which is the amount shareholders have invested in a company above the par value of its stock | [optional] 
**TreasuryStock** | Pointer to **float64** | Represents the value of shares that have been repurchased by the company and are held in its treasury | [optional] 
**MinorityInterest** | Pointer to **float64** | Represents the portion of shareholders&#39; equity that is attributable to minority shareholders in a subsidiary company | [optional] 

## Methods

### NewGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity

`func NewGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity() *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity`

NewGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquityWithDefaults

`func NewGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquityWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity`

NewGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquityWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonStock

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetCommonStock() float64`

GetCommonStock returns the CommonStock field if non-nil, zero value otherwise.

### GetCommonStockOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetCommonStockOk() (*float64, bool)`

GetCommonStockOk returns a tuple with the CommonStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonStock

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetCommonStock(v float64)`

SetCommonStock sets CommonStock field to given value.

### HasCommonStock

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasCommonStock() bool`

HasCommonStock returns a boolean if a field has been set.

### GetRetainedEarnings

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetRetainedEarnings() float64`

GetRetainedEarnings returns the RetainedEarnings field if non-nil, zero value otherwise.

### GetRetainedEarningsOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetRetainedEarningsOk() (*float64, bool)`

GetRetainedEarningsOk returns a tuple with the RetainedEarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedEarnings

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetRetainedEarnings(v float64)`

SetRetainedEarnings sets RetainedEarnings field to given value.

### HasRetainedEarnings

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasRetainedEarnings() bool`

HasRetainedEarnings returns a boolean if a field has been set.

### GetOtherShareholdersEquity

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetOtherShareholdersEquity() float64`

GetOtherShareholdersEquity returns the OtherShareholdersEquity field if non-nil, zero value otherwise.

### GetOtherShareholdersEquityOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetOtherShareholdersEquityOk() (*float64, bool)`

GetOtherShareholdersEquityOk returns a tuple with the OtherShareholdersEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherShareholdersEquity

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetOtherShareholdersEquity(v float64)`

SetOtherShareholdersEquity sets OtherShareholdersEquity field to given value.

### HasOtherShareholdersEquity

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasOtherShareholdersEquity() bool`

HasOtherShareholdersEquity returns a boolean if a field has been set.

### GetTotalShareholdersEquity

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetTotalShareholdersEquity() float64`

GetTotalShareholdersEquity returns the TotalShareholdersEquity field if non-nil, zero value otherwise.

### GetTotalShareholdersEquityOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetTotalShareholdersEquityOk() (*float64, bool)`

GetTotalShareholdersEquityOk returns a tuple with the TotalShareholdersEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShareholdersEquity

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetTotalShareholdersEquity(v float64)`

SetTotalShareholdersEquity sets TotalShareholdersEquity field to given value.

### HasTotalShareholdersEquity

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasTotalShareholdersEquity() bool`

HasTotalShareholdersEquity returns a boolean if a field has been set.

### GetAdditionalPaidInCapital

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetAdditionalPaidInCapital() float64`

GetAdditionalPaidInCapital returns the AdditionalPaidInCapital field if non-nil, zero value otherwise.

### GetAdditionalPaidInCapitalOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetAdditionalPaidInCapitalOk() (*float64, bool)`

GetAdditionalPaidInCapitalOk returns a tuple with the AdditionalPaidInCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPaidInCapital

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetAdditionalPaidInCapital(v float64)`

SetAdditionalPaidInCapital sets AdditionalPaidInCapital field to given value.

### HasAdditionalPaidInCapital

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasAdditionalPaidInCapital() bool`

HasAdditionalPaidInCapital returns a boolean if a field has been set.

### GetTreasuryStock

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetTreasuryStock() float64`

GetTreasuryStock returns the TreasuryStock field if non-nil, zero value otherwise.

### GetTreasuryStockOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetTreasuryStockOk() (*float64, bool)`

GetTreasuryStockOk returns a tuple with the TreasuryStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreasuryStock

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetTreasuryStock(v float64)`

SetTreasuryStock sets TreasuryStock field to given value.

### HasTreasuryStock

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasTreasuryStock() bool`

HasTreasuryStock returns a boolean if a field has been set.

### GetMinorityInterest

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetMinorityInterest() float64`

GetMinorityInterest returns the MinorityInterest field if non-nil, zero value otherwise.

### GetMinorityInterestOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetMinorityInterestOk() (*float64, bool)`

GetMinorityInterestOk returns a tuple with the MinorityInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorityInterest

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetMinorityInterest(v float64)`

SetMinorityInterest sets MinorityInterest field to given value.

### HasMinorityInterest

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasMinorityInterest() bool`

HasMinorityInterest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


