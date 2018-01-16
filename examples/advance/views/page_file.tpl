<!-- /views/page_file.html -->
<!doctype html>

<html>
    <head>
        <title>{{.title}}</title>
        {{template "head" .}}
    </head>

    <body>
        <a href="/"><- Back home!</a>
        <hr>
        {{include "layouts/footer"}}
    </body>
</html>