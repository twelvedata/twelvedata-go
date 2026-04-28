# MarketMoversResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | [**[]MarketMoversResponseValue**](MarketMoversResponseValue.md) | Market movers list | 
**Status** | **string** | Response status | 

## Methods

### NewMarketMoversResponseBody

`func NewMarketMoversResponseBody(values []MarketMoversResponseValue, status string, ) *MarketMoversResponseBody`

NewMarketMoversResponseBody instantiates a new MarketMoversResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketMoversResponseBodyWithDefaults

`func NewMarketMoversResponseBodyWithDefaults() *MarketMoversResponseBody`

NewMarketMoversResponseBodyWithDefaults instantiates a new MarketMoversResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *MarketMoversResponseBody) GetValues() []MarketMoversResponseValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MarketMoversResponseBody) GetValuesOk() (*[]MarketMoversResponseValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MarketMoversResponseBody) SetValues(v []MarketMoversResponseValue)`

SetValues sets Values field to given value.


### GetStatus

`func (o *MarketMoversResponseBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketMoversResponseBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketMoversResponseBody) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


