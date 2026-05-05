// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the InlineObject13MetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject13MetaIndicator{}

// InlineObject13MetaIndicator struct for InlineObject13MetaIndicator
type InlineObject13MetaIndicator struct {
	MaType     *string `json:"ma_type,omitempty"`
	MaxPeriod  *int64  `json:"max_period,omitempty"`
	MinPeriod  *int64  `json:"min_period,omitempty"`
	Name       *string `json:"name,omitempty"`
	Periods    []int64 `json:"periods,omitempty"`
	SeriesType *string `json:"series_type,omitempty"`
}

// NewInlineObject13MetaIndicator instantiates a new InlineObject13MetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject13MetaIndicator() *InlineObject13MetaIndicator {
	this := InlineObject13MetaIndicator{}
	return &this
}

// NewInlineObject13MetaIndicatorWithDefaults instantiates a new InlineObject13MetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject13MetaIndicatorWithDefaults() *InlineObject13MetaIndicator {
	this := InlineObject13MetaIndicator{}
	return &this
}

// GetMaType returns the MaType field value if set, zero value otherwise.
func (o *InlineObject13MetaIndicator) GetMaType() string {
	if o == nil || IsNil(o.MaType) {
		var ret string
		return ret
	}
	return *o.MaType
}

// GetMaTypeOk returns a tuple with the MaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject13MetaIndicator) GetMaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MaType) {
		return nil, false
	}
	return o.MaType, true
}

// HasMaType returns a boolean if a field has been set.
func (o *InlineObject13MetaIndicator) HasMaType() bool {
	if o != nil && !IsNil(o.MaType) {
		return true
	}

	return false
}

// SetMaType gets a reference to the given string and assigns it to the MaType field.
func (o *InlineObject13MetaIndicator) SetMaType(v string) {
	o.MaType = &v
}

// GetMaxPeriod returns the MaxPeriod field value if set, zero value otherwise.
func (o *InlineObject13MetaIndicator) GetMaxPeriod() int64 {
	if o == nil || IsNil(o.MaxPeriod) {
		var ret int64
		return ret
	}
	return *o.MaxPeriod
}

// GetMaxPeriodOk returns a tuple with the MaxPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject13MetaIndicator) GetMaxPeriodOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxPeriod) {
		return nil, false
	}
	return o.MaxPeriod, true
}

// HasMaxPeriod returns a boolean if a field has been set.
func (o *InlineObject13MetaIndicator) HasMaxPeriod() bool {
	if o != nil && !IsNil(o.MaxPeriod) {
		return true
	}

	return false
}

// SetMaxPeriod gets a reference to the given int64 and assigns it to the MaxPeriod field.
func (o *InlineObject13MetaIndicator) SetMaxPeriod(v int64) {
	o.MaxPeriod = &v
}

// GetMinPeriod returns the MinPeriod field value if set, zero value otherwise.
func (o *InlineObject13MetaIndicator) GetMinPeriod() int64 {
	if o == nil || IsNil(o.MinPeriod) {
		var ret int64
		return ret
	}
	return *o.MinPeriod
}

// GetMinPeriodOk returns a tuple with the MinPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject13MetaIndicator) GetMinPeriodOk() (*int64, bool) {
	if o == nil || IsNil(o.MinPeriod) {
		return nil, false
	}
	return o.MinPeriod, true
}

// HasMinPeriod returns a boolean if a field has been set.
func (o *InlineObject13MetaIndicator) HasMinPeriod() bool {
	if o != nil && !IsNil(o.MinPeriod) {
		return true
	}

	return false
}

// SetMinPeriod gets a reference to the given int64 and assigns it to the MinPeriod field.
func (o *InlineObject13MetaIndicator) SetMinPeriod(v int64) {
	o.MinPeriod = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject13MetaIndicator) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject13MetaIndicator) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject13MetaIndicator) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject13MetaIndicator) SetName(v string) {
	o.Name = &v
}

// GetPeriods returns the Periods field value if set, zero value otherwise.
func (o *InlineObject13MetaIndicator) GetPeriods() []int64 {
	if o == nil || IsNil(o.Periods) {
		var ret []int64
		return ret
	}
	return o.Periods
}

// GetPeriodsOk returns a tuple with the Periods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject13MetaIndicator) GetPeriodsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Periods) {
		return nil, false
	}
	return o.Periods, true
}

// HasPeriods returns a boolean if a field has been set.
func (o *InlineObject13MetaIndicator) HasPeriods() bool {
	if o != nil && !IsNil(o.Periods) {
		return true
	}

	return false
}

// SetPeriods gets a reference to the given []int64 and assigns it to the Periods field.
func (o *InlineObject13MetaIndicator) SetPeriods(v []int64) {
	o.Periods = v
}

// GetSeriesType returns the SeriesType field value if set, zero value otherwise.
func (o *InlineObject13MetaIndicator) GetSeriesType() string {
	if o == nil || IsNil(o.SeriesType) {
		var ret string
		return ret
	}
	return *o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject13MetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SeriesType) {
		return nil, false
	}
	return o.SeriesType, true
}

// HasSeriesType returns a boolean if a field has been set.
func (o *InlineObject13MetaIndicator) HasSeriesType() bool {
	if o != nil && !IsNil(o.SeriesType) {
		return true
	}

	return false
}

// SetSeriesType gets a reference to the given string and assigns it to the SeriesType field.
func (o *InlineObject13MetaIndicator) SetSeriesType(v string) {
	o.SeriesType = &v
}

func (o InlineObject13MetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject13MetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaType) {
		toSerialize["ma_type"] = o.MaType
	}
	if !IsNil(o.MaxPeriod) {
		toSerialize["max_period"] = o.MaxPeriod
	}
	if !IsNil(o.MinPeriod) {
		toSerialize["min_period"] = o.MinPeriod
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Periods) {
		toSerialize["periods"] = o.Periods
	}
	if !IsNil(o.SeriesType) {
		toSerialize["series_type"] = o.SeriesType
	}
	return toSerialize, nil
}

type NullableInlineObject13MetaIndicator struct {
	value *InlineObject13MetaIndicator
	isSet bool
}

func (v NullableInlineObject13MetaIndicator) Get() *InlineObject13MetaIndicator {
	return v.value
}

func (v *NullableInlineObject13MetaIndicator) Set(val *InlineObject13MetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject13MetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject13MetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject13MetaIndicator(val *InlineObject13MetaIndicator) *NullableInlineObject13MetaIndicator {
	return &NullableInlineObject13MetaIndicator{value: val, isSet: true}
}

func (v NullableInlineObject13MetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject13MetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
