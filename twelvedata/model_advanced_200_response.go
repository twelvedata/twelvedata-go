/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the Advanced200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Advanced200Response{}

// Advanced200Response struct for Advanced200Response
type Advanced200Response struct {
	// HTTP status code
	Code *int64 `json:"code,omitempty"`
	// Status of the request
	Status *string `json:"status,omitempty"`
	// Response data containing individual request results
	Data map[string]map[string]interface{} `json:"data,omitempty"`
}

// NewAdvanced200Response instantiates a new Advanced200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvanced200Response() *Advanced200Response {
	this := Advanced200Response{}
	return &this
}

// NewAdvanced200ResponseWithDefaults instantiates a new Advanced200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvanced200ResponseWithDefaults() *Advanced200Response {
	this := Advanced200Response{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Advanced200Response) GetCode() int64 {
	if o == nil || IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Advanced200Response) GetCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Advanced200Response) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *Advanced200Response) SetCode(v int64) {
	o.Code = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Advanced200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Advanced200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Advanced200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Advanced200Response) SetStatus(v string) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Advanced200Response) GetData() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Advanced200Response) GetDataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Advanced200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]map[string]interface{} and assigns it to the Data field.
func (o *Advanced200Response) SetData(v map[string]map[string]interface{}) {
	o.Data = v
}

func (o Advanced200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Advanced200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAdvanced200Response struct {
	value *Advanced200Response
	isSet bool
}

func (v NullableAdvanced200Response) Get() *Advanced200Response {
	return v.value
}

func (v *NullableAdvanced200Response) Set(val *Advanced200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvanced200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvanced200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvanced200Response(val *Advanced200Response) *NullableAdvanced200Response {
	return &NullableAdvanced200Response{value: val, isSet: true}
}

func (v NullableAdvanced200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvanced200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
