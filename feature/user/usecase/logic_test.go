package usecase

import (
	"cleanarch/domain"
	"cleanarch/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//TestAddUser tests the AddUser function
func TestAddUser(t *testing.T) {
	repo := new(mocks.UserData)

	mockData := domain.User{Nama: "red", Email: "red@red.com", Password: "aasd123"}

	returnData := mockData
	returnData.ID = 1
	returnData.Password = "$2a$10$hG8DockozzNBD4JnQ41gwOA.TOnOMLDGN0evIMR6EY1r5BAcPwhbC"
	t.Run("Success case", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		useCase := UserLogic(repo, validator.New())
		data, err := useCase.AddUser(mockData)
		assert.Nil(t, err)
		assert.Greater(t, data.ID, 0)
		assert.Equal(t, "red", data.Nama)
		assert.Equal(t, "red@red.com", data.Email)
		assert.Equal(t, "$2a$10$hG8DockozzNBD4JnQ41gwOA.TOnOMLDGN0evIMR6EY1r5BAcPwhbC", data.Password, "Password is invalid")
		repo.AssertExpectations(t)
	})
	// t.Run("Validator error", func(t *testing.T) {
	// 	repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
	// 	useCase := UserLogic(repo, validator.New())
	// 	data, err := useCase.AddUser(domain.User{})
	// 	assert.EqualError(t, err, "please make sure all fields are filled in correctly")
	// 	assert.Equal(t, data.ID, 0)
	// 	assert.Equal(t, "", data.Nama)
	// 	assert.Equal(t, "", data.Email)
	// 	assert.Equal(t, "", data.Password, "Password is invalid")
	// 	repo.AssertExpectations(t)
	// })
}
