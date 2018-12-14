/**
 * Created by wardo on 3/16/2018.
 */

import constants from './constants'

export function setUnit(unit) {
    return {type: constants.WIDGET_FILTER_SET_UNIT, unit: unit}
}
export function setWeek(week) {
    return {type: constants.WIDGET_FILTER_SET_WEEK, week: week}
}
export function setYear(year) {
    return {type: constants.WIDGET_FILTER_SET_YEAR, year: year}
}
export function setPeriod(period) {
    return {type: constants.WIDGET_FILTER_SET_PERIOD, period: period}
}
