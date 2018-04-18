<link href="/static/auth/js/validate/validate.css" rel="stylesheet">
<link rel="stylesheet" type="text/css" href="/static/auth/js/tag/jquery.tagsinput.css">
<link rel="stylesheet" href="/static/auth/css/editormd.min.css" />
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

                        <form action="/console/cate/{{.cate.Id}}" method="post" id="postCreate" class="form-horizontal createForm">
                        {{ .xsrfdata }}
                            <fieldset>
                                <input type="hidden" name="_method" value="PUT">
                                <div class="control-group">
                                    <label class="control-label" for="parentId">上级分类</label>
                                    <div class="controls">
                                        <select class="form-control" name="parentId"  id="parentId">
                                        {{range $index,$item := .cateSort }}
                                            <option value="{{ $item.Id }}" {{if eq $item.Id $.cate.ParentId }} selected {{end}}>{{str2html $item.html}}{{$item.DisplayName}}</option>
                                        {{end}}
                                        </select>
                                    </div>
                                </div>
                                <div class="control-group">
                                    <label class="control-label" for="name">分类名称</label>
                                    <div class="controls">
                                        <input type="text" class="form-control" value="{{ .cate.Name }}" name="name" id="name">
                                    </div>
                                </div>
                                <div class="control-group">
                                    <label class="control-label" for="displayName">分类别名</label>
                                    <div class="controls">
                                        <input type="text" class="form-control" value="{{ .cate.DisplayName }}" name="displayName" id="displayName">
                                    </div>
                                </div>
                                <div class="control-group" style="margin-bottom: 20px">
                                    <label class="control-label" for="description">描述</label>
                                    <div style="padding:14px;" class="form-group">
                                        <textarea style="resize: none;" name="description" id="description" rows="5" class="form-update">{{.cate.Description}}</textarea>
                                    </div>
                                </div>
                                <div class="form-actions" style="margin:20px 0 0 0;">
                                    <button type="submit" class="btn btn-primary">Submit</button>
                                    <button type="reset" class="btn">Cancel</button>
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