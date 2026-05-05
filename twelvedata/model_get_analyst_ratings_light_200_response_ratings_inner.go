// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetAnalystRatingsLight200ResponseRatingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAnalystRatingsLight200ResponseRatingsInner{}

// GetAnalystRatingsLight200ResponseRatingsInner struct for GetAnalystRatingsLight200ResponseRatingsInner
type GetAnalystRatingsLight200ResponseRatingsInner struct {
	// Date when the rating was released
	Date string `json:"date"`
	// Firm that issued the ranking
	Firm string `json:"firm"`
	// Defines the action of the firm to ranking, could be `Maintains`, `Upgrade`, `Downgrade`, `Initiates` or `Reiterates`
	RatingChange *string `json:"rating_change,omitempty"`
	// Current firm's ranking of the instrument
	RatingCurrent *string `json:"rating_current,omitempty"`
	// Prior firm's ranking of the instrument
	RatingPrior *string `json:"rating_prior,omitempty"`
}

type _GetAnalystRatingsLight200ResponseRatingsInner GetAnalystRatingsLight200ResponseRatingsInner

// NewGetAnalystRatingsLight200ResponseRatingsInner instantiates a new GetAnalystRatingsLight200ResponseRatingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAnalystRatingsLight200ResponseRatingsInner(date string, firm string) *GetAnalystRatingsLight200ResponseRatingsInner {
	this := GetAnalystRatingsLight200ResponseRatingsInner{}
	this.Date = date
	this.Firm = firm
	return &this
}

// NewGetAnalystRatingsLight200ResponseRatingsInnerWithDefaults instantiates a new GetAnalystRatingsLight200ResponseRatingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAnalystRatingsLight200ResponseRatingsInnerWithDefaults() *GetAnalystRatingsLight200ResponseRatingsInner {
	this := GetAnalystRatingsLight200ResponseRatingsInner{}
	return &this
}

// GetDate returns the Date field value
func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *GetAnalystRatingsLight200ResponseRatingsInner) SetDate(v string) {
	o.Date = v
}

// GetFirm returns the Firm field value
func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetFirm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Firm
}

// GetFirmOk returns a tuple with the Firm field value
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetFirmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Firm, true
}

// SetFirm sets field value
func (o *GetAnalystRatingsLight200ResponseRatingsInner) SetFirm(v string) {
	o.Firm = v
}

// GetRatingChange returns the RatingChange field value if set, zero value otherwise.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetRatingChange() string {
	if o == nil || IsNil(o.RatingChange) {
		var ret string
		return ret
	}
	return *o.RatingChange
}

// GetRatingChangeOk returns a tuple with the RatingChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetRatingChangeOk() (*string, bool) {
	if o == nil || IsNil(o.RatingChange) {
		return nil, false
	}
	return o.RatingChange, true
}

// HasRatingChange returns a boolean if a field has been set.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) HasRatingChange() bool {
	if o != nil && !IsNil(o.RatingChange) {
		return true
	}

	return false
}

// SetRatingChange gets a reference to the given string and assigns it to the RatingChange field.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) SetRatingChange(v string) {
	o.RatingChange = &v
}

// GetRatingCurrent returns the RatingCurrent field value if set, zero value otherwise.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetRatingCurrent() string {
	if o == nil || IsNil(o.RatingCurrent) {
		var ret string
		return ret
	}
	return *o.RatingCurrent
}

// GetRatingCurrentOk returns a tuple with the RatingCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetRatingCurrentOk() (*string, bool) {
	if o == nil || IsNil(o.RatingCurrent) {
		return nil, false
	}
	return o.RatingCurrent, true
}

// HasRatingCurrent returns a boolean if a field has been set.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) HasRatingCurrent() bool {
	if o != nil && !IsNil(o.RatingCurrent) {
		return true
	}

	return false
}

// SetRatingCurrent gets a reference to the given string and assigns it to the RatingCurrent field.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) SetRatingCurrent(v string) {
	o.RatingCurrent = &v
}

// GetRatingPrior returns the RatingPrior field value if set, zero value otherwise.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetRatingPrior() string {
	if o == nil || IsNil(o.RatingPrior) {
		var ret string
		return ret
	}
	return *o.RatingPrior
}

// GetRatingPriorOk returns a tuple with the RatingPrior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetRatingPriorOk() (*string, bool) {
	if o == nil || IsNil(o.RatingPrior) {
		return nil, false
	}
	return o.RatingPrior, true
}

// HasRatingPrior returns a boolean if a field has been set.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) HasRatingPrior() bool {
	if o != nil && !IsNil(o.RatingPrior) {
		return true
	}

	return false
}

// SetRatingPrior gets a reference to the given string and assigns it to the RatingPrior field.
func (o *GetAnalystRatingsLight200ResponseRatingsInner) SetRatingPrior(v string) {
	o.RatingPrior = &v
}

func (o GetAnalystRatingsLight200ResponseRatingsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAnalystRatingsLight200ResponseRatingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["firm"] = o.Firm
	if !IsNil(o.RatingChange) {
		toSerialize["rating_change"] = o.RatingChange
	}
	if !IsNil(o.RatingCurrent) {
		toSerialize["rating_current"] = o.RatingCurrent
	}
	if !IsNil(o.RatingPrior) {
		toSerialize["rating_prior"] = o.RatingPrior
	}
	return toSerialize, nil
}

func (o *GetAnalystRatingsLight200ResponseRatingsInner) UnmarshalJSON(data []byte) (err error) {
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

	varGetAnalystRatingsLight200ResponseRatingsInner := _GetAnalystRatingsLight200ResponseRatingsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetAnalystRatingsLight200ResponseRatingsInner)

	if err != nil {
		return err
	}

	*o = GetAnalystRatingsLight200ResponseRatingsInner(varGetAnalystRatingsLight200ResponseRatingsInner)

	return err
}

type NullableGetAnalystRatingsLight200ResponseRatingsInner struct {
	value *GetAnalystRatingsLight200ResponseRatingsInner
	isSet bool
}

func (v NullableGetAnalystRatingsLight200ResponseRatingsInner) Get() *GetAnalystRatingsLight200ResponseRatingsInner {
	return v.value
}

func (v *NullableGetAnalystRatingsLight200ResponseRatingsInner) Set(val *GetAnalystRatingsLight200ResponseRatingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAnalystRatingsLight200ResponseRatingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAnalystRatingsLight200ResponseRatingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAnalystRatingsLight200ResponseRatingsInner(val *GetAnalystRatingsLight200ResponseRatingsInner) *NullableGetAnalystRatingsLight200ResponseRatingsInner {
	return &NullableGetAnalystRatingsLight200ResponseRatingsInner{value: val, isSet: true}
}

func (v NullableGetAnalystRatingsLight200ResponseRatingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAnalystRatingsLight200ResponseRatingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
