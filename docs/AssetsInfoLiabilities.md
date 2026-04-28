# AssetsInfoLiabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalLiabilitiesNetMinorityInterest** | Pointer to **float64** | Total liabilities net minority interest | [optional] 
**CurrentLiabilities** | Pointer to [**AssetsInfoLiabilitiesCurrentLiabilities**](AssetsInfoLiabilitiesCurrentLiabilities.md) |  | [optional] 
**NonCurrentLiabilities** | Pointer to [**AssetsInfoLiabilitiesNonCurrentLiabilities**](AssetsInfoLiabilitiesNonCurrentLiabilities.md) |  | [optional] 
**Equity** | Pointer to [**EquityInfo**](EquityInfo.md) |  | [optional] 

## Methods

### NewAssetsInfoLiabilities

`func NewAssetsInfoLiabilities() *AssetsInfoLiabilities`

NewAssetsInfoLiabilities instantiates a new AssetsInfoLiabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsInfoLiabilitiesWithDefaults

`func NewAssetsInfoLiabilitiesWithDefaults() *AssetsInfoLiabilities`

NewAssetsInfoLiabilitiesWithDefaults instantiates a new AssetsInfoLiabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalLiabilitiesNetMinorityInterest

`func (o *AssetsInfoLiabilities) GetTotalLiabilitiesNetMinorityInterest() float64`

GetTotalLiabilitiesNetMinorityInterest returns the TotalLiabilitiesNetMinorityInterest field if non-nil, zero value otherwise.

### GetTotalLiabilitiesNetMinorityInterestOk

`func (o *AssetsInfoLiabilities) GetTotalLiabilitiesNetMinorityInterestOk() (*float64, bool)`

GetTotalLiabilitiesNetMinorityInterestOk returns a tuple with the TotalLiabilitiesNetMinorityInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilitiesNetMinorityInterest

`func (o *AssetsInfoLiabilities) SetTotalLiabilitiesNetMinorityInterest(v float64)`

SetTotalLiabilitiesNetMinorityInterest sets TotalLiabilitiesNetMinorityInterest field to given value.

### HasTotalLiabilitiesNetMinorityInterest

`func (o *AssetsInfoLiabilities) HasTotalLiabilitiesNetMinorityInterest() bool`

HasTotalLiabilitiesNetMinorityInterest returns a boolean if a field has been set.

### GetCurrentLiabilities

`func (o *AssetsInfoLiabilities) GetCurrentLiabilities() AssetsInfoLiabilitiesCurrentLiabilities`

GetCurrentLiabilities returns the CurrentLiabilities field if non-nil, zero value otherwise.

### GetCurrentLiabilitiesOk

`func (o *AssetsInfoLiabilities) GetCurrentLiabilitiesOk() (*AssetsInfoLiabilitiesCurrentLiabilities, bool)`

GetCurrentLiabilitiesOk returns a tuple with the CurrentLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLiabilities

`func (o *AssetsInfoLiabilities) SetCurrentLiabilities(v AssetsInfoLiabilitiesCurrentLiabilities)`

SetCurrentLiabilities sets CurrentLiabilities field to given value.

### HasCurrentLiabilities

`func (o *AssetsInfoLiabilities) HasCurrentLiabilities() bool`

HasCurrentLiabilities returns a boolean if a field has been set.

### GetNonCurrentLiabilities

`func (o *AssetsInfoLiabilities) GetNonCurrentLiabilities() AssetsInfoLiabilitiesNonCurrentLiabilities`

GetNonCurrentLiabilities returns the NonCurrentLiabilities field if non-nil, zero value otherwise.

### GetNonCurrentLiabilitiesOk

`func (o *AssetsInfoLiabilities) GetNonCurrentLiabilitiesOk() (*AssetsInfoLiabilitiesNonCurrentLiabilities, bool)`

GetNonCurrentLiabilitiesOk returns a tuple with the NonCurrentLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCurrentLiabilities

`func (o *AssetsInfoLiabilities) SetNonCurrentLiabilities(v AssetsInfoLiabilitiesNonCurrentLiabilities)`

SetNonCurrentLiabilities sets NonCurrentLiabilities field to given value.

### HasNonCurrentLiabilities

`func (o *AssetsInfoLiabilities) HasNonCurrentLiabilities() bool`

HasNonCurrentLiabilities returns a boolean if a field has been set.

### GetEquity

`func (o *AssetsInfoLiabilities) GetEquity() EquityInfo`

GetEquity returns the Equity field if non-nil, zero value otherwise.

### GetEquityOk

`func (o *AssetsInfoLiabilities) GetEquityOk() (*EquityInfo, bool)`

GetEquityOk returns a tuple with the Equity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquity

`func (o *AssetsInfoLiabilities) SetEquity(v EquityInfo)`

SetEquity sets Equity field to given value.

### HasEquity

`func (o *AssetsInfoLiabilities) HasEquity() bool`

HasEquity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


