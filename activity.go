// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbasesdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/moonbase-sdk-go/internal/apijson"
	"github.com/stainless-sdks/moonbase-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/moonbase-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/moonbase-sdk-go/option"
	"github.com/stainless-sdks/moonbase-sdk-go/packages/pagination"
	"github.com/stainless-sdks/moonbase-sdk-go/packages/param"
	"github.com/stainless-sdks/moonbase-sdk-go/packages/respjson"
	"github.com/stainless-sdks/moonbase-sdk-go/shared/constant"
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
// [ActivityActivityCallOccurred], [ActivityActivityFormSubmitted],
// [ActivityActivityInboxMessageSent], [ActivityActivityItemCreated],
// [ActivityActivityItemMentioned], [ActivityActivityMeetingHeld],
// [ActivityActivityMeetingScheduled], [ActivityActivityNoteCreated],
// [ActivityActivityProgramMessageBounced],
// [ActivityActivityProgramMessageClicked],
// [ActivityActivityProgramMessageComplained],
// [ActivityActivityProgramMessageFailed], [ActivityActivityProgramMessageOpened],
// [ActivityActivityProgramMessageSent], [ActivityActivityProgramMessageShielded],
// [ActivityActivityProgramMessageUnsubscribed].
//
// Use the [ActivityUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ActivityUnion struct {
	ID string `json:"id"`
	// This field is a union of [ActivityActivityCallOccurredLinks],
	// [ActivityActivityFormSubmittedLinks], [ActivityActivityInboxMessageSentLinks],
	// [ActivityActivityItemCreatedLinks], [ActivityActivityItemMentionedLinks],
	// [ActivityActivityMeetingHeldLinks], [ActivityActivityMeetingScheduledLinks],
	// [ActivityActivityNoteCreatedLinks],
	// [ActivityActivityProgramMessageBouncedLinks],
	// [ActivityActivityProgramMessageClickedLinks],
	// [ActivityActivityProgramMessageComplainedLinks],
	// [ActivityActivityProgramMessageFailedLinks],
	// [ActivityActivityProgramMessageOpenedLinks],
	// [ActivityActivityProgramMessageSentLinks],
	// [ActivityActivityProgramMessageShieldedLinks],
	// [ActivityActivityProgramMessageUnsubscribedLinks]
	Links      ActivityUnionLinks `json:"links"`
	OccurredAt time.Time          `json:"occurred_at"`
	// Any of "activity/call_occurred", "activity/form_submitted",
	// "activity/inbox_message_sent", "activity/item_created",
	// "activity/item_mentioned", "activity/meeting_held",
	// "activity/meeting_scheduled", "activity/note_created",
	// "activity/program_message_bounced", "activity/program_message_clicked",
	// "activity/program_message_complained", "activity/program_message_failed",
	// "activity/program_message_opened", "activity/program_message_sent",
	// "activity/program_message_shielded", "activity/program_message_unsubscribed".
	Type string `json:"type"`
	// This field is from variant [ActivityActivityCallOccurred].
	Call Call `json:"call"`
	// This field is from variant [ActivityActivityFormSubmitted].
	Collection Collection `json:"collection"`
	// This field is from variant [ActivityActivityFormSubmitted].
	Item Item `json:"item"`
	// This field is from variant [ActivityActivityInboxMessageSent].
	Message EmailMessage `json:"message"`
	// This field is from variant [ActivityActivityInboxMessageSent].
	Recipients []Address `json:"recipients"`
	// This field is from variant [ActivityActivityInboxMessageSent].
	Sender    Address    `json:"sender"`
	Attendees []Attendee `json:"attendees"`
	// This field is from variant [ActivityActivityMeetingHeld].
	Meeting Meeting `json:"meeting"`
	// This field is from variant [ActivityActivityMeetingScheduled].
	Organizer Organizer `json:"organizer"`
	// This field is from variant [ActivityActivityNoteCreated].
	Note Note `json:"note"`
	// This field is from variant [ActivityActivityNoteCreated].
	RelatedItem Item `json:"related_item"`
	// This field is from variant [ActivityActivityNoteCreated].
	RelatedMeeting Meeting `json:"related_meeting"`
	// This field is from variant [ActivityActivityProgramMessageBounced].
	Recipient Address `json:"recipient"`
	// This field is from variant [ActivityActivityProgramMessageClicked].
	LinkText string `json:"link_text"`
	// This field is from variant [ActivityActivityProgramMessageClicked].
	LinkURLUnsafe string `json:"link_url_unsafe"`
	JSON          struct {
		ID             respjson.Field
		Links          respjson.Field
		OccurredAt     respjson.Field
		Type           respjson.Field
		Call           respjson.Field
		Collection     respjson.Field
		Item           respjson.Field
		Message        respjson.Field
		Recipients     respjson.Field
		Sender         respjson.Field
		Attendees      respjson.Field
		Meeting        respjson.Field
		Organizer      respjson.Field
		Note           respjson.Field
		RelatedItem    respjson.Field
		RelatedMeeting respjson.Field
		Recipient      respjson.Field
		LinkText       respjson.Field
		LinkURLUnsafe  respjson.Field
		raw            string
	} `json:"-"`
}

// anyActivity is implemented by each variant of [ActivityUnion] to add type safety
// for the return type of [ActivityUnion.AsAny]
type anyActivity interface {
	implActivityUnion()
}

func (ActivityActivityCallOccurred) implActivityUnion()               {}
func (ActivityActivityFormSubmitted) implActivityUnion()              {}
func (ActivityActivityInboxMessageSent) implActivityUnion()           {}
func (ActivityActivityItemCreated) implActivityUnion()                {}
func (ActivityActivityItemMentioned) implActivityUnion()              {}
func (ActivityActivityMeetingHeld) implActivityUnion()                {}
func (ActivityActivityMeetingScheduled) implActivityUnion()           {}
func (ActivityActivityNoteCreated) implActivityUnion()                {}
func (ActivityActivityProgramMessageBounced) implActivityUnion()      {}
func (ActivityActivityProgramMessageClicked) implActivityUnion()      {}
func (ActivityActivityProgramMessageComplained) implActivityUnion()   {}
func (ActivityActivityProgramMessageFailed) implActivityUnion()       {}
func (ActivityActivityProgramMessageOpened) implActivityUnion()       {}
func (ActivityActivityProgramMessageSent) implActivityUnion()         {}
func (ActivityActivityProgramMessageShielded) implActivityUnion()     {}
func (ActivityActivityProgramMessageUnsubscribed) implActivityUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ActivityUnion.AsAny().(type) {
//	case moonbasesdk.ActivityActivityCallOccurred:
//	case moonbasesdk.ActivityActivityFormSubmitted:
//	case moonbasesdk.ActivityActivityInboxMessageSent:
//	case moonbasesdk.ActivityActivityItemCreated:
//	case moonbasesdk.ActivityActivityItemMentioned:
//	case moonbasesdk.ActivityActivityMeetingHeld:
//	case moonbasesdk.ActivityActivityMeetingScheduled:
//	case moonbasesdk.ActivityActivityNoteCreated:
//	case moonbasesdk.ActivityActivityProgramMessageBounced:
//	case moonbasesdk.ActivityActivityProgramMessageClicked:
//	case moonbasesdk.ActivityActivityProgramMessageComplained:
//	case moonbasesdk.ActivityActivityProgramMessageFailed:
//	case moonbasesdk.ActivityActivityProgramMessageOpened:
//	case moonbasesdk.ActivityActivityProgramMessageSent:
//	case moonbasesdk.ActivityActivityProgramMessageShielded:
//	case moonbasesdk.ActivityActivityProgramMessageUnsubscribed:
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

func (u ActivityUnion) AsActivityCallOccurred() (v ActivityActivityCallOccurred) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityFormSubmitted() (v ActivityActivityFormSubmitted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityInboxMessageSent() (v ActivityActivityInboxMessageSent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityItemCreated() (v ActivityActivityItemCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityItemMentioned() (v ActivityActivityItemMentioned) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityMeetingHeld() (v ActivityActivityMeetingHeld) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityMeetingScheduled() (v ActivityActivityMeetingScheduled) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityNoteCreated() (v ActivityActivityNoteCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageBounced() (v ActivityActivityProgramMessageBounced) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageClicked() (v ActivityActivityProgramMessageClicked) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageComplained() (v ActivityActivityProgramMessageComplained) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageFailed() (v ActivityActivityProgramMessageFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageOpened() (v ActivityActivityProgramMessageOpened) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageSent() (v ActivityActivityProgramMessageSent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageShielded() (v ActivityActivityProgramMessageShielded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActivityUnion) AsActivityProgramMessageUnsubscribed() (v ActivityActivityProgramMessageUnsubscribed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ActivityUnion) RawJSON() string { return u.JSON.raw }

func (r *ActivityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActivityUnionLinks is an implicit subunion of [ActivityUnion].
// ActivityUnionLinks provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [ActivityUnion].
type ActivityUnionLinks struct {
	Self       string `json:"self"`
	Collection string `json:"collection"`
	Item       string `json:"item"`
	// This field is from variant [ActivityActivityInboxMessageSentLinks].
	Message string `json:"message"`
	Meeting string `json:"meeting"`
	// This field is from variant [ActivityActivityNoteCreatedLinks].
	Note string `json:"note"`
	// This field is from variant [ActivityActivityNoteCreatedLinks].
	RelatedItem string `json:"related_item"`
	// This field is from variant [ActivityActivityNoteCreatedLinks].
	RelatedMeeting string `json:"related_meeting"`
	JSON           struct {
		Self           respjson.Field
		Collection     respjson.Field
		Item           respjson.Field
		Message        respjson.Field
		Meeting        respjson.Field
		Note           respjson.Field
		RelatedItem    respjson.Field
		RelatedMeeting respjson.Field
		raw            string
	} `json:"-"`
}

func (r *ActivityUnionLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when an incoming or outgoing call is logged.
type ActivityActivityCallOccurred struct {
	// Unique identifier for the object.
	ID    string                            `json:"id,required"`
	Links ActivityActivityCallOccurredLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/call_occurred`.
	Type constant.ActivityCallOccurred `json:"type,required"`
	// The `Call` object associated with this event.
	Call Call `json:"call"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Call        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityCallOccurred) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityCallOccurred) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityCallOccurredLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityCallOccurredLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityCallOccurredLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `Form` is submitted.
type ActivityActivityFormSubmitted struct {
	// Unique identifier for the object.
	ID    string                             `json:"id,required"`
	Links ActivityActivityFormSubmittedLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/form_submitted`.
	Type constant.ActivityFormSubmitted `json:"type,required"`
	// The `Collection` the new item was added to.
	Collection Collection `json:"collection"`
	// The `Item` that was created by the form submission.
	Item Item `json:"item"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Collection  respjson.Field
		Item        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityFormSubmitted) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityFormSubmitted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityFormSubmittedLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// A link to the `Collection` associated with the form.
	Collection string `json:"collection" format:"uri"`
	// A link to the `Item` created by the form submission.
	Item string `json:"item" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Collection  respjson.Field
		Item        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityFormSubmittedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityFormSubmittedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a message is sent from an `Inbox`.
type ActivityActivityInboxMessageSent struct {
	// Unique identifier for the object.
	ID    string                                `json:"id,required"`
	Links ActivityActivityInboxMessageSentLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/inbox_message_sent`.
	Type constant.ActivityInboxMessageSent `json:"type,required"`
	// The `EmailMessage` that was sent.
	Message EmailMessage `json:"message"`
	// A list of `Address` objects for the recipients.
	Recipients []Address `json:"recipients"`
	// The `Address` of the sender.
	Sender Address `json:"sender"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Message     respjson.Field
		Recipients  respjson.Field
		Sender      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityInboxMessageSent) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityInboxMessageSent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityInboxMessageSentLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// A link to the `EmailMessage` that was sent.
	Message string `json:"message" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityInboxMessageSentLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityInboxMessageSentLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when an `Item` is created.
type ActivityActivityItemCreated struct {
	// Unique identifier for the object.
	ID    string                           `json:"id,required"`
	Links ActivityActivityItemCreatedLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/item_created`.
	Type constant.ActivityItemCreated `json:"type,required"`
	// The `Collection` the item was added to.
	Collection Collection `json:"collection"`
	// The `Item` that was created.
	Item Item `json:"item"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Collection  respjson.Field
		Item        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemCreated) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityItemCreatedLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// A link to the `Collection` the item belongs to.
	Collection string `json:"collection" format:"uri"`
	// A link to the `Item` that was created.
	Item string `json:"item" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Collection  respjson.Field
		Item        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemCreatedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemCreatedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when an `Item` is mentioned.
type ActivityActivityItemMentioned struct {
	// Unique identifier for the object.
	ID    string                             `json:"id,required"`
	Links ActivityActivityItemMentionedLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/item_mentioned`.
	Type constant.ActivityItemMentioned `json:"type,required"`
	// The `Collection` the item belongs to.
	Collection Collection `json:"collection"`
	// The `Item` that was mentioned.
	Item Item `json:"item"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Collection  respjson.Field
		Item        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemMentioned) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemMentioned) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityItemMentionedLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// A link to the `Collection` the item belongs to.
	Collection string `json:"collection" format:"uri"`
	// A link to the `Item` that was mentioned.
	Item string `json:"item" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Collection  respjson.Field
		Item        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemMentionedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemMentionedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `Meeting` has concluded.
type ActivityActivityMeetingHeld struct {
	// Unique identifier for the object.
	ID    string                           `json:"id,required"`
	Links ActivityActivityMeetingHeldLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/meeting_held`.
	Type constant.ActivityMeetingHeld `json:"type,required"`
	// A list of `Attendee` objects who were part of the meeting.
	Attendees []Attendee `json:"attendees"`
	// The `Meeting` object associated with this event.
	Meeting Meeting `json:"meeting"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Attendees   respjson.Field
		Meeting     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityMeetingHeld) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityMeetingHeld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityMeetingHeldLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// A link to the `Meeting` that was held.
	Meeting string `json:"meeting" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Meeting     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityMeetingHeldLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityMeetingHeldLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `Meeting` is scheduled.
type ActivityActivityMeetingScheduled struct {
	// Unique identifier for the object.
	ID    string                                `json:"id,required"`
	Links ActivityActivityMeetingScheduledLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/meeting_scheduled`.
	Type constant.ActivityMeetingScheduled `json:"type,required"`
	// The list of `Attendee` objects invited to the meeting.
	Attendees []Attendee `json:"attendees"`
	// The `Meeting` object associated with this event.
	Meeting Meeting `json:"meeting"`
	// The `Organizer` of the meeting.
	Organizer Organizer `json:"organizer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Attendees   respjson.Field
		Meeting     respjson.Field
		Organizer   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityMeetingScheduled) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityMeetingScheduled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityMeetingScheduledLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// A link to the `Meeting` that was scheduled.
	Meeting string `json:"meeting" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Meeting     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityMeetingScheduledLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityMeetingScheduledLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `Note` is created.
type ActivityActivityNoteCreated struct {
	// Unique identifier for the object.
	ID    string                           `json:"id,required"`
	Links ActivityActivityNoteCreatedLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/note_created`.
	Type constant.ActivityNoteCreated `json:"type,required"`
	// The `Note` object that was created.
	Note Note `json:"note"`
	// The `Item` this note is related to, if any.
	RelatedItem Item `json:"related_item"`
	// The `Meeting` this note is related to, if any.
	RelatedMeeting Meeting `json:"related_meeting"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Links          respjson.Field
		OccurredAt     respjson.Field
		Type           respjson.Field
		Note           respjson.Field
		RelatedItem    respjson.Field
		RelatedMeeting respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityNoteCreated) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityNoteCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityNoteCreatedLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// A link to the `Note` that was created.
	Note string `json:"note" format:"uri"`
	// A link to the related `Item`.
	RelatedItem string `json:"related_item" format:"uri"`
	// A link to the related `Meeting`.
	RelatedMeeting string `json:"related_meeting" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self           respjson.Field
		Note           respjson.Field
		RelatedItem    respjson.Field
		RelatedMeeting respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityNoteCreatedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityNoteCreatedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `ProgramMessage` bounces.
type ActivityActivityProgramMessageBounced struct {
	// Unique identifier for the object.
	ID    string                                     `json:"id,required"`
	Links ActivityActivityProgramMessageBouncedLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/program_message_bounced`.
	Type constant.ActivityProgramMessageBounced `json:"type,required"`
	// The `Address` of the recipient whose message bounced.
	Recipient Address `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageBounced) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageBounced) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageBouncedLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageBouncedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageBouncedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a recipient clicks a tracked link in a
// `ProgramMessage`.
type ActivityActivityProgramMessageClicked struct {
	// Unique identifier for the object.
	ID    string                                     `json:"id,required"`
	Links ActivityActivityProgramMessageClickedLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/program_message_clicked`.
	Type constant.ActivityProgramMessageClicked `json:"type,required"`
	// The text of the link that was clicked.
	LinkText string `json:"link_text"`
	// The URL of the link that was clicked.
	LinkURLUnsafe string `json:"link_url_unsafe"`
	// The `Address` of the recipient who clicked the link.
	Recipient Address `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Links         respjson.Field
		OccurredAt    respjson.Field
		Type          respjson.Field
		LinkText      respjson.Field
		LinkURLUnsafe respjson.Field
		Recipient     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageClicked) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageClicked) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageClickedLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageClickedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageClickedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a recipient marks a `ProgramMessage` as
// spam.
type ActivityActivityProgramMessageComplained struct {
	// Unique identifier for the object.
	ID    string                                        `json:"id,required"`
	Links ActivityActivityProgramMessageComplainedLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/program_message_complained`.
	Type constant.ActivityProgramMessageComplained `json:"type,required"`
	// The `Address` of the recipient who complained.
	Recipient Address `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageComplained) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageComplained) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageComplainedLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageComplainedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageComplainedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `ProgramMessage` fails to be delivered
// for a technical reason.
type ActivityActivityProgramMessageFailed struct {
	// Unique identifier for the object.
	ID    string                                    `json:"id,required"`
	Links ActivityActivityProgramMessageFailedLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/program_message_failed`.
	Type constant.ActivityProgramMessageFailed `json:"type,required"`
	// The `Address` of the recipient whose message failed.
	Recipient Address `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageFailed) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageFailedLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageFailedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageFailedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a recipient opens a `ProgramMessage`.
type ActivityActivityProgramMessageOpened struct {
	// Unique identifier for the object.
	ID    string                                    `json:"id,required"`
	Links ActivityActivityProgramMessageOpenedLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/program_message_opened`.
	Type constant.ActivityProgramMessageOpened `json:"type,required"`
	// The `Address` of the recipient who opened the message.
	Recipient Address `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageOpened) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageOpened) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageOpenedLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageOpenedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageOpenedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `ProgramMessage` is successfully sent.
type ActivityActivityProgramMessageSent struct {
	// Unique identifier for the object.
	ID    string                                  `json:"id,required"`
	Links ActivityActivityProgramMessageSentLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/program_message_sent`.
	Type constant.ActivityProgramMessageSent `json:"type,required"`
	// The `Address` of the recipient the message was sent to.
	Recipient Address `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageSent) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageSent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageSentLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageSentLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageSentLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a `ProgramMessage` is prevented from being
// sent by a delivery protection rule.
type ActivityActivityProgramMessageShielded struct {
	// Unique identifier for the object.
	ID    string                                      `json:"id,required"`
	Links ActivityActivityProgramMessageShieldedLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/program_message_shielded`.
	Type constant.ActivityProgramMessageShielded `json:"type,required"`
	// The `Address` of the recipient whose message was shielded.
	Recipient Address `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageShielded) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageShielded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageShieldedLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageShieldedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageShieldedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when a recipient unsubscribes after receiving a
// `ProgramMessage`.
type ActivityActivityProgramMessageUnsubscribed struct {
	// Unique identifier for the object.
	ID    string                                          `json:"id,required"`
	Links ActivityActivityProgramMessageUnsubscribedLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/program_message_unsubscribed`.
	Type constant.ActivityProgramMessageUnsubscribed `json:"type,required"`
	// The `Address` of the recipient who unsubscribed.
	Recipient Address `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageUnsubscribed) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageUnsubscribed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageUnsubscribedLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageUnsubscribedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageUnsubscribedLinks) UnmarshalJSON(data []byte) error {
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
