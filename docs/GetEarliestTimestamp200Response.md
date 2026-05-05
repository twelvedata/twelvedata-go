# GetEarliestTimestamp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Earliest datetime, the format depends on interval | 
**UnixTime** | **int64** | Datetime converted to UNIX timestamp | 

## Methods

### NewGetEarliestTimestamp200Response

`func NewGetEarliestTimestamp200Response(datetime string, unixTime int64, ) *GetEarliestTimestamp200Response`

NewGetEarliestTimestamp200Response instantiates a new GetEarliestTimestamp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEarliestTimestamp200ResponseWithDefaults

`func NewGetEarliestTimestamp200ResponseWithDefaults() *GetEarliestTimestamp200Response`

NewGetEarliestTimestamp200ResponseWithDefaults instantiates a new GetEarliestTimestamp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetEarliestTimestamp200Response) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetEarliestTimestamp200Response) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetEarliestTimestamp200Response) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetUnixTime

`func (o *GetEarliestTimestamp200Response) GetUnixTime() int64`

GetUnixTime returns the UnixTime field if non-nil, zero value otherwise.

### GetUnixTimeOk

`func (o *GetEarliestTimestamp200Response) GetUnixTimeOk() (*int64, bool)`

GetUnixTimeOk returns a tuple with the UnixTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixTime

`func (o *GetEarliestTimestamp200Response) SetUnixTime(v int64)`

SetUnixTime sets UnixTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


