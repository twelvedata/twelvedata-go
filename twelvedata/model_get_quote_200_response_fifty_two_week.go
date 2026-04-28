/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetQuote200ResponseFiftyTwoWeek type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetQuote200ResponseFiftyTwoWeek{}

// GetQuote200ResponseFiftyTwoWeek Collection of 52-week metrics
type GetQuote200ResponseFiftyTwoWeek struct {
	// 52-week low price
	Low *string `json:"low,omitempty"`
	// 52-week high price
	High *string `json:"high,omitempty"`
	// Current price - 52-week low
	LowChange *string `json:"low_change,omitempty"`
	// Current price - 52-week high
	HighChange *string `json:"high_change,omitempty"`
	// Percentage change from 52-week low
	LowChangePercent *string `json:"low_change_percent,omitempty"`
	// Percentage change from 52-week high
	HighChangePercent *string `json:"high_change_percent,omitempty"`
	// Range between 52-week low and high
	Range *string `json:"range,omitempty"`
}

// NewGetQuote200ResponseFiftyTwoWeek instantiates a new GetQuote200ResponseFiftyTwoWeek object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetQuote200ResponseFiftyTwoWeek() *GetQuote200ResponseFiftyTwoWeek {
	this := GetQuote200ResponseFiftyTwoWeek{}
	return &this
}

// NewGetQuote200ResponseFiftyTwoWeekWithDefaults instantiates a new GetQuote200ResponseFiftyTwoWeek object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetQuote200ResponseFiftyTwoWeekWithDefaults() *GetQuote200ResponseFiftyTwoWeek {
	this := GetQuote200ResponseFiftyTwoWeek{}
	return &this
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *GetQuote200ResponseFiftyTwoWeek) GetLow() string {
	if o == nil || IsNil(o.Low) {
		var ret string
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) GetLowOk() (*string, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given string and assigns it to the Low field.
func (o *GetQuote200ResponseFiftyTwoWeek) SetLow(v string) {
	o.Low = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *GetQuote200ResponseFiftyTwoWeek) GetHigh() string {
	if o == nil || IsNil(o.High) {
		var ret string
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) GetHighOk() (*string, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given string and assigns it to the High field.
func (o *GetQuote200ResponseFiftyTwoWeek) SetHigh(v string) {
	o.High = &v
}

// GetLowChange returns the LowChange field value if set, zero value otherwise.
func (o *GetQuote200ResponseFiftyTwoWeek) GetLowChange() string {
	if o == nil || IsNil(o.LowChange) {
		var ret string
		return ret
	}
	return *o.LowChange
}

// GetLowChangeOk returns a tuple with the LowChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) GetLowChangeOk() (*string, bool) {
	if o == nil || IsNil(o.LowChange) {
		return nil, false
	}
	return o.LowChange, true
}

// HasLowChange returns a boolean if a field has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) HasLowChange() bool {
	if o != nil && !IsNil(o.LowChange) {
		return true
	}

	return false
}

// SetLowChange gets a reference to the given string and assigns it to the LowChange field.
func (o *GetQuote200ResponseFiftyTwoWeek) SetLowChange(v string) {
	o.LowChange = &v
}

// GetHighChange returns the HighChange field value if set, zero value otherwise.
func (o *GetQuote200ResponseFiftyTwoWeek) GetHighChange() string {
	if o == nil || IsNil(o.HighChange) {
		var ret string
		return ret
	}
	return *o.HighChange
}

// GetHighChangeOk returns a tuple with the HighChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) GetHighChangeOk() (*string, bool) {
	if o == nil || IsNil(o.HighChange) {
		return nil, false
	}
	return o.HighChange, true
}

// HasHighChange returns a boolean if a field has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) HasHighChange() bool {
	if o != nil && !IsNil(o.HighChange) {
		return true
	}

	return false
}

// SetHighChange gets a reference to the given string and assigns it to the HighChange field.
func (o *GetQuote200ResponseFiftyTwoWeek) SetHighChange(v string) {
	o.HighChange = &v
}

// GetLowChangePercent returns the LowChangePercent field value if set, zero value otherwise.
func (o *GetQuote200ResponseFiftyTwoWeek) GetLowChangePercent() string {
	if o == nil || IsNil(o.LowChangePercent) {
		var ret string
		return ret
	}
	return *o.LowChangePercent
}

// GetLowChangePercentOk returns a tuple with the LowChangePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) GetLowChangePercentOk() (*string, bool) {
	if o == nil || IsNil(o.LowChangePercent) {
		return nil, false
	}
	return o.LowChangePercent, true
}

// HasLowChangePercent returns a boolean if a field has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) HasLowChangePercent() bool {
	if o != nil && !IsNil(o.LowChangePercent) {
		return true
	}

	return false
}

// SetLowChangePercent gets a reference to the given string and assigns it to the LowChangePercent field.
func (o *GetQuote200ResponseFiftyTwoWeek) SetLowChangePercent(v string) {
	o.LowChangePercent = &v
}

// GetHighChangePercent returns the HighChangePercent field value if set, zero value otherwise.
func (o *GetQuote200ResponseFiftyTwoWeek) GetHighChangePercent() string {
	if o == nil || IsNil(o.HighChangePercent) {
		var ret string
		return ret
	}
	return *o.HighChangePercent
}

// GetHighChangePercentOk returns a tuple with the HighChangePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) GetHighChangePercentOk() (*string, bool) {
	if o == nil || IsNil(o.HighChangePercent) {
		return nil, false
	}
	return o.HighChangePercent, true
}

// HasHighChangePercent returns a boolean if a field has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) HasHighChangePercent() bool {
	if o != nil && !IsNil(o.HighChangePercent) {
		return true
	}

	return false
}

// SetHighChangePercent gets a reference to the given string and assigns it to the HighChangePercent field.
func (o *GetQuote200ResponseFiftyTwoWeek) SetHighChangePercent(v string) {
	o.HighChangePercent = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *GetQuote200ResponseFiftyTwoWeek) GetRange() string {
	if o == nil || IsNil(o.Range) {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) GetRangeOk() (*string, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *GetQuote200ResponseFiftyTwoWeek) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *GetQuote200ResponseFiftyTwoWeek) SetRange(v string) {
	o.Range = &v
}

func (o GetQuote200ResponseFiftyTwoWeek) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetQuote200ResponseFiftyTwoWeek) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !IsNil(o.LowChange) {
		toSerialize["low_change"] = o.LowChange
	}
	if !IsNil(o.HighChange) {
		toSerialize["high_change"] = o.HighChange
	}
	if !IsNil(o.LowChangePercent) {
		toSerialize["low_change_percent"] = o.LowChangePercent
	}
	if !IsNil(o.HighChangePercent) {
		toSerialize["high_change_percent"] = o.HighChangePercent
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	return toSerialize, nil
}

type NullableGetQuote200ResponseFiftyTwoWeek struct {
	value *GetQuote200ResponseFiftyTwoWeek
	isSet bool
}

func (v NullableGetQuote200ResponseFiftyTwoWeek) Get() *GetQuote200ResponseFiftyTwoWeek {
	return v.value
}

func (v *NullableGetQuote200ResponseFiftyTwoWeek) Set(val *GetQuote200ResponseFiftyTwoWeek) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuote200ResponseFiftyTwoWeek) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuote200ResponseFiftyTwoWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuote200ResponseFiftyTwoWeek(val *GetQuote200ResponseFiftyTwoWeek) *NullableGetQuote200ResponseFiftyTwoWeek {
	return &NullableGetQuote200ResponseFiftyTwoWeek{value: val, isSet: true}
}

func (v NullableGetQuote200ResponseFiftyTwoWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuote200ResponseFiftyTwoWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
