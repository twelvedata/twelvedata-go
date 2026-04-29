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

// checks if the ApiNotFoundErrorResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiNotFoundErrorResponseBody{}

// ApiNotFoundErrorResponseBody struct for ApiNotFoundErrorResponseBody
type ApiNotFoundErrorResponseBody struct {
	// Error code
	Code int64 `json:"code"`
	// Error message
	Message string `json:"message"`
	// Error status
	Status string `json:"status"`
}

type _ApiNotFoundErrorResponseBody ApiNotFoundErrorResponseBody

// NewApiNotFoundErrorResponseBody instantiates a new ApiNotFoundErrorResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNotFoundErrorResponseBody(code int64, message string, status string) *ApiNotFoundErrorResponseBody {
	this := ApiNotFoundErrorResponseBody{}
	this.Code = code
	this.Message = message
	this.Status = status
	return &this
}

// NewApiNotFoundErrorResponseBodyWithDefaults instantiates a new ApiNotFoundErrorResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNotFoundErrorResponseBodyWithDefaults() *ApiNotFoundErrorResponseBody {
	this := ApiNotFoundErrorResponseBody{}
	return &this
}

// GetCode returns the Code field value
func (o *ApiNotFoundErrorResponseBody) GetCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ApiNotFoundErrorResponseBody) GetCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ApiNotFoundErrorResponseBody) SetCode(v int64) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ApiNotFoundErrorResponseBody) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ApiNotFoundErrorResponseBody) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ApiNotFoundErrorResponseBody) SetMessage(v string) {
	o.Message = v
}

// GetStatus returns the Status field value
func (o *ApiNotFoundErrorResponseBody) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiNotFoundErrorResponseBody) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiNotFoundErrorResponseBody) SetStatus(v string) {
	o.Status = v
}

func (o ApiNotFoundErrorResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiNotFoundErrorResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ApiNotFoundErrorResponseBody) UnmarshalJSON(data []byte) (err error) {
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

	varApiNotFoundErrorResponseBody := _ApiNotFoundErrorResponseBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varApiNotFoundErrorResponseBody)

	if err != nil {
		return err
	}

	*o = ApiNotFoundErrorResponseBody(varApiNotFoundErrorResponseBody)

	return err
}

type NullableApiNotFoundErrorResponseBody struct {
	value *ApiNotFoundErrorResponseBody
	isSet bool
}

func (v NullableApiNotFoundErrorResponseBody) Get() *ApiNotFoundErrorResponseBody {
	return v.value
}

func (v *NullableApiNotFoundErrorResponseBody) Set(val *ApiNotFoundErrorResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNotFoundErrorResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNotFoundErrorResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNotFoundErrorResponseBody(val *ApiNotFoundErrorResponseBody) *NullableApiNotFoundErrorResponseBody {
	return &NullableApiNotFoundErrorResponseBody{value: val, isSet: true}
}

func (v NullableApiNotFoundErrorResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNotFoundErrorResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
