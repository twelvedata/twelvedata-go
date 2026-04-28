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

// checks if the GetMutualFundsList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsList200Response{}

// GetMutualFundsList200Response struct for GetMutualFundsList200Response
type GetMutualFundsList200Response struct {
	Result *GetMutualFundsList200ResponseResult `json:"result,omitempty"`
	// Response status
	Status string `json:"status"`
}

type _GetMutualFundsList200Response GetMutualFundsList200Response

// NewGetMutualFundsList200Response instantiates a new GetMutualFundsList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsList200Response(status string) *GetMutualFundsList200Response {
	this := GetMutualFundsList200Response{}
	this.Status = status
	return &this
}

// NewGetMutualFundsList200ResponseWithDefaults instantiates a new GetMutualFundsList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsList200ResponseWithDefaults() *GetMutualFundsList200Response {
	this := GetMutualFundsList200Response{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetMutualFundsList200Response) GetResult() GetMutualFundsList200ResponseResult {
	if o == nil || IsNil(o.Result) {
		var ret GetMutualFundsList200ResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsList200Response) GetResultOk() (*GetMutualFundsList200ResponseResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetMutualFundsList200Response) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given GetMutualFundsList200ResponseResult and assigns it to the Result field.
func (o *GetMutualFundsList200Response) SetResult(v GetMutualFundsList200ResponseResult) {
	o.Result = &v
}

// GetStatus returns the Status field value
func (o *GetMutualFundsList200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsList200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetMutualFundsList200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetMutualFundsList200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetMutualFundsList200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varGetMutualFundsList200Response := _GetMutualFundsList200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMutualFundsList200Response)

	if err != nil {
		return err
	}

	*o = GetMutualFundsList200Response(varGetMutualFundsList200Response)

	return err
}

type NullableGetMutualFundsList200Response struct {
	value *GetMutualFundsList200Response
	isSet bool
}

func (v NullableGetMutualFundsList200Response) Get() *GetMutualFundsList200Response {
	return v.value
}

func (v *NullableGetMutualFundsList200Response) Set(val *GetMutualFundsList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsList200Response(val *GetMutualFundsList200Response) *NullableGetMutualFundsList200Response {
	return &NullableGetMutualFundsList200Response{value: val, isSet: true}
}

func (v NullableGetMutualFundsList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
