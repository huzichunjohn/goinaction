package app

// Widget represents the widget that we create with our app.
type Widget struct {
	ID     int
	UserID int
	Name   string
	Price  int
	Color  string
}

// User represents a user in our system.
type User struct {
	ID    int
	Email string
	Token string
}
