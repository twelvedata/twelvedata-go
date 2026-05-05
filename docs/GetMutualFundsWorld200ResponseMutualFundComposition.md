# GetMutualFundsWorld200ResponseMutualFundComposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MajorMarketSectors** | Pointer to [**[]GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner**](GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner.md) | Breakdown of the fund’s portfolio by major industry sectors and their respective weights | [optional] 
**AssetAllocation** | Pointer to [**GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation**](GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation.md) |  | [optional] 
**TopHoldings** | Pointer to [**[]GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner**](GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner.md) | Top holdings of the fund with their respective weights in the overall portfolio composition | [optional] 
**BondBreakdown** | Pointer to [**GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown**](GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown.md) |  | [optional] 

## Methods

### NewGetMutualFundsWorld200ResponseMutualFundComposition

`func NewGetMutualFundsWorld200ResponseMutualFundComposition() *GetMutualFundsWorld200ResponseMutualFundComposition`

NewGetMutualFundsWorld200ResponseMutualFundComposition instantiates a new GetMutualFundsWorld200ResponseMutualFundComposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMutualFundsWorld200ResponseMutualFundCompositionWithDefaults

`func NewGetMutualFundsWorld200ResponseMutualFundCompositionWithDefaults() *GetMutualFundsWorld200ResponseMutualFundComposition`

NewGetMutualFundsWorld200ResponseMutualFundCompositionWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundComposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMajorMarketSectors

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetMajorMarketSectors() []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner`

GetMajorMarketSectors returns the MajorMarketSectors field if non-nil, zero value otherwise.

### GetMajorMarketSectorsOk

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetMajorMarketSectorsOk() (*[]GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner, bool)`

GetMajorMarketSectorsOk returns a tuple with the MajorMarketSectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorMarketSectors

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) SetMajorMarketSectors(v []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner)`

SetMajorMarketSectors sets MajorMarketSectors field to given value.

### HasMajorMarketSectors

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) HasMajorMarketSectors() bool`

HasMajorMarketSectors returns a boolean if a field has been set.

### GetAssetAllocation

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetAssetAllocation() GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation`

GetAssetAllocation returns the AssetAllocation field if non-nil, zero value otherwise.

### GetAssetAllocationOk

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetAssetAllocationOk() (*GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation, bool)`

GetAssetAllocationOk returns a tuple with the AssetAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetAllocation

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) SetAssetAllocation(v GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation)`

SetAssetAllocation sets AssetAllocation field to given value.

### HasAssetAllocation

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) HasAssetAllocation() bool`

HasAssetAllocation returns a boolean if a field has been set.

### GetTopHoldings

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetTopHoldings() []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner`

GetTopHoldings returns the TopHoldings field if non-nil, zero value otherwise.

### GetTopHoldingsOk

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetTopHoldingsOk() (*[]GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner, bool)`

GetTopHoldingsOk returns a tuple with the TopHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopHoldings

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) SetTopHoldings(v []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner)`

SetTopHoldings sets TopHoldings field to given value.

### HasTopHoldings

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) HasTopHoldings() bool`

HasTopHoldings returns a boolean if a field has been set.

### GetBondBreakdown

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetBondBreakdown() GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown`

GetBondBreakdown returns the BondBreakdown field if non-nil, zero value otherwise.

### GetBondBreakdownOk

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetBondBreakdownOk() (*GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown, bool)`

GetBondBreakdownOk returns a tuple with the BondBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondBreakdown

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) SetBondBreakdown(v GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown)`

SetBondBreakdown sets BondBreakdown field to given value.

### HasBondBreakdown

`func (o *GetMutualFundsWorld200ResponseMutualFundComposition) HasBondBreakdown() bool`

HasBondBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


