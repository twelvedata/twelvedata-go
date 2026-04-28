# GetKeyExecutives200ResponseKeyExecutivesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Full name of an executive, including first name, middle name, last name, and suffix | 
**Title** | **string** | Refers to job title | 
**Age** | Pointer to **int64** | Current age of an executive if available | [optional] 
**YearBorn** | Pointer to **int64** | Year of birth of an executive if available | [optional] 
**Pay** | Pointer to **int64** | Total salary of an executive if available | [optional] 

## Methods

### NewGetKeyExecutives200ResponseKeyExecutivesInner

`func NewGetKeyExecutives200ResponseKeyExecutivesInner(name string, title string, ) *GetKeyExecutives200ResponseKeyExecutivesInner`

NewGetKeyExecutives200ResponseKeyExecutivesInner instantiates a new GetKeyExecutives200ResponseKeyExecutivesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKeyExecutives200ResponseKeyExecutivesInnerWithDefaults

`func NewGetKeyExecutives200ResponseKeyExecutivesInnerWithDefaults() *GetKeyExecutives200ResponseKeyExecutivesInner`

NewGetKeyExecutives200ResponseKeyExecutivesInnerWithDefaults instantiates a new GetKeyExecutives200ResponseKeyExecutivesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAge

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetAge() int64`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetAgeOk() (*int64, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) SetAge(v int64)`

SetAge sets Age field to given value.

### HasAge

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetYearBorn

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetYearBorn() int64`

GetYearBorn returns the YearBorn field if non-nil, zero value otherwise.

### GetYearBornOk

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetYearBornOk() (*int64, bool)`

GetYearBornOk returns a tuple with the YearBorn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearBorn

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) SetYearBorn(v int64)`

SetYearBorn sets YearBorn field to given value.

### HasYearBorn

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) HasYearBorn() bool`

HasYearBorn returns a boolean if a field has been set.

### GetPay

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetPay() int64`

GetPay returns the Pay field if non-nil, zero value otherwise.

### GetPayOk

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetPayOk() (*int64, bool)`

GetPayOk returns a tuple with the Pay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPay

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) SetPay(v int64)`

SetPay sets Pay field to given value.

### HasPay

`func (o *GetKeyExecutives200ResponseKeyExecutivesInner) HasPay() bool`

HasPay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


