package web

type ContactSearchRequest struct {
	Name  string
	Email string
	Phone string
	Page  int
	Limit int
}