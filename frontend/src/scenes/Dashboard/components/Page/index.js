import React from 'react'
import {connect} from 'react-redux'
import {toggleSidebar} from './../../actions'

export default function withPage(BasicComponent, pageTitle = "", pageTitle2 = "") {

    class Page extends React.Component {

        componentDidMount() {
            this.props.dispatch(toggleSidebar(false)); //close sidebar
            window.scrollTo(0, 0)
        }

        componentWillMount() {

            let {config} = this.props;

            console.log(config)

            document.title = this.props.pageTitle + "";
        }

        componentDidUpdate(prevProps) {
        }

        render() {
            return <BasicComponent {...this.props}/>
        }
    }

    Page.defaultProps = {pageTitle};

    return connect(null)(Page);
}

