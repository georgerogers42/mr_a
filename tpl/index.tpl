{{ define "head" }}
<title>Mr A</title>
{{ end }}
{{ define "body" }}
{{ range . }}
<div class="article">
<h1 class="article">{{ .Serial }}&ndash;{{ .Title }}</h1>
<h2 class="article">{{ .Posted }}</h2>
<div class="contents">
{{ .HtmlContent }}
</div>
</div>
{{ end }}
{{ end }}
