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

// checks if the ApiParameterTooLongErrorResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiParameterTooLongErrorResponseBody{}

// ApiParameterTooLongErrorResponseBody struct for ApiParameterTooLongErrorResponseBody
type ApiParameterTooLongErrorResponseBody struct {
	// Error code
	Code int64 `json:"code"`
	// Error message
	Message string `json:"message"`
	// Error status
	Status string `json:"status"`
}

type _ApiParameterTooLongErrorResponseBody ApiParameterTooLongErrorResponseBody

// NewApiParameterTooLongErrorResponseBody instantiates a new ApiParameterTooLongErrorResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiParameterTooLongErrorResponseBody(code int64, message string, status string) *ApiParameterTooLongErrorResponseBody {
	this := ApiParameterTooLongErrorResponseBody{}
	this.Code = code
	this.Message = message
	this.Status = status
	return &this
}

// NewApiParameterTooLongErrorResponseBodyWithDefaults instantiates a new ApiParameterTooLongErrorResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiParameterTooLongErrorResponseBodyWithDefaults() *ApiParameterTooLongErrorResponseBody {
	this := ApiParameterTooLongErrorResponseBody{}
	return &this
}

// GetCode returns the Code field value
func (o *ApiParameterTooLongErrorResponseBody) GetCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ApiParameterTooLongErrorResponseBody) GetCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ApiParameterTooLongErrorResponseBody) SetCode(v int64) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ApiParameterTooLongErrorResponseBody) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ApiParameterTooLongErrorResponseBody) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ApiParameterTooLongErrorResponseBody) SetMessage(v string) {
	o.Message = v
}

// GetStatus returns the Status field value
func (o *ApiParameterTooLongErrorResponseBody) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiParameterTooLongErrorResponseBody) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiParameterTooLongErrorResponseBody) SetStatus(v string) {
	o.Status = v
}

func (o ApiParameterTooLongErrorResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiParameterTooLongErrorResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ApiParameterTooLongErrorResponseBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
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

	varApiParameterTooLongErrorResponseBody := _ApiParameterTooLongErrorResponseBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varApiParameterTooLongErrorResponseBody)

	if err != nil {
		return err
	}

	*o = ApiParameterTooLongErrorResponseBody(varApiParameterTooLongErrorResponseBody)

	return err
}

type NullableApiParameterTooLongErrorResponseBody struct {
	value *ApiParameterTooLongErrorResponseBody
	isSet bool
}

func (v NullableApiParameterTooLongErrorResponseBody) Get() *ApiParameterTooLongErrorResponseBody {
	return v.value
}

func (v *NullableApiParameterTooLongErrorResponseBody) Set(val *ApiParameterTooLongErrorResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableApiParameterTooLongErrorResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableApiParameterTooLongErrorResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiParameterTooLongErrorResponseBody(val *ApiParameterTooLongErrorResponseBody) *NullableApiParameterTooLongErrorResponseBody {
	return &NullableApiParameterTooLongErrorResponseBody{value: val, isSet: true}
}

func (v NullableApiParameterTooLongErrorResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiParameterTooLongErrorResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
