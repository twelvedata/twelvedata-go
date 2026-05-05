// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetETFsWorldPerformance200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorldPerformance200Response{}

// GetETFsWorldPerformance200Response struct for GetETFsWorldPerformance200Response
type GetETFsWorldPerformance200Response struct {
	Etf GetETFsWorldPerformance200ResponseEtf `json:"etf"`
	// Status of the response
	Status string `json:"status"`
}

type _GetETFsWorldPerformance200Response GetETFsWorldPerformance200Response

// NewGetETFsWorldPerformance200Response instantiates a new GetETFsWorldPerformance200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorldPerformance200Response(etf GetETFsWorldPerformance200ResponseEtf, status string) *GetETFsWorldPerformance200Response {
	this := GetETFsWorldPerformance200Response{}
	this.Etf = etf
	this.Status = status
	return &this
}

// NewGetETFsWorldPerformance200ResponseWithDefaults instantiates a new GetETFsWorldPerformance200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorldPerformance200ResponseWithDefaults() *GetETFsWorldPerformance200Response {
	this := GetETFsWorldPerformance200Response{}
	return &this
}

// GetEtf returns the Etf field value
func (o *GetETFsWorldPerformance200Response) GetEtf() GetETFsWorldPerformance200ResponseEtf {
	if o == nil {
		var ret GetETFsWorldPerformance200ResponseEtf
		return ret
	}

	return o.Etf
}

// GetEtfOk returns a tuple with the Etf field value
// and a boolean to check if the value has been set.
func (o *GetETFsWorldPerformance200Response) GetEtfOk() (*GetETFsWorldPerformance200ResponseEtf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Etf, true
}

// SetEtf sets field value
func (o *GetETFsWorldPerformance200Response) SetEtf(v GetETFsWorldPerformance200ResponseEtf) {
	o.Etf = v
}

// GetStatus returns the Status field value
func (o *GetETFsWorldPerformance200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetETFsWorldPerformance200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetETFsWorldPerformance200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetETFsWorldPerformance200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorldPerformance200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["etf"] = o.Etf
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetETFsWorldPerformance200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetETFsWorldPerformance200Response := _GetETFsWorldPerformance200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetETFsWorldPerformance200Response)

	if err != nil {
		return err
	}

	*o = GetETFsWorldPerformance200Response(varGetETFsWorldPerformance200Response)

	return err
}

type NullableGetETFsWorldPerformance200Response struct {
	value *GetETFsWorldPerformance200Response
	isSet bool
}

func (v NullableGetETFsWorldPerformance200Response) Get() *GetETFsWorldPerformance200Response {
	return v.value
}

func (v *NullableGetETFsWorldPerformance200Response) Set(val *GetETFsWorldPerformance200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorldPerformance200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorldPerformance200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorldPerformance200Response(val *GetETFsWorldPerformance200Response) *NullableGetETFsWorldPerformance200Response {
	return &NullableGetETFsWorldPerformance200Response{value: val, isSet: true}
}

func (v NullableGetETFsWorldPerformance200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorldPerformance200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
