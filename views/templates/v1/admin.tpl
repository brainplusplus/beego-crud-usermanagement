<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Dashboard</title>
    <!-- Tell the browser to be responsive to screen width -->
    <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
    <!-- Bootstrap 3.3.7 -->
    <link rel="stylesheet" href="/static/bower_components/bootstrap/dist/css/bootstrap.min.css">
    <!-- Font Awesome -->
    <link rel="stylesheet" href="/static/bower_components/font-awesome/css/font-awesome.min.css">
    <!-- Ionicons -->
    <link rel="stylesheet" href="/static/bower_components/Ionicons/css/ionicons.min.css">
    <!-- DataTables -->
    <link rel="stylesheet" href="/static/bower_components/datatables.net-bs/css/dataTables.bootstrap.min.css">
    <!-- iCheck for checkboxes and radio inputs -->
    <link rel="stylesheet" href="/static/plugins/iCheck/all.css">
    <!-- Theme style -->
    <link rel="stylesheet" href="/static/dist/css/admin.min.css">
    <link rel="stylesheet" href="/static/dist/css/skins/_all-skins.css">
    <!-- Morris chart -->
    <link rel="stylesheet" href="/static/bower_components/morris.js/morris.css">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="/static/https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="/static/https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->

    <!-- Google Font -->
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,600,700,300italic,400italic,600italic">

    <style type="text/css">
        .bg-white {
            background: #fff;
            box-shadow: 0px 6px 63px -2px rgba(158, 168, 181, 0.57);
            margin-bottom: 30px;
            color: #7171A6;
        }
        .bg-white:hover {
            color: #7171A6;
        }

        .bg-white:hover > .small-box-footer {
            color: rgb(236, 240, 245);
            background: #8f8cc1;
        }
        .bg-white > .small-box-footer {
            background: rgb(236, 240, 245);
            color: #8f8cc1;
        }

        .bg-white .icon {
            color: #f3f2ff;
        }
    </style>
    {{ block "head" . }}{{ end }}
</head>
<body class="hold-transition skin-purple sidebar-mini fixed">
<div class="wrapper">

    {{template "templates/v1/partials/header.tpl"}}
    {{template "templates/v1/partials/sidebar.tpl"}}

    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        {{ block "content" . }}{{ end }}
    </div>
    <!-- /.content-wrapper -->
    <footer class="main-footer">
        <strong>Copyright &copy; 2018  All rights reserved.</strong>
    </footer>

</div>
<!-- ./wrapper -->
{{template "templates/v1/partials/modal.tpl"}}
<!-- jQuery 3 -->
<script src="/static/bower_components/jquery/dist/jquery.min.js"></script>
<!-- jQuery UI 1.11.4 -->
<script src="/static/bower_components/jquery-ui/jquery-ui.min.js"></script>
<!-- Resolve conflict in jQuery UI tooltip with Bootstrap tooltip -->
<script>
    $.widget.bridge('uibutton', $.ui.button);
</script>
<!-- Bootstrap 3.3.7 -->
<script src="/static/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
<!-- Morris.js charts -->
<script src="/static/bower_components/raphael/raphael.min.js"></script>
<script src="/static/bower_components/morris.js/morris.min.js"></script>

<script src="/static/bower_components/datatables.net/js/jquery.dataTables.min.js"></script>
<script src="/static/bower_components/datatables.net-bs/js/dataTables.bootstrap.min.js"></script>
<!-- jQuery Knob Chart -->
<script src="/static/bower_components/jquery-knob/dist/jquery.knob.min.js"></script>
<!-- Slimscroll -->
<script src="/static/bower_components/jquery-slimscroll/jquery.slimscroll.min.js"></script>
<script src="/static/plugins/iCheck/icheck.min.js"></script>
<script src="/static/plugins/blockUI/jquery.blockUI.js"></script>
<script src="/static/plugins/parsley/parsley.min.js"></script>
<script src="/static/plugins/jqueryForm/jquery.form.min.js"></script>
<script src="/static/plugins/axios/axios.min.js"></script>
<script src="/static/bower_components/sweetalert2/dist/sweetalert2.all.js"></script>

<!-- AdminLTE App -->
<script src="/static/dist/js/admin.min.js"></script>


<script>
    var WEB = 'http://localhost:8080'
    var API = 'http://localhost:8080/api'
    var NOT_USING_AUTH_HEADER = false
</script>
<script src="/static/js/core.js"></script>

<!-- AdminLTE dashboard demo (This is only for demo purposes) -->
{{block "footer" .}}{{end}}

</body>
</html>
