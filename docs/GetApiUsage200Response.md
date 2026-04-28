# GetApiUsage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** | Current timestamp in UTC timezone | 
**CurrentUsage** | **int64** | Number of requests made in last minute | 
**PlanLimit** | **int64** | Your personal API limit (requests/minute) depending on the plan | 
**PlanCategory** | Pointer to **string** | Plan category name | [optional] 

## Methods

### NewGetApiUsage200Response

`func NewGetApiUsage200Response(timestamp string, currentUsage int64, planLimit int64, ) *GetApiUsage200Response`

NewGetApiUsage200Response instantiates a new GetApiUsage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiUsage200ResponseWithDefaults

`func NewGetApiUsage200ResponseWithDefaults() *GetApiUsage200Response`

NewGetApiUsage200ResponseWithDefaults instantiates a new GetApiUsage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *GetApiUsage200Response) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetApiUsage200Response) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetApiUsage200Response) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetCurrentUsage

`func (o *GetApiUsage200Response) GetCurrentUsage() int64`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *GetApiUsage200Response) GetCurrentUsageOk() (*int64, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *GetApiUsage200Response) SetCurrentUsage(v int64)`

SetCurrentUsage sets CurrentUsage field to given value.


### GetPlanLimit

`func (o *GetApiUsage200Response) GetPlanLimit() int64`

GetPlanLimit returns the PlanLimit field if non-nil, zero value otherwise.

### GetPlanLimitOk

`func (o *GetApiUsage200Response) GetPlanLimitOk() (*int64, bool)`

GetPlanLimitOk returns a tuple with the PlanLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanLimit

`func (o *GetApiUsage200Response) SetPlanLimit(v int64)`

SetPlanLimit sets PlanLimit field to given value.


### GetPlanCategory

`func (o *GetApiUsage200Response) GetPlanCategory() string`

GetPlanCategory returns the PlanCategory field if non-nil, zero value otherwise.

### GetPlanCategoryOk

`func (o *GetApiUsage200Response) GetPlanCategoryOk() (*string, bool)`

GetPlanCategoryOk returns a tuple with the PlanCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCategory

`func (o *GetApiUsage200Response) SetPlanCategory(v string)`

SetPlanCategory sets PlanCategory field to given value.

### HasPlanCategory

`func (o *GetApiUsage200Response) HasPlanCategory() bool`

HasPlanCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


