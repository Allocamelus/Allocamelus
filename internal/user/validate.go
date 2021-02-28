package user

import (
	"database/sql"
	"errors"
	"regexp"

	"github.com/allocamelus/allocamelus/internal/data"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jdinabox/goutils/logger"
	"github.com/nbutton23/zxcvbn-go"
)

var regexpInvalidChars = regexp.MustCompile(`^[^<>\[\]]*$`)

func initValidate(p data.Prepare) {
	preCheckUniqueName = p(`SELECT EXISTS(SELECT * FROM Users WHERE uniqueName=?)`)
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
	if err := u.ValidUniqueName(); err != nil {
		errs["unique-name"] = err
	}
	if err := u.ValidName(); err != nil {
		errs["name"] = err
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
	// ErrUniqueNameLength 5 to 64 characters
	ErrUniqueNameLength = errors.New(invalidLength + "-min5-max64")
	// ErrUniqueNameInvalid characters
	ErrUniqueNameInvalid = errors.New(invalidChars)
	// ErrUniqueNameTaken characters
	ErrUniqueNameTaken = errors.New(taken)
	regexUniqueName    = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`)
	preCheckUniqueName *sql.Stmt
)

// ValidUniqueName Validate
func (u *User) ValidUniqueName() error {
	// Check Length
	if err := validation.Validate(u.UniqueName,
		validation.Required,
		validation.Length(5, 64),
	); err != nil {
		return ErrUniqueNameLength
	}
	// Check regex for Unique Name
	if !regexUniqueName.MatchString(u.UniqueName) {
		return ErrUniqueNameInvalid
	}
	// Check Database for uniqueName
	var taken int8
	err := preCheckUniqueName.QueryRow(u.UniqueName).Scan(&taken)
	if err != nil && err != sql.ErrNoRows {
		logger.Error(err)
	}
	if taken != 0 {
		return ErrUniqueNameTaken
	}

	return nil
}

var (
	// ErrNameLength 1 to 128 characters
	ErrNameLength = errors.New(invalidLength + "--min1-max128")
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
	if len(u.Email) != 0 {
		// Check Email
		if err := validation.Validate(u.Email, is.Email); err != nil {
			return ErrEmailInvalid
		}
		// Check Database for uniqueName
		var taken int8
		err := preCheckEmail.QueryRow(u.Email).Scan(&taken)
		if err != nil && err != sql.ErrNoRows {
			logger.Error(err)
		}
		if taken != 0 {
			return ErrEmailTaken
		}
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
		[]string{u.UniqueName, u.Name, u.Email},
	); rating.Score <= 2 {
		return ErrPasswordStrength
	}
	return nil
}
