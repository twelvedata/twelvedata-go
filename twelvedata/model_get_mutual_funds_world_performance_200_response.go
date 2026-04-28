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

// checks if the GetMutualFundsWorldPerformance200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldPerformance200Response{}

// GetMutualFundsWorldPerformance200Response struct for GetMutualFundsWorldPerformance200Response
type GetMutualFundsWorldPerformance200Response struct {
	MutualFund GetMutualFundsWorldPerformance200ResponseMutualFund `json:"mutual_fund"`
	// Status of the response
	Status string `json:"status"`
}

type _GetMutualFundsWorldPerformance200Response GetMutualFundsWorldPerformance200Response

// NewGetMutualFundsWorldPerformance200Response instantiates a new GetMutualFundsWorldPerformance200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldPerformance200Response(mutualFund GetMutualFundsWorldPerformance200ResponseMutualFund, status string) *GetMutualFundsWorldPerformance200Response {
	this := GetMutualFundsWorldPerformance200Response{}
	this.MutualFund = mutualFund
	this.Status = status
	return &this
}

// NewGetMutualFundsWorldPerformance200ResponseWithDefaults instantiates a new GetMutualFundsWorldPerformance200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldPerformance200ResponseWithDefaults() *GetMutualFundsWorldPerformance200Response {
	this := GetMutualFundsWorldPerformance200Response{}
	return &this
}

// GetMutualFund returns the MutualFund field value
func (o *GetMutualFundsWorldPerformance200Response) GetMutualFund() GetMutualFundsWorldPerformance200ResponseMutualFund {
	if o == nil {
		var ret GetMutualFundsWorldPerformance200ResponseMutualFund
		return ret
	}

	return o.MutualFund
}

// GetMutualFundOk returns a tuple with the MutualFund field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldPerformance200Response) GetMutualFundOk() (*GetMutualFundsWorldPerformance200ResponseMutualFund, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MutualFund, true
}

// SetMutualFund sets field value
func (o *GetMutualFundsWorldPerformance200Response) SetMutualFund(v GetMutualFundsWorldPerformance200ResponseMutualFund) {
	o.MutualFund = v
}

// GetStatus returns the Status field value
func (o *GetMutualFundsWorldPerformance200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldPerformance200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetMutualFundsWorldPerformance200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetMutualFundsWorldPerformance200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldPerformance200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mutual_fund"] = o.MutualFund
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetMutualFundsWorldPerformance200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetMutualFundsWorldPerformance200Response := _GetMutualFundsWorldPerformance200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMutualFundsWorldPerformance200Response)

	if err != nil {
		return err
	}

	*o = GetMutualFundsWorldPerformance200Response(varGetMutualFundsWorldPerformance200Response)

	return err
}

type NullableGetMutualFundsWorldPerformance200Response struct {
	value *GetMutualFundsWorldPerformance200Response
	isSet bool
}

func (v NullableGetMutualFundsWorldPerformance200Response) Get() *GetMutualFundsWorldPerformance200Response {
	return v.value
}

func (v *NullableGetMutualFundsWorldPerformance200Response) Set(val *GetMutualFundsWorldPerformance200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldPerformance200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldPerformance200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldPerformance200Response(val *GetMutualFundsWorldPerformance200Response) *NullableGetMutualFundsWorldPerformance200Response {
	return &NullableGetMutualFundsWorldPerformance200Response{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldPerformance200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldPerformance200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
