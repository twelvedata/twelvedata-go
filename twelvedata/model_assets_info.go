/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the AssetsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfo{}

// AssetsInfo AssetsInfo represents assets information in the balance sheet
type AssetsInfo struct {
	// Total assets
	TotalAssets      *float64                    `json:"total_assets,omitempty"`
	CurrentAssets    *AssetsInfoCurrentAssets    `json:"current_assets,omitempty"`
	NonCurrentAssets *AssetsInfoNonCurrentAssets `json:"non_current_assets,omitempty"`
	Liabilities      *AssetsInfoLiabilities      `json:"liabilities,omitempty"`
}

// NewAssetsInfo instantiates a new AssetsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfo() *AssetsInfo {
	this := AssetsInfo{}
	return &this
}

// NewAssetsInfoWithDefaults instantiates a new AssetsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoWithDefaults() *AssetsInfo {
	this := AssetsInfo{}
	return &this
}

// GetTotalAssets returns the TotalAssets field value if set, zero value otherwise.
func (o *AssetsInfo) GetTotalAssets() float64 {
	if o == nil || IsNil(o.TotalAssets) {
		var ret float64
		return ret
	}
	return *o.TotalAssets
}

// GetTotalAssetsOk returns a tuple with the TotalAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfo) GetTotalAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalAssets) {
		return nil, false
	}
	return o.TotalAssets, true
}

// HasTotalAssets returns a boolean if a field has been set.
func (o *AssetsInfo) HasTotalAssets() bool {
	if o != nil && !IsNil(o.TotalAssets) {
		return true
	}

	return false
}

// SetTotalAssets gets a reference to the given float64 and assigns it to the TotalAssets field.
func (o *AssetsInfo) SetTotalAssets(v float64) {
	o.TotalAssets = &v
}

// GetCurrentAssets returns the CurrentAssets field value if set, zero value otherwise.
func (o *AssetsInfo) GetCurrentAssets() AssetsInfoCurrentAssets {
	if o == nil || IsNil(o.CurrentAssets) {
		var ret AssetsInfoCurrentAssets
		return ret
	}
	return *o.CurrentAssets
}

// GetCurrentAssetsOk returns a tuple with the CurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfo) GetCurrentAssetsOk() (*AssetsInfoCurrentAssets, bool) {
	if o == nil || IsNil(o.CurrentAssets) {
		return nil, false
	}
	return o.CurrentAssets, true
}

// HasCurrentAssets returns a boolean if a field has been set.
func (o *AssetsInfo) HasCurrentAssets() bool {
	if o != nil && !IsNil(o.CurrentAssets) {
		return true
	}

	return false
}

// SetCurrentAssets gets a reference to the given AssetsInfoCurrentAssets and assigns it to the CurrentAssets field.
func (o *AssetsInfo) SetCurrentAssets(v AssetsInfoCurrentAssets) {
	o.CurrentAssets = &v
}

// GetNonCurrentAssets returns the NonCurrentAssets field value if set, zero value otherwise.
func (o *AssetsInfo) GetNonCurrentAssets() AssetsInfoNonCurrentAssets {
	if o == nil || IsNil(o.NonCurrentAssets) {
		var ret AssetsInfoNonCurrentAssets
		return ret
	}
	return *o.NonCurrentAssets
}

// GetNonCurrentAssetsOk returns a tuple with the NonCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfo) GetNonCurrentAssetsOk() (*AssetsInfoNonCurrentAssets, bool) {
	if o == nil || IsNil(o.NonCurrentAssets) {
		return nil, false
	}
	return o.NonCurrentAssets, true
}

// HasNonCurrentAssets returns a boolean if a field has been set.
func (o *AssetsInfo) HasNonCurrentAssets() bool {
	if o != nil && !IsNil(o.NonCurrentAssets) {
		return true
	}

	return false
}

// SetNonCurrentAssets gets a reference to the given AssetsInfoNonCurrentAssets and assigns it to the NonCurrentAssets field.
func (o *AssetsInfo) SetNonCurrentAssets(v AssetsInfoNonCurrentAssets) {
	o.NonCurrentAssets = &v
}

// GetLiabilities returns the Liabilities field value if set, zero value otherwise.
func (o *AssetsInfo) GetLiabilities() AssetsInfoLiabilities {
	if o == nil || IsNil(o.Liabilities) {
		var ret AssetsInfoLiabilities
		return ret
	}
	return *o.Liabilities
}

// GetLiabilitiesOk returns a tuple with the Liabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfo) GetLiabilitiesOk() (*AssetsInfoLiabilities, bool) {
	if o == nil || IsNil(o.Liabilities) {
		return nil, false
	}
	return o.Liabilities, true
}

// HasLiabilities returns a boolean if a field has been set.
func (o *AssetsInfo) HasLiabilities() bool {
	if o != nil && !IsNil(o.Liabilities) {
		return true
	}

	return false
}

// SetLiabilities gets a reference to the given AssetsInfoLiabilities and assigns it to the Liabilities field.
func (o *AssetsInfo) SetLiabilities(v AssetsInfoLiabilities) {
	o.Liabilities = &v
}

func (o AssetsInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalAssets) {
		toSerialize["total_assets"] = o.TotalAssets
	}
	if !IsNil(o.CurrentAssets) {
		toSerialize["current_assets"] = o.CurrentAssets
	}
	if !IsNil(o.NonCurrentAssets) {
		toSerialize["non_current_assets"] = o.NonCurrentAssets
	}
	if !IsNil(o.Liabilities) {
		toSerialize["liabilities"] = o.Liabilities
	}
	return toSerialize, nil
}

type NullableAssetsInfo struct {
	value *AssetsInfo
	isSet bool
}

func (v NullableAssetsInfo) Get() *AssetsInfo {
	return v.value
}

func (v *NullableAssetsInfo) Set(val *AssetsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfo(val *AssetsInfo) *NullableAssetsInfo {
	return &NullableAssetsInfo{value: val, isSet: true}
}

func (v NullableAssetsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
