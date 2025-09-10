// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/moonbaseai/moonbase-sdk-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type ActivityCallOccurred string               // Always "activity/call_occurred"
type ActivityFormSubmitted string              // Always "activity/form_submitted"
type ActivityInboxMessageSent string           // Always "activity/inbox_message_sent"
type ActivityItemCreated string                // Always "activity/item_created"
type ActivityItemMentioned string              // Always "activity/item_mentioned"
type ActivityItemMerged string                 // Always "activity/item_merged"
type ActivityMeetingHeld string                // Always "activity/meeting_held"
type ActivityMeetingScheduled string           // Always "activity/meeting_scheduled"
type ActivityNoteCreated string                // Always "activity/note_created"
type ActivityProgramMessageBounced string      // Always "activity/program_message_bounced"
type ActivityProgramMessageClicked string      // Always "activity/program_message_clicked"
type ActivityProgramMessageComplained string   // Always "activity/program_message_complained"
type ActivityProgramMessageFailed string       // Always "activity/program_message_failed"
type ActivityProgramMessageOpened string       // Always "activity/program_message_opened"
type ActivityProgramMessageSent string         // Always "activity/program_message_sent"
type ActivityProgramMessageShielded string     // Always "activity/program_message_shielded"
type ActivityProgramMessageUnsubscribed string // Always "activity/program_message_unsubscribed"
type Call string                               // Always "call"
type CallParticipant string                    // Always "call_participant"
type ChoiceFieldOption string                  // Always "choice_field_option"
type Collection string                         // Always "collection"
type EmailMessage string                       // Always "email_message"
type Error string                              // Always "error"
type FieldBoolean string                       // Always "field/boolean"
type FieldChoice string                        // Always "field/choice"
type FieldDate string                          // Always "field/date"
type FieldDatetime string                      // Always "field/datetime"
type FieldEmail string                         // Always "field/email"
type FieldGeo string                           // Always "field/geo"
type FieldNumberMonetary string                // Always "field/number/monetary"
type FieldNumberPercentage string              // Always "field/number/percentage"
type FieldNumberUnitlessFloat string           // Always "field/number/unitless_float"
type FieldNumberUnitlessInteger string         // Always "field/number/unitless_integer"
type FieldRelation string                      // Always "field/relation"
type FieldStage string                         // Always "field/stage"
type FieldTelephoneNumber string               // Always "field/telephone_number"
type FieldTextMultiLine string                 // Always "field/text/multi_line"
type FieldTextSingleLine string                // Always "field/text/single_line"
type FieldUriDomain string                     // Always "field/uri/domain"
type FieldUriSocialLinkedIn string             // Always "field/uri/social_linked_in"
type FieldUriSocialX string                    // Always "field/uri/social_x"
type FieldUriURL string                        // Always "field/uri/url"
type File string                               // Always "file"
type Form string                               // Always "form"
type Funnel string                             // Always "funnel"
type FunnelStep string                         // Always "funnel_step"
type Inbox string                              // Always "inbox"
type InboxConversation string                  // Always "inbox_conversation"
type Item string                               // Always "item"
type List string                               // Always "list"
type Meeting string                            // Always "meeting"
type MeetingAttendee string                    // Always "meeting_attendee"
type MeetingOrganizer string                   // Always "meeting_organizer"
type MessageAddress string                     // Always "message_address"
type MessageAttachment string                  // Always "message_attachment"
type Note string                               // Always "note"
type Program string                            // Always "program"
type ProgramMessage string                     // Always "program_message"
type ProgramTemplate string                    // Always "program_template"
type Tag string                                // Always "tag"
type Tagset string                             // Always "tagset"
type ValueBoolean string                       // Always "value/boolean"
type ValueChoice string                        // Always "value/choice"
type ValueDate string                          // Always "value/date"
type ValueDatetime string                      // Always "value/datetime"
type ValueEmail string                         // Always "value/email"
type ValueFunnelStep string                    // Always "value/funnel_step"
type ValueGeo string                           // Always "value/geo"
type ValueNumberMonetary string                // Always "value/number/monetary"
type ValueNumberPercentage string              // Always "value/number/percentage"
type ValueNumberUnitlessFloat string           // Always "value/number/unitless_float"
type ValueNumberUnitlessInteger string         // Always "value/number/unitless_integer"
type ValueRelation string                      // Always "value/relation"
type ValueTelephoneNumber string               // Always "value/telephone_number"
type ValueTextMultiLine string                 // Always "value/text/multi_line"
type ValueTextSingleLine string                // Always "value/text/single_line"
type ValueUriDomain string                     // Always "value/uri/domain"
type ValueUriSocialLinkedIn string             // Always "value/uri/social_linked_in"
type ValueUriSocialX string                    // Always "value/uri/social_x"
type ValueUriURL string                        // Always "value/uri/url"
type View string                               // Always "view"
type WebhookEndpoint string                    // Always "webhook_endpoint"
type WebhookSubscription string                // Always "webhook_subscription"

func (c ActivityCallOccurred) Default() ActivityCallOccurred   { return "activity/call_occurred" }
func (c ActivityFormSubmitted) Default() ActivityFormSubmitted { return "activity/form_submitted" }
func (c ActivityInboxMessageSent) Default() ActivityInboxMessageSent {
	return "activity/inbox_message_sent"
}
func (c ActivityItemCreated) Default() ActivityItemCreated     { return "activity/item_created" }
func (c ActivityItemMentioned) Default() ActivityItemMentioned { return "activity/item_mentioned" }
func (c ActivityItemMerged) Default() ActivityItemMerged       { return "activity/item_merged" }
func (c ActivityMeetingHeld) Default() ActivityMeetingHeld     { return "activity/meeting_held" }
func (c ActivityMeetingScheduled) Default() ActivityMeetingScheduled {
	return "activity/meeting_scheduled"
}
func (c ActivityNoteCreated) Default() ActivityNoteCreated { return "activity/note_created" }
func (c ActivityProgramMessageBounced) Default() ActivityProgramMessageBounced {
	return "activity/program_message_bounced"
}
func (c ActivityProgramMessageClicked) Default() ActivityProgramMessageClicked {
	return "activity/program_message_clicked"
}
func (c ActivityProgramMessageComplained) Default() ActivityProgramMessageComplained {
	return "activity/program_message_complained"
}
func (c ActivityProgramMessageFailed) Default() ActivityProgramMessageFailed {
	return "activity/program_message_failed"
}
func (c ActivityProgramMessageOpened) Default() ActivityProgramMessageOpened {
	return "activity/program_message_opened"
}
func (c ActivityProgramMessageSent) Default() ActivityProgramMessageSent {
	return "activity/program_message_sent"
}
func (c ActivityProgramMessageShielded) Default() ActivityProgramMessageShielded {
	return "activity/program_message_shielded"
}
func (c ActivityProgramMessageUnsubscribed) Default() ActivityProgramMessageUnsubscribed {
	return "activity/program_message_unsubscribed"
}
func (c Call) Default() Call                                   { return "call" }
func (c CallParticipant) Default() CallParticipant             { return "call_participant" }
func (c ChoiceFieldOption) Default() ChoiceFieldOption         { return "choice_field_option" }
func (c Collection) Default() Collection                       { return "collection" }
func (c EmailMessage) Default() EmailMessage                   { return "email_message" }
func (c Error) Default() Error                                 { return "error" }
func (c FieldBoolean) Default() FieldBoolean                   { return "field/boolean" }
func (c FieldChoice) Default() FieldChoice                     { return "field/choice" }
func (c FieldDate) Default() FieldDate                         { return "field/date" }
func (c FieldDatetime) Default() FieldDatetime                 { return "field/datetime" }
func (c FieldEmail) Default() FieldEmail                       { return "field/email" }
func (c FieldGeo) Default() FieldGeo                           { return "field/geo" }
func (c FieldNumberMonetary) Default() FieldNumberMonetary     { return "field/number/monetary" }
func (c FieldNumberPercentage) Default() FieldNumberPercentage { return "field/number/percentage" }
func (c FieldNumberUnitlessFloat) Default() FieldNumberUnitlessFloat {
	return "field/number/unitless_float"
}
func (c FieldNumberUnitlessInteger) Default() FieldNumberUnitlessInteger {
	return "field/number/unitless_integer"
}
func (c FieldRelation) Default() FieldRelation                   { return "field/relation" }
func (c FieldStage) Default() FieldStage                         { return "field/stage" }
func (c FieldTelephoneNumber) Default() FieldTelephoneNumber     { return "field/telephone_number" }
func (c FieldTextMultiLine) Default() FieldTextMultiLine         { return "field/text/multi_line" }
func (c FieldTextSingleLine) Default() FieldTextSingleLine       { return "field/text/single_line" }
func (c FieldUriDomain) Default() FieldUriDomain                 { return "field/uri/domain" }
func (c FieldUriSocialLinkedIn) Default() FieldUriSocialLinkedIn { return "field/uri/social_linked_in" }
func (c FieldUriSocialX) Default() FieldUriSocialX               { return "field/uri/social_x" }
func (c FieldUriURL) Default() FieldUriURL                       { return "field/uri/url" }
func (c File) Default() File                                     { return "file" }
func (c Form) Default() Form                                     { return "form" }
func (c Funnel) Default() Funnel                                 { return "funnel" }
func (c FunnelStep) Default() FunnelStep                         { return "funnel_step" }
func (c Inbox) Default() Inbox                                   { return "inbox" }
func (c InboxConversation) Default() InboxConversation           { return "inbox_conversation" }
func (c Item) Default() Item                                     { return "item" }
func (c List) Default() List                                     { return "list" }
func (c Meeting) Default() Meeting                               { return "meeting" }
func (c MeetingAttendee) Default() MeetingAttendee               { return "meeting_attendee" }
func (c MeetingOrganizer) Default() MeetingOrganizer             { return "meeting_organizer" }
func (c MessageAddress) Default() MessageAddress                 { return "message_address" }
func (c MessageAttachment) Default() MessageAttachment           { return "message_attachment" }
func (c Note) Default() Note                                     { return "note" }
func (c Program) Default() Program                               { return "program" }
func (c ProgramMessage) Default() ProgramMessage                 { return "program_message" }
func (c ProgramTemplate) Default() ProgramTemplate               { return "program_template" }
func (c Tag) Default() Tag                                       { return "tag" }
func (c Tagset) Default() Tagset                                 { return "tagset" }
func (c ValueBoolean) Default() ValueBoolean                     { return "value/boolean" }
func (c ValueChoice) Default() ValueChoice                       { return "value/choice" }
func (c ValueDate) Default() ValueDate                           { return "value/date" }
func (c ValueDatetime) Default() ValueDatetime                   { return "value/datetime" }
func (c ValueEmail) Default() ValueEmail                         { return "value/email" }
func (c ValueFunnelStep) Default() ValueFunnelStep               { return "value/funnel_step" }
func (c ValueGeo) Default() ValueGeo                             { return "value/geo" }
func (c ValueNumberMonetary) Default() ValueNumberMonetary       { return "value/number/monetary" }
func (c ValueNumberPercentage) Default() ValueNumberPercentage   { return "value/number/percentage" }
func (c ValueNumberUnitlessFloat) Default() ValueNumberUnitlessFloat {
	return "value/number/unitless_float"
}
func (c ValueNumberUnitlessInteger) Default() ValueNumberUnitlessInteger {
	return "value/number/unitless_integer"
}
func (c ValueRelation) Default() ValueRelation                   { return "value/relation" }
func (c ValueTelephoneNumber) Default() ValueTelephoneNumber     { return "value/telephone_number" }
func (c ValueTextMultiLine) Default() ValueTextMultiLine         { return "value/text/multi_line" }
func (c ValueTextSingleLine) Default() ValueTextSingleLine       { return "value/text/single_line" }
func (c ValueUriDomain) Default() ValueUriDomain                 { return "value/uri/domain" }
func (c ValueUriSocialLinkedIn) Default() ValueUriSocialLinkedIn { return "value/uri/social_linked_in" }
func (c ValueUriSocialX) Default() ValueUriSocialX               { return "value/uri/social_x" }
func (c ValueUriURL) Default() ValueUriURL                       { return "value/uri/url" }
func (c View) Default() View                                     { return "view" }
func (c WebhookEndpoint) Default() WebhookEndpoint               { return "webhook_endpoint" }
func (c WebhookSubscription) Default() WebhookSubscription       { return "webhook_subscription" }

func (c ActivityCallOccurred) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c ActivityFormSubmitted) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c ActivityInboxMessageSent) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c ActivityItemCreated) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c ActivityItemMentioned) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c ActivityItemMerged) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c ActivityMeetingHeld) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c ActivityMeetingScheduled) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c ActivityNoteCreated) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c ActivityProgramMessageBounced) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ActivityProgramMessageClicked) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ActivityProgramMessageComplained) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c ActivityProgramMessageFailed) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c ActivityProgramMessageOpened) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c ActivityProgramMessageSent) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c ActivityProgramMessageShielded) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ActivityProgramMessageUnsubscribed) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Call) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c CallParticipant) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c ChoiceFieldOption) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Collection) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c EmailMessage) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c Error) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c FieldBoolean) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c FieldChoice) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c FieldDate) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c FieldDatetime) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c FieldEmail) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c FieldGeo) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c FieldNumberMonetary) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c FieldNumberPercentage) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c FieldNumberUnitlessFloat) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c FieldNumberUnitlessInteger) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c FieldRelation) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c FieldStage) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c FieldTelephoneNumber) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c FieldTextMultiLine) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c FieldTextSingleLine) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c FieldUriDomain) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c FieldUriSocialLinkedIn) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c FieldUriSocialX) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c FieldUriURL) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c File) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Form) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Funnel) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c FunnelStep) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c Inbox) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c InboxConversation) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Item) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c List) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Meeting) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c MeetingAttendee) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c MeetingOrganizer) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c MessageAddress) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c MessageAttachment) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Note) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Program) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c ProgramMessage) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c ProgramTemplate) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c Tag) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c Tagset) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c ValueBoolean) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c ValueChoice) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c ValueDate) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c ValueDatetime) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c ValueEmail) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c ValueFunnelStep) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c ValueGeo) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c ValueNumberMonetary) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c ValueNumberPercentage) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c ValueNumberUnitlessFloat) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c ValueNumberUnitlessInteger) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c ValueRelation) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c ValueTelephoneNumber) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c ValueTextMultiLine) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c ValueTextSingleLine) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c ValueUriDomain) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c ValueUriSocialLinkedIn) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c ValueUriSocialX) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c ValueUriURL) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c View) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c WebhookEndpoint) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c WebhookSubscription) MarshalJSON() ([]byte, error)                { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
