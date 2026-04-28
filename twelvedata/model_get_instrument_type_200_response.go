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

// checks if the GetInstrumentType200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInstrumentType200Response{}

// GetInstrumentType200Response struct for GetInstrumentType200Response
type GetInstrumentType200Response struct {
	// List of instrument types available at Twelve Data API.
	Result []string `json:"result"`
	// Status of the response
	Status string `json:"status"`
}

type _GetInstrumentType200Response GetInstrumentType200Response

// NewGetInstrumentType200Response instantiates a new GetInstrumentType200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInstrumentType200Response(result []string, status string) *GetInstrumentType200Response {
	this := GetInstrumentType200Response{}
	this.Result = result
	this.Status = status
	return &this
}

// NewGetInstrumentType200ResponseWithDefaults instantiates a new GetInstrumentType200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInstrumentType200ResponseWithDefaults() *GetInstrumentType200Response {
	this := GetInstrumentType200Response{}
	return &this
}

// GetResult returns the Result field value
func (o *GetInstrumentType200Response) GetResult() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetInstrumentType200Response) GetResultOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *GetInstrumentType200Response) SetResult(v []string) {
	o.Result = v
}

// GetStatus returns the Status field value
func (o *GetInstrumentType200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetInstrumentType200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetInstrumentType200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetInstrumentType200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInstrumentType200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetInstrumentType200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetInstrumentType200Response := _GetInstrumentType200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInstrumentType200Response)

	if err != nil {
		return err
	}

	*o = GetInstrumentType200Response(varGetInstrumentType200Response)

	return err
}

type NullableGetInstrumentType200Response struct {
	value *GetInstrumentType200Response
	isSet bool
}

func (v NullableGetInstrumentType200Response) Get() *GetInstrumentType200Response {
	return v.value
}

func (v *NullableGetInstrumentType200Response) Set(val *GetInstrumentType200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInstrumentType200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInstrumentType200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInstrumentType200Response(val *GetInstrumentType200Response) *NullableGetInstrumentType200Response {
	return &NullableGetInstrumentType200Response{value: val, isSet: true}
}

func (v NullableGetInstrumentType200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInstrumentType200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
