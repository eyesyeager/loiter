import moment from "moment";

// 跳转外部网站
export function jumpSite(url: string) {
    open(url);
}

// 获取以当前日期为中心，指定日期的格式化字符串
export function getDayFormat(num?: number): string {
    return num == null ? moment().format("YYYY-MM-DD") : moment().add(num, "days").format("YYYY-MM-DD");
}

// 根据换行拆分字符串
export function getListByWrap(str: string): string[] {
    if (str == null || str == "") {
        return [];
    }
    return str.split(/\r?\n/);;
}

// 校验是否为IPv4格式
export function checkIPv4(ip: string): boolean {
    return /^(\d{1,3}\.){3}\d{1,3}$/.test(ip);
}