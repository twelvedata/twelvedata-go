# PressRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Press release unique identifier | 
**Datetime** | **string** | Press release date in ISO 8601 format | 
**Title** | **string** | Press release title | 
**Body** | **string** | Press release body in html format | 
**Style** | Pointer to **string** | Custom style applied to the release | [optional] 
**Language** | Pointer to **[]string** | Press release language codes | [optional] 

## Methods

### NewPressRelease

`func NewPressRelease(id string, datetime string, title string, body string, ) *PressRelease`

NewPressRelease instantiates a new PressRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPressReleaseWithDefaults

`func NewPressReleaseWithDefaults() *PressRelease`

NewPressReleaseWithDefaults instantiates a new PressRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PressRelease) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PressRelease) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PressRelease) SetId(v string)`

SetId sets Id field to given value.


### GetDatetime

`func (o *PressRelease) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *PressRelease) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *PressRelease) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetTitle

`func (o *PressRelease) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PressRelease) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PressRelease) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *PressRelease) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PressRelease) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PressRelease) SetBody(v string)`

SetBody sets Body field to given value.


### GetStyle

`func (o *PressRelease) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *PressRelease) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *PressRelease) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *PressRelease) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetLanguage

`func (o *PressRelease) GetLanguage() []string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PressRelease) GetLanguageOk() (*[]string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PressRelease) SetLanguage(v []string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PressRelease) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


