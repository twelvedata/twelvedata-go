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

// checks if the GetETFsWorldComposition200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorldComposition200Response{}

// GetETFsWorldComposition200Response struct for GetETFsWorldComposition200Response
type GetETFsWorldComposition200Response struct {
	Etf GetETFsWorldComposition200ResponseEtf `json:"etf"`
	// Status of the response
	Status string `json:"status"`
}

type _GetETFsWorldComposition200Response GetETFsWorldComposition200Response

// NewGetETFsWorldComposition200Response instantiates a new GetETFsWorldComposition200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorldComposition200Response(etf GetETFsWorldComposition200ResponseEtf, status string) *GetETFsWorldComposition200Response {
	this := GetETFsWorldComposition200Response{}
	this.Etf = etf
	this.Status = status
	return &this
}

// NewGetETFsWorldComposition200ResponseWithDefaults instantiates a new GetETFsWorldComposition200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorldComposition200ResponseWithDefaults() *GetETFsWorldComposition200Response {
	this := GetETFsWorldComposition200Response{}
	return &this
}

// GetEtf returns the Etf field value
func (o *GetETFsWorldComposition200Response) GetEtf() GetETFsWorldComposition200ResponseEtf {
	if o == nil {
		var ret GetETFsWorldComposition200ResponseEtf
		return ret
	}

	return o.Etf
}

// GetEtfOk returns a tuple with the Etf field value
// and a boolean to check if the value has been set.
func (o *GetETFsWorldComposition200Response) GetEtfOk() (*GetETFsWorldComposition200ResponseEtf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Etf, true
}

// SetEtf sets field value
func (o *GetETFsWorldComposition200Response) SetEtf(v GetETFsWorldComposition200ResponseEtf) {
	o.Etf = v
}

// GetStatus returns the Status field value
func (o *GetETFsWorldComposition200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetETFsWorldComposition200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetETFsWorldComposition200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetETFsWorldComposition200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorldComposition200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["etf"] = o.Etf
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetETFsWorldComposition200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetETFsWorldComposition200Response := _GetETFsWorldComposition200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetETFsWorldComposition200Response)

	if err != nil {
		return err
	}

	*o = GetETFsWorldComposition200Response(varGetETFsWorldComposition200Response)

	return err
}

type NullableGetETFsWorldComposition200Response struct {
	value *GetETFsWorldComposition200Response
	isSet bool
}

func (v NullableGetETFsWorldComposition200Response) Get() *GetETFsWorldComposition200Response {
	return v.value
}

func (v *NullableGetETFsWorldComposition200Response) Set(val *GetETFsWorldComposition200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorldComposition200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorldComposition200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorldComposition200Response(val *GetETFsWorldComposition200Response) *NullableGetETFsWorldComposition200Response {
	return &NullableGetETFsWorldComposition200Response{value: val, isSet: true}
}

func (v NullableGetETFsWorldComposition200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorldComposition200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
