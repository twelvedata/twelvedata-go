// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BalanceSheetConsolidatedItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BalanceSheetConsolidatedItem{}

// BalanceSheetConsolidatedItem Balance sheet for a specific fiscal date
type BalanceSheetConsolidatedItem struct {
	// Date of fiscal period ending
	FiscalDate string      `json:"fiscal_date"`
	Assets     *AssetsInfo `json:"assets,omitempty"`
}

type _BalanceSheetConsolidatedItem BalanceSheetConsolidatedItem

// NewBalanceSheetConsolidatedItem instantiates a new BalanceSheetConsolidatedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceSheetConsolidatedItem(fiscalDate string) *BalanceSheetConsolidatedItem {
	this := BalanceSheetConsolidatedItem{}
	this.FiscalDate = fiscalDate
	return &this
}

// NewBalanceSheetConsolidatedItemWithDefaults instantiates a new BalanceSheetConsolidatedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceSheetConsolidatedItemWithDefaults() *BalanceSheetConsolidatedItem {
	this := BalanceSheetConsolidatedItem{}
	return &this
}

// GetFiscalDate returns the FiscalDate field value
func (o *BalanceSheetConsolidatedItem) GetFiscalDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FiscalDate
}

// GetFiscalDateOk returns a tuple with the FiscalDate field value
// and a boolean to check if the value has been set.
func (o *BalanceSheetConsolidatedItem) GetFiscalDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiscalDate, true
}

// SetFiscalDate sets field value
func (o *BalanceSheetConsolidatedItem) SetFiscalDate(v string) {
	o.FiscalDate = v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *BalanceSheetConsolidatedItem) GetAssets() AssetsInfo {
	if o == nil || IsNil(o.Assets) {
		var ret AssetsInfo
		return ret
	}
	return *o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceSheetConsolidatedItem) GetAssetsOk() (*AssetsInfo, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *BalanceSheetConsolidatedItem) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given AssetsInfo and assigns it to the Assets field.
func (o *BalanceSheetConsolidatedItem) SetAssets(v AssetsInfo) {
	o.Assets = &v
}

func (o BalanceSheetConsolidatedItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceSheetConsolidatedItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fiscal_date"] = o.FiscalDate
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	return toSerialize, nil
}

func (o *BalanceSheetConsolidatedItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fiscal_date",
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

	varBalanceSheetConsolidatedItem := _BalanceSheetConsolidatedItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varBalanceSheetConsolidatedItem)

	if err != nil {
		return err
	}

	*o = BalanceSheetConsolidatedItem(varBalanceSheetConsolidatedItem)

	return err
}

type NullableBalanceSheetConsolidatedItem struct {
	value *BalanceSheetConsolidatedItem
	isSet bool
}

func (v NullableBalanceSheetConsolidatedItem) Get() *BalanceSheetConsolidatedItem {
	return v.value
}

func (v *NullableBalanceSheetConsolidatedItem) Set(val *BalanceSheetConsolidatedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceSheetConsolidatedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceSheetConsolidatedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceSheetConsolidatedItem(val *BalanceSheetConsolidatedItem) *NullableBalanceSheetConsolidatedItem {
	return &NullableBalanceSheetConsolidatedItem{value: val, isSet: true}
}

func (v NullableBalanceSheetConsolidatedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceSheetConsolidatedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
