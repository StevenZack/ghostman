package views

var Str_indexlogic =`self.ready=function(){
    $(#send).on("click",function() {
        var method=$(#method).value;
        var url=$(#url).value;
        var body=$(#requestbody).value;
        view.doReq(method,url,'',body)
    })
}

function setupData(method,url,rbody) {
    $(#method).value=method;
    $(#url).value=url;
    $(#requestbody).value=rbody;
}

function response(status,header,body){
    $(#rpstatus).text=status;
    $(#rpheaders).text=header;
    $(#rpbody).value=body;
}`
