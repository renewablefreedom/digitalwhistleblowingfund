package proposals

import (
	"net/http"
	"strconv"
	"time"

	"github.com/muesli/digitalwhistleblowingfund/api/db"
	"github.com/muesli/digitalwhistleblowingfund/api/utils"

	"github.com/emicklei/go-restful"
	"github.com/muesli/smolder"
)

// PutAuthRequired returns true because all requests need authentication
func (r *ProposalResource) PutAuthRequired() bool {
	return true
}

// PutDoc returns the description of this API endpoint
func (r *ProposalResource) PutDoc() string {
	return "update an existing proposal"
}

// PutParams returns the parameters supported by this API endpoint
func (r *ProposalResource) PutParams() []*restful.Parameter {
	return nil
}

// Put processes an incoming PUT (update) request
func (r *ProposalResource) Put(context smolder.APIContext, data interface{}, request *restful.Request, response *restful.Response) {
	resp := ProposalResponse{}
	resp.Init(context)

	pps := data.(*ProposalPostStruct)

	id, err := strconv.Atoi(request.PathParameter("proposal-id"))
	if err != nil {
		smolder.ErrorResponseHandler(request, response, smolder.NewErrorResponse(
			http.StatusBadRequest,
			false,
			"Invalid proposal id",
			"ProposalResource PUT"))
		return
	}

	proposal, err := context.(*db.PollyContext).GetProposalByID(int64(id))
	if err != nil {
		r.NotFound(request, response)
		return
	}

	auth, err := context.Authentication(request)
	if err != nil || (auth.(db.User).ID != 1 && auth.(db.User).ID != proposal.UserID) {
		smolder.ErrorResponseHandler(request, response, smolder.NewErrorResponse(
			http.StatusUnauthorized,
			false,
			"Admin permission required for this operation",
			"ProposalResource PUT"))
		return
	}

	if auth.(db.User).ID != 1 && proposal.Starts.Before(time.Now()) {
		smolder.ErrorResponseHandler(request, response, smolder.NewErrorResponse(
			http.StatusUnauthorized,
			false,
			"Can't update a proposal once it has started",
			"ProposalResource PUT"))
	}

	proposal.Title = pps.Proposal.Title
	proposal.Applicant = pps.Proposal.Applicant
	proposal.Applicantdescription = pps.Proposal.Applicantdescription
	proposal.Referrerorganization = pps.Proposal.Referrerorganization
	proposal.Referrercontact = pps.Proposal.Referrercontact
	proposal.Description = pps.Proposal.Description
	proposal.Socialgoals = pps.Proposal.Socialgoals
	proposal.Activities = pps.Proposal.Activities
	proposal.Geofocus = pps.Proposal.Geofocus
	proposal.Laws = pps.Proposal.Laws
	proposal.Whistleblowingtype = pps.Proposal.Whistleblowingtype
	proposal.Motivations = pps.Proposal.Motivations
	proposal.Partners = pps.Proposal.Partners
	proposal.Board = pps.Proposal.Board
	proposal.Communication = pps.Proposal.Communication
	proposal.Information = pps.Proposal.Information
	proposal.Usage = pps.Proposal.Usage
	proposal.Dependency = pps.Proposal.Dependency
	proposal.Sustain = pps.Proposal.Sustain
	proposal.Contact = pps.Proposal.Contact
	proposal.Recipient = pps.Proposal.Recipient
	proposal.Value = pps.Proposal.Value
	proposal.RealValue = pps.Proposal.RealValue
	// proposal.Starts = pps.Proposal.Starts

	if auth.(db.User).ID == 1 {
		if !proposal.Moderated && pps.Proposal.Moderated {
			proposalAuthor, uerr := context.(*db.PollyContext).GetUserByID(proposal.UserID)
			if uerr != nil {
				panic(uerr)
			}
			utils.SendProposalAccepted(&proposalAuthor, &proposal)
		}

		proposal.Moderated = pps.Proposal.Moderated

		if pps.Proposal.Finished > 0 {
			proposal.Finished = pps.Proposal.Finished
		}
	}

	err = proposal.Update(context.(*db.PollyContext))
	if err != nil {
		smolder.ErrorResponseHandler(request, response, smolder.NewErrorResponse(
			http.StatusInternalServerError,
			true,
			"Can't update proposal",
			"ProposalResource PUT"))
		return
	}

	resp.AddProposal(&proposal)
	resp.Send(response)
}
