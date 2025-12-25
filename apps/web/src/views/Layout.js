import m from "mithril"
import 'construct-ui/lib/index.css'
import { version } from '../../package.json'

export default {
    view: v => m("main.layout", [
        m("section", v.children),
        m("footer", version)
    ])
}
