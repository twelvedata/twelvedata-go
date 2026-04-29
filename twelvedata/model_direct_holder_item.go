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

// checks if the DirectHolderItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectHolderItem{}

// DirectHolderItem struct for DirectHolderItem
type DirectHolderItem struct {
	// Refers to the legal name of the institution
	EntityName string `json:"entity_name"`
	// Refers to date reported
	DateReported *string `json:"date_reported,omitempty"`
	// Refers to the number of shares owned
	Shares *int64 `json:"shares,omitempty"`
	// Total value of shares owned, calculated by multiplying `shares` by the current price
	Value *int64 `json:"value,omitempty"`
	// Represents the percentage of shares outstanding that are owned by the financial institution
	PercentHeld *float64 `json:"percent_held,omitempty"`
}

type _DirectHolderItem DirectHolderItem

// NewDirectHolderItem instantiates a new DirectHolderItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectHolderItem(entityName string) *DirectHolderItem {
	this := DirectHolderItem{}
	this.EntityName = entityName
	return &this
}

// NewDirectHolderItemWithDefaults instantiates a new DirectHolderItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectHolderItemWithDefaults() *DirectHolderItem {
	this := DirectHolderItem{}
	return &this
}

// GetEntityName returns the EntityName field value
func (o *DirectHolderItem) GetEntityName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value
// and a boolean to check if the value has been set.
func (o *DirectHolderItem) GetEntityNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityName, true
}

// SetEntityName sets field value
func (o *DirectHolderItem) SetEntityName(v string) {
	o.EntityName = v
}

// GetDateReported returns the DateReported field value if set, zero value otherwise.
func (o *DirectHolderItem) GetDateReported() string {
	if o == nil || IsNil(o.DateReported) {
		var ret string
		return ret
	}
	return *o.DateReported
}

// GetDateReportedOk returns a tuple with the DateReported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectHolderItem) GetDateReportedOk() (*string, bool) {
	if o == nil || IsNil(o.DateReported) {
		return nil, false
	}
	return o.DateReported, true
}

// HasDateReported returns a boolean if a field has been set.
func (o *DirectHolderItem) HasDateReported() bool {
	if o != nil && !IsNil(o.DateReported) {
		return true
	}

	return false
}

// SetDateReported gets a reference to the given string and assigns it to the DateReported field.
func (o *DirectHolderItem) SetDateReported(v string) {
	o.DateReported = &v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *DirectHolderItem) GetShares() int64 {
	if o == nil || IsNil(o.Shares) {
		var ret int64
		return ret
	}
	return *o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectHolderItem) GetSharesOk() (*int64, bool) {
	if o == nil || IsNil(o.Shares) {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *DirectHolderItem) HasShares() bool {
	if o != nil && !IsNil(o.Shares) {
		return true
	}

	return false
}

// SetShares gets a reference to the given int64 and assigns it to the Shares field.
func (o *DirectHolderItem) SetShares(v int64) {
	o.Shares = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DirectHolderItem) GetValue() int64 {
	if o == nil || IsNil(o.Value) {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectHolderItem) GetValueOk() (*int64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DirectHolderItem) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *DirectHolderItem) SetValue(v int64) {
	o.Value = &v
}

// GetPercentHeld returns the PercentHeld field value if set, zero value otherwise.
func (o *DirectHolderItem) GetPercentHeld() float64 {
	if o == nil || IsNil(o.PercentHeld) {
		var ret float64
		return ret
	}
	return *o.PercentHeld
}

// GetPercentHeldOk returns a tuple with the PercentHeld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectHolderItem) GetPercentHeldOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentHeld) {
		return nil, false
	}
	return o.PercentHeld, true
}

// HasPercentHeld returns a boolean if a field has been set.
func (o *DirectHolderItem) HasPercentHeld() bool {
	if o != nil && !IsNil(o.PercentHeld) {
		return true
	}

	return false
}

// SetPercentHeld gets a reference to the given float64 and assigns it to the PercentHeld field.
func (o *DirectHolderItem) SetPercentHeld(v float64) {
	o.PercentHeld = &v
}

func (o DirectHolderItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectHolderItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_name"] = o.EntityName
	if !IsNil(o.DateReported) {
		toSerialize["date_reported"] = o.DateReported
	}
	if !IsNil(o.Shares) {
		toSerialize["shares"] = o.Shares
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.PercentHeld) {
		toSerialize["percent_held"] = o.PercentHeld
	}
	return toSerialize, nil
}

func (o *DirectHolderItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_name",
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

	varDirectHolderItem := _DirectHolderItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDirectHolderItem)

	if err != nil {
		return err
	}

	*o = DirectHolderItem(varDirectHolderItem)

	return err
}

type NullableDirectHolderItem struct {
	value *DirectHolderItem
	isSet bool
}

func (v NullableDirectHolderItem) Get() *DirectHolderItem {
	return v.value
}

func (v *NullableDirectHolderItem) Set(val *DirectHolderItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectHolderItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectHolderItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectHolderItem(val *DirectHolderItem) *NullableDirectHolderItem {
	return &NullableDirectHolderItem{value: val, isSet: true}
}

func (v NullableDirectHolderItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectHolderItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
