<!DOCTYPE html>
<head>
    <title>{{.system.Title}}</title>
    <!-- Meta data -->
    <meta http-equiv="Content-Type" content="text/html" charset="UTF-8" >
    <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta name="generator" content="{{.system.Title}}">
    <meta name="author" content="叶落山城秋">
    <meta name="description" content="{{ .system.Title }}" />
    <meta name="keywords" content="{{ .system.SeoKey }}" />
    <!-- Favicon, (keep icon in root folder) -->
    <link rel="Shortcut Icon" href="/static/index/img/favicon.ico" type="image/ico">
    <link rel="alternate" href="atom.xml" title="KIERAN&#39;S BLOG" type="application/atom+xml">
    <link rel="stylesheet" href="/static/index/css/all.css" media="screen" type="text/css">
    <link rel="stylesheet" href="/static/index/highlightjs/xcode.css" type="text/css">
    <!--[if IE 8]>
    <link rel="stylesheet" type="text/css" href="/static/index/css/ie8.css" />
    <![endif]-->
    <!-- jQuery | Load our jQuery, with an alternative source fallback to a local version if request is unavailable -->
    <script src="/static/index/js/jquery-1.11.1.min.js"></script>
    <script>window.jQuery || document.write('<script src="/static/index/js/jquery-1.11.1.min.js"><\/script>')</script>
    <!-- Load these in the <head> for quicker IE8+ load times -->
    <!-- HTML5 shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
    <script src="/static/index/js/html5shiv.min.js"></script>
    <script src="/static/index/js/respond.min.js"></script>
    <![endif]-->

    <link rel="alternate" type="application/atom+xml" title="Atom 0.3" href="atom.xml">

    <style>
        .col-md-8.col-md-offset-2.opening-statement img{display:none;}
    </style>
</head>

<!--
<body class="home-template">
-->
<body id="index" class="lightnav animsition">

<!-- ============================ Off-canvas navigation =========================== -->

<div class="sb-slidebar sb-right sb-style-overlay sb-momentum-scrolling">
    <div class="sb-close" aria-label="Close Menu" aria-hidden="true">
        <img src="/static/index/img/close.png" alt="Close"/>
    </div>
    <!-- Lists in Slidebars -->
    <ul class="sb-menu">
        <li><a href="/" class="animsition-link" title="Home">Home</a></li>
        {{/*<li><a href="archives" class="animsition-link" title="archive">archives</a></li>*/}}
        <!-- Dropdown Menu -->

        <li>
            <a class="sb-toggle-submenu">TAGS<span class="sb-caret"></span></a>
            <ul class="sb-submenu">
                {{ range $key,$value :=  .tag }}
                <li>
                    <a href="/tags/{{ map_get $value "Name"}}" class="animsition-link">{{ map_get $value "Name"}}<small>({{ map_get $value "TagNum"}})</small></a>
                </li>
                {{end}}
            </ul>
        </li>


        <li>
            <a class="sb-toggle-submenu">CATEGORIES<span class="sb-caret"></span></a>
            <ul class="sb-submenu">
                {{ range $key,$value :=  .cate }}
                    <li>
                        <a href="/categories/{{ map_get $value "DisplayName"}}" class="animsition-link">{{str2html $value.html}}{{$value.DisplayName}}</a>
                    </li>
                {{end}}
            </ul>
        </li>


        <li>
            <a class="sb-toggle-submenu">Links<span class="sb-caret"></span></a>
            <ul class="sb-submenu">
                {{ range $key,$value :=  .link }}
                    <li>
                        <a href="{{ map_get $value "Link"}}" class="animsition-link">{{ map_get $value "Name"}}</a>
                    </li>
                {{end}}

            </ul>
        </li>

    </ul>
    <!-- Lists in Slidebars -->
    <ul class="sb-menu secondary">
        <li><a href="/archive" class="animsition-link" title="Archive">Archive</a></li>
        <li><a href="about.html" class="animsition-link" title="about">About</a></li>
        <li><a href="atom.xml" class="animsition-link" title="rss">RSS</a></li>
    </ul>
</div>

<!-- ============================ END Off-canvas navigation =========================== -->

<!-- ============================ #sb-site Main Page Wrapper =========================== -->

<div id="sb-site">
    <!-- #sb-site - All page content should be contained within this id, except the off-canvas navigation itself -->

    <!-- ============================ Header & Logo bar =========================== -->

    <div id="navigation" class="navbar navbar-fixed-top">
        <div class="navbar-inner">
            <div class="container">
                <!-- Nav logo -->
                <div class="logo">
                    <a href="/" title="Logo" class="animsition-link">
                        <img src="/static/index/img/logo.png" alt="Logo" width="35px;"/>
                    </a>
                </div>
                <!-- // Nav logo -->
                <!-- Info-bar -->
                <nav>
                    <ul class="nav">
                        <li><a href="/" class="animsition-link">{{.system.Title}}</a></li>
                        <li class="nolink"><span>Always </span>Creative.</li>

                        <li><a href="https://github.com/SuperKieran" title="Github" target="_blank"><i class="icon-github"></i></a></li>


                        <li><a href="https://twitter.com/1tsKieran" title="Twitter" target="_blank"><i class="icon-twitter"></i></a></li>




                        <li><a href="http://weibo.com/taokailun" title="Sina-Weibo" target="_blank"><i class="icon-sina-weibo"></i></a></li>

                        <li class="nolink"><span>Welcome!</span></li>
                    </ul>
                </nav>
                <!--// Info-bar -->
            </div>
            <!-- // .container -->
            <div class="learnmore sb-toggle-right">More</div>
            <button type="button" class="navbar-toggle menu-icon sb-toggle-right" title="More">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar before"></span>
                <span class="icon-bar main"></span>
                <span class="icon-bar after"></span>
            </button>
        </div>
        <!-- // .navbar-inner -->
    </div>

    <!-- ============================ Header & Logo bar =========================== -->


    <!-- ============================ Hero Image =========================== -->
{{.LayoutContent }}

    <section id="statement">
        <div class="container text-center wow fadeInUp" data-wow-delay="0.5s">
            <div class="row">
                <p>每一个不曾起舞的日子都是对生命的辜负。</p>
            </div>
        </div>
    </section>
    <!-- ============================ END Content =========================== -->


    <!-- ============================ Footer =========================== -->

    <footer>
        <div class="container">
            <div class="copy">
                <p>
                    &copy; 2014<script>new Date().getFullYear()>2010&&document.write("-"+new Date().getFullYear());</script>, Content By Kieran. All Rights Reserved.
                </p>
                <p>Theme By <a href="index.html" style="color: #767D84">Kieran</a></p>
            </div>
            <div class="social">
                <ul>

                    <li><a href="https://github.com/SuperKieran" title="Github" target="_blank"><i class="icon-github"></i></a>&nbsp;</li>


                    <li><a href="https://twitter.com/1tsKieran" title="Twitter" target="_blank"><i class="icon-twitter"></i></a>&nbsp;</li>




                    <li><a href="http://weibo.com/taokailun" title="Sina-Weibo" target="_blank"><i class="icon-sina-weibo"></i></a>&nbsp;</li>

                </ul>
            </div>
            <div class="clearfix"> </div>
        </div>
    </footer>
</div>
    <!-- ============================ END Footer =========================== -->
    <!-- Load our scripts -->
    <!-- Resizable 'on-demand' full-height hero -->
    <script type="text/javascript">
        var resizeHero = function () {
            var hero = $(".cover,.heightblock"),
                    window1 = $(window);
            hero.css({
                "height": window1.height()
            });
        };

        resizeHero();

        $(window).resize(function () {
            resizeHero();
        });
    </script>
    <script src="/static/index/js/plugins.min.js"></script><!-- Bootstrap core and concatenated plugins always load here -->
    <script src="/static/index/js/scripts.js"></script><!-- Theme scripts -->


    <link rel="stylesheet" href="/static/index/fancybox/jquery.fancybox.css" media="screen" type="text/css">
    <script src="/static/index/fancybox/jquery.fancybox.pack.js"></script>
    <script type="text/javascript">
        $('#intro').find('img').each(function(){
            var alt = this.alt;
            if (alt){
                $(this).after('<span class="caption" style="display:none">' + alt + '</span>');
            }
            $(this).wrap('<a href="' + this.src + '" title="' + alt + '" class="fancybox" rel="gallery" />');
        });
        (function($){
            $('.fancybox').fancybox();
        })(jQuery);
    </script>

    <!-- Initiate flexslider plugin -->
    <script type="text/javascript">

    </script>

</body>
</html>
