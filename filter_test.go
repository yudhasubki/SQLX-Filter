package sqlxfilter_test

import (
	"fmt"
	"testing"

	filter "github.com/yudhasubki/sqlx-filter"
)

type SearchRequest struct {
	Id            string
	Name          string
	SortDirection string
	SortBy        string
	Size          int
	Page          int
}

func (r *SearchRequest) SearchBy() *SearchBy {
	searchBy := &SearchBy{}

	if r.Id != "" {
		searchBy.Ids = []string{r.Id}
	}

	if r.Name != "" {
		searchBy.Name = []string{r.Name}
	}

	if r.Size > 0 {
		searchBy.Size = r.Size
	}

	if r.Page > 0 {
		searchBy.Page = r.Page
	}

	if r.SortDirection != "" {
		searchBy.SortDirection = r.SortDirection
	}

	if r.SortBy != "" {
		searchBy.SortBy = r.SortBy
	}

	return searchBy
}

type SearchBy struct {
	Ids           []string
	Name          []string
	Size          int
	Page          int
	SortDirection string
	SortBy        string
}

func (s *SearchBy) Filter() []filter.FilterFunc {
	fn := make([]filter.FilterFunc, 0)

	if len(s.Ids) > 0 {
		fn = append(fn, filter.In("id", s.Ids))
	}

	if len(s.Name) > 0 {
		fn = append(fn, filter.In("name", s.Name))
	}

	if s.Size > 0 {
		fn = append(fn, filter.Limit(s.Size))
	}

	if s.Page > 0 && s.Size > 0 {
		fn = append(fn, filter.Paginate(s.Size, s.Page))
	}

	if s.SortDirection != "" && s.SortBy != "" {
		fn = append(fn, filter.OrderBy(s.SortDirection, s.SortBy))
	}

	return fn
}

func TestFilter(t *testing.T) {
	request := SearchRequest{
		Id:            "1",
		Name:          "Kuncoro",
		Size:          10,
		Page:          5,
		SortDirection: "asc",
		SortBy:        "name",
	}

	query := "SELECT * FROM"

	f := filter.New(request.SearchBy().Filter()...)

	args, clause := f.QueryClause("OR")
	if len(args) > 0 {
		query += " WHERE " + clause
	}

	if f.SortBy() != "" {
		query += f.SortBy()
	}

	if f.Paginate() != "" {
		query += f.Paginate()
	}

	fmt.Println(query)
}
