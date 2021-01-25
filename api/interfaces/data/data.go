package data

// DbElement represent an element of the database
type DbElement map[string]interface{}

// Matches should return true if the given element matches search criterias
type Matches func(DbElement) bool

// Finder allow to get all DbElement that Matches defined criterias
type Finder interface {
	All(f Matches) []DbElement
	First(f Matches) DbElement
}
