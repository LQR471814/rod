// This file is generated by "./lib/proto/generate"

package proto

/*

Memory

*/

// MemoryPressureLevel Memory pressure level.
type MemoryPressureLevel string

const (
	// MemoryPressureLevelModerate enum const.
	MemoryPressureLevelModerate MemoryPressureLevel = "moderate"

	// MemoryPressureLevelCritical enum const.
	MemoryPressureLevelCritical MemoryPressureLevel = "critical"
)

// MemorySamplingProfileNode Heap profile sample.
type MemorySamplingProfileNode struct {
	// Size of the sampled allocation.
	Size float64 `json:"size"`

	// Total bytes attributed to this sample.
	Total float64 `json:"total"`

	// Stack Execution stack at the point of allocation.
	Stack []string `json:"stack"`
}

// MemorySamplingProfile Array of heap profile samples.
type MemorySamplingProfile struct {
	// Samples ...
	Samples []*MemorySamplingProfileNode `json:"samples"`

	// Modules ...
	Modules []*MemoryModule `json:"modules"`
}

// MemoryModule Executable module information.
type MemoryModule struct {
	// Name of the module.
	Name string `json:"name"`

	// UUID of the module.
	UUID string `json:"uuid"`

	// BaseAddress Base address where the module is loaded into memory. Encoded as a decimal
	// or hexadecimal (0x prefixed) string.
	BaseAddress string `json:"baseAddress"`

	// Size of the module in bytes.
	Size float64 `json:"size"`
}

// MemoryGetDOMCounters ...
type MemoryGetDOMCounters struct{}

// ProtoReq name.
func (m MemoryGetDOMCounters) ProtoReq() string { return "Memory.getDOMCounters" }

// Call the request.
func (m MemoryGetDOMCounters) Call(c Client) (*MemoryGetDOMCountersResult, error) {
	var res MemoryGetDOMCountersResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// MemoryGetDOMCountersResult ...
type MemoryGetDOMCountersResult struct {
	// Documents ...
	Documents int `json:"documents"`

	// Nodes ...
	Nodes int `json:"nodes"`

	// JsEventListeners ...
	JsEventListeners int `json:"jsEventListeners"`
}

// MemoryPrepareForLeakDetection ...
type MemoryPrepareForLeakDetection struct{}

// ProtoReq name.
func (m MemoryPrepareForLeakDetection) ProtoReq() string { return "Memory.prepareForLeakDetection" }

// Call sends the request.
func (m MemoryPrepareForLeakDetection) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// MemoryForciblyPurgeJavaScriptMemory Simulate OomIntervention by purging V8 memory.
type MemoryForciblyPurgeJavaScriptMemory struct{}

// ProtoReq name.
func (m MemoryForciblyPurgeJavaScriptMemory) ProtoReq() string {
	return "Memory.forciblyPurgeJavaScriptMemory"
}

// Call sends the request.
func (m MemoryForciblyPurgeJavaScriptMemory) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// MemorySetPressureNotificationsSuppressed Enable/disable suppressing memory pressure notifications in all processes.
type MemorySetPressureNotificationsSuppressed struct {
	// Suppressed If true, memory pressure notifications will be suppressed.
	Suppressed bool `json:"suppressed"`
}

// ProtoReq name.
func (m MemorySetPressureNotificationsSuppressed) ProtoReq() string {
	return "Memory.setPressureNotificationsSuppressed"
}

// Call sends the request.
func (m MemorySetPressureNotificationsSuppressed) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// MemorySimulatePressureNotification Simulate a memory pressure notification in all processes.
type MemorySimulatePressureNotification struct {
	// Level Memory pressure level of the notification.
	Level MemoryPressureLevel `json:"level"`
}

// ProtoReq name.
func (m MemorySimulatePressureNotification) ProtoReq() string {
	return "Memory.simulatePressureNotification"
}

// Call sends the request.
func (m MemorySimulatePressureNotification) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// MemoryStartSampling Start collecting native memory profile.
type MemoryStartSampling struct {
	// SamplingInterval (optional) Average number of bytes between samples.
	SamplingInterval *int `json:"samplingInterval,omitempty"`

	// SuppressRandomness (optional) Do not randomize intervals between samples.
	SuppressRandomness *bool `json:"suppressRandomness,omitempty"`
}

// ProtoReq name.
func (m MemoryStartSampling) ProtoReq() string { return "Memory.startSampling" }

// Call sends the request.
func (m MemoryStartSampling) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// MemoryStopSampling Stop collecting native memory profile.
type MemoryStopSampling struct{}

// ProtoReq name.
func (m MemoryStopSampling) ProtoReq() string { return "Memory.stopSampling" }

// Call sends the request.
func (m MemoryStopSampling) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// MemoryGetAllTimeSamplingProfile Retrieve native memory allocations profile
// collected since renderer process startup.
type MemoryGetAllTimeSamplingProfile struct{}

// ProtoReq name.
func (m MemoryGetAllTimeSamplingProfile) ProtoReq() string { return "Memory.getAllTimeSamplingProfile" }

// Call the request.
func (m MemoryGetAllTimeSamplingProfile) Call(c Client) (*MemoryGetAllTimeSamplingProfileResult, error) {
	var res MemoryGetAllTimeSamplingProfileResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// MemoryGetAllTimeSamplingProfileResult ...
type MemoryGetAllTimeSamplingProfileResult struct {
	// Profile ...
	Profile *MemorySamplingProfile `json:"profile"`
}

// MemoryGetBrowserSamplingProfile Retrieve native memory allocations profile
// collected since browser process startup.
type MemoryGetBrowserSamplingProfile struct{}

// ProtoReq name.
func (m MemoryGetBrowserSamplingProfile) ProtoReq() string { return "Memory.getBrowserSamplingProfile" }

// Call the request.
func (m MemoryGetBrowserSamplingProfile) Call(c Client) (*MemoryGetBrowserSamplingProfileResult, error) {
	var res MemoryGetBrowserSamplingProfileResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// MemoryGetBrowserSamplingProfileResult ...
type MemoryGetBrowserSamplingProfileResult struct {
	// Profile ...
	Profile *MemorySamplingProfile `json:"profile"`
}

// MemoryGetSamplingProfile Retrieve native memory allocations profile collected since last
// `startSampling` call.
type MemoryGetSamplingProfile struct{}

// ProtoReq name.
func (m MemoryGetSamplingProfile) ProtoReq() string { return "Memory.getSamplingProfile" }

// Call the request.
func (m MemoryGetSamplingProfile) Call(c Client) (*MemoryGetSamplingProfileResult, error) {
	var res MemoryGetSamplingProfileResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// MemoryGetSamplingProfileResult ...
type MemoryGetSamplingProfileResult struct {
	// Profile ...
	Profile *MemorySamplingProfile `json:"profile"`
}
