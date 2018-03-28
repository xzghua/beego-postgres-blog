
<!-- BREADCRUMB -->
<ul id="breadcrumb">
    <li>
        <span class="entypo-home"></span>
    </li>
    <li><i class="fa fa-lg fa-angle-right"></i>
    </li>
    <li><a href="#" title="Sample page 1">标签</a>
    </li>
    <li><i class="fa fa-lg fa-angle-right"></i>
    </li>
    <li><a href="#" title="Sample page 1">标签列表</a>
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
                        标签列表</h6>
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
                    <a href="/console/tag/create" ><button type="button" class="btn btn-danger" >新建</button></a>

                    <section id="flip-scroll">
                        <div class="table-responsive">
                            <table class="table  table-hover  table-condensed table-striped ">
                                <thead class="cf">
                                <tr>
                                    <th>ID</th>
                                    <th>名称</th>
                                    <th>使用次数</th>
                                    <th>操作</th>
                                </tr>
                                </thead>
                                <tbody>
                                {{range $index,$item := .tag }}
                                <tr>
                                    <td>{{$item.Id}}</td>
                                    <td>{{$item.Name}}</td>
                                    <td>{{$item.TagNum}}</td>
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
                                <tfoot>
                                <tr>
                                    <td colspan="5">
                                        <div class="btn-group">
                                            <a href="/console/tag?page={{.lastPage}}"><button type="button" class="btn btn-warning">上一页</button></a>
                                            <a href="#" disabled="disabled"><button type="button" class="btn btn-danger" disabled>Middle</button></a>
                                            <a href="/console/tag?page={{.nextPage}}"> <button type="button" class="btn btn-info">下一页</button></a>
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





