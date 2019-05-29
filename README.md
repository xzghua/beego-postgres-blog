# 早已停更,最新重构的请左转 [https://github.com/izghua/go-blog](https://github.com/izghua/go-blog)

### 瞎搞


> 因为要熟悉熟悉`Go`,又没有其他的项目可以试试,只能自己用`Go`重构一下自己的[垃圾博客](https://www.iphpt.com)了

> 所以这次,打算用`beego`框架,数据库试试`postgresql`,给自己挖挖坑!

恩,坑什么时候能搞完还不一定!毕竟我还要去拯救世界,所以很忙的!



### yoo

> Because I want to be familiar with `Go`, there are no other projects to try, I can only use` Go` to reconstruct [my own rubbish blog](https://www.iphpt.com)

> So this time, I'm going to use `beego` framework, try` postgresql` in database, dig digging for myself!

> Well, pit when it is not necessarily finished! After all, I have to save the world, so I am very busy!(Google Translate)


--------------------------分割线--------------------------

- [x] 已完成前台显示,后台展示,页面跟老博客一毛一样,我感觉还是挺好看的,难得有这么一个页面,我两三年还喜欢
- [x] 前台所有数据都加了缓存,减少数据库的查询操作!
- [x] 自动拦截注册,当注册用户达到两个以后,停止注册!
- [x] 评论: 接入搜狐的 `畅游`,还有`github`的issue功能,需要那个,后台配置下即可(配置完以后,清除缓存即可)
- [x] 修改以前`分类`,`标签`文章数过多但不分页的情况,现在开始有分页
- [x] 修改php版本的页面分页 不省略的问题
- [x] 自定义设置 静态文件的来源(本地,7牛)
- [ ] 即将加入`sitemap`
- [ ] 想做个 `dockerfile`,然后一键安装所有的环境,毕竟这个需要`golang`+`psql`+`redis`+`nginx` 或许还要加入 `supervisor`
