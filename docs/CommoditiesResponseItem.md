# CommoditiesResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Currency pair according to ISO 4217 standard codes with slash(/) delimiter | 
**Name** | **string** | Full name of the instrument | 
**Category** | **string** | Category of commodity | 
**Description** | **string** | Short description of the commodity | 

## Methods

### NewCommoditiesResponseItem

`func NewCommoditiesResponseItem(symbol string, name string, category string, description string, ) *CommoditiesResponseItem`

NewCommoditiesResponseItem instantiates a new CommoditiesResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommoditiesResponseItemWithDefaults

`func NewCommoditiesResponseItemWithDefaults() *CommoditiesResponseItem`

NewCommoditiesResponseItemWithDefaults instantiates a new CommoditiesResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *CommoditiesResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CommoditiesResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CommoditiesResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *CommoditiesResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommoditiesResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommoditiesResponseItem) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *CommoditiesResponseItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CommoditiesResponseItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CommoditiesResponseItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetDescription

`func (o *CommoditiesResponseItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommoditiesResponseItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommoditiesResponseItem) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


