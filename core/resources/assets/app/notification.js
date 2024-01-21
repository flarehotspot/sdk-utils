(function (window) {
  var $flare = window.$flare;

  $flare.notification = {
    success: function (msg) {
      return createToast('success', msg);
    },
    info: function (msg) {
      return createToast('info', msg);
    },
    warning: function (msg) {
      return createToast('warning', msg);
    },
    error: function (msg) {
      return createToast('error', msg);
    }
  };

  function createToast(msgType, msg) {
    var colorSuccess = '#1fad45';
    var colorInfo = '#0581f5';
    var colorWarning = '#f2b211';
    var colorError = '#c72020';

    var color =
      msgType === 'success'
        ? colorSuccess
        : msgType === 'info'
        ? colorInfo
        : msgType === 'warning'
        ? colorWarning
        : msgType === 'error'
        ? colorError
        : 'gray';

    var t = Toastify({
      text: msg,
      duration: 5000,
      newWindow: true,
      close: false,
      gravity: 'bottom', // `top` or `bottom`
      position: 'right', // `left`, `center` or `right`
      style: { background: color },
      stopOnFocus: true, // Prevents dismissing of toast on hover,
      onClick: function () {
        if (t) {
          t.hideToast();
        }
      }
    }).showToast();

    return t;
  }
})(window);
