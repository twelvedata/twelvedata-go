// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtf{}

// GetETFsWorld200ResponseEtf Etf information
type GetETFsWorld200ResponseEtf struct {
	Summary     *GetETFsWorld200ResponseEtfSummary     `json:"summary,omitempty"`
	Performance *GetETFsWorld200ResponseEtfPerformance `json:"performance,omitempty"`
	Risk        *GetETFsWorld200ResponseEtfRisk        `json:"risk,omitempty"`
	Composition *GetETFsWorld200ResponseEtfComposition `json:"composition,omitempty"`
}

// NewGetETFsWorld200ResponseEtf instantiates a new GetETFsWorld200ResponseEtf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtf() *GetETFsWorld200ResponseEtf {
	this := GetETFsWorld200ResponseEtf{}
	return &this
}

// NewGetETFsWorld200ResponseEtfWithDefaults instantiates a new GetETFsWorld200ResponseEtf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfWithDefaults() *GetETFsWorld200ResponseEtf {
	this := GetETFsWorld200ResponseEtf{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtf) GetSummary() GetETFsWorld200ResponseEtfSummary {
	if o == nil || IsNil(o.Summary) {
		var ret GetETFsWorld200ResponseEtfSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtf) GetSummaryOk() (*GetETFsWorld200ResponseEtfSummary, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtf) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given GetETFsWorld200ResponseEtfSummary and assigns it to the Summary field.
func (o *GetETFsWorld200ResponseEtf) SetSummary(v GetETFsWorld200ResponseEtfSummary) {
	o.Summary = &v
}

// GetPerformance returns the Performance field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtf) GetPerformance() GetETFsWorld200ResponseEtfPerformance {
	if o == nil || IsNil(o.Performance) {
		var ret GetETFsWorld200ResponseEtfPerformance
		return ret
	}
	return *o.Performance
}

// GetPerformanceOk returns a tuple with the Performance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtf) GetPerformanceOk() (*GetETFsWorld200ResponseEtfPerformance, bool) {
	if o == nil || IsNil(o.Performance) {
		return nil, false
	}
	return o.Performance, true
}

// HasPerformance returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtf) HasPerformance() bool {
	if o != nil && !IsNil(o.Performance) {
		return true
	}

	return false
}

// SetPerformance gets a reference to the given GetETFsWorld200ResponseEtfPerformance and assigns it to the Performance field.
func (o *GetETFsWorld200ResponseEtf) SetPerformance(v GetETFsWorld200ResponseEtfPerformance) {
	o.Performance = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtf) GetRisk() GetETFsWorld200ResponseEtfRisk {
	if o == nil || IsNil(o.Risk) {
		var ret GetETFsWorld200ResponseEtfRisk
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtf) GetRiskOk() (*GetETFsWorld200ResponseEtfRisk, bool) {
	if o == nil || IsNil(o.Risk) {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtf) HasRisk() bool {
	if o != nil && !IsNil(o.Risk) {
		return true
	}

	return false
}

// SetRisk gets a reference to the given GetETFsWorld200ResponseEtfRisk and assigns it to the Risk field.
func (o *GetETFsWorld200ResponseEtf) SetRisk(v GetETFsWorld200ResponseEtfRisk) {
	o.Risk = &v
}

// GetComposition returns the Composition field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtf) GetComposition() GetETFsWorld200ResponseEtfComposition {
	if o == nil || IsNil(o.Composition) {
		var ret GetETFsWorld200ResponseEtfComposition
		return ret
	}
	return *o.Composition
}

// GetCompositionOk returns a tuple with the Composition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtf) GetCompositionOk() (*GetETFsWorld200ResponseEtfComposition, bool) {
	if o == nil || IsNil(o.Composition) {
		return nil, false
	}
	return o.Composition, true
}

// HasComposition returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtf) HasComposition() bool {
	if o != nil && !IsNil(o.Composition) {
		return true
	}

	return false
}

// SetComposition gets a reference to the given GetETFsWorld200ResponseEtfComposition and assigns it to the Composition field.
func (o *GetETFsWorld200ResponseEtf) SetComposition(v GetETFsWorld200ResponseEtfComposition) {
	o.Composition = &v
}

func (o GetETFsWorld200ResponseEtf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Performance) {
		toSerialize["performance"] = o.Performance
	}
	if !IsNil(o.Risk) {
		toSerialize["risk"] = o.Risk
	}
	if !IsNil(o.Composition) {
		toSerialize["composition"] = o.Composition
	}
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtf struct {
	value *GetETFsWorld200ResponseEtf
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtf) Get() *GetETFsWorld200ResponseEtf {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtf) Set(val *GetETFsWorld200ResponseEtf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtf(val *GetETFsWorld200ResponseEtf) *NullableGetETFsWorld200ResponseEtf {
	return &NullableGetETFsWorld200ResponseEtf{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
