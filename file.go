// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apiform"
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
// FileService contains methods and other services that help with interacting with
// the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing file.
func (r *FileService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MoonbaseFile, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a list of files that you have uploaded.
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *pagination.CursorPage[MoonbaseFile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "files"
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

// Returns a list of files that you have uploaded.
func (r *FileService) ListAutoPaging(ctx context.Context, query FileListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[MoonbaseFile] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a file.
func (r *FileService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("files/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Upload a file
func (r *FileService) Upload(ctx context.Context, body FileUploadParams, opts ...option.RequestOption) (res *MoonbaseFile, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The File object represents a file that has been uploaded to your library.
type MoonbaseFile struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// A list of items this file is associated with.
	Associations []ItemPointer `json:"associations" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A temporary, signed URL to download the file content. The URL expires after one
	// hour.
	DownloadURL string `json:"download_url" api:"required" format:"uri"`
	// The original filename of the uploaded file.
	Filename string `json:"filename" api:"required"`
	// The display name of the file.
	Name string `json:"name" api:"required"`
	// The size of the file in bytes.
	Size float64 `json:"size" api:"required"`
	// String representing the object’s type. Always `file` for this object.
	Type constant.File `json:"type" default:"file"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Associations respjson.Field
		CreatedAt    respjson.Field
		DownloadURL  respjson.Field
		Filename     respjson.Field
		Name         respjson.Field
		Size         respjson.Field
		Type         respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MoonbaseFile) RawJSON() string { return r.JSON.raw }
func (r *MoonbaseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListParams struct {
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

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileUploadParams struct {
	// The File object to be uploaded.
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	// The display name of the file.
	Name param.Opt[string] `json:"name,omitzero"`
	// Link the File to Moonbase items like a person, organization, deal, task, or an
	// item in a custom collection.
	Associations []shared.PointerParam `json:"associations,omitzero"`
	paramObj
}

func (r FileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
