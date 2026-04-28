# GetTaxInfo200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The ticker symbol of an instrument for which data was requested | 
**Name** | **string** | The instrument name | 
**Exchange** | **string** | The exchange name where the instrument is traded | 
**MicCode** | **string** | The Market Identifier Code (MIC) of the exchange where the instrument is traded | 
**Country** | **string** | The instrument country name | 

## Methods

### NewGetTaxInfo200ResponseMeta

`func NewGetTaxInfo200ResponseMeta(symbol string, name string, exchange string, micCode string, country string, ) *GetTaxInfo200ResponseMeta`

NewGetTaxInfo200ResponseMeta instantiates a new GetTaxInfo200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaxInfo200ResponseMetaWithDefaults

`func NewGetTaxInfo200ResponseMetaWithDefaults() *GetTaxInfo200ResponseMeta`

NewGetTaxInfo200ResponseMetaWithDefaults instantiates a new GetTaxInfo200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetTaxInfo200ResponseMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetTaxInfo200ResponseMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetTaxInfo200ResponseMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *GetTaxInfo200ResponseMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTaxInfo200ResponseMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTaxInfo200ResponseMeta) SetName(v string)`

SetName sets Name field to given value.


### GetExchange

`func (o *GetTaxInfo200ResponseMeta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetTaxInfo200ResponseMeta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetTaxInfo200ResponseMeta) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *GetTaxInfo200ResponseMeta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetTaxInfo200ResponseMeta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetTaxInfo200ResponseMeta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetCountry

`func (o *GetTaxInfo200ResponseMeta) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetTaxInfo200ResponseMeta) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetTaxInfo200ResponseMeta) SetCountry(v string)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


