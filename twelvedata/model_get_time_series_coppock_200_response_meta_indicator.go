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

// checks if the GetTimeSeriesCoppock200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCoppock200ResponseMetaIndicator{}

// GetTimeSeriesCoppock200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesCoppock200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// Number of periods for weighted moving average
	WmaPeriod int64 `json:"wma_period"`
	// Number of periods for long term rate of change
	LongRocPeriod int64 `json:"long_roc_period"`
	// Number of periods for short term rate of change
	ShortRocPeriod int64 `json:"short_roc_period"`
}

type _GetTimeSeriesCoppock200ResponseMetaIndicator GetTimeSeriesCoppock200ResponseMetaIndicator

// NewGetTimeSeriesCoppock200ResponseMetaIndicator instantiates a new GetTimeSeriesCoppock200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCoppock200ResponseMetaIndicator(name string, seriesType string, wmaPeriod int64, longRocPeriod int64, shortRocPeriod int64) *GetTimeSeriesCoppock200ResponseMetaIndicator {
	this := GetTimeSeriesCoppock200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.WmaPeriod = wmaPeriod
	this.LongRocPeriod = longRocPeriod
	this.ShortRocPeriod = shortRocPeriod
	return &this
}

// NewGetTimeSeriesCoppock200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesCoppock200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCoppock200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesCoppock200ResponseMetaIndicator {
	this := GetTimeSeriesCoppock200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetWmaPeriod returns the WmaPeriod field value
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) GetWmaPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.WmaPeriod
}

// GetWmaPeriodOk returns a tuple with the WmaPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) GetWmaPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WmaPeriod, true
}

// SetWmaPeriod sets field value
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) SetWmaPeriod(v int64) {
	o.WmaPeriod = v
}

// GetLongRocPeriod returns the LongRocPeriod field value
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) GetLongRocPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LongRocPeriod
}

// GetLongRocPeriodOk returns a tuple with the LongRocPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) GetLongRocPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LongRocPeriod, true
}

// SetLongRocPeriod sets field value
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) SetLongRocPeriod(v int64) {
	o.LongRocPeriod = v
}

// GetShortRocPeriod returns the ShortRocPeriod field value
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) GetShortRocPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ShortRocPeriod
}

// GetShortRocPeriodOk returns a tuple with the ShortRocPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) GetShortRocPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortRocPeriod, true
}

// SetShortRocPeriod sets field value
func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) SetShortRocPeriod(v int64) {
	o.ShortRocPeriod = v
}

func (o GetTimeSeriesCoppock200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCoppock200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["wma_period"] = o.WmaPeriod
	toSerialize["long_roc_period"] = o.LongRocPeriod
	toSerialize["short_roc_period"] = o.ShortRocPeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesCoppock200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
		"wma_period",
		"long_roc_period",
		"short_roc_period",
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

	varGetTimeSeriesCoppock200ResponseMetaIndicator := _GetTimeSeriesCoppock200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesCoppock200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCoppock200ResponseMetaIndicator(varGetTimeSeriesCoppock200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesCoppock200ResponseMetaIndicator struct {
	value *GetTimeSeriesCoppock200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesCoppock200ResponseMetaIndicator) Get() *GetTimeSeriesCoppock200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesCoppock200ResponseMetaIndicator) Set(val *GetTimeSeriesCoppock200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCoppock200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCoppock200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCoppock200ResponseMetaIndicator(val *GetTimeSeriesCoppock200ResponseMetaIndicator) *NullableGetTimeSeriesCoppock200ResponseMetaIndicator {
	return &NullableGetTimeSeriesCoppock200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCoppock200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCoppock200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
