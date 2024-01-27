import { createApp } from 'vue'
import App from './App.vue'
import router from "./router";

import 'reset-css/reset.css';
import 'animate.css';

const app = createApp(App);
app.use(router)
app.mount('#app')