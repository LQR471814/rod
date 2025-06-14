// This file is generated by "./lib/proto/generate"

package proto

/*

Fetch

A domain for letting clients substitute browser's network layer with client code.

*/

// FetchRequestID Unique request identifier.
type FetchRequestID string

// FetchRequestStage Stages of the request to handle. Request will intercept before the request is
// sent. Response will intercept after the response is received (but before response
// body is received).
type FetchRequestStage string

const (
	// FetchRequestStageRequest enum const.
	FetchRequestStageRequest FetchRequestStage = "Request"

	// FetchRequestStageResponse enum const.
	FetchRequestStageResponse FetchRequestStage = "Response"
)

// FetchRequestPattern ...
type FetchRequestPattern struct {
	// URLPattern (optional) Wildcards (`'*'` -> zero or more, `'?'` -> exactly one) are allowed. Escape character is
	// backslash. Omitting is equivalent to `"*"`.
	URLPattern string `json:"urlPattern,omitempty"`

	// ResourceType (optional) If set, only requests for matching resource types will be intercepted.
	ResourceType NetworkResourceType `json:"resourceType,omitempty"`

	// RequestStage (optional) Stage at which to begin intercepting requests. Default is Request.
	RequestStage FetchRequestStage `json:"requestStage,omitempty"`
}

// FetchHeaderEntry Response HTTP header entry.
type FetchHeaderEntry struct {
	// Name ...
	Name string `json:"name"`

	// Value ...
	Value string `json:"value"`
}

// FetchAuthChallengeSource enum.
type FetchAuthChallengeSource string

const (
	// FetchAuthChallengeSourceServer enum const.
	FetchAuthChallengeSourceServer FetchAuthChallengeSource = "Server"

	// FetchAuthChallengeSourceProxy enum const.
	FetchAuthChallengeSourceProxy FetchAuthChallengeSource = "Proxy"
)

// FetchAuthChallenge Authorization challenge for HTTP status code 401 or 407.
type FetchAuthChallenge struct {
	// Source (optional) Source of the authentication challenge.
	Source FetchAuthChallengeSource `json:"source,omitempty"`

	// Origin of the challenger.
	Origin string `json:"origin"`

	// Scheme The authentication scheme used, such as basic or digest
	Scheme string `json:"scheme"`

	// Realm The realm of the challenge. May be empty.
	Realm string `json:"realm"`
}

// FetchAuthChallengeResponseResponse enum.
type FetchAuthChallengeResponseResponse string

const (
	// FetchAuthChallengeResponseResponseDefault enum const.
	FetchAuthChallengeResponseResponseDefault FetchAuthChallengeResponseResponse = "Default"

	// FetchAuthChallengeResponseResponseCancelAuth enum const.
	FetchAuthChallengeResponseResponseCancelAuth FetchAuthChallengeResponseResponse = "CancelAuth"

	// FetchAuthChallengeResponseResponseProvideCredentials enum const.
	FetchAuthChallengeResponseResponseProvideCredentials FetchAuthChallengeResponseResponse = "ProvideCredentials"
)

// FetchAuthChallengeResponse Response to an AuthChallenge.
type FetchAuthChallengeResponse struct {
	// Response The decision on what to do in response to the authorization challenge.  Default means
	// deferring to the default behavior of the net stack, which will likely either the Cancel
	// authentication or display a popup dialog box.
	Response FetchAuthChallengeResponseResponse `json:"response"`

	// Username (optional) The username to provide, possibly empty. Should only be set if response is
	// ProvideCredentials.
	Username string `json:"username,omitempty"`

	// Password (optional) The password to provide, possibly empty. Should only be set if response is
	// ProvideCredentials.
	Password string `json:"password,omitempty"`
}

// FetchDisable Disables the fetch domain.
type FetchDisable struct{}

// ProtoReq name.
func (m FetchDisable) ProtoReq() string { return "Fetch.disable" }

// Call sends the request.
func (m FetchDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FetchEnable Enables issuing of requestPaused events. A request will be paused until client
// calls one of failRequest, fulfillRequest or continueRequest/continueWithAuth.
type FetchEnable struct {
	// Patterns (optional) If specified, only requests matching any of these patterns will produce
	// fetchRequested event and will be paused until clients response. If not set,
	// all requests will be affected.
	Patterns []*FetchRequestPattern `json:"patterns,omitempty"`

	// HandleAuthRequests (optional) If true, authRequired events will be issued and requests will be paused
	// expecting a call to continueWithAuth.
	HandleAuthRequests *bool `json:"handleAuthRequests,omitempty"`
}

// ProtoReq name.
func (m FetchEnable) ProtoReq() string { return "Fetch.enable" }

// Call sends the request.
func (m FetchEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FetchFailRequest Causes the request to fail with specified reason.
type FetchFailRequest struct {
	// RequestID An id the client received in requestPaused event.
	RequestID FetchRequestID `json:"requestId"`

	// ErrorReason Causes the request to fail with the given reason.
	ErrorReason NetworkErrorReason `json:"errorReason"`
}

// ProtoReq name.
func (m FetchFailRequest) ProtoReq() string { return "Fetch.failRequest" }

// Call sends the request.
func (m FetchFailRequest) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FetchFulfillRequest Provides response to the request.
type FetchFulfillRequest struct {
	// RequestID An id the client received in requestPaused event.
	RequestID FetchRequestID `json:"requestId"`

	// ResponseCode An HTTP response code.
	ResponseCode int `json:"responseCode"`

	// ResponseHeaders (optional) Response headers.
	ResponseHeaders []*FetchHeaderEntry `json:"responseHeaders,omitempty"`

	// BinaryResponseHeaders (optional) Alternative way of specifying response headers as a \0-separated
	// series of name: value pairs. Prefer the above method unless you
	// need to represent some non-UTF8 values that can't be transmitted
	// over the protocol as text.
	BinaryResponseHeaders []byte `json:"binaryResponseHeaders,omitempty"`

	// Body A response body. If absent, original response body will be used if
	// the request is intercepted at the response stage and empty body
	// will be used if the request is intercepted at the request stage.
	Body []byte `json:"body"`

	// ResponsePhrase (optional) A textual representation of responseCode.
	// If absent, a standard phrase matching responseCode is used.
	ResponsePhrase string `json:"responsePhrase,omitempty"`
}

// ProtoReq name.
func (m FetchFulfillRequest) ProtoReq() string { return "Fetch.fulfillRequest" }

// Call sends the request.
func (m FetchFulfillRequest) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FetchContinueRequest Continues the request, optionally modifying some of its parameters.
type FetchContinueRequest struct {
	// RequestID An id the client received in requestPaused event.
	RequestID FetchRequestID `json:"requestId"`

	// URL (optional) If set, the request url will be modified in a way that's not observable by page.
	URL string `json:"url,omitempty"`

	// Method (optional) If set, the request method is overridden.
	Method string `json:"method,omitempty"`

	// PostData (optional) If set, overrides the post data in the request.
	PostData []byte `json:"postData,omitempty"`

	// Headers (optional) If set, overrides the request headers. Note that the overrides do not
	// extend to subsequent redirect hops, if a redirect happens. Another override
	// may be applied to a different request produced by a redirect.
	Headers []*FetchHeaderEntry `json:"headers,omitempty"`

	// InterceptResponse (experimental) (optional) If set, overrides response interception behavior for this request.
	InterceptResponse *bool `json:"interceptResponse,omitempty"`
}

// ProtoReq name.
func (m FetchContinueRequest) ProtoReq() string { return "Fetch.continueRequest" }

// Call sends the request.
func (m FetchContinueRequest) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FetchContinueWithAuth Continues a request supplying authChallengeResponse following authRequired event.
type FetchContinueWithAuth struct {
	// RequestID An id the client received in authRequired event.
	RequestID FetchRequestID `json:"requestId"`

	// AuthChallengeResponse Response to  with an authChallenge.
	AuthChallengeResponse *FetchAuthChallengeResponse `json:"authChallengeResponse"`
}

// ProtoReq name.
func (m FetchContinueWithAuth) ProtoReq() string { return "Fetch.continueWithAuth" }

// Call sends the request.
func (m FetchContinueWithAuth) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FetchContinueResponse (experimental) Continues loading of the paused response, optionally modifying the
// response headers. If either responseCode or headers are modified, all of them
// must be present.
type FetchContinueResponse struct {
	// RequestID An id the client received in requestPaused event.
	RequestID FetchRequestID `json:"requestId"`

	// ResponseCode (optional) An HTTP response code. If absent, original response code will be used.
	ResponseCode *int `json:"responseCode,omitempty"`

	// ResponsePhrase (optional) A textual representation of responseCode.
	// If absent, a standard phrase matching responseCode is used.
	ResponsePhrase string `json:"responsePhrase,omitempty"`

	// ResponseHeaders (optional) Response headers. If absent, original response headers will be used.
	ResponseHeaders []*FetchHeaderEntry `json:"responseHeaders,omitempty"`

	// BinaryResponseHeaders (optional) Alternative way of specifying response headers as a \0-separated
	// series of name: value pairs. Prefer the above method unless you
	// need to represent some non-UTF8 values that can't be transmitted
	// over the protocol as text.
	BinaryResponseHeaders []byte `json:"binaryResponseHeaders,omitempty"`
}

// ProtoReq name.
func (m FetchContinueResponse) ProtoReq() string { return "Fetch.continueResponse" }

// Call sends the request.
func (m FetchContinueResponse) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FetchGetResponseBody Causes the body of the response to be received from the server and
// returned as a single string. May only be issued for a request that
// is paused in the Response stage and is mutually exclusive with
// takeResponseBodyForInterceptionAsStream. Calling other methods that
// affect the request or disabling fetch domain before body is received
// results in an undefined behavior.
// Note that the response body is not available for redirects. Requests
// paused in the _redirect received_ state may be differentiated by
// `responseCode` and presence of `location` response header, see
// comments to `requestPaused` for details.
type FetchGetResponseBody struct {
	// RequestID Identifier for the intercepted request to get body for.
	RequestID FetchRequestID `json:"requestId"`
}

// ProtoReq name.
func (m FetchGetResponseBody) ProtoReq() string { return "Fetch.getResponseBody" }

// Call the request.
func (m FetchGetResponseBody) Call(c Client) (*FetchGetResponseBodyResult, error) {
	var res FetchGetResponseBodyResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// FetchGetResponseBodyResult ...
type FetchGetResponseBodyResult struct {
	// Body Response body.
	Body string `json:"body"`

	// Base64Encoded True, if content was sent as base64.
	Base64Encoded bool `json:"base64Encoded"`
}

// FetchTakeResponseBodyAsStream Returns a handle to the stream representing the response body.
// The request must be paused in the HeadersReceived stage.
// Note that after this command the request can't be continued
// as is -- client either needs to cancel it or to provide the
// response body.
// The stream only supports sequential read, IO.read will fail if the position
// is specified.
// This method is mutually exclusive with getResponseBody.
// Calling other methods that affect the request or disabling fetch
// domain before body is received results in an undefined behavior.
type FetchTakeResponseBodyAsStream struct {
	// RequestID ...
	RequestID FetchRequestID `json:"requestId"`
}

// ProtoReq name.
func (m FetchTakeResponseBodyAsStream) ProtoReq() string { return "Fetch.takeResponseBodyAsStream" }

// Call the request.
func (m FetchTakeResponseBodyAsStream) Call(c Client) (*FetchTakeResponseBodyAsStreamResult, error) {
	var res FetchTakeResponseBodyAsStreamResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// FetchTakeResponseBodyAsStreamResult ...
type FetchTakeResponseBodyAsStreamResult struct {
	// Stream ...
	Stream IOStreamHandle `json:"stream"`
}

// FetchRequestPaused Issued when the domain is enabled and the request URL matches the
// specified filter. The request is paused until the client responds
// with one of continueRequest, failRequest or fulfillRequest.
// The stage of the request can be determined by presence of responseErrorReason
// and responseStatusCode -- the request is at the response stage if either
// of these fields is present and in the request stage otherwise.
// Redirect responses and subsequent requests are reported similarly to regular
// responses and requests. Redirect responses may be distinguished by the value
// of `responseStatusCode` (which is one of 301, 302, 303, 307, 308) along with
// presence of the `location` header. Requests resulting from a redirect will
// have `redirectedRequestId` field set.
type FetchRequestPaused struct {
	// RequestID Each request the page makes will have a unique id.
	RequestID FetchRequestID `json:"requestId"`

	// Request The details of the request.
	Request *NetworkRequest `json:"request"`

	// FrameID The id of the frame that initiated the request.
	FrameID PageFrameID `json:"frameId"`

	// ResourceType How the requested resource will be used.
	ResourceType NetworkResourceType `json:"resourceType"`

	// ResponseErrorReason (optional) Response error if intercepted at response stage.
	ResponseErrorReason NetworkErrorReason `json:"responseErrorReason,omitempty"`

	// ResponseStatusCode (optional) Response code if intercepted at response stage.
	ResponseStatusCode *int `json:"responseStatusCode,omitempty"`

	// ResponseStatusText (optional) Response status text if intercepted at response stage.
	ResponseStatusText string `json:"responseStatusText,omitempty"`

	// ResponseHeaders (optional) Response headers if intercepted at the response stage.
	ResponseHeaders []*FetchHeaderEntry `json:"responseHeaders,omitempty"`

	// NetworkID (optional) If the intercepted request had a corresponding Network.requestWillBeSent event fired for it,
	// then this networkId will be the same as the requestId present in the requestWillBeSent event.
	NetworkID NetworkRequestID `json:"networkId,omitempty"`

	// RedirectedRequestID (experimental) (optional) If the request is due to a redirect response from the server, the id of the request that
	// has caused the redirect.
	RedirectedRequestID FetchRequestID `json:"redirectedRequestId,omitempty"`
}

// ProtoEvent name.
func (evt FetchRequestPaused) ProtoEvent() string {
	return "Fetch.requestPaused"
}

// FetchAuthRequired Issued when the domain is enabled with handleAuthRequests set to true.
// The request is paused until client responds with continueWithAuth.
type FetchAuthRequired struct {
	// RequestID Each request the page makes will have a unique id.
	RequestID FetchRequestID `json:"requestId"`

	// Request The details of the request.
	Request *NetworkRequest `json:"request"`

	// FrameID The id of the frame that initiated the request.
	FrameID PageFrameID `json:"frameId"`

	// ResourceType How the requested resource will be used.
	ResourceType NetworkResourceType `json:"resourceType"`

	// AuthChallenge Details of the Authorization Challenge encountered.
	// If this is set, client should respond with continueRequest that
	// contains AuthChallengeResponse.
	AuthChallenge *FetchAuthChallenge `json:"authChallenge"`
}

// ProtoEvent name.
func (evt FetchAuthRequired) ProtoEvent() string {
	return "Fetch.authRequired"
}
