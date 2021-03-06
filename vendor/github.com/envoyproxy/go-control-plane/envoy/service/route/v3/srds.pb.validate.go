// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/route/v3/srds.proto

package envoy_service_route_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _srds_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on SrdsDummy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SrdsDummy) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// SrdsDummyValidationError is the validation error returned by
// SrdsDummy.Validate if the designated constraints aren't met.
type SrdsDummyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SrdsDummyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SrdsDummyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SrdsDummyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SrdsDummyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SrdsDummyValidationError) ErrorName() string { return "SrdsDummyValidationError" }

// Error satisfies the builtin error interface
func (e SrdsDummyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSrdsDummy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SrdsDummyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SrdsDummyValidationError{}
