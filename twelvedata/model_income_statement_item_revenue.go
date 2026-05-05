// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemRevenue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemRevenue{}

// IncomeStatementItemRevenue Revenue information
type IncomeStatementItemRevenue struct {
	// Total revenue
	TotalRevenue *float64 `json:"total_revenue,omitempty"`
	// Operating revenue
	OperatingRevenue *float64 `json:"operating_revenue,omitempty"`
}

// NewIncomeStatementItemRevenue instantiates a new IncomeStatementItemRevenue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemRevenue() *IncomeStatementItemRevenue {
	this := IncomeStatementItemRevenue{}
	return &this
}

// NewIncomeStatementItemRevenueWithDefaults instantiates a new IncomeStatementItemRevenue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemRevenueWithDefaults() *IncomeStatementItemRevenue {
	this := IncomeStatementItemRevenue{}
	return &this
}

// GetTotalRevenue returns the TotalRevenue field value if set, zero value otherwise.
func (o *IncomeStatementItemRevenue) GetTotalRevenue() float64 {
	if o == nil || IsNil(o.TotalRevenue) {
		var ret float64
		return ret
	}
	return *o.TotalRevenue
}

// GetTotalRevenueOk returns a tuple with the TotalRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemRevenue) GetTotalRevenueOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalRevenue) {
		return nil, false
	}
	return o.TotalRevenue, true
}

// HasTotalRevenue returns a boolean if a field has been set.
func (o *IncomeStatementItemRevenue) HasTotalRevenue() bool {
	if o != nil && !IsNil(o.TotalRevenue) {
		return true
	}

	return false
}

// SetTotalRevenue gets a reference to the given float64 and assigns it to the TotalRevenue field.
func (o *IncomeStatementItemRevenue) SetTotalRevenue(v float64) {
	o.TotalRevenue = &v
}

// GetOperatingRevenue returns the OperatingRevenue field value if set, zero value otherwise.
func (o *IncomeStatementItemRevenue) GetOperatingRevenue() float64 {
	if o == nil || IsNil(o.OperatingRevenue) {
		var ret float64
		return ret
	}
	return *o.OperatingRevenue
}

// GetOperatingRevenueOk returns a tuple with the OperatingRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemRevenue) GetOperatingRevenueOk() (*float64, bool) {
	if o == nil || IsNil(o.OperatingRevenue) {
		return nil, false
	}
	return o.OperatingRevenue, true
}

// HasOperatingRevenue returns a boolean if a field has been set.
func (o *IncomeStatementItemRevenue) HasOperatingRevenue() bool {
	if o != nil && !IsNil(o.OperatingRevenue) {
		return true
	}

	return false
}

// SetOperatingRevenue gets a reference to the given float64 and assigns it to the OperatingRevenue field.
func (o *IncomeStatementItemRevenue) SetOperatingRevenue(v float64) {
	o.OperatingRevenue = &v
}

func (o IncomeStatementItemRevenue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemRevenue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalRevenue) {
		toSerialize["total_revenue"] = o.TotalRevenue
	}
	if !IsNil(o.OperatingRevenue) {
		toSerialize["operating_revenue"] = o.OperatingRevenue
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemRevenue struct {
	value *IncomeStatementItemRevenue
	isSet bool
}

func (v NullableIncomeStatementItemRevenue) Get() *IncomeStatementItemRevenue {
	return v.value
}

func (v *NullableIncomeStatementItemRevenue) Set(val *IncomeStatementItemRevenue) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemRevenue) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemRevenue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemRevenue(val *IncomeStatementItemRevenue) *NullableIncomeStatementItemRevenue {
	return &NullableIncomeStatementItemRevenue{value: val, isSet: true}
}

func (v NullableIncomeStatementItemRevenue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemRevenue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
