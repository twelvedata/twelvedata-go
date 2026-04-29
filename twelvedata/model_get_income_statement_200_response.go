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

// checks if the GetIncomeStatement200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIncomeStatement200Response{}

// GetIncomeStatement200Response struct for GetIncomeStatement200Response
type GetIncomeStatement200Response struct {
	Meta GetIncomeStatement200ResponseMeta `json:"meta"`
	// Income statement data
	IncomeStatement []IncomeStatementBlock `json:"income_statement,omitempty"`
}

type _GetIncomeStatement200Response GetIncomeStatement200Response

// NewGetIncomeStatement200Response instantiates a new GetIncomeStatement200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIncomeStatement200Response(meta GetIncomeStatement200ResponseMeta) *GetIncomeStatement200Response {
	this := GetIncomeStatement200Response{}
	this.Meta = meta
	return &this
}

// NewGetIncomeStatement200ResponseWithDefaults instantiates a new GetIncomeStatement200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIncomeStatement200ResponseWithDefaults() *GetIncomeStatement200Response {
	this := GetIncomeStatement200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetIncomeStatement200Response) GetMeta() GetIncomeStatement200ResponseMeta {
	if o == nil {
		var ret GetIncomeStatement200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetIncomeStatement200Response) GetMetaOk() (*GetIncomeStatement200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetIncomeStatement200Response) SetMeta(v GetIncomeStatement200ResponseMeta) {
	o.Meta = v
}

// GetIncomeStatement returns the IncomeStatement field value if set, zero value otherwise.
func (o *GetIncomeStatement200Response) GetIncomeStatement() []IncomeStatementBlock {
	if o == nil || IsNil(o.IncomeStatement) {
		var ret []IncomeStatementBlock
		return ret
	}
	return o.IncomeStatement
}

// GetIncomeStatementOk returns a tuple with the IncomeStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIncomeStatement200Response) GetIncomeStatementOk() ([]IncomeStatementBlock, bool) {
	if o == nil || IsNil(o.IncomeStatement) {
		return nil, false
	}
	return o.IncomeStatement, true
}

// HasIncomeStatement returns a boolean if a field has been set.
func (o *GetIncomeStatement200Response) HasIncomeStatement() bool {
	if o != nil && !IsNil(o.IncomeStatement) {
		return true
	}

	return false
}

// SetIncomeStatement gets a reference to the given []IncomeStatementBlock and assigns it to the IncomeStatement field.
func (o *GetIncomeStatement200Response) SetIncomeStatement(v []IncomeStatementBlock) {
	o.IncomeStatement = v
}

func (o GetIncomeStatement200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIncomeStatement200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	if !IsNil(o.IncomeStatement) {
		toSerialize["income_statement"] = o.IncomeStatement
	}
	return toSerialize, nil
}

func (o *GetIncomeStatement200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
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

	varGetIncomeStatement200Response := _GetIncomeStatement200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetIncomeStatement200Response)

	if err != nil {
		return err
	}

	*o = GetIncomeStatement200Response(varGetIncomeStatement200Response)

	return err
}

type NullableGetIncomeStatement200Response struct {
	value *GetIncomeStatement200Response
	isSet bool
}

func (v NullableGetIncomeStatement200Response) Get() *GetIncomeStatement200Response {
	return v.value
}

func (v *NullableGetIncomeStatement200Response) Set(val *GetIncomeStatement200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIncomeStatement200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIncomeStatement200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIncomeStatement200Response(val *GetIncomeStatement200Response) *NullableGetIncomeStatement200Response {
	return &NullableGetIncomeStatement200Response{value: val, isSet: true}
}

func (v NullableGetIncomeStatement200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIncomeStatement200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
