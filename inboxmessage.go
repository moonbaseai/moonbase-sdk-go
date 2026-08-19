// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/internal/apiquery"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/packages/pagination"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// Manage your inboxes, conversations, and messages
//
// InboxMessageService contains methods and other services that help with
// interacting with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxMessageService] method instead.
type InboxMessageService struct {
	Options []option.RequestOption
	// Manage your inboxes, conversations, and messages
	Attachments InboxMessageAttachmentService
}

// NewInboxMessageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboxMessageService(opts ...option.RequestOption) (r InboxMessageService) {
	r = InboxMessageService{}
	r.Options = opts
	r.Attachments = NewInboxMessageAttachmentService(opts...)
	return
}

// Creates a new message draft.
func (r *InboxMessageService) New(ctx context.Context, body InboxMessageNewParams, opts ...option.RequestOption) (res *InboxMessageNewResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "inbox_messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the details of an existing message.
func (r *InboxMessageService) Get(ctx context.Context, id string, query InboxMessageGetParams, opts ...option.RequestOption) (res *InboxMessageGetResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("inbox_messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates an existing message draft.
func (r *InboxMessageService) Update(ctx context.Context, id string, body InboxMessageUpdateParams, opts ...option.RequestOption) (res *InboxMessageUpdateResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("inbox_messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a list of messages.
func (r *InboxMessageService) List(ctx context.Context, query InboxMessageListParams, opts ...option.RequestOption) (res *pagination.CursorPage[MessagePointer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbox_messages"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of messages.
func (r *InboxMessageService) ListAutoPaging(ctx context.Context, query InboxMessageListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[MessagePointer] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a message draft.
func (r *InboxMessageService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("inbox_messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// The Email Message object represents a single email within a `Conversation`.
type EmailMessage struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Structured content that can be rendered in multiple formats, currently
	// supporting Markdown.
	Body shared.FormattedText `json:"body" api:"required"`
	// `true` if the message appears to be part of a bulk mailing.
	Bulk bool `json:"bulk" api:"required"`
	// The time the message was received, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// `true` if the message is a draft that has not been sent.
	Draft bool `json:"draft" api:"required"`
	// The current lock version of the message for optimistic concurrency control.
	LockVersion int64 `json:"lock_version" api:"required"`
	// `true` if the message is classified as spam.
	Spam bool `json:"spam" api:"required"`
	// The subject line of the email.
	Subject string `json:"subject" api:"required"`
	// `true` if the message is in the trash.
	Trash bool `json:"trash" api:"required"`
	// String representing the object’s type. Always `email_message` for this object.
	Type constant.EmailMessage `json:"type" default:"email_message"`
	// `true` if the message has not been read.
	Unread bool `json:"unread" api:"required"`
	// A list of `Address` objects associated with the message (sender and recipients).
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Addresses []EmailMessageAddress `json:"addresses"`
	// A list of `Attachment` objects on the message.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Attachments []MessageAttachment `json:"attachments"`
	// The `Conversation` thread this message is part of.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Conversation *InboxConversation `json:"conversation"`
	// A concise, system-generated summary of the email content.
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Body         respjson.Field
		Bulk         respjson.Field
		CreatedAt    respjson.Field
		Draft        respjson.Field
		LockVersion  respjson.Field
		Spam         respjson.Field
		Subject      respjson.Field
		Trash        respjson.Field
		Type         respjson.Field
		Unread       respjson.Field
		Addresses    respjson.Field
		Attachments  respjson.Field
		Conversation respjson.Field
		Summary      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessage) RawJSON() string { return r.JSON.raw }
func (r *EmailMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The EmailMessageAddress object represents a recipient or sender of a message. It
// contains an email address and can be linked to a person and an organization in
// your collections.
type EmailMessageAddress struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// The email address.
	Email string `json:"email" api:"required" format:"email"`
	// The role of the address in the message. Can be `from`, `reply_to`, `to`, `cc`,
	// or `bcc`.
	//
	// Any of "from", "reply_to", "to", "cc", "bcc".
	Role EmailMessageAddressRole `json:"role" api:"required"`
	// String representing the object’s type. Always `message_address` for this object.
	Type constant.EmailMessageAddress `json:"type" default:"email_message_address"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Organization ItemPointer `json:"organization"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Person ItemPointer `json:"person"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Email        respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		Organization respjson.Field
		Person       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageAddress) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The role of the address in the message. Can be `from`, `reply_to`, `to`, `cc`,
// or `bcc`.
type EmailMessageAddressRole string

const (
	EmailMessageAddressRoleFrom    EmailMessageAddressRole = "from"
	EmailMessageAddressRoleReplyTo EmailMessageAddressRole = "reply_to"
	EmailMessageAddressRoleTo      EmailMessageAddressRole = "to"
	EmailMessageAddressRoleCc      EmailMessageAddressRole = "cc"
	EmailMessageAddressRoleBcc     EmailMessageAddressRole = "bcc"
)

// The property Email is required.
type EmailMessageAddressParams struct {
	// The email address.
	Email string `json:"email" api:"required" format:"email"`
	// The recipient's name.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r EmailMessageAddressParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailMessageAddressParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailMessageAddressParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Attachment object represents a file attached to a message. You can download
// the file content via the `download_url`.
type MessageAttachment struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A temporary, signed URL to download the file content. The URL expires after one
	// hour.
	DownloadURL string `json:"download_url" api:"required" format:"uri"`
	// The original name of the uploaded file, including its extension.
	Filename string `json:"filename" api:"required"`
	// The size of the file in bytes.
	Size int64 `json:"size" api:"required"`
	// String representing the object’s type. Always `message_attachment` for this
	// object.
	Type constant.MessageAttachment `json:"type" default:"message_attachment"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DownloadURL respjson.Field
		Filename    respjson.Field
		Size        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageAttachment) RawJSON() string { return r.JSON.raw }
func (r *MessageAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagePointer struct {
	ID   string           `json:"id" api:"required"`
	Type constant.Message `json:"type" default:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagePointer) RawJSON() string { return r.JSON.raw }
func (r *MessagePointer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Slack Message object represents a single Slack post within a `Conversation`.
type SlackMessage struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Structured content that can be rendered in multiple formats, currently
	// supporting Markdown.
	Body shared.FormattedText `json:"body" api:"required"`
	// `true` if the message appears to be part of a bulk mailing.
	Bulk bool `json:"bulk" api:"required"`
	// The time the message was received, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// `true` if the message is a draft that has not been sent.
	Draft bool `json:"draft" api:"required"`
	// The current lock version of the message for optimistic concurrency control.
	LockVersion int64 `json:"lock_version" api:"required"`
	// `true` if the message is classified as spam.
	Spam bool `json:"spam" api:"required"`
	// The subject line of the message (for messages received from Slack, this is a
	// snippet of the message; for messages sent to Slack, it can be set, but is not
	// sent to Slack).
	Subject string `json:"subject" api:"required"`
	// `true` if the message is in the trash.
	Trash bool `json:"trash" api:"required"`
	// String representing the object’s type. Always `slack_message` for this object.
	Type constant.SlackMessage `json:"type" default:"slack_message"`
	// `true` if the message has not been read.
	Unread bool `json:"unread" api:"required"`
	// A list of `SlackMessageAddress` objects associated with the message (sender and
	// recipients).
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Addresses []SlackMessageAddressUnion `json:"addresses"`
	// A list of `Attachment` objects on the message.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Attachments []MessageAttachment `json:"attachments"`
	// The `Conversation` thread this message is part of.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Conversation *InboxConversation `json:"conversation"`
	// A concise, system-generated summary of the message content.
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Body         respjson.Field
		Bulk         respjson.Field
		CreatedAt    respjson.Field
		Draft        respjson.Field
		LockVersion  respjson.Field
		Spam         respjson.Field
		Subject      respjson.Field
		Trash        respjson.Field
		Type         respjson.Field
		Unread       respjson.Field
		Addresses    respjson.Field
		Attachments  respjson.Field
		Conversation respjson.Field
		Summary      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SlackMessage) RawJSON() string { return r.JSON.raw }
func (r *SlackMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SlackMessageAddressUnion contains all possible properties and values from
// [SlackMessageAddressSlackMessageChannelAddress],
// [SlackMessageAddressSlackMessageUserAddress].
//
// Use the [SlackMessageAddressUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SlackMessageAddressUnion struct {
	ID         string `json:"id"`
	ProviderID string `json:"provider_id"`
	Role       string `json:"role"`
	// Any of "slack_message_channel_address", "slack_message_user_address".
	Type string `json:"type"`
	// This field is from variant [SlackMessageAddressSlackMessageChannelAddress].
	Organization ItemPointer `json:"organization"`
	// This field is from variant [SlackMessageAddressSlackMessageChannelAddress].
	Person ItemPointer `json:"person"`
	JSON   struct {
		ID           respjson.Field
		ProviderID   respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		Organization respjson.Field
		Person       respjson.Field
		raw          string
	} `json:"-"`
}

// anySlackMessageAddress is implemented by each variant of
// [SlackMessageAddressUnion] to add type safety for the return type of
// [SlackMessageAddressUnion.AsAny]
type anySlackMessageAddress interface {
	implSlackMessageAddressUnion()
}

func (SlackMessageAddressSlackMessageChannelAddress) implSlackMessageAddressUnion() {}
func (SlackMessageAddressSlackMessageUserAddress) implSlackMessageAddressUnion()    {}

// Use the following switch statement to find the correct variant
//
//	switch variant := SlackMessageAddressUnion.AsAny().(type) {
//	case moonbase.SlackMessageAddressSlackMessageChannelAddress:
//	case moonbase.SlackMessageAddressSlackMessageUserAddress:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u SlackMessageAddressUnion) AsAny() anySlackMessageAddress {
	switch u.Type {
	case "slack_message_channel_address":
		return u.AsSlackMessageChannelAddress()
	case "slack_message_user_address":
		return u.AsSlackMessageUserAddress()
	}
	return nil
}

func (u SlackMessageAddressUnion) AsSlackMessageChannelAddress() (v SlackMessageAddressSlackMessageChannelAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SlackMessageAddressUnion) AsSlackMessageUserAddress() (v SlackMessageAddressSlackMessageUserAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SlackMessageAddressUnion) RawJSON() string { return u.JSON.raw }

func (r *SlackMessageAddressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The SlackMessageChannelAddress object represents a Slack channels address on a
// message. It contains a Slack Channel ID and can be linked to a person and an
// organization in your collections.
type SlackMessageAddressSlackMessageChannelAddress struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// The Slack Channel ID.
	ProviderID string `json:"provider_id" api:"required"`
	// The role of the address in the message. Can be `from`, `reply_to`, `to`, `cc`,
	// or `bcc`.
	//
	// Any of "from", "to", "cc", "bcc".
	Role string `json:"role" api:"required"`
	// String representing the object’s type. Always `slack_message_channel_address`
	// for this object.
	Type constant.SlackMessageChannelAddress `json:"type" default:"slack_message_channel_address"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Organization ItemPointer `json:"organization"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Person ItemPointer `json:"person"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ProviderID   respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		Organization respjson.Field
		Person       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SlackMessageAddressSlackMessageChannelAddress) RawJSON() string { return r.JSON.raw }
func (r *SlackMessageAddressSlackMessageChannelAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The SlackMessageUserAddress object represents a Slack user address on a message.
// It contains a Slack User ID and can be linked to a person and an organization in
// your collections.
type SlackMessageAddressSlackMessageUserAddress struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// The Slack User ID
	ProviderID string `json:"provider_id" api:"required"`
	// The role of the address in the message. Can be `from`, `reply_to`, `to`, `cc`,
	// or `bcc`.
	//
	// Any of "from", "to", "cc", "bcc".
	Role string `json:"role" api:"required"`
	// String representing the object’s type. Always `slack_message_user_address` for
	// this object.
	Type constant.SlackMessageUserAddress `json:"type" default:"slack_message_user_address"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Organization ItemPointer `json:"organization"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Person ItemPointer `json:"person"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ProviderID   respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		Organization respjson.Field
		Person       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SlackMessageAddressSlackMessageUserAddress) RawJSON() string { return r.JSON.raw }
func (r *SlackMessageAddressSlackMessageUserAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ProviderID, Type are required.
type SlackMessageAddressParams struct {
	// The Slack channel ID.
	ProviderID string `json:"provider_id" api:"required"`
	// The channel name name.
	Name param.Opt[string] `json:"name,omitzero"`
	// This field can be elided, and will marshal its zero value as "slack_channel".
	Type constant.SlackChannel `json:"type" default:"slack_channel"`
	paramObj
}

func (r SlackMessageAddressParams) MarshalJSON() (data []byte, err error) {
	type shadow SlackMessageAddressParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SlackMessageAddressParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InboxMessageNewResponseUnion contains all possible properties and values from
// [EmailMessage], [SlackMessage].
//
// Use the [InboxMessageNewResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InboxMessageNewResponseUnion struct {
	ID string `json:"id"`
	// This field is from variant [EmailMessage].
	Body        shared.FormattedText `json:"body"`
	Bulk        bool                 `json:"bulk"`
	CreatedAt   time.Time            `json:"created_at"`
	Draft       bool                 `json:"draft"`
	LockVersion int64                `json:"lock_version"`
	Spam        bool                 `json:"spam"`
	Subject     string               `json:"subject"`
	Trash       bool                 `json:"trash"`
	// Any of "email_message", "slack_message".
	Type   string `json:"type"`
	Unread bool   `json:"unread"`
	// This field is a union of [[]EmailMessageAddress], [[]SlackMessageAddressUnion]
	Addresses   InboxMessageNewResponseUnionAddresses `json:"addresses"`
	Attachments []MessageAttachment                   `json:"attachments"`
	// This field is from variant [EmailMessage].
	Conversation InboxConversation `json:"conversation"`
	Summary      string            `json:"summary"`
	JSON         struct {
		ID           respjson.Field
		Body         respjson.Field
		Bulk         respjson.Field
		CreatedAt    respjson.Field
		Draft        respjson.Field
		LockVersion  respjson.Field
		Spam         respjson.Field
		Subject      respjson.Field
		Trash        respjson.Field
		Type         respjson.Field
		Unread       respjson.Field
		Addresses    respjson.Field
		Attachments  respjson.Field
		Conversation respjson.Field
		Summary      respjson.Field
		raw          string
	} `json:"-"`
}

// anyInboxMessageNewResponse is implemented by each variant of
// [InboxMessageNewResponseUnion] to add type safety for the return type of
// [InboxMessageNewResponseUnion.AsAny]
type anyInboxMessageNewResponse interface {
	implInboxMessageNewResponseUnion()
}

func (EmailMessage) implInboxMessageNewResponseUnion() {}
func (SlackMessage) implInboxMessageNewResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := InboxMessageNewResponseUnion.AsAny().(type) {
//	case moonbase.EmailMessage:
//	case moonbase.SlackMessage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u InboxMessageNewResponseUnion) AsAny() anyInboxMessageNewResponse {
	switch u.Type {
	case "email_message":
		return u.AsEmailMessage()
	case "slack_message":
		return u.AsSlackMessage()
	}
	return nil
}

func (u InboxMessageNewResponseUnion) AsEmailMessage() (v EmailMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InboxMessageNewResponseUnion) AsSlackMessage() (v SlackMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InboxMessageNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *InboxMessageNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InboxMessageNewResponseUnionAddresses is an implicit subunion of
// [InboxMessageNewResponseUnion]. InboxMessageNewResponseUnionAddresses provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [InboxMessageNewResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfEmailMessageAddressArray OfSlackMessageAddressArray]
type InboxMessageNewResponseUnionAddresses struct {
	// This field will be present if the value is a [[]EmailMessageAddress] instead of
	// an object.
	OfEmailMessageAddressArray []EmailMessageAddress `json:",inline"`
	// This field will be present if the value is a [[]SlackMessageAddressUnion]
	// instead of an object.
	OfSlackMessageAddressArray []SlackMessageAddressUnion `json:",inline"`
	JSON                       struct {
		OfEmailMessageAddressArray respjson.Field
		OfSlackMessageAddressArray respjson.Field
		raw                        string
	} `json:"-"`
}

func (r *InboxMessageNewResponseUnionAddresses) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InboxMessageGetResponseUnion contains all possible properties and values from
// [EmailMessage], [SlackMessage].
//
// Use the [InboxMessageGetResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InboxMessageGetResponseUnion struct {
	ID string `json:"id"`
	// This field is from variant [EmailMessage].
	Body        shared.FormattedText `json:"body"`
	Bulk        bool                 `json:"bulk"`
	CreatedAt   time.Time            `json:"created_at"`
	Draft       bool                 `json:"draft"`
	LockVersion int64                `json:"lock_version"`
	Spam        bool                 `json:"spam"`
	Subject     string               `json:"subject"`
	Trash       bool                 `json:"trash"`
	// Any of "email_message", "slack_message".
	Type   string `json:"type"`
	Unread bool   `json:"unread"`
	// This field is a union of [[]EmailMessageAddress], [[]SlackMessageAddressUnion]
	Addresses   InboxMessageGetResponseUnionAddresses `json:"addresses"`
	Attachments []MessageAttachment                   `json:"attachments"`
	// This field is from variant [EmailMessage].
	Conversation InboxConversation `json:"conversation"`
	Summary      string            `json:"summary"`
	JSON         struct {
		ID           respjson.Field
		Body         respjson.Field
		Bulk         respjson.Field
		CreatedAt    respjson.Field
		Draft        respjson.Field
		LockVersion  respjson.Field
		Spam         respjson.Field
		Subject      respjson.Field
		Trash        respjson.Field
		Type         respjson.Field
		Unread       respjson.Field
		Addresses    respjson.Field
		Attachments  respjson.Field
		Conversation respjson.Field
		Summary      respjson.Field
		raw          string
	} `json:"-"`
}

// anyInboxMessageGetResponse is implemented by each variant of
// [InboxMessageGetResponseUnion] to add type safety for the return type of
// [InboxMessageGetResponseUnion.AsAny]
type anyInboxMessageGetResponse interface {
	implInboxMessageGetResponseUnion()
}

func (EmailMessage) implInboxMessageGetResponseUnion() {}
func (SlackMessage) implInboxMessageGetResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := InboxMessageGetResponseUnion.AsAny().(type) {
//	case moonbase.EmailMessage:
//	case moonbase.SlackMessage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u InboxMessageGetResponseUnion) AsAny() anyInboxMessageGetResponse {
	switch u.Type {
	case "email_message":
		return u.AsEmailMessage()
	case "slack_message":
		return u.AsSlackMessage()
	}
	return nil
}

func (u InboxMessageGetResponseUnion) AsEmailMessage() (v EmailMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InboxMessageGetResponseUnion) AsSlackMessage() (v SlackMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InboxMessageGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *InboxMessageGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InboxMessageGetResponseUnionAddresses is an implicit subunion of
// [InboxMessageGetResponseUnion]. InboxMessageGetResponseUnionAddresses provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [InboxMessageGetResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfEmailMessageAddressArray OfSlackMessageAddressArray]
type InboxMessageGetResponseUnionAddresses struct {
	// This field will be present if the value is a [[]EmailMessageAddress] instead of
	// an object.
	OfEmailMessageAddressArray []EmailMessageAddress `json:",inline"`
	// This field will be present if the value is a [[]SlackMessageAddressUnion]
	// instead of an object.
	OfSlackMessageAddressArray []SlackMessageAddressUnion `json:",inline"`
	JSON                       struct {
		OfEmailMessageAddressArray respjson.Field
		OfSlackMessageAddressArray respjson.Field
		raw                        string
	} `json:"-"`
}

func (r *InboxMessageGetResponseUnionAddresses) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InboxMessageUpdateResponseUnion contains all possible properties and values from
// [EmailMessage], [SlackMessage].
//
// Use the [InboxMessageUpdateResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InboxMessageUpdateResponseUnion struct {
	ID string `json:"id"`
	// This field is from variant [EmailMessage].
	Body        shared.FormattedText `json:"body"`
	Bulk        bool                 `json:"bulk"`
	CreatedAt   time.Time            `json:"created_at"`
	Draft       bool                 `json:"draft"`
	LockVersion int64                `json:"lock_version"`
	Spam        bool                 `json:"spam"`
	Subject     string               `json:"subject"`
	Trash       bool                 `json:"trash"`
	// Any of "email_message", "slack_message".
	Type   string `json:"type"`
	Unread bool   `json:"unread"`
	// This field is a union of [[]EmailMessageAddress], [[]SlackMessageAddressUnion]
	Addresses   InboxMessageUpdateResponseUnionAddresses `json:"addresses"`
	Attachments []MessageAttachment                      `json:"attachments"`
	// This field is from variant [EmailMessage].
	Conversation InboxConversation `json:"conversation"`
	Summary      string            `json:"summary"`
	JSON         struct {
		ID           respjson.Field
		Body         respjson.Field
		Bulk         respjson.Field
		CreatedAt    respjson.Field
		Draft        respjson.Field
		LockVersion  respjson.Field
		Spam         respjson.Field
		Subject      respjson.Field
		Trash        respjson.Field
		Type         respjson.Field
		Unread       respjson.Field
		Addresses    respjson.Field
		Attachments  respjson.Field
		Conversation respjson.Field
		Summary      respjson.Field
		raw          string
	} `json:"-"`
}

// anyInboxMessageUpdateResponse is implemented by each variant of
// [InboxMessageUpdateResponseUnion] to add type safety for the return type of
// [InboxMessageUpdateResponseUnion.AsAny]
type anyInboxMessageUpdateResponse interface {
	implInboxMessageUpdateResponseUnion()
}

func (EmailMessage) implInboxMessageUpdateResponseUnion() {}
func (SlackMessage) implInboxMessageUpdateResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := InboxMessageUpdateResponseUnion.AsAny().(type) {
//	case moonbase.EmailMessage:
//	case moonbase.SlackMessage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u InboxMessageUpdateResponseUnion) AsAny() anyInboxMessageUpdateResponse {
	switch u.Type {
	case "email_message":
		return u.AsEmailMessage()
	case "slack_message":
		return u.AsSlackMessage()
	}
	return nil
}

func (u InboxMessageUpdateResponseUnion) AsEmailMessage() (v EmailMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InboxMessageUpdateResponseUnion) AsSlackMessage() (v SlackMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InboxMessageUpdateResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *InboxMessageUpdateResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InboxMessageUpdateResponseUnionAddresses is an implicit subunion of
// [InboxMessageUpdateResponseUnion]. InboxMessageUpdateResponseUnionAddresses
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [InboxMessageUpdateResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfEmailMessageAddressArray OfSlackMessageAddressArray]
type InboxMessageUpdateResponseUnionAddresses struct {
	// This field will be present if the value is a [[]EmailMessageAddress] instead of
	// an object.
	OfEmailMessageAddressArray []EmailMessageAddress `json:",inline"`
	// This field will be present if the value is a [[]SlackMessageAddressUnion]
	// instead of an object.
	OfSlackMessageAddressArray []SlackMessageAddressUnion `json:",inline"`
	JSON                       struct {
		OfEmailMessageAddressArray respjson.Field
		OfSlackMessageAddressArray respjson.Field
		raw                        string
	} `json:"-"`
}

func (r *InboxMessageUpdateResponseUnionAddresses) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxMessageNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a draft in a new conversation.
	OfEmailMessageNewConversationCreates *InboxMessageNewParamsMessageEmailMessageNewConversationCreateParams `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a draft in a new conversation.
	OfSlackMessageNewConversationCreates *InboxMessageNewParamsMessageSlackMessageNewConversationCreateParams `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a draft reply in an existing conversation.
	OfEmailMessageReplyCreates *InboxMessageNewParamsMessageEmailMessageReplyCreateParams `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a draft reply in an existing conversation.
	OfSlackMessageReplyCreates *InboxMessageNewParamsMessageSlackMessageReplyCreateParams `json:",inline"`

	paramObj
}

func (u InboxMessageNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEmailMessageNewConversationCreates, u.OfSlackMessageNewConversationCreates, u.OfEmailMessageReplyCreates, u.OfSlackMessageReplyCreates)
}
func (r *InboxMessageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for creating a draft in a new conversation.
//
// The properties Body, InboxID, Subject, To, Type are required.
type InboxMessageNewParamsMessageEmailMessageNewConversationCreateParams struct {
	// The email body.
	Body shared.FormattedTextParam `json:"body,omitzero" api:"required"`
	// The inbox to use for sending the email.
	InboxID string `json:"inbox_id" api:"required"`
	// The subject line of the email.
	Subject string `json:"subject" api:"required"`
	// A list of recipients.
	To []EmailMessageAddressParams `json:"to,omitzero" api:"required"`
	// A list of the BCC recipients.
	Bcc []EmailMessageAddressParams `json:"bcc,omitzero"`
	// A list of the CC recipients.
	Cc []EmailMessageAddressParams `json:"cc,omitzero"`
	// This field can be elided, and will marshal its zero value as "email_message".
	Type constant.EmailMessage `json:"type" default:"email_message"`
	paramObj
}

func (r InboxMessageNewParamsMessageEmailMessageNewConversationCreateParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxMessageNewParamsMessageEmailMessageNewConversationCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxMessageNewParamsMessageEmailMessageNewConversationCreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for creating a draft in a new conversation.
//
// The properties Body, InboxID, Subject, To, Type are required.
type InboxMessageNewParamsMessageSlackMessageNewConversationCreateParams struct {
	// The message body.
	Body shared.FormattedTextParam `json:"body,omitzero" api:"required"`
	// The inbox to use for sending the Slack message.
	InboxID string `json:"inbox_id" api:"required"`
	// The subject line of the conversation (not included in actual Slack message).
	Subject string `json:"subject" api:"required"`
	// The Slack channel to post the message in.
	To []SlackMessageAddressParams `json:"to,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "slack_message".
	Type constant.SlackMessage `json:"type" default:"slack_message"`
	paramObj
}

func (r InboxMessageNewParamsMessageSlackMessageNewConversationCreateParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxMessageNewParamsMessageSlackMessageNewConversationCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxMessageNewParamsMessageSlackMessageNewConversationCreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for creating a draft reply in an existing conversation.
//
// The properties Body, ConversationID, InboxID, Type are required.
type InboxMessageNewParamsMessageEmailMessageReplyCreateParams struct {
	// The email body.
	Body shared.FormattedTextParam `json:"body,omitzero" api:"required"`
	// The ID of the conversation to reply to.
	ConversationID string `json:"conversation_id" api:"required"`
	// The inbox to use for sending the email.
	InboxID string `json:"inbox_id" api:"required"`
	// A list of the BCC recipients.
	Bcc []EmailMessageAddressParams `json:"bcc,omitzero"`
	// A list of the CC recipients.
	Cc []EmailMessageAddressParams `json:"cc,omitzero"`
	// A list of recipients. If omitted, recipients are derived from the conversation.
	To []EmailMessageAddressParams `json:"to,omitzero"`
	// This field can be elided, and will marshal its zero value as "email_message".
	Type constant.EmailMessage `json:"type" default:"email_message"`
	paramObj
}

func (r InboxMessageNewParamsMessageEmailMessageReplyCreateParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxMessageNewParamsMessageEmailMessageReplyCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxMessageNewParamsMessageEmailMessageReplyCreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for creating a draft reply in an existing conversation.
//
// The properties Body, ConversationID, InboxID, Type are required.
type InboxMessageNewParamsMessageSlackMessageReplyCreateParams struct {
	// The message body.
	Body shared.FormattedTextParam `json:"body,omitzero" api:"required"`
	// The ID of the conversation to reply to.
	ConversationID string `json:"conversation_id" api:"required"`
	// The inbox to use for sending the Slack message.
	InboxID string `json:"inbox_id" api:"required"`
	// The Slack channel to post the message in.
	To []SlackMessageAddressParams `json:"to,omitzero"`
	// This field can be elided, and will marshal its zero value as "slack_message".
	Type constant.SlackMessage `json:"type" default:"slack_message"`
	paramObj
}

func (r InboxMessageNewParamsMessageSlackMessageReplyCreateParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxMessageNewParamsMessageSlackMessageReplyCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxMessageNewParamsMessageSlackMessageReplyCreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxMessageGetParams struct {
	// Specifies which related objects to include in the response. Valid options are
	// `addresses`, `attachments`, and `conversation`.
	//
	// Any of "addresses", "attachments", "conversation".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxMessageGetParams]'s query parameters as `url.Values`.
func (r InboxMessageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxMessageUpdateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a draft message in an existing conversation.
	OfEmailMessageUpdates *InboxMessageUpdateParamsMessageEmailMessageUpdateParams `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a draft message in an existing conversation.
	OfSlackMessageUpdates *InboxMessageUpdateParamsMessageSlackMessageUpdateParams `json:",inline"`

	paramObj
}

func (u InboxMessageUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEmailMessageUpdates, u.OfSlackMessageUpdates)
}
func (r *InboxMessageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for updating a draft message in an existing conversation.
//
// The properties LockVersion, Type are required.
type InboxMessageUpdateParamsMessageEmailMessageUpdateParams struct {
	// The current lock version of the draft for optimistic concurrency control.
	LockVersion int64 `json:"lock_version" api:"required"`
	// The subject line of the email.
	Subject param.Opt[string] `json:"subject,omitzero"`
	// A list of the BCC recipients.
	Bcc []EmailMessageAddressParams `json:"bcc,omitzero"`
	// The email body.
	Body shared.FormattedTextParam `json:"body,omitzero"`
	// A list of the CC recipients.
	Cc []EmailMessageAddressParams `json:"cc,omitzero"`
	// A list of the recipients.
	To []EmailMessageAddressParams `json:"to,omitzero"`
	// This field can be elided, and will marshal its zero value as "email_message".
	Type constant.EmailMessage `json:"type" default:"email_message"`
	paramObj
}

func (r InboxMessageUpdateParamsMessageEmailMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxMessageUpdateParamsMessageEmailMessageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxMessageUpdateParamsMessageEmailMessageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for updating a draft message in an existing conversation.
//
// The properties LockVersion, Type are required.
type InboxMessageUpdateParamsMessageSlackMessageUpdateParams struct {
	// The current lock version of the draft for optimistic concurrency control.
	LockVersion int64 `json:"lock_version" api:"required"`
	// The subject line of the conversation (not included in actual Slack message).
	Subject param.Opt[string] `json:"subject,omitzero"`
	// The message body.
	Body shared.FormattedTextParam `json:"body,omitzero"`
	// The Slack channel to post the message in.
	To []SlackMessageAddressParams `json:"to,omitzero"`
	// This field can be elided, and will marshal its zero value as "slack_message".
	Type constant.SlackMessage `json:"type" default:"slack_message"`
	paramObj
}

func (r InboxMessageUpdateParamsMessageSlackMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxMessageUpdateParamsMessageSlackMessageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxMessageUpdateParamsMessageSlackMessageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxMessageListParams struct {
	// When specified, returns results starting immediately after the item identified
	// by this cursor. Use the cursor value from the previous response's metadata to
	// fetch the next page of results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// When specified, returns results starting immediately before the item identified
	// by this cursor. Use the cursor value from the response's metadata to fetch the
	// previous page of results.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return per page. Must be between 1 and 100. Defaults
	// to 20 if not specified.
	Limit          param.Opt[int64]                     `query:"limit,omitzero" json:"-"`
	ConversationID InboxMessageListParamsConversationID `query:"conversation_id,omitzero" json:"-"`
	InboxID        InboxMessageListParamsInboxID        `query:"inbox_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxMessageListParams]'s query parameters as `url.Values`.
func (r InboxMessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxMessageListParamsConversationID struct {
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxMessageListParamsConversationID]'s query parameters as
// `url.Values`.
func (r InboxMessageListParamsConversationID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxMessageListParamsInboxID struct {
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxMessageListParamsInboxID]'s query parameters as
// `url.Values`.
func (r InboxMessageListParamsInboxID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
