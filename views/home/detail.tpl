
<section id="intro">
    <div class="container">
        <div class="row col-md-offset-2">
            <div class="col-md-8">
    			<span class="post-meta">
      <time datetime="{{.detail.CreatedAt}}" itemprop="datePublished">
          {{date .detail.CreatedAt "Y-m-d H:i:s"}}
      </time>
    |
                    {{range $k,$v := .postTag}}
                        <a href='../../tag/{{$v.id}}'>{{$v.name}}</a>,
                    {{end}}


</span>
                <h1>{{ .detail.Title}}</h1>
            </div>
        </div>
        <div class="col-md-8 col-md-offset-2">
            {{str2html .detail.Content}}
            <div class="clearfix"></div>
            <hr class="nogutter">
        </div>
        <nav class="m-pagination col-md-8 col-md-offset-2 col-sm-24" role="pagination">

            {{if .lastBeforeCond }}
                <a class="pull-left" href="{{.lastBefore.Id}}" style="float: left;max-width:35%;">
                    ←
                    <span style="overflow: hidden;text-overflow:ellipsis;white-space: nowrap;"> {{.lastBefore.Title}}</span>
                </a>
            {{end}}

            {{if  .lastPostCond }}
            <a class="pull-right" href="{{.lastPost.Id}}" style="max-width:35%;float: right;">
                <span style="overflow: hidden;text-overflow:ellipsis;white-space: nowrap;">{{.lastPost.Title}}</span>
                →
            </a>
            {{end}}
        </nav>

        <div class="col-md-8 col-md-offset-2 col-sm-24"><script type="text/javascript">
            document.write('<div id="SOHUCS" sid="' + window.location.pathname.slice(1) + '" ></div>');

            window.onload = function () {
                (function(){
                    var appid = 'cyt8DBIR5';
                    var conf = 'prod_9c1a4222cd636041ed65a5d58fc0d46f';
                    var width = window.innerWidth || document.documentElement.clientWidth;
                    var loadJs=function(d,a,id){
                        var c=document.getElementsByTagName("head")[0]||document.head||document.documentElement;
                        var b=document.createElement("script");
                        b.setAttribute("type","text/javascript");
                        b.setAttribute("charset","UTF-8");
                        b.setAttribute("src",d);
                        if (id) {
                            b.setAttribute("id", id);
                        }
                        if(typeof a==="function"){
                            if(window.attachEvent){
                                b.onreadystatechange=function(){
                                    var e=b.readyState;
                                    if(e==="loaded"||e==="complete"){
                                        b.onreadystatechange=null;a()
                                    }
                                }
                            }else{
                                b.onload=a
                            }
                        }
                        c.appendChild(b)
                    };

                    // if (width < 960) {
                    // loadJs('https://changyan.sohu.com/upload/mobile/wap-js/changyan_mobile.js?client_id=' + appid + '&conf=' + conf, null, 'changyan_mobile_js');
                    // } else {
                    loadJs("https://changyan.sohu.com/upload/changyan.js",function(){window.changyan.api.config({appid:appid,conf:conf})});
                    // }
                })();

                var _hmt = _hmt || [];
                (function() {
                    var hm = document.createElement("script");
                    hm.src = "//hm.baidu.com/hm.js?c71a239d451f8149ac6ebf67a0c992fa";
                    var s = document.getElementsByTagName("script")[0];
                    s.parentNode.insertBefore(hm, s);
                })();

                (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
                    (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
                        m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
                })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
                ga('create', 'UA-63685503-1', 'auto');
                ga('send', 'pageview');
            }
        </script>
        </div>
    </div>
</section>