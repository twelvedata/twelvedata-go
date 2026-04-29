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

// checks if the GetETFsFamily200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsFamily200Response{}

// GetETFsFamily200Response struct for GetETFsFamily200Response
type GetETFsFamily200Response struct {
	// List of ETFs by country
	Result map[string][]string `json:"result"`
	// Status of the response
	Status string `json:"status"`
}

type _GetETFsFamily200Response GetETFsFamily200Response

// NewGetETFsFamily200Response instantiates a new GetETFsFamily200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsFamily200Response(result map[string][]string, status string) *GetETFsFamily200Response {
	this := GetETFsFamily200Response{}
	this.Result = result
	this.Status = status
	return &this
}

// NewGetETFsFamily200ResponseWithDefaults instantiates a new GetETFsFamily200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsFamily200ResponseWithDefaults() *GetETFsFamily200Response {
	this := GetETFsFamily200Response{}
	return &this
}

// GetResult returns the Result field value
func (o *GetETFsFamily200Response) GetResult() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetETFsFamily200Response) GetResultOk() (map[string][]string, bool) {
	if o == nil {
		return map[string][]string{}, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *GetETFsFamily200Response) SetResult(v map[string][]string) {
	o.Result = v
}

// GetStatus returns the Status field value
func (o *GetETFsFamily200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetETFsFamily200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetETFsFamily200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetETFsFamily200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsFamily200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetETFsFamily200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetETFsFamily200Response := _GetETFsFamily200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetETFsFamily200Response)

	if err != nil {
		return err
	}

	*o = GetETFsFamily200Response(varGetETFsFamily200Response)

	return err
}

type NullableGetETFsFamily200Response struct {
	value *GetETFsFamily200Response
	isSet bool
}

func (v NullableGetETFsFamily200Response) Get() *GetETFsFamily200Response {
	return v.value
}

func (v *NullableGetETFsFamily200Response) Set(val *GetETFsFamily200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsFamily200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsFamily200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsFamily200Response(val *GetETFsFamily200Response) *NullableGetETFsFamily200Response {
	return &NullableGetETFsFamily200Response{value: val, isSet: true}
}

func (v NullableGetETFsFamily200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsFamily200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
