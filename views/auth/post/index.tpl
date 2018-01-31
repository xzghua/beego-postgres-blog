
<!-- BREADCRUMB -->
<ul id="breadcrumb">
    <li>
        <span class="entypo-home"></span>
    </li>
    <li><i class="fa fa-lg fa-angle-right"></i>
    </li>
    <li><a href="#" title="Sample page 1">文章</a>
    </li>
    <li><i class="fa fa-lg fa-angle-right"></i>
    </li>
    <li><a href="#" title="Sample page 1">文章列表</a>
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
                        文章列表</h6>
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
                                <th>作者</th>
                                <th>分类</th>
                                <th>标题</th>
                                <th>时间</th>
                                <th>操作</th>
                            </tr>
                            </thead>
                            <tbody>
                            {{range $index,$item := .post }}
                            <tr>
                                <td>{{$item.Id}}</td>
                                <td>{{$item.user_name}}</td>
                                <td>默认分类</td>
                                <td><a href="#">{{$item.Title}}</a></td>
                                <td>{{$item.CreatedAt}}</td>
                                <td>
                                    <button type="button" class="btn btn-info">
                                        <span class="entypo-pencil"></span>&nbsp;&nbsp;Edit
                                    </button>
                                    <button type="button" class="btn btn-danger">
                                        <span class="entypo-trash"></span>&nbsp;&nbsp;Delete
                                    </button>
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





