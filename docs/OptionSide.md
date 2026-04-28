# OptionSide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ask** | Pointer to **float64** |  | [optional] 
**Bid** | Pointer to **float64** |  | [optional] 
**Change** | Pointer to **float64** |  | [optional] 
**ContractName** | Pointer to **string** |  | [optional] 
**ImpliedVolatility** | Pointer to **float64** |  | [optional] 
**InTheMoney** | Pointer to **bool** |  | [optional] 
**LastPrice** | Pointer to **float64** |  | [optional] 
**LastTradeDate** | Pointer to **string** |  | [optional] 
**OpenInterest** | Pointer to **int64** |  | [optional] 
**OptionId** | Pointer to **string** |  | [optional] 
**PercentChange** | Pointer to **float64** |  | [optional] 
**Strike** | Pointer to **float64** |  | [optional] 
**Volume** | Pointer to **int64** |  | [optional] 

## Methods

### NewOptionSide

`func NewOptionSide() *OptionSide`

NewOptionSide instantiates a new OptionSide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionSideWithDefaults

`func NewOptionSideWithDefaults() *OptionSide`

NewOptionSideWithDefaults instantiates a new OptionSide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsk

`func (o *OptionSide) GetAsk() float64`

GetAsk returns the Ask field if non-nil, zero value otherwise.

### GetAskOk

`func (o *OptionSide) GetAskOk() (*float64, bool)`

GetAskOk returns a tuple with the Ask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsk

`func (o *OptionSide) SetAsk(v float64)`

SetAsk sets Ask field to given value.

### HasAsk

`func (o *OptionSide) HasAsk() bool`

HasAsk returns a boolean if a field has been set.

### GetBid

`func (o *OptionSide) GetBid() float64`

GetBid returns the Bid field if non-nil, zero value otherwise.

### GetBidOk

`func (o *OptionSide) GetBidOk() (*float64, bool)`

GetBidOk returns a tuple with the Bid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBid

`func (o *OptionSide) SetBid(v float64)`

SetBid sets Bid field to given value.

### HasBid

`func (o *OptionSide) HasBid() bool`

HasBid returns a boolean if a field has been set.

### GetChange

`func (o *OptionSide) GetChange() float64`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *OptionSide) GetChangeOk() (*float64, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *OptionSide) SetChange(v float64)`

SetChange sets Change field to given value.

### HasChange

`func (o *OptionSide) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetContractName

`func (o *OptionSide) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *OptionSide) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *OptionSide) SetContractName(v string)`

SetContractName sets ContractName field to given value.

### HasContractName

`func (o *OptionSide) HasContractName() bool`

HasContractName returns a boolean if a field has been set.

### GetImpliedVolatility

`func (o *OptionSide) GetImpliedVolatility() float64`

GetImpliedVolatility returns the ImpliedVolatility field if non-nil, zero value otherwise.

### GetImpliedVolatilityOk

`func (o *OptionSide) GetImpliedVolatilityOk() (*float64, bool)`

GetImpliedVolatilityOk returns a tuple with the ImpliedVolatility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpliedVolatility

`func (o *OptionSide) SetImpliedVolatility(v float64)`

SetImpliedVolatility sets ImpliedVolatility field to given value.

### HasImpliedVolatility

`func (o *OptionSide) HasImpliedVolatility() bool`

HasImpliedVolatility returns a boolean if a field has been set.

### GetInTheMoney

`func (o *OptionSide) GetInTheMoney() bool`

GetInTheMoney returns the InTheMoney field if non-nil, zero value otherwise.

### GetInTheMoneyOk

`func (o *OptionSide) GetInTheMoneyOk() (*bool, bool)`

GetInTheMoneyOk returns a tuple with the InTheMoney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTheMoney

`func (o *OptionSide) SetInTheMoney(v bool)`

SetInTheMoney sets InTheMoney field to given value.

### HasInTheMoney

`func (o *OptionSide) HasInTheMoney() bool`

HasInTheMoney returns a boolean if a field has been set.

### GetLastPrice

`func (o *OptionSide) GetLastPrice() float64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *OptionSide) GetLastPriceOk() (*float64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *OptionSide) SetLastPrice(v float64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *OptionSide) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastTradeDate

`func (o *OptionSide) GetLastTradeDate() string`

GetLastTradeDate returns the LastTradeDate field if non-nil, zero value otherwise.

### GetLastTradeDateOk

`func (o *OptionSide) GetLastTradeDateOk() (*string, bool)`

GetLastTradeDateOk returns a tuple with the LastTradeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTradeDate

`func (o *OptionSide) SetLastTradeDate(v string)`

SetLastTradeDate sets LastTradeDate field to given value.

### HasLastTradeDate

`func (o *OptionSide) HasLastTradeDate() bool`

HasLastTradeDate returns a boolean if a field has been set.

### GetOpenInterest

`func (o *OptionSide) GetOpenInterest() int64`

GetOpenInterest returns the OpenInterest field if non-nil, zero value otherwise.

### GetOpenInterestOk

`func (o *OptionSide) GetOpenInterestOk() (*int64, bool)`

GetOpenInterestOk returns a tuple with the OpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterest

`func (o *OptionSide) SetOpenInterest(v int64)`

SetOpenInterest sets OpenInterest field to given value.

### HasOpenInterest

`func (o *OptionSide) HasOpenInterest() bool`

HasOpenInterest returns a boolean if a field has been set.

### GetOptionId

`func (o *OptionSide) GetOptionId() string`

GetOptionId returns the OptionId field if non-nil, zero value otherwise.

### GetOptionIdOk

`func (o *OptionSide) GetOptionIdOk() (*string, bool)`

GetOptionIdOk returns a tuple with the OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionId

`func (o *OptionSide) SetOptionId(v string)`

SetOptionId sets OptionId field to given value.

### HasOptionId

`func (o *OptionSide) HasOptionId() bool`

HasOptionId returns a boolean if a field has been set.

### GetPercentChange

`func (o *OptionSide) GetPercentChange() float64`

GetPercentChange returns the PercentChange field if non-nil, zero value otherwise.

### GetPercentChangeOk

`func (o *OptionSide) GetPercentChangeOk() (*float64, bool)`

GetPercentChangeOk returns a tuple with the PercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChange

`func (o *OptionSide) SetPercentChange(v float64)`

SetPercentChange sets PercentChange field to given value.

### HasPercentChange

`func (o *OptionSide) HasPercentChange() bool`

HasPercentChange returns a boolean if a field has been set.

### GetStrike

`func (o *OptionSide) GetStrike() float64`

GetStrike returns the Strike field if non-nil, zero value otherwise.

### GetStrikeOk

`func (o *OptionSide) GetStrikeOk() (*float64, bool)`

GetStrikeOk returns a tuple with the Strike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrike

`func (o *OptionSide) SetStrike(v float64)`

SetStrike sets Strike field to given value.

### HasStrike

`func (o *OptionSide) HasStrike() bool`

HasStrike returns a boolean if a field has been set.

### GetVolume

`func (o *OptionSide) GetVolume() int64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *OptionSide) GetVolumeOk() (*int64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *OptionSide) SetVolume(v int64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *OptionSide) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


