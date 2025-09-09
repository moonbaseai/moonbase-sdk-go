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
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("inbox_conversations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of your conversations.
func (r *InboxConversationService) List(ctx context.Context, query InboxConversationListParams, opts ...option.RequestOption) (res *pagination.CursorPage[InboxConversation], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *InboxConversationService) ListAutoPaging(ctx context.Context, query InboxConversationListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[InboxConversation] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// The Conversation object represents a thread of related messages.
type InboxConversation struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// `true` if the conversation appears to be part of a bulk mailing.
	Bulk bool `json:"bulk,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// `true` if a new draft reply to this conversation has been started.
	Draft bool `json:"draft,required"`
	// Whether the conversation is marked for follow-up.
	FollowUp bool `json:"follow_up,required"`
	// The time of the most recent activity in the conversation, as an ISO 8601
	// timestamp in UTC.
	LastMessageAt time.Time `json:"last_message_at,required" format:"date-time"`
	// `true` if the conversation is marked as spam.
	Spam bool `json:"spam,required"`
	// The current state, which can be `unassigned`, `active`, `closed`, or `waiting`.
	//
	// Any of "unassigned", "active", "closed", "waiting".
	State InboxConversationState `json:"state,required"`
	// The subject line of the conversation.
	Subject string `json:"subject,required"`
	// A list of `Tag` objects applied to this conversation.
	Tags []InboxConversationTag `json:"tags,required"`
	// `true` if the conversation is in the trash.
	Trash bool `json:"trash,required"`
	// String representing the object’s type. Always `inbox_conversation` for this
	// object.
	Type constant.InboxConversation `json:"type,required"`
	// `true` if the conversation contains unread messages.
	Unread bool `json:"unread,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The `Inbox` that this conversations belongs to.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Inbox Inbox `json:"inbox"`
	// The `EmailMessage` objects that belong to this conversation.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Messages []EmailMessage `json:"messages"`
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

// A Tag is a label that can be applied to `Conversation` objects for organization
// and filtering.
type InboxConversationTag struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The name of the tag.
	Name string `json:"name,required"`
	// String representing the object’s type. Always `tag` for this object.
	Type constant.Tag `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxConversationTag) RawJSON() string { return r.JSON.raw }
func (r *InboxConversationTag) UnmarshalJSON(data []byte) error {
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
	Limit  param.Opt[int64]                  `query:"limit,omitzero" json:"-"`
	Filter InboxConversationListParamsFilter `query:"filter,omitzero" json:"-"`
	// Specifies which related objects to include in the response. Valid options are
	// `inbox`, `messages`, and `messages.addresses`.
	//
	// Any of "inbox", "messages", "messages.addresses".
	Include []string `query:"include,omitzero" json:"-"`
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

type InboxConversationListParamsFilter struct {
	ConversationID InboxConversationListParamsFilterConversationID `query:"conversation_id,omitzero" json:"-"`
	InboxID        InboxConversationListParamsFilterInboxID        `query:"inbox_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxConversationListParamsFilter]'s query parameters as
// `url.Values`.
func (r InboxConversationListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxConversationListParamsFilterConversationID struct {
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxConversationListParamsFilterConversationID]'s query
// parameters as `url.Values`.
func (r InboxConversationListParamsFilterConversationID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxConversationListParamsFilterInboxID struct {
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxConversationListParamsFilterInboxID]'s query
// parameters as `url.Values`.
func (r InboxConversationListParamsFilterInboxID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
