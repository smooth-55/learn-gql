package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/smooth/learn-gql/apps/todo"
)

type RootMutation struct {
	todoMutation todo.Mutation
}

func NewRootMutation(
	todoMutation todo.Mutation,

) RootMutation {
	return RootMutation{
		todoMutation: todoMutation,
	}
}

func (r RootMutation) GetMutation() *graphql.Object {
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation", // Always set name as Mutation
		Fields: graphql.Fields{
			"createTodo": r.todoMutation.CreateTodoField(),
			// add other mutations
		},
	})
	return rootQuery

}
