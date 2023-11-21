package core

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// User struct
type User struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name           string    `gorm:"type:varchar(100);not null" json:"name"`
	Email          string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password       string    `gorm:"type:varchar(100);not null" json:"password"`
	Role           uuid.UUID `gorm:"type:uuid;" json:"role"`
	CompanyId      uuid.UUID `gorm:"type:uuid" json:"company_id"`
	Phone          string    `gorm:"not null;" json:"phone"`
	Photo          string    `gorm:"not null;default:'default.png'" json:"photo"`
	FirebaseToken  string    `gorm:"type:varchar(255)" json:"firebase_token"`
	ActivationCode string    `gorm:"type:varchar(255)" json:"activation_code"`
	Status         bool      `gorm:"not null;default:false"`
	CreatedAt      time.Time `gorm:"not null;default:now()"`
	UpdatedAt      time.Time `gorm:"not null;default:now()"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Phone           string `json:"phone" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	Photo           string `json:"photo"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      uuid.UUID `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	CompanyId uuid.UUID `json:"company_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Photo:     user.Photo,
		CompanyId: user.CompanyId,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
