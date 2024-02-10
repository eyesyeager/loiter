let buildParamUrl = (baseUrl: string, req: any) => {
    if (req == null) return baseUrl;
    let url = baseUrl + "?";
    let propertyList = Object.keys(req);
    for (let i = 0; i < propertyList.length; i++) {
        url += propertyList[i] + "=" + req[propertyList[i]];
        if (i != propertyList.length - 1) url += "&";
    }
    return url;
};

let buildPathUrl = (url: string, req: any) => {
    if (req == null) return url;
    req.forEach((v: string | number) => {
        url += ("/" + v);
    });
    return url;
};

let buildMixUrl = (url: string, req: any) => {
    if(req == null) return url;
    if(req.param == null) return buildPathUrl(url, req.path);
    if(req.path == null) return buildParamUrl(url, req.param);
    return buildParamUrl(
        buildPathUrl(url, req.path),
        req.param
    );
}

export default [buildParamUrl, buildPathUrl, buildMixUrl];