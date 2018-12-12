import React from 'react'
import ReactDOM from 'react-dom'
import {Provider} from 'react-redux'
import {createStore, applyMiddleware} from 'redux'

import reducer from './combineReducer'
import thunk from 'redux-thunk'
import logger from 'redux-logger'
import {BrowserRouter} from 'react-router-dom'

import Dashboard from './index'

let middleware = [thunk];
let basename = "/dashboard/";

//Parse INITIAL_STATE from HTML
const appData = JSON.parse(document.getElementById("app-data").text);

if (process.env.NODE_ENV !== 'production') {
  middleware = [...middleware, logger];
  //hide json
}

const store = createStore(reducer, appData.__initial_state__, applyMiddleware(...middleware));

ReactDOM.render(
  <Provider store={store}>
      <BrowserRouter basename={basename}>
          <Dashboard {...appData.__props__}/>
      </BrowserRouter>
  </Provider>,
  document.getElementById('app')
);

if (module.hot) {
  module.hot.accept();
}