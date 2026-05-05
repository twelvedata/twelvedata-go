// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner{}

// GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner struct for GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner
type GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner struct {
	// Year of total returns
	Year *int64 `json:"year,omitempty"`
	// Fund total returns (%) generated over a given year
	ShareClassReturn *float64 `json:"share_class_return,omitempty"`
	// Same category average total returns (%) generated over a given year
	CategoryReturn *float64 `json:"category_return,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner instantiates a new GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner() *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner {
	this := GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInnerWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInnerWithDefaults() *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner {
	this := GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner{}
	return &this
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) GetYear() int64 {
	if o == nil || IsNil(o.Year) {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) GetYearOk() (*int64, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) SetYear(v int64) {
	o.Year = &v
}

// GetShareClassReturn returns the ShareClassReturn field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) GetShareClassReturn() float64 {
	if o == nil || IsNil(o.ShareClassReturn) {
		var ret float64
		return ret
	}
	return *o.ShareClassReturn
}

// GetShareClassReturnOk returns a tuple with the ShareClassReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) GetShareClassReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.ShareClassReturn) {
		return nil, false
	}
	return o.ShareClassReturn, true
}

// HasShareClassReturn returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) HasShareClassReturn() bool {
	if o != nil && !IsNil(o.ShareClassReturn) {
		return true
	}

	return false
}

// SetShareClassReturn gets a reference to the given float64 and assigns it to the ShareClassReturn field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) SetShareClassReturn(v float64) {
	o.ShareClassReturn = &v
}

// GetCategoryReturn returns the CategoryReturn field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) GetCategoryReturn() float64 {
	if o == nil || IsNil(o.CategoryReturn) {
		var ret float64
		return ret
	}
	return *o.CategoryReturn
}

// GetCategoryReturnOk returns a tuple with the CategoryReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) GetCategoryReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.CategoryReturn) {
		return nil, false
	}
	return o.CategoryReturn, true
}

// HasCategoryReturn returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) HasCategoryReturn() bool {
	if o != nil && !IsNil(o.CategoryReturn) {
		return true
	}

	return false
}

// SetCategoryReturn gets a reference to the given float64 and assigns it to the CategoryReturn field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) SetCategoryReturn(v float64) {
	o.CategoryReturn = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.ShareClassReturn) {
		toSerialize["share_class_return"] = o.ShareClassReturn
	}
	if !IsNil(o.CategoryReturn) {
		toSerialize["category_return"] = o.CategoryReturn
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner struct {
	value *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) Get() *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) Set(val *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner(val *GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner {
	return &NullableGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
