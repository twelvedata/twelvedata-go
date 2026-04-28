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

// checks if the GetCountries200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCountries200Response{}

// GetCountries200Response struct for GetCountries200Response
type GetCountries200Response struct {
	// List of countries with their ISO codes, names, capital, and currency
	Data []CountryResponseItem `json:"data"`
}

type _GetCountries200Response GetCountries200Response

// NewGetCountries200Response instantiates a new GetCountries200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCountries200Response(data []CountryResponseItem) *GetCountries200Response {
	this := GetCountries200Response{}
	this.Data = data
	return &this
}

// NewGetCountries200ResponseWithDefaults instantiates a new GetCountries200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCountries200ResponseWithDefaults() *GetCountries200Response {
	this := GetCountries200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetCountries200Response) GetData() []CountryResponseItem {
	if o == nil {
		var ret []CountryResponseItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCountries200Response) GetDataOk() ([]CountryResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetCountries200Response) SetData(v []CountryResponseItem) {
	o.Data = v
}

func (o GetCountries200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCountries200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetCountries200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGetCountries200Response := _GetCountries200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCountries200Response)

	if err != nil {
		return err
	}

	*o = GetCountries200Response(varGetCountries200Response)

	return err
}

type NullableGetCountries200Response struct {
	value *GetCountries200Response
	isSet bool
}

func (v NullableGetCountries200Response) Get() *GetCountries200Response {
	return v.value
}

func (v *NullableGetCountries200Response) Set(val *GetCountries200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCountries200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCountries200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCountries200Response(val *GetCountries200Response) *NullableGetCountries200Response {
	return &NullableGetCountries200Response{value: val, isSet: true}
}

func (v NullableGetCountries200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCountries200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
