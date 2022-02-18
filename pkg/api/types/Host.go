// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    host, err := UnmarshalHost(bytes)
//    bytes, err = host.Marshal()

package types

import "encoding/json"

func UnmarshalHost(data []byte) (Host, error) {
	var r Host
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Host) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Host struct {
	BootstrappedAt *string `json:"bootstrapped_at"`
	CreatedAt      string  `json:"created_at"`     
	HostID         string  `json:"host_id"`        
}
