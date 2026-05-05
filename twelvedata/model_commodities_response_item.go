// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CommoditiesResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommoditiesResponseItem{}

// CommoditiesResponseItem CommoditiesResponseItem represents details of a commodity
type CommoditiesResponseItem struct {
	// Currency pair according to ISO 4217 standard codes with slash(/) delimiter
	Symbol string `json:"symbol"`
	// Full name of the instrument
	Name string `json:"name"`
	// Category of commodity
	Category string `json:"category"`
	// Short description of the commodity
	Description string `json:"description"`
}

type _CommoditiesResponseItem CommoditiesResponseItem

// NewCommoditiesResponseItem instantiates a new CommoditiesResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommoditiesResponseItem(symbol string, name string, category string, description string) *CommoditiesResponseItem {
	this := CommoditiesResponseItem{}
	this.Symbol = symbol
	this.Name = name
	this.Category = category
	this.Description = description
	return &this
}

// NewCommoditiesResponseItemWithDefaults instantiates a new CommoditiesResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommoditiesResponseItemWithDefaults() *CommoditiesResponseItem {
	this := CommoditiesResponseItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *CommoditiesResponseItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *CommoditiesResponseItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *CommoditiesResponseItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *CommoditiesResponseItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CommoditiesResponseItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CommoditiesResponseItem) SetName(v string) {
	o.Name = v
}

// GetCategory returns the Category field value
func (o *CommoditiesResponseItem) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *CommoditiesResponseItem) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *CommoditiesResponseItem) SetCategory(v string) {
	o.Category = v
}

// GetDescription returns the Description field value
func (o *CommoditiesResponseItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CommoditiesResponseItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CommoditiesResponseItem) SetDescription(v string) {
	o.Description = v
}

func (o CommoditiesResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommoditiesResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["category"] = o.Category
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *CommoditiesResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"category",
		"description",
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

	varCommoditiesResponseItem := _CommoditiesResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCommoditiesResponseItem)

	if err != nil {
		return err
	}

	*o = CommoditiesResponseItem(varCommoditiesResponseItem)

	return err
}

type NullableCommoditiesResponseItem struct {
	value *CommoditiesResponseItem
	isSet bool
}

func (v NullableCommoditiesResponseItem) Get() *CommoditiesResponseItem {
	return v.value
}

func (v *NullableCommoditiesResponseItem) Set(val *CommoditiesResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCommoditiesResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCommoditiesResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommoditiesResponseItem(val *CommoditiesResponseItem) *NullableCommoditiesResponseItem {
	return &NullableCommoditiesResponseItem{value: val, isSet: true}
}

func (v NullableCommoditiesResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommoditiesResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
