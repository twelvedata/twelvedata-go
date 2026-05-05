// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesCrsi200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCrsi200ResponseMetaIndicator{}

// GetTimeSeriesCrsi200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesCrsi200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// Number of periods for RSI used to calculate price momentum
	RsiPeriod int64 `json:"rsi_period"`
	// Number of periods for RSI used to calculate up/down trend
	UpDownLength int64 `json:"up_down_length"`
	// Number of periods used to calculate PercentRank
	PercentRankPeriod int64 `json:"percent_rank_period"`
}

type _GetTimeSeriesCrsi200ResponseMetaIndicator GetTimeSeriesCrsi200ResponseMetaIndicator

// NewGetTimeSeriesCrsi200ResponseMetaIndicator instantiates a new GetTimeSeriesCrsi200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCrsi200ResponseMetaIndicator(name string, seriesType string, rsiPeriod int64, upDownLength int64, percentRankPeriod int64) *GetTimeSeriesCrsi200ResponseMetaIndicator {
	this := GetTimeSeriesCrsi200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.RsiPeriod = rsiPeriod
	this.UpDownLength = upDownLength
	this.PercentRankPeriod = percentRankPeriod
	return &this
}

// NewGetTimeSeriesCrsi200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesCrsi200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCrsi200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesCrsi200ResponseMetaIndicator {
	this := GetTimeSeriesCrsi200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetRsiPeriod returns the RsiPeriod field value
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetRsiPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RsiPeriod
}

// GetRsiPeriodOk returns a tuple with the RsiPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetRsiPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RsiPeriod, true
}

// SetRsiPeriod sets field value
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) SetRsiPeriod(v int64) {
	o.RsiPeriod = v
}

// GetUpDownLength returns the UpDownLength field value
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetUpDownLength() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpDownLength
}

// GetUpDownLengthOk returns a tuple with the UpDownLength field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetUpDownLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpDownLength, true
}

// SetUpDownLength sets field value
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) SetUpDownLength(v int64) {
	o.UpDownLength = v
}

// GetPercentRankPeriod returns the PercentRankPeriod field value
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetPercentRankPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PercentRankPeriod
}

// GetPercentRankPeriodOk returns a tuple with the PercentRankPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetPercentRankPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentRankPeriod, true
}

// SetPercentRankPeriod sets field value
func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) SetPercentRankPeriod(v int64) {
	o.PercentRankPeriod = v
}

func (o GetTimeSeriesCrsi200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCrsi200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["rsi_period"] = o.RsiPeriod
	toSerialize["up_down_length"] = o.UpDownLength
	toSerialize["percent_rank_period"] = o.PercentRankPeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
		"rsi_period",
		"up_down_length",
		"percent_rank_period",
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

	varGetTimeSeriesCrsi200ResponseMetaIndicator := _GetTimeSeriesCrsi200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesCrsi200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCrsi200ResponseMetaIndicator(varGetTimeSeriesCrsi200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesCrsi200ResponseMetaIndicator struct {
	value *GetTimeSeriesCrsi200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesCrsi200ResponseMetaIndicator) Get() *GetTimeSeriesCrsi200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesCrsi200ResponseMetaIndicator) Set(val *GetTimeSeriesCrsi200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCrsi200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCrsi200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCrsi200ResponseMetaIndicator(val *GetTimeSeriesCrsi200ResponseMetaIndicator) *NullableGetTimeSeriesCrsi200ResponseMetaIndicator {
	return &NullableGetTimeSeriesCrsi200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCrsi200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCrsi200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
