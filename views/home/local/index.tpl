<section id="hero" class="scrollme">
    <div class="container-fluid element-img" style="background: url(/static/index/img/index.jpg) no-repeat center center fixed;background-size: cover">
        <div class="row">
            <div class="col-xs-12 col-sm-8 col-sm-offset-2 col-md-8 col-md-offset-2 vertical-align cover boost text-center">
                <div class="center-me animateme" data-when="exit" data-from="0" data-to="0.6" data-opacity="0" data-translatey="100">
                    <div>

                        <h2 style="color: #e0f507;">
                            <a href="#intro" class="more scrolly" style="color: #e0f507;">
                                命定的局限尽可永在，不屈的挑战却不可须臾或缺！
                            </a>
                        </h2>

                        <h2></h2>
                        <p></p>


                    </div>
                </div>
            </div>
            <!-- // .col-md-12 -->
        </div>
        <div class="herofade beige-dk"></div>
    </div>
</section>

<!-- Height spacing helper -->
<div class="heightblock"></div>
<!-- // End height spacing helper -->

<!-- ============================ END Hero Image =========================== -->
<!-- ============================ Content =========================== -->

<section id="intro">
    <div class="container">

        {{range $index,$item := .post }}

        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3 style="white-space: pre-wrap;word-wrap: break-word;"><a href="post/{{$item.Id}}">{{$item.Title}}</a></h3>
                    <span>
                        <span class="post-meta">
                        <time datetime="{{$item.CreatedAt}}" itemprop="datePublished">
                        {{$item.CreatedAt}}
                        </time>
                        |
                            {{range $key,$value := $item.tag_list}}
                                <a href="/tags/{{ map_get $value "name"}}">
                                    {{/*绿色*/}}
                                        {{/*rgba(247,186,41,.1) #f7ba2a 黄色*/}}
                                    <span style="border: 1px solid rgba(18,206,102,.2);background-color: rgba(18,206,102,.2);color: #13ce66;padding:2px 8px;border-radius: 4px;font-size: 12px;line-height: 22px;box-sizing: border-box; "> {{ map_get $value "name"}}</span>
                                </a>
                            {{end}}
                        </span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>
                    {{$item.Abstract}} ...
                    </p>

                    <p class="pull-right readMore">
                        <a href="/detail/{{$item.Id }}">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>
        {{end}}

        <nav class="pagination" role="pagination">
            {{ if ne .currentPage   1 }}
                <a class="pull-left" href="?page={{.lastPage}}">← Newer Posts</a>
            {{end}}
            {{ if ne .currentPage  .nextPage }}
                <a class="pull-right" href="?page={{.nextPage}}">Older Posts →</a>
            {{end}}
        </nav>
    </div>
</section>