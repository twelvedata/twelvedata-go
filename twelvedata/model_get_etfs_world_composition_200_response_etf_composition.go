/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorldComposition200ResponseEtfComposition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorldComposition200ResponseEtfComposition{}

// GetETFsWorldComposition200ResponseEtfComposition Composition of a etf
type GetETFsWorldComposition200ResponseEtfComposition struct {
	// Breakdown of the fund’s portfolio by major industry sectors and their respective weights
	MajorMarketSectors []GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner `json:"major_market_sectors,omitempty"`
	// Breakdown of the fund’s portfolio by country and their respective weights
	CountryAllocation []GetETFsWorld200ResponseEtfCompositionCountryAllocationInner `json:"country_allocation,omitempty"`
	AssetAllocation   *GetETFsWorld200ResponseEtfCompositionAssetAllocation         `json:"asset_allocation,omitempty"`
	// Top holdings of a fund with their respective weights in the overall portfolio composition
	TopHoldings   []GetETFsWorld200ResponseEtfCompositionTopHoldingsInner `json:"top_holdings,omitempty"`
	BondBreakdown *GetETFsWorld200ResponseEtfCompositionBondBreakdown     `json:"bond_breakdown,omitempty"`
}

// NewGetETFsWorldComposition200ResponseEtfComposition instantiates a new GetETFsWorldComposition200ResponseEtfComposition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorldComposition200ResponseEtfComposition() *GetETFsWorldComposition200ResponseEtfComposition {
	this := GetETFsWorldComposition200ResponseEtfComposition{}
	return &this
}

// NewGetETFsWorldComposition200ResponseEtfCompositionWithDefaults instantiates a new GetETFsWorldComposition200ResponseEtfComposition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorldComposition200ResponseEtfCompositionWithDefaults() *GetETFsWorldComposition200ResponseEtfComposition {
	this := GetETFsWorldComposition200ResponseEtfComposition{}
	return &this
}

// GetMajorMarketSectors returns the MajorMarketSectors field value if set, zero value otherwise.
func (o *GetETFsWorldComposition200ResponseEtfComposition) GetMajorMarketSectors() []GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner {
	if o == nil || IsNil(o.MajorMarketSectors) {
		var ret []GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner
		return ret
	}
	return o.MajorMarketSectors
}

// GetMajorMarketSectorsOk returns a tuple with the MajorMarketSectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorldComposition200ResponseEtfComposition) GetMajorMarketSectorsOk() ([]GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner, bool) {
	if o == nil || IsNil(o.MajorMarketSectors) {
		return nil, false
	}
	return o.MajorMarketSectors, true
}

// HasMajorMarketSectors returns a boolean if a field has been set.
func (o *GetETFsWorldComposition200ResponseEtfComposition) HasMajorMarketSectors() bool {
	if o != nil && !IsNil(o.MajorMarketSectors) {
		return true
	}

	return false
}

// SetMajorMarketSectors gets a reference to the given []GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner and assigns it to the MajorMarketSectors field.
func (o *GetETFsWorldComposition200ResponseEtfComposition) SetMajorMarketSectors(v []GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) {
	o.MajorMarketSectors = v
}

// GetCountryAllocation returns the CountryAllocation field value if set, zero value otherwise.
func (o *GetETFsWorldComposition200ResponseEtfComposition) GetCountryAllocation() []GetETFsWorld200ResponseEtfCompositionCountryAllocationInner {
	if o == nil || IsNil(o.CountryAllocation) {
		var ret []GetETFsWorld200ResponseEtfCompositionCountryAllocationInner
		return ret
	}
	return o.CountryAllocation
}

// GetCountryAllocationOk returns a tuple with the CountryAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorldComposition200ResponseEtfComposition) GetCountryAllocationOk() ([]GetETFsWorld200ResponseEtfCompositionCountryAllocationInner, bool) {
	if o == nil || IsNil(o.CountryAllocation) {
		return nil, false
	}
	return o.CountryAllocation, true
}

// HasCountryAllocation returns a boolean if a field has been set.
func (o *GetETFsWorldComposition200ResponseEtfComposition) HasCountryAllocation() bool {
	if o != nil && !IsNil(o.CountryAllocation) {
		return true
	}

	return false
}

// SetCountryAllocation gets a reference to the given []GetETFsWorld200ResponseEtfCompositionCountryAllocationInner and assigns it to the CountryAllocation field.
func (o *GetETFsWorldComposition200ResponseEtfComposition) SetCountryAllocation(v []GetETFsWorld200ResponseEtfCompositionCountryAllocationInner) {
	o.CountryAllocation = v
}

// GetAssetAllocation returns the AssetAllocation field value if set, zero value otherwise.
func (o *GetETFsWorldComposition200ResponseEtfComposition) GetAssetAllocation() GetETFsWorld200ResponseEtfCompositionAssetAllocation {
	if o == nil || IsNil(o.AssetAllocation) {
		var ret GetETFsWorld200ResponseEtfCompositionAssetAllocation
		return ret
	}
	return *o.AssetAllocation
}

// GetAssetAllocationOk returns a tuple with the AssetAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorldComposition200ResponseEtfComposition) GetAssetAllocationOk() (*GetETFsWorld200ResponseEtfCompositionAssetAllocation, bool) {
	if o == nil || IsNil(o.AssetAllocation) {
		return nil, false
	}
	return o.AssetAllocation, true
}

// HasAssetAllocation returns a boolean if a field has been set.
func (o *GetETFsWorldComposition200ResponseEtfComposition) HasAssetAllocation() bool {
	if o != nil && !IsNil(o.AssetAllocation) {
		return true
	}

	return false
}

// SetAssetAllocation gets a reference to the given GetETFsWorld200ResponseEtfCompositionAssetAllocation and assigns it to the AssetAllocation field.
func (o *GetETFsWorldComposition200ResponseEtfComposition) SetAssetAllocation(v GetETFsWorld200ResponseEtfCompositionAssetAllocation) {
	o.AssetAllocation = &v
}

// GetTopHoldings returns the TopHoldings field value if set, zero value otherwise.
func (o *GetETFsWorldComposition200ResponseEtfComposition) GetTopHoldings() []GetETFsWorld200ResponseEtfCompositionTopHoldingsInner {
	if o == nil || IsNil(o.TopHoldings) {
		var ret []GetETFsWorld200ResponseEtfCompositionTopHoldingsInner
		return ret
	}
	return o.TopHoldings
}

// GetTopHoldingsOk returns a tuple with the TopHoldings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorldComposition200ResponseEtfComposition) GetTopHoldingsOk() ([]GetETFsWorld200ResponseEtfCompositionTopHoldingsInner, bool) {
	if o == nil || IsNil(o.TopHoldings) {
		return nil, false
	}
	return o.TopHoldings, true
}

// HasTopHoldings returns a boolean if a field has been set.
func (o *GetETFsWorldComposition200ResponseEtfComposition) HasTopHoldings() bool {
	if o != nil && !IsNil(o.TopHoldings) {
		return true
	}

	return false
}

// SetTopHoldings gets a reference to the given []GetETFsWorld200ResponseEtfCompositionTopHoldingsInner and assigns it to the TopHoldings field.
func (o *GetETFsWorldComposition200ResponseEtfComposition) SetTopHoldings(v []GetETFsWorld200ResponseEtfCompositionTopHoldingsInner) {
	o.TopHoldings = v
}

// GetBondBreakdown returns the BondBreakdown field value if set, zero value otherwise.
func (o *GetETFsWorldComposition200ResponseEtfComposition) GetBondBreakdown() GetETFsWorld200ResponseEtfCompositionBondBreakdown {
	if o == nil || IsNil(o.BondBreakdown) {
		var ret GetETFsWorld200ResponseEtfCompositionBondBreakdown
		return ret
	}
	return *o.BondBreakdown
}

// GetBondBreakdownOk returns a tuple with the BondBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorldComposition200ResponseEtfComposition) GetBondBreakdownOk() (*GetETFsWorld200ResponseEtfCompositionBondBreakdown, bool) {
	if o == nil || IsNil(o.BondBreakdown) {
		return nil, false
	}
	return o.BondBreakdown, true
}

// HasBondBreakdown returns a boolean if a field has been set.
func (o *GetETFsWorldComposition200ResponseEtfComposition) HasBondBreakdown() bool {
	if o != nil && !IsNil(o.BondBreakdown) {
		return true
	}

	return false
}

// SetBondBreakdown gets a reference to the given GetETFsWorld200ResponseEtfCompositionBondBreakdown and assigns it to the BondBreakdown field.
func (o *GetETFsWorldComposition200ResponseEtfComposition) SetBondBreakdown(v GetETFsWorld200ResponseEtfCompositionBondBreakdown) {
	o.BondBreakdown = &v
}

func (o GetETFsWorldComposition200ResponseEtfComposition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorldComposition200ResponseEtfComposition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MajorMarketSectors) {
		toSerialize["major_market_sectors"] = o.MajorMarketSectors
	}
	if !IsNil(o.CountryAllocation) {
		toSerialize["country_allocation"] = o.CountryAllocation
	}
	if !IsNil(o.AssetAllocation) {
		toSerialize["asset_allocation"] = o.AssetAllocation
	}
	if !IsNil(o.TopHoldings) {
		toSerialize["top_holdings"] = o.TopHoldings
	}
	if !IsNil(o.BondBreakdown) {
		toSerialize["bond_breakdown"] = o.BondBreakdown
	}
	return toSerialize, nil
}

type NullableGetETFsWorldComposition200ResponseEtfComposition struct {
	value *GetETFsWorldComposition200ResponseEtfComposition
	isSet bool
}

func (v NullableGetETFsWorldComposition200ResponseEtfComposition) Get() *GetETFsWorldComposition200ResponseEtfComposition {
	return v.value
}

func (v *NullableGetETFsWorldComposition200ResponseEtfComposition) Set(val *GetETFsWorldComposition200ResponseEtfComposition) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorldComposition200ResponseEtfComposition) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorldComposition200ResponseEtfComposition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorldComposition200ResponseEtfComposition(val *GetETFsWorldComposition200ResponseEtfComposition) *NullableGetETFsWorldComposition200ResponseEtfComposition {
	return &NullableGetETFsWorldComposition200ResponseEtfComposition{value: val, isSet: true}
}

func (v NullableGetETFsWorldComposition200ResponseEtfComposition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorldComposition200ResponseEtfComposition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
