package models

// User contains user models.
type User struct {
	ID        int64  `db:"id" json:"id,omitempty"`
	FirstName string `db:"first_name" json:"first_name,omitempty" valid:"alphanum,required"`
	LastName  string `db:"last_name" json:"last_name,omitempty" valid:"alphanum,required"`
	Email     string `db:"email" json:"email,omitempty" valid:"email,required"`
}
