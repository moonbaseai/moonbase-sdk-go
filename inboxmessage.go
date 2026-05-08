// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
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
func (r *InboxMessageService) New(ctx context.Context, body InboxMessageNewParams, opts ...option.RequestOption) (res *EmailMessage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "inbox_messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the details of an existing message.
func (r *InboxMessageService) Get(ctx context.Context, id string, query InboxMessageGetParams, opts ...option.RequestOption) (res *EmailMessage, err error) {
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
func (r *InboxMessageService) Update(ctx context.Context, id string, body InboxMessageUpdateParams, opts ...option.RequestOption) (res *EmailMessage, err error) {
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
func (r *InboxMessageService) List(ctx context.Context, query InboxMessageListParams, opts ...option.RequestOption) (res *pagination.CursorPage[EmailMessagePointer], err error) {
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
func (r *InboxMessageService) ListAutoPaging(ctx context.Context, query InboxMessageListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[EmailMessagePointer] {
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

// The Address object represents a recipient or sender of a message. It contains an
// email address and can be linked to a person and an organization in your
// collections.
type Address struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// The email address.
	Email string `json:"email" api:"required" format:"email"`
	// The role of the address in the message. Can be `from`, `reply_to`, `to`, `cc`,
	// or `bcc`.
	//
	// Any of "from", "reply_to", "to", "cc", "bcc".
	Role AddressRole `json:"role" api:"required"`
	// String representing the object’s type. Always `message_address` for this object.
	Type constant.MessageAddress `json:"type" default:"message_address"`
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
func (r Address) RawJSON() string { return r.JSON.raw }
func (r *Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The role of the address in the message. Can be `from`, `reply_to`, `to`, `cc`,
// or `bcc`.
type AddressRole string

const (
	AddressRoleFrom    AddressRole = "from"
	AddressRoleReplyTo AddressRole = "reply_to"
	AddressRoleTo      AddressRole = "to"
	AddressRoleCc      AddressRole = "cc"
	AddressRoleBcc     AddressRole = "bcc"
)

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
	Addresses []Address `json:"addresses"`
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

type EmailMessagePointer struct {
	ID   string                `json:"id" api:"required"`
	Type constant.EmailMessage `json:"type" default:"email_message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessagePointer) RawJSON() string { return r.JSON.raw }
func (r *EmailMessagePointer) UnmarshalJSON(data []byte) error {
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

type InboxMessageNewParams struct {
	// The email body.
	Body shared.FormattedTextParam `json:"body,omitzero" api:"required"`
	// The inbox to use for sending the email.
	InboxID string `json:"inbox_id" api:"required"`
	// The ID of the conversation, if responding to an existing conversation.
	ConversationID param.Opt[string] `json:"conversation_id,omitzero"`
	// The subject line of the email.
	Subject param.Opt[string] `json:"subject,omitzero"`
	// A list of the BCC recipients.
	Bcc []EmailMessageAddressParams `json:"bcc,omitzero"`
	// A list of the CC recipients.
	Cc []EmailMessageAddressParams `json:"cc,omitzero"`
	// A list of recipients.
	To []EmailMessageAddressParams `json:"to,omitzero"`
	paramObj
}

func (r InboxMessageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxMessageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxMessageNewParams) UnmarshalJSON(data []byte) error {
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
	paramObj
}

func (r InboxMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxMessageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxMessageUpdateParams) UnmarshalJSON(data []byte) error {
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
