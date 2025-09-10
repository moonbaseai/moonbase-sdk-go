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
	"github.com/moonbaseai/moonbase-sdk-go/shared"
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
	Email string `json:"email,required" format:"email"`
	// The role of the address in the message. Can be `from`, `reply_to`, `to`, `cc`,
	// or `bcc`.
	//
	// Any of "from", "reply_to", "to", "cc", "bcc".
	Role AddressRole `json:"role,required"`
	// String representing the object’s type. Always `message_address` for this object.
	Type constant.MessageAddress `json:"type,required"`
	// A lightweight reference to another resource.
	Organization shared.Pointer `json:"organization"`
	// A lightweight reference to another resource.
	Person shared.Pointer `json:"person"`
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
	ID string `json:"id,required"`
	// Structured content that can be rendered in multiple formats, currently
	// supporting Markdown.
	Body shared.FormattedText `json:"body,required"`
	// `true` if the message appears to be part of a bulk mailing.
	Bulk bool `json:"bulk,required"`
	// The time the message was received, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// `true` if the message is a draft that has not been sent.
	Draft bool `json:"draft,required"`
	// `true` if the message is classified as spam.
	Spam bool `json:"spam,required"`
	// The subject line of the email.
	Subject string `json:"subject,required"`
	// `true` if the message is in the trash.
	Trash bool `json:"trash,required"`
	// String representing the object’s type. Always `email_message` for this object.
	Type constant.EmailMessage `json:"type,required"`
	// `true` if the message has not been read.
	Unread bool `json:"unread,required"`
	// A list of `Address` objects associated with the message (sender and recipients).
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Addresses []Address `json:"addresses"`
	// A list of `Attachment` objects on the message.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Attachments []EmailMessageAttachment `json:"attachments"`
	// The `Conversation` thread this message is part of.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Conversation InboxConversation `json:"conversation"`
	// A concise, system-generated summary of the email content.
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Body         respjson.Field
		Bulk         respjson.Field
		CreatedAt    respjson.Field
		Draft        respjson.Field
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

// The Attachment object represents a file attached to a message. You can download
// the file content via the `download_url`.
type EmailMessageAttachment struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A temporary, signed URL to download the file content. The URL expires after one
	// hour.
	DownloadURL string `json:"download_url,required" format:"uri"`
	// The original name of the uploaded file, including its extension.
	Filename string `json:"filename,required"`
	// The size of the file in bytes.
	Size int64 `json:"size,required"`
	// String representing the object’s type. Always `message_attachment` for this
	// object.
	Type constant.MessageAttachment `json:"type,required"`
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
func (r EmailMessageAttachment) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageAttachment) UnmarshalJSON(data []byte) error {
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
	Limit  param.Opt[int64]             `query:"limit,omitzero" json:"-"`
	Filter InboxMessageListParamsFilter `query:"filter,omitzero" json:"-"`
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

type InboxMessageListParamsFilter struct {
	ConversationID InboxMessageListParamsFilterConversationID `query:"conversation_id,omitzero" json:"-"`
	InboxID        InboxMessageListParamsFilterInboxID        `query:"inbox_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxMessageListParamsFilter]'s query parameters as
// `url.Values`.
func (r InboxMessageListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxMessageListParamsFilterConversationID struct {
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxMessageListParamsFilterConversationID]'s query
// parameters as `url.Values`.
func (r InboxMessageListParamsFilterConversationID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxMessageListParamsFilterInboxID struct {
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxMessageListParamsFilterInboxID]'s query parameters as
// `url.Values`.
func (r InboxMessageListParamsFilterInboxID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
