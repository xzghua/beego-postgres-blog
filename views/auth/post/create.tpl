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

                        <form action="#" id="contact-form" class="form-horizontal articleForm">
                            <fieldset>

                                <div class="control-group">
                                    <label class="control-label" for="title">标题</label>
                                    <div class="controls">
                                        <input type="text" class="form-control" name="title" id="title">
                                    </div>
                                </div>
                                <div class="control-group">
                                    <label class="control-label" for="email">分类</label>
                                    <div class="controls">
                                        <select class="form-control" id="email">
                                            <option>1</option>
                                            <option>2</option>
                                        </select>
                                    </div>
                                </div>
                                <div class="control-group" style="margin-bottom: 20px">
                                    <label class="control-label" for="subject">标签</label>
                                    <div class="controls">
                                        <input type="text" name="tag" id="tag" class="form-control round-input tag">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="email">内容</label>
                                    <div id="editormd">
                                        <textarea style="display:none;">### Hello Editor.md !</textarea>
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
<script src="https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js"></script>
<script type="text/javascript" src="/static/auth/js/validate/jquery.validate.min.js"></script>
<script type="text/javascript" src="/static/auth/js/tag/jquery.tagsinput.js"></script>
<script type="text/javascript" src="/static/auth/js/jquery-ui.min.js"></script>
<script type="text/javascript" src="/static/auth/js/validate/jquery.validate.min.js"></script>

<script type="text/javascript" src="/static/auth/js/admin.create.js"></script>
<script src="/static/auth/js/editormd.min.js"></script>
<script>

    var testEditor;

    $(function() {
        testEditor = editormd("editormd", {
            width: "100%",
            height: 640,
            syncScrolling: "single",
            path: "/static/auth/js/lib/",
            imageUpload : true,
            imageFormats : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
            imageUploadURL : "/uploadPhotos",
            toolbarAutoFixed:false,
            // markdown : md,
            codeFold : true,

            saveHTMLToTextarea : true,    // 保存 HTML 到 Textarea
            searchReplace : true,
            //watch : false,                // 关闭实时预览
            htmlDecode : "style,script,iframe|on*",            // 开启 HTML 标签解析，为了安全性，默认不开启

            emoji : true,
            taskList : true,

            theme : "default | dark",
            // Preview container theme, added v1.5.0
            // You can also custom css class .editormd-preview-theme-xxxx
            previewTheme : " dark",
            // Added @v1.5.0 & after version this is CodeMirror (editor area) theme
            editorTheme : editormd.editorThemes['editormd']
        });
    });

    $(function() {

        $('#tag').tagsInput({

            autocomplete_url:"/autoCompleteTags",

            previewTheme : "dark",
            editorTheme : "pastel-on-dark",
            markdown : 'md',
            codeFold : true,

            width: 'auto',
            height:'10px',
            defaultText:'添加标签',
            minChars:'0',
            maxChars:'20',
            placeholderColor:'#ff0066'

        });

    });
</script>