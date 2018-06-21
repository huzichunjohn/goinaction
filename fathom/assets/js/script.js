'use strict';

import React from 'react';
import ReactDOM from 'react-dom';
import RealtimeVisitsCount from './components/realtime-visits';
import VisitList from './components/visits-list';
import PageviewsList from './components/pageviews';
import VisitsGraph from './components/visits-graph';
import Login from './components/login';

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = { idToken: null };
    }

    render() {
        if (this.state.idToken) {
            return (
                <div className="container">
                    <h1>Ana</h1>
                    <RealtimeVisitsCount />
                    <VisitsGraph />
                    <PageviewsList />
                </div>
            );
        } else {
            return (
                <div className="container">
                    <Login />
                </div>
            );
        }
    }
}

ReactDOM.render(
    <App />,
    document.getElementById('root')
);
