// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetRecommendations200ResponseTrends3MonthsAgo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecommendations200ResponseTrends3MonthsAgo{}

// GetRecommendations200ResponseTrends3MonthsAgo Three months ago recommendations
type GetRecommendations200ResponseTrends3MonthsAgo struct {
	// Number of analysts that give a strong buy recommendation
	StrongBuy *int64 `json:"strong_buy,omitempty"`
	// Number of analysts that give a buy recommendation
	Buy *int64 `json:"buy,omitempty"`
	// Number of analysts that give a hold recommendation
	Hold *int64 `json:"hold,omitempty"`
	// Number of analysts that give a sell recommendation
	Sell *int64 `json:"sell,omitempty"`
	// Number of analysts that give a strong sell recommendation
	StrongSell *int64 `json:"strong_sell,omitempty"`
}

// NewGetRecommendations200ResponseTrends3MonthsAgo instantiates a new GetRecommendations200ResponseTrends3MonthsAgo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecommendations200ResponseTrends3MonthsAgo() *GetRecommendations200ResponseTrends3MonthsAgo {
	this := GetRecommendations200ResponseTrends3MonthsAgo{}
	return &this
}

// NewGetRecommendations200ResponseTrends3MonthsAgoWithDefaults instantiates a new GetRecommendations200ResponseTrends3MonthsAgo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecommendations200ResponseTrends3MonthsAgoWithDefaults() *GetRecommendations200ResponseTrends3MonthsAgo {
	this := GetRecommendations200ResponseTrends3MonthsAgo{}
	return &this
}

// GetStrongBuy returns the StrongBuy field value if set, zero value otherwise.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) GetStrongBuy() int64 {
	if o == nil || IsNil(o.StrongBuy) {
		var ret int64
		return ret
	}
	return *o.StrongBuy
}

// GetStrongBuyOk returns a tuple with the StrongBuy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) GetStrongBuyOk() (*int64, bool) {
	if o == nil || IsNil(o.StrongBuy) {
		return nil, false
	}
	return o.StrongBuy, true
}

// HasStrongBuy returns a boolean if a field has been set.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) HasStrongBuy() bool {
	if o != nil && !IsNil(o.StrongBuy) {
		return true
	}

	return false
}

// SetStrongBuy gets a reference to the given int64 and assigns it to the StrongBuy field.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) SetStrongBuy(v int64) {
	o.StrongBuy = &v
}

// GetBuy returns the Buy field value if set, zero value otherwise.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) GetBuy() int64 {
	if o == nil || IsNil(o.Buy) {
		var ret int64
		return ret
	}
	return *o.Buy
}

// GetBuyOk returns a tuple with the Buy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) GetBuyOk() (*int64, bool) {
	if o == nil || IsNil(o.Buy) {
		return nil, false
	}
	return o.Buy, true
}

// HasBuy returns a boolean if a field has been set.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) HasBuy() bool {
	if o != nil && !IsNil(o.Buy) {
		return true
	}

	return false
}

// SetBuy gets a reference to the given int64 and assigns it to the Buy field.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) SetBuy(v int64) {
	o.Buy = &v
}

// GetHold returns the Hold field value if set, zero value otherwise.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) GetHold() int64 {
	if o == nil || IsNil(o.Hold) {
		var ret int64
		return ret
	}
	return *o.Hold
}

// GetHoldOk returns a tuple with the Hold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) GetHoldOk() (*int64, bool) {
	if o == nil || IsNil(o.Hold) {
		return nil, false
	}
	return o.Hold, true
}

// HasHold returns a boolean if a field has been set.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) HasHold() bool {
	if o != nil && !IsNil(o.Hold) {
		return true
	}

	return false
}

// SetHold gets a reference to the given int64 and assigns it to the Hold field.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) SetHold(v int64) {
	o.Hold = &v
}

// GetSell returns the Sell field value if set, zero value otherwise.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) GetSell() int64 {
	if o == nil || IsNil(o.Sell) {
		var ret int64
		return ret
	}
	return *o.Sell
}

// GetSellOk returns a tuple with the Sell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) GetSellOk() (*int64, bool) {
	if o == nil || IsNil(o.Sell) {
		return nil, false
	}
	return o.Sell, true
}

// HasSell returns a boolean if a field has been set.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) HasSell() bool {
	if o != nil && !IsNil(o.Sell) {
		return true
	}

	return false
}

// SetSell gets a reference to the given int64 and assigns it to the Sell field.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) SetSell(v int64) {
	o.Sell = &v
}

// GetStrongSell returns the StrongSell field value if set, zero value otherwise.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) GetStrongSell() int64 {
	if o == nil || IsNil(o.StrongSell) {
		var ret int64
		return ret
	}
	return *o.StrongSell
}

// GetStrongSellOk returns a tuple with the StrongSell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) GetStrongSellOk() (*int64, bool) {
	if o == nil || IsNil(o.StrongSell) {
		return nil, false
	}
	return o.StrongSell, true
}

// HasStrongSell returns a boolean if a field has been set.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) HasStrongSell() bool {
	if o != nil && !IsNil(o.StrongSell) {
		return true
	}

	return false
}

// SetStrongSell gets a reference to the given int64 and assigns it to the StrongSell field.
func (o *GetRecommendations200ResponseTrends3MonthsAgo) SetStrongSell(v int64) {
	o.StrongSell = &v
}

func (o GetRecommendations200ResponseTrends3MonthsAgo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecommendations200ResponseTrends3MonthsAgo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StrongBuy) {
		toSerialize["strong_buy"] = o.StrongBuy
	}
	if !IsNil(o.Buy) {
		toSerialize["buy"] = o.Buy
	}
	if !IsNil(o.Hold) {
		toSerialize["hold"] = o.Hold
	}
	if !IsNil(o.Sell) {
		toSerialize["sell"] = o.Sell
	}
	if !IsNil(o.StrongSell) {
		toSerialize["strong_sell"] = o.StrongSell
	}
	return toSerialize, nil
}

type NullableGetRecommendations200ResponseTrends3MonthsAgo struct {
	value *GetRecommendations200ResponseTrends3MonthsAgo
	isSet bool
}

func (v NullableGetRecommendations200ResponseTrends3MonthsAgo) Get() *GetRecommendations200ResponseTrends3MonthsAgo {
	return v.value
}

func (v *NullableGetRecommendations200ResponseTrends3MonthsAgo) Set(val *GetRecommendations200ResponseTrends3MonthsAgo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecommendations200ResponseTrends3MonthsAgo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecommendations200ResponseTrends3MonthsAgo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecommendations200ResponseTrends3MonthsAgo(val *GetRecommendations200ResponseTrends3MonthsAgo) *NullableGetRecommendations200ResponseTrends3MonthsAgo {
	return &NullableGetRecommendations200ResponseTrends3MonthsAgo{value: val, isSet: true}
}

func (v NullableGetRecommendations200ResponseTrends3MonthsAgo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecommendations200ResponseTrends3MonthsAgo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
