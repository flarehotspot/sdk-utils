$(document).ready(function() {
  var progressCon = $('#plugin-progress-con');
  var progressMsg = $('#plugin-progress-msg');
  var progressBar = $('#plugin-progress-bar');

  progressCon.hide();

  window.AdminEvents.addEventListener(
    'plugin:install:progress',
    function(evt) {
      var data = JSON.parse(evt.data);
      console.log('Install progress:', data);
      progressMsg.text(data.msg);
      if (data.done && data.err) {
        progressMsg.addClass('text-danger');
      } else {
        progressMsg.removeClass('text-danger');
      }
    }
  );

  $('#new-plugin-form').submit(function(e) {
    progressCon.show();
    progressMsg.text('Uploading zip file...');

    $.ajax({
      url:         $(this).attr('action'),
      type:        'POST',
      data:        new FormData(this),
      processData: false,
      contentType: false,
      success:     function(data) {
        console.log('Upload success:', data);
      },
      error: function(res) {
        console.log('Upload error: ', res);
      },
      xhr: function() {
        var xhr = new window.XMLHttpRequest();
        xhr.upload.addEventListener(
          'progress',
          function(evt) {
            if (evt.lengthComputable) {
              var percentComplete = (evt.loaded / evt.total) * 100;
              console.log('Upload progress: ', percentComplete);
              progressBar.attr('aria-valuenow', percentComplete);
              progressBar.css({ width: percentComplete + '%' });
            }
          },
          false
        );
        return xhr;
      }
    });
    e.preventDefault();
  });
});
