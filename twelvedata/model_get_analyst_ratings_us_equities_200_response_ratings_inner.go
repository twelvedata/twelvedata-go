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

// checks if the GetAnalystRatingsUsEquities200ResponseRatingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAnalystRatingsUsEquities200ResponseRatingsInner{}

// GetAnalystRatingsUsEquities200ResponseRatingsInner struct for GetAnalystRatingsUsEquities200ResponseRatingsInner
type GetAnalystRatingsUsEquities200ResponseRatingsInner struct {
	// Date when the rating was released
	Date string `json:"date"`
	// Firm that issued the ranking
	Firm string `json:"firm"`
	// Name of an analyst
	AnalystName *string `json:"analyst_name,omitempty"`
	// Defines the action of the firm to ranking, could be `Maintains`, `Upgrade`, `Downgrade`, `Initiates`, `Reiterates`, `Assumes`, or `Reinstates`
	RatingChange *string `json:"rating_change,omitempty"`
	// Current firm's ranking of the instrument
	RatingCurrent *string `json:"rating_current,omitempty"`
	// Prior firm's ranking of the instrument
	RatingPrior *string `json:"rating_prior,omitempty"`
	// Time when the rating was released or updated
	Time *string `json:"time,omitempty"`
	// Action that firm took towards target price
	ActionPriceTarget *string `json:"action_price_target,omitempty"`
	// Current firm's price target for the instrument
	PriceTargetCurrent *float64 `json:"price_target_current,omitempty"`
	// Prior firm's price target for the instrument
	PriceTargetPrior *float64 `json:"price_target_prior,omitempty"`
}

type _GetAnalystRatingsUsEquities200ResponseRatingsInner GetAnalystRatingsUsEquities200ResponseRatingsInner

// NewGetAnalystRatingsUsEquities200ResponseRatingsInner instantiates a new GetAnalystRatingsUsEquities200ResponseRatingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAnalystRatingsUsEquities200ResponseRatingsInner(date string, firm string) *GetAnalystRatingsUsEquities200ResponseRatingsInner {
	this := GetAnalystRatingsUsEquities200ResponseRatingsInner{}
	this.Date = date
	this.Firm = firm
	return &this
}

// NewGetAnalystRatingsUsEquities200ResponseRatingsInnerWithDefaults instantiates a new GetAnalystRatingsUsEquities200ResponseRatingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAnalystRatingsUsEquities200ResponseRatingsInnerWithDefaults() *GetAnalystRatingsUsEquities200ResponseRatingsInner {
	this := GetAnalystRatingsUsEquities200ResponseRatingsInner{}
	return &this
}

// GetDate returns the Date field value
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetDate(v string) {
	o.Date = v
}

// GetFirm returns the Firm field value
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetFirm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Firm
}

// GetFirmOk returns a tuple with the Firm field value
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetFirmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Firm, true
}

// SetFirm sets field value
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetFirm(v string) {
	o.Firm = v
}

// GetAnalystName returns the AnalystName field value if set, zero value otherwise.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetAnalystName() string {
	if o == nil || IsNil(o.AnalystName) {
		var ret string
		return ret
	}
	return *o.AnalystName
}

// GetAnalystNameOk returns a tuple with the AnalystName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetAnalystNameOk() (*string, bool) {
	if o == nil || IsNil(o.AnalystName) {
		return nil, false
	}
	return o.AnalystName, true
}

// HasAnalystName returns a boolean if a field has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasAnalystName() bool {
	if o != nil && !IsNil(o.AnalystName) {
		return true
	}

	return false
}

// SetAnalystName gets a reference to the given string and assigns it to the AnalystName field.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetAnalystName(v string) {
	o.AnalystName = &v
}

// GetRatingChange returns the RatingChange field value if set, zero value otherwise.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetRatingChange() string {
	if o == nil || IsNil(o.RatingChange) {
		var ret string
		return ret
	}
	return *o.RatingChange
}

// GetRatingChangeOk returns a tuple with the RatingChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetRatingChangeOk() (*string, bool) {
	if o == nil || IsNil(o.RatingChange) {
		return nil, false
	}
	return o.RatingChange, true
}

// HasRatingChange returns a boolean if a field has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasRatingChange() bool {
	if o != nil && !IsNil(o.RatingChange) {
		return true
	}

	return false
}

// SetRatingChange gets a reference to the given string and assigns it to the RatingChange field.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetRatingChange(v string) {
	o.RatingChange = &v
}

// GetRatingCurrent returns the RatingCurrent field value if set, zero value otherwise.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetRatingCurrent() string {
	if o == nil || IsNil(o.RatingCurrent) {
		var ret string
		return ret
	}
	return *o.RatingCurrent
}

// GetRatingCurrentOk returns a tuple with the RatingCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetRatingCurrentOk() (*string, bool) {
	if o == nil || IsNil(o.RatingCurrent) {
		return nil, false
	}
	return o.RatingCurrent, true
}

// HasRatingCurrent returns a boolean if a field has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasRatingCurrent() bool {
	if o != nil && !IsNil(o.RatingCurrent) {
		return true
	}

	return false
}

// SetRatingCurrent gets a reference to the given string and assigns it to the RatingCurrent field.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetRatingCurrent(v string) {
	o.RatingCurrent = &v
}

// GetRatingPrior returns the RatingPrior field value if set, zero value otherwise.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetRatingPrior() string {
	if o == nil || IsNil(o.RatingPrior) {
		var ret string
		return ret
	}
	return *o.RatingPrior
}

// GetRatingPriorOk returns a tuple with the RatingPrior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetRatingPriorOk() (*string, bool) {
	if o == nil || IsNil(o.RatingPrior) {
		return nil, false
	}
	return o.RatingPrior, true
}

// HasRatingPrior returns a boolean if a field has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasRatingPrior() bool {
	if o != nil && !IsNil(o.RatingPrior) {
		return true
	}

	return false
}

// SetRatingPrior gets a reference to the given string and assigns it to the RatingPrior field.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetRatingPrior(v string) {
	o.RatingPrior = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetTime(v string) {
	o.Time = &v
}

// GetActionPriceTarget returns the ActionPriceTarget field value if set, zero value otherwise.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetActionPriceTarget() string {
	if o == nil || IsNil(o.ActionPriceTarget) {
		var ret string
		return ret
	}
	return *o.ActionPriceTarget
}

// GetActionPriceTargetOk returns a tuple with the ActionPriceTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetActionPriceTargetOk() (*string, bool) {
	if o == nil || IsNil(o.ActionPriceTarget) {
		return nil, false
	}
	return o.ActionPriceTarget, true
}

// HasActionPriceTarget returns a boolean if a field has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasActionPriceTarget() bool {
	if o != nil && !IsNil(o.ActionPriceTarget) {
		return true
	}

	return false
}

// SetActionPriceTarget gets a reference to the given string and assigns it to the ActionPriceTarget field.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetActionPriceTarget(v string) {
	o.ActionPriceTarget = &v
}

// GetPriceTargetCurrent returns the PriceTargetCurrent field value if set, zero value otherwise.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetPriceTargetCurrent() float64 {
	if o == nil || IsNil(o.PriceTargetCurrent) {
		var ret float64
		return ret
	}
	return *o.PriceTargetCurrent
}

// GetPriceTargetCurrentOk returns a tuple with the PriceTargetCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetPriceTargetCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceTargetCurrent) {
		return nil, false
	}
	return o.PriceTargetCurrent, true
}

// HasPriceTargetCurrent returns a boolean if a field has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasPriceTargetCurrent() bool {
	if o != nil && !IsNil(o.PriceTargetCurrent) {
		return true
	}

	return false
}

// SetPriceTargetCurrent gets a reference to the given float64 and assigns it to the PriceTargetCurrent field.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetPriceTargetCurrent(v float64) {
	o.PriceTargetCurrent = &v
}

// GetPriceTargetPrior returns the PriceTargetPrior field value if set, zero value otherwise.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetPriceTargetPrior() float64 {
	if o == nil || IsNil(o.PriceTargetPrior) {
		var ret float64
		return ret
	}
	return *o.PriceTargetPrior
}

// GetPriceTargetPriorOk returns a tuple with the PriceTargetPrior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetPriceTargetPriorOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceTargetPrior) {
		return nil, false
	}
	return o.PriceTargetPrior, true
}

// HasPriceTargetPrior returns a boolean if a field has been set.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasPriceTargetPrior() bool {
	if o != nil && !IsNil(o.PriceTargetPrior) {
		return true
	}

	return false
}

// SetPriceTargetPrior gets a reference to the given float64 and assigns it to the PriceTargetPrior field.
func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetPriceTargetPrior(v float64) {
	o.PriceTargetPrior = &v
}

func (o GetAnalystRatingsUsEquities200ResponseRatingsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAnalystRatingsUsEquities200ResponseRatingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["firm"] = o.Firm
	if !IsNil(o.AnalystName) {
		toSerialize["analyst_name"] = o.AnalystName
	}
	if !IsNil(o.RatingChange) {
		toSerialize["rating_change"] = o.RatingChange
	}
	if !IsNil(o.RatingCurrent) {
		toSerialize["rating_current"] = o.RatingCurrent
	}
	if !IsNil(o.RatingPrior) {
		toSerialize["rating_prior"] = o.RatingPrior
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.ActionPriceTarget) {
		toSerialize["action_price_target"] = o.ActionPriceTarget
	}
	if !IsNil(o.PriceTargetCurrent) {
		toSerialize["price_target_current"] = o.PriceTargetCurrent
	}
	if !IsNil(o.PriceTargetPrior) {
		toSerialize["price_target_prior"] = o.PriceTargetPrior
	}
	return toSerialize, nil
}

func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"firm",
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

	varGetAnalystRatingsUsEquities200ResponseRatingsInner := _GetAnalystRatingsUsEquities200ResponseRatingsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAnalystRatingsUsEquities200ResponseRatingsInner)

	if err != nil {
		return err
	}

	*o = GetAnalystRatingsUsEquities200ResponseRatingsInner(varGetAnalystRatingsUsEquities200ResponseRatingsInner)

	return err
}

type NullableGetAnalystRatingsUsEquities200ResponseRatingsInner struct {
	value *GetAnalystRatingsUsEquities200ResponseRatingsInner
	isSet bool
}

func (v NullableGetAnalystRatingsUsEquities200ResponseRatingsInner) Get() *GetAnalystRatingsUsEquities200ResponseRatingsInner {
	return v.value
}

func (v *NullableGetAnalystRatingsUsEquities200ResponseRatingsInner) Set(val *GetAnalystRatingsUsEquities200ResponseRatingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAnalystRatingsUsEquities200ResponseRatingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAnalystRatingsUsEquities200ResponseRatingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAnalystRatingsUsEquities200ResponseRatingsInner(val *GetAnalystRatingsUsEquities200ResponseRatingsInner) *NullableGetAnalystRatingsUsEquities200ResponseRatingsInner {
	return &NullableGetAnalystRatingsUsEquities200ResponseRatingsInner{value: val, isSet: true}
}

func (v NullableGetAnalystRatingsUsEquities200ResponseRatingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAnalystRatingsUsEquities200ResponseRatingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
