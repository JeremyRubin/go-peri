<!DOCTYPE html>
<html>
  <head>
    <script src="http://ajax.googleapis.com/ajax/libs/jquery/1.4.2/jquery.min.js"></script>
    <script src="http://cdn.sockjs.org/sockjs-0.3.4.min.js"></script>
    <script>
    $(function() {
        var conn = null;

        function log(msg) {
            var control = $('#log');
            control.html(control.html() + msg + '<br/>');
            control.scrollTop(control.scrollTop() + 1000);
        }

        function disconnect() {
            if (conn != null) {
                log('Disconnecting...');

                conn.close();
                conn = null;

                updateUi();
            }
        }

        function updateUi() {
            if (conn == null || conn.readyState != SockJS.OPEN) {
                $('#status').text('disconnected');
                $('#connect').text('Connect');
            } else {
                $('#status').text('connected (' + conn.protocol + ')');
                $('#connect').text('Disconnect');
            }
        }

        $('form').submit(function() {
            var text = $('#message').val();
            conn.send(text);
            $('#message').val('').focus();
            return false;
        });

        conn = new SockJS('http://' + window.location.host + '/chat');
        log('Connecting...');

        conn.onopen = function() {
            log('Connected.');
            updateUi();
        };

        conn.onmessage = function(e) {
            log(e.data);
        };

        conn.onclose = function() {
            log('Disconnected.');
            conn = null;
            updateUi();
        };

        $('#message').val('').focus();
    });
    </script>
    <title>Sockjs-go chat</title>
  </head>
  <body>
    <h1>Sockjs-go chat</h1>

    <div>
      Status: <span id="status">disconnected</span>
    </div>
    <div id="log" style="width: 60em; height: 20em; overflow:auto; border: 1px solid black">
    </div>
    <form id="chatform">
      <label for="message">Message:</label>
      <input id="message" type="text" />
      <input type="submit" value="Send" />
    </form>
  </body>
</html>