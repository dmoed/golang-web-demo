/**
 * Created by wardo on 3/14/2018.
 */
import React from 'react'
import {connect} from 'react-redux'
import ProfilePopup from './index'
import {toggleProfilePopup} from './../../actions'
// import {unsubscribeBrowser} from './../NotificationPopup/actions'

const ProfilePopupContainer = ({pushSubscribed, showProfilePopup, displayname, email, closeProfilePopup, onLogout}) => {
    return (
        <div>
            {showProfilePopup === true &&
                <ProfilePopup displayname={displayname} email={email}
                              onLogout={() => onLogout(pushSubscribed)}
                              close={() => closeProfilePopup()}/>}
        </div>
    );
};

function mapStateToProps(store){
    return {
        showProfilePopup: store.dashboard.showProfilePopup,
        displayname: store.user.displayname,
        email: store.user.email,
        pushSubscribed: false
    }
}

function mapDispatchToProps(dispatch, ownProps){

    const {config: {routes: {logout, admin_ajax_unsubscribe_device}}} = ownProps;

    return {
        closeProfilePopup: () => {
            dispatch(toggleProfilePopup(false))
        },
        onLogout: (pushSubscribed) => {

            // if(pushSubscribed){

            //     dispatch(unsubscribeBrowser(admin_ajax_unsubscribe_device));

            //     setTimeout(() => window.location.assign(logout), 400);

            // }else{
                window.location.assign(logout);
            // }
        }
    }
}

export  default connect(mapStateToProps, mapDispatchToProps)(ProfilePopupContainer)
