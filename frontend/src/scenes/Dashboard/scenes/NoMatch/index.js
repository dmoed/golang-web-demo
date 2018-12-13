import React from 'react';

import withPage from './../../components/Page/index'
// import img from './../../notfound.png'
import './style.scss'

class NoMatch extends React.Component {
    render() {
        return (
            <div>
                <div className="page-content not-found">
                    <h1>404</h1>
                    <h3>The page you're looking for was not found</h3>
                    <div className="preview" style={{backgroundImage: `url()`}}>
                    </div>
                </div>
            </div>
        )
    }
}

export default withPage(NoMatch, "404");

