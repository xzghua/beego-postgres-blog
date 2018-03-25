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

                        <form action="/console/post" method="post" id="postCreate" class="form-horizontal articleForm">
                        {{ .xsrfdata }}
                            <fieldset>
                                <div class="control-group">
                                    <label class="control-label" for="title">标题</label>
                                    <div class="controls">
                                        <input type="text" class="form-control" name="title" id="title">
                                    </div>
                                </div>
                                <div class="control-group">
                                    <label class="control-label" for="categories">分类</label>
                                    <div class="controls">
                                        <select class="form-control" name="category"  id="categories">
                                            {{range $index,$item := .cate }}
                                                <option value="{{ $item.Id }}">{{str2html $item.html}}{{$item.DisplayName}}</option>
                                            {{end}}
                                        </select>
                                    </div>
                                </div>
                                <div class="control-group" style="margin-bottom: 20px">
                                    <label class="control-label" for="subject">标签</label>
                                    <div class="controls">
                                        <input type="text" name="tag" id="tag" class="form-control round-input tag">
                                    </div>
                                </div>
                                <div class="control-group" style="margin-bottom: 20px">
                                    <label class="control-label" for="subject">摘要</label>
                                    <div class="alert alert-info" style="margin: 1px;padding: 6px">
                                        <button data-dismiss="alert" class="close" type="button">×</button>
                                        <span class="entypo-info-circled"></span>
                                        <strong>建议写带HTML标签的摘要,为什么要呢?省得解析,还有文章断句,目前找不到组件</strong>
                                    </div>
                                    <div style="padding:14px;" class="form-group">
                                        <textarea style="resize: none;" name="abstract" id="abstract" rows="5" class="form-update"></textarea>
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="email">内容</label>
                                    <div id="editormd">
                                        <textarea  name="content" style="display:none;">### Hello Editor.md !</textarea>
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