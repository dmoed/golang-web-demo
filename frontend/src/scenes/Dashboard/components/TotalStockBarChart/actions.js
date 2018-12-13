/**
 * Created by wardo on 3/19/2018.
 */
import constants from './constants'
import axios from 'axios'

export function startFetch(){
    return {type: constants.TOTAL_STOCK_BAR_CHART_START_FETCH}
}

export function receivedData(data){
    return {type: constants.TOTAL_STOCK_BAR_CHART_RECEIVED_DATA, payload: data}
}

export function fetchData(url){
    return (dispatch, getState) => {

        dispatch(startFetch());

        //get state
        const newState = getState();

        axios.get(url)
            .then(res => {

                dispatch(receivedData(res.data.payload))

            })
            .catch(err => {

                console.log(err);

            });
    }
}
