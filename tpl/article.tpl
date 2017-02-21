{{ define "head" }}
<title>Mr A</title>
{{ end }}
{{ define "body" }}
<div class="article">
<h1 class="article"><a href="/article/{{ .Serial }}">{{ .Serial }}</a>&ndash;{{ .Title }}</h1>
<h2 class="article">{{ .Posted }}</h2>
<div class="contents">
{{ .HtmlContent }}
</div>
</div>
{{ end }}
