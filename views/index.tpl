<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="/static/components/bootstrap/dist/css/bootstrap.min.css.map">
    <link rel="stylesheet" href="https://cdn.datatables.net/1.10.13/css/jquery.dataTables.min.css">
    <link rel="stylesheet" href="/static/css/styles.css">
    <script type="text/javascript" src="/static/components/jquery/dist/jquery.js"></script>
    <script type="text/javascript" src="/static/components/bootstrap/dist/js/bootstrap.js"></script>
    <script type="text/javascript" src="https://cdn.datatables.net/1.10.13/js/jquery.dataTables.min.js"></script>
    <script type="text/javascript" src="/static/js/app.js"></script>
</head>
<body>
<div id="top" class="form-group">
    <form action="?" method="post" id="mkdir"/>
    <label for=dirname>Create New Folder</label><input id=dirname type=text name=name value="" class="form-control"/>
    <input type="submit" value="create" class="btn btn-primary"/>
    </form>
    <div id="file_drop_target">
        Drag Files Here To Upload
        <b>or</b>
        <input type="file" multiple class="form-control-file"/>
    </div>
    <div id="breadcrumb">&nbsp;</div>
</div>

<div id="upload_progress"></div>
<table id="table" class="display" cellspacing="0" width="100%">
    <thead>
    <tr>
        <th>Name</th>
        <th>Size</th>
        <th>Modified</th>
        <th>Permissions</th>
        <th>Actions</th>
    </tr>
    </thead>
</table>
<footer>Simple Go(lang) File Manager by <a href="https://github.com/mikitu">mikitu</a></footer>
</body>
</html>