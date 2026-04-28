/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemUnusualItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemUnusualItems{}

// IncomeStatementItemUnusualItems Unusual items information
type IncomeStatementItemUnusualItems struct {
	// Total unusual items
	TotalUnusualItems *float64 `json:"total_unusual_items,omitempty"`
	// Total unusual items excluding goodwill
	TotalUnusualItemsExcludingGoodwill *float64 `json:"total_unusual_items_excluding_goodwill,omitempty"`
}

// NewIncomeStatementItemUnusualItems instantiates a new IncomeStatementItemUnusualItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemUnusualItems() *IncomeStatementItemUnusualItems {
	this := IncomeStatementItemUnusualItems{}
	return &this
}

// NewIncomeStatementItemUnusualItemsWithDefaults instantiates a new IncomeStatementItemUnusualItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemUnusualItemsWithDefaults() *IncomeStatementItemUnusualItems {
	this := IncomeStatementItemUnusualItems{}
	return &this
}

// GetTotalUnusualItems returns the TotalUnusualItems field value if set, zero value otherwise.
func (o *IncomeStatementItemUnusualItems) GetTotalUnusualItems() float64 {
	if o == nil || IsNil(o.TotalUnusualItems) {
		var ret float64
		return ret
	}
	return *o.TotalUnusualItems
}

// GetTotalUnusualItemsOk returns a tuple with the TotalUnusualItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemUnusualItems) GetTotalUnusualItemsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalUnusualItems) {
		return nil, false
	}
	return o.TotalUnusualItems, true
}

// HasTotalUnusualItems returns a boolean if a field has been set.
func (o *IncomeStatementItemUnusualItems) HasTotalUnusualItems() bool {
	if o != nil && !IsNil(o.TotalUnusualItems) {
		return true
	}

	return false
}

// SetTotalUnusualItems gets a reference to the given float64 and assigns it to the TotalUnusualItems field.
func (o *IncomeStatementItemUnusualItems) SetTotalUnusualItems(v float64) {
	o.TotalUnusualItems = &v
}

// GetTotalUnusualItemsExcludingGoodwill returns the TotalUnusualItemsExcludingGoodwill field value if set, zero value otherwise.
func (o *IncomeStatementItemUnusualItems) GetTotalUnusualItemsExcludingGoodwill() float64 {
	if o == nil || IsNil(o.TotalUnusualItemsExcludingGoodwill) {
		var ret float64
		return ret
	}
	return *o.TotalUnusualItemsExcludingGoodwill
}

// GetTotalUnusualItemsExcludingGoodwillOk returns a tuple with the TotalUnusualItemsExcludingGoodwill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemUnusualItems) GetTotalUnusualItemsExcludingGoodwillOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalUnusualItemsExcludingGoodwill) {
		return nil, false
	}
	return o.TotalUnusualItemsExcludingGoodwill, true
}

// HasTotalUnusualItemsExcludingGoodwill returns a boolean if a field has been set.
func (o *IncomeStatementItemUnusualItems) HasTotalUnusualItemsExcludingGoodwill() bool {
	if o != nil && !IsNil(o.TotalUnusualItemsExcludingGoodwill) {
		return true
	}

	return false
}

// SetTotalUnusualItemsExcludingGoodwill gets a reference to the given float64 and assigns it to the TotalUnusualItemsExcludingGoodwill field.
func (o *IncomeStatementItemUnusualItems) SetTotalUnusualItemsExcludingGoodwill(v float64) {
	o.TotalUnusualItemsExcludingGoodwill = &v
}

func (o IncomeStatementItemUnusualItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemUnusualItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalUnusualItems) {
		toSerialize["total_unusual_items"] = o.TotalUnusualItems
	}
	if !IsNil(o.TotalUnusualItemsExcludingGoodwill) {
		toSerialize["total_unusual_items_excluding_goodwill"] = o.TotalUnusualItemsExcludingGoodwill
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemUnusualItems struct {
	value *IncomeStatementItemUnusualItems
	isSet bool
}

func (v NullableIncomeStatementItemUnusualItems) Get() *IncomeStatementItemUnusualItems {
	return v.value
}

func (v *NullableIncomeStatementItemUnusualItems) Set(val *IncomeStatementItemUnusualItems) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemUnusualItems) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemUnusualItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemUnusualItems(val *IncomeStatementItemUnusualItems) *NullableIncomeStatementItemUnusualItems {
	return &NullableIncomeStatementItemUnusualItems{value: val, isSet: true}
}

func (v NullableIncomeStatementItemUnusualItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemUnusualItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
