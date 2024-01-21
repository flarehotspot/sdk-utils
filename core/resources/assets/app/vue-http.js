(function ($flare) {
  var http = window.BasicHttp;
  var client = {};

  function invalidResponse(err) {
    console.error('Invalid response:', err);
  }

  function parseRespones(res) {
    var $res = res.$vue;
    if (!$res) {
      invalidResponse(res);
      return res;
    }

    if ($res.redirect) {
      $flare.router.push({ name: $res.name });
    } else if ($res.data) {
      return $res.data;
    } else {
      invalidResponse(res);
      return res;
    }
  }

  client.get = function (url, params) {
    return http.GetJson(url, params).then(function (data) {
      return parseRespones(data);
    });
  };

  client.post = function (url, params) {
    return http.PostJson(url, params).then(function (data) {
      return parseRespones(data);
    });
  };

  $flare.http = client;
})(window.$flare);
