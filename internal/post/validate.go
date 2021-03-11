package post

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ErrContentLength max 65500
var ErrContentLength = errors.New("invalid-length-min0-max65500")

// ValidateContent is content valid
func ValidateContent(content string) error {
	if err := validation.Validate(content,
		validation.Length(0, 65500),
	); err != nil {
		return ErrContentLength
	}
	return nil
}

// ContentValid is html escaped content valid
func (p *Post) ContentValid() error {
	return ValidateContent(p.Content)
}
