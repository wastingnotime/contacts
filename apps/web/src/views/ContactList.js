import m from "mithril"
import {deleteContact} from "../actions"
import {Button, EmptyState, Icons, Intent, List, ListItem} from 'construct-ui'

export default {
    view: v => v.attrs.store.getState().contacts.length > 0 ?
        [
            m(List,
                {
                    interactive: true,
                    size: 'xl'
                },
                v.attrs.store.getState().contacts.map(item =>
                    m(ListItem, {
                        contentRight: m(Button, {
                            onclick: () => v.attrs.store.dispatch(deleteContact(item.id)),
                            size: "xl",
                            iconLeft: Icons.X,
                        }),
                        label: `${item.firstName} ${item.lastName}`,
                        onclick: () => m.route.set(`/edit/${item.id}`)
                    }),
                )
            ),
            m(Button, {
                onclick: () => m.route.set("/new"),
                iconLeft: Icons.PLUS,
                intent: Intent.PRIMARY,
                fluid: true
            })
        ] :
        m(EmptyState, {
            icon: Icons.ARCHIVE,
            header: ['No data has found'],
            fill: true,
            content: m(Button, {
                onclick: () => m.route.set("/new"),
                iconLeft: Icons.PLUS,
                intent: Intent.PRIMARY,
                fluid: true
            })
        })
}
