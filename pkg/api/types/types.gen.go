// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package types

// Defines values for ErrorCode.
const (
	ErrorCodeHOSTNOTFOUND ErrorCode = "HOST_NOT_FOUND"

	ErrorCodeINTERNAL ErrorCode = "INTERNAL"

	ErrorCodeMALFORMEDPEERID ErrorCode = "MALFORMED_PEER_ID"

	ErrorCodeMALFORMEDREQUEST ErrorCode = "MALFORMED_REQUEST"

	ErrorCodePEERNOTFOUND ErrorCode = "PEER_NOT_FOUND"

	ErrorCodeROUTINGTABLENOTFOUND ErrorCode = "ROUTING_TABLE_NOT_FOUND"

	ErrorCodeSAVINGROUTINGTABLE ErrorCode = "SAVING_ROUTING_TABLE"
)

// Defines values for RoutingTableUpdateType.
const (
	RoutingTableUpdateTypeFULL RoutingTableUpdateType = "FULL"

	RoutingTableUpdateTypePEERADDED RoutingTableUpdateType = "PEER_ADDED"

	RoutingTableUpdateTypePEERREMOVED RoutingTableUpdateType = "PEER_REMOVED"
)

// AddProvider defines model for AddProvider.
type AddProvider struct {
	Distance    string  `json:"distance"`
	DurationInS float32 `json:"durationInS"`
	EndedAt     string  `json:"endedAt"`
	Error       *string `json:"error"`
	Id          int     `json:"id"`
	RemoteId    string  `json:"remoteId"`
	StartedAt   string  `json:"startedAt"`
}

// Can be any value - string, number, boolean, array or object.
type AnyValue interface{}

// Connection defines model for Connection.
type Connection struct {
	DurationInS  float32 `json:"durationInS"`
	EndedAt      string  `json:"endedAt"`
	Id           int     `json:"id"`
	MultiAddress string  `json:"multiAddress"`
	RemoteId     string  `json:"remoteId"`
	StartedAt    string  `json:"startedAt"`
}

// CreateHostRequest defines model for CreateHostRequest.
type CreateHostRequest struct {
	// An arbitrary name for this host.
	Name string `json:"name"`
}

// Dial defines model for Dial.
type Dial struct {
	DurationInS  float32 `json:"durationInS"`
	EndedAt      string  `json:"endedAt"`
	Error        *string `json:"error"`
	Id           int     `json:"id"`
	MultiAddress string  `json:"multiAddress"`
	RemoteId     string  `json:"remoteId"`
	StartedAt    string  `json:"startedAt"`
	Transport    string  `json:"transport"`
}

// ErrorCode defines model for ErrorCode.
type ErrorCode string

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Code ErrorCode `json:"code"`

	// Can be any value - string, number, boolean, array or object.
	Details *AnyValue `json:"details"`
	Message string    `json:"message"`
}

// FindNode defines model for FindNode.
type FindNode struct {
	CloserPeersCount *int    `json:"closerPeersCount"`
	DurationInS      float32 `json:"durationInS"`
	EndedAt          string  `json:"endedAt"`
	Error            *string `json:"error"`
	Id               int     `json:"id"`
	RemoteId         string  `json:"remoteId"`
	StartedAt        string  `json:"startedAt"`
}

// Host defines model for Host.
type Host struct {
	BootstrappedAt       *string `json:"bootstrappedAt"`
	CreatedAt            string  `json:"createdAt"`
	HostId               string  `json:"hostId"`
	Name                 string  `json:"name"`
	RunningProvidesCount float32 `json:"runningProvidesCount"`
}

// Peer defines model for Peer.
type Peer struct {
	AgentVersion *string  `json:"agentVersion"`
	CreatedAt    string   `json:"createdAt"`
	PeerId       string   `json:"peerId"`
	Protocols    []string `json:"protocols"`
}

// Provide defines model for Provide.
type Provide struct {
	ContentId             string  `json:"contentId"`
	EndedAt               *string `json:"endedAt"`
	Error                 *string `json:"error"`
	FinalRoutingTableId   *int    `json:"finalRoutingTableId"`
	HostId                string  `json:"hostId"`
	InitialRoutingTableId int     `json:"initialRoutingTableId"`
	ProvideId             int     `json:"provideId"`
	StartedAt             string  `json:"startedAt"`
}

// ProvideDetails defines model for ProvideDetails.
type ProvideDetails struct {
	// Embedded struct due to allOf(#/components/schemas/Provide)
	Provide `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	AddProviders []AddProvider `json:"addProviders"`
	Connections  []Connection  `json:"connections"`
	Dials        []Dial        `json:"dials"`
	FindNodes    []FindNode    `json:"findNodes"`
}

// RoutingTable defines model for RoutingTable.
type RoutingTable struct {
	BucketSize int    `json:"bucketSize"`
	CreatedAt  string `json:"createdAt"`
	EntryCount int    `json:"entryCount"`
	HostId     string `json:"hostId"`
	Id         int    `json:"id"`
}

// RoutingTableDetails defines model for RoutingTableDetails.
type RoutingTableDetails struct {
	// Embedded struct due to allOf(#/components/schemas/RoutingTable)
	RoutingTable `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Entries []RoutingTableEntry `json:"entries"`
}

// RoutingTableEntry defines model for RoutingTableEntry.
type RoutingTableEntry struct {
	AddedAt                       string  `json:"addedAt"`
	Bucket                        int     `json:"bucket"`
	ConnectedSince                *string `json:"connectedSince"`
	LastSuccessfulOutboundQueryAt string  `json:"lastSuccessfulOutboundQueryAt"`
	LastUsefulAt                  *string `json:"lastUsefulAt"`
	PeerId                        string  `json:"peerId"`
}

// RoutingTablePeer defines model for RoutingTablePeer.
type RoutingTablePeer struct {
	AddedAt                       string   `json:"addedAt"`
	AgentVersion                  *string  `json:"agentVersion"`
	Bucket                        int      `json:"bucket"`
	ConnectedSince                *string  `json:"connectedSince"`
	LastSuccessfulOutboundQueryAt string   `json:"lastSuccessfulOutboundQueryAt"`
	LastUsefulAt                  *string  `json:"lastUsefulAt"`
	PeerId                        string   `json:"peerId"`
	Protocols                     []string `json:"protocols"`
}

// RoutingTablePeers defines model for RoutingTablePeers.
type RoutingTablePeers []RoutingTablePeer

// RoutingTableUpdate defines model for RoutingTableUpdate.
type RoutingTableUpdate struct {
	Type   RoutingTableUpdateType `json:"type"`
	Update interface{}            `json:"update"`
}

// RoutingTableUpdateType defines model for RoutingTableUpdate.Type.
type RoutingTableUpdateType string

// CreateHostJSONBody defines parameters for CreateHost.
type CreateHostJSONBody CreateHostRequest

// CreateHostJSONRequestBody defines body for CreateHost for application/json ContentType.
type CreateHostJSONRequestBody CreateHostJSONBody

