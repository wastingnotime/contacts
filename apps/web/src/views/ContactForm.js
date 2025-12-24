import m from "mithril"
import { createContact, updateContact, getContact } from "../actions"

let isInsertMode = true;
let current = {}

export default {
    oninit: v => {
        isInsertMode = v.attrs.id === undefined

        current = isInsertMode ? { firstName: '', lastName: '', phoneNumber: '' } :
            v.attrs.store.getState().current

        //recover updated contact (last backend version) and avoid re-render loop
        if (!isInsertMode && isEmpty(v.attrs.store.getState().current))
            v.attrs.store.dispatch(getContact(v.attrs.id))
    },
    view: v => {
        return m("form", {
            onsubmit: e => {
                e.preventDefault()
                let action = isInsertMode ? createContact : updateContact
                v.attrs.store.dispatch(action(current))
                m.route.set("/")
            }
        }, [
            m("label.label", "First name"),
            m("input.input[type=text][placeholder=First name]", {
                oninput: e => current.firstName = e.target.value,
                value: current.firstName
            }),
            m("label.label", "Last name"),
            m("input.input[placeholder=Last name]", {
                oninput: e => current.lastName = e.target.value,
                value: current.lastName
            }),
            m("label.label", "Phone Number"),
            m("input.input[placeholder=Phone Number]", {
                oninput: e => current.phoneNumber = e.target.value,
                value: current.phoneNumber
            }),
            m("button.button[type=submit]", "Save")
        ])
    }
}

const isEmpty = o => {
    for (let i in o) return false
    return true
}