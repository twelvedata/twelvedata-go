/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the AssetsInfoCurrentAssetsInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoCurrentAssetsInventory{}

// AssetsInfoCurrentAssetsInventory Inventory information
type AssetsInfoCurrentAssetsInventory struct {
	// Total inventory
	TotalInventory *float64 `json:"total_inventory,omitempty"`
	// Inventories adjustments allowances
	InventoriesAdjustmentsAllowances *float64 `json:"inventories_adjustments_allowances,omitempty"`
	// Other inventories
	OtherInventories *float64 `json:"other_inventories,omitempty"`
	// Finished goods
	FinishedGoods *float64 `json:"finished_goods,omitempty"`
	// Work in process
	WorkInProcess *float64 `json:"work_in_process,omitempty"`
	// Raw materials
	RawMaterials *float64 `json:"raw_materials,omitempty"`
}

// NewAssetsInfoCurrentAssetsInventory instantiates a new AssetsInfoCurrentAssetsInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoCurrentAssetsInventory() *AssetsInfoCurrentAssetsInventory {
	this := AssetsInfoCurrentAssetsInventory{}
	return &this
}

// NewAssetsInfoCurrentAssetsInventoryWithDefaults instantiates a new AssetsInfoCurrentAssetsInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoCurrentAssetsInventoryWithDefaults() *AssetsInfoCurrentAssetsInventory {
	this := AssetsInfoCurrentAssetsInventory{}
	return &this
}

// GetTotalInventory returns the TotalInventory field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsInventory) GetTotalInventory() float64 {
	if o == nil || IsNil(o.TotalInventory) {
		var ret float64
		return ret
	}
	return *o.TotalInventory
}

// GetTotalInventoryOk returns a tuple with the TotalInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsInventory) GetTotalInventoryOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalInventory) {
		return nil, false
	}
	return o.TotalInventory, true
}

// HasTotalInventory returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsInventory) HasTotalInventory() bool {
	if o != nil && !IsNil(o.TotalInventory) {
		return true
	}

	return false
}

// SetTotalInventory gets a reference to the given float64 and assigns it to the TotalInventory field.
func (o *AssetsInfoCurrentAssetsInventory) SetTotalInventory(v float64) {
	o.TotalInventory = &v
}

// GetInventoriesAdjustmentsAllowances returns the InventoriesAdjustmentsAllowances field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsInventory) GetInventoriesAdjustmentsAllowances() float64 {
	if o == nil || IsNil(o.InventoriesAdjustmentsAllowances) {
		var ret float64
		return ret
	}
	return *o.InventoriesAdjustmentsAllowances
}

// GetInventoriesAdjustmentsAllowancesOk returns a tuple with the InventoriesAdjustmentsAllowances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsInventory) GetInventoriesAdjustmentsAllowancesOk() (*float64, bool) {
	if o == nil || IsNil(o.InventoriesAdjustmentsAllowances) {
		return nil, false
	}
	return o.InventoriesAdjustmentsAllowances, true
}

// HasInventoriesAdjustmentsAllowances returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsInventory) HasInventoriesAdjustmentsAllowances() bool {
	if o != nil && !IsNil(o.InventoriesAdjustmentsAllowances) {
		return true
	}

	return false
}

// SetInventoriesAdjustmentsAllowances gets a reference to the given float64 and assigns it to the InventoriesAdjustmentsAllowances field.
func (o *AssetsInfoCurrentAssetsInventory) SetInventoriesAdjustmentsAllowances(v float64) {
	o.InventoriesAdjustmentsAllowances = &v
}

// GetOtherInventories returns the OtherInventories field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsInventory) GetOtherInventories() float64 {
	if o == nil || IsNil(o.OtherInventories) {
		var ret float64
		return ret
	}
	return *o.OtherInventories
}

// GetOtherInventoriesOk returns a tuple with the OtherInventories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsInventory) GetOtherInventoriesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherInventories) {
		return nil, false
	}
	return o.OtherInventories, true
}

// HasOtherInventories returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsInventory) HasOtherInventories() bool {
	if o != nil && !IsNil(o.OtherInventories) {
		return true
	}

	return false
}

// SetOtherInventories gets a reference to the given float64 and assigns it to the OtherInventories field.
func (o *AssetsInfoCurrentAssetsInventory) SetOtherInventories(v float64) {
	o.OtherInventories = &v
}

// GetFinishedGoods returns the FinishedGoods field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsInventory) GetFinishedGoods() float64 {
	if o == nil || IsNil(o.FinishedGoods) {
		var ret float64
		return ret
	}
	return *o.FinishedGoods
}

// GetFinishedGoodsOk returns a tuple with the FinishedGoods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsInventory) GetFinishedGoodsOk() (*float64, bool) {
	if o == nil || IsNil(o.FinishedGoods) {
		return nil, false
	}
	return o.FinishedGoods, true
}

// HasFinishedGoods returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsInventory) HasFinishedGoods() bool {
	if o != nil && !IsNil(o.FinishedGoods) {
		return true
	}

	return false
}

// SetFinishedGoods gets a reference to the given float64 and assigns it to the FinishedGoods field.
func (o *AssetsInfoCurrentAssetsInventory) SetFinishedGoods(v float64) {
	o.FinishedGoods = &v
}

// GetWorkInProcess returns the WorkInProcess field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsInventory) GetWorkInProcess() float64 {
	if o == nil || IsNil(o.WorkInProcess) {
		var ret float64
		return ret
	}
	return *o.WorkInProcess
}

// GetWorkInProcessOk returns a tuple with the WorkInProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsInventory) GetWorkInProcessOk() (*float64, bool) {
	if o == nil || IsNil(o.WorkInProcess) {
		return nil, false
	}
	return o.WorkInProcess, true
}

// HasWorkInProcess returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsInventory) HasWorkInProcess() bool {
	if o != nil && !IsNil(o.WorkInProcess) {
		return true
	}

	return false
}

// SetWorkInProcess gets a reference to the given float64 and assigns it to the WorkInProcess field.
func (o *AssetsInfoCurrentAssetsInventory) SetWorkInProcess(v float64) {
	o.WorkInProcess = &v
}

// GetRawMaterials returns the RawMaterials field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsInventory) GetRawMaterials() float64 {
	if o == nil || IsNil(o.RawMaterials) {
		var ret float64
		return ret
	}
	return *o.RawMaterials
}

// GetRawMaterialsOk returns a tuple with the RawMaterials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsInventory) GetRawMaterialsOk() (*float64, bool) {
	if o == nil || IsNil(o.RawMaterials) {
		return nil, false
	}
	return o.RawMaterials, true
}

// HasRawMaterials returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsInventory) HasRawMaterials() bool {
	if o != nil && !IsNil(o.RawMaterials) {
		return true
	}

	return false
}

// SetRawMaterials gets a reference to the given float64 and assigns it to the RawMaterials field.
func (o *AssetsInfoCurrentAssetsInventory) SetRawMaterials(v float64) {
	o.RawMaterials = &v
}

func (o AssetsInfoCurrentAssetsInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoCurrentAssetsInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalInventory) {
		toSerialize["total_inventory"] = o.TotalInventory
	}
	if !IsNil(o.InventoriesAdjustmentsAllowances) {
		toSerialize["inventories_adjustments_allowances"] = o.InventoriesAdjustmentsAllowances
	}
	if !IsNil(o.OtherInventories) {
		toSerialize["other_inventories"] = o.OtherInventories
	}
	if !IsNil(o.FinishedGoods) {
		toSerialize["finished_goods"] = o.FinishedGoods
	}
	if !IsNil(o.WorkInProcess) {
		toSerialize["work_in_process"] = o.WorkInProcess
	}
	if !IsNil(o.RawMaterials) {
		toSerialize["raw_materials"] = o.RawMaterials
	}
	return toSerialize, nil
}

type NullableAssetsInfoCurrentAssetsInventory struct {
	value *AssetsInfoCurrentAssetsInventory
	isSet bool
}

func (v NullableAssetsInfoCurrentAssetsInventory) Get() *AssetsInfoCurrentAssetsInventory {
	return v.value
}

func (v *NullableAssetsInfoCurrentAssetsInventory) Set(val *AssetsInfoCurrentAssetsInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoCurrentAssetsInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoCurrentAssetsInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoCurrentAssetsInventory(val *AssetsInfoCurrentAssetsInventory) *NullableAssetsInfoCurrentAssetsInventory {
	return &NullableAssetsInfoCurrentAssetsInventory{value: val, isSet: true}
}

func (v NullableAssetsInfoCurrentAssetsInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoCurrentAssetsInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
