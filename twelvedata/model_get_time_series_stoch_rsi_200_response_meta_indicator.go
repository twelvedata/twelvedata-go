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

// checks if the GetTimeSeriesStochRsi200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesStochRsi200ResponseMetaIndicator{}

// GetTimeSeriesStochRsi200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesStochRsi200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// Length of period for calculating the RSI component
	RsiLength int64 `json:"rsi_length"`
	// Period length for computing the stochastic oscillator of the RSI
	StochLength int64 `json:"stoch_length"`
	// Period for smoothing the %K line
	KPeriod int64 `json:"k_period"`
	// Period for smoothing the %D line, which is a moving average of %K
	DPeriod int64 `json:"d_period"`
}

type _GetTimeSeriesStochRsi200ResponseMetaIndicator GetTimeSeriesStochRsi200ResponseMetaIndicator

// NewGetTimeSeriesStochRsi200ResponseMetaIndicator instantiates a new GetTimeSeriesStochRsi200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesStochRsi200ResponseMetaIndicator(name string, seriesType string, rsiLength int64, stochLength int64, kPeriod int64, dPeriod int64) *GetTimeSeriesStochRsi200ResponseMetaIndicator {
	this := GetTimeSeriesStochRsi200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.RsiLength = rsiLength
	this.StochLength = stochLength
	this.KPeriod = kPeriod
	this.DPeriod = dPeriod
	return &this
}

// NewGetTimeSeriesStochRsi200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesStochRsi200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesStochRsi200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesStochRsi200ResponseMetaIndicator {
	this := GetTimeSeriesStochRsi200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetRsiLength returns the RsiLength field value
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetRsiLength() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RsiLength
}

// GetRsiLengthOk returns a tuple with the RsiLength field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetRsiLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RsiLength, true
}

// SetRsiLength sets field value
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) SetRsiLength(v int64) {
	o.RsiLength = v
}

// GetStochLength returns the StochLength field value
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetStochLength() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StochLength
}

// GetStochLengthOk returns a tuple with the StochLength field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetStochLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StochLength, true
}

// SetStochLength sets field value
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) SetStochLength(v int64) {
	o.StochLength = v
}

// GetKPeriod returns the KPeriod field value
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetKPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.KPeriod
}

// GetKPeriodOk returns a tuple with the KPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetKPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KPeriod, true
}

// SetKPeriod sets field value
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) SetKPeriod(v int64) {
	o.KPeriod = v
}

// GetDPeriod returns the DPeriod field value
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetDPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DPeriod
}

// GetDPeriodOk returns a tuple with the DPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetDPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DPeriod, true
}

// SetDPeriod sets field value
func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) SetDPeriod(v int64) {
	o.DPeriod = v
}

func (o GetTimeSeriesStochRsi200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesStochRsi200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["rsi_length"] = o.RsiLength
	toSerialize["stoch_length"] = o.StochLength
	toSerialize["k_period"] = o.KPeriod
	toSerialize["d_period"] = o.DPeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
		"rsi_length",
		"stoch_length",
		"k_period",
		"d_period",
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

	varGetTimeSeriesStochRsi200ResponseMetaIndicator := _GetTimeSeriesStochRsi200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesStochRsi200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesStochRsi200ResponseMetaIndicator(varGetTimeSeriesStochRsi200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesStochRsi200ResponseMetaIndicator struct {
	value *GetTimeSeriesStochRsi200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesStochRsi200ResponseMetaIndicator) Get() *GetTimeSeriesStochRsi200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesStochRsi200ResponseMetaIndicator) Set(val *GetTimeSeriesStochRsi200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesStochRsi200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesStochRsi200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesStochRsi200ResponseMetaIndicator(val *GetTimeSeriesStochRsi200ResponseMetaIndicator) *NullableGetTimeSeriesStochRsi200ResponseMetaIndicator {
	return &NullableGetTimeSeriesStochRsi200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesStochRsi200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesStochRsi200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
