# GetAssetsResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Currency code | 
**Description** | Pointer to **string** | Description of the asset | [optional] 
**MicCode** | **string** | Market identifier code, e.g. DIGITAL_CURRENCY, PHYSICAL_CURRENCY, etc. | 
**Symbol** | Pointer to **string** | Currency symbol | [optional] 

## Methods

### NewGetAssetsResponseItem

`func NewGetAssetsResponseItem(code string, micCode string, ) *GetAssetsResponseItem`

NewGetAssetsResponseItem instantiates a new GetAssetsResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetsResponseItemWithDefaults

`func NewGetAssetsResponseItemWithDefaults() *GetAssetsResponseItem`

NewGetAssetsResponseItemWithDefaults instantiates a new GetAssetsResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetAssetsResponseItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetAssetsResponseItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetAssetsResponseItem) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *GetAssetsResponseItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAssetsResponseItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAssetsResponseItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAssetsResponseItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMicCode

`func (o *GetAssetsResponseItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetAssetsResponseItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetAssetsResponseItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetSymbol

`func (o *GetAssetsResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetAssetsResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetAssetsResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetAssetsResponseItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


