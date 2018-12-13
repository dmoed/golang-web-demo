
import React from 'react';
import {connect} from 'react-redux'
import {Link} from 'react-router-dom'
import {toggleSidebar, toggleProfilePopup, toggleNotificationPopup} from '../../actions'
import Avatar from './components/Avatar'
import './style.scss'

import {
    Collapse,
    Navbar,
    NavbarToggler,
    NavbarBrand,
    Nav,
    NavItem,
    NavLink
} from 'reactstrap';

class MyNavbar extends React.Component {
    render() {
        const {config: {app_name = "", app_logo = "", app_version = ""}, nCounter, displayname, openSidebar, openProfilePopup, openNotificationPopup} = this.props;

        return (
            <Navbar color="default" dark fixed="top" expand="md">
                <div className="navbar-sidebar-toggle">
                    <span className="mdi mdi-menu" onClick={openSidebar}/></div>

                <div className="navbar-branding">
                    <Link to="/" className="navbar-brand-img">
                        <img className="app-logo" src={app_logo}/>
                    </Link>

                    <Link to="/" className="navbar-brand-name">{app_name}
                        <small className="app-version">{app_version}</small>
                    </Link>
                </div>

                <ul className="navbar-nav mr-auto">
                    <NavItem>
                        <a className="nav-notification-link" onClick={openNotificationPopup}>
                                <span className="mdi mdi-bell notification-bell">
                                    {nCounter > 0
                                        ? <span className="notification-counter">{nCounter > 9 ? "9+" : nCounter}</span>
                                        : null }
                                </span>
                        </a>
                    </NavItem>
                </ul>

                <Avatar displayname={displayname} onClick={openProfilePopup}/>
            </Navbar>
        );
    }
}

function mapStateToProps(store) {
    return {
        nCounter: 0,
        displayname: store.user.displayname,
    }
}

function mapDispatchToProps(dispatch) {

    return {
        openNotificationPopup: () => {
            dispatch(toggleNotificationPopup(true));
        },
        openProfilePopup: () => {
            dispatch(toggleProfilePopup(true));
        },
        openSidebar: () => {
            dispatch(toggleSidebar(true));
        }
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(MyNavbar)