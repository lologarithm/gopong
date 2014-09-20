define(['/js/settings.js', '/js/api/tournamentAPI.js', '/js/mockServer/mockMatchAPI.js', '/js/actions/tournamentActionTypes.js', '/js/dispatcher.js', '/js/utility/cookies.js'], 

function (settings, realAPI, mockAPI, TournamentActionTypes, dispatcher, cookies) {
	var changeListeners = [];

	var tournaments = { adminTourn: [], playerTourns: [] };
	var hasFetched = false;

	var api = settings.useLocal ? mockAPI : realAPI;

	var tournamentStore = {
		addChangeListener: function (callback) {
			changeListeners.push(callback);
		},

		removeChangeListener: function (callback) {
			changeListeners = changeListeners.filter(function (a) {
				return a !== callback;
			});
		},

		notifyListeners: function () {
			changeListeners.forEach(function (a) {
				a();
			});
		},

		getTournaments: function () {
			var hasSession = cookies.getCookie('sessionid');

			if(!hasFetched && hasSession) {
				hasFetched = true;
				api.getTournaments(function (data) {
					tournaments = data || tournaments;

					this.notifyListeners();
				}.bind(this));
			}

			return tournaments;
		},

		getTournament: function (id) {
			var hasSession = cookies.getCookie('sessionid');
			var loaded = tournaments.adminTourn
				.concat(tournaments.playerTourns)
				.filter(function (a) {
					return a.Id == id;
				});

			if (loaded.length == 0 || loaded[0].Admins == null && hasSession) {
				api.getTournament(id, function (data) {
					var index = tournaments.adminTourn.indexOf(loaded[0]);
					var playerIndex = tournaments.playerTourns.indexOf(loaded[0]);

					if(index === -1 && playerIndex !== -1) {
						tournaments.playerTourns[index] = data;
					} else if (index !== -1) {
						tournaments.adminTourn[index] = data;
					}

					this.notifyListeners();
				}.bind(this));

				return null;
			} else {
				return loaded[0];	
			}
		},

		createTournament: function (model) {
			if (!model.Id) {
				api.createTournament(model, function (data) {

					tournaments.adminTourn = tournaments.adminTourn.concat(data);

					this.notifyListeners();
				}.bind(this));
			} else {
				api.updateTournament(model, function (data) {
					console.log(data);

					var index = tournaments.adminTourn.findIndex(function (a){
						return a.Id == model.Id;
					});

					tournaments.adminTourn[index] = $.extend(true, tournaments.adminTourn[index], data);

					this.notifyListeners();
				}.bind(this));
			}
		},

		generateBrackets: function (id) {
			api.generateBrackets(id, function (data) {

				var index = tournaments.adminTourn.findIndex(function (a){
					return a.Id == id;
				});

				tournaments.adminTourn[index] = $.extend(true, tournaments.adminTourn[index], data);

				this.notifyListeners();
			}.bind(this));
		},

		advanceTournament: function (id) {
			api.advanceTournament(id, function (data) {

				var index = tournaments.adminTourn.findIndex(function (a){
					return a.Id == id;
				});

				tournaments.adminTourn[index] = $.extend(true, tournaments.adminTourn[index], data);

				this.notifyListeners();
			}.bind(this));
		},

		deleteTournament: function (id) {
			api.deleteTournament(id, function (data) {

				tournaments.adminTourn = tournaments.adminTourn.filter(function (a) {
					return a.Id !== id;
				});
			}.bind(this));
		}
	}

	//Register this store to react appropriately to UI events
	dispatcher.register(function (action) {
		switch(action.type) {
			case TournamentActionTypes.CREATE_TOURNAMENT:
				tournamentStore.createTournament(action.data);
				break;
			case TournamentActionTypes.GENERATE_BRACKETS:
				tournamentStore.generateBrackets(action.data);
				break;
			case TournamentActionTypes.ADVANCE_TOURNAMENT:
				tournamentStore.advanceTournament(action.data);
				break;
			case TournamentActionTypes.DELETE_TOURNAMENT:
				tournamentStore.deleteTournament(action.data);
			default:
				break;
		}
	});

	return tournamentStore;
})