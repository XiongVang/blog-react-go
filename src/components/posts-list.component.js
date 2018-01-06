import React from "react";
import _ from "lodash";

export default class PostsList extends React.Component {

    renderList() {
        const active = this.props.active;
        if (this.props.items) {
            return this.props.items.map(function (item) {
                const bootClass = item.id === active.id ? "list-group-item list-group-item-action active" : "list-group-item list-group-item-action";
                return (
                    <a href="#" className={bootClass} key={item.id}>{item.title}</a>
                );
            });
        } else {
            return <div></div>
        }
    }

    render() {

        const items = this.renderList();

        return (
            <div>
                <h2>Posts</h2>
                <div className="list-group">
                    {items}
                </div>
            </div>
        );
    }
}