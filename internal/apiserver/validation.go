package apiserver

import (
	"errors"
	"regexp"
	"unicode"

	"github.com/ElnurKoke/web-api.git/internal/model"
)

func validUser(user model.User) error {
	for _, char := range user.Username {
		if char <= 32 || char >= 127 {
			return ErrInvalidUserName
		}
	}
	validEmail, err := regexp.MatchString(`[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`, user.Email)
	if err != nil {
		return err
	}
	if !validEmail {
		return ErrInvalidEmail
	}
	if len(user.Username) < 6 || len(user.Username) >= 36 {
		return ErrInvalidUserName
	}

	if !passIsValid(user.Password) {
		return ErrShortPassword
	}
	if user.Password != user.Repeat {
		return ErrPasswordDontMatch
	}
	return nil
}

func passIsValid(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 8 || len(s) <= 20 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

var (
	ErrCommentNotFound   = errors.New("The page does not exist or comment has been deleted")
	ErrPostNotFound      = errors.New("The page does not exist or post has been deleted")
	ErrUserNotFound      = errors.New("User does not exist or password incorrect")
	ErrInvalidUserName   = errors.New("Invalid username - your username should consist at least 6 characters")
	ErrInvalidEmail      = errors.New("Invalid email")
	ErrPasswordDontMatch = errors.New("Password didn't match")
	ErrShortPassword     = errors.New("Incorrect password - your password should be a minimum of 8 characters and consist of at least:1 lower case letter, 1 upper case letter, 1 number, 1 special symbol")
)

func validName(name string) error {
	if len(name) < 6 || len(name) > 20 {
		return ErrInvalidUserName
	}
	for _, char := range name {
		if char <= 32 || char >= 127 {
			return ErrInvalidUserName
		}
	}
	return nil
}

func validEmail(email string) error {
	validEmail, err := regexp.MatchString(`[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`, email)
	if err != nil {
		return err
	}
	if !validEmail {
		return ErrInvalidEmail
	}
	return nil
}

func mergeProjects(project, newProject model.Project) model.Project {
	if newProject.ProjectName == "" {
		newProject.ProjectName = project.ProjectName
	}
	if newProject.Category == "" {
		newProject.Category = project.Category
	}
	if newProject.ProjectType == "" {
		newProject.ProjectType = project.ProjectType
	}
	if newProject.ReleaseYear == 0 {
		newProject.ReleaseYear = project.ReleaseYear
	}
	if newProject.AgeCategory == "" {
		newProject.AgeCategory = project.AgeCategory
	}
	if newProject.Duration == "" {
		newProject.Duration = project.Duration
	}
	if newProject.Director == "" {
		newProject.Director = project.Director
	}
	if newProject.Producer == "" {
		newProject.Producer = project.Producer
	}
	return newProject
}
