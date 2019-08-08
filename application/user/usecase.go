package user

type IUseCase interface {
	ListUsers(input ListUsersInput) (ListUsersOutput, error)
}

type useCase struct {
}
