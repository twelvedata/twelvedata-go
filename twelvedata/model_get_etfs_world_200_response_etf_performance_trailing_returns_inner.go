/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner{}

// GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner struct for GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner
type GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner struct {
	// Period of trailing returns
	Period *string `json:"period,omitempty"`
	// Fund returns (%) generated over a given period
	ShareClassReturn *float64 `json:"share_class_return,omitempty"`
	// Same category average returns (%) generated over a given period
	CategoryReturn *float64 `json:"category_return,omitempty"`
}

// NewGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner instantiates a new GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner() *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner {
	this := GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner{}
	return &this
}

// NewGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInnerWithDefaults instantiates a new GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInnerWithDefaults() *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner {
	this := GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) SetPeriod(v string) {
	o.Period = &v
}

// GetShareClassReturn returns the ShareClassReturn field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) GetShareClassReturn() float64 {
	if o == nil || IsNil(o.ShareClassReturn) {
		var ret float64
		return ret
	}
	return *o.ShareClassReturn
}

// GetShareClassReturnOk returns a tuple with the ShareClassReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) GetShareClassReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.ShareClassReturn) {
		return nil, false
	}
	return o.ShareClassReturn, true
}

// HasShareClassReturn returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) HasShareClassReturn() bool {
	if o != nil && !IsNil(o.ShareClassReturn) {
		return true
	}

	return false
}

// SetShareClassReturn gets a reference to the given float64 and assigns it to the ShareClassReturn field.
func (o *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) SetShareClassReturn(v float64) {
	o.ShareClassReturn = &v
}

// GetCategoryReturn returns the CategoryReturn field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) GetCategoryReturn() float64 {
	if o == nil || IsNil(o.CategoryReturn) {
		var ret float64
		return ret
	}
	return *o.CategoryReturn
}

// GetCategoryReturnOk returns a tuple with the CategoryReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) GetCategoryReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.CategoryReturn) {
		return nil, false
	}
	return o.CategoryReturn, true
}

// HasCategoryReturn returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) HasCategoryReturn() bool {
	if o != nil && !IsNil(o.CategoryReturn) {
		return true
	}

	return false
}

// SetCategoryReturn gets a reference to the given float64 and assigns it to the CategoryReturn field.
func (o *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) SetCategoryReturn(v float64) {
	o.CategoryReturn = &v
}

func (o GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner struct {
	value *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) Get() *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) Set(val *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner(val *GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) *NullableGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner {
	return &NullableGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
