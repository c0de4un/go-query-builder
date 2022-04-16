package crm

import (
	"testing"
)

func TestSelectorsCompilation(t *testing.T) {
	builder := NewBuilder()
	builder.Select("*")
	builder.Select("projects.*")
	query := builder.compileSelectors()
	if query != "*, projects.*" {
		t.Errorf("\nunexpected selectors compilation result: %s\n", query)
	}
}

func TestJoinsCompilation(t *testing.T) {
	builder := NewBuilder()
	builder.Join("roles", "roles.user_id = users.id", "LEFT")
	query := builder.compileJoins()
	if query != "LEFT JOIN roles ON roles.user_id = users.id" {
		t.Errorf("\nunexpected joins compilation result: %s\n", query)
	}
}

func TestPredicatesCompilation(t *testing.T) {
	builder := NewBuilder()
	builder.Where("users.id", ">", "1")
	query := builder.compilePredicates()
	if query != "WHERE users.id > 1" {
		t.Errorf("\nunexpected predicates compilation result: %s\n", query)
	}

	builder.Where("users.created_at", ">=", "2022-06-04")
	query = builder.compilePredicates()
	if query != "WHERE users.id > 1 AND users.created_at >= 2022-06-04" {
		t.Errorf("\nunexpected predicates compilation result: %s\n", query)
	}
}
