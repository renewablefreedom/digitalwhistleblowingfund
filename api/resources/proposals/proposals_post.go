package proposals

import (
	"errors"
	"net/http"
	"time"

	"github.com/muesli/digitalwhistleblowingfund/api/db"
	"github.com/muesli/digitalwhistleblowingfund/api/utils"

	"github.com/emicklei/go-restful"
	"github.com/muesli/smolder"
)

// ProposalPostStruct holds all values of an incoming POST request
type ProposalPostStruct struct {
	Proposal struct {
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
		Contact              string    `json:"contact"`
		Recipient            string    `json:"recipient"`
		Value                uint64    `json:"value"`
		RealValue            uint64    `json:"realvalue"`
		Moderated            bool      `json:"moderated"`
		Starts               time.Time `json:"starts"`
		Finished             uint64    `json:"finished"`
		SpamAnswer           string    `json:"spamanswer"`
	} `json:"proposal"`
}

// PostAuthRequired returns true because all requests need authentication
func (r *ProposalResource) PostAuthRequired() bool {
	return false
}

// PostDoc returns the description of this API endpoint
func (r *ProposalResource) PostDoc() string {
	return "create a new proposal"
}

// PostParams returns the parameters supported by this API endpoint
func (r *ProposalResource) PostParams() []*restful.Parameter {
	return nil
}

// Post processes an incoming POST (create) request
func (r *ProposalResource) Post(context smolder.APIContext, data interface{}, request *restful.Request, response *restful.Response) {
	ctx := context.(*db.PollyContext)
	resp := ProposalResponse{}
	resp.Init(context)

	pps := data.(*ProposalPostStruct)

	if pps.Proposal.SpamAnswer != ctx.Config.App.SpamAnswer {
		smolder.ErrorResponseHandler(request, response, smolder.NewErrorResponse(
			http.StatusUnauthorized,
			true,
			errors.New("Please enter the correct value in the spam protection field"),
			"ProposalResource POST"))
		return
	}

	proposal := db.Proposal{
		UserID:               1,
		Title:                pps.Proposal.Title,
		Applicant:            pps.Proposal.Applicant,
		Applicantdescription: pps.Proposal.Applicantdescription,
		Referrerorganization: pps.Proposal.Referrerorganization,
		Referrercontact:      pps.Proposal.Referrercontact,
		Description:          pps.Proposal.Description,
		Socialgoals:          pps.Proposal.Socialgoals,
		Activities:           pps.Proposal.Activities,
		Geofocus:             pps.Proposal.Geofocus,
		Laws:                 pps.Proposal.Laws,
		Whistleblowingtype:   pps.Proposal.Whistleblowingtype,
		Motivations:          pps.Proposal.Motivations,
		Partners:             pps.Proposal.Partners,
		Board:                pps.Proposal.Board,
		Communication:        pps.Proposal.Communication,
		Information:          pps.Proposal.Information,
		Usage:                pps.Proposal.Usage,
		Dependency:           pps.Proposal.Dependency,
		Sustain:              pps.Proposal.Sustain,
		Contact:              pps.Proposal.Contact,
		Recipient:            pps.Proposal.Recipient,
		Value:                pps.Proposal.Value,
		Starts:               pps.Proposal.Starts,
	}
	err := proposal.Save(context.(*db.PollyContext))
	if err != nil {
		smolder.ErrorResponseHandler(request, response, smolder.NewErrorResponse(
			http.StatusInternalServerError,
			true,
			err,
			"ProposalResource POST"))
		return
	}

	utils.SendModerationRequest(&proposal)

	resp.AddProposal(&proposal)
	resp.Send(response)
}
