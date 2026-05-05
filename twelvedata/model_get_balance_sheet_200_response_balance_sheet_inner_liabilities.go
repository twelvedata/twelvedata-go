// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetBalanceSheet200ResponseBalanceSheetInnerLiabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheet200ResponseBalanceSheetInnerLiabilities{}

// GetBalanceSheet200ResponseBalanceSheetInnerLiabilities Liabilities section of the balance sheet
type GetBalanceSheet200ResponseBalanceSheetInnerLiabilities struct {
	CurrentLiabilities    *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities    `json:"current_liabilities,omitempty"`
	NonCurrentLiabilities *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities `json:"non_current_liabilities,omitempty"`
	// The sum of total_current_liabilities + total_non_current_liabilities
	TotalLiabilities *float64 `json:"total_liabilities,omitempty"`
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilities instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerLiabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilities() *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities {
	this := GetBalanceSheet200ResponseBalanceSheetInnerLiabilities{}
	return &this
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerLiabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities {
	this := GetBalanceSheet200ResponseBalanceSheetInnerLiabilities{}
	return &this
}

// GetCurrentLiabilities returns the CurrentLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) GetCurrentLiabilities() GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities {
	if o == nil || IsNil(o.CurrentLiabilities) {
		var ret GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities
		return ret
	}
	return *o.CurrentLiabilities
}

// GetCurrentLiabilitiesOk returns a tuple with the CurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) GetCurrentLiabilitiesOk() (*GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities, bool) {
	if o == nil || IsNil(o.CurrentLiabilities) {
		return nil, false
	}
	return o.CurrentLiabilities, true
}

// HasCurrentLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) HasCurrentLiabilities() bool {
	if o != nil && !IsNil(o.CurrentLiabilities) {
		return true
	}

	return false
}

// SetCurrentLiabilities gets a reference to the given GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities and assigns it to the CurrentLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) SetCurrentLiabilities(v GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) {
	o.CurrentLiabilities = &v
}

// GetNonCurrentLiabilities returns the NonCurrentLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) GetNonCurrentLiabilities() GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities {
	if o == nil || IsNil(o.NonCurrentLiabilities) {
		var ret GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities
		return ret
	}
	return *o.NonCurrentLiabilities
}

// GetNonCurrentLiabilitiesOk returns a tuple with the NonCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) GetNonCurrentLiabilitiesOk() (*GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities, bool) {
	if o == nil || IsNil(o.NonCurrentLiabilities) {
		return nil, false
	}
	return o.NonCurrentLiabilities, true
}

// HasNonCurrentLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) HasNonCurrentLiabilities() bool {
	if o != nil && !IsNil(o.NonCurrentLiabilities) {
		return true
	}

	return false
}

// SetNonCurrentLiabilities gets a reference to the given GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities and assigns it to the NonCurrentLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) SetNonCurrentLiabilities(v GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) {
	o.NonCurrentLiabilities = &v
}

// GetTotalLiabilities returns the TotalLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) GetTotalLiabilities() float64 {
	if o == nil || IsNil(o.TotalLiabilities) {
		var ret float64
		return ret
	}
	return *o.TotalLiabilities
}

// GetTotalLiabilitiesOk returns a tuple with the TotalLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) GetTotalLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalLiabilities) {
		return nil, false
	}
	return o.TotalLiabilities, true
}

// HasTotalLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) HasTotalLiabilities() bool {
	if o != nil && !IsNil(o.TotalLiabilities) {
		return true
	}

	return false
}

// SetTotalLiabilities gets a reference to the given float64 and assigns it to the TotalLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) SetTotalLiabilities(v float64) {
	o.TotalLiabilities = &v
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentLiabilities) {
		toSerialize["current_liabilities"] = o.CurrentLiabilities
	}
	if !IsNil(o.NonCurrentLiabilities) {
		toSerialize["non_current_liabilities"] = o.NonCurrentLiabilities
	}
	if !IsNil(o.TotalLiabilities) {
		toSerialize["total_liabilities"] = o.TotalLiabilities
	}
	return toSerialize, nil
}

type NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilities struct {
	value *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities
	isSet bool
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilities) Get() *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities {
	return v.value
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilities) Set(val *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilities(val *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilities {
	return &NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilities{value: val, isSet: true}
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
