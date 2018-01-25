package nimbusec

import (
	"context"
	"encoding/json"
	"net/http"
)

type IssueService service

func (srv *IssueService) List(ctx context.Context) ([]Issue, error) {
	issues := []issueDTO{}
	err := srv.client.Do(ctx, http.MethodGet, "/v3/issues", nil, &issues)
	if err != nil {
		return nil, err
	}

	converted := make([]Issue, len(issues))
	for i, issue := range issues {
		converted[i], err = translateIssue(issue)
		if err != nil {
			return nil, err
		}
	}

	return converted, nil
}

func (srv *IssueService) ListByDomain(ctx context.Context, id DomainID) ([]Issue, error) {
	issues := []issueDTO{}
	err := srv.client.Do(ctx, http.MethodGet, string(id)+"/issues", nil, &issues)
	if err != nil {
		return nil, err
	}

	converted := make([]Issue, len(issues))
	for i, issue := range issues {
		converted[i], err = translateIssue(issue)
		if err != nil {
			return nil, err
		}
	}

	return converted, nil
}

func (srv *IssueService) Get(ctx context.Context, id IssueID) (Issue, error) {
	issue := issueDTO{}
	err := srv.client.Do(ctx, http.MethodGet, string(id), nil, &issue)
	if err != nil {
		return Issue{}, err
	}

	return translateIssue(issue)
}

func (srv *IssueService) Update(ctx context.Context, id IssueID, update IssueUpdate) (Issue, error) {
	issue := issueDTO{}
	err := srv.client.Do(ctx, http.MethodPut, string(id), update, &issue)
	if err != nil {
		return Issue{}, err
	}

	return translateIssue(issue)
}

// issueDTO is a temporay struct to read json issues until they are fully parsed
// and translated into the exported Issue type.
type issueDTO struct {
	Issue
	Details json.RawMessage `json:"details,omitempty"`
}

func translateIssue(in issueDTO) (Issue, error) {
	var err error
	var out = in.Issue

	switch in.Event {
	case "blacklist":
		details := BlacklistDetails{}
		err = json.Unmarshal(in.Details, &details)
		out.Details = details

	case "cms-version":
		details := ApplicationOutdatedDetails{}
		err = json.Unmarshal(in.Details, &details)
		out.Details = details

	case "cms-vulnerable":
		details := ApplicationVulnerableDetails{}
		err = json.Unmarshal(in.Details, &details)
		out.Details = details

	case "defacement":
		details := DefacementDetails{}
		err = json.Unmarshal(in.Details, &details)
		out.Details = details

	case "malware":
		details := MalwareDetails{}
		err = json.Unmarshal(in.Details, &details)
		out.Details = details

	case "tls-ciphersuite":
		fallthrough
	case "tls-protocol":
		details := TLSConfigurationDetails{}
		err = json.Unmarshal(in.Details, &details)
		out.Details = details

	case "tls-expires":
		fallthrough
	case "tls-hostname":
		fallthrough
	case "tls-notrust":
		details := TLSCertificateDetails{}
		err = json.Unmarshal(in.Details, &details)
		out.Details = details

	case "webshell":
		details := WebshellDetails{}
		err = json.Unmarshal(in.Details, &details)
		out.Details = details
	}

	return out, err
}
