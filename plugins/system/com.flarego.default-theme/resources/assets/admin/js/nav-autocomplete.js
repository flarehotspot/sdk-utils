$(document).ready(function () {
  $.ajax({
    url: window.NAV_PREFETCH_URL,
    method: "GET",
    success: function (data) {
      data = data.map(function (d) {
        return {
          label: d.text,
          value: d.href
        };
      });

      var ac = new Autocomplete(document.getElementById("input-nav-search"), {
        data: data,
        treshold: 1,
        onSelectItem: function (d) {
          if (d.value) {
            window.location.replace(d.value);
          } else {
            // TODO: Sometimes d.value is null, need to fix.
            window.location.reload();
          }
        }
      });
    },
    error: function (err) {
      Notify.error("Error fetching navigation data.");
      console.log(err);
    }
  });
});
