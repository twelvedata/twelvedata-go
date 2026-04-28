# DirectHolderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityName** | **string** | Refers to the legal name of the institution | 
**DateReported** | Pointer to **string** | Refers to date reported | [optional] 
**Shares** | Pointer to **int64** | Refers to the number of shares owned | [optional] 
**Value** | Pointer to **int64** | Total value of shares owned, calculated by multiplying &#x60;shares&#x60; by the current price | [optional] 
**PercentHeld** | Pointer to **float64** | Represents the percentage of shares outstanding that are owned by the financial institution | [optional] 

## Methods

### NewDirectHolderItem

`func NewDirectHolderItem(entityName string, ) *DirectHolderItem`

NewDirectHolderItem instantiates a new DirectHolderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectHolderItemWithDefaults

`func NewDirectHolderItemWithDefaults() *DirectHolderItem`

NewDirectHolderItemWithDefaults instantiates a new DirectHolderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityName

`func (o *DirectHolderItem) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *DirectHolderItem) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *DirectHolderItem) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.


### GetDateReported

`func (o *DirectHolderItem) GetDateReported() string`

GetDateReported returns the DateReported field if non-nil, zero value otherwise.

### GetDateReportedOk

`func (o *DirectHolderItem) GetDateReportedOk() (*string, bool)`

GetDateReportedOk returns a tuple with the DateReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateReported

`func (o *DirectHolderItem) SetDateReported(v string)`

SetDateReported sets DateReported field to given value.

### HasDateReported

`func (o *DirectHolderItem) HasDateReported() bool`

HasDateReported returns a boolean if a field has been set.

### GetShares

`func (o *DirectHolderItem) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *DirectHolderItem) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *DirectHolderItem) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *DirectHolderItem) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetValue

`func (o *DirectHolderItem) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DirectHolderItem) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DirectHolderItem) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *DirectHolderItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPercentHeld

`func (o *DirectHolderItem) GetPercentHeld() float64`

GetPercentHeld returns the PercentHeld field if non-nil, zero value otherwise.

### GetPercentHeldOk

`func (o *DirectHolderItem) GetPercentHeldOk() (*float64, bool)`

GetPercentHeldOk returns a tuple with the PercentHeld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentHeld

`func (o *DirectHolderItem) SetPercentHeld(v float64)`

SetPercentHeld sets PercentHeld field to given value.

### HasPercentHeld

`func (o *DirectHolderItem) HasPercentHeld() bool`

HasPercentHeld returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


