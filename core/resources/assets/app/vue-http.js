(function ($flare) {
  var http = window.BasicHttp;
  var vueHttp = {};
  var rootres = '$$response';
  $flare.http = vueHttp;

  function invalidResponse(err) {
    console.error('Invalid response:', err);
  }

  function parseRespones(res) {
    var $res = res[rootres];
    if (!$res) {
      invalidResponse(res);
      return res;
    }

    if ($res.flash) {
      var f = $res.flash;
      // types are success, info, warning, error
      var colorSuccess = '#1fad45';
      var colorInfo = '#0581f5';
      var colorWarning = '#f2b211';
      var colorError = '#c72020';

      var color =
        f.type === 'success'
          ? colorSuccess
          : f.type === 'info'
          ? colorInfo
          : f.type === 'warning'
          ? colorWarning
          : f.type === 'error'
          ? colorError
          : 'gray';

      var t = Toastify({
        text: f.msg,
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
    }

    if ($res.redirect) {
      $flare.router.push({ name: $res.route_name });
    } else if ($res.data) {
      return $res.data;
    } else {
      invalidResponse(res);
      return res;
    }
  }

  vueHttp.get = function (url, params) {
    return http.GetJson(url, params).then(function (data) {
      return parseRespones(data);
    });
  };

  vueHttp.post = function (url, params) {
    return http.PostJson(url, params).then(function (data) {
      return parseRespones(data);
    });
  };
})(window.$flare);
