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

// Structured content that can be rendered in multiple formats, currently
// supporting Markdown.
//
// This is an alias to an internal type.
type FormattedTextParam = shared.FormattedTextParam

// A Tag is a label that can be applied to supported resources (such as
// conversations, calls, and meetings) for organization and filtering.
//
// This is an alias to an internal type.
type Tag = shared.Tag

// The color for the tag.
//
// This is an alias to an internal type.
type TagColor = shared.TagColor

// Equals "amber"
const TagColorAmber = shared.TagColorAmber

// Equals "blue"
const TagColorBlue = shared.TagColorBlue

// Equals "cyan"
const TagColorCyan = shared.TagColorCyan

// Equals "emerald"
const TagColorEmerald = shared.TagColorEmerald

// Equals "fuchsia"
const TagColorFuchsia = shared.TagColorFuchsia

// Equals "green"
const TagColorGreen = shared.TagColorGreen

// Equals "indigo"
const TagColorIndigo = shared.TagColorIndigo

// Equals "lime"
const TagColorLime = shared.TagColorLime

// Equals "lunar"
const TagColorLunar = shared.TagColorLunar

// Equals "orange"
const TagColorOrange = shared.TagColorOrange

// Equals "pink"
const TagColorPink = shared.TagColorPink

// Equals "purple"
const TagColorPurple = shared.TagColorPurple

// Equals "red"
const TagColorRed = shared.TagColorRed

// Equals "rose"
const TagColorRose = shared.TagColorRose

// Equals "sky"
const TagColorSky = shared.TagColorSky

// Equals "teal"
const TagColorTeal = shared.TagColorTeal

// Equals "violet"
const TagColorViolet = shared.TagColorViolet

// Equals "yellow"
const TagColorYellow = shared.TagColorYellow

// A lightweight reference to a `Tag` used in request bodies.
//
// This is an alias to an internal type.
type TagPointerParam = shared.TagPointerParam
