# ResponseSanctionItemList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The sanction list name | 
**PublishedAt** | Pointer to **string** | The sanction published date in the current sanctions list | [optional] 

## Methods

### NewResponseSanctionItemList

`func NewResponseSanctionItemList(name string, ) *ResponseSanctionItemList`

NewResponseSanctionItemList instantiates a new ResponseSanctionItemList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSanctionItemListWithDefaults

`func NewResponseSanctionItemListWithDefaults() *ResponseSanctionItemList`

NewResponseSanctionItemListWithDefaults instantiates a new ResponseSanctionItemList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseSanctionItemList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseSanctionItemList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseSanctionItemList) SetName(v string)`

SetName sets Name field to given value.


### GetPublishedAt

`func (o *ResponseSanctionItemList) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *ResponseSanctionItemList) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *ResponseSanctionItemList) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *ResponseSanctionItemList) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


