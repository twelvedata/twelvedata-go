# GetEarningsCalendar200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Earnings** | [**map[string][]GetEarningsCalendar200ResponseEarningsValueInner**](array.md) | Map of dates to earnings data | 
**Status** | **string** | Response status | 

## Methods

### NewGetEarningsCalendar200Response

`func NewGetEarningsCalendar200Response(earnings map[string][]GetEarningsCalendar200ResponseEarningsValueInner, status string, ) *GetEarningsCalendar200Response`

NewGetEarningsCalendar200Response instantiates a new GetEarningsCalendar200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEarningsCalendar200ResponseWithDefaults

`func NewGetEarningsCalendar200ResponseWithDefaults() *GetEarningsCalendar200Response`

NewGetEarningsCalendar200ResponseWithDefaults instantiates a new GetEarningsCalendar200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEarnings

`func (o *GetEarningsCalendar200Response) GetEarnings() map[string][]GetEarningsCalendar200ResponseEarningsValueInner`

GetEarnings returns the Earnings field if non-nil, zero value otherwise.

### GetEarningsOk

`func (o *GetEarningsCalendar200Response) GetEarningsOk() (*map[string][]GetEarningsCalendar200ResponseEarningsValueInner, bool)`

GetEarningsOk returns a tuple with the Earnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnings

`func (o *GetEarningsCalendar200Response) SetEarnings(v map[string][]GetEarningsCalendar200ResponseEarningsValueInner)`

SetEarnings sets Earnings field to given value.


### GetStatus

`func (o *GetEarningsCalendar200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEarningsCalendar200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEarningsCalendar200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


