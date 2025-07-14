// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"github.com/moonbaseai/moonbase-sdk-go/internal/apierror"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// A View represents a saved configuration for displaying items in a collection,
// including filters and sorting rules.
//
// This is an alias to an internal type.
type View = shared.View

// This is an alias to an internal type.
type ViewLinks = shared.ViewLinks

// The type of view, such as `table` or `board`.
//
// This is an alias to an internal type.
type ViewViewType = shared.ViewViewType

// Equals "table"
const ViewViewTypeTable = shared.ViewViewTypeTable

// Equals "board"
const ViewViewTypeBoard = shared.ViewViewTypeBoard
