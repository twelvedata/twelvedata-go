# GetCountries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CountryResponseItem**](CountryResponseItem.md) | List of countries with their ISO codes, names, capital, and currency | 

## Methods

### NewGetCountries200Response

`func NewGetCountries200Response(data []CountryResponseItem, ) *GetCountries200Response`

NewGetCountries200Response instantiates a new GetCountries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCountries200ResponseWithDefaults

`func NewGetCountries200ResponseWithDefaults() *GetCountries200Response`

NewGetCountries200ResponseWithDefaults instantiates a new GetCountries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCountries200Response) GetData() []CountryResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCountries200Response) GetDataOk() (*[]CountryResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCountries200Response) SetData(v []CountryResponseItem)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


