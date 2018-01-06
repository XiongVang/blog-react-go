import React from "react";

export default class PostsDetails extends React.Component {
    render() {
        return (
            <h2>{this.props.item.title}</h2>
        );
    }
}