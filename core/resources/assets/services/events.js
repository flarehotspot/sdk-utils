define([], function () {
  var portalEvents = new EventSource("/api/portal/events");

  window.onbeforeunload = function () {
    portalEvents.close();
  };

  return portalEvents;
});
