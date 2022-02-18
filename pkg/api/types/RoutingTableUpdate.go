// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    routingTableUpdate, err := UnmarshalRoutingTableUpdate(bytes)
//    bytes, err = routingTableUpdate.Marshal()

package types

import "encoding/json"

type RoutingTableUpdate []RoutingTablePeer

func UnmarshalRoutingTableUpdate(data []byte) (RoutingTableUpdate, error) {
	var r RoutingTableUpdate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RoutingTableUpdate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RoutingTablePeer struct {
	AddedAt                       string  `json:"added_at"`                         
	AgentVersion                  *string `json:"agent_version"`                    
	Bucket                        int64   `json:"bucket"`                           
	ConnectedAt                   *string `json:"connected_at"`                     
	LastSuccessfulOutboundQueryAt string  `json:"last_successful_outbound_query_at"`
	LastUsefulAt                  *string `json:"last_useful_at"`                   
	PeerID                        string  `json:"peer_id"`                          
}
