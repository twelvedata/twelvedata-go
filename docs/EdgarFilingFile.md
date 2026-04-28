# EdgarFilingFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | File name | 
**Size** | Pointer to **int64** | File size | [optional] 
**Type** | **string** | File type | 
**Url** | **string** | File full url | 

## Methods

### NewEdgarFilingFile

`func NewEdgarFilingFile(name string, type_ string, url string, ) *EdgarFilingFile`

NewEdgarFilingFile instantiates a new EdgarFilingFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgarFilingFileWithDefaults

`func NewEdgarFilingFileWithDefaults() *EdgarFilingFile`

NewEdgarFilingFileWithDefaults instantiates a new EdgarFilingFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgarFilingFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgarFilingFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgarFilingFile) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *EdgarFilingFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *EdgarFilingFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *EdgarFilingFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *EdgarFilingFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *EdgarFilingFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgarFilingFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgarFilingFile) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *EdgarFilingFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EdgarFilingFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EdgarFilingFile) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


