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

// checks if the GetMutualFundsWorldPurchaseInfo200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldPurchaseInfo200Response{}

// GetMutualFundsWorldPurchaseInfo200Response struct for GetMutualFundsWorldPurchaseInfo200Response
type GetMutualFundsWorldPurchaseInfo200Response struct {
	MutualFund GetMutualFundsWorldPurchaseInfo200ResponseMutualFund `json:"mutual_fund"`
	// Status of the response
	Status string `json:"status"`
}

type _GetMutualFundsWorldPurchaseInfo200Response GetMutualFundsWorldPurchaseInfo200Response

// NewGetMutualFundsWorldPurchaseInfo200Response instantiates a new GetMutualFundsWorldPurchaseInfo200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldPurchaseInfo200Response(mutualFund GetMutualFundsWorldPurchaseInfo200ResponseMutualFund, status string) *GetMutualFundsWorldPurchaseInfo200Response {
	this := GetMutualFundsWorldPurchaseInfo200Response{}
	this.MutualFund = mutualFund
	this.Status = status
	return &this
}

// NewGetMutualFundsWorldPurchaseInfo200ResponseWithDefaults instantiates a new GetMutualFundsWorldPurchaseInfo200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldPurchaseInfo200ResponseWithDefaults() *GetMutualFundsWorldPurchaseInfo200Response {
	this := GetMutualFundsWorldPurchaseInfo200Response{}
	return &this
}

// GetMutualFund returns the MutualFund field value
func (o *GetMutualFundsWorldPurchaseInfo200Response) GetMutualFund() GetMutualFundsWorldPurchaseInfo200ResponseMutualFund {
	if o == nil {
		var ret GetMutualFundsWorldPurchaseInfo200ResponseMutualFund
		return ret
	}

	return o.MutualFund
}

// GetMutualFundOk returns a tuple with the MutualFund field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldPurchaseInfo200Response) GetMutualFundOk() (*GetMutualFundsWorldPurchaseInfo200ResponseMutualFund, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MutualFund, true
}

// SetMutualFund sets field value
func (o *GetMutualFundsWorldPurchaseInfo200Response) SetMutualFund(v GetMutualFundsWorldPurchaseInfo200ResponseMutualFund) {
	o.MutualFund = v
}

// GetStatus returns the Status field value
func (o *GetMutualFundsWorldPurchaseInfo200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldPurchaseInfo200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetMutualFundsWorldPurchaseInfo200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetMutualFundsWorldPurchaseInfo200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldPurchaseInfo200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mutual_fund"] = o.MutualFund
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetMutualFundsWorldPurchaseInfo200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetMutualFundsWorldPurchaseInfo200Response := _GetMutualFundsWorldPurchaseInfo200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetMutualFundsWorldPurchaseInfo200Response)

	if err != nil {
		return err
	}

	*o = GetMutualFundsWorldPurchaseInfo200Response(varGetMutualFundsWorldPurchaseInfo200Response)

	return err
}

type NullableGetMutualFundsWorldPurchaseInfo200Response struct {
	value *GetMutualFundsWorldPurchaseInfo200Response
	isSet bool
}

func (v NullableGetMutualFundsWorldPurchaseInfo200Response) Get() *GetMutualFundsWorldPurchaseInfo200Response {
	return v.value
}

func (v *NullableGetMutualFundsWorldPurchaseInfo200Response) Set(val *GetMutualFundsWorldPurchaseInfo200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldPurchaseInfo200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldPurchaseInfo200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldPurchaseInfo200Response(val *GetMutualFundsWorldPurchaseInfo200Response) *NullableGetMutualFundsWorldPurchaseInfo200Response {
	return &NullableGetMutualFundsWorldPurchaseInfo200Response{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldPurchaseInfo200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldPurchaseInfo200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
