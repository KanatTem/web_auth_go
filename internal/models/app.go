package models

type AppView struct {
	ID          string
	Name        string
	Description string
	Roles       []string
	UserCount   int
}
