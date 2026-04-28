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

// checks if the GetMutualFundsFamily200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsFamily200Response{}

// GetMutualFundsFamily200Response struct for GetMutualFundsFamily200Response
type GetMutualFundsFamily200Response struct {
	// List of fund families by country
	Result map[string][]string `json:"result"`
	// Response status
	Status string `json:"status"`
}

type _GetMutualFundsFamily200Response GetMutualFundsFamily200Response

// NewGetMutualFundsFamily200Response instantiates a new GetMutualFundsFamily200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsFamily200Response(result map[string][]string, status string) *GetMutualFundsFamily200Response {
	this := GetMutualFundsFamily200Response{}
	this.Result = result
	this.Status = status
	return &this
}

// NewGetMutualFundsFamily200ResponseWithDefaults instantiates a new GetMutualFundsFamily200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsFamily200ResponseWithDefaults() *GetMutualFundsFamily200Response {
	this := GetMutualFundsFamily200Response{}
	return &this
}

// GetResult returns the Result field value
func (o *GetMutualFundsFamily200Response) GetResult() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsFamily200Response) GetResultOk() (map[string][]string, bool) {
	if o == nil {
		return map[string][]string{}, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *GetMutualFundsFamily200Response) SetResult(v map[string][]string) {
	o.Result = v
}

// GetStatus returns the Status field value
func (o *GetMutualFundsFamily200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsFamily200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetMutualFundsFamily200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetMutualFundsFamily200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsFamily200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetMutualFundsFamily200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varGetMutualFundsFamily200Response := _GetMutualFundsFamily200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMutualFundsFamily200Response)

	if err != nil {
		return err
	}

	*o = GetMutualFundsFamily200Response(varGetMutualFundsFamily200Response)

	return err
}

type NullableGetMutualFundsFamily200Response struct {
	value *GetMutualFundsFamily200Response
	isSet bool
}

func (v NullableGetMutualFundsFamily200Response) Get() *GetMutualFundsFamily200Response {
	return v.value
}

func (v *NullableGetMutualFundsFamily200Response) Set(val *GetMutualFundsFamily200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsFamily200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsFamily200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsFamily200Response(val *GetMutualFundsFamily200Response) *NullableGetMutualFundsFamily200Response {
	return &NullableGetMutualFundsFamily200Response{value: val, isSet: true}
}

func (v NullableGetMutualFundsFamily200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsFamily200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
