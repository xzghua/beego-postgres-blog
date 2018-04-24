<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <title>叶落山城秋</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">

    <link rel="stylesheet" href="/static/auth/css/style.css">
    <link rel="stylesheet" href="/static/auth/css/loader-style.css">
    <link rel="stylesheet" href="/static/auth/css/bootstrap.css">

    <link rel="stylesheet" type="text/css" href="/static/auth/js/progress-bar/number-pb.css">

    <style type="text/css">
        canvas#canvas4 {
            position: relative;
            top: 20px;
        }
    </style>


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
<!-- TOP NAVBAR -->
<nav role="navigation" class="navbar navbar-static-top">
    <div class="container-fluid">
        <!-- Brand and toggle get grouped for better mobile display -->
        <div class="navbar-header">
            <button data-target="#bs-example-navbar-collapse-1" data-toggle="collapse" class="navbar-toggle" type="button">
                <span class="entypo-menu"></span>
            </button>
            <button class="navbar-toggle toggle-menu-mobile toggle-left" type="button">
                <span class="entypo-list-add"></span>
            </button>
            <div id="logo-mobile" class="visible-xs">
                <h1>WEB管理<span>v1.2</span></h1>
            </div>
        </div>

        <!-- Collect the nav links, forms, and other content for toggling -->
        <div id="bs-example-navbar-collapse-1" class="collapse navbar-collapse">
            <ul class="nav navbar-nav">
                <li><a href="#"><i data-toggle="tooltip" data-placement="bottom" title="Help" style="font-size:20px;" class="icon-help tooltitle"></i></a>
                </li>
            </ul>
            <div id="nt-title-container" class="navbar-left running-text visible-lg">
                <ul class="date-top">
                    <li class="entypo-calendar" style="margin-right:5px"></li>
                    <li id="Date"></li>
                </ul>

                <ul id="digital-clock" class="digital">
                    <li class="entypo-clock" style="margin-right:5px"></li>
                    <li class="hour"></li>
                    <li>:</li>
                    <li class="min"></li>
                    <li>:</li>
                    <li class="sec"></li>
                    <li class="meridiem"></li>
                </ul>
                <ul id="nt-title">
                    <li><i class="wi-day-lightning"></i>&#160;&#160;Berlin&#160;
                        <b>85</b><i class="wi-fahrenheit"></i>&#160;; 15km/h
                    </li>
                    <li><i class="wi-day-lightning"></i>&#160;&#160;Yogyakarta&#160;
                        <b>85</b><i class="wi-fahrenheit"></i>&#160;; Tonight- 72 °F (22.2 °C)
                    </li>

                </ul>
            </div>

            <ul style="margin-right:0;" class="nav navbar-nav navbar-right">
                <li>
                    <a data-toggle="dropdown" class="dropdown-toggle" href="#">
                        <img alt="" class="admin-pic img-circle" src="http://api.randomuser.me/portraits/thumb/men/10.jpg">Hi, Dave Mattew <b class="caret"></b>
                    </a>
                    <ul style="margin-top:14px;" role="menu" class="dropdown-setting dropdown-menu">
                        <li>
                            <a href="#">
                                <span class="entypo-user"></span>&#160;&#160;My Profile</a>
                        </li>
                        <li>
                            <a href="#">
                                <span class="entypo-vcard"></span>&#160;&#160;Account Setting</a>
                        </li>
                        <li>
                            <a href="#">
                                <span class="entypo-lifebuoy"></span>&#160;&#160;Help</a>
                        </li>
                        <li class="divider"></li>
                        <li>
                            <a href="http://themeforest.net/item/apricot-navigation-admin-dashboard-template/7664475?WT.ac=category_item&WT.z_author=themesmile">
                                <span class="entypo-basket"></span>&#160;&#160; Purchase</a>
                        </li>
                    </ul>
                </li>

                <li class="hidden-xs">
                    <a class="toggle-left" href="#">
                        <span style="font-size:20px;" class="entypo-list-add"></span>
                    </a>
                </li>
            </ul>

        </div>
        <!-- /.navbar-collapse -->
    </div>
    <!-- /.container-fluid -->
</nav>

<!-- /END OF TOP NAVBAR -->

<!-- SIDE MENU -->
<div id="skin-select">
    <div id="logo">
        <h1>IPHPT<span>v2.0</span></h1>
    </div>

    <a id="toggle">
        <span class="entypo-menu"></span>
    </a>
    <div class="dark">
        <form action="#">
                <span>
                    <input type="text" name="search" value="" class="search rounded id_search" placeholder="Search Menu..." autofocus="">
                </span>
        </form>
    </div>

    <div class="search-hover">
        <form id="demo-2">
            <input type="search" placeholder="Search Menu..." class="id_search">
        </form>
    </div>

    <div class="skin-part">
        <div id="tree-wrap">
            <div class="side-bar">
                <ul class="topnav menu-left-nest">
                    <li>
                        <a href="#" style="border-left:0px solid!important;" class="title-menu-left">

                            <span class="design-kit"></span>
                            <i data-toggle="tooltip" class="entypo-cog pull-right config-wrap"></i>

                        </a>
                    </li>

                    <li>
                        <a class="tooltip-tip ajax-load" href="index.html" title="Dashboard">
                            <i class="icon-window"></i>
                            <span>Dashboard</span>
                        </a>
                    </li>

                    <li>
                        <a class="tooltip-tip" href="#" title="Extra Pages">
                            <i class="icon-document-new"></i>
                            <span>文章</span>
                        </a>
                        <ul>
                            <li>
                                <a class="tooltip-tip2 ajax-load" href="/console/post" title="Blank Page"><i class="icon-media-record"></i><span>文章列表</span></a>
                            </li>
                            <li>
                                <a class="tooltip-tip2 ajax-load" href="/console/post/create" title="Profile Page"><i class="icon-user"></i><span>新建文章</span></a>
                            </li>
                        </ul>
                    </li>

                </ul>

                <ul id="menu-showhide" class="topnav menu-left-nest">
                    <li>
                        <a href="#" style="border-left:0px solid!important;" class="title-menu-left">

                            <span class="component"></span>
                            <i data-toggle="tooltip" class="entypo-cog pull-right config-wrap"></i>

                        </a>
                    </li>


                    <li>
                        <a class="tooltip-tip" href="#" title="UI Element">
                            <i class="icon-monitor"></i>
                            <span>分类管理</span>
                        </a>
                        <ul>
                            <li>
                                <a class="tooltip-tip2 ajax-load" href="/console/cate" title="Element"><i class="icon-attachment"></i><span>分类列表</span></a>
                            </li>
                            <li><a class="tooltip-tip2 ajax-load" href="/console/cate/create" title="Button"><i class="icon-view-list-large"></i><span>分类创建</span> </a>
                            </li>
                        </ul>
                    </li>
                    <li>
                        <a class="tooltip-tip" href="#" title="Form">
                            <i class="icon-document"></i>
                            <span>标签管理</span>
                        </a>
                        <ul>
                            <li>
                                <a class="tooltip-tip2 ajax-load" href="/console/tag" title="Form Elements"><i class="icon-document-edit"></i><span>标签列表</span></a>
                            </li>
                        </ul>
                    </li>

                </ul>

                <ul id="menu-showhide" class="topnav menu-left-nest">
                    <li>
                        <a href="#" style="border-left:0px solid!important;" class="title-menu-left">

                            <span class="widget-menu"></span>
                            <i data-toggle="tooltip" class="entypo-cog pull-right config-wrap"></i>

                        </a>
                    </li>


                    <li>
                        <a class="tooltip-tip" href="#" title="UI Element">
                            <i class="icon-monitor"></i>
                            <span>友链</span>
                        </a>
                        <ul>
                            <li>
                                <a class="tooltip-tip2 ajax-load" href="/console/link" title="Element"><i class="icon-attachment"></i><span>友情链接</span></a>
                            </li>
                        </ul>
                    </li>
                    <li>
                        <a class="tooltip-tip" href="#" title="Form">
                            <i class="icon-document"></i>
                            <span>系统管理</span>
                        </a>
                        <ul>
                            <li>
                                <a class="tooltip-tip2 ajax-load" href="/console/system" title="Form Elements"><i class="icon-document-edit"></i><span>系统管理</span></a>
                            </li>
                        </ul>
                    </li>

                </ul>

                <div class="side-dash">
                    <h3>
                        <span>Device</span>
                    </h3>
                    <ul class="side-dashh-list">
                        <li>Avg. Traffic
                            <span>25k<i style="color:#44BBC1;" class="fa fa-arrow-circle-up"></i>
                                </span>
                        </li>
                        <li>Visitors
                            <span>80%<i style="color:#AB6DB0;" class="fa fa-arrow-circle-down"></i>
                                </span>
                        </li>
                        <li>Convertion Rate
                            <span>13m<i style="color:#19A1F9;" class="fa fa-arrow-circle-up"></i>
                                </span>
                        </li>
                    </ul>
                    <h3>
                        <span>Traffic</span>
                    </h3>

                    <div id="g1" style="height:180px" class="gauge"></div>
                </div>
            </div>
        </div>
    </div>
</div>
<!-- END OF SIDE MENU -->

<!--  PAPER WRAP -->
<div class="wrap-fluid">
    <div class="container-fluid paper-wrap bevel tlbr">

        <!-- CONTENT -->
        <!--TITLE -->
        <div class="row">
            <div id="paper-top">
                <div class="col-lg-3">
                    <h2 class="tittle-content-header">
                        <i class="icon-window"></i>
                        <span>Dashboard
                            </span>
                    </h2>

                </div>

                <div class="col-lg-7">
                    <div class="devider-vertical visible-lg"></div>
                    <div class="tittle-middle-header">

                        <div class="alert">
                            <button type="button" class="close" data-dismiss="alert">×</button>
                            <span class="tittle-alert entypo-info-circled"></span>
                            Welcome back,&nbsp;
                            <strong>Dave mattew!</strong>&nbsp;&nbsp;Your last sig in at Yesterday, 16:54 PM
                        </div>

                    </div>

                </div>
                <div class="col-lg-2">
                    <div class="devider-vertical visible-lg"></div>
                    <div class="btn-group btn-wigdet pull-right visible-lg">
                        <div class="btn">
                            Widget</div>
                        <button data-toggle="dropdown" class="btn dropdown-toggle" type="button">
                            <span class="caret"></span>
                            <span class="sr-only">Toggle Dropdown</span>
                        </button>
                        <ul role="menu" class="dropdown-menu">
                            <li>
                                <a href="#">
                                    <span class="entypo-plus-circled margin-iconic"></span>Add New</a>
                            </li>
                            <li>
                                <a href="#">
                                    <span class="entypo-heart margin-iconic"></span>Favorite</a>
                            </li>
                            <li>
                                <a href="#">
                                    <span class="entypo-cog margin-iconic"></span>Setting</a>
                            </li>
                        </ul>
                    </div>


                </div>
            </div>
        </div>
        <!--/ TITLE -->
        <div>
            {{.LayoutContent }}

        </div>

    </div>
</div>
<!--  END OF PAPER WRAP -->
<script type="text/javascript" src="/static/auth/js/jquery.js"></script>

<script src="/static/auth/js/progress-bar/src/jquery.velocity.min.js"></script>
<script src="/static/auth/js/progress-bar/number-pb.js"></script>
<script src="/static/auth/js/progress-bar/progress-app.js"></script>



<!-- MAIN EFFECT -->
<script type="text/javascript" src="/static/auth/js/preloader.js"></script>
<script type="text/javascript" src="/static/auth/js/bootstrap.js"></script>
<script type="text/javascript" src="/static/auth/js/app.js"></script>
<script type="text/javascript" src="/static/auth/js/load.js"></script>
<script type="text/javascript" src="/static/auth/js/main.js"></script>



<script type="text/javascript" src="/static/auth/js/validate/jquery.validate.min.js"></script>
<script type="text/javascript" src="/static/auth/js/tag/jquery.tagsinput.js"></script>
<script type="text/javascript" src="/static/auth/js/jquery-ui.min.js"></script>
<script type="text/javascript" src="/static/auth/js/validate/jquery.validate.min.js"></script>

<script type="text/javascript" src="/static/auth/js/admin.create.js"></script>
<script src="/static/auth/js/editormd.min.js"></script>


<script type="text/javascript" src="/static/auth/js/pnotify/jquery.pnotify.min.js"></script>

<script>

    var testEditor;

    function getCookie(cname)
    {
        var name = cname + "=";
        var ca = document.cookie.split(';');
        for(var i=0; i<ca.length; i++)
        {
            var c = ca[i].trim();
            if (c.indexOf(name)==0) return c.substring(name.length,c.length);
        }
        return "";
    }
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



    $(document).ready(function() {
        testEditor = editormd("editormd", {
            width: "100%",
            height: 640,
            syncScrolling: "single",
            path: "/static/auth/js/lib/",
            imageUpload : true,
            imageFormats : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
            imageUploadURL : "/uploadPhotos",
            toolbarAutoFixed:false,
            // markdown : md,
            codeFold : true,

            saveHTMLToTextarea : true,    // 保存 HTML 到 Textarea
            searchReplace : true,
            //watch : false,                // 关闭实时预览
            htmlDecode : "style,script,iframe|on*",            // 开启 HTML 标签解析，为了安全性，默认不开启

            emoji : true,
            taskList : true,

            theme : "default | dark",
            // Preview container theme, added v1.5.0
            // You can also custom css class .editormd-preview-theme-xxxx
            previewTheme : " dark",
            // Added @v1.5.0 & after version this is CodeMirror (editor area) theme
            editorTheme : editormd.editorThemes['editormd']
        });
    });

    $(function() {

        $('.tag').tagsInput({

            autocomplete_url:"/console/tag/auto",

            previewTheme : "dark",
            editorTheme : "pastel-on-dark",
            markdown : 'md',
            codeFold : true,

            width: 'auto',
            height:'10px',
            defaultText:'添加标签',
            minChars:'0',
            maxChars:'20',
            placeholderColor:'#ff0066'

        });

    });
</script>




</body>

</html>
