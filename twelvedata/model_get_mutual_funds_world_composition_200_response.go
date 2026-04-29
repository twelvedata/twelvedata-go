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

// checks if the GetMutualFundsWorldComposition200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldComposition200Response{}

// GetMutualFundsWorldComposition200Response struct for GetMutualFundsWorldComposition200Response
type GetMutualFundsWorldComposition200Response struct {
	MutualFund GetMutualFundsWorldComposition200ResponseMutualFund `json:"mutual_fund"`
	// Status of the response
	Status string `json:"status"`
}

type _GetMutualFundsWorldComposition200Response GetMutualFundsWorldComposition200Response

// NewGetMutualFundsWorldComposition200Response instantiates a new GetMutualFundsWorldComposition200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldComposition200Response(mutualFund GetMutualFundsWorldComposition200ResponseMutualFund, status string) *GetMutualFundsWorldComposition200Response {
	this := GetMutualFundsWorldComposition200Response{}
	this.MutualFund = mutualFund
	this.Status = status
	return &this
}

// NewGetMutualFundsWorldComposition200ResponseWithDefaults instantiates a new GetMutualFundsWorldComposition200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldComposition200ResponseWithDefaults() *GetMutualFundsWorldComposition200Response {
	this := GetMutualFundsWorldComposition200Response{}
	return &this
}

// GetMutualFund returns the MutualFund field value
func (o *GetMutualFundsWorldComposition200Response) GetMutualFund() GetMutualFundsWorldComposition200ResponseMutualFund {
	if o == nil {
		var ret GetMutualFundsWorldComposition200ResponseMutualFund
		return ret
	}

	return o.MutualFund
}

// GetMutualFundOk returns a tuple with the MutualFund field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldComposition200Response) GetMutualFundOk() (*GetMutualFundsWorldComposition200ResponseMutualFund, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MutualFund, true
}

// SetMutualFund sets field value
func (o *GetMutualFundsWorldComposition200Response) SetMutualFund(v GetMutualFundsWorldComposition200ResponseMutualFund) {
	o.MutualFund = v
}

// GetStatus returns the Status field value
func (o *GetMutualFundsWorldComposition200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldComposition200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetMutualFundsWorldComposition200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetMutualFundsWorldComposition200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldComposition200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mutual_fund"] = o.MutualFund
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetMutualFundsWorldComposition200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetMutualFundsWorldComposition200Response := _GetMutualFundsWorldComposition200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetMutualFundsWorldComposition200Response)

	if err != nil {
		return err
	}

	*o = GetMutualFundsWorldComposition200Response(varGetMutualFundsWorldComposition200Response)

	return err
}

type NullableGetMutualFundsWorldComposition200Response struct {
	value *GetMutualFundsWorldComposition200Response
	isSet bool
}

func (v NullableGetMutualFundsWorldComposition200Response) Get() *GetMutualFundsWorldComposition200Response {
	return v.value
}

func (v *NullableGetMutualFundsWorldComposition200Response) Set(val *GetMutualFundsWorldComposition200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldComposition200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldComposition200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldComposition200Response(val *GetMutualFundsWorldComposition200Response) *NullableGetMutualFundsWorldComposition200Response {
	return &NullableGetMutualFundsWorldComposition200Response{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldComposition200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldComposition200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
