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

// checks if the GetTaxInfo200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTaxInfo200ResponseData{}

// GetTaxInfo200ResponseData Tax information data containing the tax indicator for the requested instrument
type GetTaxInfo200ResponseData struct {
	// The instrument tax indicator, can be `null`, `us_1446f`, `spanish_ftt`, `uk_stamp_duty`, `hk_stamp_duty`, `french_ftt` or `italian_ftt`
	TaxIndicator string `json:"tax_indicator"`
}

type _GetTaxInfo200ResponseData GetTaxInfo200ResponseData

// NewGetTaxInfo200ResponseData instantiates a new GetTaxInfo200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTaxInfo200ResponseData(taxIndicator string) *GetTaxInfo200ResponseData {
	this := GetTaxInfo200ResponseData{}
	this.TaxIndicator = taxIndicator
	return &this
}

// NewGetTaxInfo200ResponseDataWithDefaults instantiates a new GetTaxInfo200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTaxInfo200ResponseDataWithDefaults() *GetTaxInfo200ResponseData {
	this := GetTaxInfo200ResponseData{}
	return &this
}

// GetTaxIndicator returns the TaxIndicator field value
func (o *GetTaxInfo200ResponseData) GetTaxIndicator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxIndicator
}

// GetTaxIndicatorOk returns a tuple with the TaxIndicator field value
// and a boolean to check if the value has been set.
func (o *GetTaxInfo200ResponseData) GetTaxIndicatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxIndicator, true
}

// SetTaxIndicator sets field value
func (o *GetTaxInfo200ResponseData) SetTaxIndicator(v string) {
	o.TaxIndicator = v
}

func (o GetTaxInfo200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTaxInfo200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tax_indicator"] = o.TaxIndicator
	return toSerialize, nil
}

func (o *GetTaxInfo200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tax_indicator",
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

	varGetTaxInfo200ResponseData := _GetTaxInfo200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTaxInfo200ResponseData)

	if err != nil {
		return err
	}

	*o = GetTaxInfo200ResponseData(varGetTaxInfo200ResponseData)

	return err
}

type NullableGetTaxInfo200ResponseData struct {
	value *GetTaxInfo200ResponseData
	isSet bool
}

func (v NullableGetTaxInfo200ResponseData) Get() *GetTaxInfo200ResponseData {
	return v.value
}

func (v *NullableGetTaxInfo200ResponseData) Set(val *GetTaxInfo200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTaxInfo200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTaxInfo200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTaxInfo200ResponseData(val *GetTaxInfo200ResponseData) *NullableGetTaxInfo200ResponseData {
	return &NullableGetTaxInfo200ResponseData{value: val, isSet: true}
}

func (v NullableGetTaxInfo200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTaxInfo200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
