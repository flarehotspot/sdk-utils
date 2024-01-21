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
      var fn = $flare.notification[$res.flash.type];
      if (fn) {
        fn($res.flash.msg);
      } else {
        console.error('Invalid flash type:', $res.flash.type);
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
