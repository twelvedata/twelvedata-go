/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner{}

// GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner struct for GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner
type GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner struct {
	// Manager name
	Name *string `json:"name,omitempty"`
	// Manager tenuring date
	TenureSince *string `json:"tenure_since,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner instantiates a new GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner() *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner {
	this := GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInnerWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInnerWithDefaults() *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner {
	this := GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) SetName(v string) {
	o.Name = &v
}

// GetTenureSince returns the TenureSince field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) GetTenureSince() string {
	if o == nil || IsNil(o.TenureSince) {
		var ret string
		return ret
	}
	return *o.TenureSince
}

// GetTenureSinceOk returns a tuple with the TenureSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) GetTenureSinceOk() (*string, bool) {
	if o == nil || IsNil(o.TenureSince) {
		return nil, false
	}
	return o.TenureSince, true
}

// HasTenureSince returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) HasTenureSince() bool {
	if o != nil && !IsNil(o.TenureSince) {
		return true
	}

	return false
}

// SetTenureSince gets a reference to the given string and assigns it to the TenureSince field.
func (o *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) SetTenureSince(v string) {
	o.TenureSince = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TenureSince) {
		toSerialize["tenure_since"] = o.TenureSince
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner struct {
	value *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) Get() *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) Set(val *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner(val *GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) *NullableGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner {
	return &NullableGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
