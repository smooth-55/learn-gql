package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/smooth/learn-gql/apps/todo"
)

type RootQuery struct {
	todoQuery todo.Query
}

func NewRootQuery(
	todoQuery todo.Query,

) RootQuery {
	return RootQuery{
		todoQuery: todoQuery,
	}
}

func (r RootQuery) GetQuery() *graphql.Object {
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query", // Always set name as Query
		Fields: graphql.Fields{
			// key: Query name
			// value: actual query
			"todos": r.todoQuery.GetAllTodoField(),
			// add other queries
		},
	})
	return rootQuery

}
