package user

import (
	"database/sql"
	"errors"
	"regexp"

	"github.com/allocamelus/allocamelus/internal/data"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jdinabox/goutils/logger"
)

var regexpInvalidChars = regexp.MustCompile(`^[^<>\[\]]*$`)

func initValidate(p data.Prepare) {
	preCheckUniqueName = p(`SELECT EXISTS(SELECT * FROM Users WHERE uniqueName=?)`)
	preCheckEmail = p(`SELECT EXISTS(SELECT * FROM Users WHERE email=?)`)
}

// ValidatePublic values
func (u *User) ValidatePublic() error {
	var errs validation.Errors
	if err := u.validUniqueName(); err != nil {
		errs["uniqueName"] = err
	}
	if err := u.validName(); err != nil {
		errs["name"] = err
	}
	if err := u.validEmail(); err != nil {
		errs["email"] = err
	}
	if err := u.validBio(); err != nil {
		errs["bio"] = err
	}
	return errs
}

var (
	// ErrUniqueNameLength 5 to 64 characters
	ErrUniqueNameLength = errors.New("user/validate: Unique Name length 5 to 64 characters")
	// ErrUniqueNameInvalid characters
	ErrUniqueNameInvalid = errors.New("user/validate: Unique Name invalid characters")
	// ErrUniqueNameTaken characters
	ErrUniqueNameTaken = errors.New("user/validate: Unique Name taken")
	regexUniqueName    = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`)
	preCheckUniqueName *sql.Stmt
)

func (u *User) validUniqueName() error {
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
	err := preCheckUniqueName.QueryRow(u.UniqueName).Scan(taken)
	if err != nil && err != sql.ErrNoRows {
		logger.Error(err)
	}
	if taken != 0 {
		return ErrUniqueNameTaken
	}

	return nil
}

var (
	// ErrNameLength 5 to 64 characters
	ErrNameLength = errors.New("user/validate: Name length 1 to 128 characters")
	// ErrNameInvalid characters
	ErrNameInvalid = errors.New("user/validate: Name invalid characters")
)

func (u *User) validName() error {
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
	ErrEmailInvalid = errors.New("user/validate: Invalid Email")
	// ErrEmailTaken characters
	ErrEmailTaken = errors.New("user/validate: Email taken")
	preCheckEmail *sql.Stmt
)

func (u *User) validEmail() error {
	if len(u.Email) != 0 {
		// Check Email
		if err := validation.Validate(u.Email, is.Email); err != nil {
			return ErrEmailInvalid
		}
		// Check Database for uniqueName
		var taken int8
		err := preCheckEmail.QueryRow(u.Email).Scan(taken)
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
	// ErrBioLength 5 to 64 characters
	ErrBioLength = errors.New("user/validate: Bio length 0 to 255 characters")
	// ErrBioInvalid characters
	ErrBioInvalid = errors.New("user/validate: Bio invalid characters")
)

func (u *User) validBio() error {
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
