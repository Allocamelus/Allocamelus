package user

import (
	"database/sql"
	"errors"
	"regexp"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/pkg/logger"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/nbutton23/zxcvbn-go"
)

func initValidate(p data.Prepare) {
	preCheckUserName = p(`SELECT EXISTS(SELECT * FROM Users WHERE userName=?)`)
	preCheckEmail = p(`SELECT EXISTS(SELECT * FROM Users WHERE email=?)`)
}

const (
	invalidLength = "invalid-length"
	taken         = "taken"
)

// ValidatePublic values
func (u *User) ValidatePublic() error {
	errs := make(validation.Errors)
	if err := u.ValidUserName(); err != nil {
		errs["userName"] = err
	}
	if err := u.ValidEmail(); err != nil {
		errs["email"] = err
	}
	if err := u.ValidBio(); err != nil {
		errs["bio"] = err
	}
	return errs.Filter()
}

var (
	// ErrUserNameLength 5 to 64 characters
	ErrUserNameLength = errors.New(invalidLength + "-min5-max64")
	// ErrUserNameTaken characters
	ErrUserNameTaken = errors.New(taken)
	regexUserName    = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`)
	preCheckUserName *sql.Stmt
)

// ValidUserName Validate
func (u *User) ValidUserName() error {
	// Check Length
	if err := validation.Validate(u.UserName,
		validation.Required,
		validation.Length(5, 64),
	); err != nil {
		return ErrUserNameLength
	}
	// Check regex for User Name
	if !regexUserName.MatchString(u.UserName) {
		return g.ErrInvalidChars
	}
	// Check Database for userName
	var isTaken bool
	err := preCheckUserName.QueryRow(u.UserName).Scan(&isTaken)
	if err != nil && err != sql.ErrNoRows {
		logger.Error(err)
	}
	if isTaken {
		return ErrUserNameTaken
	}

	return nil
}

var (
	// ErrNameLength 1 to 128 characters
	ErrNameLength = errors.New(invalidLength + "-min1-max128")
)

// ValidName Validate
func (u *User) ValidName() error {
	return ValidName(u.Name)
}

// ValidName Validate
func ValidName(name string) error {
	// Check Length
	if err := validation.Validate(name,
		validation.Required,
		validation.Length(1, 128),
	); err != nil {
		return ErrNameLength
	}
	// Check regex for Invalid characters
	if !g.ContentInvalidChars.MatchString(name) {
		return g.ErrInvalidChars
	}
	return nil
}

var (
	// ErrEmailInvalid Invalid Email
	ErrEmailInvalid = errors.New("invalid-email")
	// ErrEmailTaken characters
	ErrEmailTaken = errors.New(taken)
	preCheckEmail *sql.Stmt
)

// ValidEmail Validate
func (u *User) ValidEmail() error {
	return ValidEmail(u.Email)
}

// ValidEmail Validate
func ValidEmail(email string) error {
	// Check Email
	if err := validation.Validate(email,
		validation.Required,
		is.Email,
	); err != nil {
		return ErrEmailInvalid
	}
	return nil
}

// IsEmailUnique check if Email Exists
func (u *User) IsEmailUnique() error {
	if u.Email == "" {
		return ErrEmailInvalid
	}
	// Check Database for userName
	var isTaken bool
	err := preCheckEmail.QueryRow(u.Email).Scan(&isTaken)
	if err != nil && err != sql.ErrNoRows {
		logger.Error(err)
	}
	if isTaken {
		return ErrEmailTaken
	}
	return nil
}

var (
	// ErrBioLength 0 to 255 characters
	ErrBioLength = errors.New(invalidLength + "-min0-max255")
)

// ValidBio Validate
func (u *User) ValidBio() error {
	return ValidBio(u.Bio)
}

func ValidBio(bio string) error {
	if bio != "" {
		// Check length
		if err := validation.Validate(bio, validation.Length(0, 255)); err != nil {
			return ErrBioLength
		}
		// Check regex for Invalid characters
		if !g.ContentInvalidChars.MatchString(bio) {
			return g.ErrInvalidChars
		}
	}
	return nil
}

// ErrPasswordLength 8 to 1024
var (
	ErrPasswordLength   = errors.New(invalidLength + "-min8-max1024")
	ErrPasswordStrength = errors.New("weak-password")
)

// ValidPassword check password
func (u *User) ValidPassword(pass string) error {
	return ValidPassword(pass, u.UserName, u.Name, u.Email)
}

// ValidPassword check password
func ValidPassword(pass string, userInputs ...string) error {
	if err := validation.Validate(pass,
		validation.Required,
		validation.Length(8, 1024),
	); err != nil {
		return ErrPasswordLength
	}
	// Limit check to first 64 chars for performance
	passLen := len(pass)
	if passLen > 64 {
		passLen = 64
	}
	if rating := zxcvbn.PasswordStrength(
		pass[:passLen],
		userInputs,
	); rating.Score <= 2 {
		return ErrPasswordStrength
	}
	return nil
}
