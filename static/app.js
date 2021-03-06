var sock = null;

var wsuri = "ws://127.0.0.1:9090/websocket";

window.onload = function() {

    console.log("onload");

    sock = new WebSocket(wsuri);

    sock.onopen = function() {

        console.log("connected to " + wsuri);

    }

    sock.onclose = function(e) {

        console.log("connection closed (" + e.code + ")");

    }

    sock.onmessage = function(e) {

        console.log("message received: " + e.data);

    }

};

function send() {

    var msg = document.getElementById('message').value;
    sock.send(msg);
};