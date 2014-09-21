/** @jsx React.DOM */
define([
'jquery',
'react',
'history'], 

function ($, React, _h) {
	var IndexView = React.createClass({
		WS: new WebSocket("ws://localhost:8080/ws"),
		
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
			// TODO: make this a real object.
			WS.send("{\"\": \"\"}");
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