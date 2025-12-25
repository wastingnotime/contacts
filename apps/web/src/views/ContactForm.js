import m from "mithril"
import {createContact, updateContact, getContact} from "../actions"
import {Button, Icons, Intent, Form, FormGroup, Input, FormLabel, Icon, Classes} from 'construct-ui'

let isInsertMode = true
let current = {}
const span = {
    xs: 12,
    sm: 12,
    md: 6
}
export default {
    oninit: v => {
        isInsertMode = v.attrs.id === undefined
        current = isInsertMode ? {firstName: '', lastName: '', phoneNumber: ''} :
            v.attrs.store.getState().current

        //recover updated contact (last backend version) and avoid re-render loop
        if (!isInsertMode && isEmpty(v.attrs.store.getState().current)){
            v.attrs.store.dispatch(getContact(v.attrs.id))
        }
    },
    view: v =>
        m(Form, {
            gutter: 15,
            onsubmit: e => {
                e.preventDefault()
                let action = isInsertMode ? createContact : updateContact
                v.attrs.store.dispatch(action(current))
                m.route.set("/")
            }
        }, [
            m(FormGroup, {span}, [
                m(FormLabel, {for: 'firstName'}, 'First name'),
                m(Input, {
                    contentLeft: m(Icon, {name: Icons.USER}),
                    id: 'firstName',
                    name: 'firstName',
                    placeholder: 'First name...',
                    oninput: e => current.firstName = e.target.value,
                    value: current.firstName
                })
            ]),
            m(FormGroup, {span}, [
                m(FormLabel, {for: 'lastName'}, 'Last name'),
                m(Input, {
                    contentLeft: m(Icon, {name: Icons.USER}),
                    id: 'lastName',
                    name: 'lastName',
                    placeholder: 'Last name...',
                    oninput: e => current.lastName = e.target.value,
                    value: current.lastName
                })
            ]),
            m(FormGroup, {span}, [
                m(FormLabel, {for: 'phoneNumber'}, 'Phone Number'),
                m(Input, {
                    contentLeft: m(Icon, {name: Icons.PHONE}),
                    id: 'phoneNumber',
                    name: 'phoneNumber',
                    placeholder: 'Phone Number...',
                    oninput: e => current.phoneNumber = e.target.value,
                    value: current.phoneNumber
                })
            ]),
            m(FormGroup, {class: Classes.ALIGN_RIGHT}, [
                m(Button, {
                    iconRight: Icons.CHEVRON_RIGHT,
                    type: 'submit',
                    label: 'Submit',
                    intent: Intent.PRIMARY
                })
            ])
        ])


}

const isEmpty = o => {
    for (let i in o) return false
    return true
}
