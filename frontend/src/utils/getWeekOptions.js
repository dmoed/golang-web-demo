import moment from 'moment'

function maxWeekInYear(year) {
    return Math.max(
        moment(new Date(year, 11, 31)).isoWeek(),
        moment(new Date(year, 11, 31 - 7)).isoWeek()
    );
}

export default function getWeekOptions(year) {

    let array = [];
    let maxWeek = maxWeekInYear(year);
    let i = 0;

    while (i < maxWeek) {

        let week = maxWeek - i;
        let start_date = moment().year(year).isoWeek(maxWeek - i).isoWeekday(1).format("DD-MMM");
        let end_date = moment().year(year).isoWeek(maxWeek - i).isoWeekday(7).format("DD-MMM");

        array.push({
            label: `(WK${week}) ${start_date} - ${end_date}`,
            value: "" + week,
            start_date: start_date,
            end_date: end_date
        });
        i++;
    }

    return array;
}
