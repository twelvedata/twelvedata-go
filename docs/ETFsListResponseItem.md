# ETFsListResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Instrument symbol (ticker) | 
**Name** | **string** | Full name of the fund | 
**Country** | **string** | Country of fund incorporation | 
**MicCode** | **string** | Market identifier code (MIC) under ISO 10383 standard | 
**FundFamily** | **string** | Investment company that manages the fund | 
**FundType** | **string** | Type of fund | 

## Methods

### NewETFsListResponseItem

`func NewETFsListResponseItem(symbol string, name string, country string, micCode string, fundFamily string, fundType string, ) *ETFsListResponseItem`

NewETFsListResponseItem instantiates a new ETFsListResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewETFsListResponseItemWithDefaults

`func NewETFsListResponseItemWithDefaults() *ETFsListResponseItem`

NewETFsListResponseItemWithDefaults instantiates a new ETFsListResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ETFsListResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ETFsListResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ETFsListResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *ETFsListResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ETFsListResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ETFsListResponseItem) SetName(v string)`

SetName sets Name field to given value.


### GetCountry

`func (o *ETFsListResponseItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ETFsListResponseItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ETFsListResponseItem) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetMicCode

`func (o *ETFsListResponseItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *ETFsListResponseItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *ETFsListResponseItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetFundFamily

`func (o *ETFsListResponseItem) GetFundFamily() string`

GetFundFamily returns the FundFamily field if non-nil, zero value otherwise.

### GetFundFamilyOk

`func (o *ETFsListResponseItem) GetFundFamilyOk() (*string, bool)`

GetFundFamilyOk returns a tuple with the FundFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundFamily

`func (o *ETFsListResponseItem) SetFundFamily(v string)`

SetFundFamily sets FundFamily field to given value.


### GetFundType

`func (o *ETFsListResponseItem) GetFundType() string`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *ETFsListResponseItem) GetFundTypeOk() (*string, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *ETFsListResponseItem) SetFundType(v string)`

SetFundType sets FundType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


