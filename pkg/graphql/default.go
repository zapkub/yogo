package graphql

import (
	"net/http"
	"yogo/pkg/models"

	gql_graphql "github.com/99designs/gqlgen/graphql"
	handler "github.com/99designs/gqlgen/handler"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type container interface {
	DB() *mongo.Database
	Models() models.Models
}

type RootResolver struct {
	query    RootQueryResolver
	mutation RootMutationResolver
	user     UserTypeResolver
}

func (rootResolver *RootResolver) Query() QueryResolver {
	return &rootResolver.query
}

func (rootResolver *RootResolver) Mutation() MutationResolver {
	return &rootResolver.mutation
}

func (rootResolver *RootResolver) User() UserResolver {
	return &rootResolver.user
}

// CreateResolvers create resolver from Query and Mutation resolver
func CreateResolvers(container container) ResolverRoot {

	rootQuery := RootQueryResolver{
		container: container,
	}

	rootMutation := RootMutationResolver{
		container: container,
	}

	return &RootResolver{
		query:    rootQuery,
		mutation: rootMutation,
		user: UserTypeResolver{
			container: container,
		},
	}
}

func CreateExecutableSchema(container container) gql_graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers:  CreateResolvers(container),
		Directives: DirectiveRoot{},
		Complexity: ComplexityRoot{},
	})
}

func CreateGraphQLHandler(container container) http.HandlerFunc {
	return handler.GraphQL(CreateExecutableSchema(container))
}

func CreateGraphQLPlaygroundHandler() http.HandlerFunc {
	return handler.Playground("GraphQL playground", "/graphql")
}
