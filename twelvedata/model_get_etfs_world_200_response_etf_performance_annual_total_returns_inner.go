// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner{}

// GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner struct for GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner
type GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner struct {
	// Year of total returns
	Year *int64 `json:"year,omitempty"`
	// Fund total returns (%) generated over a given year
	ShareClassReturn *float64 `json:"share_class_return,omitempty"`
	// Same category average total returns (%) generated over a given year
	CategoryReturn *float64 `json:"category_return,omitempty"`
}

// NewGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner instantiates a new GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner() *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner {
	this := GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner{}
	return &this
}

// NewGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInnerWithDefaults instantiates a new GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInnerWithDefaults() *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner {
	this := GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner{}
	return &this
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) GetYear() int64 {
	if o == nil || IsNil(o.Year) {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) GetYearOk() (*int64, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) SetYear(v int64) {
	o.Year = &v
}

// GetShareClassReturn returns the ShareClassReturn field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) GetShareClassReturn() float64 {
	if o == nil || IsNil(o.ShareClassReturn) {
		var ret float64
		return ret
	}
	return *o.ShareClassReturn
}

// GetShareClassReturnOk returns a tuple with the ShareClassReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) GetShareClassReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.ShareClassReturn) {
		return nil, false
	}
	return o.ShareClassReturn, true
}

// HasShareClassReturn returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) HasShareClassReturn() bool {
	if o != nil && !IsNil(o.ShareClassReturn) {
		return true
	}

	return false
}

// SetShareClassReturn gets a reference to the given float64 and assigns it to the ShareClassReturn field.
func (o *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) SetShareClassReturn(v float64) {
	o.ShareClassReturn = &v
}

// GetCategoryReturn returns the CategoryReturn field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) GetCategoryReturn() float64 {
	if o == nil || IsNil(o.CategoryReturn) {
		var ret float64
		return ret
	}
	return *o.CategoryReturn
}

// GetCategoryReturnOk returns a tuple with the CategoryReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) GetCategoryReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.CategoryReturn) {
		return nil, false
	}
	return o.CategoryReturn, true
}

// HasCategoryReturn returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) HasCategoryReturn() bool {
	if o != nil && !IsNil(o.CategoryReturn) {
		return true
	}

	return false
}

// SetCategoryReturn gets a reference to the given float64 and assigns it to the CategoryReturn field.
func (o *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) SetCategoryReturn(v float64) {
	o.CategoryReturn = &v
}

func (o GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) ToMap() (map[string]interface{}, error) {
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

type NullableGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner struct {
	value *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) Get() *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) Set(val *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner(val *GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) *NullableGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner {
	return &NullableGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
