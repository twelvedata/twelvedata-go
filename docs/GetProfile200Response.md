# GetProfile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Ticker of the company | 
**Name** | **string** | Name of the company | 
**Exchange** | **string** | Exchange name where the company is listed | 
**MicCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
**Sector** | Pointer to **string** | Sector at which the company operates | [optional] 
**Industry** | Pointer to **string** | Industry at which company operates | [optional] 
**Employees** | Pointer to **int64** | Number of employees in the company | [optional] 
**Website** | Pointer to **string** | Website of the company | [optional] 
**Description** | Pointer to **string** | Description of the company activities | [optional] 
**Type** | Pointer to **string** | Issue type of the stock | [optional] 
**CEO** | Pointer to **string** | Name of the CEO of the company | [optional] 
**Address** | Pointer to **string** | Street address of the company if presented | [optional] 
**Address2** | Pointer to **string** | Secondary address of the company if presented | [optional] 
**City** | Pointer to **string** | City of the company if presented | [optional] 
**Zip** | Pointer to **string** | ZIP code of the company if presented | [optional] 
**State** | Pointer to **string** | State of the company if presented | [optional] 
**Country** | Pointer to **string** | Country of the company if presented | [optional] 
**Phone** | Pointer to **string** | Phone number of the company if presented | [optional] 

## Methods

### NewGetProfile200Response

`func NewGetProfile200Response(symbol string, name string, exchange string, micCode string, ) *GetProfile200Response`

NewGetProfile200Response instantiates a new GetProfile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfile200ResponseWithDefaults

`func NewGetProfile200ResponseWithDefaults() *GetProfile200Response`

NewGetProfile200ResponseWithDefaults instantiates a new GetProfile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetProfile200Response) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetProfile200Response) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetProfile200Response) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *GetProfile200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetProfile200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetProfile200Response) SetName(v string)`

SetName sets Name field to given value.


### GetExchange

`func (o *GetProfile200Response) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetProfile200Response) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetProfile200Response) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *GetProfile200Response) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetProfile200Response) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetProfile200Response) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetSector

`func (o *GetProfile200Response) GetSector() string`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *GetProfile200Response) GetSectorOk() (*string, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *GetProfile200Response) SetSector(v string)`

SetSector sets Sector field to given value.

### HasSector

`func (o *GetProfile200Response) HasSector() bool`

HasSector returns a boolean if a field has been set.

### GetIndustry

`func (o *GetProfile200Response) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *GetProfile200Response) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *GetProfile200Response) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *GetProfile200Response) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetEmployees

`func (o *GetProfile200Response) GetEmployees() int64`

GetEmployees returns the Employees field if non-nil, zero value otherwise.

### GetEmployeesOk

`func (o *GetProfile200Response) GetEmployeesOk() (*int64, bool)`

GetEmployeesOk returns a tuple with the Employees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployees

`func (o *GetProfile200Response) SetEmployees(v int64)`

SetEmployees sets Employees field to given value.

### HasEmployees

`func (o *GetProfile200Response) HasEmployees() bool`

HasEmployees returns a boolean if a field has been set.

### GetWebsite

`func (o *GetProfile200Response) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *GetProfile200Response) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *GetProfile200Response) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *GetProfile200Response) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetDescription

`func (o *GetProfile200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetProfile200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetProfile200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetProfile200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *GetProfile200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetProfile200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetProfile200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetProfile200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCEO

`func (o *GetProfile200Response) GetCEO() string`

GetCEO returns the CEO field if non-nil, zero value otherwise.

### GetCEOOk

`func (o *GetProfile200Response) GetCEOOk() (*string, bool)`

GetCEOOk returns a tuple with the CEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCEO

`func (o *GetProfile200Response) SetCEO(v string)`

SetCEO sets CEO field to given value.

### HasCEO

`func (o *GetProfile200Response) HasCEO() bool`

HasCEO returns a boolean if a field has been set.

### GetAddress

`func (o *GetProfile200Response) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetProfile200Response) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetProfile200Response) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetProfile200Response) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddress2

`func (o *GetProfile200Response) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *GetProfile200Response) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *GetProfile200Response) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *GetProfile200Response) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *GetProfile200Response) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GetProfile200Response) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GetProfile200Response) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GetProfile200Response) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetZip

`func (o *GetProfile200Response) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *GetProfile200Response) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *GetProfile200Response) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *GetProfile200Response) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetState

`func (o *GetProfile200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetProfile200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetProfile200Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetProfile200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *GetProfile200Response) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetProfile200Response) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetProfile200Response) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetProfile200Response) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPhone

`func (o *GetProfile200Response) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GetProfile200Response) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GetProfile200Response) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GetProfile200Response) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


