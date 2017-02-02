import Ember from 'ember';

const { inject: { service } } = Ember;

export default Ember.Controller.extend({
    session: Ember.inject.service('session'),
    currentUser: service('current-user'),

    actions: {
      invalidateSession() {
        this.store.unloadAll('vote');
        this.store.unloadAll('proposal');
        this.get('session').invalidate();
      }
    },

    selfupdate: Ember.inject.service(),
    willInitSelfUpdate: function() {
      this.get('selfupdate').watchUpdates();
    }.on('init')
});
