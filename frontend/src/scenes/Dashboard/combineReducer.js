import {combineReducers} from 'redux'

import dashboardReducer from './reducer'
import userReducer from './userReducer'
import totalStockBarChartReducer from './components/TotalStockBarChart/reducer'

export default combineReducers({
    dashboard: dashboardReducer,
    user: userReducer,
    totalStockBarChart: totalStockBarChartReducer
})