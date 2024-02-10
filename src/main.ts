import { createApp } from "vue"
import App from "./App.vue"
import router from "./router";

import "reset-css/reset.css";
import "@/assets/css/global.css";
import "element-plus/dist/index.css";

const app = createApp(App);
app.use(router)
app.mount('#app')