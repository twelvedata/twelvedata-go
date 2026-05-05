// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetETFsWorldSummary200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorldSummary200Response{}

// GetETFsWorldSummary200Response struct for GetETFsWorldSummary200Response
type GetETFsWorldSummary200Response struct {
	Etf GetETFsWorldSummary200ResponseEtf `json:"etf"`
	// Status of the response
	Status string `json:"status"`
}

type _GetETFsWorldSummary200Response GetETFsWorldSummary200Response

// NewGetETFsWorldSummary200Response instantiates a new GetETFsWorldSummary200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorldSummary200Response(etf GetETFsWorldSummary200ResponseEtf, status string) *GetETFsWorldSummary200Response {
	this := GetETFsWorldSummary200Response{}
	this.Etf = etf
	this.Status = status
	return &this
}

// NewGetETFsWorldSummary200ResponseWithDefaults instantiates a new GetETFsWorldSummary200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorldSummary200ResponseWithDefaults() *GetETFsWorldSummary200Response {
	this := GetETFsWorldSummary200Response{}
	return &this
}

// GetEtf returns the Etf field value
func (o *GetETFsWorldSummary200Response) GetEtf() GetETFsWorldSummary200ResponseEtf {
	if o == nil {
		var ret GetETFsWorldSummary200ResponseEtf
		return ret
	}

	return o.Etf
}

// GetEtfOk returns a tuple with the Etf field value
// and a boolean to check if the value has been set.
func (o *GetETFsWorldSummary200Response) GetEtfOk() (*GetETFsWorldSummary200ResponseEtf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Etf, true
}

// SetEtf sets field value
func (o *GetETFsWorldSummary200Response) SetEtf(v GetETFsWorldSummary200ResponseEtf) {
	o.Etf = v
}

// GetStatus returns the Status field value
func (o *GetETFsWorldSummary200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetETFsWorldSummary200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetETFsWorldSummary200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetETFsWorldSummary200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorldSummary200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["etf"] = o.Etf
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetETFsWorldSummary200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetETFsWorldSummary200Response := _GetETFsWorldSummary200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetETFsWorldSummary200Response)

	if err != nil {
		return err
	}

	*o = GetETFsWorldSummary200Response(varGetETFsWorldSummary200Response)

	return err
}

type NullableGetETFsWorldSummary200Response struct {
	value *GetETFsWorldSummary200Response
	isSet bool
}

func (v NullableGetETFsWorldSummary200Response) Get() *GetETFsWorldSummary200Response {
	return v.value
}

func (v *NullableGetETFsWorldSummary200Response) Set(val *GetETFsWorldSummary200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorldSummary200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorldSummary200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorldSummary200Response(val *GetETFsWorldSummary200Response) *NullableGetETFsWorldSummary200Response {
	return &NullableGetETFsWorldSummary200Response{value: val, isSet: true}
}

func (v NullableGetETFsWorldSummary200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorldSummary200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
