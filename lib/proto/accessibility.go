// This file is generated by "./lib/proto/generate"

package proto

import (
	"github.com/ysmood/gson"
)

/*

Accessibility

*/

// AccessibilityAXNodeID Unique accessibility node identifier.
type AccessibilityAXNodeID string

// AccessibilityAXValueType Enum of possible property types.
type AccessibilityAXValueType string

const (
	// AccessibilityAXValueTypeBoolean enum const.
	AccessibilityAXValueTypeBoolean AccessibilityAXValueType = "boolean"

	// AccessibilityAXValueTypeTristate enum const.
	AccessibilityAXValueTypeTristate AccessibilityAXValueType = "tristate"

	// AccessibilityAXValueTypeBooleanOrUndefined enum const.
	AccessibilityAXValueTypeBooleanOrUndefined AccessibilityAXValueType = "booleanOrUndefined"

	// AccessibilityAXValueTypeIdref enum const.
	AccessibilityAXValueTypeIdref AccessibilityAXValueType = "idref"

	// AccessibilityAXValueTypeIdrefList enum const.
	AccessibilityAXValueTypeIdrefList AccessibilityAXValueType = "idrefList"

	// AccessibilityAXValueTypeInteger enum const.
	AccessibilityAXValueTypeInteger AccessibilityAXValueType = "integer"

	// AccessibilityAXValueTypeNode enum const.
	AccessibilityAXValueTypeNode AccessibilityAXValueType = "node"

	// AccessibilityAXValueTypeNodeList enum const.
	AccessibilityAXValueTypeNodeList AccessibilityAXValueType = "nodeList"

	// AccessibilityAXValueTypeNumber enum const.
	AccessibilityAXValueTypeNumber AccessibilityAXValueType = "number"

	// AccessibilityAXValueTypeString enum const.
	AccessibilityAXValueTypeString AccessibilityAXValueType = "string"

	// AccessibilityAXValueTypeComputedString enum const.
	AccessibilityAXValueTypeComputedString AccessibilityAXValueType = "computedString"

	// AccessibilityAXValueTypeToken enum const.
	AccessibilityAXValueTypeToken AccessibilityAXValueType = "token"

	// AccessibilityAXValueTypeTokenList enum const.
	AccessibilityAXValueTypeTokenList AccessibilityAXValueType = "tokenList"

	// AccessibilityAXValueTypeDomRelation enum const.
	AccessibilityAXValueTypeDomRelation AccessibilityAXValueType = "domRelation"

	// AccessibilityAXValueTypeRole enum const.
	AccessibilityAXValueTypeRole AccessibilityAXValueType = "role"

	// AccessibilityAXValueTypeInternalRole enum const.
	AccessibilityAXValueTypeInternalRole AccessibilityAXValueType = "internalRole"

	// AccessibilityAXValueTypeValueUndefined enum const.
	AccessibilityAXValueTypeValueUndefined AccessibilityAXValueType = "valueUndefined"
)

// AccessibilityAXValueSourceType Enum of possible property sources.
type AccessibilityAXValueSourceType string

const (
	// AccessibilityAXValueSourceTypeAttribute enum const.
	AccessibilityAXValueSourceTypeAttribute AccessibilityAXValueSourceType = "attribute"

	// AccessibilityAXValueSourceTypeImplicit enum const.
	AccessibilityAXValueSourceTypeImplicit AccessibilityAXValueSourceType = "implicit"

	// AccessibilityAXValueSourceTypeStyle enum const.
	AccessibilityAXValueSourceTypeStyle AccessibilityAXValueSourceType = "style"

	// AccessibilityAXValueSourceTypeContents enum const.
	AccessibilityAXValueSourceTypeContents AccessibilityAXValueSourceType = "contents"

	// AccessibilityAXValueSourceTypePlaceholder enum const.
	AccessibilityAXValueSourceTypePlaceholder AccessibilityAXValueSourceType = "placeholder"

	// AccessibilityAXValueSourceTypeRelatedElement enum const.
	AccessibilityAXValueSourceTypeRelatedElement AccessibilityAXValueSourceType = "relatedElement"
)

// AccessibilityAXValueNativeSourceType Enum of possible native property sources (as a subtype of a particular AXValueSourceType).
type AccessibilityAXValueNativeSourceType string

const (
	// AccessibilityAXValueNativeSourceTypeDescription enum const.
	AccessibilityAXValueNativeSourceTypeDescription AccessibilityAXValueNativeSourceType = "description"

	// AccessibilityAXValueNativeSourceTypeFigcaption enum const.
	AccessibilityAXValueNativeSourceTypeFigcaption AccessibilityAXValueNativeSourceType = "figcaption"

	// AccessibilityAXValueNativeSourceTypeLabel enum const.
	AccessibilityAXValueNativeSourceTypeLabel AccessibilityAXValueNativeSourceType = "label"

	// AccessibilityAXValueNativeSourceTypeLabelfor enum const.
	AccessibilityAXValueNativeSourceTypeLabelfor AccessibilityAXValueNativeSourceType = "labelfor"

	// AccessibilityAXValueNativeSourceTypeLabelwrapped enum const.
	AccessibilityAXValueNativeSourceTypeLabelwrapped AccessibilityAXValueNativeSourceType = "labelwrapped"

	// AccessibilityAXValueNativeSourceTypeLegend enum const.
	AccessibilityAXValueNativeSourceTypeLegend AccessibilityAXValueNativeSourceType = "legend"

	// AccessibilityAXValueNativeSourceTypeRubyannotation enum const.
	AccessibilityAXValueNativeSourceTypeRubyannotation AccessibilityAXValueNativeSourceType = "rubyannotation"

	// AccessibilityAXValueNativeSourceTypeTablecaption enum const.
	AccessibilityAXValueNativeSourceTypeTablecaption AccessibilityAXValueNativeSourceType = "tablecaption"

	// AccessibilityAXValueNativeSourceTypeTitle enum const.
	AccessibilityAXValueNativeSourceTypeTitle AccessibilityAXValueNativeSourceType = "title"

	// AccessibilityAXValueNativeSourceTypeOther enum const.
	AccessibilityAXValueNativeSourceTypeOther AccessibilityAXValueNativeSourceType = "other"
)

// AccessibilityAXValueSource A single source for a computed AX property.
type AccessibilityAXValueSource struct {
	// Type What type of source this is.
	Type AccessibilityAXValueSourceType `json:"type"`

	// Value (optional) The value of this property source.
	Value *AccessibilityAXValue `json:"value,omitempty"`

	// Attribute (optional) The name of the relevant attribute, if any.
	Attribute string `json:"attribute,omitempty"`

	// AttributeValue (optional) The value of the relevant attribute, if any.
	AttributeValue *AccessibilityAXValue `json:"attributeValue,omitempty"`

	// Superseded (optional) Whether this source is superseded by a higher priority source.
	Superseded *bool `json:"superseded,omitempty"`

	// NativeSource (optional) The native markup source for this value, e.g. a `<label>` element.
	NativeSource AccessibilityAXValueNativeSourceType `json:"nativeSource,omitempty"`

	// NativeSourceValue (optional) The value, such as a node or node list, of the native source.
	NativeSourceValue *AccessibilityAXValue `json:"nativeSourceValue,omitempty"`

	// Invalid (optional) Whether the value for this property is invalid.
	Invalid *bool `json:"invalid,omitempty"`

	// InvalidReason (optional) Reason for the value being invalid, if it is.
	InvalidReason string `json:"invalidReason,omitempty"`
}

// AccessibilityAXRelatedNode ...
type AccessibilityAXRelatedNode struct {
	// BackendDOMNodeID The BackendNodeId of the related DOM node.
	BackendDOMNodeID DOMBackendNodeID `json:"backendDOMNodeId"`

	// Idref (optional) The IDRef value provided, if any.
	Idref string `json:"idref,omitempty"`

	// Text (optional) The text alternative of this node in the current context.
	Text string `json:"text,omitempty"`
}

// AccessibilityAXProperty ...
type AccessibilityAXProperty struct {
	// Name The name of this property.
	Name AccessibilityAXPropertyName `json:"name"`

	// Value The value of this property.
	Value *AccessibilityAXValue `json:"value"`
}

// AccessibilityAXValue A single computed AX property.
type AccessibilityAXValue struct {
	// Type The type of this value.
	Type AccessibilityAXValueType `json:"type"`

	// Value (optional) The computed value of this property.
	Value gson.JSON `json:"value,omitempty"`

	// RelatedNodes (optional) One or more related nodes, if applicable.
	RelatedNodes []*AccessibilityAXRelatedNode `json:"relatedNodes,omitempty"`

	// Sources (optional) The sources which contributed to the computation of this property.
	Sources []*AccessibilityAXValueSource `json:"sources,omitempty"`
}

// AccessibilityAXPropertyName Values of AXProperty name:
// - from 'busy' to 'roledescription': states which apply to every AX node
// - from 'live' to 'root': attributes which apply to nodes in live regions
// - from 'autocomplete' to 'valuetext': attributes which apply to widgets
// - from 'checked' to 'selected': states which apply to widgets
// - from 'activedescendant' to 'owns' - relationships between elements other than parent/child/sibling.
type AccessibilityAXPropertyName string

const (
	// AccessibilityAXPropertyNameBusy enum const.
	AccessibilityAXPropertyNameBusy AccessibilityAXPropertyName = "busy"

	// AccessibilityAXPropertyNameDisabled enum const.
	AccessibilityAXPropertyNameDisabled AccessibilityAXPropertyName = "disabled"

	// AccessibilityAXPropertyNameEditable enum const.
	AccessibilityAXPropertyNameEditable AccessibilityAXPropertyName = "editable"

	// AccessibilityAXPropertyNameFocusable enum const.
	AccessibilityAXPropertyNameFocusable AccessibilityAXPropertyName = "focusable"

	// AccessibilityAXPropertyNameFocused enum const.
	AccessibilityAXPropertyNameFocused AccessibilityAXPropertyName = "focused"

	// AccessibilityAXPropertyNameHidden enum const.
	AccessibilityAXPropertyNameHidden AccessibilityAXPropertyName = "hidden"

	// AccessibilityAXPropertyNameHiddenRoot enum const.
	AccessibilityAXPropertyNameHiddenRoot AccessibilityAXPropertyName = "hiddenRoot"

	// AccessibilityAXPropertyNameInvalid enum const.
	AccessibilityAXPropertyNameInvalid AccessibilityAXPropertyName = "invalid"

	// AccessibilityAXPropertyNameKeyshortcuts enum const.
	AccessibilityAXPropertyNameKeyshortcuts AccessibilityAXPropertyName = "keyshortcuts"

	// AccessibilityAXPropertyNameSettable enum const.
	AccessibilityAXPropertyNameSettable AccessibilityAXPropertyName = "settable"

	// AccessibilityAXPropertyNameRoledescription enum const.
	AccessibilityAXPropertyNameRoledescription AccessibilityAXPropertyName = "roledescription"

	// AccessibilityAXPropertyNameLive enum const.
	AccessibilityAXPropertyNameLive AccessibilityAXPropertyName = "live"

	// AccessibilityAXPropertyNameAtomic enum const.
	AccessibilityAXPropertyNameAtomic AccessibilityAXPropertyName = "atomic"

	// AccessibilityAXPropertyNameRelevant enum const.
	AccessibilityAXPropertyNameRelevant AccessibilityAXPropertyName = "relevant"

	// AccessibilityAXPropertyNameRoot enum const.
	AccessibilityAXPropertyNameRoot AccessibilityAXPropertyName = "root"

	// AccessibilityAXPropertyNameAutocomplete enum const.
	AccessibilityAXPropertyNameAutocomplete AccessibilityAXPropertyName = "autocomplete"

	// AccessibilityAXPropertyNameHasPopup enum const.
	AccessibilityAXPropertyNameHasPopup AccessibilityAXPropertyName = "hasPopup"

	// AccessibilityAXPropertyNameLevel enum const.
	AccessibilityAXPropertyNameLevel AccessibilityAXPropertyName = "level"

	// AccessibilityAXPropertyNameMultiselectable enum const.
	AccessibilityAXPropertyNameMultiselectable AccessibilityAXPropertyName = "multiselectable"

	// AccessibilityAXPropertyNameOrientation enum const.
	AccessibilityAXPropertyNameOrientation AccessibilityAXPropertyName = "orientation"

	// AccessibilityAXPropertyNameMultiline enum const.
	AccessibilityAXPropertyNameMultiline AccessibilityAXPropertyName = "multiline"

	// AccessibilityAXPropertyNameReadonly enum const.
	AccessibilityAXPropertyNameReadonly AccessibilityAXPropertyName = "readonly"

	// AccessibilityAXPropertyNameRequired enum const.
	AccessibilityAXPropertyNameRequired AccessibilityAXPropertyName = "required"

	// AccessibilityAXPropertyNameValuemin enum const.
	AccessibilityAXPropertyNameValuemin AccessibilityAXPropertyName = "valuemin"

	// AccessibilityAXPropertyNameValuemax enum const.
	AccessibilityAXPropertyNameValuemax AccessibilityAXPropertyName = "valuemax"

	// AccessibilityAXPropertyNameValuetext enum const.
	AccessibilityAXPropertyNameValuetext AccessibilityAXPropertyName = "valuetext"

	// AccessibilityAXPropertyNameChecked enum const.
	AccessibilityAXPropertyNameChecked AccessibilityAXPropertyName = "checked"

	// AccessibilityAXPropertyNameExpanded enum const.
	AccessibilityAXPropertyNameExpanded AccessibilityAXPropertyName = "expanded"

	// AccessibilityAXPropertyNameModal enum const.
	AccessibilityAXPropertyNameModal AccessibilityAXPropertyName = "modal"

	// AccessibilityAXPropertyNamePressed enum const.
	AccessibilityAXPropertyNamePressed AccessibilityAXPropertyName = "pressed"

	// AccessibilityAXPropertyNameSelected enum const.
	AccessibilityAXPropertyNameSelected AccessibilityAXPropertyName = "selected"

	// AccessibilityAXPropertyNameActivedescendant enum const.
	AccessibilityAXPropertyNameActivedescendant AccessibilityAXPropertyName = "activedescendant"

	// AccessibilityAXPropertyNameControls enum const.
	AccessibilityAXPropertyNameControls AccessibilityAXPropertyName = "controls"

	// AccessibilityAXPropertyNameDescribedby enum const.
	AccessibilityAXPropertyNameDescribedby AccessibilityAXPropertyName = "describedby"

	// AccessibilityAXPropertyNameDetails enum const.
	AccessibilityAXPropertyNameDetails AccessibilityAXPropertyName = "details"

	// AccessibilityAXPropertyNameErrormessage enum const.
	AccessibilityAXPropertyNameErrormessage AccessibilityAXPropertyName = "errormessage"

	// AccessibilityAXPropertyNameFlowto enum const.
	AccessibilityAXPropertyNameFlowto AccessibilityAXPropertyName = "flowto"

	// AccessibilityAXPropertyNameLabelledby enum const.
	AccessibilityAXPropertyNameLabelledby AccessibilityAXPropertyName = "labelledby"

	// AccessibilityAXPropertyNameOwns enum const.
	AccessibilityAXPropertyNameOwns AccessibilityAXPropertyName = "owns"

	// AccessibilityAXPropertyNameURL enum const.
	AccessibilityAXPropertyNameURL AccessibilityAXPropertyName = "url"
)

// AccessibilityAXNode A node in the accessibility tree.
type AccessibilityAXNode struct {
	// NodeID Unique identifier for this node.
	NodeID AccessibilityAXNodeID `json:"nodeId"`

	// Ignored Whether this node is ignored for accessibility
	Ignored bool `json:"ignored"`

	// IgnoredReasons (optional) Collection of reasons why this node is hidden.
	IgnoredReasons []*AccessibilityAXProperty `json:"ignoredReasons,omitempty"`

	// Role (optional) This `Node`'s role, whether explicit or implicit.
	Role *AccessibilityAXValue `json:"role,omitempty"`

	// ChromeRole (optional) This `Node`'s Chrome raw role.
	ChromeRole *AccessibilityAXValue `json:"chromeRole,omitempty"`

	// Name (optional) The accessible name for this `Node`.
	Name *AccessibilityAXValue `json:"name,omitempty"`

	// Description (optional) The accessible description for this `Node`.
	Description *AccessibilityAXValue `json:"description,omitempty"`

	// Value (optional) The value for this `Node`.
	Value *AccessibilityAXValue `json:"value,omitempty"`

	// Properties (optional) All other properties
	Properties []*AccessibilityAXProperty `json:"properties,omitempty"`

	// ParentID (optional) ID for this node's parent.
	ParentID AccessibilityAXNodeID `json:"parentId,omitempty"`

	// ChildIDs (optional) IDs for each of this node's child nodes.
	ChildIDs []AccessibilityAXNodeID `json:"childIds,omitempty"`

	// BackendDOMNodeID (optional) The backend ID for the associated DOM node, if any.
	BackendDOMNodeID DOMBackendNodeID `json:"backendDOMNodeId,omitempty"`

	// FrameID (optional) The frame ID for the frame associated with this nodes document.
	FrameID PageFrameID `json:"frameId,omitempty"`
}

// AccessibilityDisable Disables the accessibility domain.
type AccessibilityDisable struct{}

// ProtoReq name.
func (m AccessibilityDisable) ProtoReq() string { return "Accessibility.disable" }

// Call sends the request.
func (m AccessibilityDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// AccessibilityEnable Enables the accessibility domain which causes `AXNodeId`s to remain consistent between method calls.
// This turns on accessibility for the page, which can impact performance until accessibility is disabled.
type AccessibilityEnable struct{}

// ProtoReq name.
func (m AccessibilityEnable) ProtoReq() string { return "Accessibility.enable" }

// Call sends the request.
func (m AccessibilityEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// AccessibilityGetPartialAXTree (experimental) Fetches the accessibility node and partial accessibility tree for this DOM node, if it exists.
type AccessibilityGetPartialAXTree struct {
	// NodeID (optional) Identifier of the node to get the partial accessibility tree for.
	NodeID DOMNodeID `json:"nodeId,omitempty"`

	// BackendNodeID (optional) Identifier of the backend node to get the partial accessibility tree for.
	BackendNodeID DOMBackendNodeID `json:"backendNodeId,omitempty"`

	// ObjectID (optional) JavaScript object id of the node wrapper to get the partial accessibility tree for.
	ObjectID RuntimeRemoteObjectID `json:"objectId,omitempty"`

	// FetchRelatives (optional) Whether to fetch this node's ancestors, siblings and children. Defaults to true.
	FetchRelatives *bool `json:"fetchRelatives,omitempty"`
}

// ProtoReq name.
func (m AccessibilityGetPartialAXTree) ProtoReq() string { return "Accessibility.getPartialAXTree" }

// Call the request.
func (m AccessibilityGetPartialAXTree) Call(c Client) (*AccessibilityGetPartialAXTreeResult, error) {
	var res AccessibilityGetPartialAXTreeResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// AccessibilityGetPartialAXTreeResult (experimental) ...
type AccessibilityGetPartialAXTreeResult struct {
	// Nodes The `Accessibility.AXNode` for this DOM node, if it exists, plus its ancestors, siblings and
	// children, if requested.
	Nodes []*AccessibilityAXNode `json:"nodes"`
}

// AccessibilityGetFullAXTree (experimental) Fetches the entire accessibility tree for the root Document.
type AccessibilityGetFullAXTree struct {
	// Depth (optional) The maximum depth at which descendants of the root node should be retrieved.
	// If omitted, the full tree is returned.
	Depth *int `json:"depth,omitempty"`

	// FrameID (optional) The frame for whose document the AX tree should be retrieved.
	// If omitted, the root frame is used.
	FrameID PageFrameID `json:"frameId,omitempty"`
}

// ProtoReq name.
func (m AccessibilityGetFullAXTree) ProtoReq() string { return "Accessibility.getFullAXTree" }

// Call the request.
func (m AccessibilityGetFullAXTree) Call(c Client) (*AccessibilityGetFullAXTreeResult, error) {
	var res AccessibilityGetFullAXTreeResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// AccessibilityGetFullAXTreeResult (experimental) ...
type AccessibilityGetFullAXTreeResult struct {
	// Nodes ...
	Nodes []*AccessibilityAXNode `json:"nodes"`
}

// AccessibilityGetRootAXNode (experimental) Fetches the root node.
// Requires `enable()` to have been called previously.
type AccessibilityGetRootAXNode struct {
	// FrameID (optional) The frame in whose document the node resides.
	// If omitted, the root frame is used.
	FrameID PageFrameID `json:"frameId,omitempty"`
}

// ProtoReq name.
func (m AccessibilityGetRootAXNode) ProtoReq() string { return "Accessibility.getRootAXNode" }

// Call the request.
func (m AccessibilityGetRootAXNode) Call(c Client) (*AccessibilityGetRootAXNodeResult, error) {
	var res AccessibilityGetRootAXNodeResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// AccessibilityGetRootAXNodeResult (experimental) ...
type AccessibilityGetRootAXNodeResult struct {
	// Node ...
	Node *AccessibilityAXNode `json:"node"`
}

// AccessibilityGetAXNodeAndAncestors (experimental) Fetches a node and all ancestors up to and including the root.
// Requires `enable()` to have been called previously.
type AccessibilityGetAXNodeAndAncestors struct {
	// NodeID (optional) Identifier of the node to get.
	NodeID DOMNodeID `json:"nodeId,omitempty"`

	// BackendNodeID (optional) Identifier of the backend node to get.
	BackendNodeID DOMBackendNodeID `json:"backendNodeId,omitempty"`

	// ObjectID (optional) JavaScript object id of the node wrapper to get.
	ObjectID RuntimeRemoteObjectID `json:"objectId,omitempty"`
}

// ProtoReq name.
func (m AccessibilityGetAXNodeAndAncestors) ProtoReq() string {
	return "Accessibility.getAXNodeAndAncestors"
}

// Call the request.
func (m AccessibilityGetAXNodeAndAncestors) Call(c Client) (*AccessibilityGetAXNodeAndAncestorsResult, error) {
	var res AccessibilityGetAXNodeAndAncestorsResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// AccessibilityGetAXNodeAndAncestorsResult (experimental) ...
type AccessibilityGetAXNodeAndAncestorsResult struct {
	// Nodes ...
	Nodes []*AccessibilityAXNode `json:"nodes"`
}

// AccessibilityGetChildAXNodes (experimental) Fetches a particular accessibility node by AXNodeId.
// Requires `enable()` to have been called previously.
type AccessibilityGetChildAXNodes struct {
	// ID ...
	ID AccessibilityAXNodeID `json:"id"`

	// FrameID (optional) The frame in whose document the node resides.
	// If omitted, the root frame is used.
	FrameID PageFrameID `json:"frameId,omitempty"`
}

// ProtoReq name.
func (m AccessibilityGetChildAXNodes) ProtoReq() string { return "Accessibility.getChildAXNodes" }

// Call the request.
func (m AccessibilityGetChildAXNodes) Call(c Client) (*AccessibilityGetChildAXNodesResult, error) {
	var res AccessibilityGetChildAXNodesResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// AccessibilityGetChildAXNodesResult (experimental) ...
type AccessibilityGetChildAXNodesResult struct {
	// Nodes ...
	Nodes []*AccessibilityAXNode `json:"nodes"`
}

// AccessibilityQueryAXTree (experimental) Query a DOM node's accessibility subtree for accessible name and role.
// This command computes the name and role for all nodes in the subtree, including those that are
// ignored for accessibility, and returns those that match the specified name and role. If no DOM
// node is specified, or the DOM node does not exist, the command returns an error. If neither
// `accessibleName` or `role` is specified, it returns all the accessibility nodes in the subtree.
type AccessibilityQueryAXTree struct {
	// NodeID (optional) Identifier of the node for the root to query.
	NodeID DOMNodeID `json:"nodeId,omitempty"`

	// BackendNodeID (optional) Identifier of the backend node for the root to query.
	BackendNodeID DOMBackendNodeID `json:"backendNodeId,omitempty"`

	// ObjectID (optional) JavaScript object id of the node wrapper for the root to query.
	ObjectID RuntimeRemoteObjectID `json:"objectId,omitempty"`

	// AccessibleName (optional) Find nodes with this computed name.
	AccessibleName string `json:"accessibleName,omitempty"`

	// Role (optional) Find nodes with this computed role.
	Role string `json:"role,omitempty"`
}

// ProtoReq name.
func (m AccessibilityQueryAXTree) ProtoReq() string { return "Accessibility.queryAXTree" }

// Call the request.
func (m AccessibilityQueryAXTree) Call(c Client) (*AccessibilityQueryAXTreeResult, error) {
	var res AccessibilityQueryAXTreeResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// AccessibilityQueryAXTreeResult (experimental) ...
type AccessibilityQueryAXTreeResult struct {
	// Nodes A list of `Accessibility.AXNode` matching the specified attributes,
	// including nodes that are ignored for accessibility.
	Nodes []*AccessibilityAXNode `json:"nodes"`
}

// AccessibilityLoadComplete (experimental) The loadComplete event mirrors the load complete event sent by the browser to assistive
// technology when the web page has finished loading.
type AccessibilityLoadComplete struct {
	// Root New document root node.
	Root *AccessibilityAXNode `json:"root"`
}

// ProtoEvent name.
func (evt AccessibilityLoadComplete) ProtoEvent() string {
	return "Accessibility.loadComplete"
}

// AccessibilityNodesUpdated (experimental) The nodesUpdated event is sent every time a previously requested node has changed the in tree.
type AccessibilityNodesUpdated struct {
	// Nodes Updated node data.
	Nodes []*AccessibilityAXNode `json:"nodes"`
}

// ProtoEvent name.
func (evt AccessibilityNodesUpdated) ProtoEvent() string {
	return "Accessibility.nodesUpdated"
}
