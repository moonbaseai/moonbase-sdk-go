// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"encoding/json"
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
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// View activities and capture calls
//
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
func (r *ActivityService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Activity, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("activities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a list of activities.
func (r *ActivityService) List(ctx context.Context, query ActivityListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Activity], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *ActivityService) ListAutoPaging(ctx context.Context, query ActivityListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Activity] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// The Activity object represents a specific event that has occurred, such as a
// meeting being scheduled or a form being submitted.
type Activity struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// An array of entities involved along with each entity's relation to the activity.
	Constituents []Constituent `json:"constituents" api:"required"`
	// The time at which the event occurred, as an ISO 8601 timestamp in UTC.
	OccurredAt time.Time `json:"occurred_at" api:"required" format:"date-time"`
	// The type of activity.
	//
	// Any of "activity/call_occurred", "activity/file_created",
	// "activity/form_submitted", "activity/inbox_message_sent",
	// "activity/item_created", "activity/item_mentioned", "activity/item_merged",
	// "activity/meeting_held", "activity/meeting_scheduled", "activity/note_created",
	// "activity/program_message_bounced", "activity/program_message_clicked",
	// "activity/program_message_complained", "activity/program_message_failed",
	// "activity/program_message_opened", "activity/program_message_sent",
	// "activity/program_message_shielded", "activity/program_message_unsubscribed".
	Type ActivityType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Constituents respjson.Field
		OccurredAt   respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Activity) RawJSON() string { return r.JSON.raw }
func (r *Activity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of activity.
type ActivityType string

const (
	ActivityTypeActivityCallOccurred               ActivityType = "activity/call_occurred"
	ActivityTypeActivityFileCreated                ActivityType = "activity/file_created"
	ActivityTypeActivityFormSubmitted              ActivityType = "activity/form_submitted"
	ActivityTypeActivityInboxMessageSent           ActivityType = "activity/inbox_message_sent"
	ActivityTypeActivityItemCreated                ActivityType = "activity/item_created"
	ActivityTypeActivityItemMentioned              ActivityType = "activity/item_mentioned"
	ActivityTypeActivityItemMerged                 ActivityType = "activity/item_merged"
	ActivityTypeActivityMeetingHeld                ActivityType = "activity/meeting_held"
	ActivityTypeActivityMeetingScheduled           ActivityType = "activity/meeting_scheduled"
	ActivityTypeActivityNoteCreated                ActivityType = "activity/note_created"
	ActivityTypeActivityProgramMessageBounced      ActivityType = "activity/program_message_bounced"
	ActivityTypeActivityProgramMessageClicked      ActivityType = "activity/program_message_clicked"
	ActivityTypeActivityProgramMessageComplained   ActivityType = "activity/program_message_complained"
	ActivityTypeActivityProgramMessageFailed       ActivityType = "activity/program_message_failed"
	ActivityTypeActivityProgramMessageOpened       ActivityType = "activity/program_message_opened"
	ActivityTypeActivityProgramMessageSent         ActivityType = "activity/program_message_sent"
	ActivityTypeActivityProgramMessageShielded     ActivityType = "activity/program_message_shielded"
	ActivityTypeActivityProgramMessageUnsubscribed ActivityType = "activity/program_message_unsubscribed"
)

// The Constituent object represents information about something that was involved
// in a particular activity.
type Constituent struct {
	// A lightweight reference to the entity of `Constituent`, containing information
	// about what type of entity it is as well as the entity's id.
	Entity ConstituentEntityPointerUnion `json:"entity" api:"required"`
	// Any of "actor", "object", "target".
	Relation ConstituentRelation  `json:"relation" api:"required"`
	Type     constant.Constituent `json:"type" default:"constituent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entity      respjson.Field
		Relation    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Constituent) RawJSON() string { return r.JSON.raw }
func (r *Constituent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConstituentRelation string

const (
	ConstituentRelationActor  ConstituentRelation = "actor"
	ConstituentRelationObject ConstituentRelation = "object"
	ConstituentRelationTarget ConstituentRelation = "target"
)

// ConstituentEntityPointerUnion contains all possible properties and values from
// [CallPointer], [CollectionPointer], [ItemPointer], [FilePointer],
// [MeetingPointer], [EmailMessagePointer], [NotePointer], [ProgramPointer],
// [ProgramMessagePointer], [ProgramTemplatePointer], [UnsubscribePointer].
//
// Use the [ConstituentEntityPointerUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConstituentEntityPointerUnion struct {
	ID string `json:"id"`
	// Any of "call", "collection", "item", "file", "meeting", "email_message", "note",
	// "program", "program_message", "program_template", "unsubscribe".
	Type string `json:"type"`
	// This field is from variant [CollectionPointer].
	Ref string `json:"ref"`
	// This field is from variant [ItemPointer].
	Collection CollectionPointer `json:"collection"`
	JSON       struct {
		ID         respjson.Field
		Type       respjson.Field
		Ref        respjson.Field
		Collection respjson.Field
		raw        string
	} `json:"-"`
}

// anyConstituentEntityPointer is implemented by each variant of
// [ConstituentEntityPointerUnion] to add type safety for the return type of
// [ConstituentEntityPointerUnion.AsAny]
type anyConstituentEntityPointer interface {
	implConstituentEntityPointerUnion()
}

func (CallPointer) implConstituentEntityPointerUnion()            {}
func (CollectionPointer) implConstituentEntityPointerUnion()      {}
func (ItemPointer) implConstituentEntityPointerUnion()            {}
func (FilePointer) implConstituentEntityPointerUnion()            {}
func (MeetingPointer) implConstituentEntityPointerUnion()         {}
func (EmailMessagePointer) implConstituentEntityPointerUnion()    {}
func (NotePointer) implConstituentEntityPointerUnion()            {}
func (ProgramPointer) implConstituentEntityPointerUnion()         {}
func (ProgramMessagePointer) implConstituentEntityPointerUnion()  {}
func (ProgramTemplatePointer) implConstituentEntityPointerUnion() {}
func (UnsubscribePointer) implConstituentEntityPointerUnion()     {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConstituentEntityPointerUnion.AsAny().(type) {
//	case moonbase.CallPointer:
//	case moonbase.CollectionPointer:
//	case moonbase.ItemPointer:
//	case moonbase.FilePointer:
//	case moonbase.MeetingPointer:
//	case moonbase.EmailMessagePointer:
//	case moonbase.NotePointer:
//	case moonbase.ProgramPointer:
//	case moonbase.ProgramMessagePointer:
//	case moonbase.ProgramTemplatePointer:
//	case moonbase.UnsubscribePointer:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConstituentEntityPointerUnion) AsAny() anyConstituentEntityPointer {
	switch u.Type {
	case "call":
		return u.AsCall()
	case "collection":
		return u.AsCollection()
	case "item":
		return u.AsItem()
	case "file":
		return u.AsFile()
	case "meeting":
		return u.AsMeeting()
	case "email_message":
		return u.AsEmailMessage()
	case "note":
		return u.AsNote()
	case "program":
		return u.AsProgram()
	case "program_message":
		return u.AsProgramMessage()
	case "program_template":
		return u.AsProgramTemplate()
	case "unsubscribe":
		return u.AsUnsubscribe()
	}
	return nil
}

func (u ConstituentEntityPointerUnion) AsCall() (v CallPointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConstituentEntityPointerUnion) AsCollection() (v CollectionPointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConstituentEntityPointerUnion) AsItem() (v ItemPointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConstituentEntityPointerUnion) AsFile() (v FilePointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConstituentEntityPointerUnion) AsMeeting() (v MeetingPointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConstituentEntityPointerUnion) AsEmailMessage() (v EmailMessagePointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConstituentEntityPointerUnion) AsNote() (v NotePointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConstituentEntityPointerUnion) AsProgram() (v ProgramPointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConstituentEntityPointerUnion) AsProgramMessage() (v ProgramMessagePointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConstituentEntityPointerUnion) AsProgramTemplate() (v ProgramTemplatePointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConstituentEntityPointerUnion) AsUnsubscribe() (v UnsubscribePointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConstituentEntityPointerUnion) RawJSON() string { return u.JSON.raw }

func (r *ConstituentEntityPointerUnion) UnmarshalJSON(data []byte) error {
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
	// Filter activities by which entities were involved. Must be paired with
	// constituent_entity_type.
	ConstituentEntityID ActivityListParamsConstituentEntityID `query:"constituent_entity_id,omitzero" json:"-"`
	// Filter activities by which entities were involved. Must be paired with
	// constituent_entity_id.
	ConstituentEntityType ActivityListParamsConstituentEntityType `query:"constituent_entity_type,omitzero" json:"-"`
	// Filter activities by which entities were involved via specific relations. Must
	// be paired with constituent_entity_type and constituent_entity_id.
	ConstituentRelation ActivityListParamsConstituentRelation `query:"constituent_relation,omitzero" json:"-"`
	// Filter activities by when they occurred.
	OccurredAt ActivityListParamsOccurredAt `query:"occurred_at,omitzero" json:"-"`
	// Filter activities by type.
	Type ActivityListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListParams]'s query parameters as `url.Values`.
func (r ActivityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter activities by which entities were involved. Must be paired with
// constituent_entity_type.
type ActivityListParamsConstituentEntityID struct {
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListParamsConstituentEntityID]'s query parameters
// as `url.Values`.
func (r ActivityListParamsConstituentEntityID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter activities by which entities were involved. Must be paired with
// constituent_entity_id.
type ActivityListParamsConstituentEntityType struct {
	// The type of the entity involved as a constituent of the activity.
	//
	// Any of "call", "collection", "email_message", "file", "item", "meeting", "note",
	// "program", "program_message", "program_template", "unsubscribe".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListParamsConstituentEntityType]'s query parameters
// as `url.Values`.
func (r ActivityListParamsConstituentEntityType) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter activities by which entities were involved via specific relations. Must
// be paired with constituent_entity_type and constituent_entity_id.
type ActivityListParamsConstituentRelation struct {
	// Any of "actor", "object", "target".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListParamsConstituentRelation]'s query parameters
// as `url.Values`.
func (r ActivityListParamsConstituentRelation) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter activities by when they occurred.
type ActivityListParamsOccurredAt struct {
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListParamsOccurredAt]'s query parameters as
// `url.Values`.
func (r ActivityListParamsOccurredAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter activities by type.
type ActivityListParamsType struct {
	// Any of "activity/call_occurred", "activity/form_submitted",
	// "activity/inbox_message_sent", "activity/item_created",
	// "activity/item_mentioned", "activity/item_merged", "activity/file_created",
	// "activity/meeting_held", "activity/meeting_scheduled", "activity/note_created",
	// "activity/program_message_bounced", "activity/program_message_clicked",
	// "activity/program_message_complained", "activity/program_message_failed",
	// "activity/program_message_opened", "activity/program_message_sent",
	// "activity/program_message_shielded", "activity/program_message_unsubscribed".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListParamsType]'s query parameters as `url.Values`.
func (r ActivityListParamsType) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
