// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetCrossListings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCrossListings200Response{}

// GetCrossListings200Response struct for GetCrossListings200Response
type GetCrossListings200Response struct {
	Result *CrossListingsResult `json:"result,omitempty"`
}

// NewGetCrossListings200Response instantiates a new GetCrossListings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCrossListings200Response() *GetCrossListings200Response {
	this := GetCrossListings200Response{}
	return &this
}

// NewGetCrossListings200ResponseWithDefaults instantiates a new GetCrossListings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCrossListings200ResponseWithDefaults() *GetCrossListings200Response {
	this := GetCrossListings200Response{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetCrossListings200Response) GetResult() CrossListingsResult {
	if o == nil || IsNil(o.Result) {
		var ret CrossListingsResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossListings200Response) GetResultOk() (*CrossListingsResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetCrossListings200Response) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given CrossListingsResult and assigns it to the Result field.
func (o *GetCrossListings200Response) SetResult(v CrossListingsResult) {
	o.Result = &v
}

func (o GetCrossListings200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCrossListings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetCrossListings200Response struct {
	value *GetCrossListings200Response
	isSet bool
}

func (v NullableGetCrossListings200Response) Get() *GetCrossListings200Response {
	return v.value
}

func (v *NullableGetCrossListings200Response) Set(val *GetCrossListings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCrossListings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCrossListings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCrossListings200Response(val *GetCrossListings200Response) *NullableGetCrossListings200Response {
	return &NullableGetCrossListings200Response{value: val, isSet: true}
}

func (v NullableGetCrossListings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCrossListings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
