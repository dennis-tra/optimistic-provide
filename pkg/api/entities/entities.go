package entities

import "time"

type HostResponse struct {
	HostID         string     `json:"host_id"`
	BootstrappedAt *time.Time `json:"bootstrapped_at"`
	CreatedAt      time.Time  `json:"created_at"`
}

type ProvideCreateRequest struct {
	HostID string `json:"host_id"`
}

type ProvideResponse struct {
	ProvideID             int        `json:"provide_id"`
	HostID                string     `json:"host_id"`
	ContentID             string     `json:"content_id"`
	InitialRoutingTableID int        `json:"initial_routing_table_id"`
	FinalRoutingTableID   int        `json:"final_routing_table_id,omitempty"`
	StartedAt             time.Time  `json:"started_at"`
	EndedAt               *time.Time `json:"ended_at"`
}
