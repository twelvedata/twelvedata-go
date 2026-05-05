// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetBalanceSheet200ResponseBalanceSheetInnerAssets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheet200ResponseBalanceSheetInnerAssets{}

// GetBalanceSheet200ResponseBalanceSheetInnerAssets Assets section of the balance sheet
type GetBalanceSheet200ResponseBalanceSheetInnerAssets struct {
	CurrentAssets    *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets    `json:"current_assets,omitempty"`
	NonCurrentAssets *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets `json:"non_current_assets,omitempty"`
	// The sum of total_current_assets + total_non_current_assets
	TotalAssets *float64 `json:"total_assets,omitempty"`
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerAssets instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheet200ResponseBalanceSheetInnerAssets() *GetBalanceSheet200ResponseBalanceSheetInnerAssets {
	this := GetBalanceSheet200ResponseBalanceSheetInnerAssets{}
	return &this
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInnerAssets {
	this := GetBalanceSheet200ResponseBalanceSheetInnerAssets{}
	return &this
}

// GetCurrentAssets returns the CurrentAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) GetCurrentAssets() GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets {
	if o == nil || IsNil(o.CurrentAssets) {
		var ret GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets
		return ret
	}
	return *o.CurrentAssets
}

// GetCurrentAssetsOk returns a tuple with the CurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) GetCurrentAssetsOk() (*GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets, bool) {
	if o == nil || IsNil(o.CurrentAssets) {
		return nil, false
	}
	return o.CurrentAssets, true
}

// HasCurrentAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) HasCurrentAssets() bool {
	if o != nil && !IsNil(o.CurrentAssets) {
		return true
	}

	return false
}

// SetCurrentAssets gets a reference to the given GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets and assigns it to the CurrentAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) SetCurrentAssets(v GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) {
	o.CurrentAssets = &v
}

// GetNonCurrentAssets returns the NonCurrentAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) GetNonCurrentAssets() GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets {
	if o == nil || IsNil(o.NonCurrentAssets) {
		var ret GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets
		return ret
	}
	return *o.NonCurrentAssets
}

// GetNonCurrentAssetsOk returns a tuple with the NonCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) GetNonCurrentAssetsOk() (*GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets, bool) {
	if o == nil || IsNil(o.NonCurrentAssets) {
		return nil, false
	}
	return o.NonCurrentAssets, true
}

// HasNonCurrentAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) HasNonCurrentAssets() bool {
	if o != nil && !IsNil(o.NonCurrentAssets) {
		return true
	}

	return false
}

// SetNonCurrentAssets gets a reference to the given GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets and assigns it to the NonCurrentAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) SetNonCurrentAssets(v GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) {
	o.NonCurrentAssets = &v
}

// GetTotalAssets returns the TotalAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) GetTotalAssets() float64 {
	if o == nil || IsNil(o.TotalAssets) {
		var ret float64
		return ret
	}
	return *o.TotalAssets
}

// GetTotalAssetsOk returns a tuple with the TotalAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) GetTotalAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalAssets) {
		return nil, false
	}
	return o.TotalAssets, true
}

// HasTotalAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) HasTotalAssets() bool {
	if o != nil && !IsNil(o.TotalAssets) {
		return true
	}

	return false
}

// SetTotalAssets gets a reference to the given float64 and assigns it to the TotalAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) SetTotalAssets(v float64) {
	o.TotalAssets = &v
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerAssets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerAssets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentAssets) {
		toSerialize["current_assets"] = o.CurrentAssets
	}
	if !IsNil(o.NonCurrentAssets) {
		toSerialize["non_current_assets"] = o.NonCurrentAssets
	}
	if !IsNil(o.TotalAssets) {
		toSerialize["total_assets"] = o.TotalAssets
	}
	return toSerialize, nil
}

type NullableGetBalanceSheet200ResponseBalanceSheetInnerAssets struct {
	value *GetBalanceSheet200ResponseBalanceSheetInnerAssets
	isSet bool
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssets) Get() *GetBalanceSheet200ResponseBalanceSheetInnerAssets {
	return v.value
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssets) Set(val *GetBalanceSheet200ResponseBalanceSheetInnerAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheet200ResponseBalanceSheetInnerAssets(val *GetBalanceSheet200ResponseBalanceSheetInnerAssets) *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssets {
	return &NullableGetBalanceSheet200ResponseBalanceSheetInnerAssets{value: val, isSet: true}
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
