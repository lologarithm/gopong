/** @jsx React.DOM */
define(['react', './validationInput.js'], function (React, ValidationInput) {
	var MemberComposer = React.createClass({
		propTypes: {
			header: React.PropTypes.string,
			buttonText: React.PropTypes.string, //The text on the button, defaults to 'Add'
			validationFunc: React.PropTypes.func,

			onMembersChange: React.PropTypes.func
		},

		getInitialState: function(){
            return {
                members: this.props.members || [],
                memberText: '',
                memberValid: false
            };
        },

        componentWillUpdate: function (next) {
			if(next.members) {
				this.state.members = next.members.slice();
			}
        },

		render: function () {
			var controls = [];

			if(this.props.header) {
				controls.push(<h4>{this.props.header}</h4>);
			}

			var mapFunction = this.props.memberTemplate ? this.props.memberTemplate.bind(this) : function (a,i) {
				return <div className="margin-bottom-m">
							<span>{a}</span>
							<button type="button" className="btn btn-danger btn-xs right" onClick={this.removeMember(i)}>X</button>
						</div>;
			}.bind(this); 
			controls.push(this.state.members.map(mapFunction));

			if(this.props.validationFunc) {
				var buttonText = this.props.buttonText || 'Add';
				
				controls.push(<div className="flexContainer margin-bottom-m">
					<ValidationInput className='flex1' style={{'margin-bottom':'0px'}} isValid={this.state.memberValid} value={this.state.memberText} onChange={this.onMemberTextChange} onKeyDown={this.onKeyDown} id="member" />
					<button type="button" className="btn btn-primary flexNone margin-left-l" onClick={this.addMember} >{buttonText}</button>
				</div>);
			} else {
				var buttonText = this.props.buttonText || 'Add';

				controls.push(<div className="flexContainer margin-bottom-m">
					<input className='flex1' style={{'margin-bottom':'0px'}} value={this.state.memberText} onChange={this.onMemberTextChange} />
					<button type="button" className="btn btn-primary flexNone margin-left-l" onClick={this.addMember} >{buttonText}</button>
				</div>);
			}

			return this.transferPropsTo(
							<div>
								{controls}
							</div>);
		},

		onMemberTextChange: function (e) {
			if (e.target.value !== this.state.memberText) {
				this.setState({ memberText : e.target.value });

				if(this.props.validationFunc) {
					this.setState({ memberValid : this.props.validationFunc(e.target.value) });
				}	
			}
		},

		addMember: function (e) {
			if(this.props.validationFunc && this.props.validationFunc(this.state.memberText)) {
				var members = this.state.members.concat(this.state.memberText);
				this.setState({ members : members, memberText: '' });

				this.props.onMembersChange(members);	
			}
			
			e.stopPropagation();
		},

		removeMember: function (index) {
			return function () {
				var members = this.state.members.filter(function (a,i) { return i !== index; });
				this.setState({ members: members });

				this.props.onMembersChange(members);
			}.bind(this);
		},

		onKeyDown: function (e) {
			if (e.keyCode === 13) {
				this.addMember(e);
			}
		}
	});

	return MemberComposer;
})