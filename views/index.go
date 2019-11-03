package views

var Str_index =`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Ghost Man</title>
    <style>
        .row {
            flow: horizontal;
            width: *;
        }
    </style>
</head>

<body>
    <div class="row">
        <select id="method">
            <option value="GET">GET</option>
            <option value="POST">POST</option>
            <option value="PUT">PUT</option>
        </select>
        <input type="text" id="url" placeholder="http://example.com/api" value="http://" style="width: *;">
        <input type="button" id="send" value="send">
    </div>
    <h2>
        Request:
    </h2>
    <div>
        <span>Cypher:</span>
        <input type="text" id="cypher">
    </div>
    <div>body:</div>
    <textarea id="requestbody" cols="30" rows="10"></textarea>
    <h2>
        Response: 
    </h2>
    <br>
    <div>status:</div>
    <div id="rpstatus"></div>
    <br>
    <div>headers:</div>
    <div id="rpheaders"></div>
    <br>
    <div>body:</div>
    <textarea id="rpbody" cols="30" rows="10" readonly style="width: *;"></textarea>
    <script type="text/tiscript">
        self.ready=function(){
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
}
    </script>
</body>

</html>`
