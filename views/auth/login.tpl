<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>login - Auth</title>
    <!-- Tell the browser to be responsive to screen width -->
    <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">

    <!-- Bootstrap 3.3.7 -->
    <link rel="stylesheet" href="/static/bower_components/bootstrap/dist/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/dist/css/auth.css">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->

    <!-- Google Font -->
    <link href="https://fonts.googleapis.com/css?family=Muli:400,600" rel="stylesheet">
</head>
<body>
<section class="section-auth">
    <h1>Sign in to ScaleAff</h1>
    <p></p>
    <div class="box">
        <form id="login_form" action="/api/auth/login" method="post">
            <div class="form-group">
                <label for="exampleInputEmail1">Email address</label>
                <input type="email" class="form-control" id="Username" name="Username" placeholder="Enter email" required>
            </div>
            <div class="form-group">
                <label for="exampleInputPassword1">Password</label>
                <input type="password" class="form-control" id="Password" name="Password" placeholder="Password" required>
            </div>
            <div class="form-group">
                <button type="submit" class="btn btn-ui">Login</button>
            </div>
            <div class="form-group pass">
                <a href="forgot.html">Forgot password?</a>
            </div>
        </form>
    </div>
    <p>Need an account? <a href="register.html"> Sign up</a></p>
</section>
<script>
    var WEB = 'http://localhost:8080'
    var API = 'http://localhost:8080/api'
    var NOT_USING_AUTH_HEADER = false
</script>
<script src="/static/bower_components/jquery/dist/jquery.min.js"></script>
<script src="/static/bower_components/jquery-ui/jquery-ui.min.js"></script>
<script src="/static/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
<script src="/static/plugins/blockUI/jquery.blockUI.js"></script>
<script src="/static/plugins/parsley/parsley.min.js"></script>
<script src="/static/plugins/jqueryForm/jquery.form.min.js"></script>
<script src="/static/plugins/axios/axios.min.js"></script>
<script src="/static/bower_components/sweetalert2/dist/sweetalert2.all.js"></script>
<script src="/static/js/core.js"></script>
<script src="/static/js/modules/auth/login.js"></script>
</body>
</html>