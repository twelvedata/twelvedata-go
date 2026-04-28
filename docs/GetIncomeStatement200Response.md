# GetIncomeStatement200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetIncomeStatement200ResponseMeta**](GetIncomeStatement200ResponseMeta.md) |  | 
**IncomeStatement** | Pointer to [**[]IncomeStatementBlock**](IncomeStatementBlock.md) | Income statement data | [optional] 

## Methods

### NewGetIncomeStatement200Response

`func NewGetIncomeStatement200Response(meta GetIncomeStatement200ResponseMeta, ) *GetIncomeStatement200Response`

NewGetIncomeStatement200Response instantiates a new GetIncomeStatement200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIncomeStatement200ResponseWithDefaults

`func NewGetIncomeStatement200ResponseWithDefaults() *GetIncomeStatement200Response`

NewGetIncomeStatement200ResponseWithDefaults instantiates a new GetIncomeStatement200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetIncomeStatement200Response) GetMeta() GetIncomeStatement200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetIncomeStatement200Response) GetMetaOk() (*GetIncomeStatement200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetIncomeStatement200Response) SetMeta(v GetIncomeStatement200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetIncomeStatement

`func (o *GetIncomeStatement200Response) GetIncomeStatement() []IncomeStatementBlock`

GetIncomeStatement returns the IncomeStatement field if non-nil, zero value otherwise.

### GetIncomeStatementOk

`func (o *GetIncomeStatement200Response) GetIncomeStatementOk() (*[]IncomeStatementBlock, bool)`

GetIncomeStatementOk returns a tuple with the IncomeStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeStatement

`func (o *GetIncomeStatement200Response) SetIncomeStatement(v []IncomeStatementBlock)`

SetIncomeStatement sets IncomeStatement field to given value.

### HasIncomeStatement

`func (o *GetIncomeStatement200Response) HasIncomeStatement() bool`

HasIncomeStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


