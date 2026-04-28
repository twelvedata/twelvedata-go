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

// checks if the GetDividends200ResponseDividendsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDividends200ResponseDividendsInner{}

// GetDividends200ResponseDividendsInner struct for GetDividends200ResponseDividendsInner
type GetDividends200ResponseDividendsInner struct {
	// Stands for the ex date
	ExDate string `json:"ex_date"`
	// Stands for the payment amount
	Amount float64 `json:"amount"`
}

type _GetDividends200ResponseDividendsInner GetDividends200ResponseDividendsInner

// NewGetDividends200ResponseDividendsInner instantiates a new GetDividends200ResponseDividendsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDividends200ResponseDividendsInner(exDate string, amount float64) *GetDividends200ResponseDividendsInner {
	this := GetDividends200ResponseDividendsInner{}
	this.ExDate = exDate
	this.Amount = amount
	return &this
}

// NewGetDividends200ResponseDividendsInnerWithDefaults instantiates a new GetDividends200ResponseDividendsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDividends200ResponseDividendsInnerWithDefaults() *GetDividends200ResponseDividendsInner {
	this := GetDividends200ResponseDividendsInner{}
	return &this
}

// GetExDate returns the ExDate field value
func (o *GetDividends200ResponseDividendsInner) GetExDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExDate
}

// GetExDateOk returns a tuple with the ExDate field value
// and a boolean to check if the value has been set.
func (o *GetDividends200ResponseDividendsInner) GetExDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExDate, true
}

// SetExDate sets field value
func (o *GetDividends200ResponseDividendsInner) SetExDate(v string) {
	o.ExDate = v
}

// GetAmount returns the Amount field value
func (o *GetDividends200ResponseDividendsInner) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetDividends200ResponseDividendsInner) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetDividends200ResponseDividendsInner) SetAmount(v float64) {
	o.Amount = v
}

func (o GetDividends200ResponseDividendsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDividends200ResponseDividendsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ex_date"] = o.ExDate
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *GetDividends200ResponseDividendsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ex_date",
		"amount",
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

	varGetDividends200ResponseDividendsInner := _GetDividends200ResponseDividendsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetDividends200ResponseDividendsInner)

	if err != nil {
		return err
	}

	*o = GetDividends200ResponseDividendsInner(varGetDividends200ResponseDividendsInner)

	return err
}

type NullableGetDividends200ResponseDividendsInner struct {
	value *GetDividends200ResponseDividendsInner
	isSet bool
}

func (v NullableGetDividends200ResponseDividendsInner) Get() *GetDividends200ResponseDividendsInner {
	return v.value
}

func (v *NullableGetDividends200ResponseDividendsInner) Set(val *GetDividends200ResponseDividendsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDividends200ResponseDividendsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDividends200ResponseDividendsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDividends200ResponseDividendsInner(val *GetDividends200ResponseDividendsInner) *NullableGetDividends200ResponseDividendsInner {
	return &NullableGetDividends200ResponseDividendsInner{value: val, isSet: true}
}

func (v NullableGetDividends200ResponseDividendsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDividends200ResponseDividendsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
