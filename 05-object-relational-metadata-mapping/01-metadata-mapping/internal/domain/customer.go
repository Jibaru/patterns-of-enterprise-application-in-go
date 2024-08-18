package domain

type Customer struct {
	ID       int    `db:"id"`
	FullName string `db:"name"`
	Email    string `db:"email"`
}
