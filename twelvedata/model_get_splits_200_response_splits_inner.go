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

// checks if the GetSplits200ResponseSplitsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSplits200ResponseSplitsInner{}

// GetSplits200ResponseSplitsInner struct for GetSplits200ResponseSplitsInner
type GetSplits200ResponseSplitsInner struct {
	// Stands for the split date
	Date string `json:"date"`
	// Specification of the split event
	Description string `json:"description"`
	// The ratio by which the number of a company's outstanding shares of stock are increased following a stock split. For example, a `4-for-1 split` results in four times as many outstanding shares, with each share selling at one forth of its pre-split price
	Ratio float64 `json:"ratio"`
	// From factor of the split
	FromFactor float64 `json:"from_factor"`
	// To factor of the split
	ToFactor float64 `json:"to_factor"`
}

type _GetSplits200ResponseSplitsInner GetSplits200ResponseSplitsInner

// NewGetSplits200ResponseSplitsInner instantiates a new GetSplits200ResponseSplitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSplits200ResponseSplitsInner(date string, description string, ratio float64, fromFactor float64, toFactor float64) *GetSplits200ResponseSplitsInner {
	this := GetSplits200ResponseSplitsInner{}
	this.Date = date
	this.Description = description
	this.Ratio = ratio
	this.FromFactor = fromFactor
	this.ToFactor = toFactor
	return &this
}

// NewGetSplits200ResponseSplitsInnerWithDefaults instantiates a new GetSplits200ResponseSplitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSplits200ResponseSplitsInnerWithDefaults() *GetSplits200ResponseSplitsInner {
	this := GetSplits200ResponseSplitsInner{}
	return &this
}

// GetDate returns the Date field value
func (o *GetSplits200ResponseSplitsInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *GetSplits200ResponseSplitsInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *GetSplits200ResponseSplitsInner) SetDate(v string) {
	o.Date = v
}

// GetDescription returns the Description field value
func (o *GetSplits200ResponseSplitsInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GetSplits200ResponseSplitsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GetSplits200ResponseSplitsInner) SetDescription(v string) {
	o.Description = v
}

// GetRatio returns the Ratio field value
func (o *GetSplits200ResponseSplitsInner) GetRatio() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value
// and a boolean to check if the value has been set.
func (o *GetSplits200ResponseSplitsInner) GetRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ratio, true
}

// SetRatio sets field value
func (o *GetSplits200ResponseSplitsInner) SetRatio(v float64) {
	o.Ratio = v
}

// GetFromFactor returns the FromFactor field value
func (o *GetSplits200ResponseSplitsInner) GetFromFactor() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.FromFactor
}

// GetFromFactorOk returns a tuple with the FromFactor field value
// and a boolean to check if the value has been set.
func (o *GetSplits200ResponseSplitsInner) GetFromFactorOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromFactor, true
}

// SetFromFactor sets field value
func (o *GetSplits200ResponseSplitsInner) SetFromFactor(v float64) {
	o.FromFactor = v
}

// GetToFactor returns the ToFactor field value
func (o *GetSplits200ResponseSplitsInner) GetToFactor() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ToFactor
}

// GetToFactorOk returns a tuple with the ToFactor field value
// and a boolean to check if the value has been set.
func (o *GetSplits200ResponseSplitsInner) GetToFactorOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToFactor, true
}

// SetToFactor sets field value
func (o *GetSplits200ResponseSplitsInner) SetToFactor(v float64) {
	o.ToFactor = v
}

func (o GetSplits200ResponseSplitsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSplits200ResponseSplitsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["description"] = o.Description
	toSerialize["ratio"] = o.Ratio
	toSerialize["from_factor"] = o.FromFactor
	toSerialize["to_factor"] = o.ToFactor
	return toSerialize, nil
}

func (o *GetSplits200ResponseSplitsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"description",
		"ratio",
		"from_factor",
		"to_factor",
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

	varGetSplits200ResponseSplitsInner := _GetSplits200ResponseSplitsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSplits200ResponseSplitsInner)

	if err != nil {
		return err
	}

	*o = GetSplits200ResponseSplitsInner(varGetSplits200ResponseSplitsInner)

	return err
}

type NullableGetSplits200ResponseSplitsInner struct {
	value *GetSplits200ResponseSplitsInner
	isSet bool
}

func (v NullableGetSplits200ResponseSplitsInner) Get() *GetSplits200ResponseSplitsInner {
	return v.value
}

func (v *NullableGetSplits200ResponseSplitsInner) Set(val *GetSplits200ResponseSplitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSplits200ResponseSplitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSplits200ResponseSplitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSplits200ResponseSplitsInner(val *GetSplits200ResponseSplitsInner) *NullableGetSplits200ResponseSplitsInner {
	return &NullableGetSplits200ResponseSplitsInner{value: val, isSet: true}
}

func (v NullableGetSplits200ResponseSplitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSplits200ResponseSplitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
