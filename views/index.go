package views

var Str_index =`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Ghost Man</title>
    <style>
        .row{
            flow: horizontal;
            width: *;            
        }
    </style>
</head>

<body>
    <div class="row">
        <select id="method">
            <option value="get">GET</option>
            <option value="post">POST</option>
            <option value="put">PUT</option>
        </select>
        <input type="text" id="url" placeholder="http://example.com/api" value="http://" style="width: *;">
        <input type="button" id="send" value="send">
    </div>
    <script type="text/tiscript">
        self.ready=function(){
    
}
    </script>
</body>

</html>`
