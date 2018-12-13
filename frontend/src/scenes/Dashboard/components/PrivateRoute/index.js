/**
 * Created by wardo on 8/13/2018.
 */
import React from "react";
import {connect} from "react-redux";
import {Route, Redirect} from 'react-router-dom'
// import ForbiddenPage from './../../scenes/Forbidden';

const PrivateRoute = ({ component: Component, userRoles, accessRoles, config, ...rest,  }) => {
    return (
        <Route {...rest} render={(props) => (
            accessRoles.some(access => userRoles.includes(access))
              ? <Component {...props} config={config} userRoles={userRoles}/>
              : <p>Forbidden</p>
        )}/>
    );
};

export default PrivateRoute