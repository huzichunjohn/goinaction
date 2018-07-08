'use strict';

import { h, render, Component } from 'preact';
import LoginForm from '../components/LoginForm';
import HeaderBar from '../components/HeaderBar';

class Login extends Component {
    render() {
        return (
            <div>
                <HeaderBar showLogout={false} />
                <div class="container">
                    <LoginForm onSuccess={this.props.onLogin} />
                </div>
            </div>
        );
    }
}

export default Login;