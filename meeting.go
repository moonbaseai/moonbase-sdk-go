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

// MeetingService contains methods and other services that help with interacting
// with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingService] method instead.
type MeetingService struct {
	Options []option.RequestOption
}

// NewMeetingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeetingService(opts ...option.RequestOption) (r MeetingService) {
	r = MeetingService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing meeting.
func (r *MeetingService) Get(ctx context.Context, id string, query MeetingGetParams, opts ...option.RequestOption) (res *Meeting, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("meetings/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of meetings.
func (r *MeetingService) List(ctx context.Context, query MeetingListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Meeting], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "meetings"
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

// Returns a list of meetings.
func (r *MeetingService) ListAutoPaging(ctx context.Context, query MeetingListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Meeting] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// The Attendee object represents a participant in a meeting. It includes their
// email address and links to associated `Person` and `Organization` items, if they
// exist in your collections.
type Attendee struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The email address of the attendee.
	Email string `json:"email,required"`
	// A hash of related links.
	Links AttendeeLinks `json:"links,required"`
	// String representing the object’s type. Always `attendee` for this object.
	Type constant.Attendee `json:"type,required"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		Links       respjson.Field
		Type        respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Attendee) RawJSON() string { return r.JSON.raw }
func (r *Attendee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A hash of related links.
type AttendeeLinks struct {
	// A link to the associated `Organization` item.
	Organization string `json:"organization,required" format:"uri"`
	// A link to the associated `Person` item.
	Person string `json:"person,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		Person       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttendeeLinks) RawJSON() string { return r.JSON.raw }
func (r *AttendeeLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Meeting object represents a calendar event. It includes details about the
// participants, timing, and associated content like summaries and recordings.
type Meeting struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The end time of the meeting, as an RFC 3339 timestamp.
	EndAt time.Time `json:"end_at,required" format:"date-time"`
	// The globally unique iCalendar UID for the meeting event.
	ICalUid string       `json:"i_cal_uid,required"`
	Links   MeetingLinks `json:"links,required"`
	// The unique identifier for the meeting from the external calendar provider (e.g.,
	// Google Calendar).
	ProviderID string `json:"provider_id,required"`
	// The start time of the meeting, as an RFC 3339 timestamp.
	StartAt time.Time `json:"start_at,required" format:"date-time"`
	// The IANA time zone in which the meeting is scheduled (e.g.,
	// `America/Los_Angeles`).
	TimeZone string `json:"time_zone,required"`
	// String representing the object’s type. Always `meeting` for this object.
	Type constant.Meeting `json:"type,required"`
	// A list of `Attendee` objects for the meeting.
	Attendees []Attendee `json:"attendees"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A detailed description or agenda for the meeting.
	Description string `json:"description"`
	// The duration of the meeting in seconds.
	Duration float64 `json:"duration"`
	// The physical or virtual location of the meeting.
	Location string `json:"location"`
	// The `Organizer` of the meeting.
	Organizer Organizer `json:"organizer"`
	// A URL to access the meeting in the external provider's system.
	ProviderUri string `json:"provider_uri" format:"uri"`
	// A summary or notes generated before the meeting.
	SummaryAnte string `json:"summary_ante"`
	// A summary or notes generated after the meeting.
	SummaryPost string `json:"summary_post"`
	// The title or subject of the meeting.
	Title string `json:"title"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EndAt       respjson.Field
		ICalUid     respjson.Field
		Links       respjson.Field
		ProviderID  respjson.Field
		StartAt     respjson.Field
		TimeZone    respjson.Field
		Type        respjson.Field
		Attendees   respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Duration    respjson.Field
		Location    respjson.Field
		Organizer   respjson.Field
		ProviderUri respjson.Field
		SummaryAnte respjson.Field
		SummaryPost respjson.Field
		Title       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Meeting) RawJSON() string { return r.JSON.raw }
func (r *Meeting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// A link to an associated `Note` object.
	Note string `json:"note" format:"uri"`
	// A temporary, signed URL to download the meeting recording. The URL expires after
	// one hour.
	RecordingURL string `json:"recording_url" format:"uri"`
	// A link to a long-form summary of the meeting.
	Summary string `json:"summary" format:"uri"`
	// A temporary, signed URL to download the meeting transcript. The URL expires
	// after one hour.
	TranscriptURL string `json:"transcript_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self          respjson.Field
		Note          respjson.Field
		RecordingURL  respjson.Field
		Summary       respjson.Field
		TranscriptURL respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingLinks) RawJSON() string { return r.JSON.raw }
func (r *MeetingLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents the organizer of a meeting.
type Organizer struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The email address of the organizer.
	Email string         `json:"email,required"`
	Links OrganizerLinks `json:"links,required"`
	// String representing the object’s type. Always `organizer` for this object.
	Type constant.Organizer `json:"type,required"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		Links       respjson.Field
		Type        respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Organizer) RawJSON() string { return r.JSON.raw }
func (r *Organizer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizerLinks struct {
	// A link to the associated `Organization` item.
	Organization string `json:"organization,required" format:"uri"`
	// A link to the associated `Person` item.
	Person string `json:"person,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		Person       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizerLinks) RawJSON() string { return r.JSON.raw }
func (r *OrganizerLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingGetParams struct {
	// Specifies which related objects to include in the response. Valid options are
	// `organizer` and `attendees`.
	//
	// Any of "organizer", "attendees".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeetingGetParams]'s query parameters as `url.Values`.
func (r MeetingGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeetingListParams struct {
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

// URLQuery serializes [MeetingListParams]'s query parameters as `url.Values`.
func (r MeetingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
