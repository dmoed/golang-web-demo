import React from "react"
import {connect} from 'react-redux'
import {withRouter} from 'react-router-dom'

import Routes from './Routes'
import MyNavbar from './components/NavBar/index'
import ProfilePopupContainer from './components/ProfilePopup/container.js'

import 'bootstrap/dist/css/bootstrap.css'
import './style.scss'

class Dashboard extends React.Component {
  render(){

    const {roles, ...rest} = this.props;

    return (
      <div id="dashboard">

        <ProfilePopupContainer {...rest}/>

        <header>
          <MyNavbar {...rest}/>
        </header>
        <div className="page-wrapper">
          <div className="page-content-wrapper">
              <Routes {...rest} userRoles={roles}/>
          </div>
        </div>
      </div>
    )
  }
}

function mapStateToProps(store) {
  return {
      roles: store.user.roles
  };
}

export default withRouter(connect(mapStateToProps)(Dashboard))