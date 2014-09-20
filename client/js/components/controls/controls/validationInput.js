
/** @jsx React.DOM */
define(['react'], function (React) {
	var Modal = React.createClass({

		render: function () {
			var label;
			if(this.props.label) {
				label = <label className="control-label" for={this.props.id}>{this.props.label}</label>;	
			}
				

			return this.transferPropsTo(
				<div className={this.generateFormgroupClass(this.props.isValid, this.props.label)}>
					{label}
					<input type="text" className="form-control" value={this.props.value} onChange={this.props.onChange} id={this.props.id} />
					<span style={{'top': (this.props.label ? '' : '0px') }} className={this.generateIconClass(this.props.isValid)}></span>
				</div>);
		},

		generateIconClass: function (isAuth) {
			return 'glyphicon form-control-feedback ' + (isAuth ? 'glyphicon-ok' : 'glyphicon-remove');
		},

		generateFormgroupClass: function (isAuth) {
			return 'form-group has-feedback ' + (isAuth ? 'has-success ' : 'has-error ');
		},
	});

	return Modal;
})