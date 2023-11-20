package todo

import "github.com/graphql-go/graphql"

type Mutation struct{}

func NewMutation() Mutation { return Mutation{} }

func (m Mutation) CreateTodoInputType() *graphql.InputObject {
	obj := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateTodoInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String), // required
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.String, // optional
			},
		},
	})
	return obj
}

func (m Mutation) TodoResponseType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "TodoType",
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

func (m Mutation) CreateTodoResolver(p graphql.ResolveParams) (interface{}, error) {
	inputs := p.Args["input"].(map[string]interface{})
	todo := Todo{
		Id:          1,
		Title:       inputs["title"].(string),
		Description: inputs["description"].(string),
		IsCompleted: false,
	}
	// implement logic to save this object to database
	return todo, nil
}

func (m Mutation) CreateTodoField() *graphql.Field {
	return &graphql.Field{
		Type:    m.TodoResponseType(),
		Resolve: m.CreateTodoResolver,
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: m.CreateTodoInputType(),
			},
		},
	}
}
