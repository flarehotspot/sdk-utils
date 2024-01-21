(function ($flare) {
  var http = window.BasicHttp;
  var vueHttp = {};
  var rootres = '$$response';
  $flare.http = vueHttp;

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

  function invalidResponse(res) {
    var err = res.error || res.message || 'Invalid server response';
    $flare.notification.error(err);
  }

  function parseRespones(res) {
    var $res = res[rootres];
    if (!$res) {
      invalidResponse(res);
      return res;
    }

    if ($res.flash) {
      var toastFn = $flare.notification[$res.flash.type];
      if (toastFn) {
        toastFn($res.flash.msg);
      } else {
        console.error('Invalid flash type:', $res.flash.type);
        $flare.notification.info($res.flash.msg);
      }
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
})(window.$flare);
