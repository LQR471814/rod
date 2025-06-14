// This file is generated by "./lib/proto/generate"

package proto

/*

IO

Input/Output operations for streams produced by DevTools.

*/

// IOStreamHandle This is either obtained from another method or specified as `blob:<uuid>` where
// `<uuid>` is an UUID of a Blob.
type IOStreamHandle string

// IOClose Close the stream, discard any temporary backing storage.
type IOClose struct {
	// Handle of the stream to close.
	Handle IOStreamHandle `json:"handle"`
}

// ProtoReq name.
func (m IOClose) ProtoReq() string { return "IO.close" }

// Call sends the request.
func (m IOClose) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// IORead Read a chunk of the stream.
type IORead struct {
	// Handle of the stream to read.
	Handle IOStreamHandle `json:"handle"`

	// Offset (optional) Seek to the specified offset before reading (if not specified, proceed with offset
	// following the last read). Some types of streams may only support sequential reads.
	Offset *int `json:"offset,omitempty"`

	// Size (optional) Maximum number of bytes to read (left upon the agent discretion if not specified).
	Size *int `json:"size,omitempty"`
}

// ProtoReq name.
func (m IORead) ProtoReq() string { return "IO.read" }

// Call the request.
func (m IORead) Call(c Client) (*IOReadResult, error) {
	var res IOReadResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// IOReadResult ...
type IOReadResult struct {
	// Base64Encoded (optional) Set if the data is base64-encoded
	Base64Encoded *bool `json:"base64Encoded,omitempty"`

	// Data that were read.
	Data string `json:"data"`

	// EOF Set if the end-of-file condition occurred while reading.
	EOF bool `json:"eof"`
}

// IOResolveBlob Return UUID of Blob object specified by a remote object id.
type IOResolveBlob struct {
	// ObjectID Object id of a Blob object wrapper.
	ObjectID RuntimeRemoteObjectID `json:"objectId"`
}

// ProtoReq name.
func (m IOResolveBlob) ProtoReq() string { return "IO.resolveBlob" }

// Call the request.
func (m IOResolveBlob) Call(c Client) (*IOResolveBlobResult, error) {
	var res IOResolveBlobResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// IOResolveBlobResult ...
type IOResolveBlobResult struct {
	// UUID of the specified Blob.
	UUID string `json:"uuid"`
}
