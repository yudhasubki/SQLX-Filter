# Sqlx Filter

sqlxfilter is a simple package that helps a user to generate clauses, sort by, limit, offset, or group by query—using an option's function to make it easy.

## Installation

Use to get to install this package.

```bash
go get -u github.com/yudhasubki/sqlxfilter
```

## Usage

```go

type IncomingSearchRequest struct {
	Id            string
	Name          string
	SortDirection string // ASC or DESC
	SortBy        string // column that needs to be sorted
	Size          int    // Limit size
}

// transform to SearchByModel struct
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

	if r.SortDirection != "" {
		searchBy.SortDirection = r.SortDirection
	}

	if r.SortBy != "" {
		searchBy.SortBy = r.SortBy
	}

	return searchBy
}

// struct used for filtering columns in the database
type SearchByModel struct {
        Ids           []string
	Name          []string
	Size          int
	Page          int
	SortDirection string
	SortBy        string
}

// extract to opts function
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

	if s.SortDirection != "" && s.SortBy != "" {
		fn = append(fn, filter.OrderBy(s.SortDirection, s.SortBy))
	}

	return fn
}

func main() {
    request := SearchRequest{
		Id:            "1",
		Size:          10,
		SortDirection: "asc",
		SortBy:        "name",
    }

    f := filter.New(request.SearchBy().Filter()...)

    f.QueryClause("OR") // output : args = [[1] [Kuncoro]] || query = id IN(?) OR name IN(?)
    
    f.SortBy() // ORDER BY name ASC

    f.Limit() // LIMIT 10
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
