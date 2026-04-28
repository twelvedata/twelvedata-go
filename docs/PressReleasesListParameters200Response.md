# PressReleasesListParameters200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PressReleases** | [**[]PressRelease**](PressRelease.md) | List of press releases | 
**Status** | **string** | Response status | 

## Methods

### NewPressReleasesListParameters200Response

`func NewPressReleasesListParameters200Response(pressReleases []PressRelease, status string, ) *PressReleasesListParameters200Response`

NewPressReleasesListParameters200Response instantiates a new PressReleasesListParameters200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPressReleasesListParameters200ResponseWithDefaults

`func NewPressReleasesListParameters200ResponseWithDefaults() *PressReleasesListParameters200Response`

NewPressReleasesListParameters200ResponseWithDefaults instantiates a new PressReleasesListParameters200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPressReleases

`func (o *PressReleasesListParameters200Response) GetPressReleases() []PressRelease`

GetPressReleases returns the PressReleases field if non-nil, zero value otherwise.

### GetPressReleasesOk

`func (o *PressReleasesListParameters200Response) GetPressReleasesOk() (*[]PressRelease, bool)`

GetPressReleasesOk returns a tuple with the PressReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPressReleases

`func (o *PressReleasesListParameters200Response) SetPressReleases(v []PressRelease)`

SetPressReleases sets PressReleases field to given value.


### GetStatus

`func (o *PressReleasesListParameters200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PressReleasesListParameters200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PressReleasesListParameters200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


