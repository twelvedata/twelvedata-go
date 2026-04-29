/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BondsResponseItemAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BondsResponseItemAccess{}

// BondsResponseItemAccess Info on which plan symbol is available (displayed then `show_plan` is `true`)
type BondsResponseItemAccess struct {
	// Level of access to the symbol
	Global string `json:"global"`
	// The individual plan name for the symbol
	Plan string `json:"plan"`
	// The business plan name for the symbol
	PlanBusiness string `json:"plan_business"`
}

type _BondsResponseItemAccess BondsResponseItemAccess

// NewBondsResponseItemAccess instantiates a new BondsResponseItemAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBondsResponseItemAccess(global string, plan string, planBusiness string) *BondsResponseItemAccess {
	this := BondsResponseItemAccess{}
	this.Global = global
	this.Plan = plan
	this.PlanBusiness = planBusiness
	return &this
}

// NewBondsResponseItemAccessWithDefaults instantiates a new BondsResponseItemAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBondsResponseItemAccessWithDefaults() *BondsResponseItemAccess {
	this := BondsResponseItemAccess{}
	return &this
}

// GetGlobal returns the Global field value
func (o *BondsResponseItemAccess) GetGlobal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Global
}

// GetGlobalOk returns a tuple with the Global field value
// and a boolean to check if the value has been set.
func (o *BondsResponseItemAccess) GetGlobalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Global, true
}

// SetGlobal sets field value
func (o *BondsResponseItemAccess) SetGlobal(v string) {
	o.Global = v
}

// GetPlan returns the Plan field value
func (o *BondsResponseItemAccess) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *BondsResponseItemAccess) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *BondsResponseItemAccess) SetPlan(v string) {
	o.Plan = v
}

// GetPlanBusiness returns the PlanBusiness field value
func (o *BondsResponseItemAccess) GetPlanBusiness() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanBusiness
}

// GetPlanBusinessOk returns a tuple with the PlanBusiness field value
// and a boolean to check if the value has been set.
func (o *BondsResponseItemAccess) GetPlanBusinessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanBusiness, true
}

// SetPlanBusiness sets field value
func (o *BondsResponseItemAccess) SetPlanBusiness(v string) {
	o.PlanBusiness = v
}

func (o BondsResponseItemAccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BondsResponseItemAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["global"] = o.Global
	toSerialize["plan"] = o.Plan
	toSerialize["plan_business"] = o.PlanBusiness
	return toSerialize, nil
}

func (o *BondsResponseItemAccess) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"global",
		"plan",
		"plan_business",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBondsResponseItemAccess := _BondsResponseItemAccess{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varBondsResponseItemAccess)

	if err != nil {
		return err
	}

	*o = BondsResponseItemAccess(varBondsResponseItemAccess)

	return err
}

type NullableBondsResponseItemAccess struct {
	value *BondsResponseItemAccess
	isSet bool
}

func (v NullableBondsResponseItemAccess) Get() *BondsResponseItemAccess {
	return v.value
}

func (v *NullableBondsResponseItemAccess) Set(val *BondsResponseItemAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableBondsResponseItemAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableBondsResponseItemAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBondsResponseItemAccess(val *BondsResponseItemAccess) *NullableBondsResponseItemAccess {
	return &NullableBondsResponseItemAccess{value: val, isSet: true}
}

func (v NullableBondsResponseItemAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBondsResponseItemAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
