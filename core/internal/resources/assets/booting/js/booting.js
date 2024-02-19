/**
 * Copyright 2021-2022 Flarego Technologies Corp. <business@flarego.ph>
 * @file             : script.js
 * @author           : Adones Pitogo <pitogo.adones@gmail.com>
 * Date              : Nov 29, 2022
 * Last Modified Date: Nov 29, 2022
 */
window.addEventListener("load", function () {
  function checkStatus(callback) {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/", true);

    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4) {
        callback(xhr.status);
      }
    };

    xhr.send();
  }

  function redirectHome() {
    checkStatus(function (status) {
      if (status === 200) {
        window.location.href = "/";
      } else {
        setTimeout(redirectHome, 1000); // Check again after 1 second
      }
    });
  }

  function callBackHook(data) {
    document.getElementById("status-text").innerText = data.status;
  }

  var evt = new EventSource("/boot/status");

  evt.addEventListener("boot:progress", function (res) {
    var data = JSON.parse(res.data);
    console.log(data);
    callBackHook(data);
    if (data.done) {
      redirectHome();
    }
  });

  evt.onerror = function (res) {
    console.error(res);
    setTimeout(redirectHome, 1000);
  };
});
