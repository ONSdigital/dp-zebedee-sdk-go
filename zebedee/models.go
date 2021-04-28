package zebedee

//User defines the CMS user structure
type User struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Inactive          bool   `json:"inactive"`
	LastAdmin         string `json:"lastAdmin"`
	TemporaryPassword bool   `json:"temporaryPassword"`
}

// Credentials is the model representing the user login details
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Session is the model of a CMS user session.
type Session struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

// Permissions is the model representing user's CMS permissions
type Permissions struct {
	Email  string `json:"email"`
	Admin  bool   `json:"admin"`
	Editor bool   `json:"editor"`
}
