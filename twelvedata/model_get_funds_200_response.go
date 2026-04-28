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

// checks if the GetFunds200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFunds200Response{}

// GetFunds200Response struct for GetFunds200Response
type GetFunds200Response struct {
	Result GetFunds200ResponseResult `json:"result"`
	// Status of the response
	Status string `json:"status"`
}

type _GetFunds200Response GetFunds200Response

// NewGetFunds200Response instantiates a new GetFunds200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFunds200Response(result GetFunds200ResponseResult, status string) *GetFunds200Response {
	this := GetFunds200Response{}
	this.Result = result
	this.Status = status
	return &this
}

// NewGetFunds200ResponseWithDefaults instantiates a new GetFunds200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFunds200ResponseWithDefaults() *GetFunds200Response {
	this := GetFunds200Response{}
	return &this
}

// GetResult returns the Result field value
func (o *GetFunds200Response) GetResult() GetFunds200ResponseResult {
	if o == nil {
		var ret GetFunds200ResponseResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetFunds200Response) GetResultOk() (*GetFunds200ResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *GetFunds200Response) SetResult(v GetFunds200ResponseResult) {
	o.Result = v
}

// GetStatus returns the Status field value
func (o *GetFunds200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetFunds200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetFunds200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetFunds200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFunds200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetFunds200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetFunds200Response := _GetFunds200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFunds200Response)

	if err != nil {
		return err
	}

	*o = GetFunds200Response(varGetFunds200Response)

	return err
}

type NullableGetFunds200Response struct {
	value *GetFunds200Response
	isSet bool
}

func (v NullableGetFunds200Response) Get() *GetFunds200Response {
	return v.value
}

func (v *NullableGetFunds200Response) Set(val *GetFunds200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFunds200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFunds200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFunds200Response(val *GetFunds200Response) *NullableGetFunds200Response {
	return &NullableGetFunds200Response{value: val, isSet: true}
}

func (v NullableGetFunds200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFunds200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
