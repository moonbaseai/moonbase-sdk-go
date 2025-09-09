// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"encoding/json"
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

// ActivityService contains methods and other services that help with interacting
// with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActivityService] method instead.
type ActivityService struct {
	Options []option.RequestOption
}

// NewActivityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActivityService(opts ...option.RequestOption) (r ActivityService) {
	r = ActivityService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing activity.
func (r *ActivityService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ActivityUnion, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("activities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of activities.
func (r *ActivityService) List(ctx context.Context, query ActivityListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ActivityUnion], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "activities"
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

// Returns a list of activities.
func (r *ActivityService) ListAutoPaging(ctx context.Context, query ActivityListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ActivityUnion] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// ActivityUnion contains all possible properties and values from
// [ActivityCallOccurred], [ActivityFormSubmitted], [ActivityInboxMessageSent],
// [ActivityItemCreated], [ActivityItemMentioned], [ActivityItemMerged],
// [ActivityMeetingHeld], [ActivityMeetingScheduled], [ActivityNoteCreated],
// [ActivityProgramMessageBounced], [ActivityProgramMessageClicked],
// [ActivityProgramMessageComplained], [ActivityProgramMessageFailed],
// [ActivityProgramMessageOpened], [ActivityProgramMessageSent],
// [ActivityProgramMessageShielded], [ActivityProgramMessageUnsubscribed].
//
// Use the [ActivityUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ActivityUnion struct {
	ID string `json:"id"`
	// This field is from variant [ActivityCallOccurred].
	Call       shared.Pointer `json:"call"`
	OccurredAt time.Time      `json:"occurred_at"`
	// Any of "activity/call_occurred", "activity/form_submitted",
	// "activity/inbox_message_sent", "activity/item_created",
	// "activity/item_mentioned", "activity/item_merged", "activity/meeting_held",
	// "activity/meeting_scheduled", "activity/note_created",
	// "activity/program_message_bounced", "activity/program_message_clicked",
	// "activity/program_message_complained", "activity/program_message_failed",
	// "activity/program_message_opened", "activity/program_message_sent",
	// "activity/program_message_shielded", "activity/program_message_unsubscribed".
	Type string `json:"type"`
	// This field is from variant [ActivityFormSubmitted].
	Item ItemPointer `json:"item"`
	// This field is from variant [ActivityInboxMessageSent].
	Message shared.Pointer `json:"message"`
	// This field is from variant [ActivityItemMentioned].
	Author ItemPointer `json:"author"`
	// This field is from variant [ActivityItemMentioned].
	Note shared.Pointer `json:"note"`
	// This field is from variant [ActivityItemMerged].
	Destination ItemPointer `json:"destination"`
	// This field is from variant [ActivityItemMerged].
	Initiator ItemPointer `json:"initiator"`
	// This field is from variant [ActivityItemMerged].
	Source ItemPointer `json:"source"`
	// This field is from variant [ActivityMeetingHeld].
	Meeting shared.Pointer `json:"meeting"`
	// This field is from variant [ActivityNoteCreated].
	RelatedItem ItemPointer `json:"related_item"`
	// This field is from variant [ActivityNoteCreated].
	RelatedMeeting shared.Pointer `json:"related_meeting"`
	// This field is from variant [ActivityProgramMessageBounced].
	ProgramMessage shared.Pointer `json:"program_message"`
	// This field is from variant [ActivityProgramMessageBounced].
	Recipient ItemPointer `json:"recipient"`
	// This field is from variant [ActivityProgramMessageBounced].
	BounceType string `json:"bounce_type"`
	// This field is from variant [ActivityProgramMessageBounced].
	BouncedRecipientEmails []string `json:"bounced_recipient_emails"`
	// This field is from variant [ActivityProgramMessageClicked].
	LinkText string `json:"link_text"`
	// This field is from variant [ActivityProgramMessageClicked].
	LinkURLUnsafe string `json:"link_url_unsafe"`
	ReasonCode    string `json:"reason_code"`
	// This field is from variant [ActivityProgramMessageSent].
	RecipientEmails []string `json:"recipient_emails"`
	// This field is from variant [ActivityProgramMessageUnsubscribed].
	Email string `json:"email"`
	JSON  struct {
		ID                     respjson.Field
		Call                   respjson.Field
		OccurredAt             respjson.Field
		Type                   respjson.Field
		Item                   respjson.Field
		Message                respjson.Field
		Author                 respjson.Field
		Note                   respjson.Field
		Destination            respjson.Field
		Initiator              respjson.Field
		Source                 respjson.Field
		Meeting                respjson.Field
		RelatedItem            respjson.Field
		RelatedMeeting         respjson.Field
		ProgramMessage         respjson.Field
		Recipient              respjson.Field
		BounceType             respjson.Field
		BouncedRecipientEmails respjson.Field
		LinkText               respjson.Field
		LinkURLUnsafe          respjson.Field
		ReasonCode             respjson.Field
		RecipientEmails        respjson.Field
		Email                  respjson.Field
		raw                    string
	} `json:"-"`
}

// anyActivity is implemented by each variant of [ActivityUnion] to add type safety
// for the return type of [ActivityUnion.AsAny]
type anyActivity interface {
	implActivityUnion()
}

func (ActivityCallOccurred) implActivityUnion()               {}
func (ActivityFormSubmitted) implActivityUnion()              {}
func (ActivityInboxMessageSent) implActivityUnion()           {}
func (ActivityItemCreated) implActivityUnion()                {}
func (ActivityItemMentioned) implActivityUnion()              {}
func (ActivityItemMerged) implActivityUnion()                 {}
func (ActivityMeetingHeld) implActivityUnion()                {}
func (ActivityMeetingScheduled) implActivityUnion()           {}
func (ActivityNoteCreated) implActivityUnion()                {}
func (ActivityProgramMessageBounced) implActivityUnion()      {}
func (ActivityProgramMessageClicked) implActivityUnion()      {}
func (ActivityProgramMessageComplained) implActivityUnion()   {}
func (ActivityProgramMessageFailed) implActivityUnion()       {}
func (ActivityProgramMessageOpened) implActivityUnion()       {}
func (ActivityProgramMessageSent) implActivityUnion()         {}
func (ActivityProgramMessageShielded) implActivityUnion()     {}
func (ActivityProgramMessageUnsubscribed) implActivityUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ActivityUnion.AsAny().(type) {
//	case moonbase.ActivityCallOccurred:
//	case moonbase.ActivityFormSubmitted:
//	case moonbase.ActivityInboxMessageSent:
//	case moonbase.ActivityItemCreated:
//	case moonbase.ActivityItemMentioned:
//	case moonbase.ActivityItemMerged:
//	case moonbase.ActivityMeetingHeld:
//	case moonbase.ActivityMeetingScheduled:
//	case moonbase.ActivityNoteCreated:
//	case moonbase.ActivityProgramMessageBounced:
//	case moonbase.ActivityProgramMessageClicked:
//	case moonbase.ActivityProgramMessageComplained:
//	case moonbase.ActivityProgramMessageFailed:
//	case moonbase.ActivityProgramMessageOpened:
//	case moonbase.ActivityProgramMessageSent:
//	case moonbase.ActivityProgramMessageShielded:
//	case moonbase.ActivityProgramMessageUnsubscribed:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ActivityUnion) AsAny() anyActivity {
	switch u.Type {
	case "activity/call_occurred":
		return u.AsActivityCallOccurred()
	case "activity/form_submitted":
		return u.AsActivityFormSubmitted()
	case "activity/inbox_message_sent":
		return u.AsActivityInboxMessageSent()
	case "activity/item_created":
		return u.AsActivityItemCreated()
	case "activity/item_mentioned":
		return u.AsActivityItemMentioned()
	case "activity/item_merged":
		return u.AsActivityItemMerged()
	case "activity/meeting_held":
		return u.AsActivityMeetingHeld()
	case "activity/meeting_scheduled":
		return u.AsActivityMeetingScheduled()
	case "activity/note_created":
		return u.AsActivityNoteCreated()
	case "activity/program_message_bounced":
		return u.AsActivityProgramMessageBounced()
	case "activity/program_message_clicked":
		return u.AsActivityProgramMessageClicked()
	case "activity/program_message_complained":
		return u.AsActivityProgramMessageComplained()
	case "activity/program_message_failed":
		return u.AsActivityProgramMessageFailed()
	case "activity/program_message_opened":
		return u.AsActivityProgramMessageOpened()
	case "activity/program_message_sent":
		return u.AsActivityProgramMessageSent()
	case "activity/program_message_shielded":
		return u.AsActivityProgramMessageShielded()
	case "activity/program_message_unsubscribed":
		return u.AsActivityProgramMessageUnsubscribed()
	}
	return nil
}

func (u ActivityUnion) AsActivityCallOccurred() (v ActivityCallOccurred) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityFormSubmitted() (v ActivityFormSubmitted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityInboxMessageSent() (v ActivityInboxMessageSent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityItemCreated() (v ActivityItemCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityItemMentioned() (v ActivityItemMentioned) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityItemMerged() (v ActivityItemMerged) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityMeetingHeld() (v ActivityMeetingHeld) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityMeetingScheduled() (v ActivityMeetingScheduled) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityNoteCreated() (v ActivityNoteCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageBounced() (v ActivityProgramMessageBounced) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageClicked() (v ActivityProgramMessageClicked) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageComplained() (v ActivityProgramMessageComplained) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageFailed() (v ActivityProgramMessageFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageOpened() (v ActivityProgramMessageOpened) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageSent() (v ActivityProgramMessageSent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageShielded() (v ActivityProgramMessageShielded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageUnsubscribed() (v ActivityProgramMessageUnsubscribed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ActivityUnion) RawJSON() string { return u.JSON.raw }

func (r *ActivityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when an incoming or outgoing call is logged.
type ActivityCallOccurred struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// A lightweight reference to another resource.
	Call shared.Pointer `json:"call,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/call_occurred`.
	Type constant.ActivityCallOccurred `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Call        respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityCallOccurred) RawJSON() string { return r.JSON.raw }
func (r *ActivityCallOccurred) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `Form` is submitted.
type ActivityFormSubmitted struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Item ItemPointer `json:"item,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/form_submitted`.
	Type constant.ActivityFormSubmitted `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Item        respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityFormSubmitted) RawJSON() string { return r.JSON.raw }
func (r *ActivityFormSubmitted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a message is sent from an `Inbox`.
type ActivityInboxMessageSent struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// A lightweight reference to another resource.
	Message shared.Pointer `json:"message,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/inbox_message_sent`.
	Type constant.ActivityInboxMessageSent `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Message     respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityInboxMessageSent) RawJSON() string { return r.JSON.raw }
func (r *ActivityInboxMessageSent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when an `Item` is created.
type ActivityItemCreated struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Item ItemPointer `json:"item,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/item_created`.
	Type constant.ActivityItemCreated `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Item        respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityItemCreated) RawJSON() string { return r.JSON.raw }
func (r *ActivityItemCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when an `Item` is mentioned.
type ActivityItemMentioned struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Author ItemPointer `json:"author,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Item ItemPointer `json:"item,required"`
	// A lightweight reference to another resource.
	Note shared.Pointer `json:"note,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/item_mentioned`.
	Type constant.ActivityItemMentioned `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Author      respjson.Field
		Item        respjson.Field
		Note        respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityItemMentioned) RawJSON() string { return r.JSON.raw }
func (r *ActivityItemMentioned) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when an `Item` is merged into another item.
type ActivityItemMerged struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Destination ItemPointer `json:"destination,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Initiator ItemPointer `json:"initiator,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Source ItemPointer `json:"source,required"`
	// The type of activity. Always `activity/item_merged`.
	Type constant.ActivityItemMerged `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Destination respjson.Field
		Initiator   respjson.Field
		OccurredAt  respjson.Field
		Source      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityItemMerged) RawJSON() string { return r.JSON.raw }
func (r *ActivityItemMerged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `Meeting` has concluded.
type ActivityMeetingHeld struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// A lightweight reference to another resource.
	Meeting shared.Pointer `json:"meeting,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/meeting_held`.
	Type constant.ActivityMeetingHeld `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Meeting     respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityMeetingHeld) RawJSON() string { return r.JSON.raw }
func (r *ActivityMeetingHeld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `Meeting` is scheduled.
type ActivityMeetingScheduled struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// A lightweight reference to another resource.
	Meeting shared.Pointer `json:"meeting,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/meeting_scheduled`.
	Type constant.ActivityMeetingScheduled `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Meeting     respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityMeetingScheduled) RawJSON() string { return r.JSON.raw }
func (r *ActivityMeetingScheduled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `Note` is created.
type ActivityNoteCreated struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// A lightweight reference to another resource.
	Note shared.Pointer `json:"note,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	RelatedItem ItemPointer `json:"related_item,required"`
	// A lightweight reference to another resource.
	RelatedMeeting shared.Pointer `json:"related_meeting,required"`
	// The type of activity. Always `activity/note_created`.
	Type constant.ActivityNoteCreated `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Note           respjson.Field
		OccurredAt     respjson.Field
		RelatedItem    respjson.Field
		RelatedMeeting respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityNoteCreated) RawJSON() string { return r.JSON.raw }
func (r *ActivityNoteCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `ProgramMessage` bounces.
type ActivityProgramMessageBounced struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// A lightweight reference to another resource.
	ProgramMessage shared.Pointer `json:"program_message,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Recipient ItemPointer `json:"recipient,required"`
	// The type of activity. Always `activity/program_message_bounced`.
	Type constant.ActivityProgramMessageBounced `json:"type,required"`
	// The type of bounce (e.g., `Permanent` for hard bounces, `Temporary` for soft
	// bounces).
	BounceType string `json:"bounce_type"`
	// List of email addresses that bounced.
	BouncedRecipientEmails []string `json:"bounced_recipient_emails"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		OccurredAt             respjson.Field
		ProgramMessage         respjson.Field
		Recipient              respjson.Field
		Type                   respjson.Field
		BounceType             respjson.Field
		BouncedRecipientEmails respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityProgramMessageBounced) RawJSON() string { return r.JSON.raw }
func (r *ActivityProgramMessageBounced) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a recipient clicks a tracked link in a
// `ProgramMessage`.
type ActivityProgramMessageClicked struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// A lightweight reference to another resource.
	ProgramMessage shared.Pointer `json:"program_message,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Recipient ItemPointer `json:"recipient,required"`
	// The type of activity. Always `activity/program_message_clicked`.
	Type constant.ActivityProgramMessageClicked `json:"type,required"`
	// The text of the link that was clicked.
	LinkText string `json:"link_text"`
	// The URL of the link that was clicked.
	LinkURLUnsafe string `json:"link_url_unsafe"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		OccurredAt     respjson.Field
		ProgramMessage respjson.Field
		Recipient      respjson.Field
		Type           respjson.Field
		LinkText       respjson.Field
		LinkURLUnsafe  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityProgramMessageClicked) RawJSON() string { return r.JSON.raw }
func (r *ActivityProgramMessageClicked) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a recipient marks a `ProgramMessage` as
// spam.
type ActivityProgramMessageComplained struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// A lightweight reference to another resource.
	ProgramMessage shared.Pointer `json:"program_message,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Recipient ItemPointer `json:"recipient,required"`
	// The type of activity. Always `activity/program_message_complained`.
	Type constant.ActivityProgramMessageComplained `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		OccurredAt     respjson.Field
		ProgramMessage respjson.Field
		Recipient      respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityProgramMessageComplained) RawJSON() string { return r.JSON.raw }
func (r *ActivityProgramMessageComplained) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `ProgramMessage` fails to be delivered
// for a technical reason.
type ActivityProgramMessageFailed struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// A lightweight reference to another resource.
	ProgramMessage shared.Pointer `json:"program_message,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Recipient ItemPointer `json:"recipient,required"`
	// The type of activity. Always `activity/program_message_failed`.
	Type constant.ActivityProgramMessageFailed `json:"type,required"`
	// A code indicating the reason for the failure (e.g., `message_contained_virus`).
	ReasonCode string `json:"reason_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		OccurredAt     respjson.Field
		ProgramMessage respjson.Field
		Recipient      respjson.Field
		Type           respjson.Field
		ReasonCode     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityProgramMessageFailed) RawJSON() string { return r.JSON.raw }
func (r *ActivityProgramMessageFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a recipient opens a `ProgramMessage`.
type ActivityProgramMessageOpened struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// A lightweight reference to another resource.
	ProgramMessage shared.Pointer `json:"program_message,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Recipient ItemPointer `json:"recipient,required"`
	// The type of activity. Always `activity/program_message_opened`.
	Type constant.ActivityProgramMessageOpened `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		OccurredAt     respjson.Field
		ProgramMessage respjson.Field
		Recipient      respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityProgramMessageOpened) RawJSON() string { return r.JSON.raw }
func (r *ActivityProgramMessageOpened) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `ProgramMessage` is successfully sent.
type ActivityProgramMessageSent struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// A lightweight reference to another resource.
	ProgramMessage shared.Pointer `json:"program_message,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Recipient ItemPointer `json:"recipient,required"`
	// The type of activity. Always `activity/program_message_sent`.
	Type constant.ActivityProgramMessageSent `json:"type,required"`
	// List of email addresses the message was sent to.
	RecipientEmails []string `json:"recipient_emails"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		OccurredAt      respjson.Field
		ProgramMessage  respjson.Field
		Recipient       respjson.Field
		Type            respjson.Field
		RecipientEmails respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityProgramMessageSent) RawJSON() string { return r.JSON.raw }
func (r *ActivityProgramMessageSent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `ProgramMessage` is prevented from being
// sent by a delivery protection rule.
type ActivityProgramMessageShielded struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// A lightweight reference to another resource.
	ProgramMessage shared.Pointer `json:"program_message,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Recipient ItemPointer `json:"recipient,required"`
	// The type of activity. Always `activity/program_message_shielded`.
	Type constant.ActivityProgramMessageShielded `json:"type,required"`
	// A code indicating why the message was shielded (e.g.,
	// `person_previously_unsubscribed`).
	ReasonCode string `json:"reason_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		OccurredAt     respjson.Field
		ProgramMessage respjson.Field
		Recipient      respjson.Field
		Type           respjson.Field
		ReasonCode     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityProgramMessageShielded) RawJSON() string { return r.JSON.raw }
func (r *ActivityProgramMessageShielded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a recipient unsubscribes after receiving a
// `ProgramMessage`.
type ActivityProgramMessageUnsubscribed struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// A lightweight reference to another resource.
	ProgramMessage shared.Pointer `json:"program_message,required"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Recipient ItemPointer `json:"recipient,required"`
	// The type of activity. Always `activity/program_message_unsubscribed`.
	Type constant.ActivityProgramMessageUnsubscribed `json:"type,required"`
	// The email address of the person who unsubscribed.
	Email string `json:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		OccurredAt     respjson.Field
		ProgramMessage respjson.Field
		Recipient      respjson.Field
		Type           respjson.Field
		Email          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityProgramMessageUnsubscribed) RawJSON() string { return r.JSON.raw }
func (r *ActivityProgramMessageUnsubscribed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityListParams struct {
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

// URLQuery serializes [ActivityListParams]'s query parameters as `url.Values`.
func (r ActivityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
