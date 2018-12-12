import {combineReducers} from 'redux'

import dashboardReducer from './reducer'
import userReducer from './userReducer'

export default combineReducers({
    dashboard: dashboardReducer,
    user: userReducer,
})