package utils

import "errors"

const (
	// One operation returns a single document from the database
	One string = "one"

	// All operation returns multiple documents from the database
	All string = "all"

	// Count operation returns the number of documents which match the condition
	Count string = "count"

	// Distinct operation returns distinct values
	Distinct string = "distinct"

	// Upsert creates a new document if it doesn't exist, else it updates exiting document
	Upsert string = "upsert"
)

var ErrInvalidParams = errors.New("Invalid parameter provided")

// GraphQLGroupByArgument is used by graphql group clause
const GraphQLGroupByArgument = "group"

// GraphQLAggregate is used by graphql aggregate clause
const GraphQLAggregate = "aggregate"
