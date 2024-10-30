package models

type Operator interface {
	isOperator()
}

type StringOperator string

func (StringOperator) isOperator() {}

type DateOperator string

func (DateOperator) isOperator() {}

type NumberOperator string

func (NumberOperator) isOperator() {}

type NullOperator string

func (NullOperator) isOperator() {}

type BoolOperator string

func (BoolOperator) isOperator() {}

const (
	STRING_IS_EXACTLY     StringOperator = "is exactly"
	STRING_IS_NOT_EXACTLY StringOperator = "is not exactly"
	STRING_CONTAINS       StringOperator = "contains"
	STRING_IS             StringOperator = "is"
	STRING_IS_NOT         StringOperator = "is not"
)

const (
	DATE_CUSTOM                   DateOperator = "custom"
	DATE_IN_THE_LAST              DateOperator = "in the last"
	DATE_IN_THE_PREVIOUS          DateOperator = "in the previous"
	DATE_IN_THE_CURRENT           DateOperator = "in the current"
	DATE_EQUAL_TO                 DateOperator = "equal to"
	DATE_NOT_EQUAL_TO             DateOperator = "not equal to"
	DATE_GREATER_THAN             DateOperator = "greater than"
	DATE_LESS_THAN                DateOperator = "less than"
	DATE_GREATER_THAN_OR_EQUAL_TO DateOperator = "greater than or equal to"
	DATE_LESS_THAN_OR_EQUAL_TO    DateOperator = "less than or equal to"
)

const (
	NUMBER_EQUAL_TO                 NumberOperator = "equal to"
	NUMBER_NOT_EQUAL_TO             NumberOperator = "not equal to"
	NUMBER_GREATER_THAN             NumberOperator = "greater than"
	NUMBER_LESS_THAN                NumberOperator = "less than"
	NUMBER_GREATER_THAN_OR_EQUAL_TO NumberOperator = "greater than or equal to"
	NUMBER_LESS_THAN_OR_EQUAL_TO    NumberOperator = "less than or equal to"
)

const (
	IS_NOT_NULL NullOperator = "is not null"
	IS_NULL     NullOperator = "is null"
)

const (
	BOOL_EQUAL_TO     BoolOperator = "equal to"
	BOOL_NOT_EQUAL_TO BoolOperator = "not equal to"
)

type TimeUnit string

const (
	YEAR    TimeUnit = "year"
	QUARTER TimeUnit = "quarter"
	MONTH   TimeUnit = "month"
	WEEK    TimeUnit = "week"
	DAY     TimeUnit = "day"
	HOUR    TimeUnit = "hour"
)

type FieldType string

const (
	STRING  FieldType = "string"
	NUMBER  FieldType = "number"
	DATE    FieldType = "date"
	NULL    FieldType = "null"
	CUSTOM  FieldType = "custom"
	BOOLEAN FieldType = "boolean"
)

type FilterType string

const (
	STRING_FILTER          FilterType = "string-filter"
	DATE_FILTER            FilterType = "date-filter"
	DATE_CUSTOM_FILTER     FilterType = "date-custom-filter"
	DATE_COMPARISON_FILTER FilterType = "date-comparison-filter"
	NUMERIC_FILTER         FilterType = "numeric-filter"
	NULL_FILTER            FilterType = "null-filter"
	STRING_IN_FILTER       FilterType = "string-in-filter"
	BOOLEAN_FILTER         FilterType = "boolean-filter"
)

type DateRange struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type DateValue struct {
	Value int      `json:"value"`
	Unit  TimeUnit `json:"unit"`
}

type BaseFilter struct {
	FilterType FilterType  `json:"filterType"`
	FieldType  FieldType   `json:"fieldType"`
	Operator   Operator    `json:"operator"`
	Field      string      `json:"field"`
	Value      interface{} `json:"value"`
	Table      *string     `json:"table,omitempty"`
}

type Filter struct {
	FilterType FilterType  `json:"filterType"`
	Operator   Operator    `json:"operator"`
	Value      interface{} `json:"value"`
	Field      string      `json:"field"`
	Table      string      `json:"table"`
}

func ConvertCustomFilter(filter Filter) BaseFilter {
	switch filter.FilterType {
	case STRING_FILTER:
		if _, ok := filter.Value.(string); !ok {
			panic("Invalid value for StringFilter, expected string")
		}
		if _, ok := filter.Operator.(StringOperator); !ok {
			panic("Invalid operator for StringFilter, expected StringOperator")
		}
		return BaseFilter{
			FilterType: filter.FilterType,
			FieldType:  STRING,
			Operator:   filter.Operator,
			Field:      filter.Field,
			Value:      filter.Value,
			Table:      &filter.Table,
		}
	case STRING_IN_FILTER:
		if _, ok := filter.Value.([]string); !ok {
			panic("Invalid value for StringInFilter, expected list")
		}
		if _, ok := filter.Operator.(StringOperator); !ok {
			panic("Invalid operator for StringInFilter, expected StringOperator")
		}
		return BaseFilter{
			FilterType: filter.FilterType,
			FieldType:  STRING,
			Operator:   filter.Operator,
			Field:      filter.Field,
			Value:      filter.Value,
			Table:      &filter.Table,
		}
	case NUMERIC_FILTER:
		if _, ok := filter.Value.(int); !ok {
			panic("Invalid value for NumericFilter, expected int")
		}
		if _, ok := filter.Operator.(NumberOperator); !ok {
			panic("Invalid operator for NumericFilter, expected NumberOperator")
		}
		return BaseFilter{
			FilterType: filter.FilterType,
			FieldType:  NUMBER,
			Operator:   filter.Operator,
			Field:      filter.Field,
			Value:      filter.Value,
			Table:      &filter.Table,
		}
	case DATE_FILTER:
		if _, ok := filter.Value.(DateValue); !ok {
			panic("Invalid value for DateFilter, expected DateValue")
		}
		if _, ok := filter.Operator.(DateOperator); !ok {
			panic("Invalid operator for DateFilter, expected DateOperator")
		}
		return BaseFilter{
			FilterType: filter.FilterType,
			FieldType:  DATE,
			Operator:   filter.Operator,
			Field:      filter.Field,
			Value:      filter.Value,
			Table:      &filter.Table,
		}
	case DATE_CUSTOM_FILTER:
		if _, ok := filter.Value.(DateRange); !ok {
			panic("Invalid value for DateCustomFilter, expected DateRange")
		}
		if _, ok := filter.Operator.(DateOperator); !ok {
			panic("Invalid operator for DateCustomFilter, expected DateOperator")
		}
		return BaseFilter{
			FilterType: filter.FilterType,
			FieldType:  DATE,
			Operator:   filter.Operator,
			Field:      filter.Field,
			Value:      filter.Value,
			Table:      &filter.Table,
		}
	case DATE_COMPARISON_FILTER:
		if _, ok := filter.Value.(string); !ok {
			panic("Invalid value for DateComparisonFilter, expected str")
		}
		if _, ok := filter.Operator.(DateOperator); !ok {
			panic("Invalid operator for DateComparisonFilter, expected DateOperator")
		}
		return BaseFilter{
			FilterType: filter.FilterType,
			FieldType:  DATE,
			Operator:   filter.Operator,
			Field:      filter.Field,
			Value:      filter.Value,
			Table:      &filter.Table,
		}
	case NULL_FILTER:
		if filter.Value != nil {
			panic("Invalid value for NullFilter, expected None")
		}
		if _, ok := filter.Operator.(NullOperator); !ok {
			panic("Invalid operator for NullFilter, expected NullOperator")
		}
		return BaseFilter{
			FilterType: filter.FilterType,
			FieldType:  NULL,
			Operator:   filter.Operator,
			Field:      filter.Field,
			Value:      filter.Value,
			Table:      &filter.Table,
		}
	case BOOLEAN_FILTER:
		if _, ok := filter.Value.(bool); !ok {
			panic("Invalid value for BooleanFilter, expected bool")
		}
		if _, ok := filter.Operator.(BoolOperator); !ok {
			panic("Invalid operator for BooleanFilter, expected BoolOperator")
		}
		return BaseFilter{
			FilterType: filter.FilterType,
			FieldType:  BOOLEAN,
			Operator:   filter.Operator,
			Field:      filter.Field,
			Value:      filter.Value,
			Table:      &filter.Table,
		}
	default:
		panic("Unknown filter type: " + filter.FilterType)
	}
}
