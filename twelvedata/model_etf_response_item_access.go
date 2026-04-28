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

// checks if the EtfResponseItemAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EtfResponseItemAccess{}

// EtfResponseItemAccess Info on which plan symbol is available (displayed then `show_plan` is `true`)
type EtfResponseItemAccess struct {
	// Level of access to the symbol
	Global string `json:"global"`
	// The individual plan name for the symbol
	Plan string `json:"plan"`
	// The business plan name for the symbol
	PlanBusiness string `json:"plan_business"`
}

type _EtfResponseItemAccess EtfResponseItemAccess

// NewEtfResponseItemAccess instantiates a new EtfResponseItemAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEtfResponseItemAccess(global string, plan string, planBusiness string) *EtfResponseItemAccess {
	this := EtfResponseItemAccess{}
	this.Global = global
	this.Plan = plan
	this.PlanBusiness = planBusiness
	return &this
}

// NewEtfResponseItemAccessWithDefaults instantiates a new EtfResponseItemAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEtfResponseItemAccessWithDefaults() *EtfResponseItemAccess {
	this := EtfResponseItemAccess{}
	return &this
}

// GetGlobal returns the Global field value
func (o *EtfResponseItemAccess) GetGlobal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Global
}

// GetGlobalOk returns a tuple with the Global field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItemAccess) GetGlobalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Global, true
}

// SetGlobal sets field value
func (o *EtfResponseItemAccess) SetGlobal(v string) {
	o.Global = v
}

// GetPlan returns the Plan field value
func (o *EtfResponseItemAccess) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItemAccess) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *EtfResponseItemAccess) SetPlan(v string) {
	o.Plan = v
}

// GetPlanBusiness returns the PlanBusiness field value
func (o *EtfResponseItemAccess) GetPlanBusiness() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanBusiness
}

// GetPlanBusinessOk returns a tuple with the PlanBusiness field value
// and a boolean to check if the value has been set.
func (o *EtfResponseItemAccess) GetPlanBusinessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanBusiness, true
}

// SetPlanBusiness sets field value
func (o *EtfResponseItemAccess) SetPlanBusiness(v string) {
	o.PlanBusiness = v
}

func (o EtfResponseItemAccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EtfResponseItemAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["global"] = o.Global
	toSerialize["plan"] = o.Plan
	toSerialize["plan_business"] = o.PlanBusiness
	return toSerialize, nil
}

func (o *EtfResponseItemAccess) UnmarshalJSON(data []byte) (err error) {
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

	varEtfResponseItemAccess := _EtfResponseItemAccess{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEtfResponseItemAccess)

	if err != nil {
		return err
	}

	*o = EtfResponseItemAccess(varEtfResponseItemAccess)

	return err
}

type NullableEtfResponseItemAccess struct {
	value *EtfResponseItemAccess
	isSet bool
}

func (v NullableEtfResponseItemAccess) Get() *EtfResponseItemAccess {
	return v.value
}

func (v *NullableEtfResponseItemAccess) Set(val *EtfResponseItemAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableEtfResponseItemAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableEtfResponseItemAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtfResponseItemAccess(val *EtfResponseItemAccess) *NullableEtfResponseItemAccess {
	return &NullableEtfResponseItemAccess{value: val, isSet: true}
}

func (v NullableEtfResponseItemAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtfResponseItemAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
