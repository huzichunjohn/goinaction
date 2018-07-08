'use strict';

import { h, render, Component } from 'preact';
import LogoutButton from '../components/LogoutButton.js';
import Pageviews from '../components/Pageviews.js';
import Realtime from '../components/Realtime.js';
import Graph from '../components/Graph.js';
import DatePicker from '../components/DatePicker.js';
import Table from '../components/Table';
import HeaderBar from '../components/HeaderBar';

class Dashboard extends Component {
  constructor(props) {
    super(props);

    this.state = {
      period: 7
    }
  }

  render() {
    return (
        <div>
            <HeaderBar showLogout={true} onLogout={this.props.onLogout} />
            <div class="container">
              <Realtime />
              <div class="clear">
                <DatePicker period={this.state.period} onChoose={(p) => { this.setState({ period: p })}} />
              </div>
              <Graph period={this.state.period} />
              <div class="row">
                <div class="col-4">
                  <Pageviews period={this.state.period} />
                </div>
                <div class="col-2">
                  <Table period={this.state.period} endpoint="languages" title="Languages" headers={["#", "Language", "Count", "%"]} /> 
                </div>     
              </div>
              <div class="row">
                <div class="col-2">
                  <Table period={this.state.period} endpoint="screen-resolutions" title="Screen Resolutions" headers={["#", "Screen Resolution", "Count", "%"]} />
                </div>
                <div class="col-2">
                  <Table period={this.state.period} endpoint="countries" title="Countries" headers={["#", "Country", "Count", "%"]} />              
                </div>
                <div class="col-2">
                  <Table period={this.state.period} endpoint="browsers" title="Browsers" headers={["#", "Browser", "Count", "%"]} />              
                </div>
              </div>
            </div>
        </div>
    );
  }
}

export default Dashboard;