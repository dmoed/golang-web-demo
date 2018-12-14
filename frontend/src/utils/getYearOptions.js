import moment from 'moment'

export default function getYearOptions(minYear = 2012){

    let currentYear = (new Date()).getFullYear();

    return Array.apply(null, new Array(parseInt(currentYear) - minYear + 1)).map((un, i) => ({
        label: "" + (currentYear - i),
        value: "" + (currentYear - i)
    }));
}