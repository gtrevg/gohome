var React = require('react');

var ZonePicker = React.createClass({
    getInitialState: function() {
        return {
            value: ''
        };
    },

    selected: function(evt) {
        this.setState({ value: evt.target.value });
        this.props.changed && this.props.changed(evt.target.value);
    },
    
    render: function() {
        var options = [];
        this.props.zones.forEach(function(zone) {
            options.push(<option key={zone.id} value={zone.id}>{zone.name}</option>);
        });
        return (
            <div className="cmp-ZonePicker">
              <select className="form-control" onChange={this.selected} value={this.state.value}>
                <option value="">Select a zone...</option>
                {options}
              </select>
            </div>
        );
    }
});
module.exports = ZonePicker;