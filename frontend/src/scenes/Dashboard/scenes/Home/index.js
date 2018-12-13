import React from 'react'
import moment from 'moment'

import withPage from './../../components/Page/index'
import TotalStockBarChart from './../../components/TotalStockBarChart'
const currentWeek = moment().format('w');

class Home extends React.Component {

    render(){

        const {config: {routes: {url_ajax_total_stock_bar_chart}}} = this.props;

        return (
            <div className="page-content-empty">
                <h1>Report</h1>

                <div>
                    <TotalStockBarChart url={url_ajax_total_stock_bar_chart} currentWeek={currentWeek}/>
                </div>
            </div>
        )
    }
}

export default withPage(Home, "Home");
