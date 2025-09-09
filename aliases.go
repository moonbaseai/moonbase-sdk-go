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

// Structured content that can be rendered in multiple formats, currently
// supporting Markdown.
//
// This is an alias to an internal type.
type FormattedText = shared.FormattedText

// A lightweight reference to another resource.
//
// This is an alias to an internal type.
type Pointer = shared.Pointer

// A lightweight reference to another resource.
//
// This is an alias to an internal type.
type PointerParam = shared.PointerParam
