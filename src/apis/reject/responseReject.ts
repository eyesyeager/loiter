import progress from "@/utils/nprogress";
import { ElMessage } from "element-plus";
import { responseCode } from "@/config";

export default (error: any) => {
    // 关闭进度条
    progress.close();
    // 特殊状态码处理
    let code = error.response.status;
    switch (code) {
        case responseCode.tooManyRequest: 
            requestLimit();
            return;
        case responseCode.badGateway:
            badGateway();
            return;
        default:
            ElMessage({
                type: "error", 
                message: error.message
            });
    }
}

// 限制请求
function requestLimit() {
    alert("请求频率过高，请稍后再试")
}

// 网关异常
function badGateway() {
    alert("网关异常，请联系管理员处理")
}
