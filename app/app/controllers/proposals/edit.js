import Ember from 'ember';
import moment from 'moment';

const {
  inject: {
    service
  }
} = Ember;

export default Ember.Controller.extend({
  currentUser: service('current-user'),

  responseMessage: "",
  errorMessage: "",

  title: Ember.computed('proposal', function () {
    return this.get('proposal').get('title');
  }),
  applicant: Ember.computed('proposal', function () {
    return this.get('proposal').get('applicant');
  }),
  applicantdescription: Ember.computed('proposal', function () {
    return this.get('proposal').get('applicantdescription');
  }),
  referrerorganization: Ember.computed('proposal', function () {
    return this.get('proposal').get('referrerorganization');
  }),
  referrercontact: Ember.computed('proposal', function () {
    return this.get('proposal').get('referrercontact');
  }),
  description: Ember.computed('proposal', function () {
    return this.get('proposal').get('description');
  }),
  socialgoals: Ember.computed('proposal', function () {
    return this.get('proposal').get('socialgoals');
  }),
  activities: Ember.computed('proposal', function () {
    return this.get('proposal').get('activities');
  }),
  geofocus: Ember.computed('proposal', function () {
    return this.get('proposal').get('geofocus');
  }),
  laws: Ember.computed('proposal', function () {
    return this.get('proposal').get('laws');
  }),
  whistleblowingtype: Ember.computed('proposal', function () {
    return this.get('proposal').get('whistleblowingtype');
  }),
  motivations: Ember.computed('proposal', function () {
    return this.get('proposal').get('motivations');
  }),
  partners: Ember.computed('proposal', function () {
    return this.get('proposal').get('partners');
  }),
  board: Ember.computed('proposal', function () {
    return this.get('proposal').get('board');
  }),
  communication: Ember.computed('proposal', function () {
    return this.get('proposal').get('communication');
  }),
  information: Ember.computed('proposal', function () {
    return this.get('proposal').get('information');
  }),
  usage: Ember.computed('proposal', function () {
    return this.get('proposal').get('usage');
  }),
  dependency: Ember.computed('proposal', function () {
    return this.get('proposal').get('dependency');
  }),
  sustain: Ember.computed('proposal', function () {
    return this.get('proposal').get('sustain');
  }),
  contact: Ember.computed('proposal', function () {
    return this.get('proposal').get('contact');
  }),
  recipient: Ember.computed('proposal', function () {
    return this.get('proposal').get('recipient');
  }),
  value: Ember.computed('proposal', function () {
    return this.get('proposal').get('value');
  }),
  realvalue: Ember.computed('proposal', function () {
    return this.get('proposal').get('realvalue');
  }),
  startdate: Ember.computed('proposal', function () {
    return this.get('proposal').get('starts');
  }),

  isValid: Ember.computed('title', 'applicant', 'applicantdescription', 'referrerorganization', 'referrercontact', 'description', 'socialgoals',
    'activities', 'geofocus', 'laws', 'whistleblowingtype', 'motivations', 'partners', 'board', 'communication', 'information', 'usage',
    'dependency', 'sustain', 'value', 'recipient', 'contact', 'startdate',
    function () {

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
      const contact = this.get('contact');
      const recipient = this.get('recipient');
      const value = this.get('value');
      const realvalue = this.get('realvalue');
      const startdate = this.get('startdate');

      return title.length > 0 && applicant.length > 0 &&
        applicantdescription.length > 0 && referrerorganization.length > 0 &&
        referrercontact.length > 0 && description.length > 0 &&
        socialgoals.length > 0 && activities.length > 0 &&
        geofocus.length > 0 && laws.length > 0 &&
        whistleblowingtype.length > 0 && motivations.length > 0 &&
        partners.length > 0 && board.length > 0 &&
        communication.length > 0 && information.length > 0 &&
        usage.length > 0 && dependency.length > 0 &&
        sustain.length > 0 && recipient.length > 0 &&
        contact.length > 0 &&
        parseInt(value) > 0 &&
        startdate.getFullYear() > 0;
    }),
  isDisabled: Ember.computed.not('isValid'),

  actions: {
    saveProposal() {
      this.set('errorMessage', '');
      this.set('responseMessage', '');
      this.set('progressMessage', `Saving proposal...`);

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
      const realvalue = this.get('realvalue');
      const recipient = this.get('recipient');
      const contact = this.get('contact');
      const startdate = this.get('startdate');

      var proposal = this.get('proposal');

      proposal.set('title', title);
      proposal.set('applicant', applicant);
      proposal.set('applicantdescription', applicantdescription);
      proposal.set('referrerorganization', referrerorganization);
      proposal.set('referrercontact', referrercontact);
      proposal.set('description', description);
      proposal.set('socialgoals', socialgoals);
      proposal.set('activities', activities);
      proposal.set('geofocus', geofocus);
      proposal.set('laws', laws);
      proposal.set('whistleblowingtype', whistleblowingtype);
      proposal.set('motivations', motivations);
      proposal.set('partners', partners);
      proposal.set('board', board);
      proposal.set('communication', communication);
      proposal.set('information', information);
      proposal.set('usage', usage);
      proposal.set('dependency', dependency);
      proposal.set('sustain', sustain);
      proposal.set('contact', contact);
      proposal.set('recipient', recipient);
      proposal.set('value', value);
      proposal.set('realvalue', realvalue);
      proposal.set('starts', startdate);

      proposal.save().then(
        ( /*proposal*/ ) => {
          this.set('responseMessage', `Your proposal has been updated. Thank you!`);
          this.set('progressMessage', '');
          /*          this.set('title', '');
                    this.set('description', '');
                    this.set('recipient', '');
                    this.set('value', '');
                    this.set('startdate', new Date());*/
        },
        error => {
          this.set('errorMessage', `Failed updating your proposal: ` + error);
          this.set('progressMessage', '');
        }
      );
    }
  }
});
