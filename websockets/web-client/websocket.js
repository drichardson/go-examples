(function(window) {
  console.log("Running...");

  var ws = new WebSocket("ws://localhost:12345/echo");
  ws.onopen = function() {
    console.log("onopen");
    console.log("sending echo request");
    ws.send("hello, world");
  };
  ws.onmessage = function(e) {
    console.log("onmessage: data=" + e.data);
  };
  ws.onerror = function() {
    console.log("onerror");
  };
  ws.onclose = function(e) {
    console.log("onclose: wasClean=" + e.wasClean + ", code=" + e.code + ", reason=" + e.reason);
  };
})(window);
