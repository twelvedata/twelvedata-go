// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetMutualFundsWorldRisk200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldRisk200Response{}

// GetMutualFundsWorldRisk200Response struct for GetMutualFundsWorldRisk200Response
type GetMutualFundsWorldRisk200Response struct {
	MutualFund GetMutualFundsWorldRisk200ResponseMutualFund `json:"mutual_fund"`
	// Status of the response
	Status string `json:"status"`
}

type _GetMutualFundsWorldRisk200Response GetMutualFundsWorldRisk200Response

// NewGetMutualFundsWorldRisk200Response instantiates a new GetMutualFundsWorldRisk200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldRisk200Response(mutualFund GetMutualFundsWorldRisk200ResponseMutualFund, status string) *GetMutualFundsWorldRisk200Response {
	this := GetMutualFundsWorldRisk200Response{}
	this.MutualFund = mutualFund
	this.Status = status
	return &this
}

// NewGetMutualFundsWorldRisk200ResponseWithDefaults instantiates a new GetMutualFundsWorldRisk200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldRisk200ResponseWithDefaults() *GetMutualFundsWorldRisk200Response {
	this := GetMutualFundsWorldRisk200Response{}
	return &this
}

// GetMutualFund returns the MutualFund field value
func (o *GetMutualFundsWorldRisk200Response) GetMutualFund() GetMutualFundsWorldRisk200ResponseMutualFund {
	if o == nil {
		var ret GetMutualFundsWorldRisk200ResponseMutualFund
		return ret
	}

	return o.MutualFund
}

// GetMutualFundOk returns a tuple with the MutualFund field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldRisk200Response) GetMutualFundOk() (*GetMutualFundsWorldRisk200ResponseMutualFund, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MutualFund, true
}

// SetMutualFund sets field value
func (o *GetMutualFundsWorldRisk200Response) SetMutualFund(v GetMutualFundsWorldRisk200ResponseMutualFund) {
	o.MutualFund = v
}

// GetStatus returns the Status field value
func (o *GetMutualFundsWorldRisk200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldRisk200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetMutualFundsWorldRisk200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetMutualFundsWorldRisk200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldRisk200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mutual_fund"] = o.MutualFund
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetMutualFundsWorldRisk200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetMutualFundsWorldRisk200Response := _GetMutualFundsWorldRisk200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetMutualFundsWorldRisk200Response)

	if err != nil {
		return err
	}

	*o = GetMutualFundsWorldRisk200Response(varGetMutualFundsWorldRisk200Response)

	return err
}

type NullableGetMutualFundsWorldRisk200Response struct {
	value *GetMutualFundsWorldRisk200Response
	isSet bool
}

func (v NullableGetMutualFundsWorldRisk200Response) Get() *GetMutualFundsWorldRisk200Response {
	return v.value
}

func (v *NullableGetMutualFundsWorldRisk200Response) Set(val *GetMutualFundsWorldRisk200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldRisk200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldRisk200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldRisk200Response(val *GetMutualFundsWorldRisk200Response) *NullableGetMutualFundsWorldRisk200Response {
	return &NullableGetMutualFundsWorldRisk200Response{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldRisk200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldRisk200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
