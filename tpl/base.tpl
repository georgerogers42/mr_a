<!doctype html>
<html>
    <head>
        {{ block "css" . }}
        <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Inconsolata|Noto+Serif">
        <link rel="stylesheet" href="/css/style.css">
        {{ end }}
        {{ template "head" . }}
    </head>
    <body>
        <div class="masthead">
            <h1 class="masthead"><a href="/">Mr A's Journal</a></h1>
        </div>
        {{ template "body" . }}
    </body>
</html>
