{{ define "head" }}
<title>Mr A</title>
{{ end }}
{{ define "body" }}
{{ range . }}
{{ template "article" . }}
{{ end }}
{{ end }}
