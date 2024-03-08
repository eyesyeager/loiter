<template>
    <div class="login">
        <div class="title">{{ site.name }}</div>
        <div class="form">
            <el-input class="input" v-model="username" :prefix-icon="User" placeholder="请输入账号" />
            <el-input class="input" type="password" v-model="password" show-password :prefix-icon="Lock" placeholder="请输入密码" />
        </div>
        <div class="forget">
            <span @click="forgetPsd">忘记密码？</span>
        </div>
        <el-button plain class="action" @click="doLogin">登录</el-button>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { site, responseCode, storage } from "@/config";
import { User, Lock } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import api from "@/apis/api";
import { useRouter } from "vue-router";
import { publicPath } from "@/router/path";

const router = useRouter();
const username = ref('');
const password = ref('');

// 点击忘记密码，提示不支持修改
function forgetPsd() {
    ElMessage('暂不支持修改密码，请联系超级管理员重置！');
}

// 执行登录操作
function doLogin() {
    // 校验请求信息
    if (username.value.length > site.usernameMaxLen || username.value.length < site.usernameMinLen 
        || password.value.length > site.userPasMaxLen || password.value.length < site.userPasMinLen) {
        ElMessage({ type: "warning", message: "登录信息格式不规范" });
        return;
    }
    // 发送登录请求
    api.doLogin({
        username: username.value,
        password: password.value
    }).then(({code, msg}) => {
        if (code == responseCode.success) {
            location.reload();
        } else {
            ElMessage({ type: "error", message: "登录失败：" + msg });
        }
    });
}

// 不允许已登录用户进入登录页
function goHome() {
    if (localStorage.getItem(storage.token)) {
        router.replace(publicPath.home);
    }
}

onMounted(() => {
    goHome();
});
</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";
@import "@/assets/css/size.scss";

.login {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: $baseBackground;
    margin: auto;
    width: 420px;
    height: 450px;
    border: 1px solid $baseBorder;
    border-radius: $cardBorder;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    .title {
        color: $primaryText;
        font-size: 32px;
        font-weight: bold;
        height: 80px;
        line-height: 80px;
    }

    .form {
        height: 126px;
        text-align: center;
        margin: 10px 0;

        .input {
            width: 280px;
            height: 48px;
            margin: 10px auto;
        }
    }

    .forget {
        width: 280px;
        text-align: right;

        span {
            font-size: 12px;
            color: $secondaryText;
            cursor: pointer;
            transition: 0.3s;

            &:hover {
                color: $primaryText;
            }
        }
    }

    .action {
        width: 280px;
        height: 48px;
        margin: 30px 0;
    }
}</style>