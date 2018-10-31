import Ember from 'ember';

const {
  inject: {
    service
  }
} = Ember;

export default Ember.Route.extend({
  currentUser: service('current-user'),

  setupController(controller, models) {
    //       controller.set('contact', this.get('currentUser').get('user').get('email'));
    controller.set('errorMessage', '');
    controller.set('responseMessage', '');
    controller.set('progressMessage', '');
  }
});
