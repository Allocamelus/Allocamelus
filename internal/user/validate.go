package user

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var regexpInvalidChars = regexp.MustCompile(`^[^<>\[\]]*$`)

// ValidatePublic values
func (u *User) ValidatePublic() (errs []error) {
	if err := u.validUniqueName(); err != nil {
		errs = append(errs, err)
	}
	if err := u.validName(); err != nil {
		errs = append(errs, err)
	}
	if err := u.validEmail(); err != nil {
		errs = append(errs, err)
	}
	if err := u.validBio(); err != nil {
		errs = append(errs, err)
	}
	return errs
}

var (
	// ErrUniqueNameLength 5 to 64 characters
	ErrUniqueNameLength = errors.New("user/validate: Unique Name 5 to 64 characters")
	// ErrUniqueNameInvalid characters
	ErrUniqueNameInvalid = errors.New("user/validate: Unique Name invalid characters")
	regexUniqueName      = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`)
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
	return nil
}

var (
	// ErrNameLength 5 to 64 characters
	ErrNameLength = errors.New("user/validate: Name 5 to 64 characters")
	// ErrNameInvalid characters
	ErrNameInvalid = errors.New("user/validate: Name invalid characters")
)

func (u *User) validName() error {
	// Check Length
	if err := validation.Validate(u.Name,
		validation.Required,
		validation.Length(3, 128),
	); err != nil {
		return ErrNameLength
	}
	// Check regex for Invalid characters
	if !regexpInvalidChars.MatchString(u.Name) {
		return ErrNameInvalid
	}
	return nil
}

// ErrEmailInvalid Invalid Email
var ErrEmailInvalid = errors.New("user/validate: Invalid Email")

func (u *User) validEmail() error {
	// Check Email
	if err := validation.Validate(u.Email, is.Email); err != nil {
		return ErrEmailInvalid
	}
	return nil
}

var (
	// ErrBioLength 5 to 64 characters
	ErrBioLength = errors.New("user/validate: Bio 0 to 255 characters")
	// ErrBioInvalid characters
	ErrBioInvalid = errors.New("user/validate: Bio invalid characters")
)

func (u *User) validBio() error {
	// Check length
	if err := validation.Validate(u.Bio, validation.Length(0, 255)); err != nil {
		return ErrBioLength
	}
	// Check regex for Invalid characters
	if !regexpInvalidChars.MatchString(u.Bio) {
		return ErrBioInvalid
	}
	return nil
}
