
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
                        <a href='../../tags/{{$v.name}}'>{{$v.name}}</a>,
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

        <div class="col-md-8 col-md-offset-2 col-sm-24">

            {{if  eq .system.CommentType  1}}

            <div id="SOHUCS" sid="' + window.location.pathname.slice(1) + '" ></div>
                    <script type="text/javascript">
                    (function(){
                        var appid = {{.system.CyAppId}};
                        var conf = {{.system.CyAppKey}};
                        var width = window.innerWidth || document.documentElement.clientWidth;
                        if (width < 960) {
                            window.document.write('<script id="changyan_mobile_js" charset="utf-8" type="text/javascript" src="https://changyan.sohu.com/upload/mobile/wap-js/changyan_mobile.js?client_id=' + appid + '&conf=' + conf + '"><\/script>'); } else { var loadJs=function(d,a){var c=document.getElementsByTagName("head")[0]||document.head||document.documentElement;var b=document.createElement("script");b.setAttribute("type","text/javascript");b.setAttribute("charset","UTF-8");b.setAttribute("src",d);if(typeof a==="function"){if(window.attachEvent){b.onreadystatechange=function(){var e=b.readyState;if(e==="loaded"||e==="complete"){b.onreadystatechange=null;a()}}}else{b.onload=a}}c.appendChild(b)};loadJs("https://changyan.sohu.com/upload/changyan.js",function(){window.changyan.api.config({appid:appid,conf:conf})}); } })(); </script>

            {{else}}
                <div id="container"></div>
                <link rel="stylesheet" href="https://imsun.github.io/gitment/style/default.css">
                <script src="https://imsun.github.io/gitment/dist/gitment.browser.js"></script>
                <script>
                    var gitment = new Gitment({
                        id: {{.detail.Id}}, // 可选。默认为 location.href
                        owner: {{.system.GithubName}},
                        repo:  {{.system.GithubRepo}},
                        oauth: {
                            client_id: {{.system.GithubClientId}},
                            client_secret: {{.system.GithubClientSecret}},
                        },
                    })
                    gitment.render('container')
                </script>
            {{end}}


        </div>
    </div>
</section>