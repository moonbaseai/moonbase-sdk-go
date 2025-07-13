// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbasesdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/moonbase-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/moonbase-sdk-go/option"
)

// CollectionFieldService contains methods and other services that help with
// interacting with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCollectionFieldService] method instead.
type CollectionFieldService struct {
	Options []option.RequestOption
}

// NewCollectionFieldService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCollectionFieldService(opts ...option.RequestOption) (r CollectionFieldService) {
	r = CollectionFieldService{}
	r.Options = opts
	return
}

// Retrieves the details of a field in a collection.
func (r *CollectionFieldService) Get(ctx context.Context, id string, query CollectionFieldGetParams, opts ...option.RequestOption) (res *Field, err error) {
	opts = append(r.Options[:], opts...)
	if query.CollectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("collections/%s/fields/%s", query.CollectionID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CollectionFieldGetParams struct {
	CollectionID string `path:"collection_id,required" json:"-"`
	paramObj
}
