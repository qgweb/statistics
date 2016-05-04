<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>投放数据统计</title>
    <!-- 新 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
    <!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
    <script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="//cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
</head>
<body>
<div class="btn-group" role="group">
    <a type="button" class="btn btn-default" href="/stats?qy=zj">浙江</a>
    <a type="button" class="btn btn-default" href="/stats?qy=js">江苏</a>
</div>
<table class="table table-hover table-bordered">
    <thead>
    <tr>
        <th colspan="4" class="text-center">投放流程数据统计</th>
    </tr>
    <tr>
        <th>时间</th>
        <th>来源数据</th>
        <th>广告数据量</th>
        <th>电信写入量</th>
    </tr>
    </thead>
    <tbody>
        {{range $k,$v:= .info}}
        <tr>
            <td>{{$v.Timestamp | unix}}</td>
            <td>
                {{range $kk,$vv:=$v.SourceData}}
                    {{$kk}} - {{$vv}}<br>
                {{end}}
            </td>
            <td>
                {{range $kk,$vv:=$v.AdvertData}}
                {{$kk}} - {{$vv}}<br>
                {{end}}
            </td>
            <td>
                {{$v.DXData}}
            </td>
        </tr>
        {{end}}
    </tbody>
</table>
</body>
</html>