# GetInsiderTransactions200ResponseInsiderTransactionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | Full name of an individual, including first name, middle name, last name, and suffix | 
**Position** | **string** | Job position of insider | 
**DateReported** | **string** | Date the transaction was reported | 
**IsDirect** | Pointer to **bool** | &#x60;true&#x60; if direct, &#x60;false&#x60; if indirect | [optional] 
**Shares** | Pointer to **int64** | As per report the number of shares acquired or disposed of the transaction | [optional] 
**Value** | Pointer to **int64** | Represents the value of transaction, calculated as price multiplied by the volume | [optional] 
**Description** | **string** | Exact price or price range of the transaction if available | 

## Methods

### NewGetInsiderTransactions200ResponseInsiderTransactionsInner

`func NewGetInsiderTransactions200ResponseInsiderTransactionsInner(fullName string, position string, dateReported string, description string, ) *GetInsiderTransactions200ResponseInsiderTransactionsInner`

NewGetInsiderTransactions200ResponseInsiderTransactionsInner instantiates a new GetInsiderTransactions200ResponseInsiderTransactionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInsiderTransactions200ResponseInsiderTransactionsInnerWithDefaults

`func NewGetInsiderTransactions200ResponseInsiderTransactionsInnerWithDefaults() *GetInsiderTransactions200ResponseInsiderTransactionsInner`

NewGetInsiderTransactions200ResponseInsiderTransactionsInnerWithDefaults instantiates a new GetInsiderTransactions200ResponseInsiderTransactionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetPosition

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetDateReported

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetDateReported() string`

GetDateReported returns the DateReported field if non-nil, zero value otherwise.

### GetDateReportedOk

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetDateReportedOk() (*string, bool)`

GetDateReportedOk returns a tuple with the DateReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateReported

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetDateReported(v string)`

SetDateReported sets DateReported field to given value.


### GetIsDirect

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetIsDirect() bool`

GetIsDirect returns the IsDirect field if non-nil, zero value otherwise.

### GetIsDirectOk

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetIsDirectOk() (*bool, bool)`

GetIsDirectOk returns a tuple with the IsDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirect

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetIsDirect(v bool)`

SetIsDirect sets IsDirect field to given value.

### HasIsDirect

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) HasIsDirect() bool`

HasIsDirect returns a boolean if a field has been set.

### GetShares

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetValue

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDescription

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


