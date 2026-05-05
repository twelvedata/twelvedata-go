// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ApiInternalServerErrorResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInternalServerErrorResponseBody{}

// ApiInternalServerErrorResponseBody struct for ApiInternalServerErrorResponseBody
type ApiInternalServerErrorResponseBody struct {
	// Error code
	Code int64 `json:"code"`
	// Error message
	Message string `json:"message"`
	// Error status
	Status string `json:"status"`
}

type _ApiInternalServerErrorResponseBody ApiInternalServerErrorResponseBody

// NewApiInternalServerErrorResponseBody instantiates a new ApiInternalServerErrorResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInternalServerErrorResponseBody(code int64, message string, status string) *ApiInternalServerErrorResponseBody {
	this := ApiInternalServerErrorResponseBody{}
	this.Code = code
	this.Message = message
	this.Status = status
	return &this
}

// NewApiInternalServerErrorResponseBodyWithDefaults instantiates a new ApiInternalServerErrorResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInternalServerErrorResponseBodyWithDefaults() *ApiInternalServerErrorResponseBody {
	this := ApiInternalServerErrorResponseBody{}
	return &this
}

// GetCode returns the Code field value
func (o *ApiInternalServerErrorResponseBody) GetCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ApiInternalServerErrorResponseBody) GetCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ApiInternalServerErrorResponseBody) SetCode(v int64) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ApiInternalServerErrorResponseBody) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ApiInternalServerErrorResponseBody) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ApiInternalServerErrorResponseBody) SetMessage(v string) {
	o.Message = v
}

// GetStatus returns the Status field value
func (o *ApiInternalServerErrorResponseBody) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiInternalServerErrorResponseBody) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiInternalServerErrorResponseBody) SetStatus(v string) {
	o.Status = v
}

func (o ApiInternalServerErrorResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInternalServerErrorResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ApiInternalServerErrorResponseBody) UnmarshalJSON(data []byte) (err error) {
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

	varApiInternalServerErrorResponseBody := _ApiInternalServerErrorResponseBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varApiInternalServerErrorResponseBody)

	if err != nil {
		return err
	}

	*o = ApiInternalServerErrorResponseBody(varApiInternalServerErrorResponseBody)

	return err
}

type NullableApiInternalServerErrorResponseBody struct {
	value *ApiInternalServerErrorResponseBody
	isSet bool
}

func (v NullableApiInternalServerErrorResponseBody) Get() *ApiInternalServerErrorResponseBody {
	return v.value
}

func (v *NullableApiInternalServerErrorResponseBody) Set(val *ApiInternalServerErrorResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInternalServerErrorResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInternalServerErrorResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInternalServerErrorResponseBody(val *ApiInternalServerErrorResponseBody) *NullableApiInternalServerErrorResponseBody {
	return &NullableApiInternalServerErrorResponseBody{value: val, isSet: true}
}

func (v NullableApiInternalServerErrorResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInternalServerErrorResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
