# TimeSeriesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime at local exchange time referring to when the bar with specified interval was opened. | 
**Open** | **string** | Price at the opening of current bar | 
**High** | **string** | Highest price which occurred during the current bar. | 
**Low** | **string** | Lowest price which occurred during the current bar. | 
**Close** | **string** | Close price at the end of the bar. | 
**Volume** | Pointer to **string** | Trading volume which occurred during the current bar | [optional] 

## Methods

### NewTimeSeriesItem

`func NewTimeSeriesItem(datetime string, open string, high string, low string, close string, ) *TimeSeriesItem`

NewTimeSeriesItem instantiates a new TimeSeriesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesItemWithDefaults

`func NewTimeSeriesItemWithDefaults() *TimeSeriesItem`

NewTimeSeriesItemWithDefaults instantiates a new TimeSeriesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *TimeSeriesItem) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *TimeSeriesItem) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *TimeSeriesItem) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetOpen

`func (o *TimeSeriesItem) GetOpen() string`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *TimeSeriesItem) GetOpenOk() (*string, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *TimeSeriesItem) SetOpen(v string)`

SetOpen sets Open field to given value.


### GetHigh

`func (o *TimeSeriesItem) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *TimeSeriesItem) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *TimeSeriesItem) SetHigh(v string)`

SetHigh sets High field to given value.


### GetLow

`func (o *TimeSeriesItem) GetLow() string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *TimeSeriesItem) GetLowOk() (*string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *TimeSeriesItem) SetLow(v string)`

SetLow sets Low field to given value.


### GetClose

`func (o *TimeSeriesItem) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *TimeSeriesItem) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *TimeSeriesItem) SetClose(v string)`

SetClose sets Close field to given value.


### GetVolume

`func (o *TimeSeriesItem) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *TimeSeriesItem) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *TimeSeriesItem) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *TimeSeriesItem) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


