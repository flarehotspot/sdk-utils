window.DomReady(function () {
  var events = window.PortalEvents;

  events.addEventListener("client:connected", function (res) {
    var data = JSON.parse(res.data);
    console.log(data);
  });

  events.addEventListener("client:disconnected", function (res) {
    var data = JSON.parse(res.data);
    console.log(data);
  });

  events.onerror = function (e) {
    console.log("Socket error:", e);
  };
});
