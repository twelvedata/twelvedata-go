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

// checks if the GetTimeSeriesIchimoku200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesIchimoku200ResponseValuesInner{}

// GetTimeSeriesIchimoku200ResponseValuesInner struct for GetTimeSeriesIchimoku200ResponseValuesInner
type GetTimeSeriesIchimoku200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Tenkan-sen value
	TenkanSen *string `json:"tenkan_sen,omitempty"`
	// Kijun-sen value
	KijunSen *string `json:"kijun_sen,omitempty"`
	// Senkou span A value
	SenkouSpanA string `json:"senkou_span_a"`
	// Senkou span B value
	SenkouSpanB string `json:"senkou_span_b"`
	// Chikou span value
	ChikouSpan *string `json:"chikou_span,omitempty"`
}

type _GetTimeSeriesIchimoku200ResponseValuesInner GetTimeSeriesIchimoku200ResponseValuesInner

// NewGetTimeSeriesIchimoku200ResponseValuesInner instantiates a new GetTimeSeriesIchimoku200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesIchimoku200ResponseValuesInner(datetime string, senkouSpanA string, senkouSpanB string) *GetTimeSeriesIchimoku200ResponseValuesInner {
	this := GetTimeSeriesIchimoku200ResponseValuesInner{}
	this.Datetime = datetime
	this.SenkouSpanA = senkouSpanA
	this.SenkouSpanB = senkouSpanB
	return &this
}

// NewGetTimeSeriesIchimoku200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesIchimoku200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesIchimoku200ResponseValuesInnerWithDefaults() *GetTimeSeriesIchimoku200ResponseValuesInner {
	this := GetTimeSeriesIchimoku200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetTenkanSen returns the TenkanSen field value if set, zero value otherwise.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetTenkanSen() string {
	if o == nil || IsNil(o.TenkanSen) {
		var ret string
		return ret
	}
	return *o.TenkanSen
}

// GetTenkanSenOk returns a tuple with the TenkanSen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetTenkanSenOk() (*string, bool) {
	if o == nil || IsNil(o.TenkanSen) {
		return nil, false
	}
	return o.TenkanSen, true
}

// HasTenkanSen returns a boolean if a field has been set.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) HasTenkanSen() bool {
	if o != nil && !IsNil(o.TenkanSen) {
		return true
	}

	return false
}

// SetTenkanSen gets a reference to the given string and assigns it to the TenkanSen field.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) SetTenkanSen(v string) {
	o.TenkanSen = &v
}

// GetKijunSen returns the KijunSen field value if set, zero value otherwise.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetKijunSen() string {
	if o == nil || IsNil(o.KijunSen) {
		var ret string
		return ret
	}
	return *o.KijunSen
}

// GetKijunSenOk returns a tuple with the KijunSen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetKijunSenOk() (*string, bool) {
	if o == nil || IsNil(o.KijunSen) {
		return nil, false
	}
	return o.KijunSen, true
}

// HasKijunSen returns a boolean if a field has been set.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) HasKijunSen() bool {
	if o != nil && !IsNil(o.KijunSen) {
		return true
	}

	return false
}

// SetKijunSen gets a reference to the given string and assigns it to the KijunSen field.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) SetKijunSen(v string) {
	o.KijunSen = &v
}

// GetSenkouSpanA returns the SenkouSpanA field value
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetSenkouSpanA() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenkouSpanA
}

// GetSenkouSpanAOk returns a tuple with the SenkouSpanA field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetSenkouSpanAOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenkouSpanA, true
}

// SetSenkouSpanA sets field value
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) SetSenkouSpanA(v string) {
	o.SenkouSpanA = v
}

// GetSenkouSpanB returns the SenkouSpanB field value
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetSenkouSpanB() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenkouSpanB
}

// GetSenkouSpanBOk returns a tuple with the SenkouSpanB field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetSenkouSpanBOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenkouSpanB, true
}

// SetSenkouSpanB sets field value
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) SetSenkouSpanB(v string) {
	o.SenkouSpanB = v
}

// GetChikouSpan returns the ChikouSpan field value if set, zero value otherwise.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetChikouSpan() string {
	if o == nil || IsNil(o.ChikouSpan) {
		var ret string
		return ret
	}
	return *o.ChikouSpan
}

// GetChikouSpanOk returns a tuple with the ChikouSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetChikouSpanOk() (*string, bool) {
	if o == nil || IsNil(o.ChikouSpan) {
		return nil, false
	}
	return o.ChikouSpan, true
}

// HasChikouSpan returns a boolean if a field has been set.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) HasChikouSpan() bool {
	if o != nil && !IsNil(o.ChikouSpan) {
		return true
	}

	return false
}

// SetChikouSpan gets a reference to the given string and assigns it to the ChikouSpan field.
func (o *GetTimeSeriesIchimoku200ResponseValuesInner) SetChikouSpan(v string) {
	o.ChikouSpan = &v
}

func (o GetTimeSeriesIchimoku200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesIchimoku200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	if !IsNil(o.TenkanSen) {
		toSerialize["tenkan_sen"] = o.TenkanSen
	}
	if !IsNil(o.KijunSen) {
		toSerialize["kijun_sen"] = o.KijunSen
	}
	toSerialize["senkou_span_a"] = o.SenkouSpanA
	toSerialize["senkou_span_b"] = o.SenkouSpanB
	if !IsNil(o.ChikouSpan) {
		toSerialize["chikou_span"] = o.ChikouSpan
	}
	return toSerialize, nil
}

func (o *GetTimeSeriesIchimoku200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"senkou_span_a",
		"senkou_span_b",
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

	varGetTimeSeriesIchimoku200ResponseValuesInner := _GetTimeSeriesIchimoku200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesIchimoku200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesIchimoku200ResponseValuesInner(varGetTimeSeriesIchimoku200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesIchimoku200ResponseValuesInner struct {
	value *GetTimeSeriesIchimoku200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesIchimoku200ResponseValuesInner) Get() *GetTimeSeriesIchimoku200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesIchimoku200ResponseValuesInner) Set(val *GetTimeSeriesIchimoku200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesIchimoku200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesIchimoku200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesIchimoku200ResponseValuesInner(val *GetTimeSeriesIchimoku200ResponseValuesInner) *NullableGetTimeSeriesIchimoku200ResponseValuesInner {
	return &NullableGetTimeSeriesIchimoku200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesIchimoku200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesIchimoku200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
