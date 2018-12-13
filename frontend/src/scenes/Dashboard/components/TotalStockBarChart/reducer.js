/**
 * Created by wardo on 3/19/2018.
 */
import constants from './constants'

const initialState = {
    data: [],
    isLoading: false
};

function reducer(state = initialState, action) {
    switch (action.type) {
        case constants.TOTAL_STOCK_BAR_CHART_START_FETCH:

            return Object.assign({}, state, {isLoading: true});

        case constants.TOTAL_STOCK_BAR_CHART_RECEIVED_DATA:

            const sData = action.payload;

            return Object.assign({}, state, {data: sData, isLoading: false});
    }

    return state;
}

export default reducer