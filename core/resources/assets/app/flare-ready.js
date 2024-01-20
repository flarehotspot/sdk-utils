(function ($flare) {
  var callbacks = [];
  $flare.ready = function (cb) {
    callbacks.push(cb);
  };

  $flare._triggerReady = function () {
    for (var i = 0; i < callbacks.length; i++) {
      callbacks[i]();
    }
  };
})(window.$flare);
