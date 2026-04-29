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

// checks if the GetBalanceSheetConsolidated200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheetConsolidated200Response{}

// GetBalanceSheetConsolidated200Response struct for GetBalanceSheetConsolidated200Response
type GetBalanceSheetConsolidated200Response struct {
	// Balance sheet data
	BalanceSheet []BalanceSheetConsolidatedItem `json:"balance_sheet"`
	// Response status
	Status *string `json:"status,omitempty"`
}

type _GetBalanceSheetConsolidated200Response GetBalanceSheetConsolidated200Response

// NewGetBalanceSheetConsolidated200Response instantiates a new GetBalanceSheetConsolidated200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheetConsolidated200Response(balanceSheet []BalanceSheetConsolidatedItem) *GetBalanceSheetConsolidated200Response {
	this := GetBalanceSheetConsolidated200Response{}
	this.BalanceSheet = balanceSheet
	return &this
}

// NewGetBalanceSheetConsolidated200ResponseWithDefaults instantiates a new GetBalanceSheetConsolidated200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheetConsolidated200ResponseWithDefaults() *GetBalanceSheetConsolidated200Response {
	this := GetBalanceSheetConsolidated200Response{}
	return &this
}

// GetBalanceSheet returns the BalanceSheet field value
func (o *GetBalanceSheetConsolidated200Response) GetBalanceSheet() []BalanceSheetConsolidatedItem {
	if o == nil {
		var ret []BalanceSheetConsolidatedItem
		return ret
	}

	return o.BalanceSheet
}

// GetBalanceSheetOk returns a tuple with the BalanceSheet field value
// and a boolean to check if the value has been set.
func (o *GetBalanceSheetConsolidated200Response) GetBalanceSheetOk() ([]BalanceSheetConsolidatedItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.BalanceSheet, true
}

// SetBalanceSheet sets field value
func (o *GetBalanceSheetConsolidated200Response) SetBalanceSheet(v []BalanceSheetConsolidatedItem) {
	o.BalanceSheet = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetBalanceSheetConsolidated200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheetConsolidated200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetBalanceSheetConsolidated200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetBalanceSheetConsolidated200Response) SetStatus(v string) {
	o.Status = &v
}

func (o GetBalanceSheetConsolidated200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheetConsolidated200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balance_sheet"] = o.BalanceSheet
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *GetBalanceSheetConsolidated200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"balance_sheet",
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

	varGetBalanceSheetConsolidated200Response := _GetBalanceSheetConsolidated200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetBalanceSheetConsolidated200Response)

	if err != nil {
		return err
	}

	*o = GetBalanceSheetConsolidated200Response(varGetBalanceSheetConsolidated200Response)

	return err
}

type NullableGetBalanceSheetConsolidated200Response struct {
	value *GetBalanceSheetConsolidated200Response
	isSet bool
}

func (v NullableGetBalanceSheetConsolidated200Response) Get() *GetBalanceSheetConsolidated200Response {
	return v.value
}

func (v *NullableGetBalanceSheetConsolidated200Response) Set(val *GetBalanceSheetConsolidated200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheetConsolidated200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheetConsolidated200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheetConsolidated200Response(val *GetBalanceSheetConsolidated200Response) *NullableGetBalanceSheetConsolidated200Response {
	return &NullableGetBalanceSheetConsolidated200Response{value: val, isSet: true}
}

func (v NullableGetBalanceSheetConsolidated200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheetConsolidated200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
