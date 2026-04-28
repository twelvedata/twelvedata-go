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

// checks if the GetETFsType200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsType200Response{}

// GetETFsType200Response struct for GetETFsType200Response
type GetETFsType200Response struct {
	// List of ETFs by market
	Result map[string][]string `json:"result"`
	// Status of the response
	Status string `json:"status"`
}

type _GetETFsType200Response GetETFsType200Response

// NewGetETFsType200Response instantiates a new GetETFsType200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsType200Response(result map[string][]string, status string) *GetETFsType200Response {
	this := GetETFsType200Response{}
	this.Result = result
	this.Status = status
	return &this
}

// NewGetETFsType200ResponseWithDefaults instantiates a new GetETFsType200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsType200ResponseWithDefaults() *GetETFsType200Response {
	this := GetETFsType200Response{}
	return &this
}

// GetResult returns the Result field value
func (o *GetETFsType200Response) GetResult() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetETFsType200Response) GetResultOk() (map[string][]string, bool) {
	if o == nil {
		return map[string][]string{}, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *GetETFsType200Response) SetResult(v map[string][]string) {
	o.Result = v
}

// GetStatus returns the Status field value
func (o *GetETFsType200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetETFsType200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetETFsType200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetETFsType200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsType200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetETFsType200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetETFsType200Response := _GetETFsType200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetETFsType200Response)

	if err != nil {
		return err
	}

	*o = GetETFsType200Response(varGetETFsType200Response)

	return err
}

type NullableGetETFsType200Response struct {
	value *GetETFsType200Response
	isSet bool
}

func (v NullableGetETFsType200Response) Get() *GetETFsType200Response {
	return v.value
}

func (v *NullableGetETFsType200Response) Set(val *GetETFsType200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsType200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsType200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsType200Response(val *GetETFsType200Response) *NullableGetETFsType200Response {
	return &NullableGetETFsType200Response{value: val, isSet: true}
}

func (v NullableGetETFsType200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsType200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
