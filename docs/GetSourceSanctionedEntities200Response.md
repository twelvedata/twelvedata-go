# GetSourceSanctionedEntities200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sanctions** | [**[]ResponseSanctionedEntitiy**](ResponseSanctionedEntitiy.md) | List of sanctioned entities | 
**Count** | **int64** | Total number of sanctioned entities | 
**Status** | **string** | Response status | 

## Methods

### NewGetSourceSanctionedEntities200Response

`func NewGetSourceSanctionedEntities200Response(sanctions []ResponseSanctionedEntitiy, count int64, status string, ) *GetSourceSanctionedEntities200Response`

NewGetSourceSanctionedEntities200Response instantiates a new GetSourceSanctionedEntities200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSourceSanctionedEntities200ResponseWithDefaults

`func NewGetSourceSanctionedEntities200ResponseWithDefaults() *GetSourceSanctionedEntities200Response`

NewGetSourceSanctionedEntities200ResponseWithDefaults instantiates a new GetSourceSanctionedEntities200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSanctions

`func (o *GetSourceSanctionedEntities200Response) GetSanctions() []ResponseSanctionedEntitiy`

GetSanctions returns the Sanctions field if non-nil, zero value otherwise.

### GetSanctionsOk

`func (o *GetSourceSanctionedEntities200Response) GetSanctionsOk() (*[]ResponseSanctionedEntitiy, bool)`

GetSanctionsOk returns a tuple with the Sanctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctions

`func (o *GetSourceSanctionedEntities200Response) SetSanctions(v []ResponseSanctionedEntitiy)`

SetSanctions sets Sanctions field to given value.


### GetCount

`func (o *GetSourceSanctionedEntities200Response) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetSourceSanctionedEntities200Response) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetSourceSanctionedEntities200Response) SetCount(v int64)`

SetCount sets Count field to given value.


### GetStatus

`func (o *GetSourceSanctionedEntities200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSourceSanctionedEntities200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSourceSanctionedEntities200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


