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

// checks if the GetETFsWorldRisk200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorldRisk200Response{}

// GetETFsWorldRisk200Response struct for GetETFsWorldRisk200Response
type GetETFsWorldRisk200Response struct {
	Etf GetETFsWorldRisk200ResponseEtf `json:"etf"`
	// Status of the response
	Status string `json:"status"`
}

type _GetETFsWorldRisk200Response GetETFsWorldRisk200Response

// NewGetETFsWorldRisk200Response instantiates a new GetETFsWorldRisk200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorldRisk200Response(etf GetETFsWorldRisk200ResponseEtf, status string) *GetETFsWorldRisk200Response {
	this := GetETFsWorldRisk200Response{}
	this.Etf = etf
	this.Status = status
	return &this
}

// NewGetETFsWorldRisk200ResponseWithDefaults instantiates a new GetETFsWorldRisk200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorldRisk200ResponseWithDefaults() *GetETFsWorldRisk200Response {
	this := GetETFsWorldRisk200Response{}
	return &this
}

// GetEtf returns the Etf field value
func (o *GetETFsWorldRisk200Response) GetEtf() GetETFsWorldRisk200ResponseEtf {
	if o == nil {
		var ret GetETFsWorldRisk200ResponseEtf
		return ret
	}

	return o.Etf
}

// GetEtfOk returns a tuple with the Etf field value
// and a boolean to check if the value has been set.
func (o *GetETFsWorldRisk200Response) GetEtfOk() (*GetETFsWorldRisk200ResponseEtf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Etf, true
}

// SetEtf sets field value
func (o *GetETFsWorldRisk200Response) SetEtf(v GetETFsWorldRisk200ResponseEtf) {
	o.Etf = v
}

// GetStatus returns the Status field value
func (o *GetETFsWorldRisk200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetETFsWorldRisk200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetETFsWorldRisk200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetETFsWorldRisk200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorldRisk200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["etf"] = o.Etf
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetETFsWorldRisk200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetETFsWorldRisk200Response := _GetETFsWorldRisk200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetETFsWorldRisk200Response)

	if err != nil {
		return err
	}

	*o = GetETFsWorldRisk200Response(varGetETFsWorldRisk200Response)

	return err
}

type NullableGetETFsWorldRisk200Response struct {
	value *GetETFsWorldRisk200Response
	isSet bool
}

func (v NullableGetETFsWorldRisk200Response) Get() *GetETFsWorldRisk200Response {
	return v.value
}

func (v *NullableGetETFsWorldRisk200Response) Set(val *GetETFsWorldRisk200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorldRisk200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorldRisk200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorldRisk200Response(val *GetETFsWorldRisk200Response) *NullableGetETFsWorldRisk200Response {
	return &NullableGetETFsWorldRisk200Response{value: val, isSet: true}
}

func (v NullableGetETFsWorldRisk200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorldRisk200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
