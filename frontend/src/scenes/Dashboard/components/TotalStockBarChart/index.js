/**
 * Created by wardo on 3/19/2018.
 */
import React from 'react';
import ReactDOM from 'react-dom';
import {connect} from 'react-redux'
import {fetchData} from './actions'
import Loader from './../../components/Loader'
import Url from 'domurl'

import BarChart from 'recharts/lib/chart/BarChart';
import Bar from 'recharts/lib/cartesian/Bar';
import XAxis from 'recharts/lib/cartesian/XAxis';
import YAxis from 'recharts/lib/cartesian/YAxis';
import Tooltip from 'recharts/lib/component/Tooltip';
import ReferenceLine from 'recharts/lib/cartesian/ReferenceLine';
import ResponsiveContainer from 'recharts/lib/component/ResponsiveContainer';
const TAG = "[ STOCK BAR WIDGET ]";
import moment from 'moment';
import domtoimage from 'dom-to-image';
import {saveAs} from 'file-saver';

function displayTotal(n, d) {
    let x = ('' + n).length;
    let p = Math.pow;
    d = p(10, d);

    x -= x % 3;
    return Math.round(n * d / p(10, x)) / d + " kMGTPE"[x / 3];
}

const CustomizedLabel = ({x, y, value, width}) => {
    return <text x={x+width/2} y={y} dy={-4}
                 fontSize='11' fontFamily='sans-serif' fontWeight="light"
                 fill={'blue'} textAnchor="middle">
        {displayTotal(value, 1).toLocaleString()}
    </text>
};

const StockTooltip = ({active, payload, unit}) => {
    if (active) {
        return (
            <div className="custom-tooltip"
                 style={{backgroundColor: '#fff', border: '1px solid rgb(204, 204, 204)', padding: '12px 8px', borderRadius: '4px'}}>
                <p><strong>{payload[0]['payload']['dates'][0]} to {payload[0]['payload']['dates'][1]}</strong></p>
                <p><strong>{payload[0]['payload']['year_week']}</strong></p>
                <p>{Number(payload[0]['value']).toLocaleString()} {unit}</p>
            </div>
        );
    }

    return null;
};

const CustomerTooltip = ({active, payload}) => {
    if (active) {
        return (
            <div className="custom-tooltip"
                 style={{backgroundColor: '#fff', border: '1px solid rgb(204, 204, 204)', padding: '12px 8px', borderRadius: '4px'}}>
                <p><strong>{payload[0]['payload']['dates'][0]} to {payload[0]['payload']['dates'][1]}</strong></p>
                <p>{payload[0]['value']} klanten</p>
            </div>
        );
    }

    return null;
};

const renderCustomizedXTick = ({ x = 0, y = 0, payload }) => {
    return (
        <text x={x} y={y + 5} textAnchor="middle" dominantBaseline="hanging"
              fontSize='11' fontFamily='sans-serif' fontWeight="bold">
            {payload.value}
        </text>
    );
};

class TotalStockBarChart extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            exportIsLoading: false
        };
    }

    componentDidMount() {
        let {week, year, fetchData, period} = this.props;

        setTimeout(() => fetchData(week, year, period), 100);
    }

    componentDidUpdate(prevProps) {

        let {period, week, year, fetchData} = this.props;

        console.log(TAG, 'updated');
        let dirty = false;

        if (prevProps.week !== week) {
            dirty = true;
        }
        if (prevProps.year !== year) {
            dirty = true;
        }
        if (prevProps.period !== period) {
            dirty = true;
        }

        if (dirty) {
            setTimeout(() => fetchData(week, year, period), 100);
        }
    }

    render() {

        const {isLoading, data, period, unit, avgQty, avgVisits} = this.props;
        let {exportIsLoading} = this.state;

        return (
            <div>
                <div className="panel panel-default panel-widget">
                    <div className="panel-body">

                        <div className="row">
                            <div className="col-md-6">
                                <h4>Stock Meting</h4>
                            </div>
                            <div className="col-md-6">
                                { data.length > 0 ? <p style={{textAlign: 'right'}}>
                                    <button className="btn btn-xs" disabled={exportIsLoading}
                                            onClick={() => this.exportChart('stock')}>{exportIsLoading ? 'saving as PNG' : 'Download'}
                                    </button>
                                </p> : null}
                            </div>
                        </div>

                        {isLoading
                            ?
                            <Loader />
                            :
                            <div className="text-center">
                                {this.renderStockChart(data, avgQty, unit, period)}
                                <span>
                                    Gemiddeld voorraad <em
                                    style={{color: 'red'}}>{avgQty.toLocaleString()} {unit}</em> per week.
                                </span>
                            </div>}
                    </div>
                </div>
            </div>
        );
    }

    renderStockChart(data, avgQty, unit = 'trays', period) {

        if (data.length < 1) {
            return null;
        }

        let dataKey = 'total_' + unit;

        return (
                <BarChart height={240} width={period * 100} data={data} margin={{top: 15, right: 0, left: 0, bottom: 5}}
                          ref={(chart) => this.stockBarChart = chart}>
                    <XAxis dataKey="label" axisLine={false} tickSize={0} tick={renderCustomizedXTick}/>
                    <Tooltip content={<StockTooltip unit={unit}/>}/>
                    <ReferenceLine y={avgQty} label="" stroke="red" strokeDasharray="3 3"/>
                    <Bar dataKey={dataKey} fill='#f26522' label={<CustomizedLabel />} isAnimationActive={false}/>
                </BarChart>
        );
    }

    renderCustomerChart(data, avgVisits) {

        if (data.length < 1) {
            return null;
        }

        return (
            <ResponsiveContainer width='100%' minHeight={180} style={{margin: '0 auto'}}>
                <BarChart height={100} data={data} margin={{top: 15, right: 0, left: 0, bottom: 5}}
                          ref={(chart) => this.customersBarChart = chart}>
                    <XAxis dataKey="label" axisLine={false} tickSize={0} tick={renderCustomizedXTick}/>
                    <Tooltip content={<CustomerTooltip/>}/>
                    <ReferenceLine y={avgVisits} label="" stroke="red" alwaysShow={true}
                                   strokeDasharray="3 3"/>
                    <Bar dataKey={'total_customers'} fill='#47ACB1' label={<CustomizedLabel />}
                         isAnimationActive={false}/>
                </BarChart>
            </ResponsiveContainer>
        );
    }

    exportChart(type) {

        this.setState({exportIsLoading: true});

        let {week, year} = this.props;

        let date_1 = moment().day("Monday").year(year).week(week);
        let date_2 = moment().day("Monday").year(year).week(week).add(6, 'days');
        let chartNode;

        switch (type) {
            case 'customers':
                chartNode = ReactDOM.findDOMNode(this.customersBarChart);
                break;
            default:
                chartNode = ReactDOM.findDOMNode(this.stockBarChart);
        }

        domtoimage.toBlob(chartNode)
            .then((blob) => {
                this.setState({exportIsLoading: false});
                saveAs(blob, `total_stock_meting_${date_1.format('D-MMMM')}_${date_2.format('D-MMMM')}_${year}.png`);
            })
            .catch((error) => {
                console.error('oops, something went wrong!', error);
                this.setState({exportIsLoading: false});
            });
    }
}

function mapStateToProps({totalStockBarChart, filterWidget}, ownProps) {

    let data = Object.values(totalStockBarChart.data);
    let unit = filterWidget.unit;
    let avgQty = 0;
    let avgVisits = 0;

    data.sort((a,b) => (parseInt(a.year + "" + a.week) < parseInt(b.year + "" + b.week)) ? 1 : (parseInt(a.year + "" + a.week) > parseInt(b.year + "" + b.week) ? -1 : 0)); 

     if (data.length) {

        let e = 0;
        data.forEach((q) => {
            if (parseInt(q['total_' + unit]) === 0) {
                return;
            }
            if (parseInt(q.week) === parseInt(ownProps.currentWeek)) {
                return;
            }
            e++;
            avgQty += q['total_' + unit];
            //avgVisits += q['total_customers'];
        });
        avgQty = Math.round(avgQty / e);
        //avgVisits = Math.round(avgVisits / e);
     }

    return {
        isLoading:false,
        data: data,
        unit: unit,
        period: filterWidget.period,
        week: filterWidget.week,
        year: filterWidget.year,
        avgQty: avgQty,
        avgVisits: avgVisits
    }
}

function mapDispatchToProps(dispatch, ownProps) {
    return {
        fetchData: (week, year, period) => {
            let url = new Url(ownProps.url);
            url.query.week = week;
            url.query.year = year;
            url.query.max = period;
            dispatch(fetchData(url.toString()))
        }
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(TotalStockBarChart)