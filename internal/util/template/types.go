package template

type Data struct {
	User    User
	Account Account
}

type User struct {
	Id       string
	Name     string
	FullName string
	Email    string
}

type Account struct {
	Id        string
	Name      string
	LoginName string
	Subject   string
	Email     string
}
