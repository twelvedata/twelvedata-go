# ResponseSanctionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | The sanction source | 
**Program** | **string** | The sanction program | 
**Notes** | **string** | Notes for the sanction | 
**Lists** | [**[]ResponseSanctionItemList**](ResponseSanctionItemList.md) | Sanction lists | 

## Methods

### NewResponseSanctionItem

`func NewResponseSanctionItem(source string, program string, notes string, lists []ResponseSanctionItemList, ) *ResponseSanctionItem`

NewResponseSanctionItem instantiates a new ResponseSanctionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSanctionItemWithDefaults

`func NewResponseSanctionItemWithDefaults() *ResponseSanctionItem`

NewResponseSanctionItemWithDefaults instantiates a new ResponseSanctionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ResponseSanctionItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ResponseSanctionItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ResponseSanctionItem) SetSource(v string)`

SetSource sets Source field to given value.


### GetProgram

`func (o *ResponseSanctionItem) GetProgram() string`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *ResponseSanctionItem) GetProgramOk() (*string, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *ResponseSanctionItem) SetProgram(v string)`

SetProgram sets Program field to given value.


### GetNotes

`func (o *ResponseSanctionItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ResponseSanctionItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ResponseSanctionItem) SetNotes(v string)`

SetNotes sets Notes field to given value.


### GetLists

`func (o *ResponseSanctionItem) GetLists() []ResponseSanctionItemList`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *ResponseSanctionItem) GetListsOk() (*[]ResponseSanctionItemList, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *ResponseSanctionItem) SetLists(v []ResponseSanctionItemList)`

SetLists sets Lists field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


