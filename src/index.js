import "babel-polyfill";
import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { createStore, applyMiddlware } from "redux";
import reduxPromise from "redux-promise";

import reducers from "./reducers";
import { applyMiddleware } from "../../../Library/Caches/typescript/2.6/node_modules/redux";

import App from "./components/app.component"
const createStoreWithMiddleware = applyMiddleware(reduxPromise)(createStore);

ReactDOM.render(
    <Provider store={createStoreWithMiddleware(reducers)}>
        <App />
    </Provider>,
    document.getElementById("root")
);