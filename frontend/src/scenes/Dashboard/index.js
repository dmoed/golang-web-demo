import React from "react"
import MyNavbar from './components/NavBar/index'

import 'bootstrap/dist/css/bootstrap.css'
import './style.scss'

class Dashboard extends React.Component {
  render(){

    const {config} = this.props;

    return (
      <div id="dashboard">
        <header>
          <MyNavbar config={config}/>
          </header>
        

        
      </div>
    )
  }
}

export default Dashboard