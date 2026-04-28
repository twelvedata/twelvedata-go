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

// checks if the GetMutualFundsWorldSummary200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorldSummary200Response{}

// GetMutualFundsWorldSummary200Response struct for GetMutualFundsWorldSummary200Response
type GetMutualFundsWorldSummary200Response struct {
	MutualFund GetMutualFundsWorldSummary200ResponseMutualFund `json:"mutual_fund"`
	// Status of the response
	Status string `json:"status"`
}

type _GetMutualFundsWorldSummary200Response GetMutualFundsWorldSummary200Response

// NewGetMutualFundsWorldSummary200Response instantiates a new GetMutualFundsWorldSummary200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorldSummary200Response(mutualFund GetMutualFundsWorldSummary200ResponseMutualFund, status string) *GetMutualFundsWorldSummary200Response {
	this := GetMutualFundsWorldSummary200Response{}
	this.MutualFund = mutualFund
	this.Status = status
	return &this
}

// NewGetMutualFundsWorldSummary200ResponseWithDefaults instantiates a new GetMutualFundsWorldSummary200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorldSummary200ResponseWithDefaults() *GetMutualFundsWorldSummary200Response {
	this := GetMutualFundsWorldSummary200Response{}
	return &this
}

// GetMutualFund returns the MutualFund field value
func (o *GetMutualFundsWorldSummary200Response) GetMutualFund() GetMutualFundsWorldSummary200ResponseMutualFund {
	if o == nil {
		var ret GetMutualFundsWorldSummary200ResponseMutualFund
		return ret
	}

	return o.MutualFund
}

// GetMutualFundOk returns a tuple with the MutualFund field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldSummary200Response) GetMutualFundOk() (*GetMutualFundsWorldSummary200ResponseMutualFund, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MutualFund, true
}

// SetMutualFund sets field value
func (o *GetMutualFundsWorldSummary200Response) SetMutualFund(v GetMutualFundsWorldSummary200ResponseMutualFund) {
	o.MutualFund = v
}

// GetStatus returns the Status field value
func (o *GetMutualFundsWorldSummary200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorldSummary200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetMutualFundsWorldSummary200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetMutualFundsWorldSummary200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorldSummary200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mutual_fund"] = o.MutualFund
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetMutualFundsWorldSummary200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetMutualFundsWorldSummary200Response := _GetMutualFundsWorldSummary200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMutualFundsWorldSummary200Response)

	if err != nil {
		return err
	}

	*o = GetMutualFundsWorldSummary200Response(varGetMutualFundsWorldSummary200Response)

	return err
}

type NullableGetMutualFundsWorldSummary200Response struct {
	value *GetMutualFundsWorldSummary200Response
	isSet bool
}

func (v NullableGetMutualFundsWorldSummary200Response) Get() *GetMutualFundsWorldSummary200Response {
	return v.value
}

func (v *NullableGetMutualFundsWorldSummary200Response) Set(val *GetMutualFundsWorldSummary200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorldSummary200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorldSummary200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorldSummary200Response(val *GetMutualFundsWorldSummary200Response) *NullableGetMutualFundsWorldSummary200Response {
	return &NullableGetMutualFundsWorldSummary200Response{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorldSummary200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorldSummary200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
