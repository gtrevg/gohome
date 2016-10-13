var Constants = require('../constants.js');
var initialState = require('../initialState.js');

module.exports = function(state, action) {
    var newState = state;

    switch(action.type) {
    case Constants.ZONE_LOAD_ALL:
        break;

    case Constants.ZONE_LOAD_ALL_FAIL:
        break;

    case Constants.ZONE_LOAD_ALL_RAW:
        newState = action.data;
        break;

    case Constants.ZONE_CREATE:
        break;
    case Constants.ZONE_CREATE_RAW:
        break;
    case Constants.ZONE_CREATE_FAIL:
        break;

    case Constants.ZONE_IMPORT:
        break;
    case Constants.ZONE_IMPORT_RAW:
        newState = [action.data].concat(newState);
        break;
    case Constants.ZONE_IMPORT_FAIL:
        break;
        
    default:
        newState = state || initialState().zones;
    }

    return newState;
};
