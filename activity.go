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
// [ActivityActivityCallOccurred], [ActivityActivityFormSubmitted],
// [ActivityActivityInboxMessageSent], [ActivityActivityItemCreated],
// [ActivityActivityItemMentioned], [ActivityActivityItemMerged],
// [ActivityActivityMeetingHeld], [ActivityActivityMeetingScheduled],
// [ActivityActivityNoteCreated], [ActivityActivityProgramMessageBounced],
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
	// [ActivityActivityItemMergedLinks], [ActivityActivityMeetingHeldLinks],
	// [ActivityActivityMeetingScheduledLinks], [ActivityActivityNoteCreatedLinks],
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
	// "activity/item_mentioned", "activity/item_merged", "activity/meeting_held",
	// "activity/meeting_scheduled", "activity/note_created",
	// "activity/program_message_bounced", "activity/program_message_clicked",
	// "activity/program_message_complained", "activity/program_message_failed",
	// "activity/program_message_opened", "activity/program_message_sent",
	// "activity/program_message_shielded", "activity/program_message_unsubscribed".
	Type string `json:"type"`
	// This field is from variant [ActivityActivityCallOccurred].
	Call ActivityActivityCallOccurredCall `json:"call"`
	// This field is a union of [ActivityActivityFormSubmittedCollection],
	// [ActivityActivityItemCreatedCollection]
	Collection ActivityUnionCollection `json:"collection"`
	// This field is a union of [ActivityActivityFormSubmittedRelatedItem],
	// [ActivityActivityNoteCreatedRelatedItem]
	RelatedItem ActivityUnionRelatedItem `json:"related_item"`
	// This field is from variant [ActivityActivityInboxMessageSent].
	Message ActivityActivityInboxMessageSentMessage `json:"message"`
	// This field is from variant [ActivityActivityItemCreated].
	CreatedItem ActivityActivityItemCreatedCreatedItem `json:"created_item"`
	// This field is from variant [ActivityActivityItemMentioned].
	Author ActivityActivityItemMentionedAuthor `json:"author"`
	// This field is from variant [ActivityActivityItemMentioned].
	MentionedItem ActivityActivityItemMentionedMentionedItem `json:"mentioned_item"`
	// This field is a union of [ActivityActivityItemMentionedNote],
	// [ActivityActivityNoteCreatedNote]
	Note ActivityUnionNote `json:"note"`
	// This field is from variant [ActivityActivityItemMerged].
	Destination ActivityActivityItemMergedDestination `json:"destination"`
	// This field is from variant [ActivityActivityItemMerged].
	Initiator ActivityActivityItemMergedInitiator `json:"initiator"`
	// This field is from variant [ActivityActivityItemMerged].
	Source ActivityActivityItemMergedSource `json:"source"`
	// This field is a union of [ActivityActivityMeetingHeldMeeting],
	// [ActivityActivityMeetingScheduledMeeting]
	Meeting ActivityUnionMeeting `json:"meeting"`
	// This field is from variant [ActivityActivityNoteCreated].
	RelatedMeeting ActivityActivityNoteCreatedRelatedMeeting `json:"related_meeting"`
	// This field is from variant [ActivityActivityProgramMessageBounced].
	BounceType string `json:"bounce_type"`
	// This field is from variant [ActivityActivityProgramMessageBounced].
	BouncedRecipientEmails []string `json:"bounced_recipient_emails"`
	// This field is a union of [ActivityActivityProgramMessageBouncedProgramMessage],
	// [ActivityActivityProgramMessageClickedProgramMessage],
	// [ActivityActivityProgramMessageComplainedProgramMessage],
	// [ActivityActivityProgramMessageFailedProgramMessage],
	// [ActivityActivityProgramMessageOpenedProgramMessage],
	// [ActivityActivityProgramMessageSentProgramMessage],
	// [ActivityActivityProgramMessageShieldedProgramMessage],
	// [ActivityActivityProgramMessageUnsubscribedProgramMessage]
	ProgramMessage ActivityUnionProgramMessage `json:"program_message"`
	// This field is a union of [ActivityActivityProgramMessageBouncedRecipient],
	// [ActivityActivityProgramMessageClickedRecipient],
	// [ActivityActivityProgramMessageComplainedRecipient],
	// [ActivityActivityProgramMessageFailedRecipient],
	// [ActivityActivityProgramMessageOpenedRecipient],
	// [ActivityActivityProgramMessageSentRecipient],
	// [ActivityActivityProgramMessageShieldedRecipient],
	// [ActivityActivityProgramMessageUnsubscribedRecipient]
	Recipient ActivityUnionRecipient `json:"recipient"`
	// This field is from variant [ActivityActivityProgramMessageClicked].
	LinkText string `json:"link_text"`
	// This field is from variant [ActivityActivityProgramMessageClicked].
	LinkURLUnsafe string `json:"link_url_unsafe"`
	ReasonCode    string `json:"reason_code"`
	// This field is from variant [ActivityActivityProgramMessageSent].
	RecipientEmails []string `json:"recipient_emails"`
	// This field is from variant [ActivityActivityProgramMessageUnsubscribed].
	Email string `json:"email"`
	JSON  struct {
		ID                     respjson.Field
		Links                  respjson.Field
		OccurredAt             respjson.Field
		Type                   respjson.Field
		Call                   respjson.Field
		Collection             respjson.Field
		RelatedItem            respjson.Field
		Message                respjson.Field
		CreatedItem            respjson.Field
		Author                 respjson.Field
		MentionedItem          respjson.Field
		Note                   respjson.Field
		Destination            respjson.Field
		Initiator              respjson.Field
		Source                 respjson.Field
		Meeting                respjson.Field
		RelatedMeeting         respjson.Field
		BounceType             respjson.Field
		BouncedRecipientEmails respjson.Field
		ProgramMessage         respjson.Field
		Recipient              respjson.Field
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

func (ActivityActivityCallOccurred) implActivityUnion()               {}
func (ActivityActivityFormSubmitted) implActivityUnion()              {}
func (ActivityActivityInboxMessageSent) implActivityUnion()           {}
func (ActivityActivityItemCreated) implActivityUnion()                {}
func (ActivityActivityItemMentioned) implActivityUnion()              {}
func (ActivityActivityItemMerged) implActivityUnion()                 {}
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
//	case moonbase.ActivityActivityCallOccurred:
//	case moonbase.ActivityActivityFormSubmitted:
//	case moonbase.ActivityActivityInboxMessageSent:
//	case moonbase.ActivityActivityItemCreated:
//	case moonbase.ActivityActivityItemMentioned:
//	case moonbase.ActivityActivityItemMerged:
//	case moonbase.ActivityActivityMeetingHeld:
//	case moonbase.ActivityActivityMeetingScheduled:
//	case moonbase.ActivityActivityNoteCreated:
//	case moonbase.ActivityActivityProgramMessageBounced:
//	case moonbase.ActivityActivityProgramMessageClicked:
//	case moonbase.ActivityActivityProgramMessageComplained:
//	case moonbase.ActivityActivityProgramMessageFailed:
//	case moonbase.ActivityActivityProgramMessageOpened:
//	case moonbase.ActivityActivityProgramMessageSent:
//	case moonbase.ActivityActivityProgramMessageShielded:
//	case moonbase.ActivityActivityProgramMessageUnsubscribed:
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

func (u ActivityUnion) AsActivityItemMerged() (v ActivityActivityItemMerged) {
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
	// This field is from variant [ActivityActivityItemMentionedLinks].
	Author string `json:"author"`
	Note   string `json:"note"`
	// This field is from variant [ActivityActivityItemMergedLinks].
	Destination string `json:"destination"`
	// This field is from variant [ActivityActivityItemMergedLinks].
	Initiator string `json:"initiator"`
	Meeting   string `json:"meeting"`
	// This field is from variant [ActivityActivityNoteCreatedLinks].
	RelatedItem string `json:"related_item"`
	// This field is from variant [ActivityActivityNoteCreatedLinks].
	RelatedMeeting string `json:"related_meeting"`
	Recipient      string `json:"recipient"`
	JSON           struct {
		Self           respjson.Field
		Collection     respjson.Field
		Item           respjson.Field
		Message        respjson.Field
		Author         respjson.Field
		Note           respjson.Field
		Destination    respjson.Field
		Initiator      respjson.Field
		Meeting        respjson.Field
		RelatedItem    respjson.Field
		RelatedMeeting respjson.Field
		Recipient      respjson.Field
		raw            string
	} `json:"-"`
}

func (r *ActivityUnionLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActivityUnionCollection is an implicit subunion of [ActivityUnion].
// ActivityUnionCollection provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [ActivityUnion].
type ActivityUnionCollection struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	JSON struct {
		ID   respjson.Field
		Type respjson.Field
		raw  string
	} `json:"-"`
}

func (r *ActivityUnionCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActivityUnionRelatedItem is an implicit subunion of [ActivityUnion].
// ActivityUnionRelatedItem provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [ActivityUnion].
type ActivityUnionRelatedItem struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	JSON struct {
		ID   respjson.Field
		Type respjson.Field
		raw  string
	} `json:"-"`
}

func (r *ActivityUnionRelatedItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActivityUnionNote is an implicit subunion of [ActivityUnion]. ActivityUnionNote
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ActivityUnion].
type ActivityUnionNote struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	JSON struct {
		ID   respjson.Field
		Type respjson.Field
		raw  string
	} `json:"-"`
}

func (r *ActivityUnionNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActivityUnionMeeting is an implicit subunion of [ActivityUnion].
// ActivityUnionMeeting provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [ActivityUnion].
type ActivityUnionMeeting struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	JSON struct {
		ID   respjson.Field
		Type respjson.Field
		raw  string
	} `json:"-"`
}

func (r *ActivityUnionMeeting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActivityUnionProgramMessage is an implicit subunion of [ActivityUnion].
// ActivityUnionProgramMessage provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [ActivityUnion].
type ActivityUnionProgramMessage struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	JSON struct {
		ID   respjson.Field
		Type respjson.Field
		raw  string
	} `json:"-"`
}

func (r *ActivityUnionProgramMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActivityUnionRecipient is an implicit subunion of [ActivityUnion].
// ActivityUnionRecipient provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [ActivityUnion].
type ActivityUnionRecipient struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	JSON struct {
		ID   respjson.Field
		Type respjson.Field
		raw  string
	} `json:"-"`
}

func (r *ActivityUnionRecipient) UnmarshalJSON(data []byte) error {
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
	Call ActivityActivityCallOccurredCall `json:"call"`
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

// The `Call` object associated with this event.
type ActivityActivityCallOccurredCall struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityCallOccurredCall) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityCallOccurredCall) UnmarshalJSON(data []byte) error {
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
	Collection  ActivityActivityFormSubmittedCollection  `json:"collection"`
	RelatedItem ActivityActivityFormSubmittedRelatedItem `json:"related_item"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Collection  respjson.Field
		RelatedItem respjson.Field
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

// The `Collection` the new item was added to.
type ActivityActivityFormSubmittedCollection struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityFormSubmittedCollection) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityFormSubmittedCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityFormSubmittedRelatedItem struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityFormSubmittedRelatedItem) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityFormSubmittedRelatedItem) UnmarshalJSON(data []byte) error {
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
	Message ActivityActivityInboxMessageSentMessage `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Message     respjson.Field
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

// The `EmailMessage` that was sent.
type ActivityActivityInboxMessageSentMessage struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityInboxMessageSentMessage) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityInboxMessageSentMessage) UnmarshalJSON(data []byte) error {
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
	Collection  ActivityActivityItemCreatedCollection  `json:"collection"`
	CreatedItem ActivityActivityItemCreatedCreatedItem `json:"created_item"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Collection  respjson.Field
		CreatedItem respjson.Field
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

// The `Collection` the item was added to.
type ActivityActivityItemCreatedCollection struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemCreatedCollection) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemCreatedCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityItemCreatedCreatedItem struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemCreatedCreatedItem) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemCreatedCreatedItem) UnmarshalJSON(data []byte) error {
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
	Type          constant.ActivityItemMentioned             `json:"type,required"`
	Author        ActivityActivityItemMentionedAuthor        `json:"author"`
	MentionedItem ActivityActivityItemMentionedMentionedItem `json:"mentioned_item"`
	Note          ActivityActivityItemMentionedNote          `json:"note"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Links         respjson.Field
		OccurredAt    respjson.Field
		Type          respjson.Field
		Author        respjson.Field
		MentionedItem respjson.Field
		Note          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
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
	// A link to the `Person` who mentioned the item.
	Author string `json:"author" format:"uri"`
	// A link to the `Item` that was mentioned.
	Item string `json:"item" format:"uri"`
	// A link to the `Note` where the item was mentioned.
	Note string `json:"note" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Author      respjson.Field
		Item        respjson.Field
		Note        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemMentionedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemMentionedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityItemMentionedAuthor struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemMentionedAuthor) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemMentionedAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityItemMentionedMentionedItem struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemMentionedMentionedItem) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemMentionedMentionedItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityItemMentionedNote struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemMentionedNote) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemMentionedNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an event that occurs when an `Item` is merged into another item.
type ActivityActivityItemMerged struct {
	// Unique identifier for the object.
	ID    string                          `json:"id,required"`
	Links ActivityActivityItemMergedLinks `json:"links,required"`
	// The time at which the event occurred, as an RFC 3339 timestamp.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The type of activity. Always `activity/item_merged`.
	Type constant.ActivityItemMerged `json:"type,required"`
	// A pointer to the `Item` that the data was merged into.
	Destination ActivityActivityItemMergedDestination `json:"destination"`
	// The person that performed the merge.
	Initiator ActivityActivityItemMergedInitiator `json:"initiator"`
	// A pointer to the source `Item`.
	Source ActivityActivityItemMergedSource `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Destination respjson.Field
		Initiator   respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemMerged) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemMerged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityItemMergedLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// A link to the `Item` that received the data from the source.
	Destination string `json:"destination" format:"uri"`
	// A link to the person that performed the merge.
	Initiator string `json:"initiator" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Destination respjson.Field
		Initiator   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemMergedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemMergedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A pointer to the `Item` that the data was merged into.
type ActivityActivityItemMergedDestination struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemMergedDestination) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemMergedDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The person that performed the merge.
type ActivityActivityItemMergedInitiator struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemMergedInitiator) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemMergedInitiator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A pointer to the source `Item`.
type ActivityActivityItemMergedSource struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityItemMergedSource) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityItemMergedSource) UnmarshalJSON(data []byte) error {
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
	// The `Meeting` object associated with this event.
	Meeting ActivityActivityMeetingHeldMeeting `json:"meeting"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
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

// The `Meeting` object associated with this event.
type ActivityActivityMeetingHeldMeeting struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityMeetingHeldMeeting) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityMeetingHeldMeeting) UnmarshalJSON(data []byte) error {
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
	// The `Meeting` object associated with this event.
	Meeting ActivityActivityMeetingScheduledMeeting `json:"meeting"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Meeting     respjson.Field
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

// The `Meeting` object associated with this event.
type ActivityActivityMeetingScheduledMeeting struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityMeetingScheduledMeeting) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityMeetingScheduledMeeting) UnmarshalJSON(data []byte) error {
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
	Note ActivityActivityNoteCreatedNote `json:"note"`
	// The `Item` this note is related to, if any.
	RelatedItem ActivityActivityNoteCreatedRelatedItem `json:"related_item"`
	// The `Meeting` this note is related to, if any.
	RelatedMeeting ActivityActivityNoteCreatedRelatedMeeting `json:"related_meeting"`
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

// The `Note` object that was created.
type ActivityActivityNoteCreatedNote struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityNoteCreatedNote) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityNoteCreatedNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The `Item` this note is related to, if any.
type ActivityActivityNoteCreatedRelatedItem struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityNoteCreatedRelatedItem) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityNoteCreatedRelatedItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The `Meeting` this note is related to, if any.
type ActivityActivityNoteCreatedRelatedMeeting struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityNoteCreatedRelatedMeeting) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityNoteCreatedRelatedMeeting) UnmarshalJSON(data []byte) error {
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
	Type                   constant.ActivityProgramMessageBounced              `json:"type,required"`
	BounceType             string                                              `json:"bounce_type"`
	BouncedRecipientEmails []string                                            `json:"bounced_recipient_emails"`
	ProgramMessage         ActivityActivityProgramMessageBouncedProgramMessage `json:"program_message"`
	// A link to the `Address` of the recipient whose message bounced.
	Recipient ActivityActivityProgramMessageBouncedRecipient `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Links                  respjson.Field
		OccurredAt             respjson.Field
		Type                   respjson.Field
		BounceType             respjson.Field
		BouncedRecipientEmails respjson.Field
		ProgramMessage         respjson.Field
		Recipient              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
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
	// A link to the `Address` of the recipient whose message bounced.
	Recipient string `json:"recipient" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageBouncedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageBouncedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageBouncedProgramMessage struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageBouncedProgramMessage) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageBouncedProgramMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A link to the `Address` of the recipient whose message bounced.
type ActivityActivityProgramMessageBouncedRecipient struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageBouncedRecipient) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageBouncedRecipient) UnmarshalJSON(data []byte) error {
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
	LinkURLUnsafe  string                                              `json:"link_url_unsafe"`
	ProgramMessage ActivityActivityProgramMessageClickedProgramMessage `json:"program_message"`
	// A link to the `Address` of the recipient who clicked the link.
	Recipient ActivityActivityProgramMessageClickedRecipient `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Links          respjson.Field
		OccurredAt     respjson.Field
		Type           respjson.Field
		LinkText       respjson.Field
		LinkURLUnsafe  respjson.Field
		ProgramMessage respjson.Field
		Recipient      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
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
	// A link to the `Address` of the recipient who clicked the link.
	Recipient string `json:"recipient" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageClickedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageClickedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageClickedProgramMessage struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageClickedProgramMessage) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageClickedProgramMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A link to the `Address` of the recipient who clicked the link.
type ActivityActivityProgramMessageClickedRecipient struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageClickedRecipient) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageClickedRecipient) UnmarshalJSON(data []byte) error {
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
	Type           constant.ActivityProgramMessageComplained              `json:"type,required"`
	ProgramMessage ActivityActivityProgramMessageComplainedProgramMessage `json:"program_message"`
	// A link to the `Address` of the recipient who complained.
	Recipient ActivityActivityProgramMessageComplainedRecipient `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Links          respjson.Field
		OccurredAt     respjson.Field
		Type           respjson.Field
		ProgramMessage respjson.Field
		Recipient      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
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
	// A link to the `Address` of the recipient who complained.
	Recipient string `json:"recipient" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageComplainedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageComplainedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageComplainedProgramMessage struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageComplainedProgramMessage) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageComplainedProgramMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A link to the `Address` of the recipient who complained.
type ActivityActivityProgramMessageComplainedRecipient struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageComplainedRecipient) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageComplainedRecipient) UnmarshalJSON(data []byte) error {
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
	Type           constant.ActivityProgramMessageFailed              `json:"type,required"`
	ProgramMessage ActivityActivityProgramMessageFailedProgramMessage `json:"program_message"`
	ReasonCode     string                                             `json:"reason_code"`
	// A link to the `Address` of the recipient whose message failed.
	Recipient ActivityActivityProgramMessageFailedRecipient `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Links          respjson.Field
		OccurredAt     respjson.Field
		Type           respjson.Field
		ProgramMessage respjson.Field
		ReasonCode     respjson.Field
		Recipient      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
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
	// A link to the `Address` of the recipient whose message failed.
	Recipient string `json:"recipient" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageFailedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageFailedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageFailedProgramMessage struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageFailedProgramMessage) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageFailedProgramMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A link to the `Address` of the recipient whose message failed.
type ActivityActivityProgramMessageFailedRecipient struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageFailedRecipient) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageFailedRecipient) UnmarshalJSON(data []byte) error {
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
	Type           constant.ActivityProgramMessageOpened              `json:"type,required"`
	ProgramMessage ActivityActivityProgramMessageOpenedProgramMessage `json:"program_message"`
	// A link to the `Address` of the recipient who opened the message.
	Recipient ActivityActivityProgramMessageOpenedRecipient `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Links          respjson.Field
		OccurredAt     respjson.Field
		Type           respjson.Field
		ProgramMessage respjson.Field
		Recipient      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
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
	// A link to the `Address` of the recipient who opened the message.
	Recipient string `json:"recipient" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageOpenedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageOpenedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageOpenedProgramMessage struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageOpenedProgramMessage) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageOpenedProgramMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A link to the `Address` of the recipient who opened the message.
type ActivityActivityProgramMessageOpenedRecipient struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageOpenedRecipient) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageOpenedRecipient) UnmarshalJSON(data []byte) error {
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
	Type           constant.ActivityProgramMessageSent              `json:"type,required"`
	ProgramMessage ActivityActivityProgramMessageSentProgramMessage `json:"program_message"`
	// A link to the `Address` of the recipient the message was sent to.
	Recipient       ActivityActivityProgramMessageSentRecipient `json:"recipient"`
	RecipientEmails []string                                    `json:"recipient_emails"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Links           respjson.Field
		OccurredAt      respjson.Field
		Type            respjson.Field
		ProgramMessage  respjson.Field
		Recipient       respjson.Field
		RecipientEmails respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
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
	// A link to the `Address` of the recipient the message was sent to.
	Recipient string `json:"recipient" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageSentLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageSentLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageSentProgramMessage struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageSentProgramMessage) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageSentProgramMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A link to the `Address` of the recipient the message was sent to.
type ActivityActivityProgramMessageSentRecipient struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageSentRecipient) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageSentRecipient) UnmarshalJSON(data []byte) error {
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
	Type           constant.ActivityProgramMessageShielded              `json:"type,required"`
	ProgramMessage ActivityActivityProgramMessageShieldedProgramMessage `json:"program_message"`
	ReasonCode     string                                               `json:"reason_code"`
	// A link to the `Address` of the recipient whose message was shielded.
	Recipient ActivityActivityProgramMessageShieldedRecipient `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Links          respjson.Field
		OccurredAt     respjson.Field
		Type           respjson.Field
		ProgramMessage respjson.Field
		ReasonCode     respjson.Field
		Recipient      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
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
	// A link to the `Address` of the recipient whose message was shielded.
	Recipient string `json:"recipient" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageShieldedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageShieldedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageShieldedProgramMessage struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageShieldedProgramMessage) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageShieldedProgramMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A link to the `Address` of the recipient whose message was shielded.
type ActivityActivityProgramMessageShieldedRecipient struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageShieldedRecipient) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageShieldedRecipient) UnmarshalJSON(data []byte) error {
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
	Type           constant.ActivityProgramMessageUnsubscribed              `json:"type,required"`
	Email          string                                                   `json:"email"`
	ProgramMessage ActivityActivityProgramMessageUnsubscribedProgramMessage `json:"program_message"`
	// A link to the `Address` of the recipient who unsubscribed.
	Recipient ActivityActivityProgramMessageUnsubscribedRecipient `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Links          respjson.Field
		OccurredAt     respjson.Field
		Type           respjson.Field
		Email          respjson.Field
		ProgramMessage respjson.Field
		Recipient      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
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
	// A link to the `Address` of the recipient who unsubscribed.
	Recipient string `json:"recipient" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageUnsubscribedLinks) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageUnsubscribedLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityActivityProgramMessageUnsubscribedProgramMessage struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageUnsubscribedProgramMessage) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageUnsubscribedProgramMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A link to the `Address` of the recipient who unsubscribed.
type ActivityActivityProgramMessageUnsubscribedRecipient struct {
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActivityActivityProgramMessageUnsubscribedRecipient) RawJSON() string { return r.JSON.raw }
func (r *ActivityActivityProgramMessageUnsubscribedRecipient) UnmarshalJSON(data []byte) error {
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
