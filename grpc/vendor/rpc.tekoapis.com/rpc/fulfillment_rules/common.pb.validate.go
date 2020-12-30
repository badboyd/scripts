// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/fulfillment_rules/v1/common.proto

package fulfillment_rules

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
var _common_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Inventory with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Inventory) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Northern

	// no validation rules for Central

	// no validation rules for Southern

	return nil
}

// InventoryValidationError is the validation error returned by
// Inventory.Validate if the designated constraints aren't met.
type InventoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InventoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InventoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InventoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InventoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InventoryValidationError) ErrorName() string { return "InventoryValidationError" }

// Error satisfies the builtin error interface
func (e InventoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInventory.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InventoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InventoryValidationError{}

// Validate checks the field values on Seller with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Seller) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return SellerValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	if m.GetSellerCenterId() <= 0 {
		return SellerValidationError{
			field:  "SellerCenterId",
			reason: "value must be greater than 0",
		}
	}

	if utf8.RuneCountInString(m.GetCode()) < 1 {
		return SellerValidationError{
			field:  "Code",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for IsAllowAutoProcess

	// no validation rules for IsImportWithoutStore

	// no validation rules for FirebasePath

	if m.GetDefaultInventory() == nil {
		return SellerValidationError{
			field:  "DefaultInventory",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDefaultInventory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SellerValidationError{
				field:  "DefaultInventory",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PaymentEpsilon

	// no validation rules for IsSkipCoupon

	// no validation rules for IsAllocateOrderPromotion

	// no validation rules for IsAllocateNoneOrderPromotion

	// no validation rules for IsApplyRelativeAllocation

	return nil
}

// SellerValidationError is the validation error returned by Seller.Validate if
// the designated constraints aren't met.
type SellerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SellerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SellerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SellerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SellerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SellerValidationError) ErrorName() string { return "SellerValidationError" }

// Error satisfies the builtin error interface
func (e SellerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSeller.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SellerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SellerValidationError{}

// Validate checks the field values on Address with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Address) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ApartmentNumber

	// no validation rules for Street

	// no validation rules for Ward

	// no validation rules for District

	if utf8.RuneCountInString(m.GetCity()) < 1 {
		return AddressValidationError{
			field:  "City",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetProvince()) < 1 {
		return AddressValidationError{
			field:  "Province",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// AddressValidationError is the validation error returned by Address.Validate
// if the designated constraints aren't met.
type AddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressValidationError) ErrorName() string { return "AddressValidationError" }

// Error satisfies the builtin error interface
func (e AddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressValidationError{}

// Validate checks the field values on ShippingAddress with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ShippingAddress) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return ShippingAddressValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return ShippingAddressValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if !_ShippingAddress_Phone_Pattern.MatchString(m.GetPhone()) {
		return ShippingAddressValidationError{
			field:  "Phone",
			reason: "value does not match regex pattern \"^(0|\\\\+)[1-9](\\\\d{8}|\\\\d{10})$\"",
		}
	}

	// no validation rules for Note

	if m.GetAddress() == nil {
		return ShippingAddressValidationError{
			field:  "Address",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ShippingAddressValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ShippingAddressValidationError is the validation error returned by
// ShippingAddress.Validate if the designated constraints aren't met.
type ShippingAddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShippingAddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShippingAddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShippingAddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShippingAddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShippingAddressValidationError) ErrorName() string { return "ShippingAddressValidationError" }

// Error satisfies the builtin error interface
func (e ShippingAddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShippingAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShippingAddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShippingAddressValidationError{}

var _ShippingAddress_Phone_Pattern = regexp.MustCompile("^(0|\\+)[1-9](\\d{8}|\\d{10})$")

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *User) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return UserValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for Name

	// no validation rules for Phone

	// no validation rules for Email

	for idx, item := range m.GetShippingAddress() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserValidationError{
					field:  fmt.Sprintf("ShippingAddress[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on ShippingFeeRule with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ShippingFeeRule) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Rule

	// no validation rules for Fee

	// no validation rules for Priority

	// no validation rules for Comment

	return nil
}

// ShippingFeeRuleValidationError is the validation error returned by
// ShippingFeeRule.Validate if the designated constraints aren't met.
type ShippingFeeRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShippingFeeRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShippingFeeRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShippingFeeRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShippingFeeRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShippingFeeRuleValidationError) ErrorName() string { return "ShippingFeeRuleValidationError" }

// Error satisfies the builtin error interface
func (e ShippingFeeRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShippingFeeRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShippingFeeRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShippingFeeRuleValidationError{}
