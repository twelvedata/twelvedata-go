# GetETFsWorld200ResponseEtfComposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MajorMarketSectors** | Pointer to [**[]GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner**](GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner.md) | Breakdown of the fund’s portfolio by major industry sectors and their respective weights | [optional] 
**CountryAllocation** | Pointer to [**[]GetETFsWorld200ResponseEtfCompositionCountryAllocationInner**](GetETFsWorld200ResponseEtfCompositionCountryAllocationInner.md) | Breakdown of the fund’s portfolio by country and their respective weights | [optional] 
**AssetAllocation** | Pointer to [**GetETFsWorld200ResponseEtfCompositionAssetAllocation**](GetETFsWorld200ResponseEtfCompositionAssetAllocation.md) |  | [optional] 
**TopHoldings** | Pointer to [**[]GetETFsWorld200ResponseEtfCompositionTopHoldingsInner**](GetETFsWorld200ResponseEtfCompositionTopHoldingsInner.md) | Top holdings of a fund with their respective weights in the overall portfolio composition | [optional] 
**BondBreakdown** | Pointer to [**GetETFsWorld200ResponseEtfCompositionBondBreakdown**](GetETFsWorld200ResponseEtfCompositionBondBreakdown.md) |  | [optional] 

## Methods

### NewGetETFsWorld200ResponseEtfComposition

`func NewGetETFsWorld200ResponseEtfComposition() *GetETFsWorld200ResponseEtfComposition`

NewGetETFsWorld200ResponseEtfComposition instantiates a new GetETFsWorld200ResponseEtfComposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetETFsWorld200ResponseEtfCompositionWithDefaults

`func NewGetETFsWorld200ResponseEtfCompositionWithDefaults() *GetETFsWorld200ResponseEtfComposition`

NewGetETFsWorld200ResponseEtfCompositionWithDefaults instantiates a new GetETFsWorld200ResponseEtfComposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMajorMarketSectors

`func (o *GetETFsWorld200ResponseEtfComposition) GetMajorMarketSectors() []GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner`

GetMajorMarketSectors returns the MajorMarketSectors field if non-nil, zero value otherwise.

### GetMajorMarketSectorsOk

`func (o *GetETFsWorld200ResponseEtfComposition) GetMajorMarketSectorsOk() (*[]GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner, bool)`

GetMajorMarketSectorsOk returns a tuple with the MajorMarketSectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorMarketSectors

`func (o *GetETFsWorld200ResponseEtfComposition) SetMajorMarketSectors(v []GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner)`

SetMajorMarketSectors sets MajorMarketSectors field to given value.

### HasMajorMarketSectors

`func (o *GetETFsWorld200ResponseEtfComposition) HasMajorMarketSectors() bool`

HasMajorMarketSectors returns a boolean if a field has been set.

### GetCountryAllocation

`func (o *GetETFsWorld200ResponseEtfComposition) GetCountryAllocation() []GetETFsWorld200ResponseEtfCompositionCountryAllocationInner`

GetCountryAllocation returns the CountryAllocation field if non-nil, zero value otherwise.

### GetCountryAllocationOk

`func (o *GetETFsWorld200ResponseEtfComposition) GetCountryAllocationOk() (*[]GetETFsWorld200ResponseEtfCompositionCountryAllocationInner, bool)`

GetCountryAllocationOk returns a tuple with the CountryAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryAllocation

`func (o *GetETFsWorld200ResponseEtfComposition) SetCountryAllocation(v []GetETFsWorld200ResponseEtfCompositionCountryAllocationInner)`

SetCountryAllocation sets CountryAllocation field to given value.

### HasCountryAllocation

`func (o *GetETFsWorld200ResponseEtfComposition) HasCountryAllocation() bool`

HasCountryAllocation returns a boolean if a field has been set.

### GetAssetAllocation

`func (o *GetETFsWorld200ResponseEtfComposition) GetAssetAllocation() GetETFsWorld200ResponseEtfCompositionAssetAllocation`

GetAssetAllocation returns the AssetAllocation field if non-nil, zero value otherwise.

### GetAssetAllocationOk

`func (o *GetETFsWorld200ResponseEtfComposition) GetAssetAllocationOk() (*GetETFsWorld200ResponseEtfCompositionAssetAllocation, bool)`

GetAssetAllocationOk returns a tuple with the AssetAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetAllocation

`func (o *GetETFsWorld200ResponseEtfComposition) SetAssetAllocation(v GetETFsWorld200ResponseEtfCompositionAssetAllocation)`

SetAssetAllocation sets AssetAllocation field to given value.

### HasAssetAllocation

`func (o *GetETFsWorld200ResponseEtfComposition) HasAssetAllocation() bool`

HasAssetAllocation returns a boolean if a field has been set.

### GetTopHoldings

`func (o *GetETFsWorld200ResponseEtfComposition) GetTopHoldings() []GetETFsWorld200ResponseEtfCompositionTopHoldingsInner`

GetTopHoldings returns the TopHoldings field if non-nil, zero value otherwise.

### GetTopHoldingsOk

`func (o *GetETFsWorld200ResponseEtfComposition) GetTopHoldingsOk() (*[]GetETFsWorld200ResponseEtfCompositionTopHoldingsInner, bool)`

GetTopHoldingsOk returns a tuple with the TopHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopHoldings

`func (o *GetETFsWorld200ResponseEtfComposition) SetTopHoldings(v []GetETFsWorld200ResponseEtfCompositionTopHoldingsInner)`

SetTopHoldings sets TopHoldings field to given value.

### HasTopHoldings

`func (o *GetETFsWorld200ResponseEtfComposition) HasTopHoldings() bool`

HasTopHoldings returns a boolean if a field has been set.

### GetBondBreakdown

`func (o *GetETFsWorld200ResponseEtfComposition) GetBondBreakdown() GetETFsWorld200ResponseEtfCompositionBondBreakdown`

GetBondBreakdown returns the BondBreakdown field if non-nil, zero value otherwise.

### GetBondBreakdownOk

`func (o *GetETFsWorld200ResponseEtfComposition) GetBondBreakdownOk() (*GetETFsWorld200ResponseEtfCompositionBondBreakdown, bool)`

GetBondBreakdownOk returns a tuple with the BondBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondBreakdown

`func (o *GetETFsWorld200ResponseEtfComposition) SetBondBreakdown(v GetETFsWorld200ResponseEtfCompositionBondBreakdown)`

SetBondBreakdown sets BondBreakdown field to given value.

### HasBondBreakdown

`func (o *GetETFsWorld200ResponseEtfComposition) HasBondBreakdown() bool`

HasBondBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


