<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="/static/components/bootstrap/dist/css/bootstrap.min.css">
    <link rel="stylesheet" href="https://cdn.datatables.net/1.10.13/css/dataTables.bootstrap.min.css">
    <link rel="stylesheet" href="/static/css/styles.css">
    <script type="text/javascript" src="/static/components/jquery/dist/jquery.js"></script>
    <script type="text/javascript" src="/static/components/bootstrap/dist/js/bootstrap.js"></script>
    <script type="text/javascript" src="https://cdn.datatables.net/1.10.13/js/jquery.dataTables.min.js"></script>
    <script type="text/javascript" src="https://cdn.datatables.net/1.10.13/js/dataTables.bootstrap.min.js"></script>

    <script type="text/javascript" src="/static/js/app.js"></script>
</head>
<body>

<div class="container-fluid">
	<div id="top" class="row">
		<div class="col-sm-12">
		<div id="file_drop_target" class="col-xs-12 col-md-8">
			Drag Files Here To Upload
			<b>or</b>
			<input type="file" multiple class="form-control-file"/>
		</div>
		<div class="col-xs-6 col-md-4">
			<form action="?" method="post" id="mkdir" class="form-inline"/>
				<div class="form-group">
					<label for=dirname>Create New Folder</label><input id=dirname type=text name=name value="" class="form-control"/>
				    <input type="submit" value="create" class="btn btn-primary"/>
				</div>
		    </form>
		</div>
		</div>
	</div>
	<div id="breadcrumb">&nbsp;</div>

	<div id="upload_progress"></div>

	<table id="table" class="table table-striped table-bordered" cellspacing="0" width="100%">
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
</div>
<footer>Simple Go(lang) File Manager by <a href="https://github.com/mikitu">mikitu</a></footer>
</body>
</html>