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

// checks if the GetMutualFundsList200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsList200ResponseResult{}

// GetMutualFundsList200ResponseResult Response result
type GetMutualFundsList200ResponseResult struct {
	// Total number of matching funds
	Count int64 `json:"count"`
	// List of mutual funds
	List []MutualFundsListResponseListItem `json:"list"`
}

type _GetMutualFundsList200ResponseResult GetMutualFundsList200ResponseResult

// NewGetMutualFundsList200ResponseResult instantiates a new GetMutualFundsList200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsList200ResponseResult(count int64, list []MutualFundsListResponseListItem) *GetMutualFundsList200ResponseResult {
	this := GetMutualFundsList200ResponseResult{}
	this.Count = count
	this.List = list
	return &this
}

// NewGetMutualFundsList200ResponseResultWithDefaults instantiates a new GetMutualFundsList200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsList200ResponseResultWithDefaults() *GetMutualFundsList200ResponseResult {
	this := GetMutualFundsList200ResponseResult{}
	return &this
}

// GetCount returns the Count field value
func (o *GetMutualFundsList200ResponseResult) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsList200ResponseResult) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetMutualFundsList200ResponseResult) SetCount(v int64) {
	o.Count = v
}

// GetList returns the List field value
func (o *GetMutualFundsList200ResponseResult) GetList() []MutualFundsListResponseListItem {
	if o == nil {
		var ret []MutualFundsListResponseListItem
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *GetMutualFundsList200ResponseResult) GetListOk() ([]MutualFundsListResponseListItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.List, true
}

// SetList sets field value
func (o *GetMutualFundsList200ResponseResult) SetList(v []MutualFundsListResponseListItem) {
	o.List = v
}

func (o GetMutualFundsList200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsList200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["list"] = o.List
	return toSerialize, nil
}

func (o *GetMutualFundsList200ResponseResult) UnmarshalJSON(data []byte) (err error) {
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

	varGetMutualFundsList200ResponseResult := _GetMutualFundsList200ResponseResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMutualFundsList200ResponseResult)

	if err != nil {
		return err
	}

	*o = GetMutualFundsList200ResponseResult(varGetMutualFundsList200ResponseResult)

	return err
}

type NullableGetMutualFundsList200ResponseResult struct {
	value *GetMutualFundsList200ResponseResult
	isSet bool
}

func (v NullableGetMutualFundsList200ResponseResult) Get() *GetMutualFundsList200ResponseResult {
	return v.value
}

func (v *NullableGetMutualFundsList200ResponseResult) Set(val *GetMutualFundsList200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsList200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsList200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsList200ResponseResult(val *GetMutualFundsList200ResponseResult) *NullableGetMutualFundsList200ResponseResult {
	return &NullableGetMutualFundsList200ResponseResult{value: val, isSet: true}
}

func (v NullableGetMutualFundsList200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsList200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
