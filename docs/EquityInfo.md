# EquityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalEquityGrossMinorityInterest** | Pointer to **float64** | Total equity gross minority interest | [optional] 
**StockholdersEquity** | Pointer to **float64** | Stockholders equity | [optional] 
**CommonStockEquity** | Pointer to **float64** | Common stock equity | [optional] 
**PreferredStockEquity** | Pointer to **float64** | Preferred stock equity | [optional] 
**OtherEquityInterest** | Pointer to **float64** | Other equity interest | [optional] 
**MinorityInterest** | Pointer to **float64** | Minority interest | [optional] 
**TotalCapitalization** | Pointer to **float64** | Total capitalization | [optional] 
**NetTangibleAssets** | Pointer to **float64** | Net tangible assets | [optional] 
**TangibleBookValue** | Pointer to **float64** | Tangible book value | [optional] 
**InvestedCapital** | Pointer to **float64** | Invested capital | [optional] 
**WorkingCapital** | Pointer to **float64** | Working capital | [optional] 
**CapitalStock** | Pointer to [**EquityInfoCapitalStock**](EquityInfoCapitalStock.md) |  | [optional] 
**EquityAdjustments** | Pointer to [**EquityInfoEquityAdjustments**](EquityInfoEquityAdjustments.md) |  | [optional] 
**NetDebt** | Pointer to **float64** | Net debt | [optional] 
**TotalDebt** | Pointer to **float64** | Total debt | [optional] 

## Methods

### NewEquityInfo

`func NewEquityInfo() *EquityInfo`

NewEquityInfo instantiates a new EquityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityInfoWithDefaults

`func NewEquityInfoWithDefaults() *EquityInfo`

NewEquityInfoWithDefaults instantiates a new EquityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalEquityGrossMinorityInterest

`func (o *EquityInfo) GetTotalEquityGrossMinorityInterest() float64`

GetTotalEquityGrossMinorityInterest returns the TotalEquityGrossMinorityInterest field if non-nil, zero value otherwise.

### GetTotalEquityGrossMinorityInterestOk

`func (o *EquityInfo) GetTotalEquityGrossMinorityInterestOk() (*float64, bool)`

GetTotalEquityGrossMinorityInterestOk returns a tuple with the TotalEquityGrossMinorityInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEquityGrossMinorityInterest

`func (o *EquityInfo) SetTotalEquityGrossMinorityInterest(v float64)`

SetTotalEquityGrossMinorityInterest sets TotalEquityGrossMinorityInterest field to given value.

### HasTotalEquityGrossMinorityInterest

`func (o *EquityInfo) HasTotalEquityGrossMinorityInterest() bool`

HasTotalEquityGrossMinorityInterest returns a boolean if a field has been set.

### GetStockholdersEquity

`func (o *EquityInfo) GetStockholdersEquity() float64`

GetStockholdersEquity returns the StockholdersEquity field if non-nil, zero value otherwise.

### GetStockholdersEquityOk

`func (o *EquityInfo) GetStockholdersEquityOk() (*float64, bool)`

GetStockholdersEquityOk returns a tuple with the StockholdersEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockholdersEquity

`func (o *EquityInfo) SetStockholdersEquity(v float64)`

SetStockholdersEquity sets StockholdersEquity field to given value.

### HasStockholdersEquity

`func (o *EquityInfo) HasStockholdersEquity() bool`

HasStockholdersEquity returns a boolean if a field has been set.

### GetCommonStockEquity

`func (o *EquityInfo) GetCommonStockEquity() float64`

GetCommonStockEquity returns the CommonStockEquity field if non-nil, zero value otherwise.

### GetCommonStockEquityOk

`func (o *EquityInfo) GetCommonStockEquityOk() (*float64, bool)`

GetCommonStockEquityOk returns a tuple with the CommonStockEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonStockEquity

`func (o *EquityInfo) SetCommonStockEquity(v float64)`

SetCommonStockEquity sets CommonStockEquity field to given value.

### HasCommonStockEquity

`func (o *EquityInfo) HasCommonStockEquity() bool`

HasCommonStockEquity returns a boolean if a field has been set.

### GetPreferredStockEquity

`func (o *EquityInfo) GetPreferredStockEquity() float64`

GetPreferredStockEquity returns the PreferredStockEquity field if non-nil, zero value otherwise.

### GetPreferredStockEquityOk

`func (o *EquityInfo) GetPreferredStockEquityOk() (*float64, bool)`

GetPreferredStockEquityOk returns a tuple with the PreferredStockEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredStockEquity

`func (o *EquityInfo) SetPreferredStockEquity(v float64)`

SetPreferredStockEquity sets PreferredStockEquity field to given value.

### HasPreferredStockEquity

`func (o *EquityInfo) HasPreferredStockEquity() bool`

HasPreferredStockEquity returns a boolean if a field has been set.

### GetOtherEquityInterest

`func (o *EquityInfo) GetOtherEquityInterest() float64`

GetOtherEquityInterest returns the OtherEquityInterest field if non-nil, zero value otherwise.

### GetOtherEquityInterestOk

`func (o *EquityInfo) GetOtherEquityInterestOk() (*float64, bool)`

GetOtherEquityInterestOk returns a tuple with the OtherEquityInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherEquityInterest

`func (o *EquityInfo) SetOtherEquityInterest(v float64)`

SetOtherEquityInterest sets OtherEquityInterest field to given value.

### HasOtherEquityInterest

`func (o *EquityInfo) HasOtherEquityInterest() bool`

HasOtherEquityInterest returns a boolean if a field has been set.

### GetMinorityInterest

`func (o *EquityInfo) GetMinorityInterest() float64`

GetMinorityInterest returns the MinorityInterest field if non-nil, zero value otherwise.

### GetMinorityInterestOk

`func (o *EquityInfo) GetMinorityInterestOk() (*float64, bool)`

GetMinorityInterestOk returns a tuple with the MinorityInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorityInterest

`func (o *EquityInfo) SetMinorityInterest(v float64)`

SetMinorityInterest sets MinorityInterest field to given value.

### HasMinorityInterest

`func (o *EquityInfo) HasMinorityInterest() bool`

HasMinorityInterest returns a boolean if a field has been set.

### GetTotalCapitalization

`func (o *EquityInfo) GetTotalCapitalization() float64`

GetTotalCapitalization returns the TotalCapitalization field if non-nil, zero value otherwise.

### GetTotalCapitalizationOk

`func (o *EquityInfo) GetTotalCapitalizationOk() (*float64, bool)`

GetTotalCapitalizationOk returns a tuple with the TotalCapitalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapitalization

`func (o *EquityInfo) SetTotalCapitalization(v float64)`

SetTotalCapitalization sets TotalCapitalization field to given value.

### HasTotalCapitalization

`func (o *EquityInfo) HasTotalCapitalization() bool`

HasTotalCapitalization returns a boolean if a field has been set.

### GetNetTangibleAssets

`func (o *EquityInfo) GetNetTangibleAssets() float64`

GetNetTangibleAssets returns the NetTangibleAssets field if non-nil, zero value otherwise.

### GetNetTangibleAssetsOk

`func (o *EquityInfo) GetNetTangibleAssetsOk() (*float64, bool)`

GetNetTangibleAssetsOk returns a tuple with the NetTangibleAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTangibleAssets

`func (o *EquityInfo) SetNetTangibleAssets(v float64)`

SetNetTangibleAssets sets NetTangibleAssets field to given value.

### HasNetTangibleAssets

`func (o *EquityInfo) HasNetTangibleAssets() bool`

HasNetTangibleAssets returns a boolean if a field has been set.

### GetTangibleBookValue

`func (o *EquityInfo) GetTangibleBookValue() float64`

GetTangibleBookValue returns the TangibleBookValue field if non-nil, zero value otherwise.

### GetTangibleBookValueOk

`func (o *EquityInfo) GetTangibleBookValueOk() (*float64, bool)`

GetTangibleBookValueOk returns a tuple with the TangibleBookValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTangibleBookValue

`func (o *EquityInfo) SetTangibleBookValue(v float64)`

SetTangibleBookValue sets TangibleBookValue field to given value.

### HasTangibleBookValue

`func (o *EquityInfo) HasTangibleBookValue() bool`

HasTangibleBookValue returns a boolean if a field has been set.

### GetInvestedCapital

`func (o *EquityInfo) GetInvestedCapital() float64`

GetInvestedCapital returns the InvestedCapital field if non-nil, zero value otherwise.

### GetInvestedCapitalOk

`func (o *EquityInfo) GetInvestedCapitalOk() (*float64, bool)`

GetInvestedCapitalOk returns a tuple with the InvestedCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestedCapital

`func (o *EquityInfo) SetInvestedCapital(v float64)`

SetInvestedCapital sets InvestedCapital field to given value.

### HasInvestedCapital

`func (o *EquityInfo) HasInvestedCapital() bool`

HasInvestedCapital returns a boolean if a field has been set.

### GetWorkingCapital

`func (o *EquityInfo) GetWorkingCapital() float64`

GetWorkingCapital returns the WorkingCapital field if non-nil, zero value otherwise.

### GetWorkingCapitalOk

`func (o *EquityInfo) GetWorkingCapitalOk() (*float64, bool)`

GetWorkingCapitalOk returns a tuple with the WorkingCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingCapital

`func (o *EquityInfo) SetWorkingCapital(v float64)`

SetWorkingCapital sets WorkingCapital field to given value.

### HasWorkingCapital

`func (o *EquityInfo) HasWorkingCapital() bool`

HasWorkingCapital returns a boolean if a field has been set.

### GetCapitalStock

`func (o *EquityInfo) GetCapitalStock() EquityInfoCapitalStock`

GetCapitalStock returns the CapitalStock field if non-nil, zero value otherwise.

### GetCapitalStockOk

`func (o *EquityInfo) GetCapitalStockOk() (*EquityInfoCapitalStock, bool)`

GetCapitalStockOk returns a tuple with the CapitalStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalStock

`func (o *EquityInfo) SetCapitalStock(v EquityInfoCapitalStock)`

SetCapitalStock sets CapitalStock field to given value.

### HasCapitalStock

`func (o *EquityInfo) HasCapitalStock() bool`

HasCapitalStock returns a boolean if a field has been set.

### GetEquityAdjustments

`func (o *EquityInfo) GetEquityAdjustments() EquityInfoEquityAdjustments`

GetEquityAdjustments returns the EquityAdjustments field if non-nil, zero value otherwise.

### GetEquityAdjustmentsOk

`func (o *EquityInfo) GetEquityAdjustmentsOk() (*EquityInfoEquityAdjustments, bool)`

GetEquityAdjustmentsOk returns a tuple with the EquityAdjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityAdjustments

`func (o *EquityInfo) SetEquityAdjustments(v EquityInfoEquityAdjustments)`

SetEquityAdjustments sets EquityAdjustments field to given value.

### HasEquityAdjustments

`func (o *EquityInfo) HasEquityAdjustments() bool`

HasEquityAdjustments returns a boolean if a field has been set.

### GetNetDebt

`func (o *EquityInfo) GetNetDebt() float64`

GetNetDebt returns the NetDebt field if non-nil, zero value otherwise.

### GetNetDebtOk

`func (o *EquityInfo) GetNetDebtOk() (*float64, bool)`

GetNetDebtOk returns a tuple with the NetDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetDebt

`func (o *EquityInfo) SetNetDebt(v float64)`

SetNetDebt sets NetDebt field to given value.

### HasNetDebt

`func (o *EquityInfo) HasNetDebt() bool`

HasNetDebt returns a boolean if a field has been set.

### GetTotalDebt

`func (o *EquityInfo) GetTotalDebt() float64`

GetTotalDebt returns the TotalDebt field if non-nil, zero value otherwise.

### GetTotalDebtOk

`func (o *EquityInfo) GetTotalDebtOk() (*float64, bool)`

GetTotalDebtOk returns a tuple with the TotalDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebt

`func (o *EquityInfo) SetTotalDebt(v float64)`

SetTotalDebt sets TotalDebt field to given value.

### HasTotalDebt

`func (o *EquityInfo) HasTotalDebt() bool`

HasTotalDebt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


