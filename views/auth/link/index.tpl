
<!-- BREADCRUMB -->
<ul id="breadcrumb">
    <li>
        <span class="entypo-home"></span>
    </li>
    <li><i class="fa fa-lg fa-angle-right"></i>
    </li>
    <li><a href="#" title="Sample page 1">友链</a>
    </li>
    <li><i class="fa fa-lg fa-angle-right"></i>
    </li>
    <li><a href="#" title="Sample page 1">友链列表</a>
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
                        友链列表</h6>
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
                            <a href="/console/link/create" style="margin-left:10px;" class="pull-left btn btn-info " title="新增友链">新增友链</a>

                            <table class="table  table-hover  table-condensed table-striped ">
                                <thead class="cf">
                                <tr>
                                    <th>ID</th>
                                    <th>友链名</th>
                                    <th>友链链接</th>
                                    <th>友链排名</th>
                                    <th>创建时间</th>
                                    <th>操作</th>
                                </tr>
                                </thead>
                                <tbody>
                                {{range $index,$item := .link }}
                                <tr>
                                    <td>{{$item.Id}}</td>
                                    <td>{{$item.Name}}</td>
                                    <td>{{$item.Link}}</td>
                                    <td>{{$item.Sort}}</td>
                                    <td>{{$item.CreatedAt}}</td>
                                    <td>
                                        <form action="/console/link/{{$item.Id}}" method="post">
                                        {{ $.xsrfdata }}
                                            <input type="hidden" name="_method" value="DELETE">
                                            <a href="/console/link/{{$item.Id}}/edit">
                                                <button type="button" class="btn btn-info">
                                                    <span class="entypo-pencil"></span>&nbsp;&nbsp;Edit
                                                </button>
                                            </a>
                                            <button  style="margin-left:10px;" class=" btn btn-danger " title="分类删除">
                                                <span class="entypo-trash"></span>&nbsp;&nbsp;删除</button>

                                        </form>
                                    </td>
                                </tr>
                                {{end}}
                                </tbody>
                                <tfoot>
                                <tr>
                                    <td colspan="5">
                                        <div class="btn-group">
                                            <a href="/console/link?page={{.lastPage}}"><button type="button" class="btn btn-warning">上一页</button></a>
                                            <a href="#" disabled="disabled"><button type="button" class="btn btn-danger" disabled>Middle</button></a>
                                            <a href="/console/link?page={{.nextPage}}"> <button type="button" class="btn btn-info">下一页</button></a>
                                        </div>
                                    </td>
                                </tr>
                                </tfoot>
                            </table>
                        </div>
                    </section>
                </div>
            </div>
        </div>
    </div>
</div>





