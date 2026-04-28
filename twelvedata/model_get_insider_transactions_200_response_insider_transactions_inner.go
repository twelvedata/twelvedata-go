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

// checks if the GetInsiderTransactions200ResponseInsiderTransactionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsiderTransactions200ResponseInsiderTransactionsInner{}

// GetInsiderTransactions200ResponseInsiderTransactionsInner struct for GetInsiderTransactions200ResponseInsiderTransactionsInner
type GetInsiderTransactions200ResponseInsiderTransactionsInner struct {
	// Full name of an individual, including first name, middle name, last name, and suffix
	FullName string `json:"full_name"`
	// Job position of insider
	Position string `json:"position"`
	// Date the transaction was reported
	DateReported string `json:"date_reported"`
	// `true` if direct, `false` if indirect
	IsDirect *bool `json:"is_direct,omitempty"`
	// As per report the number of shares acquired or disposed of the transaction
	Shares *int64 `json:"shares,omitempty"`
	// Represents the value of transaction, calculated as price multiplied by the volume
	Value *int64 `json:"value,omitempty"`
	// Exact price or price range of the transaction if available
	Description string `json:"description"`
}

type _GetInsiderTransactions200ResponseInsiderTransactionsInner GetInsiderTransactions200ResponseInsiderTransactionsInner

// NewGetInsiderTransactions200ResponseInsiderTransactionsInner instantiates a new GetInsiderTransactions200ResponseInsiderTransactionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsiderTransactions200ResponseInsiderTransactionsInner(fullName string, position string, dateReported string, description string) *GetInsiderTransactions200ResponseInsiderTransactionsInner {
	this := GetInsiderTransactions200ResponseInsiderTransactionsInner{}
	this.FullName = fullName
	this.Position = position
	this.DateReported = dateReported
	this.Description = description
	return &this
}

// NewGetInsiderTransactions200ResponseInsiderTransactionsInnerWithDefaults instantiates a new GetInsiderTransactions200ResponseInsiderTransactionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsiderTransactions200ResponseInsiderTransactionsInnerWithDefaults() *GetInsiderTransactions200ResponseInsiderTransactionsInner {
	this := GetInsiderTransactions200ResponseInsiderTransactionsInner{}
	return &this
}

// GetFullName returns the FullName field value
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetFullName(v string) {
	o.FullName = v
}

// GetPosition returns the Position field value
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetPosition(v string) {
	o.Position = v
}

// GetDateReported returns the DateReported field value
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetDateReported() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateReported
}

// GetDateReportedOk returns a tuple with the DateReported field value
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetDateReportedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateReported, true
}

// SetDateReported sets field value
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetDateReported(v string) {
	o.DateReported = v
}

// GetIsDirect returns the IsDirect field value if set, zero value otherwise.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetIsDirect() bool {
	if o == nil || IsNil(o.IsDirect) {
		var ret bool
		return ret
	}
	return *o.IsDirect
}

// GetIsDirectOk returns a tuple with the IsDirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetIsDirectOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDirect) {
		return nil, false
	}
	return o.IsDirect, true
}

// HasIsDirect returns a boolean if a field has been set.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) HasIsDirect() bool {
	if o != nil && !IsNil(o.IsDirect) {
		return true
	}

	return false
}

// SetIsDirect gets a reference to the given bool and assigns it to the IsDirect field.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetIsDirect(v bool) {
	o.IsDirect = &v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetShares() int64 {
	if o == nil || IsNil(o.Shares) {
		var ret int64
		return ret
	}
	return *o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetSharesOk() (*int64, bool) {
	if o == nil || IsNil(o.Shares) {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) HasShares() bool {
	if o != nil && !IsNil(o.Shares) {
		return true
	}

	return false
}

// SetShares gets a reference to the given int64 and assigns it to the Shares field.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetShares(v int64) {
	o.Shares = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetValue() int64 {
	if o == nil || IsNil(o.Value) {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetValueOk() (*int64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetValue(v int64) {
	o.Value = &v
}

// GetDescription returns the Description field value
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) SetDescription(v string) {
	o.Description = v
}

func (o GetInsiderTransactions200ResponseInsiderTransactionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsiderTransactions200ResponseInsiderTransactionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["full_name"] = o.FullName
	toSerialize["position"] = o.Position
	toSerialize["date_reported"] = o.DateReported
	if !IsNil(o.IsDirect) {
		toSerialize["is_direct"] = o.IsDirect
	}
	if !IsNil(o.Shares) {
		toSerialize["shares"] = o.Shares
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *GetInsiderTransactions200ResponseInsiderTransactionsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"full_name",
		"position",
		"date_reported",
		"description",
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

	varGetInsiderTransactions200ResponseInsiderTransactionsInner := _GetInsiderTransactions200ResponseInsiderTransactionsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInsiderTransactions200ResponseInsiderTransactionsInner)

	if err != nil {
		return err
	}

	*o = GetInsiderTransactions200ResponseInsiderTransactionsInner(varGetInsiderTransactions200ResponseInsiderTransactionsInner)

	return err
}

type NullableGetInsiderTransactions200ResponseInsiderTransactionsInner struct {
	value *GetInsiderTransactions200ResponseInsiderTransactionsInner
	isSet bool
}

func (v NullableGetInsiderTransactions200ResponseInsiderTransactionsInner) Get() *GetInsiderTransactions200ResponseInsiderTransactionsInner {
	return v.value
}

func (v *NullableGetInsiderTransactions200ResponseInsiderTransactionsInner) Set(val *GetInsiderTransactions200ResponseInsiderTransactionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsiderTransactions200ResponseInsiderTransactionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsiderTransactions200ResponseInsiderTransactionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsiderTransactions200ResponseInsiderTransactionsInner(val *GetInsiderTransactions200ResponseInsiderTransactionsInner) *NullableGetInsiderTransactions200ResponseInsiderTransactionsInner {
	return &NullableGetInsiderTransactions200ResponseInsiderTransactionsInner{value: val, isSet: true}
}

func (v NullableGetInsiderTransactions200ResponseInsiderTransactionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsiderTransactions200ResponseInsiderTransactionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
