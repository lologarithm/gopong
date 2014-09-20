define(['./accountActionTypes.js', '/js/dispatcher.js'], function (actionTypes, dispatcher) {

	var accountActions = {
		updateAccount: function (data) {
			dispatcher.dispatch({ type: actionTypes.UPDATE_ACCOUNT, data: data });
		},
		changeRegion: function (data) {
			dispatcher.dispatch({ 
				type: actionTypes.CHANGE_REGION,
				data: data
			});	
		}
	};

	return accountActions;
});