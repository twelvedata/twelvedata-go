/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner{}

// GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner struct for GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner
type GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner struct {
	// Sector category of a fund exposure
	Sector *string `json:"sector,omitempty"`
	// Weight of a fund exposure in a sector
	Weight *float64 `json:"weight,omitempty"`
}

// NewGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner instantiates a new GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner() *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner {
	this := GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner{}
	return &this
}

// NewGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInnerWithDefaults instantiates a new GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInnerWithDefaults() *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner {
	this := GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner{}
	return &this
}

// GetSector returns the Sector field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) GetSector() string {
	if o == nil || IsNil(o.Sector) {
		var ret string
		return ret
	}
	return *o.Sector
}

// GetSectorOk returns a tuple with the Sector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) GetSectorOk() (*string, bool) {
	if o == nil || IsNil(o.Sector) {
		return nil, false
	}
	return o.Sector, true
}

// HasSector returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) HasSector() bool {
	if o != nil && !IsNil(o.Sector) {
		return true
	}

	return false
}

// SetSector gets a reference to the given string and assigns it to the Sector field.
func (o *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) SetSector(v string) {
	o.Sector = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) SetWeight(v float64) {
	o.Weight = &v
}

func (o GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sector) {
		toSerialize["sector"] = o.Sector
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner struct {
	value *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) Get() *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) Set(val *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner(val *GetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) *NullableGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner {
	return &NullableGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionMajorMarketSectorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
