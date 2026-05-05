// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner{}

// GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner struct for GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner
type GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner struct {
	// Period of trailing returns
	Period *string `json:"period,omitempty"`
	// Fund returns (%) generated over a given period
	ShareClassReturn *float64 `json:"share_class_return,omitempty"`
	// Same category average returns (%) generated over a given period
	CategoryReturn *float64 `json:"category_return,omitempty"`
	// Rank of a fund in category by total returns
	RankInCategory *int64 `json:"rank_in_category,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner instantiates a new GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner() *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner {
	this := GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInnerWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInnerWithDefaults() *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner {
	this := GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) SetPeriod(v string) {
	o.Period = &v
}

// GetShareClassReturn returns the ShareClassReturn field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) GetShareClassReturn() float64 {
	if o == nil || IsNil(o.ShareClassReturn) {
		var ret float64
		return ret
	}
	return *o.ShareClassReturn
}

// GetShareClassReturnOk returns a tuple with the ShareClassReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) GetShareClassReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.ShareClassReturn) {
		return nil, false
	}
	return o.ShareClassReturn, true
}

// HasShareClassReturn returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) HasShareClassReturn() bool {
	if o != nil && !IsNil(o.ShareClassReturn) {
		return true
	}

	return false
}

// SetShareClassReturn gets a reference to the given float64 and assigns it to the ShareClassReturn field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) SetShareClassReturn(v float64) {
	o.ShareClassReturn = &v
}

// GetCategoryReturn returns the CategoryReturn field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) GetCategoryReturn() float64 {
	if o == nil || IsNil(o.CategoryReturn) {
		var ret float64
		return ret
	}
	return *o.CategoryReturn
}

// GetCategoryReturnOk returns a tuple with the CategoryReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) GetCategoryReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.CategoryReturn) {
		return nil, false
	}
	return o.CategoryReturn, true
}

// HasCategoryReturn returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) HasCategoryReturn() bool {
	if o != nil && !IsNil(o.CategoryReturn) {
		return true
	}

	return false
}

// SetCategoryReturn gets a reference to the given float64 and assigns it to the CategoryReturn field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) SetCategoryReturn(v float64) {
	o.CategoryReturn = &v
}

// GetRankInCategory returns the RankInCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) GetRankInCategory() int64 {
	if o == nil || IsNil(o.RankInCategory) {
		var ret int64
		return ret
	}
	return *o.RankInCategory
}

// GetRankInCategoryOk returns a tuple with the RankInCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) GetRankInCategoryOk() (*int64, bool) {
	if o == nil || IsNil(o.RankInCategory) {
		return nil, false
	}
	return o.RankInCategory, true
}

// HasRankInCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) HasRankInCategory() bool {
	if o != nil && !IsNil(o.RankInCategory) {
		return true
	}

	return false
}

// SetRankInCategory gets a reference to the given int64 and assigns it to the RankInCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) SetRankInCategory(v int64) {
	o.RankInCategory = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.ShareClassReturn) {
		toSerialize["share_class_return"] = o.ShareClassReturn
	}
	if !IsNil(o.CategoryReturn) {
		toSerialize["category_return"] = o.CategoryReturn
	}
	if !IsNil(o.RankInCategory) {
		toSerialize["rank_in_category"] = o.RankInCategory
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner struct {
	value *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) Get() *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) Set(val *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner(val *GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner {
	return &NullableGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
