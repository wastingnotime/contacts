import "./style.css"
import m from "mithril"
import ContactList from "./views/ContactList"
import ContactForm from "./views/ContactForm"
import Layout from "./views/Layout"
import { createStore, applyMiddleware } from 'redux'
import { thunk } from 'redux-thunk'
import reducer from './reducers'
import { getAllContacts } from './actions'

const store = createStore(reducer, applyMiddleware(thunk))

store.dispatch(getAllContacts());

const render = () => {
    m.route(document.body, "/", {
        "/": {
            render: () => m(Layout, m(ContactList, { store }))
        },
        "/edit/:id": {
            render: v => m(Layout, m(ContactForm, { ...v.attrs, store }))
        },
        "/new": {
            render: v => m(Layout, m(ContactForm, { ...v.attrs, store }))
        }
    })
}

store.subscribe(render)

render()