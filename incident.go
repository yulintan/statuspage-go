package statuspage

import (
	"context"
)

// IncidentService handles communication with the page related methods
// of the Statuspage API.
//
// Statuspage API docs: https://developer.statuspage.io/#tag/incidents
type IncidentService service

// Incident is the Statuspage API incident representation
type Incident struct {
	ID                            *string     `json:"id,omitempty"`
	Components                    []Component `json:"components,omitempty"`
	CreatedAt                     *Timestamp  `json:"created_at,omitempty"`
	Impact                        *string     `json:"impact,omitempty"`
	ImpactOverride                *string     `json:"impact_override,omitempty"`
	MonitoringAt                  *Timestamp  `json:"monitoring_at,omitempty"`
	Name                          *string     `json:"name,omitempty"`
	PageID                        *string     `json:"page_id,omitempty"`
	PortmortemBody                *string     `json:"portmortem_body,omitempty"`
	PostmortemBodyLastUpdatedAt   *Timestamp  `json:"postmortem_body_last_updated_at,omitempty"`
	PostmortemIgnored             *bool       `json:"postmortem_ignored,omitempty"`
	PostmortemNotifiedSubscribers *bool       `json:"postmortem_notified_subscribers,omitempty"`
	PostmortemNotifiedTwitter     *bool       `json:"postmortem_notified_twitter,omitempty"`
	PostmortemPublishedAt         *string     `json:"postmortem_published_at,omitempty"`
	ResolvedAt                    *Timestamp  `json:"resolved_at,omitempty"`
	ScheduledAutoCompleted        *bool       `json:"scheduled_auto_completed,omitempty"`
	ScheduledAutoInProgress       *bool       `json:"scheduled_auto_in_progress,omitempty"`
	ScheduledFor                  *Timestamp  `json:"scheduled_for,omitempty"`
	ScheduledRemindPrior          *bool       `json:"scheduled_remind_prior,omitempty"`
	ScheduledRemindedAt           *Timestamp  `json:"scheduled_reminded_at,omitempty"`
	ScheduledUntil                *Timestamp  `json:"scheduled_until,omitempty"`
	Shortlink                     *string     `json:"shortlink,omitempty"`
	Status                        *string     `json:"status,omitempty"`
	UpdatedAt                     *Timestamp  `json:"updated_at,omitempty"`
}

// UpdateIncidentParams are the parameters that can be changed using the update incident API endpoint
type UpdateIncidentParams struct {
	Name         string            `json:"name,omitempty"`
	Status       string            `json:"status,omitempty"`
	Body         string            `json:"body,omitempty"`
	Components   map[string]string `json:"components,omitempty"`
	ComponentIDs []string          `json:"component_i_ds,omitempty"`
}

// UpdateIncidentRequestBody is the update incident request body representation
type UpdateIncidentRequestBody struct {
	Incident UpdateIncidentParams `json:"incident"`
}

func (i Incident) String() string {
	return Stringify(i)
}

// ListIncidents returns incident information for a given page id
func (s *IncidentService) ListIncidents(ctx context.Context, pageID string) ([]Incident, error) {
	path := "v1/pages/" + pageID + "/incidents"
	req, err := s.client.newRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	var incident []Incident
	_, err = s.client.do(ctx, req, &incident)

	return incident, err
}

// ListIncidents returns incident information for a given page id
func (s *IncidentService) CreateIncident(ctx context.Context, pageID string, incident UpdateIncidentParams) (*Incident, error) {
	path := "v1/pages/" + pageID + "/incidents"
	payload := UpdateIncidentRequestBody{Incident: incident}
	req, err := s.client.newRequest("POST", path, payload)
	if err != nil {
		return nil, err
	}

	var updatedIncident Incident
	_, err = s.client.do(ctx, req, &updatedIncident)

	return &updatedIncident, err
}
