<header class="main-header">
    <!-- Logo -->
    <a href="" class="logo">
        <!-- mini logo for sidebar mini 50x50 pixels -->
        <span class="logo-mini"><b>S</b>Aff</span>
        <!-- logo for regular state and mobile devices -->
        <span class="logo-lg"><b>Scale</b>Aff</span>
    </a>
    <!-- Header Navbar: style can be found in header.less -->
    <nav class="navbar navbar-static-top">
        <!-- Sidebar toggle button-->
        <a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">
            <span class="sr-only">Toggle navigation</span>
        </a>

        <div class="navbar-custom-menu">
            <ul class="nav navbar-nav">
                {{template "templates/v1/partials/notifications.tpl"}}
                {{template "templates/v1/partials/usermenu.tpl"}}
            </ul>
        </div>
    </nav>
</header>