/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets{}

// AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets Goodwill and other intangible assets information
type AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets struct {
	// Goodwill
	Goodwill *float64 `json:"goodwill,omitempty"`
	// Other intangible assets
	OtherIntangibleAssets *float64 `json:"other_intangible_assets,omitempty"`
	// Total goodwill and intangible assets
	TotalGoodwillAndIntangibleAssets *float64 `json:"total_goodwill_and_intangible_assets,omitempty"`
}

// NewAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets instantiates a new AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets() *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets {
	this := AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets{}
	return &this
}

// NewAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssetsWithDefaults instantiates a new AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssetsWithDefaults() *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets {
	this := AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets{}
	return &this
}

// GetGoodwill returns the Goodwill field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) GetGoodwill() float64 {
	if o == nil || IsNil(o.Goodwill) {
		var ret float64
		return ret
	}
	return *o.Goodwill
}

// GetGoodwillOk returns a tuple with the Goodwill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) GetGoodwillOk() (*float64, bool) {
	if o == nil || IsNil(o.Goodwill) {
		return nil, false
	}
	return o.Goodwill, true
}

// HasGoodwill returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) HasGoodwill() bool {
	if o != nil && !IsNil(o.Goodwill) {
		return true
	}

	return false
}

// SetGoodwill gets a reference to the given float64 and assigns it to the Goodwill field.
func (o *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) SetGoodwill(v float64) {
	o.Goodwill = &v
}

// GetOtherIntangibleAssets returns the OtherIntangibleAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) GetOtherIntangibleAssets() float64 {
	if o == nil || IsNil(o.OtherIntangibleAssets) {
		var ret float64
		return ret
	}
	return *o.OtherIntangibleAssets
}

// GetOtherIntangibleAssetsOk returns a tuple with the OtherIntangibleAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) GetOtherIntangibleAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherIntangibleAssets) {
		return nil, false
	}
	return o.OtherIntangibleAssets, true
}

// HasOtherIntangibleAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) HasOtherIntangibleAssets() bool {
	if o != nil && !IsNil(o.OtherIntangibleAssets) {
		return true
	}

	return false
}

// SetOtherIntangibleAssets gets a reference to the given float64 and assigns it to the OtherIntangibleAssets field.
func (o *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) SetOtherIntangibleAssets(v float64) {
	o.OtherIntangibleAssets = &v
}

// GetTotalGoodwillAndIntangibleAssets returns the TotalGoodwillAndIntangibleAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) GetTotalGoodwillAndIntangibleAssets() float64 {
	if o == nil || IsNil(o.TotalGoodwillAndIntangibleAssets) {
		var ret float64
		return ret
	}
	return *o.TotalGoodwillAndIntangibleAssets
}

// GetTotalGoodwillAndIntangibleAssetsOk returns a tuple with the TotalGoodwillAndIntangibleAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) GetTotalGoodwillAndIntangibleAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalGoodwillAndIntangibleAssets) {
		return nil, false
	}
	return o.TotalGoodwillAndIntangibleAssets, true
}

// HasTotalGoodwillAndIntangibleAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) HasTotalGoodwillAndIntangibleAssets() bool {
	if o != nil && !IsNil(o.TotalGoodwillAndIntangibleAssets) {
		return true
	}

	return false
}

// SetTotalGoodwillAndIntangibleAssets gets a reference to the given float64 and assigns it to the TotalGoodwillAndIntangibleAssets field.
func (o *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) SetTotalGoodwillAndIntangibleAssets(v float64) {
	o.TotalGoodwillAndIntangibleAssets = &v
}

func (o AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Goodwill) {
		toSerialize["goodwill"] = o.Goodwill
	}
	if !IsNil(o.OtherIntangibleAssets) {
		toSerialize["other_intangible_assets"] = o.OtherIntangibleAssets
	}
	if !IsNil(o.TotalGoodwillAndIntangibleAssets) {
		toSerialize["total_goodwill_and_intangible_assets"] = o.TotalGoodwillAndIntangibleAssets
	}
	return toSerialize, nil
}

type NullableAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets struct {
	value *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets
	isSet bool
}

func (v NullableAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) Get() *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets {
	return v.value
}

func (v *NullableAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) Set(val *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets(val *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) *NullableAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets {
	return &NullableAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets{value: val, isSet: true}
}

func (v NullableAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
