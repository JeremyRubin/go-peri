{{/* Brought to you by Jeremy Rubin, 2013 */}}
<html>
<head>
	<title>This Site</title>

    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link href='/static/bootstrap/css/bootstrap.css' rel='stylesheet' media='screen'>
    <link href='/static/css/custom.css' rel='stylesheet'>
    <link href="/static/bootstrap/css/bootstrap-responsive.css" rel="stylesheet">
    <link href="/static/bootstrap/css/slider.css" rel="stylesheet">
    <link href="/static/bootstrap/css/switch.css" rel="stylesheet">
</head>

<body>

    <div class='container-fluid' id='main_content'>
        {{ range $row, $row_data := .Layout }}

        <div class = 'row-fluid'>
            {{ range $column, $column_data := $row_data }}
            <div class='span3 {{ $column_data.type }}' data-device='{{ $column_data.name }}' id="box{{ $row }}x{{ $column }}"> </div>
            {{ end }}
        </div>
        {{ end }}

        <!--<div class='well'id='debug'>debug</div> -->
    </div>


    {{/* SCRIPTS: */}}
    <script src='/static/js/jquery-2.0.2.min.js'></script>
    <script src='/static/js/sockjs-0.3.4.min.js'></script>
    <script src='/static/bootstrap/js/bootstrap.min.js'></script>
    <script src='/static/bootstrap/js/bootstrap-slider.js'></script>
    <script src='/static/bootstrap/js/switch.js'></script>
    <script src='/static/js/slider.js'></script>



</body>
</html>