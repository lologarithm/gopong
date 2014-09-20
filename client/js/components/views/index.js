/** @jsx React.DOM */
define([
'jquery',
'react',
'history',
'/js/components/views/summoner/master.js',
'/js/components/views/accountView.js',
'/js/components/views/aboutView.js',
'/js/components/views/tournament/tournamentView.js',
'/js/components/controls/controls/tabManager/verticalTabManager.js',
'/js/components/controls/loginHeader.js'], 

function ($, React, _h, SummonerView, AccountView, AboutView, TournamentView, TabManager, LoginHeader) {
	var IndexView = React.createClass({
		ACTIVE_VIEW: {
			ROOT: '/'
		},

		getInitialState : function () {
			return {
			};
		},

		componentDidMount: function () {
			this.navigateURL();
			window.addEventListener('popstate', function(event) {
				this.navigateURL ();
			}.bind(this));

			History.replaceState({}, 'Go Pong!', window.location.pathname);
		},

		navigateURL: function () {
			var tokens = window.location.pathname.split('/');

			switch(tokens[1]) {
				default:
					this.setState({ activeView : this.ACTIVE_VIEW.ROOT });
					break;
			}
		},

		render: function () {

			return <div className="flexContainer flexColumn">
						GO PONG!
					</div>
		},

		changeView: function (index) {
			return $.proxy(function () {
				History.pushState({}, 'Go Pong!', index);	
				this.setState({ activeView : index });

				if(index === this.ACTIVE_VIEW.ROOT) {
					this.setState(this.getInitialState());
				}
			}, this);
		}
	});

	function init() {
		React.renderComponent(<IndexView />, $('#container')[0]);
	}


	return init;
})