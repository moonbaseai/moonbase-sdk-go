// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apiform"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
)

// Manage your inboxes, conversations, and messages
//
// InboxMessageAttachmentService contains methods and other services that help with
// interacting with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxMessageAttachmentService] method instead.
type InboxMessageAttachmentService struct {
	Options []option.RequestOption
}

// NewInboxMessageAttachmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInboxMessageAttachmentService(opts ...option.RequestOption) (r InboxMessageAttachmentService) {
	r = InboxMessageAttachmentService{}
	r.Options = opts
	return
}

// Add an attachment to a draft message. You can send either a multipart/form-data
// request with the raw file content, or a JSON request with a file ID.
func (r *InboxMessageAttachmentService) New(ctx context.Context, inboxMessageID string, body InboxMessageAttachmentNewParams, opts ...option.RequestOption) (res *MessageAttachment, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboxMessageID == "" {
		err = errors.New("missing required inbox_message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("inbox_messages/%s/attachments", inboxMessageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Removes an attachment from a draft message.
func (r *InboxMessageAttachmentService) Delete(ctx context.Context, id string, body InboxMessageAttachmentDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.InboxMessageID == "" {
		err = errors.New("missing required inbox_message_id parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("inbox_messages/%s/attachments/%s", body.InboxMessageID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type InboxMessageAttachmentNewParams struct {
	MessageAttachmentCreateParams MessageAttachmentCreateParams
	paramObj
}

func (r InboxMessageAttachmentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r.MessageAttachmentCreateParams, writer)
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

type InboxMessageAttachmentDeleteParams struct {
	InboxMessageID string `path:"inbox_message_id" api:"required" json:"-"`
	paramObj
}
