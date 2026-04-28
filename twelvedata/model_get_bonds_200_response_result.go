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

// checks if the GetBonds200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBonds200ResponseResult{}

// GetBonds200ResponseResult struct for GetBonds200ResponseResult
type GetBonds200ResponseResult struct {
	// Total number of matching instruments
	Count int64              `json:"count"`
	List  []BondResponseItem `json:"list,omitempty"`
}

type _GetBonds200ResponseResult GetBonds200ResponseResult

// NewGetBonds200ResponseResult instantiates a new GetBonds200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBonds200ResponseResult(count int64) *GetBonds200ResponseResult {
	this := GetBonds200ResponseResult{}
	this.Count = count
	return &this
}

// NewGetBonds200ResponseResultWithDefaults instantiates a new GetBonds200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBonds200ResponseResultWithDefaults() *GetBonds200ResponseResult {
	this := GetBonds200ResponseResult{}
	return &this
}

// GetCount returns the Count field value
func (o *GetBonds200ResponseResult) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetBonds200ResponseResult) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetBonds200ResponseResult) SetCount(v int64) {
	o.Count = v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *GetBonds200ResponseResult) GetList() []BondResponseItem {
	if o == nil || IsNil(o.List) {
		var ret []BondResponseItem
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBonds200ResponseResult) GetListOk() ([]BondResponseItem, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *GetBonds200ResponseResult) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []BondResponseItem and assigns it to the List field.
func (o *GetBonds200ResponseResult) SetList(v []BondResponseItem) {
	o.List = v
}

func (o GetBonds200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBonds200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

func (o *GetBonds200ResponseResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
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

	varGetBonds200ResponseResult := _GetBonds200ResponseResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetBonds200ResponseResult)

	if err != nil {
		return err
	}

	*o = GetBonds200ResponseResult(varGetBonds200ResponseResult)

	return err
}

type NullableGetBonds200ResponseResult struct {
	value *GetBonds200ResponseResult
	isSet bool
}

func (v NullableGetBonds200ResponseResult) Get() *GetBonds200ResponseResult {
	return v.value
}

func (v *NullableGetBonds200ResponseResult) Set(val *GetBonds200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBonds200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBonds200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBonds200ResponseResult(val *GetBonds200ResponseResult) *NullableGetBonds200ResponseResult {
	return &NullableGetBonds200ResponseResult{value: val, isSet: true}
}

func (v NullableGetBonds200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBonds200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
