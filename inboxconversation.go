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
	ID    string                 `json:"id,required"`
	Links InboxConversationLinks `json:"links,required"`
	// String representing the object’s type. Always `inbox_conversation` for this
	// object.
	Type constant.InboxConversation `json:"type,required"`
	// A list of `Address` objects (participants) in this conversation.
	Addresses []Address `json:"addresses"`
	// `true` if the conversation appears to be part of a bulk mailing.
	Bulk bool `json:"bulk"`
	// `true` if the conversation is currently closed.
	Closed bool `json:"closed"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the conversation is marked for follow-up.
	FollowUp bool `json:"follow_up"`
	// `true` if a new draft reply to this conversation has been started.
	NewDraftConversation bool `json:"new_draft_conversation"`
	// `true` if the conversation is currently open.
	Open bool `json:"open"`
	// `true` if the conversation is marked as spam.
	Spam bool `json:"spam"`
	// The subject line of the conversation.
	Subject string `json:"subject"`
	// A list of `Tag` objects applied to this conversation.
	Tags []InboxConversationTag `json:"tags"`
	// The time of the most recent activity in the conversation, as an RFC 3339
	// timestamp.
	Timestamp string `json:"timestamp"`
	// `true` if the conversation is in the trash.
	Trash bool `json:"trash"`
	// `true` if the conversation contains unread messages.
	Unread bool `json:"unread"`
	// If the conversation is snoozed, this is the time it will reappear in the inbox,
	// as an RFC 3339 timestamp.
	UnsnoozeAt time.Time `json:"unsnooze_at" format:"date-time"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Links                respjson.Field
		Type                 respjson.Field
		Addresses            respjson.Field
		Bulk                 respjson.Field
		Closed               respjson.Field
		CreatedAt            respjson.Field
		FollowUp             respjson.Field
		NewDraftConversation respjson.Field
		Open                 respjson.Field
		Spam                 respjson.Field
		Subject              respjson.Field
		Tags                 respjson.Field
		Timestamp            respjson.Field
		Trash                respjson.Field
		Unread               respjson.Field
		UnsnoozeAt           respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxConversation) RawJSON() string { return r.JSON.raw }
func (r *InboxConversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxConversationLinks struct {
	// A link to the `Inbox` this conversation belongs to.
	Inbox string `json:"inbox,required" format:"uri"`
	// A link to the list of `Message` objects in this conversation.
	Messages string `json:"messages,required" format:"uri"`
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inbox       respjson.Field
		Messages    respjson.Field
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxConversationLinks) RawJSON() string { return r.JSON.raw }
func (r *InboxConversationLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// `addresses` and `tags`.
	//
	// Any of "addresses", "tags".
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
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter conversations by one or more inbox IDs.
	Inbox []string `query:"inbox,omitzero" json:"-"`
	// Specifies which related objects to include in the response. Valid options are
	// `addresses` and `tags`.
	//
	// Any of "addresses", "tags".
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
