import moment from "moment";

// 跳转外部网站
export function jumpSite(url: string) {
    open(url);
}

// 获取以当前日期为中心，指定日期的格式化字符串
export function getDayFormat(num?: number): string {
    return num == null ? moment().format("YYYY-MM-DD") : moment().add(num, "days").format("YYYY-MM-DD");
}