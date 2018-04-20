
<!-- BREADCRUMB -->
<ul id="breadcrumb">
    <li>
        <span class="entypo-home"></span>
    </li>
    <li><i class="fa fa-lg fa-angle-right"></i>
    </li>
    <li><a href="#" title="Sample page 1">分类</a>
    </li>
    <li><i class="fa fa-lg fa-angle-right"></i>
    </li>
    <li><a href="#" title="Sample page 1">分类列表</a>
    </li>
    <li class="pull-right">
        <div class="input-group input-widget">

            <input style="border-radius:15px" type="text" placeholder="Search..." class="form-control">
        </div>
    </li>
</ul>

<!-- END OF BREADCRUMB -->
<div class="content-wrap">
    <div class="row">
        <div class="col-sm-12">
            <div class="nest" id="tableStaticClose">
                <div class="title-alt">
                    <h6>
                        分类列表</h6>
                    <div class="titleClose">
                        <a class="gone" href="#tableStaticClose">
                            <span class="entypo-cancel"></span>
                        </a>
                    </div>
                    <div class="titleToggle">
                        <a class="nav-toggle-alt" href="#tableStatic">
                            <span class="entypo-up-open"></span>
                        </a>
                    </div>

                </div>

                <div class="body-nest" id="tableStatic">

                    <section id="flip-scroll">
                        <div class="table-responsive">
                            <table class="table  table-hover  table-condensed table-striped ">
                                <thead class="cf">
                                <tr>
                                    <th>ID</th>
                                    <th>分类</th>
                                    <th>创建时间</th>
                                    <th>操作</th>
                                </tr>
                                </thead>
                                <tbody>
                                {{range $index,$item := .cate }}
                                <tr>
                                    <td>{{$item.Id}}</td>
                                    <td>{{str2html $item.html}}{{$item.DisplayName}}</td>
                                    <td>{{$item.CreatedAt}}</td>
                                    <td>
                                        <form action="/console/cate/{{$item.Id}}" method="post">
                                        {{ $.xsrfdata }}
                                            <a href="/console/cate/{{$item.Id}}/edit">
                                                <button type="button" class="btn btn-info">
                                                    <span class="entypo-pencil"></span>&nbsp;&nbsp;Edit
                                                </button>
                                            </a>
                                            <input type="hidden" name="_method" value="DELETE">
                                            <button  style="margin-left:10px;" class=" btn btn-danger " title="文章删除">
                                                <span class="entypo-trash"></span>&nbsp;&nbsp;删除</button>
                                        </form>
                                    </td>
                                </tr>
                                {{end}}
                                </tbody>

                            </table>
                        </div>
                    </section>
                </div>
            </div>
        </div>
    </div>
</div>





