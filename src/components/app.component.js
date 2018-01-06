import React from "react";
import { BrowserRouter, Switch, Route } from "react-router-dom";

import PostsHome from "./posts-home.component";

export default class App extends React.Component {
    render() {
        return (
            <div className="container">
                <BrowserRouter>
                    <Switch>
                        <Route path="/" component={PostsHome} />
                    </Switch>
                </BrowserRouter>
            </div>
        );
    }
}