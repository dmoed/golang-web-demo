import React from 'react'
import {Switch, Route} from 'react-router-dom'

//get pages
import HomePage from './scenes/Home';
import NoMatch from './scenes/NoMatch';
import PrivateRoute from './components/PrivateRoute';

//set AUTH
import auth from './constants/auth'
const {USER} = auth;

const Routes = (RouteProps) => {
    return (
        <Switch>
            <PrivateRoute exact path='/' component={HomePage} config={RouteProps.config}
                          accessRoles={USER} userRoles={RouteProps.userRoles}/>

            <Route component={NoMatch}/>
        </Switch>
    );
};

export default Routes