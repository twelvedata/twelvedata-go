# MarketStateResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The full name of exchange | 
**Code** | **string** | Market identifier code (MIC) under ISO 10383 standard | 
**Country** | **string** | Country where the exchange is located | 
**IsMarketOpen** | **bool** | Indicates if the market is currently open | 
**TimeAfterOpen** | **string** | Time after market opening in &lt;code&gt;HH:MM:SS&lt;/code&gt; format; if currently closed - returns &lt;code&gt;00:00:00&lt;/code&gt; | 
**TimeToOpen** | **string** | Time to market opening in &lt;code&gt;HH:MM:SS&lt;/code&gt; format; if currently open - returns &lt;code&gt;00:00:00&lt;/code&gt; | 
**TimeToClose** | **string** | Time to market closing in &lt;code&gt;HH:MM:SS&lt;/code&gt; format; if currently closed - returns &lt;code&gt;00:00:00&lt;/code&gt; | 

## Methods

### NewMarketStateResponseItem

`func NewMarketStateResponseItem(name string, code string, country string, isMarketOpen bool, timeAfterOpen string, timeToOpen string, timeToClose string, ) *MarketStateResponseItem`

NewMarketStateResponseItem instantiates a new MarketStateResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketStateResponseItemWithDefaults

`func NewMarketStateResponseItemWithDefaults() *MarketStateResponseItem`

NewMarketStateResponseItemWithDefaults instantiates a new MarketStateResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MarketStateResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketStateResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketStateResponseItem) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *MarketStateResponseItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MarketStateResponseItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MarketStateResponseItem) SetCode(v string)`

SetCode sets Code field to given value.


### GetCountry

`func (o *MarketStateResponseItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MarketStateResponseItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MarketStateResponseItem) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetIsMarketOpen

`func (o *MarketStateResponseItem) GetIsMarketOpen() bool`

GetIsMarketOpen returns the IsMarketOpen field if non-nil, zero value otherwise.

### GetIsMarketOpenOk

`func (o *MarketStateResponseItem) GetIsMarketOpenOk() (*bool, bool)`

GetIsMarketOpenOk returns a tuple with the IsMarketOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketOpen

`func (o *MarketStateResponseItem) SetIsMarketOpen(v bool)`

SetIsMarketOpen sets IsMarketOpen field to given value.


### GetTimeAfterOpen

`func (o *MarketStateResponseItem) GetTimeAfterOpen() string`

GetTimeAfterOpen returns the TimeAfterOpen field if non-nil, zero value otherwise.

### GetTimeAfterOpenOk

`func (o *MarketStateResponseItem) GetTimeAfterOpenOk() (*string, bool)`

GetTimeAfterOpenOk returns a tuple with the TimeAfterOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAfterOpen

`func (o *MarketStateResponseItem) SetTimeAfterOpen(v string)`

SetTimeAfterOpen sets TimeAfterOpen field to given value.


### GetTimeToOpen

`func (o *MarketStateResponseItem) GetTimeToOpen() string`

GetTimeToOpen returns the TimeToOpen field if non-nil, zero value otherwise.

### GetTimeToOpenOk

`func (o *MarketStateResponseItem) GetTimeToOpenOk() (*string, bool)`

GetTimeToOpenOk returns a tuple with the TimeToOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToOpen

`func (o *MarketStateResponseItem) SetTimeToOpen(v string)`

SetTimeToOpen sets TimeToOpen field to given value.


### GetTimeToClose

`func (o *MarketStateResponseItem) GetTimeToClose() string`

GetTimeToClose returns the TimeToClose field if non-nil, zero value otherwise.

### GetTimeToCloseOk

`func (o *MarketStateResponseItem) GetTimeToCloseOk() (*string, bool)`

GetTimeToCloseOk returns a tuple with the TimeToClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToClose

`func (o *MarketStateResponseItem) SetTimeToClose(v string)`

SetTimeToClose sets TimeToClose field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


