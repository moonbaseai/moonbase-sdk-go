# Activities

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ActivityUnion">ActivityUnion</a>

Methods:

- <code title="get /activities/{id}">client.Activities.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ActivityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ActivityUnion">ActivityUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /activities">client.Activities.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ActivityService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ActivityListParams">ActivityListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ActivityUnion">ActivityUnion</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Calls

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Call">Call</a>

Methods:

- <code title="post /calls">client.Calls.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#CallService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#CallNewParams">CallNewParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Call">Call</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /calls/upsert">client.Calls.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#CallService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#CallUpsertParams">CallUpsertParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Call">Call</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Collections

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Collection">Collection</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Field">Field</a>

Methods:

- <code title="get /collections/{id}">client.Collections.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#CollectionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#CollectionGetParams">CollectionGetParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Collection">Collection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /collections">client.Collections.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#CollectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#CollectionListParams">CollectionListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Collection">Collection</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Fields

Methods:

- <code title="get /collections/{collection_id}/fields/{id}">client.Collections.Fields.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#CollectionFieldService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#CollectionFieldGetParams">CollectionFieldGetParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Field">Field</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#MoonbaseFile">MoonbaseFile</a>

Methods:

- <code title="get /files/{id}">client.Files.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#MoonbaseFile">MoonbaseFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files">client.Files.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FileListParams">FileListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#MoonbaseFile">MoonbaseFile</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Forms

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Form">Form</a>

Methods:

- <code title="get /forms/{id}">client.Forms.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FormService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FormGetParams">FormGetParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Form">Form</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /forms">client.Forms.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FormService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FormListParams">FormListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Form">Form</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InboxConversations

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxConversation">InboxConversation</a>

Methods:

- <code title="get /inbox_conversations/{id}">client.InboxConversations.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxConversationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxConversationGetParams">InboxConversationGetParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxConversation">InboxConversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /inbox_conversations">client.InboxConversations.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxConversationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxConversationListParams">InboxConversationListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxConversation">InboxConversation</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InboxMessages

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#EmailMessage">EmailMessage</a>

Methods:

- <code title="get /inbox_messages/{id}">client.InboxMessages.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxMessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxMessageGetParams">InboxMessageGetParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#EmailMessage">EmailMessage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /inbox_messages">client.InboxMessages.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxMessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxMessageListParams">InboxMessageListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#EmailMessage">EmailMessage</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Inboxes

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Inbox">Inbox</a>

Methods:

- <code title="get /inboxes/{id}">client.Inboxes.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxGetParams">InboxGetParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Inbox">Inbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /inboxes">client.Inboxes.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#InboxListParams">InboxListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Inbox">Inbox</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Items

Params Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#BooleanValueParam">BooleanValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ChoiceParam">ChoiceParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#DateValueParam">DateValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#DatetimeValueParam">DatetimeValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#DomainValueParam">DomainValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#EmailValueParam">EmailValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FieldValueUnionParam">FieldValueUnionParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FloatValueParam">FloatValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FunnelStepParam">FunnelStepParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#GeoValueParam">GeoValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#IntegerValueParam">IntegerValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ItemParam">ItemParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#MonetaryValueParam">MonetaryValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#MultiLineTextValueParam">MultiLineTextValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#PercentageValueParam">PercentageValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#RelationValueParam">RelationValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#SingleLineTextValueParam">SingleLineTextValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#SocialLinkedInValueParam">SocialLinkedInValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#SocialXValueParam">SocialXValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#TelephoneNumberParam">TelephoneNumberParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#URLValueParam">URLValueParam</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ValueUnionParam">ValueUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#BooleanValue">BooleanValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Choice">Choice</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#DateValue">DateValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#DatetimeValue">DatetimeValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#DomainValue">DomainValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#EmailValue">EmailValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FieldValueUnion">FieldValueUnion</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FloatValue">FloatValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#FunnelStep">FunnelStep</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#GeoValue">GeoValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#IntegerValue">IntegerValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Item">Item</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#MonetaryValue">MonetaryValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#MultiLineTextValue">MultiLineTextValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#PercentageValue">PercentageValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#RelationValue">RelationValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#SingleLineTextValue">SingleLineTextValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#SocialLinkedInValue">SocialLinkedInValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#SocialXValue">SocialXValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#TelephoneNumber">TelephoneNumber</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#URLValue">URLValue</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ValueUnion">ValueUnion</a>

Methods:

- <code title="post /items">client.Items.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ItemService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ItemNewParams">ItemNewParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /items/{id}">client.Items.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ItemService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /items/{id}">client.Items.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ItemService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ItemUpdateParams">ItemUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /items/{id}">client.Items.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ItemService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /items/upsert">client.Items.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ItemService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ItemUpsertParams">ItemUpsertParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Meetings

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Attendee">Attendee</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Meeting">Meeting</a>
- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Organizer">Organizer</a>

Methods:

- <code title="get /meetings/{id}">client.Meetings.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#MeetingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#MeetingGetParams">MeetingGetParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Meeting">Meeting</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /meetings">client.Meetings.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#MeetingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#MeetingListParams">MeetingListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Meeting">Meeting</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Notes

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Note">Note</a>

Methods:

- <code title="get /notes/{id}">client.Notes.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#NoteService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Note">Note</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /notes">client.Notes.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#NoteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#NoteListParams">NoteListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Note">Note</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ProgramMessages

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramMessageNewResponse">ProgramMessageNewResponse</a>

Methods:

- <code title="post /program_messages">client.ProgramMessages.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramMessageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramMessageNewParams">ProgramMessageNewParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramMessageNewResponse">ProgramMessageNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ProgramTemplates

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramTemplate">ProgramTemplate</a>

Methods:

- <code title="get /program_templates/{id}">client.ProgramTemplates.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramTemplateService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramTemplateGetParams">ProgramTemplateGetParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramTemplate">ProgramTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /program_templates">client.ProgramTemplates.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramTemplateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramTemplateListParams">ProgramTemplateListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramTemplate">ProgramTemplate</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Programs

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Program">Program</a>

Methods:

- <code title="get /programs/{id}">client.Programs.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramGetParams">ProgramGetParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Program">Program</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /programs">client.Programs.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ProgramListParams">ProgramListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Program">Program</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tagsets

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Tagset">Tagset</a>

Methods:

- <code title="get /tagsets/{id}">client.Tagsets.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#TagsetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Tagset">Tagset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /tagsets">client.Tagsets.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#TagsetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#TagsetListParams">TagsetListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Tagset">Tagset</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Views

Response Types:

- <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#View">View</a>

Methods:

- <code title="get /views/{id}">client.Views.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ViewService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ViewGetParams">ViewGetParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#View">View</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Items

Methods:

- <code title="get /views/{id}/items">client.Views.Items.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ViewItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#ViewItemListParams">ViewItemListParams</a>) (<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go">moonbase</a>.<a href="https://pkg.go.dev/github.com/moonbaseai/moonbase-sdk-go#Item">Item</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
