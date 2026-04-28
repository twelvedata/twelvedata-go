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

// checks if the GetFunds200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFunds200ResponseResult{}

// GetFunds200ResponseResult struct for GetFunds200ResponseResult
type GetFunds200ResponseResult struct {
	// Total number of matching instruments
	Count int64 `json:"count"`
	// List of funds
	List []FundResponseItem `json:"list"`
}

type _GetFunds200ResponseResult GetFunds200ResponseResult

// NewGetFunds200ResponseResult instantiates a new GetFunds200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFunds200ResponseResult(count int64, list []FundResponseItem) *GetFunds200ResponseResult {
	this := GetFunds200ResponseResult{}
	this.Count = count
	this.List = list
	return &this
}

// NewGetFunds200ResponseResultWithDefaults instantiates a new GetFunds200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFunds200ResponseResultWithDefaults() *GetFunds200ResponseResult {
	this := GetFunds200ResponseResult{}
	return &this
}

// GetCount returns the Count field value
func (o *GetFunds200ResponseResult) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetFunds200ResponseResult) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetFunds200ResponseResult) SetCount(v int64) {
	o.Count = v
}

// GetList returns the List field value
func (o *GetFunds200ResponseResult) GetList() []FundResponseItem {
	if o == nil {
		var ret []FundResponseItem
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *GetFunds200ResponseResult) GetListOk() ([]FundResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.List, true
}

// SetList sets field value
func (o *GetFunds200ResponseResult) SetList(v []FundResponseItem) {
	o.List = v
}

func (o GetFunds200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFunds200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["list"] = o.List
	return toSerialize, nil
}

func (o *GetFunds200ResponseResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"list",
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

	varGetFunds200ResponseResult := _GetFunds200ResponseResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFunds200ResponseResult)

	if err != nil {
		return err
	}

	*o = GetFunds200ResponseResult(varGetFunds200ResponseResult)

	return err
}

type NullableGetFunds200ResponseResult struct {
	value *GetFunds200ResponseResult
	isSet bool
}

func (v NullableGetFunds200ResponseResult) Get() *GetFunds200ResponseResult {
	return v.value
}

func (v *NullableGetFunds200ResponseResult) Set(val *GetFunds200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFunds200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFunds200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFunds200ResponseResult(val *GetFunds200ResponseResult) *NullableGetFunds200ResponseResult {
	return &NullableGetFunds200ResponseResult{value: val, isSet: true}
}

func (v NullableGetFunds200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFunds200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
