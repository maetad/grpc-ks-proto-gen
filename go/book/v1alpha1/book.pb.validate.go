// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: book/v1alpha1/book.proto

package book

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ListBooksRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListBooksRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListBooksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListBooksRequestMultiError, or nil if none found.
func (m *ListBooksRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListBooksRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageSize

	// no validation rules for PageToken

	if len(errors) > 0 {
		return ListBooksRequestMultiError(errors)
	}

	return nil
}

// ListBooksRequestMultiError is an error wrapping multiple validation errors
// returned by ListBooksRequest.ValidateAll() if the designated constraints
// aren't met.
type ListBooksRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListBooksRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListBooksRequestMultiError) AllErrors() []error { return m }

// ListBooksRequestValidationError is the validation error returned by
// ListBooksRequest.Validate if the designated constraints aren't met.
type ListBooksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBooksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBooksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBooksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBooksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBooksRequestValidationError) ErrorName() string { return "ListBooksRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListBooksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBooksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBooksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBooksRequestValidationError{}

// Validate checks the field values on ListBooksResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListBooksResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListBooksResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListBooksResponseMultiError, or nil if none found.
func (m *ListBooksResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListBooksResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBooks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListBooksResponseValidationError{
						field:  fmt.Sprintf("Books[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListBooksResponseValidationError{
						field:  fmt.Sprintf("Books[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListBooksResponseValidationError{
					field:  fmt.Sprintf("Books[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return ListBooksResponseMultiError(errors)
	}

	return nil
}

// ListBooksResponseMultiError is an error wrapping multiple validation errors
// returned by ListBooksResponse.ValidateAll() if the designated constraints
// aren't met.
type ListBooksResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListBooksResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListBooksResponseMultiError) AllErrors() []error { return m }

// ListBooksResponseValidationError is the validation error returned by
// ListBooksResponse.Validate if the designated constraints aren't met.
type ListBooksResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBooksResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBooksResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBooksResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBooksResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBooksResponseValidationError) ErrorName() string {
	return "ListBooksResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListBooksResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBooksResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBooksResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBooksResponseValidationError{}

// Validate checks the field values on CreateBookRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateBookRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateBookRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateBookRequestMultiError, or nil if none found.
func (m *CreateBookRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateBookRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBook()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateBookRequestValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateBookRequestValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBook()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateBookRequestValidationError{
				field:  "Book",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateBookRequestMultiError(errors)
	}

	return nil
}

// CreateBookRequestMultiError is an error wrapping multiple validation errors
// returned by CreateBookRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateBookRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateBookRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateBookRequestMultiError) AllErrors() []error { return m }

// CreateBookRequestValidationError is the validation error returned by
// CreateBookRequest.Validate if the designated constraints aren't met.
type CreateBookRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateBookRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateBookRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateBookRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateBookRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateBookRequestValidationError) ErrorName() string {
	return "CreateBookRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateBookRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateBookRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateBookRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateBookRequestValidationError{}

// Validate checks the field values on CreateBookResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateBookResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateBookResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateBookResponseMultiError, or nil if none found.
func (m *CreateBookResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateBookResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBook()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateBookResponseValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateBookResponseValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBook()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateBookResponseValidationError{
				field:  "Book",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateBookResponseMultiError(errors)
	}

	return nil
}

// CreateBookResponseMultiError is an error wrapping multiple validation errors
// returned by CreateBookResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateBookResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateBookResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateBookResponseMultiError) AllErrors() []error { return m }

// CreateBookResponseValidationError is the validation error returned by
// CreateBookResponse.Validate if the designated constraints aren't met.
type CreateBookResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateBookResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateBookResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateBookResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateBookResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateBookResponseValidationError) ErrorName() string {
	return "CreateBookResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateBookResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateBookResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateBookResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateBookResponseValidationError{}

// Validate checks the field values on GetBookRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetBookRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetBookRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetBookRequestMultiError,
// or nil if none found.
func (m *GetBookRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetBookRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return GetBookRequestMultiError(errors)
	}

	return nil
}

// GetBookRequestMultiError is an error wrapping multiple validation errors
// returned by GetBookRequest.ValidateAll() if the designated constraints
// aren't met.
type GetBookRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetBookRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetBookRequestMultiError) AllErrors() []error { return m }

// GetBookRequestValidationError is the validation error returned by
// GetBookRequest.Validate if the designated constraints aren't met.
type GetBookRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBookRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBookRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBookRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBookRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBookRequestValidationError) ErrorName() string { return "GetBookRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetBookRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBookRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBookRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBookRequestValidationError{}

// Validate checks the field values on GetBookResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetBookResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetBookResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetBookResponseMultiError, or nil if none found.
func (m *GetBookResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetBookResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBook()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetBookResponseValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetBookResponseValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBook()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetBookResponseValidationError{
				field:  "Book",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetBookResponseMultiError(errors)
	}

	return nil
}

// GetBookResponseMultiError is an error wrapping multiple validation errors
// returned by GetBookResponse.ValidateAll() if the designated constraints
// aren't met.
type GetBookResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetBookResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetBookResponseMultiError) AllErrors() []error { return m }

// GetBookResponseValidationError is the validation error returned by
// GetBookResponse.Validate if the designated constraints aren't met.
type GetBookResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBookResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBookResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBookResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBookResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBookResponseValidationError) ErrorName() string { return "GetBookResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetBookResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBookResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBookResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBookResponseValidationError{}

// Validate checks the field values on UpdateBookRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateBookRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateBookRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateBookRequestMultiError, or nil if none found.
func (m *UpdateBookRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateBookRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBook()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateBookRequestValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateBookRequestValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBook()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateBookRequestValidationError{
				field:  "Book",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateBookRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateBookRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateBookRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateBookRequestMultiError(errors)
	}

	return nil
}

// UpdateBookRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateBookRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateBookRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateBookRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateBookRequestMultiError) AllErrors() []error { return m }

// UpdateBookRequestValidationError is the validation error returned by
// UpdateBookRequest.Validate if the designated constraints aren't met.
type UpdateBookRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateBookRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateBookRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateBookRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateBookRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateBookRequestValidationError) ErrorName() string {
	return "UpdateBookRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateBookRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateBookRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateBookRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateBookRequestValidationError{}

// Validate checks the field values on UpdateBookResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateBookResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateBookResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateBookResponseMultiError, or nil if none found.
func (m *UpdateBookResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateBookResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBook()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateBookResponseValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateBookResponseValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBook()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateBookResponseValidationError{
				field:  "Book",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateBookResponseMultiError(errors)
	}

	return nil
}

// UpdateBookResponseMultiError is an error wrapping multiple validation errors
// returned by UpdateBookResponse.ValidateAll() if the designated constraints
// aren't met.
type UpdateBookResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateBookResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateBookResponseMultiError) AllErrors() []error { return m }

// UpdateBookResponseValidationError is the validation error returned by
// UpdateBookResponse.Validate if the designated constraints aren't met.
type UpdateBookResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateBookResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateBookResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateBookResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateBookResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateBookResponseValidationError) ErrorName() string {
	return "UpdateBookResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateBookResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateBookResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateBookResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateBookResponseValidationError{}

// Validate checks the field values on DeleteBookRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteBookRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteBookRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteBookRequestMultiError, or nil if none found.
func (m *DeleteBookRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteBookRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return DeleteBookRequestMultiError(errors)
	}

	return nil
}

// DeleteBookRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteBookRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteBookRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteBookRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteBookRequestMultiError) AllErrors() []error { return m }

// DeleteBookRequestValidationError is the validation error returned by
// DeleteBookRequest.Validate if the designated constraints aren't met.
type DeleteBookRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteBookRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteBookRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteBookRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteBookRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteBookRequestValidationError) ErrorName() string {
	return "DeleteBookRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteBookRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteBookRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteBookRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteBookRequestValidationError{}

// Validate checks the field values on Book with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Book) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Book with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BookMultiError, or nil if none found.
func (m *Book) ValidateAll() error {
	return m.validate(true)
}

func (m *Book) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Title

	// no validation rules for Author

	// no validation rules for Isbn

	// no validation rules for Publisher

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDeletedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookValidationError{
				field:  "DeletedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BookMultiError(errors)
	}

	return nil
}

// BookMultiError is an error wrapping multiple validation errors returned by
// Book.ValidateAll() if the designated constraints aren't met.
type BookMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookMultiError) AllErrors() []error { return m }

// BookValidationError is the validation error returned by Book.Validate if the
// designated constraints aren't met.
type BookValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookValidationError) ErrorName() string { return "BookValidationError" }

// Error satisfies the builtin error interface
func (e BookValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBook.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookValidationError{}
