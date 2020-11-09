package main

var defaultHtmlTemplate = `
<html>
    <head>
        <title>GPractice</title>
        <style>
            tr, .row {
                width: 100%;
                margin: auto;
                display: flex;
                padding: 1px;
            }
            td, th, .col {
                padding: 5px;
            }
        </style>
        <script type="application/javascript">
            // stupid HTML doesn't know anything about DELETE methods
            // so have to use js...
            function deleteFunc(id) {
                console.log("deleting item " + id);
                var xhttp = new XMLHttpRequest();
                xhttp.open("DELETE", "/app/" + id);
                xhttp.send();
                console.log("sent");
                window.location.assign("/app");
            }
        </script>
        <link rel="icon" type="image/png" href="data:image/png;base64,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">
    </head>

    <body>
        <a href="/app">home</a>

        <h1>Default template</h1>

        <form action="/app" method="post">
            {{ if and .Item .Item.Id }}
            <p>Update item: </p>
            <div class="trow">
                <label for="idInput">ID</label>
                <input type="text" id="idInput" name="idInput" value={{ .Item.Id }} readonly>
            </div>
            {{ else }}
            <p>Add item: </p>
            {{ end }}
            <div class="trow">
                <label for="dateInput">Date</label>
                <input type="date" id="dateInput" name="dateInput" value={{ .Item.Date }}>
            </div>
            <div class="trow">
                <label for="durationInput">Duration</label>
                <input type="text" id="durationInput"  name="durationInput"  value={{ .Item.Duration }}>
            </div>
            <input type="submit" id="saveButton" value="save">
        </form>

        {{ if and .Report .Report.Days }}
        <div>
            <h3>Report</h3>
            <div class="row">
                <h4 class="col">days</h4>
                <h4 class="col">since</h4>
                <h4 class="col">total</h4>
            </div>
            <div class="row">
                <p class="col">{{ .Report.Days }}</p>
                <p class="col">{{ .Report.Since }}</p>
                <p class="col">{{ .Report.Total.Minutes }}</p>
            </div>
        </div>
        {{ end }}

        {{ if .Items }}
        <div>
            <h3>History</h3>
            <table>
                <thead>
                    <tr>
                        <th>id</th>
                        <th>date</th>
                        <th>duration</th>
                    </tr>
                </thead>
                <tbody>
                {{ range .Items }}
                <tr>
                    <td>{{ .Id }}</td>
                    <td>{{ .Date }}</td>
                    <td>{{ .Duration }}</td>
                    <td><a href="/app/{{ .Id }}">Edit</a></td>
                    <td><input type="button" value="Del" onclick="deleteFunc({{ .Id }});"></td>
                </tr>
                {{ end }}
                </tbody>
            </table>
        </div>
        {{ end }}

    </body>
</html>
`
