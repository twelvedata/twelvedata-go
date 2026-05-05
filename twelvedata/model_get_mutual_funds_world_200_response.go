// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetMutualFundsWorld200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200Response{}

// GetMutualFundsWorld200Response struct for GetMutualFundsWorld200Response
type GetMutualFundsWorld200Response struct {
	MutualFund GetMutualFundsWorld200ResponseMutualFund `json:"mutual_fund"`
	// Status of the response
	Status string `json:"status"`
}

type _GetMutualFundsWorld200Response GetMutualFundsWorld200Response

// NewGetMutualFundsWorld200Response instantiates a new GetMutualFundsWorld200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200Response(mutualFund GetMutualFundsWorld200ResponseMutualFund, status string) *GetMutualFundsWorld200Response {
	this := GetMutualFundsWorld200Response{}
	this.MutualFund = mutualFund
	this.Status = status
	return &this
}

// NewGetMutualFundsWorld200ResponseWithDefaults instantiates a new GetMutualFundsWorld200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseWithDefaults() *GetMutualFundsWorld200Response {
	this := GetMutualFundsWorld200Response{}
	return &this
}

// GetMutualFund returns the MutualFund field value
func (o *GetMutualFundsWorld200Response) GetMutualFund() GetMutualFundsWorld200ResponseMutualFund {
	if o == nil {
		var ret GetMutualFundsWorld200ResponseMutualFund
		return ret
	}

	return o.MutualFund
}

// GetMutualFundOk returns a tuple with the MutualFund field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200Response) GetMutualFundOk() (*GetMutualFundsWorld200ResponseMutualFund, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MutualFund, true
}

// SetMutualFund sets field value
func (o *GetMutualFundsWorld200Response) SetMutualFund(v GetMutualFundsWorld200ResponseMutualFund) {
	o.MutualFund = v
}

// GetStatus returns the Status field value
func (o *GetMutualFundsWorld200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetMutualFundsWorld200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetMutualFundsWorld200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mutual_fund"] = o.MutualFund
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetMutualFundsWorld200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetMutualFundsWorld200Response := _GetMutualFundsWorld200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetMutualFundsWorld200Response)

	if err != nil {
		return err
	}

	*o = GetMutualFundsWorld200Response(varGetMutualFundsWorld200Response)

	return err
}

type NullableGetMutualFundsWorld200Response struct {
	value *GetMutualFundsWorld200Response
	isSet bool
}

func (v NullableGetMutualFundsWorld200Response) Get() *GetMutualFundsWorld200Response {
	return v.value
}

func (v *NullableGetMutualFundsWorld200Response) Set(val *GetMutualFundsWorld200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200Response(val *GetMutualFundsWorld200Response) *NullableGetMutualFundsWorld200Response {
	return &NullableGetMutualFundsWorld200Response{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
