// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars{}

// GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars Corporate ESG pillars
type GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars struct {
	// ESG environmental score
	Environmental *float64 `json:"environmental,omitempty"`
	// ESG social score
	Social *float64 `json:"social,omitempty"`
	// ESG governance score
	Governance *float64 `json:"governance,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars instantiates a new GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars() *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars {
	this := GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillarsWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillarsWithDefaults() *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars {
	this := GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars{}
	return &this
}

// GetEnvironmental returns the Environmental field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) GetEnvironmental() float64 {
	if o == nil || IsNil(o.Environmental) {
		var ret float64
		return ret
	}
	return *o.Environmental
}

// GetEnvironmentalOk returns a tuple with the Environmental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) GetEnvironmentalOk() (*float64, bool) {
	if o == nil || IsNil(o.Environmental) {
		return nil, false
	}
	return o.Environmental, true
}

// HasEnvironmental returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) HasEnvironmental() bool {
	if o != nil && !IsNil(o.Environmental) {
		return true
	}

	return false
}

// SetEnvironmental gets a reference to the given float64 and assigns it to the Environmental field.
func (o *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) SetEnvironmental(v float64) {
	o.Environmental = &v
}

// GetSocial returns the Social field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) GetSocial() float64 {
	if o == nil || IsNil(o.Social) {
		var ret float64
		return ret
	}
	return *o.Social
}

// GetSocialOk returns a tuple with the Social field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) GetSocialOk() (*float64, bool) {
	if o == nil || IsNil(o.Social) {
		return nil, false
	}
	return o.Social, true
}

// HasSocial returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) HasSocial() bool {
	if o != nil && !IsNil(o.Social) {
		return true
	}

	return false
}

// SetSocial gets a reference to the given float64 and assigns it to the Social field.
func (o *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) SetSocial(v float64) {
	o.Social = &v
}

// GetGovernance returns the Governance field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) GetGovernance() float64 {
	if o == nil || IsNil(o.Governance) {
		var ret float64
		return ret
	}
	return *o.Governance
}

// GetGovernanceOk returns a tuple with the Governance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) GetGovernanceOk() (*float64, bool) {
	if o == nil || IsNil(o.Governance) {
		return nil, false
	}
	return o.Governance, true
}

// HasGovernance returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) HasGovernance() bool {
	if o != nil && !IsNil(o.Governance) {
		return true
	}

	return false
}

// SetGovernance gets a reference to the given float64 and assigns it to the Governance field.
func (o *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) SetGovernance(v float64) {
	o.Governance = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Environmental) {
		toSerialize["environmental"] = o.Environmental
	}
	if !IsNil(o.Social) {
		toSerialize["social"] = o.Social
	}
	if !IsNil(o.Governance) {
		toSerialize["governance"] = o.Governance
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars struct {
	value *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) Get() *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) Set(val *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars(val *GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) *NullableGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars {
	return &NullableGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
