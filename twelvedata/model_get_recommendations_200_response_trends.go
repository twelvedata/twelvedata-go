/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetRecommendations200ResponseTrends type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecommendations200ResponseTrends{}

// GetRecommendations200ResponseTrends Analyst recommendations trends
type GetRecommendations200ResponseTrends struct {
	CurrentMonth  *GetRecommendations200ResponseTrendsCurrentMonth  `json:"current_month,omitempty"`
	PreviousMonth *GetRecommendations200ResponseTrendsPreviousMonth `json:"previous_month,omitempty"`
	Var2MonthsAgo *GetRecommendations200ResponseTrends2MonthsAgo    `json:"2_months_ago,omitempty"`
	Var3MonthsAgo *GetRecommendations200ResponseTrends3MonthsAgo    `json:"3_months_ago,omitempty"`
}

// NewGetRecommendations200ResponseTrends instantiates a new GetRecommendations200ResponseTrends object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecommendations200ResponseTrends() *GetRecommendations200ResponseTrends {
	this := GetRecommendations200ResponseTrends{}
	return &this
}

// NewGetRecommendations200ResponseTrendsWithDefaults instantiates a new GetRecommendations200ResponseTrends object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecommendations200ResponseTrendsWithDefaults() *GetRecommendations200ResponseTrends {
	this := GetRecommendations200ResponseTrends{}
	return &this
}

// GetCurrentMonth returns the CurrentMonth field value if set, zero value otherwise.
func (o *GetRecommendations200ResponseTrends) GetCurrentMonth() GetRecommendations200ResponseTrendsCurrentMonth {
	if o == nil || IsNil(o.CurrentMonth) {
		var ret GetRecommendations200ResponseTrendsCurrentMonth
		return ret
	}
	return *o.CurrentMonth
}

// GetCurrentMonthOk returns a tuple with the CurrentMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecommendations200ResponseTrends) GetCurrentMonthOk() (*GetRecommendations200ResponseTrendsCurrentMonth, bool) {
	if o == nil || IsNil(o.CurrentMonth) {
		return nil, false
	}
	return o.CurrentMonth, true
}

// HasCurrentMonth returns a boolean if a field has been set.
func (o *GetRecommendations200ResponseTrends) HasCurrentMonth() bool {
	if o != nil && !IsNil(o.CurrentMonth) {
		return true
	}

	return false
}

// SetCurrentMonth gets a reference to the given GetRecommendations200ResponseTrendsCurrentMonth and assigns it to the CurrentMonth field.
func (o *GetRecommendations200ResponseTrends) SetCurrentMonth(v GetRecommendations200ResponseTrendsCurrentMonth) {
	o.CurrentMonth = &v
}

// GetPreviousMonth returns the PreviousMonth field value if set, zero value otherwise.
func (o *GetRecommendations200ResponseTrends) GetPreviousMonth() GetRecommendations200ResponseTrendsPreviousMonth {
	if o == nil || IsNil(o.PreviousMonth) {
		var ret GetRecommendations200ResponseTrendsPreviousMonth
		return ret
	}
	return *o.PreviousMonth
}

// GetPreviousMonthOk returns a tuple with the PreviousMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecommendations200ResponseTrends) GetPreviousMonthOk() (*GetRecommendations200ResponseTrendsPreviousMonth, bool) {
	if o == nil || IsNil(o.PreviousMonth) {
		return nil, false
	}
	return o.PreviousMonth, true
}

// HasPreviousMonth returns a boolean if a field has been set.
func (o *GetRecommendations200ResponseTrends) HasPreviousMonth() bool {
	if o != nil && !IsNil(o.PreviousMonth) {
		return true
	}

	return false
}

// SetPreviousMonth gets a reference to the given GetRecommendations200ResponseTrendsPreviousMonth and assigns it to the PreviousMonth field.
func (o *GetRecommendations200ResponseTrends) SetPreviousMonth(v GetRecommendations200ResponseTrendsPreviousMonth) {
	o.PreviousMonth = &v
}

// GetVar2MonthsAgo returns the Var2MonthsAgo field value if set, zero value otherwise.
func (o *GetRecommendations200ResponseTrends) GetVar2MonthsAgo() GetRecommendations200ResponseTrends2MonthsAgo {
	if o == nil || IsNil(o.Var2MonthsAgo) {
		var ret GetRecommendations200ResponseTrends2MonthsAgo
		return ret
	}
	return *o.Var2MonthsAgo
}

// GetVar2MonthsAgoOk returns a tuple with the Var2MonthsAgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecommendations200ResponseTrends) GetVar2MonthsAgoOk() (*GetRecommendations200ResponseTrends2MonthsAgo, bool) {
	if o == nil || IsNil(o.Var2MonthsAgo) {
		return nil, false
	}
	return o.Var2MonthsAgo, true
}

// HasVar2MonthsAgo returns a boolean if a field has been set.
func (o *GetRecommendations200ResponseTrends) HasVar2MonthsAgo() bool {
	if o != nil && !IsNil(o.Var2MonthsAgo) {
		return true
	}

	return false
}

// SetVar2MonthsAgo gets a reference to the given GetRecommendations200ResponseTrends2MonthsAgo and assigns it to the Var2MonthsAgo field.
func (o *GetRecommendations200ResponseTrends) SetVar2MonthsAgo(v GetRecommendations200ResponseTrends2MonthsAgo) {
	o.Var2MonthsAgo = &v
}

// GetVar3MonthsAgo returns the Var3MonthsAgo field value if set, zero value otherwise.
func (o *GetRecommendations200ResponseTrends) GetVar3MonthsAgo() GetRecommendations200ResponseTrends3MonthsAgo {
	if o == nil || IsNil(o.Var3MonthsAgo) {
		var ret GetRecommendations200ResponseTrends3MonthsAgo
		return ret
	}
	return *o.Var3MonthsAgo
}

// GetVar3MonthsAgoOk returns a tuple with the Var3MonthsAgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecommendations200ResponseTrends) GetVar3MonthsAgoOk() (*GetRecommendations200ResponseTrends3MonthsAgo, bool) {
	if o == nil || IsNil(o.Var3MonthsAgo) {
		return nil, false
	}
	return o.Var3MonthsAgo, true
}

// HasVar3MonthsAgo returns a boolean if a field has been set.
func (o *GetRecommendations200ResponseTrends) HasVar3MonthsAgo() bool {
	if o != nil && !IsNil(o.Var3MonthsAgo) {
		return true
	}

	return false
}

// SetVar3MonthsAgo gets a reference to the given GetRecommendations200ResponseTrends3MonthsAgo and assigns it to the Var3MonthsAgo field.
func (o *GetRecommendations200ResponseTrends) SetVar3MonthsAgo(v GetRecommendations200ResponseTrends3MonthsAgo) {
	o.Var3MonthsAgo = &v
}

func (o GetRecommendations200ResponseTrends) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecommendations200ResponseTrends) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentMonth) {
		toSerialize["current_month"] = o.CurrentMonth
	}
	if !IsNil(o.PreviousMonth) {
		toSerialize["previous_month"] = o.PreviousMonth
	}
	if !IsNil(o.Var2MonthsAgo) {
		toSerialize["2_months_ago"] = o.Var2MonthsAgo
	}
	if !IsNil(o.Var3MonthsAgo) {
		toSerialize["3_months_ago"] = o.Var3MonthsAgo
	}
	return toSerialize, nil
}

type NullableGetRecommendations200ResponseTrends struct {
	value *GetRecommendations200ResponseTrends
	isSet bool
}

func (v NullableGetRecommendations200ResponseTrends) Get() *GetRecommendations200ResponseTrends {
	return v.value
}

func (v *NullableGetRecommendations200ResponseTrends) Set(val *GetRecommendations200ResponseTrends) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecommendations200ResponseTrends) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecommendations200ResponseTrends) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecommendations200ResponseTrends(val *GetRecommendations200ResponseTrends) *NullableGetRecommendations200ResponseTrends {
	return &NullableGetRecommendations200ResponseTrends{value: val, isSet: true}
}

func (v NullableGetRecommendations200ResponseTrends) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecommendations200ResponseTrends) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
