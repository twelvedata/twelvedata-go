// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfCompositionAssetAllocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfCompositionAssetAllocation{}

// GetETFsWorld200ResponseEtfCompositionAssetAllocation Asset allocation of a fund by different asset classes and their respective weights
type GetETFsWorld200ResponseEtfCompositionAssetAllocation struct {
	// Percentage of overall portfolio composition in cash
	Cash *float64 `json:"cash,omitempty"`
	// Percentage of overall portfolio composition in stocks
	Stocks *float64 `json:"stocks,omitempty"`
	// Percentage of overall portfolio composition in preferred stocks
	PreferredStocks *float64 `json:"preferred_stocks,omitempty"`
	// Percentage of overall portfolio composition in convertable securities
	Convertables *float64 `json:"convertables,omitempty"`
	// Percentage of overall portfolio composition in bond
	Bonds *float64 `json:"bonds,omitempty"`
	// Percentage of overall portfolio composition in other forms of holding
	Others *float64 `json:"others,omitempty"`
}

// NewGetETFsWorld200ResponseEtfCompositionAssetAllocation instantiates a new GetETFsWorld200ResponseEtfCompositionAssetAllocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfCompositionAssetAllocation() *GetETFsWorld200ResponseEtfCompositionAssetAllocation {
	this := GetETFsWorld200ResponseEtfCompositionAssetAllocation{}
	return &this
}

// NewGetETFsWorld200ResponseEtfCompositionAssetAllocationWithDefaults instantiates a new GetETFsWorld200ResponseEtfCompositionAssetAllocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfCompositionAssetAllocationWithDefaults() *GetETFsWorld200ResponseEtfCompositionAssetAllocation {
	this := GetETFsWorld200ResponseEtfCompositionAssetAllocation{}
	return &this
}

// GetCash returns the Cash field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) GetCash() float64 {
	if o == nil || IsNil(o.Cash) {
		var ret float64
		return ret
	}
	return *o.Cash
}

// GetCashOk returns a tuple with the Cash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) GetCashOk() (*float64, bool) {
	if o == nil || IsNil(o.Cash) {
		return nil, false
	}
	return o.Cash, true
}

// HasCash returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) HasCash() bool {
	if o != nil && !IsNil(o.Cash) {
		return true
	}

	return false
}

// SetCash gets a reference to the given float64 and assigns it to the Cash field.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) SetCash(v float64) {
	o.Cash = &v
}

// GetStocks returns the Stocks field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) GetStocks() float64 {
	if o == nil || IsNil(o.Stocks) {
		var ret float64
		return ret
	}
	return *o.Stocks
}

// GetStocksOk returns a tuple with the Stocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) GetStocksOk() (*float64, bool) {
	if o == nil || IsNil(o.Stocks) {
		return nil, false
	}
	return o.Stocks, true
}

// HasStocks returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) HasStocks() bool {
	if o != nil && !IsNil(o.Stocks) {
		return true
	}

	return false
}

// SetStocks gets a reference to the given float64 and assigns it to the Stocks field.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) SetStocks(v float64) {
	o.Stocks = &v
}

// GetPreferredStocks returns the PreferredStocks field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) GetPreferredStocks() float64 {
	if o == nil || IsNil(o.PreferredStocks) {
		var ret float64
		return ret
	}
	return *o.PreferredStocks
}

// GetPreferredStocksOk returns a tuple with the PreferredStocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) GetPreferredStocksOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredStocks) {
		return nil, false
	}
	return o.PreferredStocks, true
}

// HasPreferredStocks returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) HasPreferredStocks() bool {
	if o != nil && !IsNil(o.PreferredStocks) {
		return true
	}

	return false
}

// SetPreferredStocks gets a reference to the given float64 and assigns it to the PreferredStocks field.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) SetPreferredStocks(v float64) {
	o.PreferredStocks = &v
}

// GetConvertables returns the Convertables field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) GetConvertables() float64 {
	if o == nil || IsNil(o.Convertables) {
		var ret float64
		return ret
	}
	return *o.Convertables
}

// GetConvertablesOk returns a tuple with the Convertables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) GetConvertablesOk() (*float64, bool) {
	if o == nil || IsNil(o.Convertables) {
		return nil, false
	}
	return o.Convertables, true
}

// HasConvertables returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) HasConvertables() bool {
	if o != nil && !IsNil(o.Convertables) {
		return true
	}

	return false
}

// SetConvertables gets a reference to the given float64 and assigns it to the Convertables field.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) SetConvertables(v float64) {
	o.Convertables = &v
}

// GetBonds returns the Bonds field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) GetBonds() float64 {
	if o == nil || IsNil(o.Bonds) {
		var ret float64
		return ret
	}
	return *o.Bonds
}

// GetBondsOk returns a tuple with the Bonds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) GetBondsOk() (*float64, bool) {
	if o == nil || IsNil(o.Bonds) {
		return nil, false
	}
	return o.Bonds, true
}

// HasBonds returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) HasBonds() bool {
	if o != nil && !IsNil(o.Bonds) {
		return true
	}

	return false
}

// SetBonds gets a reference to the given float64 and assigns it to the Bonds field.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) SetBonds(v float64) {
	o.Bonds = &v
}

// GetOthers returns the Others field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) GetOthers() float64 {
	if o == nil || IsNil(o.Others) {
		var ret float64
		return ret
	}
	return *o.Others
}

// GetOthersOk returns a tuple with the Others field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) GetOthersOk() (*float64, bool) {
	if o == nil || IsNil(o.Others) {
		return nil, false
	}
	return o.Others, true
}

// HasOthers returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) HasOthers() bool {
	if o != nil && !IsNil(o.Others) {
		return true
	}

	return false
}

// SetOthers gets a reference to the given float64 and assigns it to the Others field.
func (o *GetETFsWorld200ResponseEtfCompositionAssetAllocation) SetOthers(v float64) {
	o.Others = &v
}

func (o GetETFsWorld200ResponseEtfCompositionAssetAllocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfCompositionAssetAllocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cash) {
		toSerialize["cash"] = o.Cash
	}
	if !IsNil(o.Stocks) {
		toSerialize["stocks"] = o.Stocks
	}
	if !IsNil(o.PreferredStocks) {
		toSerialize["preferred_stocks"] = o.PreferredStocks
	}
	if !IsNil(o.Convertables) {
		toSerialize["convertables"] = o.Convertables
	}
	if !IsNil(o.Bonds) {
		toSerialize["bonds"] = o.Bonds
	}
	if !IsNil(o.Others) {
		toSerialize["others"] = o.Others
	}
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtfCompositionAssetAllocation struct {
	value *GetETFsWorld200ResponseEtfCompositionAssetAllocation
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfCompositionAssetAllocation) Get() *GetETFsWorld200ResponseEtfCompositionAssetAllocation {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionAssetAllocation) Set(val *GetETFsWorld200ResponseEtfCompositionAssetAllocation) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfCompositionAssetAllocation) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionAssetAllocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfCompositionAssetAllocation(val *GetETFsWorld200ResponseEtfCompositionAssetAllocation) *NullableGetETFsWorld200ResponseEtfCompositionAssetAllocation {
	return &NullableGetETFsWorld200ResponseEtfCompositionAssetAllocation{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfCompositionAssetAllocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionAssetAllocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
