package views

var Str_error =`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Infomation</title>
</head>
<body>
    map[web/error.html:<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Infomation</title>
</head>
<body>
    {{.}}
</body>
</html> web/index.html:<!DOCTYPE html>
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
    <div>body:</div>
    <textarea id="requestbody" cols="30" rows="10"></textarea>
    <h2>
        Response: 
    </h2>
    <div>status:</div>
    <div id="rpstatus"></div>
    <div>headers:</div>
    <div id="rpheaders"></div>
    <div>body:</div>
    <textarea id="rpbody" cols="30" rows="10" readonly></textarea>
    <script type="text/tiscript">
        {{index . "web/indexlogic.tis"}}
    </script>
</body>

</html> web/indexlogic.tis:self.ready=function(){
    $(#send).on("click",function() {
        var method=$(#method).value;
        var url=$(#url).value;
        var body=$(#requestbody).value;
        view.doReq(method,url,'',body)
    })
}

function response(status,header,body){
    $(#rpstatus).text=status;
    $(#rpheaders).text=header;
    $(#rpbody).value=body;
}]
</body>
</html>`
