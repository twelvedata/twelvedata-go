/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums{}

// GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums Minimum investment amounts required to purchase or add to the mutual fund, including IRA minimums
type GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums struct {
	// Investment minimum
	InitialInvestment *int64 `json:"initial_investment,omitempty"`
	// Minimum amount of additional investment
	AdditionalInvestment *int64 `json:"additional_investment,omitempty"`
	// Investment minimum for IRA
	InitialIraInvestment *int64 `json:"initial_ira_investment,omitempty"`
	// Minimum amount of additional investment for IRA
	AdditionalIraInvestment *int64 `json:"additional_ira_investment,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums instantiates a new GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums() *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums {
	this := GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimumsWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimumsWithDefaults() *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums {
	this := GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums{}
	return &this
}

// GetInitialInvestment returns the InitialInvestment field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) GetInitialInvestment() int64 {
	if o == nil || IsNil(o.InitialInvestment) {
		var ret int64
		return ret
	}
	return *o.InitialInvestment
}

// GetInitialInvestmentOk returns a tuple with the InitialInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) GetInitialInvestmentOk() (*int64, bool) {
	if o == nil || IsNil(o.InitialInvestment) {
		return nil, false
	}
	return o.InitialInvestment, true
}

// HasInitialInvestment returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) HasInitialInvestment() bool {
	if o != nil && !IsNil(o.InitialInvestment) {
		return true
	}

	return false
}

// SetInitialInvestment gets a reference to the given int64 and assigns it to the InitialInvestment field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) SetInitialInvestment(v int64) {
	o.InitialInvestment = &v
}

// GetAdditionalInvestment returns the AdditionalInvestment field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) GetAdditionalInvestment() int64 {
	if o == nil || IsNil(o.AdditionalInvestment) {
		var ret int64
		return ret
	}
	return *o.AdditionalInvestment
}

// GetAdditionalInvestmentOk returns a tuple with the AdditionalInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) GetAdditionalInvestmentOk() (*int64, bool) {
	if o == nil || IsNil(o.AdditionalInvestment) {
		return nil, false
	}
	return o.AdditionalInvestment, true
}

// HasAdditionalInvestment returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) HasAdditionalInvestment() bool {
	if o != nil && !IsNil(o.AdditionalInvestment) {
		return true
	}

	return false
}

// SetAdditionalInvestment gets a reference to the given int64 and assigns it to the AdditionalInvestment field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) SetAdditionalInvestment(v int64) {
	o.AdditionalInvestment = &v
}

// GetInitialIraInvestment returns the InitialIraInvestment field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) GetInitialIraInvestment() int64 {
	if o == nil || IsNil(o.InitialIraInvestment) {
		var ret int64
		return ret
	}
	return *o.InitialIraInvestment
}

// GetInitialIraInvestmentOk returns a tuple with the InitialIraInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) GetInitialIraInvestmentOk() (*int64, bool) {
	if o == nil || IsNil(o.InitialIraInvestment) {
		return nil, false
	}
	return o.InitialIraInvestment, true
}

// HasInitialIraInvestment returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) HasInitialIraInvestment() bool {
	if o != nil && !IsNil(o.InitialIraInvestment) {
		return true
	}

	return false
}

// SetInitialIraInvestment gets a reference to the given int64 and assigns it to the InitialIraInvestment field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) SetInitialIraInvestment(v int64) {
	o.InitialIraInvestment = &v
}

// GetAdditionalIraInvestment returns the AdditionalIraInvestment field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) GetAdditionalIraInvestment() int64 {
	if o == nil || IsNil(o.AdditionalIraInvestment) {
		var ret int64
		return ret
	}
	return *o.AdditionalIraInvestment
}

// GetAdditionalIraInvestmentOk returns a tuple with the AdditionalIraInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) GetAdditionalIraInvestmentOk() (*int64, bool) {
	if o == nil || IsNil(o.AdditionalIraInvestment) {
		return nil, false
	}
	return o.AdditionalIraInvestment, true
}

// HasAdditionalIraInvestment returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) HasAdditionalIraInvestment() bool {
	if o != nil && !IsNil(o.AdditionalIraInvestment) {
		return true
	}

	return false
}

// SetAdditionalIraInvestment gets a reference to the given int64 and assigns it to the AdditionalIraInvestment field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) SetAdditionalIraInvestment(v int64) {
	o.AdditionalIraInvestment = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InitialInvestment) {
		toSerialize["initial_investment"] = o.InitialInvestment
	}
	if !IsNil(o.AdditionalInvestment) {
		toSerialize["additional_investment"] = o.AdditionalInvestment
	}
	if !IsNil(o.InitialIraInvestment) {
		toSerialize["initial_ira_investment"] = o.InitialIraInvestment
	}
	if !IsNil(o.AdditionalIraInvestment) {
		toSerialize["additional_ira_investment"] = o.AdditionalIraInvestment
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums struct {
	value *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) Get() *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) Set(val *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums(val *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums {
	return &NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
