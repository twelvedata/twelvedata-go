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

// checks if the GetETFsList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsList200Response{}

// GetETFsList200Response struct for GetETFsList200Response
type GetETFsList200Response struct {
	Result GetETFsList200ResponseResult `json:"result"`
	// Status of the response
	Status string `json:"status"`
}

type _GetETFsList200Response GetETFsList200Response

// NewGetETFsList200Response instantiates a new GetETFsList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsList200Response(result GetETFsList200ResponseResult, status string) *GetETFsList200Response {
	this := GetETFsList200Response{}
	this.Result = result
	this.Status = status
	return &this
}

// NewGetETFsList200ResponseWithDefaults instantiates a new GetETFsList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsList200ResponseWithDefaults() *GetETFsList200Response {
	this := GetETFsList200Response{}
	return &this
}

// GetResult returns the Result field value
func (o *GetETFsList200Response) GetResult() GetETFsList200ResponseResult {
	if o == nil {
		var ret GetETFsList200ResponseResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetETFsList200Response) GetResultOk() (*GetETFsList200ResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *GetETFsList200Response) SetResult(v GetETFsList200ResponseResult) {
	o.Result = v
}

// GetStatus returns the Status field value
func (o *GetETFsList200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetETFsList200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetETFsList200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetETFsList200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetETFsList200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetETFsList200Response := _GetETFsList200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetETFsList200Response)

	if err != nil {
		return err
	}

	*o = GetETFsList200Response(varGetETFsList200Response)

	return err
}

type NullableGetETFsList200Response struct {
	value *GetETFsList200Response
	isSet bool
}

func (v NullableGetETFsList200Response) Get() *GetETFsList200Response {
	return v.value
}

func (v *NullableGetETFsList200Response) Set(val *GetETFsList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsList200Response(val *GetETFsList200Response) *NullableGetETFsList200Response {
	return &NullableGetETFsList200Response{value: val, isSet: true}
}

func (v NullableGetETFsList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
