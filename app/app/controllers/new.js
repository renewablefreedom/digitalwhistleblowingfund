import Ember from 'ember';
import moment from 'moment';

export default Ember.Controller.extend({
  title: "",
  applicant: "",
  applicantdescription: "",
  referrerorganization: "",
  referrercontact: "",
  description: "",
  socialgoals: "",
  activities: "",
  geofocus: "",
  laws: "",
  communication: "",
  information: "",
  usage: "",
  sustain: "",
  recipient: "",
  contact: "",
  spamanswer: "",
  responseMessage: "",
  errorMessage: "",

  isValid: Ember.computed('title', 'applicant', 'applicantdescription', 'referrerorganization', 'referrercontact', 'description', 'socialgoals',
    'activities', 'geofocus', 'laws', 'communication', 'information', 'usage',
    'sustain', 'recipient', 'contact', 'spamanswer',
    function () {
      return this.title.length > 0 && this.applicant.length > 0 &&
        this.applicantdescription.length > 0 && this.referrerorganization.length > 0 &&
        this.referrercontact.length > 0 && this.description.length > 0 &&
        this.socialgoals.length > 0 && this.activities.length > 0 &&
        this.geofocus.length > 0 && this.laws.length > 0 &&
        this.communication.length > 0 && this.information.length > 0 &&
        this.usage.length > 0 &&
        this.sustain.length > 0 && this.recipient.length > 0 &&
        this.contact.length > 0 && this.spamanswer == '42';
    }),
  isDisabled: Ember.computed.not('isValid'),

  actions: {
    createProposal() {
      this.set('errorMessage', '');
      this.set('responseMessage', '');
      this.set('progressMessage', `Creating proposal...`);

      const title = this.get('title');
      const applicant = this.get('applicant');
      const applicantdescription = this.get('applicantdescription');
      const referrerorganization = this.get('referrerorganization');
      const referrercontact = this.get('referrercontact');
      const description = this.get('description');
      const socialgoals = this.get('socialgoals');
      const activities = this.get('activities');
      const geofocus = this.get('geofocus');
      const laws = this.get('laws');
      const communication = this.get('communication');
      const information = this.get('information');
      const usage = this.get('usage');
      const sustain = this.get('sustain');
      const recipient = this.get('recipient');
      const contact = this.get('contact');
      const spamanswer = this.get('spamanswer');
      const startdate = new Date();
      const board = false;
      const dependency = false;
      const value = 3000;
      const whistleblowingtype = '';
      const motivations = '';
      const partners = '';

      const newProposal = this.store.createRecord('proposal', {
        title: title,
        applicant: applicant,
        applicantdescription: applicantdescription,
        referrerorganization: referrerorganization,
        referrercontact: referrercontact,
        description: description,
        socialgoals: socialgoals,
        activities: activities,
        geofocus: geofocus,
        laws: laws,
        whistleblowingtype: whistleblowingtype,
        motivations: motivations,
        partners: partners,
        board: board,
        communication: communication,
        information: information,
        usage: usage,
        dependency: dependency,
        sustain: sustain,
        value: value,
        recipient: recipient,
        contact: contact,
        starts: startdate,
        spamanswer: spamanswer
      });
      newProposal.save().then(
        ( /*proposal*/ ) => {
          this.set('responseMessage', `Your proposal is now awaiting moderation. Thank you!`);
          this.set('progressMessage', '');
          this.set('title', '');
          this.set('applicant', '');
          this.set('applicantdescription', '');
          this.set('referrerorganization', '');
          this.set('referrercontact', '');
          this.set('description', '');
          this.set('socialgoals', '');
          this.set('activities', '');
          this.set('geofocus', '');
          this.set('laws', '');
          this.set('communication', '');
          this.set('information', '');
          this.set('usage', '');
          this.set('sustain', '');
          this.set('recipient', '');
          this.set('contact', '');
          this.set('spamanswer', '');
        },
        error => {
          this.set('errorMessage', `Failed adding your proposal: ` + error);
          this.set('progressMessage', '');
        }
      );
    }
  }
});
