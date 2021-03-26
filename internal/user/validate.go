package user

import (
	"database/sql"
	"errors"
	"regexp"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/pkg/logger"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/nbutton23/zxcvbn-go"
)

var regexpInvalidChars = regexp.MustCompile(`^[^<>\[\]]*$`)

func initValidate(p data.Prepare) {
	preCheckUserName = p(`SELECT EXISTS(SELECT * FROM Users WHERE userName=?)`)
	preCheckEmail = p(`SELECT EXISTS(SELECT * FROM Users WHERE email=?)`)
}

const (
	invalidLength = "invalid-length"
	invalidChars  = "invalid-characters"
	invalidEmail  = "invalid-email"
	weakPassword  = "weak-password"
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
	// ErrUserNameInvalid characters
	ErrUserNameInvalid = errors.New(invalidChars)
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
		return ErrUserNameInvalid
	}
	// Check Database for userName
	var taken bool
	err := preCheckUserName.QueryRow(u.UserName).Scan(&taken)
	if err != nil && err != sql.ErrNoRows {
		logger.Error(err)
	}
	if taken {
		return ErrUserNameTaken
	}

	return nil
}

var (
	// ErrNameLength 1 to 128 characters
	ErrNameLength = errors.New(invalidLength + "-min1-max128")
	// ErrNameInvalid characters
	ErrNameInvalid = errors.New(invalidChars)
)

// ValidName Validate
func (u *User) ValidName() error {
	// Check Length
	if err := validation.Validate(u.Name,
		validation.Required,
		validation.Length(1, 128),
	); err != nil {
		return ErrNameLength
	}
	// Check regex for Invalid characters
	if !regexpInvalidChars.MatchString(u.Name) {
		return ErrNameInvalid
	}
	return nil
}

var (
	// ErrEmailInvalid Invalid Email
	ErrEmailInvalid = errors.New(invalidEmail)
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
	if len(u.Email) == 0 {
		return ErrEmailInvalid
	}
	// Check Database for userName
	var taken bool
	err := preCheckEmail.QueryRow(u.Email).Scan(&taken)
	if err != nil && err != sql.ErrNoRows {
		logger.Error(err)
	}
	if taken {
		return ErrEmailTaken
	}
	return nil
}

var (
	// ErrBioLength 0 to 255 characters
	ErrBioLength = errors.New(invalidLength + "-min0-max255")
	// ErrBioInvalid characters
	ErrBioInvalid = errors.New(invalidChars)
)

// ValidBio Validate
func (u *User) ValidBio() error {
	if len(u.Bio) != 0 {
		// Check length
		if err := validation.Validate(u.Bio, validation.Length(0, 255)); err != nil {
			return ErrBioLength
		}
		// Check regex for Invalid characters
		if !regexpInvalidChars.MatchString(u.Bio) {
			return ErrBioInvalid
		}
	}
	return nil
}

// ErrPasswordLength 8 to 1024
var (
	ErrPasswordLength   = errors.New(invalidLength + "-min8-max1024")
	ErrPasswordStrength = errors.New(weakPassword)
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
