{{ template "templates/v1/admin.tpl" . }}
{{ define "content" }}
<!-- Content Header (Page header) -->
<section class="content-header">
    <h1>
        Role
        <small>Manager</small>
    </h1>
    <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> Home</a></li>
        <li class="active">Role Manager</li>
    </ol>
</section>

<!-- Main content -->
<section class="content">

    <!-- Main row -->
    <div class="row">
        <!-- Left col -->
        <section class="col-lg-12">

            <div class="box">
                <div class="box-header">
                    <div class="pull-left">
                        <h3 class="box-title">List Role</h3>
                    </div>
                    <div class="pull-right">
                        {{/*<a href="#" class="btn bg-purple" onclick="AddFormIndex()">Add Role</a>*/}}
                    </div>
                    <div class="clearfix"></div>
                    <div class="box-filter-table">
                        <div class="row">
                            <div class="form-group">
                                <div class="col-md-1" style="min-width: 140px">
                                    <label style="top:5px; position: relative;">Search</label>
                                </div>
                                <div class="col-md-11">
                                    <input type="text" id="searchTextIndex" class="form-control">
                                    </input>
                                </div>
                            </div>
                            <div class="form-group">
                                <div class="col-md-2" style="min-width: 140px">
                                    <a href="#" class="btn bg-purple" onclick="SearchTableIndex()">Search</a>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <!-- /.box-header -->
                <div class="box-body">
                    <table id="tableIndex" class="table table-bordered table-hover" cellspacing="0" width="100%"></table>
                </div>
                <!-- /.box-body -->

        </section>
        <!-- right col -->
    </div>
    <!-- /.row (main row) -->

</section>
<!-- /.content -->
{{end}}
<!-- /.content -->
{{define "footer"}}
<script src="/static/js/modules/role/index.js"></script>
<script src="/static/js/modules/role/form.js"></script>
{{end}}