import { ElMessage } from "element-plus";

export default (error: any) => {
    ElMessage({
        type: "error", 
        message: error.message
    })
}