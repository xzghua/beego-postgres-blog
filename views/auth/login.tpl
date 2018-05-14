<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <title>叶落山城秋 v2.0</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">

    <!-- Le styles -->
    <script type="text/javascript" src="/static/auth/js/jquery.min.js"></script>

    <!--  <link rel="stylesheet" href="assets/css/style.css"> -->
    <link rel="stylesheet" href="/static/auth/css/loader-style.css">
    <link rel="stylesheet" href="/static/auth/css/bootstrap.css">
    <link rel="stylesheet" href="/static/auth/css/signin.css">

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
                    <form role="form" action="/console/login" method="post">
                    {{ .xsrfdata }}
                        <div class="form-group">
                            <!--a href="#" class="pull-right label-forgot">Forgot email?</a-->
                            <label for="name">用户名</label>
                            <input type="text" name="name" id="name" class="form-control">
                        </div>
                        <div class="form-group">
                            <!--a href="#" class="pull-right label-forgot">Forgot password?</a-->
                            <label for="password">密码</label>
                            <input type="password" name="password" id="password" class="form-control">
                        </div>
                        <div class="checkbox pull-left">
                            <label>
                                <input type="checkbox">记住用户名</label>
                        </div>
                        <button class="btn btn btn-primary pull-right" type="submit">
                            登 录
                        </button>
                    </form>
                    <a class="forgotLnk" href="#"></a>
                    <div class="or-box">
                        <center><span class="text-center login-with">Login or <b>Sign Up</b></span></center>
                    </div>

                    <div class="row-block">
                        <div class="row">
                            <div class="col-md-12 row-block">
                                <a href="/console/register" class="btn btn-primary btn-block">Create New Account</a>
                            </div>
                        </div>
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
