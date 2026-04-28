/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetMutualFundsWorldSustainability200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldSustainability200Response{}

// GetMutualFundsWorldSustainability200Response struct for GetMutualFundsWorldSustainability200Response
type GetMutualFundsWorldSustainability200Response struct {
	MutualFund GetMutualFundsWorldSustainability200ResponseMutualFund `json:"mutual_fund"`
	// Status of the response
	Status string `json:"status"`
}

type _GetMutualFundsWorldSustainability200Response GetMutualFundsWorldSustainability200Response

// NewGetMutualFundsWorldSustainability200Response instantiates a new GetMutualFundsWorldSustainability200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldSustainability200Response(mutualFund GetMutualFundsWorldSustainability200ResponseMutualFund, status string) *GetMutualFundsWorldSustainability200Response {
	this := GetMutualFundsWorldSustainability200Response{}
	this.MutualFund = mutualFund
	this.Status = status
	return &this
}

// NewGetMutualFundsWorldSustainability200ResponseWithDefaults instantiates a new GetMutualFundsWorldSustainability200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldSustainability200ResponseWithDefaults() *GetMutualFundsWorldSustainability200Response {
	this := GetMutualFundsWorldSustainability200Response{}
	return &this
}

// GetMutualFund returns the MutualFund field value
func (o *GetMutualFundsWorldSustainability200Response) GetMutualFund() GetMutualFundsWorldSustainability200ResponseMutualFund {
	if o == nil {
		var ret GetMutualFundsWorldSustainability200ResponseMutualFund
		return ret
	}

	return o.MutualFund
}

// GetMutualFundOk returns a tuple with the MutualFund field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldSustainability200Response) GetMutualFundOk() (*GetMutualFundsWorldSustainability200ResponseMutualFund, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MutualFund, true
}

// SetMutualFund sets field value
func (o *GetMutualFundsWorldSustainability200Response) SetMutualFund(v GetMutualFundsWorldSustainability200ResponseMutualFund) {
	o.MutualFund = v
}

// GetStatus returns the Status field value
func (o *GetMutualFundsWorldSustainability200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldSustainability200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetMutualFundsWorldSustainability200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetMutualFundsWorldSustainability200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldSustainability200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mutual_fund"] = o.MutualFund
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetMutualFundsWorldSustainability200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mutual_fund",
		"status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetMutualFundsWorldSustainability200Response := _GetMutualFundsWorldSustainability200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMutualFundsWorldSustainability200Response)

	if err != nil {
		return err
	}

	*o = GetMutualFundsWorldSustainability200Response(varGetMutualFundsWorldSustainability200Response)

	return err
}

type NullableGetMutualFundsWorldSustainability200Response struct {
	value *GetMutualFundsWorldSustainability200Response
	isSet bool
}

func (v NullableGetMutualFundsWorldSustainability200Response) Get() *GetMutualFundsWorldSustainability200Response {
	return v.value
}

func (v *NullableGetMutualFundsWorldSustainability200Response) Set(val *GetMutualFundsWorldSustainability200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldSustainability200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldSustainability200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldSustainability200Response(val *GetMutualFundsWorldSustainability200Response) *NullableGetMutualFundsWorldSustainability200Response {
	return &NullableGetMutualFundsWorldSustainability200Response{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldSustainability200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldSustainability200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
