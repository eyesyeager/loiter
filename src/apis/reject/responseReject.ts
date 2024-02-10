import progress from "@/utils/nprogress";
import { ElMessage } from "element-plus";

export default (error: any) => {
    // 关闭进度条
    progress.close();
    // 弹出报错
    ElMessage({
        type: "error", 
        message: error.message
    })
}