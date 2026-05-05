// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the OptionSide type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionSide{}

// OptionSide struct for OptionSide
type OptionSide struct {
	Ask               *float64 `json:"ask,omitempty"`
	Bid               *float64 `json:"bid,omitempty"`
	Change            *float64 `json:"change,omitempty"`
	ContractName      *string  `json:"contract_name,omitempty"`
	ImpliedVolatility *float64 `json:"implied_volatility,omitempty"`
	InTheMoney        *bool    `json:"in_the_money,omitempty"`
	LastPrice         *float64 `json:"last_price,omitempty"`
	LastTradeDate     *string  `json:"last_trade_date,omitempty"`
	OpenInterest      *int64   `json:"open_interest,omitempty"`
	OptionId          *string  `json:"option_id,omitempty"`
	PercentChange     *float64 `json:"percent_change,omitempty"`
	Strike            *float64 `json:"strike,omitempty"`
	Volume            *int64   `json:"volume,omitempty"`
}

// NewOptionSide instantiates a new OptionSide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionSide() *OptionSide {
	this := OptionSide{}
	return &this
}

// NewOptionSideWithDefaults instantiates a new OptionSide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionSideWithDefaults() *OptionSide {
	this := OptionSide{}
	return &this
}

// GetAsk returns the Ask field value if set, zero value otherwise.
func (o *OptionSide) GetAsk() float64 {
	if o == nil || IsNil(o.Ask) {
		var ret float64
		return ret
	}
	return *o.Ask
}

// GetAskOk returns a tuple with the Ask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetAskOk() (*float64, bool) {
	if o == nil || IsNil(o.Ask) {
		return nil, false
	}
	return o.Ask, true
}

// HasAsk returns a boolean if a field has been set.
func (o *OptionSide) HasAsk() bool {
	if o != nil && !IsNil(o.Ask) {
		return true
	}

	return false
}

// SetAsk gets a reference to the given float64 and assigns it to the Ask field.
func (o *OptionSide) SetAsk(v float64) {
	o.Ask = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *OptionSide) GetBid() float64 {
	if o == nil || IsNil(o.Bid) {
		var ret float64
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetBidOk() (*float64, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *OptionSide) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float64 and assigns it to the Bid field.
func (o *OptionSide) SetBid(v float64) {
	o.Bid = &v
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *OptionSide) GetChange() float64 {
	if o == nil || IsNil(o.Change) {
		var ret float64
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetChangeOk() (*float64, bool) {
	if o == nil || IsNil(o.Change) {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *OptionSide) HasChange() bool {
	if o != nil && !IsNil(o.Change) {
		return true
	}

	return false
}

// SetChange gets a reference to the given float64 and assigns it to the Change field.
func (o *OptionSide) SetChange(v float64) {
	o.Change = &v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *OptionSide) GetContractName() string {
	if o == nil || IsNil(o.ContractName) {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetContractNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContractName) {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *OptionSide) HasContractName() bool {
	if o != nil && !IsNil(o.ContractName) {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *OptionSide) SetContractName(v string) {
	o.ContractName = &v
}

// GetImpliedVolatility returns the ImpliedVolatility field value if set, zero value otherwise.
func (o *OptionSide) GetImpliedVolatility() float64 {
	if o == nil || IsNil(o.ImpliedVolatility) {
		var ret float64
		return ret
	}
	return *o.ImpliedVolatility
}

// GetImpliedVolatilityOk returns a tuple with the ImpliedVolatility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetImpliedVolatilityOk() (*float64, bool) {
	if o == nil || IsNil(o.ImpliedVolatility) {
		return nil, false
	}
	return o.ImpliedVolatility, true
}

// HasImpliedVolatility returns a boolean if a field has been set.
func (o *OptionSide) HasImpliedVolatility() bool {
	if o != nil && !IsNil(o.ImpliedVolatility) {
		return true
	}

	return false
}

// SetImpliedVolatility gets a reference to the given float64 and assigns it to the ImpliedVolatility field.
func (o *OptionSide) SetImpliedVolatility(v float64) {
	o.ImpliedVolatility = &v
}

// GetInTheMoney returns the InTheMoney field value if set, zero value otherwise.
func (o *OptionSide) GetInTheMoney() bool {
	if o == nil || IsNil(o.InTheMoney) {
		var ret bool
		return ret
	}
	return *o.InTheMoney
}

// GetInTheMoneyOk returns a tuple with the InTheMoney field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetInTheMoneyOk() (*bool, bool) {
	if o == nil || IsNil(o.InTheMoney) {
		return nil, false
	}
	return o.InTheMoney, true
}

// HasInTheMoney returns a boolean if a field has been set.
func (o *OptionSide) HasInTheMoney() bool {
	if o != nil && !IsNil(o.InTheMoney) {
		return true
	}

	return false
}

// SetInTheMoney gets a reference to the given bool and assigns it to the InTheMoney field.
func (o *OptionSide) SetInTheMoney(v bool) {
	o.InTheMoney = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *OptionSide) GetLastPrice() float64 {
	if o == nil || IsNil(o.LastPrice) {
		var ret float64
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetLastPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.LastPrice) {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *OptionSide) HasLastPrice() bool {
	if o != nil && !IsNil(o.LastPrice) {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given float64 and assigns it to the LastPrice field.
func (o *OptionSide) SetLastPrice(v float64) {
	o.LastPrice = &v
}

// GetLastTradeDate returns the LastTradeDate field value if set, zero value otherwise.
func (o *OptionSide) GetLastTradeDate() string {
	if o == nil || IsNil(o.LastTradeDate) {
		var ret string
		return ret
	}
	return *o.LastTradeDate
}

// GetLastTradeDateOk returns a tuple with the LastTradeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetLastTradeDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastTradeDate) {
		return nil, false
	}
	return o.LastTradeDate, true
}

// HasLastTradeDate returns a boolean if a field has been set.
func (o *OptionSide) HasLastTradeDate() bool {
	if o != nil && !IsNil(o.LastTradeDate) {
		return true
	}

	return false
}

// SetLastTradeDate gets a reference to the given string and assigns it to the LastTradeDate field.
func (o *OptionSide) SetLastTradeDate(v string) {
	o.LastTradeDate = &v
}

// GetOpenInterest returns the OpenInterest field value if set, zero value otherwise.
func (o *OptionSide) GetOpenInterest() int64 {
	if o == nil || IsNil(o.OpenInterest) {
		var ret int64
		return ret
	}
	return *o.OpenInterest
}

// GetOpenInterestOk returns a tuple with the OpenInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetOpenInterestOk() (*int64, bool) {
	if o == nil || IsNil(o.OpenInterest) {
		return nil, false
	}
	return o.OpenInterest, true
}

// HasOpenInterest returns a boolean if a field has been set.
func (o *OptionSide) HasOpenInterest() bool {
	if o != nil && !IsNil(o.OpenInterest) {
		return true
	}

	return false
}

// SetOpenInterest gets a reference to the given int64 and assigns it to the OpenInterest field.
func (o *OptionSide) SetOpenInterest(v int64) {
	o.OpenInterest = &v
}

// GetOptionId returns the OptionId field value if set, zero value otherwise.
func (o *OptionSide) GetOptionId() string {
	if o == nil || IsNil(o.OptionId) {
		var ret string
		return ret
	}
	return *o.OptionId
}

// GetOptionIdOk returns a tuple with the OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetOptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptionId) {
		return nil, false
	}
	return o.OptionId, true
}

// HasOptionId returns a boolean if a field has been set.
func (o *OptionSide) HasOptionId() bool {
	if o != nil && !IsNil(o.OptionId) {
		return true
	}

	return false
}

// SetOptionId gets a reference to the given string and assigns it to the OptionId field.
func (o *OptionSide) SetOptionId(v string) {
	o.OptionId = &v
}

// GetPercentChange returns the PercentChange field value if set, zero value otherwise.
func (o *OptionSide) GetPercentChange() float64 {
	if o == nil || IsNil(o.PercentChange) {
		var ret float64
		return ret
	}
	return *o.PercentChange
}

// GetPercentChangeOk returns a tuple with the PercentChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetPercentChangeOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentChange) {
		return nil, false
	}
	return o.PercentChange, true
}

// HasPercentChange returns a boolean if a field has been set.
func (o *OptionSide) HasPercentChange() bool {
	if o != nil && !IsNil(o.PercentChange) {
		return true
	}

	return false
}

// SetPercentChange gets a reference to the given float64 and assigns it to the PercentChange field.
func (o *OptionSide) SetPercentChange(v float64) {
	o.PercentChange = &v
}

// GetStrike returns the Strike field value if set, zero value otherwise.
func (o *OptionSide) GetStrike() float64 {
	if o == nil || IsNil(o.Strike) {
		var ret float64
		return ret
	}
	return *o.Strike
}

// GetStrikeOk returns a tuple with the Strike field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetStrikeOk() (*float64, bool) {
	if o == nil || IsNil(o.Strike) {
		return nil, false
	}
	return o.Strike, true
}

// HasStrike returns a boolean if a field has been set.
func (o *OptionSide) HasStrike() bool {
	if o != nil && !IsNil(o.Strike) {
		return true
	}

	return false
}

// SetStrike gets a reference to the given float64 and assigns it to the Strike field.
func (o *OptionSide) SetStrike(v float64) {
	o.Strike = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *OptionSide) GetVolume() int64 {
	if o == nil || IsNil(o.Volume) {
		var ret int64
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetVolumeOk() (*int64, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *OptionSide) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given int64 and assigns it to the Volume field.
func (o *OptionSide) SetVolume(v int64) {
	o.Volume = &v
}

func (o OptionSide) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionSide) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ask) {
		toSerialize["ask"] = o.Ask
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	if !IsNil(o.Change) {
		toSerialize["change"] = o.Change
	}
	if !IsNil(o.ContractName) {
		toSerialize["contract_name"] = o.ContractName
	}
	if !IsNil(o.ImpliedVolatility) {
		toSerialize["implied_volatility"] = o.ImpliedVolatility
	}
	if !IsNil(o.InTheMoney) {
		toSerialize["in_the_money"] = o.InTheMoney
	}
	if !IsNil(o.LastPrice) {
		toSerialize["last_price"] = o.LastPrice
	}
	if !IsNil(o.LastTradeDate) {
		toSerialize["last_trade_date"] = o.LastTradeDate
	}
	if !IsNil(o.OpenInterest) {
		toSerialize["open_interest"] = o.OpenInterest
	}
	if !IsNil(o.OptionId) {
		toSerialize["option_id"] = o.OptionId
	}
	if !IsNil(o.PercentChange) {
		toSerialize["percent_change"] = o.PercentChange
	}
	if !IsNil(o.Strike) {
		toSerialize["strike"] = o.Strike
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableOptionSide struct {
	value *OptionSide
	isSet bool
}

func (v NullableOptionSide) Get() *OptionSide {
	return v.value
}

func (v *NullableOptionSide) Set(val *OptionSide) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionSide) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionSide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionSide(val *OptionSide) *NullableOptionSide {
	return &NullableOptionSide{value: val, isSet: true}
}

func (v NullableOptionSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionSide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
