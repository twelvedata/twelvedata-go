// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetStatistics200ResponseStatisticsStockStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatisticsStockStatistics{}

// GetStatistics200ResponseStatisticsStockStatistics Stock statistics of the company
type GetStatistics200ResponseStatisticsStockStatistics struct {
	// Refers for the shares outstanding currently held by all its shareholders
	SharesOutstanding *float64 `json:"shares_outstanding,omitempty"`
	// Refers to floating stock is the number of public shares a company has available for trading on the open market
	FloatShares *float64 `json:"float_shares,omitempty"`
	// Refers to the average 10 days volume
	Avg10Volume *int64 `json:"avg_10_volume,omitempty"`
	// Refers to the average 90 days volume
	Avg90Volume *int64 `json:"avg_90_volume,omitempty"`
	// Refers to the number of shares that have been shorted
	SharesShort *int64 `json:"shares_short,omitempty"`
	// Refers to short ratio measure
	ShortRatio *float64 `json:"short_ratio,omitempty"`
	// Refers to the number of shorted shares divided by the number of shares outstanding
	ShortPercentOfSharesOutstanding *float64 `json:"short_percent_of_shares_outstanding,omitempty"`
	// Refers to the percentage of shares held by the company insiders
	PercentHeldByInsiders *float64 `json:"percent_held_by_insiders,omitempty"`
	// Refers to the percentage of shares held by the institutions
	PercentHeldByInstitutions *float64 `json:"percent_held_by_institutions,omitempty"`
}

// NewGetStatistics200ResponseStatisticsStockStatistics instantiates a new GetStatistics200ResponseStatisticsStockStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatisticsStockStatistics() *GetStatistics200ResponseStatisticsStockStatistics {
	this := GetStatistics200ResponseStatisticsStockStatistics{}
	return &this
}

// NewGetStatistics200ResponseStatisticsStockStatisticsWithDefaults instantiates a new GetStatistics200ResponseStatisticsStockStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsStockStatisticsWithDefaults() *GetStatistics200ResponseStatisticsStockStatistics {
	this := GetStatistics200ResponseStatisticsStockStatistics{}
	return &this
}

// GetSharesOutstanding returns the SharesOutstanding field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetSharesOutstanding() float64 {
	if o == nil || IsNil(o.SharesOutstanding) {
		var ret float64
		return ret
	}
	return *o.SharesOutstanding
}

// GetSharesOutstandingOk returns a tuple with the SharesOutstanding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetSharesOutstandingOk() (*float64, bool) {
	if o == nil || IsNil(o.SharesOutstanding) {
		return nil, false
	}
	return o.SharesOutstanding, true
}

// HasSharesOutstanding returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasSharesOutstanding() bool {
	if o != nil && !IsNil(o.SharesOutstanding) {
		return true
	}

	return false
}

// SetSharesOutstanding gets a reference to the given float64 and assigns it to the SharesOutstanding field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetSharesOutstanding(v float64) {
	o.SharesOutstanding = &v
}

// GetFloatShares returns the FloatShares field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetFloatShares() float64 {
	if o == nil || IsNil(o.FloatShares) {
		var ret float64
		return ret
	}
	return *o.FloatShares
}

// GetFloatSharesOk returns a tuple with the FloatShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetFloatSharesOk() (*float64, bool) {
	if o == nil || IsNil(o.FloatShares) {
		return nil, false
	}
	return o.FloatShares, true
}

// HasFloatShares returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasFloatShares() bool {
	if o != nil && !IsNil(o.FloatShares) {
		return true
	}

	return false
}

// SetFloatShares gets a reference to the given float64 and assigns it to the FloatShares field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetFloatShares(v float64) {
	o.FloatShares = &v
}

// GetAvg10Volume returns the Avg10Volume field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetAvg10Volume() int64 {
	if o == nil || IsNil(o.Avg10Volume) {
		var ret int64
		return ret
	}
	return *o.Avg10Volume
}

// GetAvg10VolumeOk returns a tuple with the Avg10Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetAvg10VolumeOk() (*int64, bool) {
	if o == nil || IsNil(o.Avg10Volume) {
		return nil, false
	}
	return o.Avg10Volume, true
}

// HasAvg10Volume returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasAvg10Volume() bool {
	if o != nil && !IsNil(o.Avg10Volume) {
		return true
	}

	return false
}

// SetAvg10Volume gets a reference to the given int64 and assigns it to the Avg10Volume field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetAvg10Volume(v int64) {
	o.Avg10Volume = &v
}

// GetAvg90Volume returns the Avg90Volume field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetAvg90Volume() int64 {
	if o == nil || IsNil(o.Avg90Volume) {
		var ret int64
		return ret
	}
	return *o.Avg90Volume
}

// GetAvg90VolumeOk returns a tuple with the Avg90Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetAvg90VolumeOk() (*int64, bool) {
	if o == nil || IsNil(o.Avg90Volume) {
		return nil, false
	}
	return o.Avg90Volume, true
}

// HasAvg90Volume returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasAvg90Volume() bool {
	if o != nil && !IsNil(o.Avg90Volume) {
		return true
	}

	return false
}

// SetAvg90Volume gets a reference to the given int64 and assigns it to the Avg90Volume field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetAvg90Volume(v int64) {
	o.Avg90Volume = &v
}

// GetSharesShort returns the SharesShort field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetSharesShort() int64 {
	if o == nil || IsNil(o.SharesShort) {
		var ret int64
		return ret
	}
	return *o.SharesShort
}

// GetSharesShortOk returns a tuple with the SharesShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetSharesShortOk() (*int64, bool) {
	if o == nil || IsNil(o.SharesShort) {
		return nil, false
	}
	return o.SharesShort, true
}

// HasSharesShort returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasSharesShort() bool {
	if o != nil && !IsNil(o.SharesShort) {
		return true
	}

	return false
}

// SetSharesShort gets a reference to the given int64 and assigns it to the SharesShort field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetSharesShort(v int64) {
	o.SharesShort = &v
}

// GetShortRatio returns the ShortRatio field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetShortRatio() float64 {
	if o == nil || IsNil(o.ShortRatio) {
		var ret float64
		return ret
	}
	return *o.ShortRatio
}

// GetShortRatioOk returns a tuple with the ShortRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetShortRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortRatio) {
		return nil, false
	}
	return o.ShortRatio, true
}

// HasShortRatio returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasShortRatio() bool {
	if o != nil && !IsNil(o.ShortRatio) {
		return true
	}

	return false
}

// SetShortRatio gets a reference to the given float64 and assigns it to the ShortRatio field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetShortRatio(v float64) {
	o.ShortRatio = &v
}

// GetShortPercentOfSharesOutstanding returns the ShortPercentOfSharesOutstanding field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetShortPercentOfSharesOutstanding() float64 {
	if o == nil || IsNil(o.ShortPercentOfSharesOutstanding) {
		var ret float64
		return ret
	}
	return *o.ShortPercentOfSharesOutstanding
}

// GetShortPercentOfSharesOutstandingOk returns a tuple with the ShortPercentOfSharesOutstanding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetShortPercentOfSharesOutstandingOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortPercentOfSharesOutstanding) {
		return nil, false
	}
	return o.ShortPercentOfSharesOutstanding, true
}

// HasShortPercentOfSharesOutstanding returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasShortPercentOfSharesOutstanding() bool {
	if o != nil && !IsNil(o.ShortPercentOfSharesOutstanding) {
		return true
	}

	return false
}

// SetShortPercentOfSharesOutstanding gets a reference to the given float64 and assigns it to the ShortPercentOfSharesOutstanding field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetShortPercentOfSharesOutstanding(v float64) {
	o.ShortPercentOfSharesOutstanding = &v
}

// GetPercentHeldByInsiders returns the PercentHeldByInsiders field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetPercentHeldByInsiders() float64 {
	if o == nil || IsNil(o.PercentHeldByInsiders) {
		var ret float64
		return ret
	}
	return *o.PercentHeldByInsiders
}

// GetPercentHeldByInsidersOk returns a tuple with the PercentHeldByInsiders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetPercentHeldByInsidersOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentHeldByInsiders) {
		return nil, false
	}
	return o.PercentHeldByInsiders, true
}

// HasPercentHeldByInsiders returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasPercentHeldByInsiders() bool {
	if o != nil && !IsNil(o.PercentHeldByInsiders) {
		return true
	}

	return false
}

// SetPercentHeldByInsiders gets a reference to the given float64 and assigns it to the PercentHeldByInsiders field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetPercentHeldByInsiders(v float64) {
	o.PercentHeldByInsiders = &v
}

// GetPercentHeldByInstitutions returns the PercentHeldByInstitutions field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetPercentHeldByInstitutions() float64 {
	if o == nil || IsNil(o.PercentHeldByInstitutions) {
		var ret float64
		return ret
	}
	return *o.PercentHeldByInstitutions
}

// GetPercentHeldByInstitutionsOk returns a tuple with the PercentHeldByInstitutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetPercentHeldByInstitutionsOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentHeldByInstitutions) {
		return nil, false
	}
	return o.PercentHeldByInstitutions, true
}

// HasPercentHeldByInstitutions returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasPercentHeldByInstitutions() bool {
	if o != nil && !IsNil(o.PercentHeldByInstitutions) {
		return true
	}

	return false
}

// SetPercentHeldByInstitutions gets a reference to the given float64 and assigns it to the PercentHeldByInstitutions field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetPercentHeldByInstitutions(v float64) {
	o.PercentHeldByInstitutions = &v
}

func (o GetStatistics200ResponseStatisticsStockStatistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatisticsStockStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SharesOutstanding) {
		toSerialize["shares_outstanding"] = o.SharesOutstanding
	}
	if !IsNil(o.FloatShares) {
		toSerialize["float_shares"] = o.FloatShares
	}
	if !IsNil(o.Avg10Volume) {
		toSerialize["avg_10_volume"] = o.Avg10Volume
	}
	if !IsNil(o.Avg90Volume) {
		toSerialize["avg_90_volume"] = o.Avg90Volume
	}
	if !IsNil(o.SharesShort) {
		toSerialize["shares_short"] = o.SharesShort
	}
	if !IsNil(o.ShortRatio) {
		toSerialize["short_ratio"] = o.ShortRatio
	}
	if !IsNil(o.ShortPercentOfSharesOutstanding) {
		toSerialize["short_percent_of_shares_outstanding"] = o.ShortPercentOfSharesOutstanding
	}
	if !IsNil(o.PercentHeldByInsiders) {
		toSerialize["percent_held_by_insiders"] = o.PercentHeldByInsiders
	}
	if !IsNil(o.PercentHeldByInstitutions) {
		toSerialize["percent_held_by_institutions"] = o.PercentHeldByInstitutions
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatisticsStockStatistics struct {
	value *GetStatistics200ResponseStatisticsStockStatistics
	isSet bool
}

func (v NullableGetStatistics200ResponseStatisticsStockStatistics) Get() *GetStatistics200ResponseStatisticsStockStatistics {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatisticsStockStatistics) Set(val *GetStatistics200ResponseStatisticsStockStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatisticsStockStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatisticsStockStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatisticsStockStatistics(val *GetStatistics200ResponseStatisticsStockStatistics) *NullableGetStatistics200ResponseStatisticsStockStatistics {
	return &NullableGetStatistics200ResponseStatisticsStockStatistics{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatisticsStockStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatisticsStockStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
