// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetETFsWorld200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200Response{}

// GetETFsWorld200Response struct for GetETFsWorld200Response
type GetETFsWorld200Response struct {
	Etf GetETFsWorld200ResponseEtf `json:"etf"`
	// Status of the response
	Status string `json:"status"`
}

type _GetETFsWorld200Response GetETFsWorld200Response

// NewGetETFsWorld200Response instantiates a new GetETFsWorld200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200Response(etf GetETFsWorld200ResponseEtf, status string) *GetETFsWorld200Response {
	this := GetETFsWorld200Response{}
	this.Etf = etf
	this.Status = status
	return &this
}

// NewGetETFsWorld200ResponseWithDefaults instantiates a new GetETFsWorld200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseWithDefaults() *GetETFsWorld200Response {
	this := GetETFsWorld200Response{}
	return &this
}

// GetEtf returns the Etf field value
func (o *GetETFsWorld200Response) GetEtf() GetETFsWorld200ResponseEtf {
	if o == nil {
		var ret GetETFsWorld200ResponseEtf
		return ret
	}

	return o.Etf
}

// GetEtfOk returns a tuple with the Etf field value
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200Response) GetEtfOk() (*GetETFsWorld200ResponseEtf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Etf, true
}

// SetEtf sets field value
func (o *GetETFsWorld200Response) SetEtf(v GetETFsWorld200ResponseEtf) {
	o.Etf = v
}

// GetStatus returns the Status field value
func (o *GetETFsWorld200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetETFsWorld200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetETFsWorld200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["etf"] = o.Etf
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetETFsWorld200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"etf",
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

	varGetETFsWorld200Response := _GetETFsWorld200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetETFsWorld200Response)

	if err != nil {
		return err
	}

	*o = GetETFsWorld200Response(varGetETFsWorld200Response)

	return err
}

type NullableGetETFsWorld200Response struct {
	value *GetETFsWorld200Response
	isSet bool
}

func (v NullableGetETFsWorld200Response) Get() *GetETFsWorld200Response {
	return v.value
}

func (v *NullableGetETFsWorld200Response) Set(val *GetETFsWorld200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200Response(val *GetETFsWorld200Response) *NullableGetETFsWorld200Response {
	return &NullableGetETFsWorld200Response{value: val, isSet: true}
}

func (v NullableGetETFsWorld200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
