<link rel="stylesheet" href="/static/auth/css/bootstrap.css">
<link rel="stylesheet" href="/static/index/css/archive.css">
<section id="intro">
    <div class="container" >
        <div class="content-wrap">
            <div class="time-wrap">
                <div class="row">
                    <div class="col-md-12">
                        <ul class="timeline">
                            {{range $key,$value := .timeArr}}
                                <li class="time-label">
                                        <span class="bg-red">
                                            {{$value}}
                                        </span>
                                </li>
                                {{ range $k,$v := $.postList }}
                                        {{if eq $value $k}}
                                                {{range $k2,$v2 := $v}}
                                                    <li>
                                                        <i class="fa fa-envelope bg-time"></i>
                                                        <div class="timeline-item">
                                                            <span class="time"><i class="fa fa-clock-o"></i> {{ map_get $v2 "CreatedAt" }}</span>
                                                            <a href="#">
                                                                <h4 class="timeline-header" style="overflow: hidden;text-overflow:ellipsis;white-space: nowrap;">
                                                                    {{map_get $v2 "Title"}}
                                                                </h4>
                                                            </a>
                                                        </div>
                                                    </li>
                                                {{end}}
                                        {{end}}
                                {{end}}
                            {{end}}
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>
</section>