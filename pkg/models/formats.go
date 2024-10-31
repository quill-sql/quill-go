package models

type FieldFormat string

const (
	WholeNumber      FieldFormat = "whole_number"
	OneDecimalPlace  FieldFormat = "one_decimal_place"
	TwoDecimalPlaces FieldFormat = "two_decimal_places"
	DollarAmount     FieldFormat = "dollar_amount"
	MMMYYYY          FieldFormat = "MMM_yyyy"
	MMMDDYYYY        FieldFormat = "MMM_dd_yyyy"
	MMMDDToMMMDD     FieldFormat = "MMM_dd-MMM_dd"
	MMMDDHHMMAPPM    FieldFormat = "MMM_dd_hh:mm_ap_pm"
	HHAPPM           FieldFormat = "hh_ap_pm"
	Percent          FieldFormat = "percent"
	StringFormat     FieldFormat = "string"
)

type FormattedColumn struct {
	Label     string      `json:"label"`
	Field     string      `json:"field,omitempty"`
	ChartType string      `json:"chartype,omitempty"`
	Format    FieldFormat `json:"format,omitempty"`
	FieldType string      `json:"fieldType,omitempty"`
}
