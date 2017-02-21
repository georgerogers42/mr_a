<!doctype html>
<html>
    <head>
        {{ block "css" . }}
        <link rel="stylesheet" href="/css/style.css">
        {{ end }}
        {{ template "head" . }}
    </head>
    <body>
        <h1 class="masthead"><a href="/">Mr A's Journal</a></h1>
        {{ template "body" . }}
    </body>
</html>
