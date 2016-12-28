/**
 * Created by mihaibucse on 22/12/2016.
 */
var table;
var baseUrl = "/list?file="
$(document).ready(function () {
  var hashval = window.location.hash.substr(1);
  $('#breadcrumb').empty().html(renderBreadcrumbs(hashval));
  table = $('#table').DataTable({
    "rowCallback": function (nRow, aData, iDisplayIndex) {
      $(nRow).addClass(aData.is_dir ? 'is_dir' : '');
    },
    "ajax": {
      "url": baseUrl + hashval,
      "dataSrc": function (json) {
        $.each(json.data, function (idx, data) {
          var $link = $('<a class="name" />').text(data.name);
          if (data.is_dir) {
            $link = $('<a class="name" />')
                .attr('href', '#' + data.path)
                .text(data.name);
          }
          var $dl_link = $('<a/>').attr('href', '/download?file=' + encodeURIComponent(data.path))
              .addClass('download').text('download');
          var $delete_link = $('<a href="#" />').attr('data-file', data.path).addClass('delete').text('delete');
          var perms = [];
          if (data.is_readable) perms.push('read');
          if (data.is_writable) perms.push('write');
          if (data.is_executable) perms.push('exec');
          $size = $('<span class="size" />').text(formatFileSize(data.size))
          data.name = $('<div/>').append($link).html();
          data.sizef = $('<div/>').append($size).html()
          data.fmtime = formatTimestamp(data.mtime);
          data.perm = perms.join('+');
          data.actions = $('<div/>').append($dl_link).append(data.is_deleteable ? $delete_link : '').html()
        })
        return json.data
      }
    },
    "aaSorting": [[4, 'desc']],
    "responsive": true,
    "columnDefs": [ {
        "className": 'actions',
        "orderable": false,
        "targets":   6
      },
      { responsivePriority: 1, targets: 0 },
      { responsivePriority: 2, targets: -1 },
    ],
    "columns": [
      {"data": "name"},
      {"data": "size", "visible": false},
      {"data": "sizef", "orderData": 1, "targets": 1},
      {"data": "mtime", "visible": false},
      {"data": "fmtime", "orderData": 3, "targets": 3},
      {"data": "perm", "orderable": false},
      {"data": "actions"}
    ]
  });
  $(window).bind('hashchange', reloadTable).trigger('hashchange');
});
function reloadTable() {
  var hashval = window.location.hash.substr(1);
  $('#breadcrumb').empty().html(renderBreadcrumbs(hashval));
  table.ajax.url(baseUrl + hashval)
  table.ajax.reload()
}
function renderBreadcrumbs(path) {
  var base = "",
      $html = $('<ol class="breadcrumb"/>').append($('<li class="breadcrumb-item"><a href="#">Home</a></li>'));
  $.each(path.split('/'), function (k, v) {
    if (v) {
      $html.append($('<li class="breadcrumb-item"/>')
          .append($('<a/>').attr('href', '#' + base + v).text(v)));
      base += v + '/';
    }
  });
  return $html;
}
function formatFileSize(bytes) {
  var s = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB'];
  for (var pos = 0; bytes >= 1000; pos++, bytes /= 1024);
  var d = Math.round(bytes * 10);
  return pos ? [parseInt(d / 10), ".", d % 10, " ", s[pos]].join('') : bytes + ' bytes';
}
function formatTimestamp(unix_timestamp) {
  var m = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
  var d = new Date(unix_timestamp * 1000);
  return [m[d.getMonth()], ' ', d.getDate(), ', ', d.getFullYear(), " ",
    (d.getHours() % 12 || 12), ":", (d.getMinutes() < 10 ? '0' : '') + d.getMinutes(),
    " ", d.getHours() >= 12 ? 'PM' : 'AM'].join('');
}
