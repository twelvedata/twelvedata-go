# GetLogo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetLogo200ResponseMeta**](GetLogo200ResponseMeta.md) |  | 
**Url** | Pointer to **string** | Link to download the logo (for stocks only) | [optional] 
**LogoBase** | Pointer to **string** | Link to download the base currency logo (for &#x60;forex&#x60; and &#x60;crypto&#x60; only) | [optional] 
**LogoQuote** | Pointer to **string** | Link to download the quote currency logo (for &#x60;forex&#x60; and &#x60;crypto&#x60; only) | [optional] 

## Methods

### NewGetLogo200Response

`func NewGetLogo200Response(meta GetLogo200ResponseMeta, ) *GetLogo200Response`

NewGetLogo200Response instantiates a new GetLogo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogo200ResponseWithDefaults

`func NewGetLogo200ResponseWithDefaults() *GetLogo200Response`

NewGetLogo200ResponseWithDefaults instantiates a new GetLogo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetLogo200Response) GetMeta() GetLogo200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetLogo200Response) GetMetaOk() (*GetLogo200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetLogo200Response) SetMeta(v GetLogo200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetUrl

`func (o *GetLogo200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetLogo200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetLogo200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetLogo200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLogoBase

`func (o *GetLogo200Response) GetLogoBase() string`

GetLogoBase returns the LogoBase field if non-nil, zero value otherwise.

### GetLogoBaseOk

`func (o *GetLogo200Response) GetLogoBaseOk() (*string, bool)`

GetLogoBaseOk returns a tuple with the LogoBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoBase

`func (o *GetLogo200Response) SetLogoBase(v string)`

SetLogoBase sets LogoBase field to given value.

### HasLogoBase

`func (o *GetLogo200Response) HasLogoBase() bool`

HasLogoBase returns a boolean if a field has been set.

### GetLogoQuote

`func (o *GetLogo200Response) GetLogoQuote() string`

GetLogoQuote returns the LogoQuote field if non-nil, zero value otherwise.

### GetLogoQuoteOk

`func (o *GetLogo200Response) GetLogoQuoteOk() (*string, bool)`

GetLogoQuoteOk returns a tuple with the LogoQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoQuote

`func (o *GetLogo200Response) SetLogoQuote(v string)`

SetLogoQuote sets LogoQuote field to given value.

### HasLogoQuote

`func (o *GetLogo200Response) HasLogoQuote() bool`

HasLogoQuote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


