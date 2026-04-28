# ResponseMutualFundWorldComposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MajorMarketSectors** | Pointer to [**[]GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner**](GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner.md) | Breakdown of the fund’s portfolio by major industry sectors and their respective weights | [optional] 
**AssetAllocation** | Pointer to [**GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation**](GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation.md) |  | [optional] 
**TopHoldings** | Pointer to [**[]GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner**](GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner.md) | Top holdings of the fund with their respective weights in the overall portfolio composition | [optional] 
**BondBreakdown** | Pointer to [**GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown**](GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown.md) |  | [optional] 

## Methods

### NewResponseMutualFundWorldComposition

`func NewResponseMutualFundWorldComposition() *ResponseMutualFundWorldComposition`

NewResponseMutualFundWorldComposition instantiates a new ResponseMutualFundWorldComposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseMutualFundWorldCompositionWithDefaults

`func NewResponseMutualFundWorldCompositionWithDefaults() *ResponseMutualFundWorldComposition`

NewResponseMutualFundWorldCompositionWithDefaults instantiates a new ResponseMutualFundWorldComposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMajorMarketSectors

`func (o *ResponseMutualFundWorldComposition) GetMajorMarketSectors() []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner`

GetMajorMarketSectors returns the MajorMarketSectors field if non-nil, zero value otherwise.

### GetMajorMarketSectorsOk

`func (o *ResponseMutualFundWorldComposition) GetMajorMarketSectorsOk() (*[]GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner, bool)`

GetMajorMarketSectorsOk returns a tuple with the MajorMarketSectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorMarketSectors

`func (o *ResponseMutualFundWorldComposition) SetMajorMarketSectors(v []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner)`

SetMajorMarketSectors sets MajorMarketSectors field to given value.

### HasMajorMarketSectors

`func (o *ResponseMutualFundWorldComposition) HasMajorMarketSectors() bool`

HasMajorMarketSectors returns a boolean if a field has been set.

### GetAssetAllocation

`func (o *ResponseMutualFundWorldComposition) GetAssetAllocation() GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation`

GetAssetAllocation returns the AssetAllocation field if non-nil, zero value otherwise.

### GetAssetAllocationOk

`func (o *ResponseMutualFundWorldComposition) GetAssetAllocationOk() (*GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation, bool)`

GetAssetAllocationOk returns a tuple with the AssetAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetAllocation

`func (o *ResponseMutualFundWorldComposition) SetAssetAllocation(v GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation)`

SetAssetAllocation sets AssetAllocation field to given value.

### HasAssetAllocation

`func (o *ResponseMutualFundWorldComposition) HasAssetAllocation() bool`

HasAssetAllocation returns a boolean if a field has been set.

### GetTopHoldings

`func (o *ResponseMutualFundWorldComposition) GetTopHoldings() []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner`

GetTopHoldings returns the TopHoldings field if non-nil, zero value otherwise.

### GetTopHoldingsOk

`func (o *ResponseMutualFundWorldComposition) GetTopHoldingsOk() (*[]GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner, bool)`

GetTopHoldingsOk returns a tuple with the TopHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopHoldings

`func (o *ResponseMutualFundWorldComposition) SetTopHoldings(v []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner)`

SetTopHoldings sets TopHoldings field to given value.

### HasTopHoldings

`func (o *ResponseMutualFundWorldComposition) HasTopHoldings() bool`

HasTopHoldings returns a boolean if a field has been set.

### GetBondBreakdown

`func (o *ResponseMutualFundWorldComposition) GetBondBreakdown() GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown`

GetBondBreakdown returns the BondBreakdown field if non-nil, zero value otherwise.

### GetBondBreakdownOk

`func (o *ResponseMutualFundWorldComposition) GetBondBreakdownOk() (*GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown, bool)`

GetBondBreakdownOk returns a tuple with the BondBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondBreakdown

`func (o *ResponseMutualFundWorldComposition) SetBondBreakdown(v GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown)`

SetBondBreakdown sets BondBreakdown field to given value.

### HasBondBreakdown

`func (o *ResponseMutualFundWorldComposition) HasBondBreakdown() bool`

HasBondBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


