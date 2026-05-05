// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner{}

// GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner struct for GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner
type GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner struct {
	// Sector category of a fund exposure
	Sector *string `json:"sector,omitempty"`
	// Weight of a fund exposure in a sector
	Weight *float64 `json:"weight,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner instantiates a new GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner() *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner {
	this := GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInnerWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInnerWithDefaults() *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner {
	this := GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner{}
	return &this
}

// GetSector returns the Sector field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) GetSector() string {
	if o == nil || IsNil(o.Sector) {
		var ret string
		return ret
	}
	return *o.Sector
}

// GetSectorOk returns a tuple with the Sector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) GetSectorOk() (*string, bool) {
	if o == nil || IsNil(o.Sector) {
		return nil, false
	}
	return o.Sector, true
}

// HasSector returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) HasSector() bool {
	if o != nil && !IsNil(o.Sector) {
		return true
	}

	return false
}

// SetSector gets a reference to the given string and assigns it to the Sector field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) SetSector(v string) {
	o.Sector = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) SetWeight(v float64) {
	o.Weight = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sector) {
		toSerialize["sector"] = o.Sector
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner struct {
	value *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) Get() *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) Set(val *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner(val *GetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) *NullableGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner {
	return &NullableGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionMajorMarketSectorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
