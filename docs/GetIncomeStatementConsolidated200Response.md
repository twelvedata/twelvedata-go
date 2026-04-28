# GetIncomeStatementConsolidated200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomeStatement** | Pointer to [**[]IncomeStatementItem**](IncomeStatementItem.md) | Income statement data | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetIncomeStatementConsolidated200Response

`func NewGetIncomeStatementConsolidated200Response() *GetIncomeStatementConsolidated200Response`

NewGetIncomeStatementConsolidated200Response instantiates a new GetIncomeStatementConsolidated200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIncomeStatementConsolidated200ResponseWithDefaults

`func NewGetIncomeStatementConsolidated200ResponseWithDefaults() *GetIncomeStatementConsolidated200Response`

NewGetIncomeStatementConsolidated200ResponseWithDefaults instantiates a new GetIncomeStatementConsolidated200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomeStatement

`func (o *GetIncomeStatementConsolidated200Response) GetIncomeStatement() []IncomeStatementItem`

GetIncomeStatement returns the IncomeStatement field if non-nil, zero value otherwise.

### GetIncomeStatementOk

`func (o *GetIncomeStatementConsolidated200Response) GetIncomeStatementOk() (*[]IncomeStatementItem, bool)`

GetIncomeStatementOk returns a tuple with the IncomeStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeStatement

`func (o *GetIncomeStatementConsolidated200Response) SetIncomeStatement(v []IncomeStatementItem)`

SetIncomeStatement sets IncomeStatement field to given value.

### HasIncomeStatement

`func (o *GetIncomeStatementConsolidated200Response) HasIncomeStatement() bool`

HasIncomeStatement returns a boolean if a field has been set.

### GetStatus

`func (o *GetIncomeStatementConsolidated200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIncomeStatementConsolidated200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIncomeStatementConsolidated200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIncomeStatementConsolidated200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


