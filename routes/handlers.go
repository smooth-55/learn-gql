package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/smooth/learn-gql/gql"
	"github.com/smooth/learn-gql/infrastructure"
)

type Handler struct {
	router   infrastructure.Router
	query    gql.RootQuery
	mutation gql.RootMutation
}

// NewUserRoutes creates new user controller
func NewHandler(
	router infrastructure.Router,
	query gql.RootQuery,
	mutation gql.RootMutation,

) Handler {
	return Handler{
		router:   router,
		query:    query,
		mutation: mutation,
	}
}

func (h Handler) GetSchema() *graphql.Schema {
	var Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    h.query.GetQuery(),
		Mutation: h.mutation.GetMutation(),
	})
	if err != nil {
		panic(err)
	}
	return &Schema
}

func (h Handler) RequestHandler(c *gin.Context) {
	// Parse the incoming JSON data
	var requestData map[string]interface{}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}

	// Extract the GraphQL query from the request
	query, _ := requestData["query"].(string)

	result := graphql.Do(graphql.Params{
		Schema:        *h.GetSchema(),
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		c.JSON(400, gin.H{"errors": result.Errors})
		return
	}

	// Return the JSON response
	c.JSON(200, gin.H{"data": result.Data})
}

// Setup user routes
func (i Handler) Setup() {
	gql := i.router.Gin.Group("")
	{
		gql.GET("/playground", func(c *gin.Context) {
			c.HTML(200, "playground.html", gin.H{
				"message": "Graphql playground",
			})
		})
		gql.POST("/gql", i.RequestHandler)
	}
}
