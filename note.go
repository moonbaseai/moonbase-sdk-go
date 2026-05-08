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
	"github.com/moonbaseai/moonbase-sdk-go/shared"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// Manage your meetings, files, and notes
//
// NoteService contains methods and other services that help with interacting with
// the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNoteService] method instead.
type NoteService struct {
	Options []option.RequestOption
}

// NewNoteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNoteService(opts ...option.RequestOption) (r NoteService) {
	r = NoteService{}
	r.Options = opts
	return
}

// Create a new note.
func (r *NoteService) New(ctx context.Context, body NoteNewParams, opts ...option.RequestOption) (res *Note, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the details of an existing note.
func (r *NoteService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Note, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("notes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing note.
func (r *NoteService) Update(ctx context.Context, id string, body NoteUpdateParams, opts ...option.RequestOption) (res *Note, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("notes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a list of your notes.
func (r *NoteService) List(ctx context.Context, query NoteListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Note], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "notes"
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

// Returns a list of your notes.
func (r *NoteService) ListAutoPaging(ctx context.Context, query NoteListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Note] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a note.
func (r *NoteService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("notes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// The Note object represents a block of text content, often used for meeting notes
// or summaries.
type Note struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// A list of items, meetings or calls this note is associated with.
	Associations []NoteAssociationPointerUnion `json:"associations" api:"required"`
	// The main content of the note.
	Body shared.FormattedText `json:"body" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The current lock version of the note for optimistic concurrency control.
	LockVersion int64 `json:"lock_version" api:"required"`
	// String representing the object’s type. Always `note` for this object.
	Type constant.Note `json:"type" default:"note"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Creator ItemPointer `json:"creator" api:"nullable"`
	// A short, system-generated summary of the note's content.
	Summary string `json:"summary"`
	// An optional title for the note.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Associations respjson.Field
		Body         respjson.Field
		CreatedAt    respjson.Field
		LockVersion  respjson.Field
		Type         respjson.Field
		UpdatedAt    respjson.Field
		Creator      respjson.Field
		Summary      respjson.Field
		Title        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Note) RawJSON() string { return r.JSON.raw }
func (r *Note) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NoteAssociationParamPointerOfCall(id string) NoteAssociationParamPointerUnion {
	var call CallPointerParam
	call.ID = id
	return NoteAssociationParamPointerUnion{OfCall: &call}
}

func NoteAssociationParamPointerOfItem(id string) NoteAssociationParamPointerUnion {
	var item ItemPointerParam
	item.ID = id
	return NoteAssociationParamPointerUnion{OfItem: &item}
}

func NoteAssociationParamPointerOfMeeting(id string) NoteAssociationParamPointerUnion {
	var meeting MeetingPointerParam
	meeting.ID = id
	return NoteAssociationParamPointerUnion{OfMeeting: &meeting}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NoteAssociationParamPointerUnion struct {
	OfCall    *CallPointerParam    `json:",omitzero,inline"`
	OfItem    *ItemPointerParam    `json:",omitzero,inline"`
	OfMeeting *MeetingPointerParam `json:",omitzero,inline"`
	paramUnion
}

func (u NoteAssociationParamPointerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCall, u.OfItem, u.OfMeeting)
}
func (u *NoteAssociationParamPointerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *NoteAssociationParamPointerUnion) asAny() any {
	if !param.IsOmitted(u.OfCall) {
		return u.OfCall
	} else if !param.IsOmitted(u.OfItem) {
		return u.OfItem
	} else if !param.IsOmitted(u.OfMeeting) {
		return u.OfMeeting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u NoteAssociationParamPointerUnion) GetID() *string {
	if vt := u.OfCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfItem; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfMeeting; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u NoteAssociationParamPointerUnion) GetType() *string {
	if vt := u.OfCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfItem; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMeeting; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[NoteAssociationParamPointerUnion](
		"type",
		apijson.Discriminator[CallPointerParam]("call"),
		apijson.Discriminator[ItemPointerParam]("item"),
		apijson.Discriminator[MeetingPointerParam]("meeting"),
	)
}

// NoteAssociationPointerUnion contains all possible properties and values from
// [CallPointer], [ItemPointer], [MeetingPointer].
//
// Use the [NoteAssociationPointerUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type NoteAssociationPointerUnion struct {
	ID string `json:"id"`
	// Any of "call", "item", "meeting".
	Type string `json:"type"`
	// This field is from variant [ItemPointer].
	Collection CollectionPointer `json:"collection"`
	JSON       struct {
		ID         respjson.Field
		Type       respjson.Field
		Collection respjson.Field
		raw        string
	} `json:"-"`
}

// anyNoteAssociationPointer is implemented by each variant of
// [NoteAssociationPointerUnion] to add type safety for the return type of
// [NoteAssociationPointerUnion.AsAny]
type anyNoteAssociationPointer interface {
	implNoteAssociationPointerUnion()
}

func (CallPointer) implNoteAssociationPointerUnion()    {}
func (ItemPointer) implNoteAssociationPointerUnion()    {}
func (MeetingPointer) implNoteAssociationPointerUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := NoteAssociationPointerUnion.AsAny().(type) {
//	case moonbase.CallPointer:
//	case moonbase.ItemPointer:
//	case moonbase.MeetingPointer:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u NoteAssociationPointerUnion) AsAny() anyNoteAssociationPointer {
	switch u.Type {
	case "call":
		return u.AsCall()
	case "item":
		return u.AsItem()
	case "meeting":
		return u.AsMeeting()
	}
	return nil
}

func (u NoteAssociationPointerUnion) AsCall() (v CallPointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteAssociationPointerUnion) AsItem() (v ItemPointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteAssociationPointerUnion) AsMeeting() (v MeetingPointer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NoteAssociationPointerUnion) RawJSON() string { return u.JSON.raw }

func (r *NoteAssociationPointerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotePointer struct {
	ID   string        `json:"id" api:"required"`
	Type constant.Note `json:"type" default:"note"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotePointer) RawJSON() string { return r.JSON.raw }
func (r *NotePointer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteNewParams struct {
	// The main content of the note.
	Body shared.FormattedTextParam `json:"body,omitzero" api:"required"`
	// Link the Note to Moonbase items (person, organization, deal, task, or an item in
	// a custom collection), meetings, or calls.
	Associations []NoteAssociationParamPointerUnion `json:"associations,omitzero"`
	paramObj
}

func (r NoteNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NoteNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NoteNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteUpdateParams struct {
	// The main content of the note.
	Body shared.FormattedTextParam `json:"body,omitzero" api:"required"`
	// The current lock version of the note for optimistic concurrency control.
	LockVersion int64 `json:"lock_version" api:"required"`
	paramObj
}

func (r NoteUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NoteUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NoteUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteListParams struct {
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

// URLQuery serializes [NoteListParams]'s query parameters as `url.Values`.
func (r NoteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
