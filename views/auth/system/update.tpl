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

                        <form action="/console/cate" method="post" id="postCreate" class="form-horizontal createForm">
                        {{ .xsrfdata }}
                            <fieldset>

                                <div class="control-group">
                                    <label class="control-label" for="name">网站标题</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.Title}}" class="form-control" name="name" id="name">
                                    </div>
                                </div>
                                <div class="control-group">
                                    <label class="control-label" for="displayName">标题SEO名</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.STitle}}" class="form-control" name="displayName" id="displayName">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">网站描述</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.Description}}" class="form-control" name="description" id="description">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">网站SEO关键词</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.SeoKey}}" class="form-control" name="description" id="description">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">网站SEO描述</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.SeoDes}}" class="form-control" name="description" id="description">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">网站备案号</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.RecordNumber}}" class="form-control" name="description" id="description">
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