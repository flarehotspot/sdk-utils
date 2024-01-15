(function(window) {
  window.AdminEvents = new EventSource('/api/admin/events');
  window.onbeforeunload = function() {
    window.AdminEvents.close();
  };
})(window);
