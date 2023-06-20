import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
// 导入element-plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import { FontAwesomeIcon } from './plugins/font-awesome'

const app = createApp(App)
app.use(store)
app.use(router)
app.use(ElementPlus)
app.component("font-awesome-icon", FontAwesomeIcon)
app.mount("#app");

// 1.实例化Vue。
// 2.放置项目中经常会用到的插件和CSS样式。例如： 网络请求插件:axios和vue-resource、图片懒加载插件：vue-lazyload
// 3.存储全局变量。例如（用于的基本信息）

