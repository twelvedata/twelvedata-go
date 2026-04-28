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

// checks if the GetApiUsage200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetApiUsage200Response{}

// GetApiUsage200Response struct for GetApiUsage200Response
type GetApiUsage200Response struct {
	// Current timestamp in UTC timezone
	Timestamp string `json:"timestamp"`
	// Number of requests made in last minute
	CurrentUsage int64 `json:"current_usage"`
	// Your personal API limit (requests/minute) depending on the plan
	PlanLimit int64 `json:"plan_limit"`
	// Plan category name
	PlanCategory *string `json:"plan_category,omitempty"`
}

type _GetApiUsage200Response GetApiUsage200Response

// NewGetApiUsage200Response instantiates a new GetApiUsage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApiUsage200Response(timestamp string, currentUsage int64, planLimit int64) *GetApiUsage200Response {
	this := GetApiUsage200Response{}
	this.Timestamp = timestamp
	this.CurrentUsage = currentUsage
	this.PlanLimit = planLimit
	return &this
}

// NewGetApiUsage200ResponseWithDefaults instantiates a new GetApiUsage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApiUsage200ResponseWithDefaults() *GetApiUsage200Response {
	this := GetApiUsage200Response{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *GetApiUsage200Response) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *GetApiUsage200Response) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *GetApiUsage200Response) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetCurrentUsage returns the CurrentUsage field value
func (o *GetApiUsage200Response) GetCurrentUsage() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CurrentUsage
}

// GetCurrentUsageOk returns a tuple with the CurrentUsage field value
// and a boolean to check if the value has been set.
func (o *GetApiUsage200Response) GetCurrentUsageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentUsage, true
}

// SetCurrentUsage sets field value
func (o *GetApiUsage200Response) SetCurrentUsage(v int64) {
	o.CurrentUsage = v
}

// GetPlanLimit returns the PlanLimit field value
func (o *GetApiUsage200Response) GetPlanLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanLimit
}

// GetPlanLimitOk returns a tuple with the PlanLimit field value
// and a boolean to check if the value has been set.
func (o *GetApiUsage200Response) GetPlanLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanLimit, true
}

// SetPlanLimit sets field value
func (o *GetApiUsage200Response) SetPlanLimit(v int64) {
	o.PlanLimit = v
}

// GetPlanCategory returns the PlanCategory field value if set, zero value otherwise.
func (o *GetApiUsage200Response) GetPlanCategory() string {
	if o == nil || IsNil(o.PlanCategory) {
		var ret string
		return ret
	}
	return *o.PlanCategory
}

// GetPlanCategoryOk returns a tuple with the PlanCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiUsage200Response) GetPlanCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.PlanCategory) {
		return nil, false
	}
	return o.PlanCategory, true
}

// HasPlanCategory returns a boolean if a field has been set.
func (o *GetApiUsage200Response) HasPlanCategory() bool {
	if o != nil && !IsNil(o.PlanCategory) {
		return true
	}

	return false
}

// SetPlanCategory gets a reference to the given string and assigns it to the PlanCategory field.
func (o *GetApiUsage200Response) SetPlanCategory(v string) {
	o.PlanCategory = &v
}

func (o GetApiUsage200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApiUsage200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["current_usage"] = o.CurrentUsage
	toSerialize["plan_limit"] = o.PlanLimit
	if !IsNil(o.PlanCategory) {
		toSerialize["plan_category"] = o.PlanCategory
	}
	return toSerialize, nil
}

func (o *GetApiUsage200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timestamp",
		"current_usage",
		"plan_limit",
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

	varGetApiUsage200Response := _GetApiUsage200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetApiUsage200Response)

	if err != nil {
		return err
	}

	*o = GetApiUsage200Response(varGetApiUsage200Response)

	return err
}

type NullableGetApiUsage200Response struct {
	value *GetApiUsage200Response
	isSet bool
}

func (v NullableGetApiUsage200Response) Get() *GetApiUsage200Response {
	return v.value
}

func (v *NullableGetApiUsage200Response) Set(val *GetApiUsage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApiUsage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApiUsage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApiUsage200Response(val *GetApiUsage200Response) *NullableGetApiUsage200Response {
	return &NullableGetApiUsage200Response{value: val, isSet: true}
}

func (v NullableGetApiUsage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApiUsage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
