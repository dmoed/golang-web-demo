/**
 * Created by wardo on 3/16/2018.
 */

import constants from './constants'
import moment from 'moment'

const initialState = {
    unit: "trays",
    week: moment().format("W") + "",
    year: moment().format('Y') + "",
    period: "4"
};

function reducer(state = initialState, action) {
    switch (action.type) {
        case constants.WIDGET_FILTER_SET_UNIT:

            return Object.assign({}, state, {unit: action.unit});

        case constants.WIDGET_FILTER_SET_WEEK:

            return Object.assign({}, state, {week: action.week});

        case constants.WIDGET_FILTER_SET_YEAR:

            return Object.assign({}, state, {year: action.year});

        case constants.WIDGET_FILTER_SET_PERIOD:

        return Object.assign({}, state, {period: action.period});
    }

    return state;
}


export default reducer