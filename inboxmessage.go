// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/internal/apiquery"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/packages/pagination"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// InboxMessageService contains methods and other services that help with
// interacting with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxMessageService] method instead.
type InboxMessageService struct {
	Options []option.RequestOption
}

// NewInboxMessageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboxMessageService(opts ...option.RequestOption) (r InboxMessageService) {
	r = InboxMessageService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing message.
func (r *InboxMessageService) Get(ctx context.Context, id string, query InboxMessageGetParams, opts ...option.RequestOption) (res *EmailMessage, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("inbox_messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of messages.
func (r *InboxMessageService) List(ctx context.Context, query InboxMessageListParams, opts ...option.RequestOption) (res *pagination.CursorPage[EmailMessage], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *InboxMessageService) ListAutoPaging(ctx context.Context, query InboxMessageListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[EmailMessage] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// The Address object represents a recipient or sender of a message. It contains an
// email address and can be linked to a person and an organization in your
// collections.
type Address struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The email address.
	Email string `json:"email,required"`
	// String representing the object’s type. Always `address` for this object.
	Type constant.Address `json:"type,required"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A hash of related links.
	Links AddressLinks `json:"links"`
	// The role of the address in the message. Can be `from`, `reply_to`, `to`, `cc`,
	// or `bcc`.
	//
	// Any of "from", "reply_to", "to", "cc", "bcc".
	Role AddressRole `json:"role"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		Type        respjson.Field
		CreatedAt   respjson.Field
		Links       respjson.Field
		Role        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Address) RawJSON() string { return r.JSON.raw }
func (r *Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A hash of related links.
type AddressLinks struct {
	// A link to the associated `Organization` item.
	Organization string `json:"organization" format:"uri"`
	// A link to the associated `Person` item.
	Person string `json:"person" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		Person       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressLinks) RawJSON() string { return r.JSON.raw }
func (r *AddressLinks) UnmarshalJSON(data []byte) error {
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
	ID    string            `json:"id,required"`
	Links EmailMessageLinks `json:"links,required"`
	// The globally unique `Message-ID` header of the email.
	Rfc822MessageID string `json:"rfc822_message_id,required"`
	// The subject line of the email.
	Subject string `json:"subject,required"`
	// String representing the object’s type. Always `email_message` for this object.
	Type constant.EmailMessage `json:"type,required"`
	// A list of `Address` objects associated with the message (sender and recipients).
	Addresses []Address `json:"addresses"`
	// A list of `Attachment` objects on the message.
	Attachments []EmailMessageAttachment `json:"attachments"`
	// The HTML content of the email body.
	BodyHTML string `json:"body_html"`
	// The plain text content of the email body.
	BodyPlain string `json:"body_plain"`
	// `true` if the message appears to be part of a bulk mailing.
	Bulk bool `json:"bulk"`
	// The `Conversation` thread this message is part of.
	Conversation InboxConversation `json:"conversation"`
	// The time the message was received, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// `true` if the message is a draft that has not been sent.
	Draft bool `json:"draft"`
	// The `Message-ID` of the email this message is a reply to.
	InReplyToRfc822MessageID string `json:"in_reply_to_rfc822_message_id"`
	// `true` if the message is classified as spam.
	Spam bool `json:"spam"`
	// A concise, system-generated summary of the email content.
	Summary string `json:"summary"`
	// `true` if the message is in the trash.
	Trash bool `json:"trash"`
	// `true` if the message has not been read.
	Unread bool `json:"unread"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Links                    respjson.Field
		Rfc822MessageID          respjson.Field
		Subject                  respjson.Field
		Type                     respjson.Field
		Addresses                respjson.Field
		Attachments              respjson.Field
		BodyHTML                 respjson.Field
		BodyPlain                respjson.Field
		Bulk                     respjson.Field
		Conversation             respjson.Field
		CreatedAt                respjson.Field
		Draft                    respjson.Field
		InReplyToRfc822MessageID respjson.Field
		Spam                     respjson.Field
		Summary                  respjson.Field
		Trash                    respjson.Field
		Unread                   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessage) RawJSON() string { return r.JSON.raw }
func (r *EmailMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageLinks struct {
	// A link to the `Conversation` this message belongs to.
	Conversation string `json:"conversation,required" format:"uri"`
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conversation respjson.Field
		Self         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageLinks) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Attachment object represents a file attached to a message. You can download
// the file content via the `download_url`.
type EmailMessageAttachment struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The original name of the uploaded file, including its extension.
	Filename string `json:"filename,required"`
	// A hash of related links.
	Links EmailMessageAttachmentLinks `json:"links,required"`
	// The size of the file in bytes.
	Size int64 `json:"size,required"`
	// String representing the object’s type. Always `attachment` for this object.
	Type constant.Attachment `json:"type,required"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Filename    respjson.Field
		Links       respjson.Field
		Size        respjson.Field
		Type        respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageAttachment) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A hash of related links.
type EmailMessageAttachmentLinks struct {
	// A link to the `Conversation` this attachment belongs to.
	Conversation string `json:"conversation,required" format:"uri"`
	// A temporary, signed URL to download the file content. The URL expires after one
	// hour.
	DownloadURL string `json:"download_url,required" format:"uri"`
	// A link to the `Message` this attachment belongs to.
	Message string `json:"message,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conversation respjson.Field
		DownloadURL  respjson.Field
		Message      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageAttachmentLinks) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageAttachmentLinks) UnmarshalJSON(data []byte) error {
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
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter messages by one or more conversation IDs.
	Conversation []string `query:"conversation,omitzero" json:"-"`
	// Filter messages by one or more inbox IDs.
	Inbox []string `query:"inbox,omitzero" json:"-"`
	// Specifies which related objects to include in the response. Valid options are
	// `addresses`, `attachments`, and `conversation`.
	//
	// Any of "addresses", "attachments", "conversation".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxMessageListParams]'s query parameters as `url.Values`.
func (r InboxMessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
