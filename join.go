package crm

type join struct {
	tableName  string
	conditions []string
	side       string
}

func newJoin(tableName string, condition string, side string) *join {
	var join join
	join.tableName = tableName
	join.conditions = make([]string, 1)
	join.conditions = append(join.conditions, condition)
	join.side = side
	return &join
}

func (join *join) compile() string {
	output := join.side + " JOIN " + join.tableName

	output += " ON "
	for i, v := range join.conditions {
		if i > 1 {
			output += " AND " + v
		} else {
			output += v
		}
	}

	return output
}
