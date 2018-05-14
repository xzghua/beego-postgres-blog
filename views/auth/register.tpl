<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <title>叶落山城秋 v2.0</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">

    <!-- Le styles -->

    <!--  <link rel="stylesheet" href="assets/css/style.css"> -->
    <link rel="stylesheet" href="/static/auth/css/loader-style.css">
    <link rel="stylesheet" href="/static/auth/css/bootstrap.css">
    <link rel="stylesheet" href="/static/auth/css/signin.css">
    <link href="/static/auth/js/validate/validate.css" rel="stylesheet">

    <!-- HTML5 shim, for IE6-8 support of HTML5 elements -->
    <!--[if lt IE 9]>
    <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->
    <!-- Fav and touch icons -->
    <link rel="shortcut icon" href="/static/auth/ico/minus.png">
</head>

<body>
<!-- Preloader -->
<div id="preloader">
    <div id="status">&nbsp;</div>
</div>

<div class="container">
    <div class="" id="login-wrapper">
        <div class="row">
            <div class="col-md-4 col-md-offset-4">
                <div id="logo-login">
                    <h1>叶落山城秋
                        <span>v2.0</span>
                    </h1>
                </div>
            </div>
        </div>

        <div class="row">
            <div class="col-md-4 col-md-offset-4">
                <div class="account-box">
                    <form role="form" action="/console/register" method="post" class="registerForm">
                    {{ .xsrfdata }}
                        <div class="form-group">
                            <!--a href="#" class="pull-right label-forgot">Forgot email?</a-->
                            <label for="name">用户名</label>
                            <input type="text" name="name" id="name" class="form-control">
                        </div>
                        <div class="form-group">
                            <!--a href="#" class="pull-right label-forgot">Forgot email?</a-->
                            <label for="email">邮箱</label>
                            <input type="text" name="email" id="email" class="form-control">
                        </div>
                        <div class="form-group">
                            <!--a href="#" class="pull-right label-forgot">Forgot password?</a-->
                            <label for="password">密码</label>
                            <input type="password" name="password" id="password" class="form-control">
                        </div>
                        <div class="form-group">
                            <!--a href="#" class="pull-right label-forgot">Forgot password?</a-->
                            <label for="rePassword">确认密码</label>
                            <input type="password" name="rePassword" id="rePassword" class="form-control">
                        </div>
                        <button class="btn btn btn-primary pull-right" type="submit">
                            注 册
                        </button>
                    </form>
                    <a class="forgotLnk" href="#"></a>
                    <div class="or-box">
                        <center><span class="text-center login-with">Login or <b><a href="/console/login">Sign Up</a></b></span></center>
                    </div>
                    <div class="row-block">
                        <div class="row">
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <p>&nbsp;</p>
    <div style="text-align:center;margin:0 auto;">
        <h6 style="color:#fff;">Copyright(C)2018 iphpt.com All Rights Reserved<br />
            叶落山城秋</h6>
    </div>

</div>
<div id="test1" class="gmap3"></div>



<!--  END OF PAPER WRAP -->




<!-- MAIN EFFECT -->
<script type="text/javascript" src="/static/auth/js/jquery.min.js"></script>
<script type="text/javascript" src="/static/auth/js/validate/jquery.validate.min.js"></script>
<script type="text/javascript" src="/static/auth/js/bootstrap.js"></script>
<script type="text/javascript" src="/static/auth/js/admin.create.js"></script>


<script>
    if ("{{.flash.error}}") {
        $.pnotify({
            title: '哎哟',
            text: "{{.flash.error}}",
            type: 'error'
        });
    } else if ("{{.flash.success}}") {
        $.pnotify({
            title: '哎哟',
            text: "{{.flash.success}}",
            type: 'success'
        });
    } else if ("{{.flash.info}}") {
        $.pnotify({
            title: '哎哟',
            text: "{{.flash.info}}",
            type: 'info'
        });
    }

</script>


</body>

</html>
