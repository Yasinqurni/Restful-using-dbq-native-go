package entity_test

import (
	"restful-api/domain/entity"
	"restful-api/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserEntity(t *testing.T) {

	type expected struct {
		Name     string
		Email    string
		Phone    string
		Password string
		Role     string
		Gender   string
	}
	type testCases struct {
		name     string
		dto      entity.UserDTO
		expected *expected
		isError  bool
		err      any
	}

	testcases := []testCases{
		{
			name: "positive",
			dto: entity.UserDTO{
				Name:     "sholeh",
				Email:    "email1@gmail.com",
				Phone:    "082272615810",
				Password: "123456",
				Role:     "user",
				Gender:   "male",
			},
			expected: &expected{
				Name:     "sholeh",
				Email:    "email1@gmail.com",
				Phone:    "082272615810",
				Password: "123456",
				Role:     "user",
				Gender:   "male",
			},
			isError: false,
			err:     nil,
		},
		{
			name: "negative with name is empty string",
			dto: entity.UserDTO{
				Name:     "",
				Email:    "email2@gmail.com",
				Phone:    "082272615810",
				Password: "123456",
				Role:     "user",
				Gender:   "male",
			},
			expected: nil,
			isError:  true,
			err:      pkg.EMPTY_NAME,
		},
		{
			name: "negative with email is empty string",
			dto: entity.UserDTO{
				Name:     "sholeh",
				Email:    "",
				Phone:    "082272615810",
				Password: "123456",
				Role:     "user",
				Gender:   "male",
			},
			expected: nil,
			isError:  true,
			err:      pkg.ERROR_EMAIL,
		},
		{
			name: "negative with phone length is 9",
			dto: entity.UserDTO{
				Name:     "sholeh",
				Email:    "email3@gmail.com",
				Phone:    "082278989",
				Password: "123456",
				Role:     "user",
				Gender:   "male",
			},
			expected: nil,
			isError:  true,
			err:      pkg.ERROR_PHONE,
		},
		{
			name: "negative with password length is 5",
			dto: entity.UserDTO{
				Name:     "sholeh",
				Email:    "email4@gmail.com",
				Phone:    "082272615810",
				Password: "12345",
				Role:     "user",
				Gender:   "male",
			},
			expected: nil,
			isError:  true,
			err:      pkg.ERROR_PASSWORD,
		},
		{
			name: "negative with role is something else",
			dto: entity.UserDTO{
				Name:     "sholeh",
				Email:    "email5@gmail.com",
				Phone:    "082272615810",
				Password: "123456",
				Role:     "role",
				Gender:   "male",
			},
			expected: nil,
			isError:  true,
			err:      pkg.ERROR_ROLE,
		},
		{
			name: "negative with gender is something else",
			dto: entity.UserDTO{
				Name:     "sholeh",
				Email:    "email6@gmail.com",
				Phone:    "082272615810",
				Password: "123456",
				Role:     "user",
				Gender:   "other",
			},
			expected: nil,
			isError:  true,
			err:      pkg.ERROR_GENDER,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {

			user, err := entity.NewUser(test.dto)

			if !test.isError {
				assert.Equal(t, test.expected.Name, user.GetName())
				assert.Equal(t, test.expected.Email, user.GetEmail())
				assert.Equal(t, test.expected.Password, user.GetPassword())
				assert.Equal(t, test.expected.Phone, user.GetPhone())
				assert.Equal(t, test.expected.Role, user.GetRole())
			} else {
				assert.Equal(t, test.err, err.Error())
			}

		})
	}
}
