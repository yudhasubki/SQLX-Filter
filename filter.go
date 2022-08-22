package sqlxfilter

import (
	"fmt"
	"strings"
)

type Filter struct {
	conditions
	order
	limit
	group
}

type FilterFunc func(f *Filter)

func New(filterFunc ...FilterFunc) *Filter {
	f := &Filter{}
	for _, filter := range filterFunc {
		filter(f)
	}

	return f
}

// QueryClause that generate clause and its arguments
// the operator has three conditions "AND", "OR", and "NOT".
func (f *Filter) QueryClause(operator string) ([]interface{}, string) {
	var (
		args    = make([]interface{}, 0)
		clauses = make([]string, 0)
	)

	if len(f.conditions) == 0 {
		return args, ""
	}

	for _, c := range f.conditions {
		args = append(args, c.Value)
		clauses = append(clauses, c.buildQuery())
	}

	return args, strings.Join(clauses, fmt.Sprintf(" %s ", operator))
}

// Limit function that generate Limit Query
func (f *Filter) Limit() string {
	if f.limit.Size == 0 {
		return ""
	}

	return fmt.Sprintf(" LIMIT %d", f.limit.Size)
}

// Paginate function that generate limit query and its offset.
func (f Filter) Paginate() string {
	if f.limit.Page == 0 && f.limit.Size == 0 {
		return ""
	}

	offset := (f.limit.Page - 1) * f.limit.Size

	return fmt.Sprintf(" LIMIT %d OFFSET %d", f.limit.Size, offset)
}

// SortBy function that generate Order By query.
func (o Filter) SortBy() string {
	if len(o.order.Columns) == 0 {
		return ""
	}

	return fmt.Sprintf(" ORDER BY %s %s", strings.Join(o.order.Columns, ","), o.Direction)
}

// Group function that generate Group By column.
func (o Filter) Group() string {
	if len(o.group.Columns) == 0 {
		return ""
	}

	return fmt.Sprintf(" GROUP BY %s", strings.Join(o.group.Columns, ","))
}

type condition struct {
	Field    string
	Operator string
	Value    interface{}
}

type conditions []condition

func (filter condition) buildQuery() string {
	var clause string

	switch strings.ToLower(filter.Operator) {
	case "in":
		clause = fmt.Sprintf("%s IN(?)", filter.Field)
	default:
		clause = fmt.Sprintf("%s %s ?", filter.Field, filter.Operator)
	}

	return clause
}

type order struct {
	Direction string
	Columns   []string
}

type limit struct {
	Size int
	Page int
}

type group struct {
	Columns []string
}
