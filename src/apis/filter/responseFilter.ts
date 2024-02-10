import progress from "@/utils/nprogress";
import { AxiosResponse } from "axios";

export default (response: AxiosResponse<any, any>) => {
    progress.close();
}