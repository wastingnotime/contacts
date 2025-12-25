import m from "mithril"
import {createContact, updateContact, getContact} from "../actions"
import {Button, Icons, Intent, Form, FormGroup, Input, FormLabel, Icon, Classes} from 'construct-ui'

const span = {
    xs: 12,
    sm: 12,
    md: 6
}

const emptyContact = () => ({
    firstName: '',
    lastName: '',
    phoneNumber: ''
})
export default {
    oninit: v => {
        v.state.isInsertMode = v.attrs.id === undefined
        v.state.current = emptyContact()
        if (!v.state.isInsertMode) {
            v.state.lastId = v.attrs.id
            const fromStore = v.attrs.store.getState().current
            if (fromStore && String(fromStore.id) === String(v.attrs.id)) {
                v.state.current = { ...fromStore }
            }
            v.attrs.store.dispatch(getContact(v.attrs.id))
        }
    },
    onbeforeupdate: v => {
        const isInsertMode = v.attrs.id === undefined
        if (v.state.isInsertMode !== isInsertMode) {
            v.state.isInsertMode = isInsertMode
            if (isInsertMode) {
                v.state.current = emptyContact()
                v.state.lastId = undefined
                return
            }
        }
        if (isInsertMode) {
            return
        }
        if (v.state.lastId !== v.attrs.id) {
            v.state.lastId = v.attrs.id
            v.state.current = emptyContact()
            v.attrs.store.dispatch(getContact(v.attrs.id))
            return
        }
        const fromStore = v.attrs.store.getState().current
        if (fromStore && fromStore.id && fromStore.id !== v.state.current.id) {
            v.state.current = { ...fromStore }
        }
    },
    view: v =>
        m(Form, {
            gutter: 15,
            onsubmit: e => {
                e.preventDefault()
                let action = v.state.isInsertMode ? createContact : updateContact
                v.attrs.store.dispatch(action(v.state.current))
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
                    oninput: e => v.state.current.firstName = e.target.value,
                    value: v.state.current.firstName
                })
            ]),
            m(FormGroup, {span}, [
                m(FormLabel, {for: 'lastName'}, 'Last name'),
                m(Input, {
                    contentLeft: m(Icon, {name: Icons.USER}),
                    id: 'lastName',
                    name: 'lastName',
                    placeholder: 'Last name...',
                    oninput: e => v.state.current.lastName = e.target.value,
                    value: v.state.current.lastName
                })
            ]),
            m(FormGroup, {span}, [
                m(FormLabel, {for: 'phoneNumber'}, 'Phone Number'),
                m(Input, {
                    contentLeft: m(Icon, {name: Icons.PHONE}),
                    id: 'phoneNumber',
                    name: 'phoneNumber',
                    placeholder: 'Phone Number...',
                    oninput: e => v.state.current.phoneNumber = e.target.value,
                    value: v.state.current.phoneNumber
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
