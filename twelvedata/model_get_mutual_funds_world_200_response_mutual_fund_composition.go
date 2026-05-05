// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundComposition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundComposition{}

// GetMutualFundsWorld200ResponseMutualFundComposition Composition of a mutual fund
type GetMutualFundsWorld200ResponseMutualFundComposition struct {
	// Breakdown of the fund’s portfolio by major industry sectors and their respective weights
	MajorMarketSectors []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner `json:"major_market_sectors,omitempty"`
	// Top holdings of the fund with their respective weights in the overall portfolio composition
	TopHoldings     []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner `json:"top_holdings,omitempty"`
	AssetAllocation *GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation   `json:"asset_allocation,omitempty"`
	BondBreakdown   *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown     `json:"bond_breakdown,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundComposition instantiates a new GetMutualFundsWorld200ResponseMutualFundComposition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundComposition() *GetMutualFundsWorld200ResponseMutualFundComposition {
	this := GetMutualFundsWorld200ResponseMutualFundComposition{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundCompositionWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundComposition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundCompositionWithDefaults() *GetMutualFundsWorld200ResponseMutualFundComposition {
	this := GetMutualFundsWorld200ResponseMutualFundComposition{}
	return &this
}

// GetMajorMarketSectors returns the MajorMarketSectors field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetMajorMarketSectors() []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner {
	if o == nil || IsNil(o.MajorMarketSectors) {
		var ret []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner
		return ret
	}
	return o.MajorMarketSectors
}

// GetMajorMarketSectorsOk returns a tuple with the MajorMarketSectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetMajorMarketSectorsOk() ([]GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner, bool) {
	if o == nil || IsNil(o.MajorMarketSectors) {
		return nil, false
	}
	return o.MajorMarketSectors, true
}

// HasMajorMarketSectors returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) HasMajorMarketSectors() bool {
	if o != nil && !IsNil(o.MajorMarketSectors) {
		return true
	}

	return false
}

// SetMajorMarketSectors gets a reference to the given []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner and assigns it to the MajorMarketSectors field.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) SetMajorMarketSectors(v []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) {
	o.MajorMarketSectors = v
}

// GetTopHoldings returns the TopHoldings field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetTopHoldings() []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner {
	if o == nil || IsNil(o.TopHoldings) {
		var ret []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner
		return ret
	}
	return o.TopHoldings
}

// GetTopHoldingsOk returns a tuple with the TopHoldings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetTopHoldingsOk() ([]GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner, bool) {
	if o == nil || IsNil(o.TopHoldings) {
		return nil, false
	}
	return o.TopHoldings, true
}

// HasTopHoldings returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) HasTopHoldings() bool {
	if o != nil && !IsNil(o.TopHoldings) {
		return true
	}

	return false
}

// SetTopHoldings gets a reference to the given []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner and assigns it to the TopHoldings field.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) SetTopHoldings(v []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) {
	o.TopHoldings = v
}

// GetAssetAllocation returns the AssetAllocation field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetAssetAllocation() GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation {
	if o == nil || IsNil(o.AssetAllocation) {
		var ret GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation
		return ret
	}
	return *o.AssetAllocation
}

// GetAssetAllocationOk returns a tuple with the AssetAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetAssetAllocationOk() (*GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation, bool) {
	if o == nil || IsNil(o.AssetAllocation) {
		return nil, false
	}
	return o.AssetAllocation, true
}

// HasAssetAllocation returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) HasAssetAllocation() bool {
	if o != nil && !IsNil(o.AssetAllocation) {
		return true
	}

	return false
}

// SetAssetAllocation gets a reference to the given GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation and assigns it to the AssetAllocation field.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) SetAssetAllocation(v GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation) {
	o.AssetAllocation = &v
}

// GetBondBreakdown returns the BondBreakdown field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetBondBreakdown() GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown {
	if o == nil || IsNil(o.BondBreakdown) {
		var ret GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown
		return ret
	}
	return *o.BondBreakdown
}

// GetBondBreakdownOk returns a tuple with the BondBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) GetBondBreakdownOk() (*GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown, bool) {
	if o == nil || IsNil(o.BondBreakdown) {
		return nil, false
	}
	return o.BondBreakdown, true
}

// HasBondBreakdown returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) HasBondBreakdown() bool {
	if o != nil && !IsNil(o.BondBreakdown) {
		return true
	}

	return false
}

// SetBondBreakdown gets a reference to the given GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown and assigns it to the BondBreakdown field.
func (o *GetMutualFundsWorld200ResponseMutualFundComposition) SetBondBreakdown(v GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown) {
	o.BondBreakdown = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundComposition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundComposition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MajorMarketSectors) {
		toSerialize["major_market_sectors"] = o.MajorMarketSectors
	}
	if !IsNil(o.TopHoldings) {
		toSerialize["top_holdings"] = o.TopHoldings
	}
	if !IsNil(o.AssetAllocation) {
		toSerialize["asset_allocation"] = o.AssetAllocation
	}
	if !IsNil(o.BondBreakdown) {
		toSerialize["bond_breakdown"] = o.BondBreakdown
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundComposition struct {
	value *GetMutualFundsWorld200ResponseMutualFundComposition
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundComposition) Get() *GetMutualFundsWorld200ResponseMutualFundComposition {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundComposition) Set(val *GetMutualFundsWorld200ResponseMutualFundComposition) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundComposition) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundComposition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundComposition(val *GetMutualFundsWorld200ResponseMutualFundComposition) *NullableGetMutualFundsWorld200ResponseMutualFundComposition {
	return &NullableGetMutualFundsWorld200ResponseMutualFundComposition{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundComposition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundComposition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
