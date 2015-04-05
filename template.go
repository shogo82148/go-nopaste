package nopaste

import "html/template"

var tmplRoot = template.Must(template.New("tmplRoot").Parse(`{{define "index"}}
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Nopaste</title>
    </head>
  <body>
  <h1>Nopaste</h1>
    <form method="post">
    <p>
       <textarea id="text" name="text" rows="24" cols="100"></textarea>
    </p>
    <p>
       <input type="submit" value="Paste it">
    </p>
    </form>
  </body>
</html>
{{end}}
`))
