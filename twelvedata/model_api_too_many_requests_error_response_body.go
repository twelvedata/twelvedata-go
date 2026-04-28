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

// checks if the ApiTooManyRequestsErrorResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTooManyRequestsErrorResponseBody{}

// ApiTooManyRequestsErrorResponseBody struct for ApiTooManyRequestsErrorResponseBody
type ApiTooManyRequestsErrorResponseBody struct {
	// Error code
	Code int64 `json:"code"`
	// Error message
	Message string `json:"message"`
	// Error status
	Status string `json:"status"`
}

type _ApiTooManyRequestsErrorResponseBody ApiTooManyRequestsErrorResponseBody

// NewApiTooManyRequestsErrorResponseBody instantiates a new ApiTooManyRequestsErrorResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTooManyRequestsErrorResponseBody(code int64, message string, status string) *ApiTooManyRequestsErrorResponseBody {
	this := ApiTooManyRequestsErrorResponseBody{}
	this.Code = code
	this.Message = message
	this.Status = status
	return &this
}

// NewApiTooManyRequestsErrorResponseBodyWithDefaults instantiates a new ApiTooManyRequestsErrorResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTooManyRequestsErrorResponseBodyWithDefaults() *ApiTooManyRequestsErrorResponseBody {
	this := ApiTooManyRequestsErrorResponseBody{}
	return &this
}

// GetCode returns the Code field value
func (o *ApiTooManyRequestsErrorResponseBody) GetCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ApiTooManyRequestsErrorResponseBody) GetCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ApiTooManyRequestsErrorResponseBody) SetCode(v int64) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ApiTooManyRequestsErrorResponseBody) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ApiTooManyRequestsErrorResponseBody) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ApiTooManyRequestsErrorResponseBody) SetMessage(v string) {
	o.Message = v
}

// GetStatus returns the Status field value
func (o *ApiTooManyRequestsErrorResponseBody) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiTooManyRequestsErrorResponseBody) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiTooManyRequestsErrorResponseBody) SetStatus(v string) {
	o.Status = v
}

func (o ApiTooManyRequestsErrorResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTooManyRequestsErrorResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ApiTooManyRequestsErrorResponseBody) UnmarshalJSON(data []byte) (err error) {
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

	varApiTooManyRequestsErrorResponseBody := _ApiTooManyRequestsErrorResponseBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTooManyRequestsErrorResponseBody)

	if err != nil {
		return err
	}

	*o = ApiTooManyRequestsErrorResponseBody(varApiTooManyRequestsErrorResponseBody)

	return err
}

type NullableApiTooManyRequestsErrorResponseBody struct {
	value *ApiTooManyRequestsErrorResponseBody
	isSet bool
}

func (v NullableApiTooManyRequestsErrorResponseBody) Get() *ApiTooManyRequestsErrorResponseBody {
	return v.value
}

func (v *NullableApiTooManyRequestsErrorResponseBody) Set(val *ApiTooManyRequestsErrorResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTooManyRequestsErrorResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTooManyRequestsErrorResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTooManyRequestsErrorResponseBody(val *ApiTooManyRequestsErrorResponseBody) *NullableApiTooManyRequestsErrorResponseBody {
	return &NullableApiTooManyRequestsErrorResponseBody{value: val, isSet: true}
}

func (v NullableApiTooManyRequestsErrorResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTooManyRequestsErrorResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
