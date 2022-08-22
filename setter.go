package sqlxfilter

import "strings"

func Where(field, operator string, value interface{}) FilterFunc {
	return func(f *Filter) {
		f.conditions = append(f.conditions, condition{
			Field:    field,
			Operator: operator,
			Value:    value,
		})
	}
}

func Equal(field string, value interface{}) FilterFunc {
	return func(f *Filter) {
		f.conditions = append(f.conditions, condition{
			Field:    field,
			Operator: "=",
			Value:    value,
		})
	}
}

func In(field string, value interface{}) FilterFunc {
	return func(f *Filter) {
		f.conditions = append(f.conditions, condition{
			Field:    field,
			Operator: "IN",
			Value:    value,
		})
	}
}

func OrderBy(direction string, columns ...string) FilterFunc {
	return func(f *Filter) {
		if len(columns) > 0 {
			f.order = order{
				Direction: strings.ToUpper(direction),
				Columns:   columns,
			}
		}
	}
}

func Limit(size int) FilterFunc {
	return func(f *Filter) {
		f.limit.Size = size
	}
}

func Paginate(size int, page int) FilterFunc {
	return func(f *Filter) {
		f.limit.Size = size
		f.limit.Page = page
	}
}
