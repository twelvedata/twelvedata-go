# CrossListingsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Ticker symbol of instrument | 
**Name** | **string** | Name of symbol | 
**Exchange** | **string** | Exchange where instrument is traded | 
**MicCode** | **string** | Market identifier code (MIC) under ISO 10383 standard | 

## Methods

### NewCrossListingsItem

`func NewCrossListingsItem(symbol string, name string, exchange string, micCode string, ) *CrossListingsItem`

NewCrossListingsItem instantiates a new CrossListingsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossListingsItemWithDefaults

`func NewCrossListingsItemWithDefaults() *CrossListingsItem`

NewCrossListingsItemWithDefaults instantiates a new CrossListingsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *CrossListingsItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CrossListingsItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CrossListingsItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *CrossListingsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrossListingsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrossListingsItem) SetName(v string)`

SetName sets Name field to given value.


### GetExchange

`func (o *CrossListingsItem) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *CrossListingsItem) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *CrossListingsItem) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *CrossListingsItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *CrossListingsItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *CrossListingsItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


