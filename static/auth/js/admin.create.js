/**
 * Created by ylsc on 16-8-8.
 */

/**
 * 校验分类列表的表单
 * @type {any}
 */
var createForm = $('.createForm').validate({
    rules:{
        name:{
            required:true,
            maxlength:20
        },
        displayName:{
            required:true,
            maxlength:20
        },
        description:{
            maxlength:120
        }

    },
    messages:{
        name:{
            required:'请输入分类名称',
            maxlength:'长度超出最大范围'
        },
        displayName:{
            required:'请输入分类别名',
            maxlength:'长度超出最大范围'
        },
        description:{
            maxlength:'长度超出最大范围'
        }
    }

});

/**
 * 校验标签列表
 * @type {any}
 */
var tagForm = $('.tagForm').validate({
    rules:{
        name:{
            required:true,
            maxlength:20
        }
    },
    messages:{
        name:{
            required:'请输入标签名称',
            maxlength:'标签长度超出'
        }
    }
});

/**
 * 写文章表单
 * @type {any}
 */
var articleForm = $('.articleForm').validate({
    rules:{
        title:{
            required:true,
            maxlength:100
        },
        category:{
            required:true,
            min:1
        },
        tag:{
            required:true
        },
        abstract:{
            required:true
        }
    },
    messages:{
        title:{
            required:'请输入标题',
            maxlength:'已超出范围'
        },
        category:{
            required:'请选择分类',
            min:'请选择分类'
        },
        tag:{
            required:'请输入标签'
        },
        abstract:{
            required:'请输入摘要'
        }
    }
});

var linkForm = $('.linkForm').validate({
   rules:{
       name:{
           required:true,
           maxlength:30
       },
       link:{
           required:true,
           url:true
       },
       ordering:{
           required:true,
           digits:true,
           min:0
       }
   },
    messages:{
        name:{
            required:'请输入友链名',
            maxlength:'超出最大长度'
        },
        link:{
            required:'请输入友链地址',
            url:'请输入正确的地址'
        },
        ordering:{
            required:'请输入序号(重叠不影响)',
            digits:'请输入整数',
            min:'请输入大于或等于0的整数'
        }
    }
});