package schema

// Schema is the GraphQL standard definition of our schema
var Schema = `
	schema {
		query: Query
	}

	type User {
		id: ID!
		email: String!
		username: String!
		password: String!
	}

	type Query {
		users: [User]!
	}
`

// Resolver will contain all resolvers for our schema
type Resolver struct{}
