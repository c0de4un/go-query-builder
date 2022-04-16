package crm

type Builder struct {
	selectors []string
	from      string
	joins     []*join
	wheres    []*where
}

func NewBuilder() *Builder {
	var builder Builder
	builder.selectors = make([]string, 1)
	builder.joins = make([]*join, 0)
	builder.wheres = make([]*where, 0)

	return &builder
}

func (builder *Builder) Compile() (string, error) {
	query := builder.compileSelectors()
	query += " " + builder.from
	if len(builder.joins) > 0 {
		query += " " + builder.compileJoins()
	}

	if len(builder.wheres) > 0 {
		query += " " + builder.compilePredicates()
	}

	return query + ";", nil
}

func (builder *Builder) Select(query string) {
	builder.selectors = append(builder.selectors, query)
}

func (builder *Builder) From(from string) {
	builder.from = from
}

func (builder *Builder) compileSelectors() string {
	var selector string
	for i, val := range builder.selectors {
		if i > 1 {
			selector += ", "
		}
		selector += val
	}

	return selector
}

func (builder *Builder) Join(tableName string, condition string, side string) {
	command := newJoin(tableName, condition, side)
	builder.joins = append(builder.joins, command)
}

func (builder *Builder) compileJoins() string {
	var joins string

	for i, join := range builder.joins {
		if i > 1 {
			joins += ", "
		}

		joins += join.compile()
	}

	return joins
}

func (builder *Builder) Where(column string, operator string, value string) {
	predicate := newWhere(column, operator, value)
	builder.wheres = append(builder.wheres, predicate)
}

func (builder *Builder) compilePredicates() string {
	if len(builder.wheres) < 1 {
		return ""
	}

	output := "WHERE "
	for i, v := range builder.wheres {
		if i > 0 {
			output += " AND " + v.compileWhere()
		} else {
			output += v.compileWhere()
		}
	}

	return output
}
