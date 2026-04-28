# InstitutionalHolderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityName** | **string** | Refers to the legal name of the institution | 
**DateReported** | Pointer to **string** | Refers to date reported | [optional] 
**Shares** | Pointer to **int64** | Refers to the number of shares owned | [optional] 
**Value** | Pointer to **int64** | Total value of shares owned, calculated by multiplying &#x60;shares&#x60; by the current price | [optional] 
**PercentHeld** | Pointer to **float64** | Represents the percentage of shares outstanding that are owned by the financial institution | [optional] 

## Methods

### NewInstitutionalHolderItem

`func NewInstitutionalHolderItem(entityName string, ) *InstitutionalHolderItem`

NewInstitutionalHolderItem instantiates a new InstitutionalHolderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionalHolderItemWithDefaults

`func NewInstitutionalHolderItemWithDefaults() *InstitutionalHolderItem`

NewInstitutionalHolderItemWithDefaults instantiates a new InstitutionalHolderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityName

`func (o *InstitutionalHolderItem) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *InstitutionalHolderItem) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *InstitutionalHolderItem) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.


### GetDateReported

`func (o *InstitutionalHolderItem) GetDateReported() string`

GetDateReported returns the DateReported field if non-nil, zero value otherwise.

### GetDateReportedOk

`func (o *InstitutionalHolderItem) GetDateReportedOk() (*string, bool)`

GetDateReportedOk returns a tuple with the DateReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateReported

`func (o *InstitutionalHolderItem) SetDateReported(v string)`

SetDateReported sets DateReported field to given value.

### HasDateReported

`func (o *InstitutionalHolderItem) HasDateReported() bool`

HasDateReported returns a boolean if a field has been set.

### GetShares

`func (o *InstitutionalHolderItem) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *InstitutionalHolderItem) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *InstitutionalHolderItem) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *InstitutionalHolderItem) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetValue

`func (o *InstitutionalHolderItem) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InstitutionalHolderItem) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InstitutionalHolderItem) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *InstitutionalHolderItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPercentHeld

`func (o *InstitutionalHolderItem) GetPercentHeld() float64`

GetPercentHeld returns the PercentHeld field if non-nil, zero value otherwise.

### GetPercentHeldOk

`func (o *InstitutionalHolderItem) GetPercentHeldOk() (*float64, bool)`

GetPercentHeldOk returns a tuple with the PercentHeld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentHeld

`func (o *InstitutionalHolderItem) SetPercentHeld(v float64)`

SetPercentHeld sets PercentHeld field to given value.

### HasPercentHeld

`func (o *InstitutionalHolderItem) HasPercentHeld() bool`

HasPercentHeld returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


