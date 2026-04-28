# ResponseSanctionedEntitiy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The instrument symbol ticker | 
**Name** | **string** | The instrument name | 
**MicCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
**Country** | **string** | Country name | 
**Sanction** | [**ResponseSanctionItem**](ResponseSanctionItem.md) |  | 

## Methods

### NewResponseSanctionedEntitiy

`func NewResponseSanctionedEntitiy(symbol string, name string, micCode string, country string, sanction ResponseSanctionItem, ) *ResponseSanctionedEntitiy`

NewResponseSanctionedEntitiy instantiates a new ResponseSanctionedEntitiy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSanctionedEntitiyWithDefaults

`func NewResponseSanctionedEntitiyWithDefaults() *ResponseSanctionedEntitiy`

NewResponseSanctionedEntitiyWithDefaults instantiates a new ResponseSanctionedEntitiy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ResponseSanctionedEntitiy) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ResponseSanctionedEntitiy) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ResponseSanctionedEntitiy) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *ResponseSanctionedEntitiy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseSanctionedEntitiy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseSanctionedEntitiy) SetName(v string)`

SetName sets Name field to given value.


### GetMicCode

`func (o *ResponseSanctionedEntitiy) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *ResponseSanctionedEntitiy) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *ResponseSanctionedEntitiy) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetCountry

`func (o *ResponseSanctionedEntitiy) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ResponseSanctionedEntitiy) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ResponseSanctionedEntitiy) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetSanction

`func (o *ResponseSanctionedEntitiy) GetSanction() ResponseSanctionItem`

GetSanction returns the Sanction field if non-nil, zero value otherwise.

### GetSanctionOk

`func (o *ResponseSanctionedEntitiy) GetSanctionOk() (*ResponseSanctionItem, bool)`

GetSanctionOk returns a tuple with the Sanction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanction

`func (o *ResponseSanctionedEntitiy) SetSanction(v ResponseSanctionItem)`

SetSanction sets Sanction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


