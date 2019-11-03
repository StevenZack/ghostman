package views

var Str_indexlogic =`self.ready=function(){
    $(#send).on("click",function() {
        var method=$(#method).value;
        var url=$(#url).value;
        var cypher=$(#cypher).value;
        var body=$(#requestbody).value;
        view.doReq(method,url,cypher,'',body)
    })
}

function setupData(method,url,cypher,header,rbody) {
    $(#method).value=method;
    $(#url).value=url;
    $(#cypher).value=cypher;
    $(#requestbody).value=rbody;
}

function response(status,header,body){
    $(#rpstatus).text=status;
    $(#rpheaders).text=header;
    $(#rpbody).value=body;
}

function showErr(e){
    view.showErr(e);
}`
