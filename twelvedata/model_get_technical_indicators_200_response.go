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

// checks if the GetTechnicalIndicators200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTechnicalIndicators200Response{}

// GetTechnicalIndicators200Response struct for GetTechnicalIndicators200Response
type GetTechnicalIndicators200Response struct {
	// Map of technical indicators available at Twelve Data API
	Data map[string]GetTechnicalIndicators200ResponseDataValue `json:"data"`
	// Response status
	Status string `json:"status"`
}

type _GetTechnicalIndicators200Response GetTechnicalIndicators200Response

// NewGetTechnicalIndicators200Response instantiates a new GetTechnicalIndicators200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTechnicalIndicators200Response(data map[string]GetTechnicalIndicators200ResponseDataValue, status string) *GetTechnicalIndicators200Response {
	this := GetTechnicalIndicators200Response{}
	this.Data = data
	this.Status = status
	return &this
}

// NewGetTechnicalIndicators200ResponseWithDefaults instantiates a new GetTechnicalIndicators200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTechnicalIndicators200ResponseWithDefaults() *GetTechnicalIndicators200Response {
	this := GetTechnicalIndicators200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetTechnicalIndicators200Response) GetData() map[string]GetTechnicalIndicators200ResponseDataValue {
	if o == nil {
		var ret map[string]GetTechnicalIndicators200ResponseDataValue
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetTechnicalIndicators200Response) GetDataOk() (map[string]GetTechnicalIndicators200ResponseDataValue, bool) {
	if o == nil {
		return map[string]GetTechnicalIndicators200ResponseDataValue{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetTechnicalIndicators200Response) SetData(v map[string]GetTechnicalIndicators200ResponseDataValue) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *GetTechnicalIndicators200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTechnicalIndicators200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTechnicalIndicators200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTechnicalIndicators200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTechnicalIndicators200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTechnicalIndicators200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGetTechnicalIndicators200Response := _GetTechnicalIndicators200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTechnicalIndicators200Response)

	if err != nil {
		return err
	}

	*o = GetTechnicalIndicators200Response(varGetTechnicalIndicators200Response)

	return err
}

type NullableGetTechnicalIndicators200Response struct {
	value *GetTechnicalIndicators200Response
	isSet bool
}

func (v NullableGetTechnicalIndicators200Response) Get() *GetTechnicalIndicators200Response {
	return v.value
}

func (v *NullableGetTechnicalIndicators200Response) Set(val *GetTechnicalIndicators200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTechnicalIndicators200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTechnicalIndicators200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTechnicalIndicators200Response(val *GetTechnicalIndicators200Response) *NullableGetTechnicalIndicators200Response {
	return &NullableGetTechnicalIndicators200Response{value: val, isSet: true}
}

func (v NullableGetTechnicalIndicators200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTechnicalIndicators200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
