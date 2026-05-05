// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetIncomeStatementConsolidated200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIncomeStatementConsolidated200Response{}

// GetIncomeStatementConsolidated200Response struct for GetIncomeStatementConsolidated200Response
type GetIncomeStatementConsolidated200Response struct {
	// Income statement data
	IncomeStatement []IncomeStatementItem `json:"income_statement,omitempty"`
	// Response status
	Status *string `json:"status,omitempty"`
}

// NewGetIncomeStatementConsolidated200Response instantiates a new GetIncomeStatementConsolidated200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIncomeStatementConsolidated200Response() *GetIncomeStatementConsolidated200Response {
	this := GetIncomeStatementConsolidated200Response{}
	return &this
}

// NewGetIncomeStatementConsolidated200ResponseWithDefaults instantiates a new GetIncomeStatementConsolidated200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIncomeStatementConsolidated200ResponseWithDefaults() *GetIncomeStatementConsolidated200Response {
	this := GetIncomeStatementConsolidated200Response{}
	return &this
}

// GetIncomeStatement returns the IncomeStatement field value if set, zero value otherwise.
func (o *GetIncomeStatementConsolidated200Response) GetIncomeStatement() []IncomeStatementItem {
	if o == nil || IsNil(o.IncomeStatement) {
		var ret []IncomeStatementItem
		return ret
	}
	return o.IncomeStatement
}

// GetIncomeStatementOk returns a tuple with the IncomeStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIncomeStatementConsolidated200Response) GetIncomeStatementOk() ([]IncomeStatementItem, bool) {
	if o == nil || IsNil(o.IncomeStatement) {
		return nil, false
	}
	return o.IncomeStatement, true
}

// HasIncomeStatement returns a boolean if a field has been set.
func (o *GetIncomeStatementConsolidated200Response) HasIncomeStatement() bool {
	if o != nil && !IsNil(o.IncomeStatement) {
		return true
	}

	return false
}

// SetIncomeStatement gets a reference to the given []IncomeStatementItem and assigns it to the IncomeStatement field.
func (o *GetIncomeStatementConsolidated200Response) SetIncomeStatement(v []IncomeStatementItem) {
	o.IncomeStatement = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetIncomeStatementConsolidated200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIncomeStatementConsolidated200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetIncomeStatementConsolidated200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetIncomeStatementConsolidated200Response) SetStatus(v string) {
	o.Status = &v
}

func (o GetIncomeStatementConsolidated200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIncomeStatementConsolidated200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncomeStatement) {
		toSerialize["income_statement"] = o.IncomeStatement
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableGetIncomeStatementConsolidated200Response struct {
	value *GetIncomeStatementConsolidated200Response
	isSet bool
}

func (v NullableGetIncomeStatementConsolidated200Response) Get() *GetIncomeStatementConsolidated200Response {
	return v.value
}

func (v *NullableGetIncomeStatementConsolidated200Response) Set(val *GetIncomeStatementConsolidated200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIncomeStatementConsolidated200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIncomeStatementConsolidated200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIncomeStatementConsolidated200Response(val *GetIncomeStatementConsolidated200Response) *NullableGetIncomeStatementConsolidated200Response {
	return &NullableGetIncomeStatementConsolidated200Response{value: val, isSet: true}
}

func (v NullableGetIncomeStatementConsolidated200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIncomeStatementConsolidated200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
