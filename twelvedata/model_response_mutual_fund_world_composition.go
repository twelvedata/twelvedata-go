// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the ResponseMutualFundWorldComposition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMutualFundWorldComposition{}

// ResponseMutualFundWorldComposition Mutual fund composition
type ResponseMutualFundWorldComposition struct {
	// Breakdown of the fund’s portfolio by major industry sectors and their respective weights
	MajorMarketSectors []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner `json:"major_market_sectors,omitempty"`
	// Top holdings of the fund with their respective weights in the overall portfolio composition
	TopHoldings     []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner `json:"top_holdings,omitempty"`
	AssetAllocation *GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation   `json:"asset_allocation,omitempty"`
	BondBreakdown   *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown     `json:"bond_breakdown,omitempty"`
}

// NewResponseMutualFundWorldComposition instantiates a new ResponseMutualFundWorldComposition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMutualFundWorldComposition() *ResponseMutualFundWorldComposition {
	this := ResponseMutualFundWorldComposition{}
	return &this
}

// NewResponseMutualFundWorldCompositionWithDefaults instantiates a new ResponseMutualFundWorldComposition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMutualFundWorldCompositionWithDefaults() *ResponseMutualFundWorldComposition {
	this := ResponseMutualFundWorldComposition{}
	return &this
}

// GetMajorMarketSectors returns the MajorMarketSectors field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldComposition) GetMajorMarketSectors() []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner {
	if o == nil || IsNil(o.MajorMarketSectors) {
		var ret []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner
		return ret
	}
	return o.MajorMarketSectors
}

// GetMajorMarketSectorsOk returns a tuple with the MajorMarketSectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldComposition) GetMajorMarketSectorsOk() ([]GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner, bool) {
	if o == nil || IsNil(o.MajorMarketSectors) {
		return nil, false
	}
	return o.MajorMarketSectors, true
}

// HasMajorMarketSectors returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldComposition) HasMajorMarketSectors() bool {
	if o != nil && !IsNil(o.MajorMarketSectors) {
		return true
	}

	return false
}

// SetMajorMarketSectors gets a reference to the given []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner and assigns it to the MajorMarketSectors field.
func (o *ResponseMutualFundWorldComposition) SetMajorMarketSectors(v []GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) {
	o.MajorMarketSectors = v
}

// GetTopHoldings returns the TopHoldings field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldComposition) GetTopHoldings() []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner {
	if o == nil || IsNil(o.TopHoldings) {
		var ret []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner
		return ret
	}
	return o.TopHoldings
}

// GetTopHoldingsOk returns a tuple with the TopHoldings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldComposition) GetTopHoldingsOk() ([]GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner, bool) {
	if o == nil || IsNil(o.TopHoldings) {
		return nil, false
	}
	return o.TopHoldings, true
}

// HasTopHoldings returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldComposition) HasTopHoldings() bool {
	if o != nil && !IsNil(o.TopHoldings) {
		return true
	}

	return false
}

// SetTopHoldings gets a reference to the given []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner and assigns it to the TopHoldings field.
func (o *ResponseMutualFundWorldComposition) SetTopHoldings(v []GetMutualFundsWorld200ResponseMutualFundCompositionTopHoldingsInner) {
	o.TopHoldings = v
}

// GetAssetAllocation returns the AssetAllocation field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldComposition) GetAssetAllocation() GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation {
	if o == nil || IsNil(o.AssetAllocation) {
		var ret GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation
		return ret
	}
	return *o.AssetAllocation
}

// GetAssetAllocationOk returns a tuple with the AssetAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldComposition) GetAssetAllocationOk() (*GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation, bool) {
	if o == nil || IsNil(o.AssetAllocation) {
		return nil, false
	}
	return o.AssetAllocation, true
}

// HasAssetAllocation returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldComposition) HasAssetAllocation() bool {
	if o != nil && !IsNil(o.AssetAllocation) {
		return true
	}

	return false
}

// SetAssetAllocation gets a reference to the given GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation and assigns it to the AssetAllocation field.
func (o *ResponseMutualFundWorldComposition) SetAssetAllocation(v GetMutualFundsWorld200ResponseMutualFundCompositionAssetAllocation) {
	o.AssetAllocation = &v
}

// GetBondBreakdown returns the BondBreakdown field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldComposition) GetBondBreakdown() GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown {
	if o == nil || IsNil(o.BondBreakdown) {
		var ret GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown
		return ret
	}
	return *o.BondBreakdown
}

// GetBondBreakdownOk returns a tuple with the BondBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldComposition) GetBondBreakdownOk() (*GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown, bool) {
	if o == nil || IsNil(o.BondBreakdown) {
		return nil, false
	}
	return o.BondBreakdown, true
}

// HasBondBreakdown returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldComposition) HasBondBreakdown() bool {
	if o != nil && !IsNil(o.BondBreakdown) {
		return true
	}

	return false
}

// SetBondBreakdown gets a reference to the given GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown and assigns it to the BondBreakdown field.
func (o *ResponseMutualFundWorldComposition) SetBondBreakdown(v GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdown) {
	o.BondBreakdown = &v
}

func (o ResponseMutualFundWorldComposition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMutualFundWorldComposition) ToMap() (map[string]interface{}, error) {
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

type NullableResponseMutualFundWorldComposition struct {
	value *ResponseMutualFundWorldComposition
	isSet bool
}

func (v NullableResponseMutualFundWorldComposition) Get() *ResponseMutualFundWorldComposition {
	return v.value
}

func (v *NullableResponseMutualFundWorldComposition) Set(val *ResponseMutualFundWorldComposition) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMutualFundWorldComposition) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMutualFundWorldComposition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMutualFundWorldComposition(val *ResponseMutualFundWorldComposition) *NullableResponseMutualFundWorldComposition {
	return &NullableResponseMutualFundWorldComposition{value: val, isSet: true}
}

func (v NullableResponseMutualFundWorldComposition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMutualFundWorldComposition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
