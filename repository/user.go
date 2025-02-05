package repositories

import (
	"fmt"
	"grpc-practise/models"
)

type UserRepository struct {
	Users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Users: []models.User{},
	}
}

func (repository *UserRepository) CreateUser(user models.User) (models.User, error) {
	newUser := models.User{
		Id:       fmt.Sprintf("%d", len(repository.Users)+1),
		Name:     user.Name,
		Location: user.Location,
		Title:    user.Title,
	}

	repository.Users = append(repository.Users, newUser)

	return newUser, nil
}

func (repository *UserRepository) GetUser(id string) (models.User, error) {
	for _, user := range repository.Users {
		if user.Id == id {
			return user, nil
		}
	}
	return models.User{}, fmt.Errorf("user not found")
}

func (repository *UserRepository) UpdateUser(id string, newUser models.User) (models.User, error) {
	for index, user := range repository.Users {
		if user.Id == id {
			repository.Users[index].Name = newUser.Name
			repository.Users[index].Location = newUser.Location
			repository.Users[index].Title = newUser.Title
			return repository.Users[index], nil
		}
	}
	return models.User{}, fmt.Errorf("user not found")
}

func (repository *UserRepository) DeleteUser(id string) error {
	for index, user := range repository.Users {
		if user.Id == id {
			repository.Users = append(repository.Users[:index], repository.Users[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func (repository *UserRepository) GetAllUsers() ([]models.User, error) {
	return repository.Users, nil
}
