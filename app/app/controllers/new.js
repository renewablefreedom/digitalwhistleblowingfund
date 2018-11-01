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
  whistleblowingtype: "",
  motivations: "",
  partners: "",
  board: "",
  communication: "",
  information: "",
  usage: "",
  dependency: "",
  sustain: "",
  value: "",
  recipient: "",
  contact: "",
  spamanswer: "",
  responseMessage: "",
  errorMessage: "",
  startdate: new Date(),

  isValid: Ember.computed('title', 'applicant', 'applicantdescription', 'referrerorganization', 'referrercontact', 'description', 'socialgoals',
    'activities', 'geofocus', 'laws', 'whistleblowingtype', 'motivations', 'partners', 'board', 'communication', 'information', 'usage',
    'dependency', 'sustain', 'value', 'recipient', 'contact', 'startdate', 'spamanswer',
    function () {
      return this.title.length > 0 && this.applicant.length > 0 &&
        this.applicantdescription.length > 0 && this.referrerorganization.length > 0 &&
        this.referrercontact.length > 0 && this.description.length > 0 &&
        this.socialgoals.length > 0 && this.activities.length > 0 &&
        this.geofocus.length > 0 && this.laws.length > 0 &&
        this.whistleblowingtype.length > 0 && this.motivations.length > 0 &&
        this.partners.length > 0 && this.board.length > 0 &&
        this.communication.length > 0 && this.information.length > 0 &&
        this.usage.length > 0 && this.dependency.length > 0 &&
        this.sustain.length > 0 && this.recipient.length > 0 &&
        this.contact.length > 0 && this.spamanswer == '42' &&
        parseInt(this.value) > 0 &&
        this.startdate.getFullYear() > 0;
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
      const whistleblowingtype = this.get('whistleblowingtype');
      const motivations = this.get('motivations');
      const partners = this.get('partners');
      const board = this.get('board');
      const communication = this.get('communication');
      const information = this.get('information');
      const usage = this.get('usage');
      const dependency = this.get('dependency');
      const sustain = this.get('sustain');
      const value = this.get('value');
      const recipient = this.get('recipient');
      const contact = this.get('contact');
      const startdate = this.get('startdate');
      const spamanswer = this.get('spamanswer');

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
          this.set('whistleblowingtype', '');
          this.set('motivations', '');
          this.set('partners', '');
          this.set('board', '');
          this.set('communication', '');
          this.set('information', '');
          this.set('usage', '');
          this.set('dependency', '');
          this.set('sustain', '');
          this.set('value', '');
          this.set('recipient', '');
          this.set('contact', '');
          this.set('spamanswer', '');
          this.set('startdate', new Date());
        },
        error => {
          this.set('errorMessage', `Failed adding your proposal: ` + error);
          this.set('progressMessage', '');
        }
      );
    }
  }
});
