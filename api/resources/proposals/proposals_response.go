package proposals

import (
	"time"

	"github.com/muesli/digitalwhistleblowingfund/api/db"
	"github.com/muesli/digitalwhistleblowingfund/api/utils"

	"github.com/muesli/smolder"
)

// ProposalResponse is the common response to 'proposal' requests
type ProposalResponse struct {
	smolder.Response

	Proposals []proposalInfoResponse `json:"proposals,omitempty"`
	proposals []db.Proposal
}

type proposalInfoResponse struct {
	ID                   int64     `json:"id"`
	Title                string    `json:"title"`
	Applicant            string    `json:"applicant"`
	Applicantdescription string    `json:"applicantdescription"`
	Referrerorganization string    `json:"referrerorganization"`
	Referrercontact      string    `json:"referrercontact"`
	Description          string    `json:"description"`
	Socialgoals          string    `json:"socialgoals"`
	Activities           string    `json:"activities"`
	Geofocus             string    `json:"geofocus"`
	Laws                 string    `json:"laws"`
	Whistleblowingtype   string    `json:"whistleblowingtype"`
	Motivations          string    `json:"motivations"`
	Partners             string    `json:"partners"`
	Board                string    `json:"board"`
	Communication        string    `json:"communication"`
	Information          string    `json:"information"`
	Usage                string    `json:"usage"`
	Dependency           string    `json:"dependency"`
	Sustain              string    `json:"sustain"`
	User                 int64     `json:"user"`
	Contact              string    `json:"contact"`
	Recipient            string    `json:"recipient"`
	Value                uint64    `json:"value"`
	RealValue            uint64    `json:"realvalue"`
	GrantType            string    `json:"granttype"`
	URL                  string    `json:"url"`
	Starts               time.Time `json:"starts"`
	Ends                 time.Time `json:"ends"`
	Ended                bool      `json:"ended"`
	Accepted             bool      `json:"accepted"`
	Moderated            bool      `json:"moderated"`
	Votes                uint64    `json:"votes"`
	Vetos                uint64    `json:"vetos"`
}

// Init a new response
func (r *ProposalResponse) Init(context smolder.APIContext) {
	r.Parent = r
	r.Context = context

	r.Proposals = []proposalInfoResponse{}
}

// AddProposal adds a proposal to the response
func (r *ProposalResponse) AddProposal(proposal *db.Proposal) {
	r.proposals = append(r.proposals, *proposal)
	r.Proposals = append(r.Proposals, prepareProposalResponse(r.Context, proposal))
}

// EmptyResponse returns an empty API response for this endpoint if there's no data to respond with
func (r *ProposalResponse) EmptyResponse() interface{} {
	if len(r.proposals) == 0 {
		var out struct {
			Proposals interface{} `json:"proposals"`
		}
		out.Proposals = []proposalInfoResponse{}
		return out
	}
	return nil
}

func prepareProposalResponse(context smolder.APIContext, proposal *db.Proposal) proposalInfoResponse {
	ctx := context.(*db.PollyContext)
	resp := proposalInfoResponse{
		ID:                   proposal.ID,
		Title:                proposal.Title,
		Applicant:            proposal.Applicant,
		Applicantdescription: proposal.Applicantdescription,
		Referrerorganization: proposal.Referrerorganization,
		Referrercontact:      proposal.Referrercontact,
		Description:          proposal.Description,
		Socialgoals:          proposal.Socialgoals,
		Activities:           proposal.Activities,
		Geofocus:             proposal.Geofocus,
		Laws:                 proposal.Laws,
		Whistleblowingtype:   proposal.Whistleblowingtype,
		Motivations:          proposal.Motivations,
		Partners:             proposal.Partners,
		Board:                proposal.Board,
		Communication:        proposal.Communication,
		Information:          proposal.Information,
		Usage:                proposal.Usage,
		Dependency:           proposal.Dependency,
		Sustain:              proposal.Sustain,
		User:                 proposal.UserID,
		Contact:              proposal.Contact,
		Recipient:            proposal.Recipient,
		Value:                proposal.Value,
		RealValue:            proposal.RealValue,
		Starts:               proposal.Starts,
		Ends:                 proposal.Ends(ctx),
		Ended:                proposal.Ended(ctx),
		Votes:                proposal.Votes,
		Vetos:                proposal.Vetos,
		Accepted:             proposal.Accepted(ctx),
		Moderated:            proposal.Moderated,
		URL:                  utils.BuildURL(ctx.Config.Web.BaseURL, *proposal),
	}

	if proposal.Value < uint64(ctx.Config.App.Proposals.SmallGrantValueThreshold) {
		resp.GrantType = "small"
	} else {
		resp.GrantType = "large"
	}

	return resp
}
