import m from "mithril"
import {deleteContact} from "../actions"

export default {
    view: v => m(".contact-list", v.attrs.store.getState().contacts.map(o =>
        m("o", [
            m(m.route.Link, { class: 'contact-list-item', href: "/edit/" + o.id }, o.firstName + " " + o.lastName),
            m("a.button", { onclick: () => v.attrs.store.dispatch(deleteContact(o.id)) }, "X"),
        ])
    ))
}