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
// InboxConversationService contains methods and other services that help with
// interacting with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxConversationService] method instead.
type InboxConversationService struct {
	Options []option.RequestOption
}

// NewInboxConversationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInboxConversationService(opts ...option.RequestOption) (r InboxConversationService) {
	r = InboxConversationService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing conversation.
func (r *InboxConversationService) Get(ctx context.Context, id string, query InboxConversationGetParams, opts ...option.RequestOption) (res *InboxConversation, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("inbox_conversations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a list of your conversations.
func (r *InboxConversationService) List(ctx context.Context, query InboxConversationListParams, opts ...option.RequestOption) (res *pagination.CursorPage[InboxConversationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbox_conversations"
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

// Returns a list of your conversations.
func (r *InboxConversationService) ListAutoPaging(ctx context.Context, query InboxConversationListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[InboxConversationListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// The Conversation object represents a thread of related messages.
type InboxConversation struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// `true` if the conversation appears to be part of a bulk mailing.
	Bulk bool `json:"bulk" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// `true` if a new draft reply to this conversation has been started.
	Draft bool `json:"draft" api:"required"`
	// Whether the conversation is marked for follow-up.
	FollowUp bool `json:"follow_up" api:"required"`
	// The time of the most recent activity in the conversation, as an ISO 8601
	// timestamp in UTC.
	LastMessageAt time.Time `json:"last_message_at" api:"required" format:"date-time"`
	// `true` if the conversation is marked as spam.
	Spam bool `json:"spam" api:"required"`
	// The current state, which can be `unassigned`, `active`, `closed`, or `waiting`.
	//
	// Any of "unassigned", "active", "closed", "waiting".
	State InboxConversationState `json:"state" api:"required"`
	// The subject line of the conversation.
	Subject string `json:"subject" api:"required"`
	// A list of `Tag` objects applied to this conversation.
	Tags []shared.Tag `json:"tags" api:"required"`
	// `true` if the conversation is in the trash.
	Trash bool `json:"trash" api:"required"`
	// String representing the object’s type. Always `inbox_conversation` for this
	// object.
	Type constant.InboxConversation `json:"type" default:"inbox_conversation"`
	// `true` if the conversation contains unread messages.
	Unread bool `json:"unread" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The `Inbox` that this conversations belongs to.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Inbox Inbox `json:"inbox"`
	// The `Message` objects that belong to this conversation.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Messages []any `json:"messages"`
	// If the conversation is snoozed, this is the time it will reappear in the inbox,
	// as an ISO 8601 timestamp in UTC.
	UnsnoozeAt time.Time `json:"unsnooze_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Bulk          respjson.Field
		CreatedAt     respjson.Field
		Draft         respjson.Field
		FollowUp      respjson.Field
		LastMessageAt respjson.Field
		Spam          respjson.Field
		State         respjson.Field
		Subject       respjson.Field
		Tags          respjson.Field
		Trash         respjson.Field
		Type          respjson.Field
		Unread        respjson.Field
		UpdatedAt     respjson.Field
		Inbox         respjson.Field
		Messages      respjson.Field
		UnsnoozeAt    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxConversation) RawJSON() string { return r.JSON.raw }
func (r *InboxConversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current state, which can be `unassigned`, `active`, `closed`, or `waiting`.
type InboxConversationState string

const (
	InboxConversationStateUnassigned InboxConversationState = "unassigned"
	InboxConversationStateActive     InboxConversationState = "active"
	InboxConversationStateClosed     InboxConversationState = "closed"
	InboxConversationStateWaiting    InboxConversationState = "waiting"
)

type InboxConversationListResponse struct {
	ID   string                     `json:"id" api:"required"`
	Type constant.InboxConversation `json:"type" default:"inbox_conversation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxConversationListResponse) RawJSON() string { return r.JSON.raw }
func (r *InboxConversationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxConversationGetParams struct {
	// Specifies which related objects to include in the response. Valid options are
	// `inbox`, `messages`, and `messages.addresses`.
	//
	// Any of "inbox", "messages", "messages.addresses".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxConversationGetParams]'s query parameters as
// `url.Values`.
func (r InboxConversationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxConversationListParams struct {
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
	Limit   param.Opt[int64]                   `query:"limit,omitzero" json:"-"`
	InboxID InboxConversationListParamsInboxID `query:"inbox_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxConversationListParams]'s query parameters as
// `url.Values`.
func (r InboxConversationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxConversationListParamsInboxID struct {
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxConversationListParamsInboxID]'s query parameters as
// `url.Values`.
func (r InboxConversationListParamsInboxID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
