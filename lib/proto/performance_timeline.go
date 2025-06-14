// This file is generated by "./lib/proto/generate"

package proto

/*

PerformanceTimeline

Reporting of performance timeline events, as specified in
https://w3c.github.io/performance-timeline/#dom-performanceobserver.

*/

// PerformanceTimelineLargestContentfulPaint See https://github.com/WICG/LargestContentfulPaint and largest_contentful_paint.idl.
type PerformanceTimelineLargestContentfulPaint struct {
	// RenderTime ...
	RenderTime TimeSinceEpoch `json:"renderTime"`

	// LoadTime ...
	LoadTime TimeSinceEpoch `json:"loadTime"`

	// Size The number of pixels being painted.
	Size float64 `json:"size"`

	// ElementID (optional) The id attribute of the element, if available.
	ElementID string `json:"elementId,omitempty"`

	// URL (optional) The URL of the image (may be trimmed).
	URL string `json:"url,omitempty"`

	// NodeID (optional) ...
	NodeID DOMBackendNodeID `json:"nodeId,omitempty"`
}

// PerformanceTimelineLayoutShiftAttribution ...
type PerformanceTimelineLayoutShiftAttribution struct {
	// PreviousRect ...
	PreviousRect *DOMRect `json:"previousRect"`

	// CurrentRect ...
	CurrentRect *DOMRect `json:"currentRect"`

	// NodeID (optional) ...
	NodeID DOMBackendNodeID `json:"nodeId,omitempty"`
}

// PerformanceTimelineLayoutShift See https://wicg.github.io/layout-instability/#sec-layout-shift and layout_shift.idl.
type PerformanceTimelineLayoutShift struct {
	// Value Score increment produced by this event.
	Value float64 `json:"value"`

	// HadRecentInput ...
	HadRecentInput bool `json:"hadRecentInput"`

	// LastInputTime ...
	LastInputTime TimeSinceEpoch `json:"lastInputTime"`

	// Sources ...
	Sources []*PerformanceTimelineLayoutShiftAttribution `json:"sources"`
}

// PerformanceTimelineTimelineEvent ...
type PerformanceTimelineTimelineEvent struct {
	// FrameID Identifies the frame that this event is related to. Empty for non-frame targets.
	FrameID PageFrameID `json:"frameId"`

	// Type The event type, as specified in https://w3c.github.io/performance-timeline/#dom-performanceentry-entrytype
	// This determines which of the optional "details" fields is present.
	Type string `json:"type"`

	// Name may be empty depending on the type.
	Name string `json:"name"`

	// Time in seconds since Epoch, monotonically increasing within document lifetime.
	Time TimeSinceEpoch `json:"time"`

	// Duration (optional) Event duration, if applicable.
	Duration *float64 `json:"duration,omitempty"`

	// LcpDetails (optional) ...
	LcpDetails *PerformanceTimelineLargestContentfulPaint `json:"lcpDetails,omitempty"`

	// LayoutShiftDetails (optional) ...
	LayoutShiftDetails *PerformanceTimelineLayoutShift `json:"layoutShiftDetails,omitempty"`
}

// See also: timelineEventAdded.
type PerformanceTimelineEnable struct {
	// EventTypes The types of event to report, as specified in
	// https://w3c.github.io/performance-timeline/#dom-performanceentry-entrytype
	// The specified filter overrides any previous filters, passing empty
	// filter disables recording.
	// Note that not all types exposed to the web platform are currently supported.
	EventTypes []string `json:"eventTypes"`
}

// ProtoReq name.
func (m PerformanceTimelineEnable) ProtoReq() string { return "PerformanceTimeline.enable" }

// Call sends the request.
func (m PerformanceTimelineEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// PerformanceTimelineTimelineEventAdded Sent when a performance timeline event is added. See reportPerformanceTimeline method.
type PerformanceTimelineTimelineEventAdded struct {
	// Event ...
	Event *PerformanceTimelineTimelineEvent `json:"event"`
}

// ProtoEvent name.
func (evt PerformanceTimelineTimelineEventAdded) ProtoEvent() string {
	return "PerformanceTimeline.timelineEventAdded"
}
