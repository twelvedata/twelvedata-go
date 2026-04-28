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

// checks if the GetBonds200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBonds200Response{}

// GetBonds200Response struct for GetBonds200Response
type GetBonds200Response struct {
	Result *GetBonds200ResponseResult `json:"result,omitempty"`
	// Response status
	Status string `json:"status"`
}

type _GetBonds200Response GetBonds200Response

// NewGetBonds200Response instantiates a new GetBonds200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBonds200Response(status string) *GetBonds200Response {
	this := GetBonds200Response{}
	this.Status = status
	return &this
}

// NewGetBonds200ResponseWithDefaults instantiates a new GetBonds200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBonds200ResponseWithDefaults() *GetBonds200Response {
	this := GetBonds200Response{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetBonds200Response) GetResult() GetBonds200ResponseResult {
	if o == nil || IsNil(o.Result) {
		var ret GetBonds200ResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBonds200Response) GetResultOk() (*GetBonds200ResponseResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetBonds200Response) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given GetBonds200ResponseResult and assigns it to the Result field.
func (o *GetBonds200Response) SetResult(v GetBonds200ResponseResult) {
	o.Result = &v
}

// GetStatus returns the Status field value
func (o *GetBonds200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetBonds200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetBonds200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetBonds200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBonds200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetBonds200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetBonds200Response := _GetBonds200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetBonds200Response)

	if err != nil {
		return err
	}

	*o = GetBonds200Response(varGetBonds200Response)

	return err
}

type NullableGetBonds200Response struct {
	value *GetBonds200Response
	isSet bool
}

func (v NullableGetBonds200Response) Get() *GetBonds200Response {
	return v.value
}

func (v *NullableGetBonds200Response) Set(val *GetBonds200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBonds200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBonds200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBonds200Response(val *GetBonds200Response) *NullableGetBonds200Response {
	return &NullableGetBonds200Response{value: val, isSet: true}
}

func (v NullableGetBonds200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBonds200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
