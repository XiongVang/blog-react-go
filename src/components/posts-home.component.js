import React from "react";

import List from "./posts-list.component";
import Details from "./posts-details.component"

export default class PostsHome extends React.Component {

    render() {

        const items = _.values({
            1: { id: "1", title: "Item 1" },
            2: { id: "2", title: "Item 2" },
            3: { id: "3", title: "Item 3" },
            4: { id: "4", title: "Item 4" },
            5: { id: "5", title: "Item 5" },
            6: { id: "6", title: "Item 6" },
            7: { id: "7", title: "Item 7" }
        }).sort(function (a, b) {
            if (a.id >= b.id) {
                return -1
            }
            return 0;
        });

        const active = { id: "3", title: "Item 3" };

        return (
            <div className="row">
                <nav className="col-md-3">
                    <List items={items} active={active} />
                </nav>
                <main className="col-md-9">
                    <Details item={active} />
                </main>
            </div>
        );
    }
}