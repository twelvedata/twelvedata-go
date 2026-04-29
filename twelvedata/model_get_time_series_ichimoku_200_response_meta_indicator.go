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

// checks if the GetTimeSeriesIchimoku200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesIchimoku200ResponseMetaIndicator{}

// GetTimeSeriesIchimoku200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesIchimoku200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// The time period used for generating the conversation line
	ConversionLinePeriod int64 `json:"conversion_line_period"`
	// The time period used for generating the base line
	BaseLinePeriod int64 `json:"base_line_period"`
	// The time period used for generating the leading span B line
	LeadingSpanBPeriod int64 `json:"leading_span_b_period"`
	// The time period used for generating the lagging span line
	LaggingSpanPeriod int64 `json:"lagging_span_period"`
	// Indicates whether to include ahead span period
	IncludeAheadSpanPeriod bool `json:"include_ahead_span_period"`
}

type _GetTimeSeriesIchimoku200ResponseMetaIndicator GetTimeSeriesIchimoku200ResponseMetaIndicator

// NewGetTimeSeriesIchimoku200ResponseMetaIndicator instantiates a new GetTimeSeriesIchimoku200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesIchimoku200ResponseMetaIndicator(name string, conversionLinePeriod int64, baseLinePeriod int64, leadingSpanBPeriod int64, laggingSpanPeriod int64, includeAheadSpanPeriod bool) *GetTimeSeriesIchimoku200ResponseMetaIndicator {
	this := GetTimeSeriesIchimoku200ResponseMetaIndicator{}
	this.Name = name
	this.ConversionLinePeriod = conversionLinePeriod
	this.BaseLinePeriod = baseLinePeriod
	this.LeadingSpanBPeriod = leadingSpanBPeriod
	this.LaggingSpanPeriod = laggingSpanPeriod
	this.IncludeAheadSpanPeriod = includeAheadSpanPeriod
	return &this
}

// NewGetTimeSeriesIchimoku200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesIchimoku200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesIchimoku200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesIchimoku200ResponseMetaIndicator {
	this := GetTimeSeriesIchimoku200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetConversionLinePeriod returns the ConversionLinePeriod field value
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetConversionLinePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ConversionLinePeriod
}

// GetConversionLinePeriodOk returns a tuple with the ConversionLinePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetConversionLinePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversionLinePeriod, true
}

// SetConversionLinePeriod sets field value
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) SetConversionLinePeriod(v int64) {
	o.ConversionLinePeriod = v
}

// GetBaseLinePeriod returns the BaseLinePeriod field value
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetBaseLinePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BaseLinePeriod
}

// GetBaseLinePeriodOk returns a tuple with the BaseLinePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetBaseLinePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseLinePeriod, true
}

// SetBaseLinePeriod sets field value
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) SetBaseLinePeriod(v int64) {
	o.BaseLinePeriod = v
}

// GetLeadingSpanBPeriod returns the LeadingSpanBPeriod field value
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetLeadingSpanBPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LeadingSpanBPeriod
}

// GetLeadingSpanBPeriodOk returns a tuple with the LeadingSpanBPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetLeadingSpanBPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeadingSpanBPeriod, true
}

// SetLeadingSpanBPeriod sets field value
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) SetLeadingSpanBPeriod(v int64) {
	o.LeadingSpanBPeriod = v
}

// GetLaggingSpanPeriod returns the LaggingSpanPeriod field value
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetLaggingSpanPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LaggingSpanPeriod
}

// GetLaggingSpanPeriodOk returns a tuple with the LaggingSpanPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetLaggingSpanPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LaggingSpanPeriod, true
}

// SetLaggingSpanPeriod sets field value
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) SetLaggingSpanPeriod(v int64) {
	o.LaggingSpanPeriod = v
}

// GetIncludeAheadSpanPeriod returns the IncludeAheadSpanPeriod field value
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetIncludeAheadSpanPeriod() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeAheadSpanPeriod
}

// GetIncludeAheadSpanPeriodOk returns a tuple with the IncludeAheadSpanPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetIncludeAheadSpanPeriodOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeAheadSpanPeriod, true
}

// SetIncludeAheadSpanPeriod sets field value
func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) SetIncludeAheadSpanPeriod(v bool) {
	o.IncludeAheadSpanPeriod = v
}

func (o GetTimeSeriesIchimoku200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesIchimoku200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["conversion_line_period"] = o.ConversionLinePeriod
	toSerialize["base_line_period"] = o.BaseLinePeriod
	toSerialize["leading_span_b_period"] = o.LeadingSpanBPeriod
	toSerialize["lagging_span_period"] = o.LaggingSpanPeriod
	toSerialize["include_ahead_span_period"] = o.IncludeAheadSpanPeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"conversion_line_period",
		"base_line_period",
		"leading_span_b_period",
		"lagging_span_period",
		"include_ahead_span_period",
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

	varGetTimeSeriesIchimoku200ResponseMetaIndicator := _GetTimeSeriesIchimoku200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesIchimoku200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesIchimoku200ResponseMetaIndicator(varGetTimeSeriesIchimoku200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesIchimoku200ResponseMetaIndicator struct {
	value *GetTimeSeriesIchimoku200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesIchimoku200ResponseMetaIndicator) Get() *GetTimeSeriesIchimoku200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesIchimoku200ResponseMetaIndicator) Set(val *GetTimeSeriesIchimoku200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesIchimoku200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesIchimoku200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesIchimoku200ResponseMetaIndicator(val *GetTimeSeriesIchimoku200ResponseMetaIndicator) *NullableGetTimeSeriesIchimoku200ResponseMetaIndicator {
	return &NullableGetTimeSeriesIchimoku200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesIchimoku200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesIchimoku200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
