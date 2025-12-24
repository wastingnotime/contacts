import m from "mithril"

export default {
    view: v => m("main.layout", [
        m("nav.menu", [
            m(m.route.Link, { class: 'button', href: "/new" }, "New")
        ]),
        m("section", v.children)
    ])
}