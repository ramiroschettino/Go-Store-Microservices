//go:generate mockgen -source=../domain/user_repository.go -destination=./mock_user_repository.go -package=mocks
//go:generate mockgen -source=../domain/token_repository.go -destination=./mock_token_repository.go -package=mocks
package mocks