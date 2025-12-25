import "./style.css"
import m from "mithril"
import ContactList from "./views/ContactList"
import ContactForm from "./views/ContactForm"
import Layout from "./views/Layout"
import reducer from './reducers'
import {getAllContacts} from './actions'
import {configureStore} from '@reduxjs/toolkit'

const store = configureStore({reducer})

store.dispatch(getAllContacts());

const render = () => m.route(document.body, "/", {
    "/": {
        render: () => m(Layout, m(ContactList, {store}))
    },
    "/edit/:id": {
        render: v => m(Layout, m(ContactForm, {...v.attrs, store}))
    },
    "/new": {
        render: v => m(Layout, m(ContactForm, {...v.attrs, store}))
    }
})

store.subscribe(render)

render()
