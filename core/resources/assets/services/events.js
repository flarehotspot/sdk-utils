/**
 * @file             : events.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.ph>
 * Date              : Jan 25, 2024
 * Last Modified Date: Jan 25, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */
define([], function () {
  var portalEvents = new EventSource("/api/portal/events");

  window.onbeforeunload = function () {
    portalEvents.close();
  };

  return portalEvents;
});
