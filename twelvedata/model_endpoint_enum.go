/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
	"fmt"
)

// EndpointEnum the model 'EndpointEnum'
type EndpointEnum string

// List of EndpointEnum
const (
	ENDPOINTENUM_PRICE_TARGET                      EndpointEnum = "price_target"
	ENDPOINTENUM_RECOMMENDATIONS                   EndpointEnum = "recommendations"
	ENDPOINTENUM_STATISTICS                        EndpointEnum = "statistics"
	ENDPOINTENUM_INSIDER_TRANSACTIONS              EndpointEnum = "insider_transactions"
	ENDPOINTENUM_PROFILE                           EndpointEnum = "profile"
	ENDPOINTENUM_MUTUAL_FUNDS_WORLD_SUMMARY        EndpointEnum = "mutual_funds_world_summary"
	ENDPOINTENUM_MUTUAL_FUNDS_WORLD                EndpointEnum = "mutual_funds_world"
	ENDPOINTENUM_INSTITUTIONAL_HOLDERS             EndpointEnum = "institutional_holders"
	ENDPOINTENUM_ANALYST_RATING                    EndpointEnum = "analyst_rating"
	ENDPOINTENUM_INCOME_STATEMENT                  EndpointEnum = "income_statement"
	ENDPOINTENUM_INCOME_STATEMENT_QUARTERLY        EndpointEnum = "income_statement_quarterly"
	ENDPOINTENUM_CASH_FLOW                         EndpointEnum = "cash_flow"
	ENDPOINTENUM_CASH_FLOW_QUARTERLY               EndpointEnum = "cash_flow_quarterly"
	ENDPOINTENUM_BALANCE_SHEET                     EndpointEnum = "balance_sheet"
	ENDPOINTENUM_BALANCE_SHEET_QUARTERLY           EndpointEnum = "balance_sheet_quarterly"
	ENDPOINTENUM_MUTUAL_FUNDS_LIST                 EndpointEnum = "mutual_funds_list"
	ENDPOINTENUM_MUTUAL_FUNDS_WORLD_SUSTAINABILITY EndpointEnum = "mutual_funds_world_sustainability"
	ENDPOINTENUM_MUTUAL_FUNDS_WORLD_SUMMARY2       EndpointEnum = "mutual_funds_world_summary"
	ENDPOINTENUM_MUTUAL_FUNDS_WORLD_RISK           EndpointEnum = "mutual_funds_world_risk"
	ENDPOINTENUM_MUTUAL_FUNDS_WORLD_PURCHASE_INFO  EndpointEnum = "mutual_funds_world_purchase_info"
	ENDPOINTENUM_MUTUAL_FUNDS_WORLD_COMPOSITION    EndpointEnum = "mutual_funds_world_composition"
	ENDPOINTENUM_MUTUAL_FUNDS_WORLD_PERFORMANCE    EndpointEnum = "mutual_funds_world_performance"
	ENDPOINTENUM_MUTUAL_FUNDS_WORLD2               EndpointEnum = "mutual_funds_world"
	ENDPOINTENUM_ETFS_LIST                         EndpointEnum = "etfs_list"
	ENDPOINTENUM_ETFS_WORLD                        EndpointEnum = "etfs_world"
	ENDPOINTENUM_ETFS_WORLD_SUMMARY                EndpointEnum = "etfs_world_summary"
	ENDPOINTENUM_ETFS_WORLD_PERFORMANCE            EndpointEnum = "etfs_world_performance"
	ENDPOINTENUM_ETFS_WORLD_RISK                   EndpointEnum = "etfs_world_risk"
	ENDPOINTENUM_ETFS_WORLD_COMPOSITION            EndpointEnum = "etfs_world_composition"
	ENDPOINTENUM_DIVIDENDS                         EndpointEnum = "dividends"
	ENDPOINTENUM_SPLITS                            EndpointEnum = "splits"
)

// All allowed values of EndpointEnum enum
var AllowedEndpointEnumEnumValues = []EndpointEnum{
	"price_target",
	"recommendations",
	"statistics",
	"insider_transactions",
	"profile",
	"mutual_funds_world_summary",
	"mutual_funds_world",
	"institutional_holders",
	"analyst_rating",
	"income_statement",
	"income_statement_quarterly",
	"cash_flow",
	"cash_flow_quarterly",
	"balance_sheet",
	"balance_sheet_quarterly",
	"mutual_funds_list",
	"mutual_funds_world_sustainability",
	"mutual_funds_world_summary",
	"mutual_funds_world_risk",
	"mutual_funds_world_purchase_info",
	"mutual_funds_world_composition",
	"mutual_funds_world_performance",
	"mutual_funds_world",
	"etfs_list",
	"etfs_world",
	"etfs_world_summary",
	"etfs_world_performance",
	"etfs_world_risk",
	"etfs_world_composition",
	"dividends",
	"splits",
}

func (v *EndpointEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EndpointEnum(value)
	for _, existing := range AllowedEndpointEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EndpointEnum", value)
}

// NewEndpointEnumFromValue returns a pointer to a valid EndpointEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEndpointEnumFromValue(v string) (*EndpointEnum, error) {
	ev := EndpointEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EndpointEnum: valid values are %v", v, AllowedEndpointEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EndpointEnum) IsValid() bool {
	for _, existing := range AllowedEndpointEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EndpointEnum value
func (v EndpointEnum) Ptr() *EndpointEnum {
	return &v
}

type NullableEndpointEnum struct {
	value *EndpointEnum
	isSet bool
}

func (v NullableEndpointEnum) Get() *EndpointEnum {
	return v.value
}

func (v *NullableEndpointEnum) Set(val *EndpointEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointEnum(val *EndpointEnum) *NullableEndpointEnum {
	return &NullableEndpointEnum{value: val, isSet: true}
}

func (v NullableEndpointEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
