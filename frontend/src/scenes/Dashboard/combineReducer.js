import {combineReducers} from 'redux'

import dashboardReducer from './reducer'
import userReducer from './userReducer'
import filterWidgetReducer from './components/Filter/reducer'
import totalStockBarChartReducer from './components/TotalStockBarChart/reducer'

export default combineReducers({
    dashboard: dashboardReducer,
    user: userReducer,
    filterWidget: filterWidgetReducer,
    totalStockBarChart: totalStockBarChartReducer,
})