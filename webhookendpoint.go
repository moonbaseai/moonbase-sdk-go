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

// WebhookEndpointService contains methods and other services that help with
// interacting with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookEndpointService] method instead.
type WebhookEndpointService struct {
	Options []option.RequestOption
}

// NewWebhookEndpointService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookEndpointService(opts ...option.RequestOption) (r WebhookEndpointService) {
	r = WebhookEndpointService{}
	r.Options = opts
	return
}

// Create a new endpoint.
func (r *WebhookEndpointService) New(ctx context.Context, body WebhookEndpointNewParams, opts ...option.RequestOption) (res *Endpoint, err error) {
	opts = append(r.Options[:], opts...)
	path := "webhook_endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing endpoint.
func (r *WebhookEndpointService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Endpoint, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("webhook_endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an endpoint.
func (r *WebhookEndpointService) Update(ctx context.Context, id string, body WebhookEndpointUpdateParams, opts ...option.RequestOption) (res *Endpoint, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("webhook_endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of endpoints.
func (r *WebhookEndpointService) List(ctx context.Context, query WebhookEndpointListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Endpoint], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "webhook_endpoints"
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

// Returns a list of endpoints.
func (r *WebhookEndpointService) ListAutoPaging(ctx context.Context, query WebhookEndpointListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Endpoint] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes an endpoint.
func (r *WebhookEndpointService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("webhook_endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// A Webhook Endpoint is an HTTP endpoint that receives webhooks. You can configure
// which events are sent to each endpoint by creating `WebhookSubscription`
// objects.
type Endpoint struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Indicates whether the endpoint is enabled.
	//
	// Any of "disabled", "enabled".
	Status EndpointStatus `json:"status,required"`
	// An array of `WebhookSubscription` objects representing the events this endpoint
	// will receive.
	Subscriptions []Subscription `json:"subscriptions,required"`
	// String representing the object’s type. Always `webhook_endpoint` for this
	// object.
	Type constant.WebhookEndpoint `json:"type,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The HTTPS URL where webhook events will be sent.
	URL string `json:"url,required"`
	// The signing secret used to verify webhook authenticity. This value is only shown
	// when creating the endpoint and starts with `whsec_`.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Status        respjson.Field
		Subscriptions respjson.Field
		Type          respjson.Field
		UpdatedAt     respjson.Field
		URL           respjson.Field
		Secret        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Endpoint) RawJSON() string { return r.JSON.raw }
func (r *Endpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether the endpoint is enabled.
type EndpointStatus string

const (
	EndpointStatusDisabled EndpointStatus = "disabled"
	EndpointStatusEnabled  EndpointStatus = "enabled"
)

// A Webhook Subscription defines which event type should be sent to a specific
// webhook endpoint.
type Subscription struct {
	// The type of event that will trigger notifications to the endpoint (e.g.,
	// `activity/item_created`).
	//
	// Any of "activity/call_occurred", "activity/form_submitted",
	// "activity/inbox_message_sent", "activity/item_created",
	// "activity/item_mentioned", "activity/item_merged", "activity/meeting_held",
	// "activity/meeting_scheduled", "activity/note_created",
	// "activity/program_message_bounced", "activity/program_message_clicked",
	// "activity/program_message_complained", "activity/program_message_failed",
	// "activity/program_message_opened", "activity/program_message_sent",
	// "activity/program_message_shielded", "activity/program_message_unsubscribed".
	EventType SubscriptionEventType `json:"event_type,required"`
	// String representing the object’s type. Always `webhook_subscription` for this
	// object.
	Type constant.WebhookSubscription `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventType   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Subscription) RawJSON() string { return r.JSON.raw }
func (r *Subscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event that will trigger notifications to the endpoint (e.g.,
// `activity/item_created`).
type SubscriptionEventType string

const (
	SubscriptionEventTypeActivityCallOccurred               SubscriptionEventType = "activity/call_occurred"
	SubscriptionEventTypeActivityFormSubmitted              SubscriptionEventType = "activity/form_submitted"
	SubscriptionEventTypeActivityInboxMessageSent           SubscriptionEventType = "activity/inbox_message_sent"
	SubscriptionEventTypeActivityItemCreated                SubscriptionEventType = "activity/item_created"
	SubscriptionEventTypeActivityItemMentioned              SubscriptionEventType = "activity/item_mentioned"
	SubscriptionEventTypeActivityItemMerged                 SubscriptionEventType = "activity/item_merged"
	SubscriptionEventTypeActivityMeetingHeld                SubscriptionEventType = "activity/meeting_held"
	SubscriptionEventTypeActivityMeetingScheduled           SubscriptionEventType = "activity/meeting_scheduled"
	SubscriptionEventTypeActivityNoteCreated                SubscriptionEventType = "activity/note_created"
	SubscriptionEventTypeActivityProgramMessageBounced      SubscriptionEventType = "activity/program_message_bounced"
	SubscriptionEventTypeActivityProgramMessageClicked      SubscriptionEventType = "activity/program_message_clicked"
	SubscriptionEventTypeActivityProgramMessageComplained   SubscriptionEventType = "activity/program_message_complained"
	SubscriptionEventTypeActivityProgramMessageFailed       SubscriptionEventType = "activity/program_message_failed"
	SubscriptionEventTypeActivityProgramMessageOpened       SubscriptionEventType = "activity/program_message_opened"
	SubscriptionEventTypeActivityProgramMessageSent         SubscriptionEventType = "activity/program_message_sent"
	SubscriptionEventTypeActivityProgramMessageShielded     SubscriptionEventType = "activity/program_message_shielded"
	SubscriptionEventTypeActivityProgramMessageUnsubscribed SubscriptionEventType = "activity/program_message_unsubscribed"
)

type WebhookEndpointNewParams struct {
	// Indicates whether the endpoint is enabled.
	//
	// Any of "disabled", "enabled".
	Status WebhookEndpointNewParamsStatus `json:"status,omitzero,required"`
	// The HTTPS URL where webhook events will be sent.
	URL string `json:"url,required"`
	// An array of event types that this endpoint should receive notifications for.
	Subscriptions []WebhookEndpointNewParamsSubscription `json:"subscriptions,omitzero"`
	paramObj
}

func (r WebhookEndpointNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookEndpointNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookEndpointNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether the endpoint is enabled.
type WebhookEndpointNewParamsStatus string

const (
	WebhookEndpointNewParamsStatusDisabled WebhookEndpointNewParamsStatus = "disabled"
	WebhookEndpointNewParamsStatusEnabled  WebhookEndpointNewParamsStatus = "enabled"
)

// Parameters for creating a webhook subscription.
//
// The property EventType is required.
type WebhookEndpointNewParamsSubscription struct {
	// The type of event that will trigger notifications to the endpoint (e.g.,
	// `activity/item_created`).
	//
	// Any of "activity/call_occurred", "activity/form_submitted",
	// "activity/inbox_message_sent", "activity/item_created",
	// "activity/item_mentioned", "activity/item_merged", "activity/meeting_held",
	// "activity/meeting_scheduled", "activity/note_created",
	// "activity/program_message_bounced", "activity/program_message_clicked",
	// "activity/program_message_complained", "activity/program_message_failed",
	// "activity/program_message_opened", "activity/program_message_sent",
	// "activity/program_message_shielded", "activity/program_message_unsubscribed".
	EventType string `json:"event_type,omitzero,required"`
	paramObj
}

func (r WebhookEndpointNewParamsSubscription) MarshalJSON() (data []byte, err error) {
	type shadow WebhookEndpointNewParamsSubscription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookEndpointNewParamsSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookEndpointNewParamsSubscription](
		"event_type", "activity/call_occurred", "activity/form_submitted", "activity/inbox_message_sent", "activity/item_created", "activity/item_mentioned", "activity/item_merged", "activity/meeting_held", "activity/meeting_scheduled", "activity/note_created", "activity/program_message_bounced", "activity/program_message_clicked", "activity/program_message_complained", "activity/program_message_failed", "activity/program_message_opened", "activity/program_message_sent", "activity/program_message_shielded", "activity/program_message_unsubscribed",
	)
}

type WebhookEndpointUpdateParams struct {
	// The HTTPS URL where webhook events will be sent.
	URL param.Opt[string] `json:"url,omitzero"`
	// Indicates whether the endpoint is enabled.
	//
	// Any of "disabled", "enabled".
	Status WebhookEndpointUpdateParamsStatus `json:"status,omitzero"`
	// An array of event types that this endpoint should receive notifications for.
	Subscriptions []WebhookEndpointUpdateParamsSubscription `json:"subscriptions,omitzero"`
	paramObj
}

func (r WebhookEndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookEndpointUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookEndpointUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether the endpoint is enabled.
type WebhookEndpointUpdateParamsStatus string

const (
	WebhookEndpointUpdateParamsStatusDisabled WebhookEndpointUpdateParamsStatus = "disabled"
	WebhookEndpointUpdateParamsStatusEnabled  WebhookEndpointUpdateParamsStatus = "enabled"
)

// Parameters for updating a webhook subscription.
//
// The property EventType is required.
type WebhookEndpointUpdateParamsSubscription struct {
	// The type of event that will trigger notifications to the endpoint (e.g.,
	// `activity/item_created`).
	//
	// Any of "activity/call_occurred", "activity/form_submitted",
	// "activity/inbox_message_sent", "activity/item_created",
	// "activity/item_mentioned", "activity/item_merged", "activity/meeting_held",
	// "activity/meeting_scheduled", "activity/note_created",
	// "activity/program_message_bounced", "activity/program_message_clicked",
	// "activity/program_message_complained", "activity/program_message_failed",
	// "activity/program_message_opened", "activity/program_message_sent",
	// "activity/program_message_shielded", "activity/program_message_unsubscribed".
	EventType string `json:"event_type,omitzero,required"`
	// Unique identifier for the object.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r WebhookEndpointUpdateParamsSubscription) MarshalJSON() (data []byte, err error) {
	type shadow WebhookEndpointUpdateParamsSubscription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookEndpointUpdateParamsSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookEndpointUpdateParamsSubscription](
		"event_type", "activity/call_occurred", "activity/form_submitted", "activity/inbox_message_sent", "activity/item_created", "activity/item_mentioned", "activity/item_merged", "activity/meeting_held", "activity/meeting_scheduled", "activity/note_created", "activity/program_message_bounced", "activity/program_message_clicked", "activity/program_message_complained", "activity/program_message_failed", "activity/program_message_opened", "activity/program_message_sent", "activity/program_message_shielded", "activity/program_message_unsubscribed",
	)
}

type WebhookEndpointListParams struct {
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
	paramObj
}

// URLQuery serializes [WebhookEndpointListParams]'s query parameters as
// `url.Values`.
func (r WebhookEndpointListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
