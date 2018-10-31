package utils

import (
	"bytes"
	"io"
	"text/template"

	"github.com/muesli/digitalwhistleblowingfund/api/config"
	"github.com/muesli/digitalwhistleblowingfund/api/db"

	"github.com/muesli/gomail"
)

var (
	templates = make(map[string]config.EmailTemplate)
	settings  config.Data
)

// TemplateHelper combines multiple db-structs to make them become
// accessible from the template
type TemplateHelper struct {
	User     *db.User
	Proposal *db.Proposal
	BaseURL  string
}

// SetupEmailTemplates compiles the email templates
func SetupEmailTemplates(c config.Data) {
	settings = c
	templates["invitation"] = config.EmailTemplate{
		Subject: c.App.Templates.Invitation.Subject,
		Text:    c.App.Templates.Invitation.Text,
		HTML:    c.App.Templates.Invitation.HTML,
	}
	templates["moderation_proposal"] = config.EmailTemplate{
		Subject: c.App.Templates.ModerationProposal.Subject,
		Text:    c.App.Templates.ModerationProposal.Text,
		HTML:    c.App.Templates.ModerationProposal.HTML,
	}
	templates["proposal_accepted"] = config.EmailTemplate{
		Subject: c.App.Templates.ProposalAccepted.Subject,
		Text:    c.App.Templates.ProposalAccepted.Text,
		HTML:    c.App.Templates.ProposalAccepted.HTML,
	}
	templates["proposal_started"] = config.EmailTemplate{
		Subject: c.App.Templates.ProposalStarted.Subject,
		Text:    c.App.Templates.ProposalStarted.Text,
		HTML:    c.App.Templates.ProposalStarted.HTML,
	}
}

// SendInvitation sends out an email, inviting a user to join polly
func SendInvitation(user *db.User) {
	tmpl := templates["invitation"]

	th := TemplateHelper{
		User:    user,
		BaseURL: settings.Web.BaseURL,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", settings.Connections.Email.ReplyTo)
	m.SetHeader("To", user.Email)
	//	m.SetAddressHeader("Cc", "foo@foobar.com", "Joe")
	//	m.Attach("/tmp/attachment.jpg")

	var buf bytes.Buffer
	t := template.Must(template.New("invitation_subject").Parse(tmpl.Subject))
	t.Execute(&buf, th)
	m.SetHeader("Subject", buf.String())

	m.AddAlternativeWriter("text/plain", func(w io.Writer) error {
		t := template.Must(template.New("invitation_text").Parse(tmpl.Text))
		return t.Execute(w, th)
	})
	/*TODO: if HTML in template is empty, do not write html content

	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
	t := template.Must(template.New("invitation_html").Parse(tmpl.HTML))
		return t.Execute(w, th)
	})*/

	d := gomail.NewDialer(settings.Connections.Email.SMTP.Server, settings.Connections.Email.SMTP.Port,
		settings.Connections.Email.SMTP.User, settings.Connections.Email.SMTP.Password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

// SendModerationRequest sends out an email to the admin, asking for moderation of a newly posted proposal
func SendModerationRequest(proposal *db.Proposal) {
	tmpl := templates["moderation_proposal"]

	th := TemplateHelper{
		Proposal: proposal,
		BaseURL:  settings.Web.BaseURL,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", settings.Connections.Email.ReplyTo)
	m.SetHeader("To", settings.Connections.Email.AdminEmail)

	var buf bytes.Buffer
	t := template.Must(template.New("moderation_proposal_subject").Parse(tmpl.Subject))
	t.Execute(&buf, th)
	m.SetHeader("Subject", buf.String())

	m.AddAlternativeWriter("text/plain", func(w io.Writer) error {
		t := template.Must(template.New("moderation_proposal_text").Parse(tmpl.Text))
		return t.Execute(w, th)
	})
	/*TODO
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		t := template.Must(template.New("moderation_proposal_html").Parse(tmpl.HTML))
		return t.Execute(w, th)
	})*/

	d := gomail.NewDialer(settings.Connections.Email.SMTP.Server, settings.Connections.Email.SMTP.Port,
		settings.Connections.Email.SMTP.User, settings.Connections.Email.SMTP.Password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

// SendProposalAccepted sends out an email to the proposal author, when their proposal got accepted
func SendProposalAccepted(user *db.User, proposal *db.Proposal) {
	tmpl := templates["proposal_accepted"]

	th := TemplateHelper{
		User:     user,
		Proposal: proposal,
		BaseURL:  settings.Web.BaseURL,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", settings.Connections.Email.ReplyTo)
	m.SetHeader("To", user.Email)

	var buf bytes.Buffer
	t := template.Must(template.New("proposal_accepted_subject").Parse(tmpl.Subject))
	t.Execute(&buf, th)
	m.SetHeader("Subject", buf.String())

	m.AddAlternativeWriter("text/plain", func(w io.Writer) error {
		t := template.Must(template.New("proposal_accepted_text").Parse(tmpl.Text))
		return t.Execute(w, th)
	})
	/*TODO
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		t := template.Must(template.New("proposal_accepted_html").Parse(tmpl.HTML))
		return t.Execute(w, th)
	})*/

	d := gomail.NewDialer(settings.Connections.Email.SMTP.Server, settings.Connections.Email.SMTP.Port,
		settings.Connections.Email.SMTP.User, settings.Connections.Email.SMTP.Password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

// SendProposalStarted sends out an email to mailman, when a proposal got started
func SendProposalStarted(proposal db.Proposal) {
	tmpl := templates["proposal_started"]

	th := TemplateHelper{
		Proposal: &proposal,
		BaseURL:  settings.Web.BaseURL,
	}

	m := gomail.NewMessage()
	m.SetAddressHeader("From", settings.Connections.Email.Mailman.Address, settings.Connections.Email.Mailman.Name)
	m.SetAddressHeader("To", settings.Connections.Email.Mailman.Address, settings.Connections.Email.Mailman.Name)

	var buf bytes.Buffer
	t := template.Must(template.New("proposal_started_subject").Parse(tmpl.Subject))
	t.Execute(&buf, th)
	m.SetHeader("Subject", buf.String())

	m.AddAlternativeWriter("text/plain", func(w io.Writer) error {
		t := template.Must(template.New("proposal_started_text").Parse(tmpl.Text))
		return t.Execute(w, th)
	})
	/*TODO
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		t := template.Must(template.New("proposal_started_html").Parse(tmpl.HTML))
		return t.Execute(w, th)
	})*/

	d := gomail.NewDialer(settings.Connections.Email.SMTP.Server, settings.Connections.Email.SMTP.Port,
		settings.Connections.Email.SMTP.User, settings.Connections.Email.SMTP.Password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
