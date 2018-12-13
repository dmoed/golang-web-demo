import React from 'react';

import withPage from './../../components/Page/index'
// import img from './../../notfound.png'
import './style.scss'

class ForbiddenPage extends React.Component {
    render() {
        return (
            <div>
                <div className="page-content forbidden">
                    <h1>403</h1>
                    <h3>You are not authorized to access this page.</h3>
                    <div className="preview" style={{backgroundImage: `url()`}}>
                    </div>
                </div>
            </div>
        )
    }
}

export default withPage(ForbiddenPage, '403');

