// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetEarningsCalendar200ResponseEarningsValueInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEarningsCalendar200ResponseEarningsValueInner{}

// GetEarningsCalendar200ResponseEarningsValueInner struct for GetEarningsCalendar200ResponseEarningsValueInner
type GetEarningsCalendar200ResponseEarningsValueInner struct {
	// Instrument symbol (ticker)
	Symbol *string `json:"symbol,omitempty"`
	// Full name of instrument
	Name *string `json:"name,omitempty"`
	// Currency in which instrument is traded by ISO 4217 standard
	Currency *string `json:"currency,omitempty"`
	// Exchange where instrument is traded
	Exchange *string `json:"exchange,omitempty"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode *string `json:"mic_code,omitempty"`
	// Country where exchange is located
	Country *string `json:"country,omitempty"`
	// Can be either of the following values: `Pre Market`, `After Hours`, `Time Not Supplied`
	Time *string `json:"time,omitempty"`
	// Analyst estimate of the future company earning
	EpsEstimate *float64 `json:"eps_estimate,omitempty"`
	// Actual value of reported earning
	EpsActual *float64 `json:"eps_actual,omitempty"`
	// Delta between `eps_actual` and `eps_estimate`
	Difference *float64 `json:"difference,omitempty"`
	// Surprise in percentage of the `eps_actual` related to `eps_estimate`
	SurprisePrc *float64 `json:"surprise_prc,omitempty"`
}

// NewGetEarningsCalendar200ResponseEarningsValueInner instantiates a new GetEarningsCalendar200ResponseEarningsValueInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEarningsCalendar200ResponseEarningsValueInner() *GetEarningsCalendar200ResponseEarningsValueInner {
	this := GetEarningsCalendar200ResponseEarningsValueInner{}
	return &this
}

// NewGetEarningsCalendar200ResponseEarningsValueInnerWithDefaults instantiates a new GetEarningsCalendar200ResponseEarningsValueInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEarningsCalendar200ResponseEarningsValueInnerWithDefaults() *GetEarningsCalendar200ResponseEarningsValueInner {
	this := GetEarningsCalendar200ResponseEarningsValueInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetName(v string) {
	o.Name = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetExchange(v string) {
	o.Exchange = &v
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetMicCode(v string) {
	o.MicCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetCountry(v string) {
	o.Country = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetTime(v string) {
	o.Time = &v
}

// GetEpsEstimate returns the EpsEstimate field value if set, zero value otherwise.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetEpsEstimate() float64 {
	if o == nil || IsNil(o.EpsEstimate) {
		var ret float64
		return ret
	}
	return *o.EpsEstimate
}

// GetEpsEstimateOk returns a tuple with the EpsEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetEpsEstimateOk() (*float64, bool) {
	if o == nil || IsNil(o.EpsEstimate) {
		return nil, false
	}
	return o.EpsEstimate, true
}

// HasEpsEstimate returns a boolean if a field has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasEpsEstimate() bool {
	if o != nil && !IsNil(o.EpsEstimate) {
		return true
	}

	return false
}

// SetEpsEstimate gets a reference to the given float64 and assigns it to the EpsEstimate field.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetEpsEstimate(v float64) {
	o.EpsEstimate = &v
}

// GetEpsActual returns the EpsActual field value if set, zero value otherwise.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetEpsActual() float64 {
	if o == nil || IsNil(o.EpsActual) {
		var ret float64
		return ret
	}
	return *o.EpsActual
}

// GetEpsActualOk returns a tuple with the EpsActual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetEpsActualOk() (*float64, bool) {
	if o == nil || IsNil(o.EpsActual) {
		return nil, false
	}
	return o.EpsActual, true
}

// HasEpsActual returns a boolean if a field has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasEpsActual() bool {
	if o != nil && !IsNil(o.EpsActual) {
		return true
	}

	return false
}

// SetEpsActual gets a reference to the given float64 and assigns it to the EpsActual field.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetEpsActual(v float64) {
	o.EpsActual = &v
}

// GetDifference returns the Difference field value if set, zero value otherwise.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetDifference() float64 {
	if o == nil || IsNil(o.Difference) {
		var ret float64
		return ret
	}
	return *o.Difference
}

// GetDifferenceOk returns a tuple with the Difference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetDifferenceOk() (*float64, bool) {
	if o == nil || IsNil(o.Difference) {
		return nil, false
	}
	return o.Difference, true
}

// HasDifference returns a boolean if a field has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasDifference() bool {
	if o != nil && !IsNil(o.Difference) {
		return true
	}

	return false
}

// SetDifference gets a reference to the given float64 and assigns it to the Difference field.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetDifference(v float64) {
	o.Difference = &v
}

// GetSurprisePrc returns the SurprisePrc field value if set, zero value otherwise.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetSurprisePrc() float64 {
	if o == nil || IsNil(o.SurprisePrc) {
		var ret float64
		return ret
	}
	return *o.SurprisePrc
}

// GetSurprisePrcOk returns a tuple with the SurprisePrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetSurprisePrcOk() (*float64, bool) {
	if o == nil || IsNil(o.SurprisePrc) {
		return nil, false
	}
	return o.SurprisePrc, true
}

// HasSurprisePrc returns a boolean if a field has been set.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasSurprisePrc() bool {
	if o != nil && !IsNil(o.SurprisePrc) {
		return true
	}

	return false
}

// SetSurprisePrc gets a reference to the given float64 and assigns it to the SurprisePrc field.
func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetSurprisePrc(v float64) {
	o.SurprisePrc = &v
}

func (o GetEarningsCalendar200ResponseEarningsValueInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEarningsCalendar200ResponseEarningsValueInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.MicCode) {
		toSerialize["mic_code"] = o.MicCode
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.EpsEstimate) {
		toSerialize["eps_estimate"] = o.EpsEstimate
	}
	if !IsNil(o.EpsActual) {
		toSerialize["eps_actual"] = o.EpsActual
	}
	if !IsNil(o.Difference) {
		toSerialize["difference"] = o.Difference
	}
	if !IsNil(o.SurprisePrc) {
		toSerialize["surprise_prc"] = o.SurprisePrc
	}
	return toSerialize, nil
}

type NullableGetEarningsCalendar200ResponseEarningsValueInner struct {
	value *GetEarningsCalendar200ResponseEarningsValueInner
	isSet bool
}

func (v NullableGetEarningsCalendar200ResponseEarningsValueInner) Get() *GetEarningsCalendar200ResponseEarningsValueInner {
	return v.value
}

func (v *NullableGetEarningsCalendar200ResponseEarningsValueInner) Set(val *GetEarningsCalendar200ResponseEarningsValueInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEarningsCalendar200ResponseEarningsValueInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEarningsCalendar200ResponseEarningsValueInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEarningsCalendar200ResponseEarningsValueInner(val *GetEarningsCalendar200ResponseEarningsValueInner) *NullableGetEarningsCalendar200ResponseEarningsValueInner {
	return &NullableGetEarningsCalendar200ResponseEarningsValueInner{value: val, isSet: true}
}

func (v NullableGetEarningsCalendar200ResponseEarningsValueInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEarningsCalendar200ResponseEarningsValueInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
