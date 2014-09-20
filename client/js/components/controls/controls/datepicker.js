/** @jsx React.DOM */
define(['react'], function (React) {
	var Datepicker = React.createClass({

		getInitialState: function(){
            return {
                startDate : '',
                endDate : ''
            };
        },

        componentDidMount: function () {
			// $('#sandbox-container .input-group.date').datepicker({
			//     format: "mm/dd/yyyy"
			// });
        },

        componentWillReceiveProps: function(nextProps) {
			this.setState( { isVisible : nextProps.isVisible });
        },

		render: function () {
			return this.transferPropsTo(
				<div id="datepicker">
					<label for="start">Tournament Dates</label>
					<div className="input-daterange input-group">
					    <input id="start" placeholder="mm/dd/yyyy" type="text" className="input-sm form-control" value={this.props.startDate || this.state.startDate} onChange={this.onStartChange} name="start" />
					    <span className="input-group-addon">to</span>
					    <input placeholder="mm/dd/yyyy" type="text" className="input-sm form-control" value={this.props.endDate || this.state.endDate} onChange={this.onEndChange} name="end" />
				    </div>
				</div>);
		},

		onStartChange: function (e) {
			this.setState ({ startDate : e.target.value });

			this.props.onStartChange(e.target.value);
		},

		onEndChange: function (e) {
			this.setState ({ endDate : e.target.value });

			this.props.onEndChange(e.target.value);
		},
	});

	return Datepicker;
})