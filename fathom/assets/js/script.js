'use strict';

import React from 'react';
import ReactDOM from 'react-dom';
import RealtimeVisitsCount from './components/realtime-visits';
import VisitList from './components/visits-list';
import PageviewsList from './components/pageviews';
import VisitsGraph from './components/visits-graph';

ReactDOM.render(
    <div className="container">
        <h1>Ana</h1>
        <RealtimeVisitsCount />
        <VisitsGraph />
        <PageviewsList />
    </div>,
    document.getElementById('root')
);
