/**
 * Created by wardo on 3/16/2018.
 */
import React from 'react'
import {connect} from 'react-redux'
import {setUnit, setWeek, setYear, setPeriod} from './actions'

import ToggleButton from 'react-toggle-button';
import Select from 'react-select';
import getYearOptions from './../../../../utils/getYearOptions'
import getWeekOptions from './../../../../utils/getWeekOptions'

const yearOptions = getYearOptions();

const periodOptions = [
    {label: "last 4 weeks", value: "4"},
    {label: "last 8 weeks", value: "8"},
    {label: "last 6 months", value: "24"},
    {label: "last 9 months", value: "36"},
    {label: "last 12 month", value: "52"},
]

const findValue = (value, options) =>  {
    return typeof value === 'string' && value.length
                ? options.find(option => option.value === value)
                : value;
}

const Filter = ({unit, week, year, period, setUnit, setWeek, setYear, setPeriod}) => {

    var weekOptions = [];
    
    if (year) {
        weekOptions = getWeekOptions(year);
    }
    
    return (
        <div className="row" style={{marginBottom: '20px'}}>
            <div className="col-md-2 col-sm-3">
                <label>Year</label>
                <Select
                    className="select-red"
                    name="inventory-year"
                    value={findValue(year, yearOptions)} clearable={false}
                    options={yearOptions}
                    onChange={(e) => setYear(e)}
                />
            </div>

            <div className="col-md-4 col-sm-6">
                <label>Week</label>
                <Select
                    className="select-red"
                    name="inventory-week"
                    value={findValue(week, weekOptions)} clearable={false}
                    options={weekOptions}
                    onChange={(e) => setWeek(e)}
                />
            </div>

            <div className="col-md-4 col-sm-6">
                <label>Period</label>
                <Select
                    className="select-red"
                    name="inventory-period"
                    value={findValue(period, periodOptions)} clearable={false}
                    options={periodOptions}
                    onChange={(e) => setPeriod(e)}
                />
            </div>

            <div className="col-md-2 col-sm-2">
                <label>Unit: {unit}</label>
                <ToggleButton
                    inactiveLabel={"trays"}
                    activeLabel={"bottles"}
                    value={unit !== 'trays'}
                    onToggle={(e) => setUnit(e)}/>
            </div>
        </div>
    );
};

function mapStateToProps(store) {
    return {
        unit: store.filterWidget.unit,
        week: store.filterWidget.week,
        year: store.filterWidget.year,
        period: store.filterWidget.period,
    }
}

function mapDispatchToProps(dispatch) {
    return {
        setWeek: (value) => {
            //console.log(value)
            dispatch(setWeek(value.value))
        },
        setYear: (value) => {
            //console.log(value)
            dispatch(setYear(value.value))
        },
        setPeriod: (value) => {
            //console.log(value)
            dispatch(setPeriod(value.value))
        },
        setUnit: (value) => {
            //console.log(value)
            let unit = !value ? 'bottles' : 'trays';
            dispatch(setUnit(unit))
        }
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Filter)