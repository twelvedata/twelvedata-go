// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemDividendsAndShares type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemDividendsAndShares{}

// IncomeStatementItemDividendsAndShares Dividends and shares information
type IncomeStatementItemDividendsAndShares struct {
	// Dividend per share
	DividendPerShare *float64 `json:"dividend_per_share,omitempty"`
	// Diluted average shares
	DilutedAverageShares *float64 `json:"diluted_average_shares,omitempty"`
	// Basic average shares
	BasicAverageShares *float64 `json:"basic_average_shares,omitempty"`
	// Preferred stock dividends
	PreferredStockDividends *float64 `json:"preferred_stock_dividends,omitempty"`
	// Other under preferred stock dividend
	OtherUnderPreferredStockDividend *float64 `json:"other_under_preferred_stock_dividend,omitempty"`
}

// NewIncomeStatementItemDividendsAndShares instantiates a new IncomeStatementItemDividendsAndShares object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemDividendsAndShares() *IncomeStatementItemDividendsAndShares {
	this := IncomeStatementItemDividendsAndShares{}
	return &this
}

// NewIncomeStatementItemDividendsAndSharesWithDefaults instantiates a new IncomeStatementItemDividendsAndShares object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemDividendsAndSharesWithDefaults() *IncomeStatementItemDividendsAndShares {
	this := IncomeStatementItemDividendsAndShares{}
	return &this
}

// GetDividendPerShare returns the DividendPerShare field value if set, zero value otherwise.
func (o *IncomeStatementItemDividendsAndShares) GetDividendPerShare() float64 {
	if o == nil || IsNil(o.DividendPerShare) {
		var ret float64
		return ret
	}
	return *o.DividendPerShare
}

// GetDividendPerShareOk returns a tuple with the DividendPerShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemDividendsAndShares) GetDividendPerShareOk() (*float64, bool) {
	if o == nil || IsNil(o.DividendPerShare) {
		return nil, false
	}
	return o.DividendPerShare, true
}

// HasDividendPerShare returns a boolean if a field has been set.
func (o *IncomeStatementItemDividendsAndShares) HasDividendPerShare() bool {
	if o != nil && !IsNil(o.DividendPerShare) {
		return true
	}

	return false
}

// SetDividendPerShare gets a reference to the given float64 and assigns it to the DividendPerShare field.
func (o *IncomeStatementItemDividendsAndShares) SetDividendPerShare(v float64) {
	o.DividendPerShare = &v
}

// GetDilutedAverageShares returns the DilutedAverageShares field value if set, zero value otherwise.
func (o *IncomeStatementItemDividendsAndShares) GetDilutedAverageShares() float64 {
	if o == nil || IsNil(o.DilutedAverageShares) {
		var ret float64
		return ret
	}
	return *o.DilutedAverageShares
}

// GetDilutedAverageSharesOk returns a tuple with the DilutedAverageShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemDividendsAndShares) GetDilutedAverageSharesOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedAverageShares) {
		return nil, false
	}
	return o.DilutedAverageShares, true
}

// HasDilutedAverageShares returns a boolean if a field has been set.
func (o *IncomeStatementItemDividendsAndShares) HasDilutedAverageShares() bool {
	if o != nil && !IsNil(o.DilutedAverageShares) {
		return true
	}

	return false
}

// SetDilutedAverageShares gets a reference to the given float64 and assigns it to the DilutedAverageShares field.
func (o *IncomeStatementItemDividendsAndShares) SetDilutedAverageShares(v float64) {
	o.DilutedAverageShares = &v
}

// GetBasicAverageShares returns the BasicAverageShares field value if set, zero value otherwise.
func (o *IncomeStatementItemDividendsAndShares) GetBasicAverageShares() float64 {
	if o == nil || IsNil(o.BasicAverageShares) {
		var ret float64
		return ret
	}
	return *o.BasicAverageShares
}

// GetBasicAverageSharesOk returns a tuple with the BasicAverageShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemDividendsAndShares) GetBasicAverageSharesOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicAverageShares) {
		return nil, false
	}
	return o.BasicAverageShares, true
}

// HasBasicAverageShares returns a boolean if a field has been set.
func (o *IncomeStatementItemDividendsAndShares) HasBasicAverageShares() bool {
	if o != nil && !IsNil(o.BasicAverageShares) {
		return true
	}

	return false
}

// SetBasicAverageShares gets a reference to the given float64 and assigns it to the BasicAverageShares field.
func (o *IncomeStatementItemDividendsAndShares) SetBasicAverageShares(v float64) {
	o.BasicAverageShares = &v
}

// GetPreferredStockDividends returns the PreferredStockDividends field value if set, zero value otherwise.
func (o *IncomeStatementItemDividendsAndShares) GetPreferredStockDividends() float64 {
	if o == nil || IsNil(o.PreferredStockDividends) {
		var ret float64
		return ret
	}
	return *o.PreferredStockDividends
}

// GetPreferredStockDividendsOk returns a tuple with the PreferredStockDividends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemDividendsAndShares) GetPreferredStockDividendsOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredStockDividends) {
		return nil, false
	}
	return o.PreferredStockDividends, true
}

// HasPreferredStockDividends returns a boolean if a field has been set.
func (o *IncomeStatementItemDividendsAndShares) HasPreferredStockDividends() bool {
	if o != nil && !IsNil(o.PreferredStockDividends) {
		return true
	}

	return false
}

// SetPreferredStockDividends gets a reference to the given float64 and assigns it to the PreferredStockDividends field.
func (o *IncomeStatementItemDividendsAndShares) SetPreferredStockDividends(v float64) {
	o.PreferredStockDividends = &v
}

// GetOtherUnderPreferredStockDividend returns the OtherUnderPreferredStockDividend field value if set, zero value otherwise.
func (o *IncomeStatementItemDividendsAndShares) GetOtherUnderPreferredStockDividend() float64 {
	if o == nil || IsNil(o.OtherUnderPreferredStockDividend) {
		var ret float64
		return ret
	}
	return *o.OtherUnderPreferredStockDividend
}

// GetOtherUnderPreferredStockDividendOk returns a tuple with the OtherUnderPreferredStockDividend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemDividendsAndShares) GetOtherUnderPreferredStockDividendOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherUnderPreferredStockDividend) {
		return nil, false
	}
	return o.OtherUnderPreferredStockDividend, true
}

// HasOtherUnderPreferredStockDividend returns a boolean if a field has been set.
func (o *IncomeStatementItemDividendsAndShares) HasOtherUnderPreferredStockDividend() bool {
	if o != nil && !IsNil(o.OtherUnderPreferredStockDividend) {
		return true
	}

	return false
}

// SetOtherUnderPreferredStockDividend gets a reference to the given float64 and assigns it to the OtherUnderPreferredStockDividend field.
func (o *IncomeStatementItemDividendsAndShares) SetOtherUnderPreferredStockDividend(v float64) {
	o.OtherUnderPreferredStockDividend = &v
}

func (o IncomeStatementItemDividendsAndShares) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemDividendsAndShares) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DividendPerShare) {
		toSerialize["dividend_per_share"] = o.DividendPerShare
	}
	if !IsNil(o.DilutedAverageShares) {
		toSerialize["diluted_average_shares"] = o.DilutedAverageShares
	}
	if !IsNil(o.BasicAverageShares) {
		toSerialize["basic_average_shares"] = o.BasicAverageShares
	}
	if !IsNil(o.PreferredStockDividends) {
		toSerialize["preferred_stock_dividends"] = o.PreferredStockDividends
	}
	if !IsNil(o.OtherUnderPreferredStockDividend) {
		toSerialize["other_under_preferred_stock_dividend"] = o.OtherUnderPreferredStockDividend
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemDividendsAndShares struct {
	value *IncomeStatementItemDividendsAndShares
	isSet bool
}

func (v NullableIncomeStatementItemDividendsAndShares) Get() *IncomeStatementItemDividendsAndShares {
	return v.value
}

func (v *NullableIncomeStatementItemDividendsAndShares) Set(val *IncomeStatementItemDividendsAndShares) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemDividendsAndShares) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemDividendsAndShares) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemDividendsAndShares(val *IncomeStatementItemDividendsAndShares) *NullableIncomeStatementItemDividendsAndShares {
	return &NullableIncomeStatementItemDividendsAndShares{value: val, isSet: true}
}

func (v NullableIncomeStatementItemDividendsAndShares) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemDividendsAndShares) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
