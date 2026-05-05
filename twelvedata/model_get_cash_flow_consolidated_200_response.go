// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetCashFlowConsolidated200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCashFlowConsolidated200Response{}

// GetCashFlowConsolidated200Response struct for GetCashFlowConsolidated200Response
type GetCashFlowConsolidated200Response struct {
	// Cash flow data
	CashFlow []CashFlowData `json:"cash_flow"`
	// Response status
	Status *string `json:"status,omitempty"`
}

type _GetCashFlowConsolidated200Response GetCashFlowConsolidated200Response

// NewGetCashFlowConsolidated200Response instantiates a new GetCashFlowConsolidated200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCashFlowConsolidated200Response(cashFlow []CashFlowData) *GetCashFlowConsolidated200Response {
	this := GetCashFlowConsolidated200Response{}
	this.CashFlow = cashFlow
	return &this
}

// NewGetCashFlowConsolidated200ResponseWithDefaults instantiates a new GetCashFlowConsolidated200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCashFlowConsolidated200ResponseWithDefaults() *GetCashFlowConsolidated200Response {
	this := GetCashFlowConsolidated200Response{}
	return &this
}

// GetCashFlow returns the CashFlow field value
func (o *GetCashFlowConsolidated200Response) GetCashFlow() []CashFlowData {
	if o == nil {
		var ret []CashFlowData
		return ret
	}

	return o.CashFlow
}

// GetCashFlowOk returns a tuple with the CashFlow field value
// and a boolean to check if the value has been set.
func (o *GetCashFlowConsolidated200Response) GetCashFlowOk() ([]CashFlowData, bool) {
	if o == nil {
		return nil, false
	}
	return o.CashFlow, true
}

// SetCashFlow sets field value
func (o *GetCashFlowConsolidated200Response) SetCashFlow(v []CashFlowData) {
	o.CashFlow = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetCashFlowConsolidated200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCashFlowConsolidated200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetCashFlowConsolidated200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetCashFlowConsolidated200Response) SetStatus(v string) {
	o.Status = &v
}

func (o GetCashFlowConsolidated200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCashFlowConsolidated200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cash_flow"] = o.CashFlow
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *GetCashFlowConsolidated200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cash_flow",
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

	varGetCashFlowConsolidated200Response := _GetCashFlowConsolidated200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetCashFlowConsolidated200Response)

	if err != nil {
		return err
	}

	*o = GetCashFlowConsolidated200Response(varGetCashFlowConsolidated200Response)

	return err
}

type NullableGetCashFlowConsolidated200Response struct {
	value *GetCashFlowConsolidated200Response
	isSet bool
}

func (v NullableGetCashFlowConsolidated200Response) Get() *GetCashFlowConsolidated200Response {
	return v.value
}

func (v *NullableGetCashFlowConsolidated200Response) Set(val *GetCashFlowConsolidated200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCashFlowConsolidated200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCashFlowConsolidated200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCashFlowConsolidated200Response(val *GetCashFlowConsolidated200Response) *NullableGetCashFlowConsolidated200Response {
	return &NullableGetCashFlowConsolidated200Response{value: val, isSet: true}
}

func (v NullableGetCashFlowConsolidated200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCashFlowConsolidated200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
