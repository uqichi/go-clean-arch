package iam

type IAuthenticationService interface {
	Login()
	Logout()
	IsUserAuthenticated() bool
}

type authenticationService struct {
}
