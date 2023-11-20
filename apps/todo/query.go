package todo

import "github.com/graphql-go/graphql"

type Query struct {
}

func NewQuery() Query {
	return Query{}
}

func (q Query) GetAllTodoResolver(p graphql.ResolveParams) (interface{}, error) {
	todos := []Todo{
		{
			Id:          1,
			Title:       "todo 1",
			Description: "todo 1 desc",
			IsCompleted: true,
		},
		{
			Id:          2,
			Title:       "todo 2",
			Description: "todo 2 desc",
			IsCompleted: false,
		},
	}
	return todos, nil
}

func (q Query) TodoType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "GetAllTodo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"is_completed": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	})

}

func (q Query) GetAllTodoField() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(q.TodoType()),
		// we can add arguments if needed
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: q.GetAllTodoResolver,
	}
}
