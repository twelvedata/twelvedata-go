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

// checks if the GetMutualFundsWorldRatings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldRatings200Response{}

// GetMutualFundsWorldRatings200Response struct for GetMutualFundsWorldRatings200Response
type GetMutualFundsWorldRatings200Response struct {
	MutualFund GetMutualFundsWorldRatings200ResponseMutualFund `json:"mutual_fund"`
	// Status of the response
	Status string `json:"status"`
}

type _GetMutualFundsWorldRatings200Response GetMutualFundsWorldRatings200Response

// NewGetMutualFundsWorldRatings200Response instantiates a new GetMutualFundsWorldRatings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldRatings200Response(mutualFund GetMutualFundsWorldRatings200ResponseMutualFund, status string) *GetMutualFundsWorldRatings200Response {
	this := GetMutualFundsWorldRatings200Response{}
	this.MutualFund = mutualFund
	this.Status = status
	return &this
}

// NewGetMutualFundsWorldRatings200ResponseWithDefaults instantiates a new GetMutualFundsWorldRatings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldRatings200ResponseWithDefaults() *GetMutualFundsWorldRatings200Response {
	this := GetMutualFundsWorldRatings200Response{}
	return &this
}

// GetMutualFund returns the MutualFund field value
func (o *GetMutualFundsWorldRatings200Response) GetMutualFund() GetMutualFundsWorldRatings200ResponseMutualFund {
	if o == nil {
		var ret GetMutualFundsWorldRatings200ResponseMutualFund
		return ret
	}

	return o.MutualFund
}

// GetMutualFundOk returns a tuple with the MutualFund field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldRatings200Response) GetMutualFundOk() (*GetMutualFundsWorldRatings200ResponseMutualFund, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MutualFund, true
}

// SetMutualFund sets field value
func (o *GetMutualFundsWorldRatings200Response) SetMutualFund(v GetMutualFundsWorldRatings200ResponseMutualFund) {
	o.MutualFund = v
}

// GetStatus returns the Status field value
func (o *GetMutualFundsWorldRatings200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldRatings200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetMutualFundsWorldRatings200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetMutualFundsWorldRatings200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldRatings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mutual_fund"] = o.MutualFund
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetMutualFundsWorldRatings200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetMutualFundsWorldRatings200Response := _GetMutualFundsWorldRatings200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMutualFundsWorldRatings200Response)

	if err != nil {
		return err
	}

	*o = GetMutualFundsWorldRatings200Response(varGetMutualFundsWorldRatings200Response)

	return err
}

type NullableGetMutualFundsWorldRatings200Response struct {
	value *GetMutualFundsWorldRatings200Response
	isSet bool
}

func (v NullableGetMutualFundsWorldRatings200Response) Get() *GetMutualFundsWorldRatings200Response {
	return v.value
}

func (v *NullableGetMutualFundsWorldRatings200Response) Set(val *GetMutualFundsWorldRatings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldRatings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldRatings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldRatings200Response(val *GetMutualFundsWorldRatings200Response) *NullableGetMutualFundsWorldRatings200Response {
	return &NullableGetMutualFundsWorldRatings200Response{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldRatings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldRatings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
