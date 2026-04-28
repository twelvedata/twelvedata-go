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

// checks if the ExchangesResponseItemAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangesResponseItemAccess{}

// ExchangesResponseItemAccess Info on which plan symbol is available (displayed then `show_plan` is `true`)
type ExchangesResponseItemAccess struct {
	// Level of access to the symbol
	Global string `json:"global"`
	// The individual plan name for the symbol
	Plan string `json:"plan"`
	// The business plan name for the symbol
	PlanBusiness string `json:"plan_business"`
}

type _ExchangesResponseItemAccess ExchangesResponseItemAccess

// NewExchangesResponseItemAccess instantiates a new ExchangesResponseItemAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangesResponseItemAccess(global string, plan string, planBusiness string) *ExchangesResponseItemAccess {
	this := ExchangesResponseItemAccess{}
	this.Global = global
	this.Plan = plan
	this.PlanBusiness = planBusiness
	return &this
}

// NewExchangesResponseItemAccessWithDefaults instantiates a new ExchangesResponseItemAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangesResponseItemAccessWithDefaults() *ExchangesResponseItemAccess {
	this := ExchangesResponseItemAccess{}
	return &this
}

// GetGlobal returns the Global field value
func (o *ExchangesResponseItemAccess) GetGlobal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Global
}

// GetGlobalOk returns a tuple with the Global field value
// and a boolean to check if the value has been set.
func (o *ExchangesResponseItemAccess) GetGlobalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Global, true
}

// SetGlobal sets field value
func (o *ExchangesResponseItemAccess) SetGlobal(v string) {
	o.Global = v
}

// GetPlan returns the Plan field value
func (o *ExchangesResponseItemAccess) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *ExchangesResponseItemAccess) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *ExchangesResponseItemAccess) SetPlan(v string) {
	o.Plan = v
}

// GetPlanBusiness returns the PlanBusiness field value
func (o *ExchangesResponseItemAccess) GetPlanBusiness() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanBusiness
}

// GetPlanBusinessOk returns a tuple with the PlanBusiness field value
// and a boolean to check if the value has been set.
func (o *ExchangesResponseItemAccess) GetPlanBusinessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanBusiness, true
}

// SetPlanBusiness sets field value
func (o *ExchangesResponseItemAccess) SetPlanBusiness(v string) {
	o.PlanBusiness = v
}

func (o ExchangesResponseItemAccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangesResponseItemAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["global"] = o.Global
	toSerialize["plan"] = o.Plan
	toSerialize["plan_business"] = o.PlanBusiness
	return toSerialize, nil
}

func (o *ExchangesResponseItemAccess) UnmarshalJSON(data []byte) (err error) {
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

	varExchangesResponseItemAccess := _ExchangesResponseItemAccess{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExchangesResponseItemAccess)

	if err != nil {
		return err
	}

	*o = ExchangesResponseItemAccess(varExchangesResponseItemAccess)

	return err
}

type NullableExchangesResponseItemAccess struct {
	value *ExchangesResponseItemAccess
	isSet bool
}

func (v NullableExchangesResponseItemAccess) Get() *ExchangesResponseItemAccess {
	return v.value
}

func (v *NullableExchangesResponseItemAccess) Set(val *ExchangesResponseItemAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangesResponseItemAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangesResponseItemAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangesResponseItemAccess(val *ExchangesResponseItemAccess) *NullableExchangesResponseItemAccess {
	return &NullableExchangesResponseItemAccess{value: val, isSet: true}
}

func (v NullableExchangesResponseItemAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangesResponseItemAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
