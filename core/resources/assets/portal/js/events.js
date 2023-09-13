(function(window) {
  window.PortalEvents = new EventSource('/api/portal/events');
  window.onbeforeunload = function() {
    window.PortalEvents.close();
  };
})(window);
