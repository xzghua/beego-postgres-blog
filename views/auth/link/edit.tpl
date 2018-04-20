<link href="/static/auth/js/validate/validate.css" rel="stylesheet">
<link rel="stylesheet" type="text/css" href="/static/auth/css/jquery-ui.min.css">

<div class="content-wrap">
    <div class="row">
        <div class="col-sm-12">
            <div class="nest" id="validationClose">
                <div class="title-alt">
                    <h6>
                        新建</h6>
                    <div class="titleClose">
                        <a class="gone" href="#validationClose">
                            <span class="entypo-cancel"></span>
                        </a>
                    </div>
                    <div class="titleToggle">
                        <a class="nav-toggle-alt" href="#validation">
                            <span class="entypo-up-open"></span>
                        </a>
                    </div>

                </div>

                <div class="body-nest" id="validation">
                    <div class="form_center">

                        <form action="/console/link/{{.link.Id}}" method="post" id="linkCreate" class="form-horizontal linkForm">
                        {{ .xsrfdata }}
                            <input type="hidden" name="_method" value="PUT">
                            <fieldset>
                                <div class="form-group">
                                    <label class="col-sm-3 control-label" for="linkName">友链名称</label>
                                    <div class="col-sm-6">
                                        <input type="text" value="{{.link.Name}}" id="linkName" name="name"  class="form-control linkName">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-sm-3 control-label">友链链接</label>
                                    <div class="col-sm-6">
                                        <input type="text" value="{{.link.Link}}" id="link" name="link" class="form-control round-input link">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-sm-3 control-label">顺序</label>
                                    <div class="col-sm-6">
                                        <input type="text" value="{{.link.Sort}}" id="ordering" name="ordering" class="form-control ordering">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-sm-3 control-label"></label>
                                    <div class="col-sm-6">
                                        <button class="btn" type="submit">
                                            <span class="fontawesome-save"></span>&nbsp;&nbsp;修改</button>
                                    </div>
                                </div>
                            </fieldset>
                        </form>

                    </div>
                </div>

            </div>
        </div>
    </div>
</div>




</div>