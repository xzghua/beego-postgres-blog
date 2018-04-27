<section id="hero" class="scrollme">
    <div class="container-fluid element-img" style="background: url(/static/index/img/1.jpeg) no-repeat center center fixed;background-size: cover">
        <div class="row">
            <div class="col-xs-12 col-sm-8 col-sm-offset-2 col-md-8 col-md-offset-2 vertical-align cover boost text-center">
                <div class="center-me animateme" data-when="exit" data-from="0" data-to="0.6" data-opacity="0" data-translatey="100">
                    <div>

                        <h2>叶落山城秋</h2>
                        <p></p>

                        <h2></h2>
                        <p></p>


                    </div>
                </div>
            </div>
            <!-- // .col-md-12 -->
        </div>
        <div class="herofade beige-dk"></div>
    </div>
</section>

<!-- Height spacing helper -->
<div class="heightblock"></div>
<!-- // End height spacing helper -->

<!-- ============================ END Hero Image =========================== -->
<!-- ============================ Content =========================== -->

<section id="intro">
    <div class="container">

        {{range $index,$item := .post }}

        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3 style="white-space: pre-wrap;word-wrap: break-word;"><a href="post/43/index.html">{{$item.Title}}</a></h3>
                    <span>
                        <span class="post-meta">
                        <time datetime="{{$item.CreatedAt}}" itemprop="datePublished">
                        {{$item.CreatedAt}}
                        </time>
                        |
                            {{range $key,$value := $item.tag_list}}
                                    {{ range $k,$v := $value }}
                                        <span class="el-tag el-tag--success"> {{$v}}</span>
                                    {{end}}
                            {{end}}
                        </span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>
                    {{$item.Abstract}} ...
                    </p>

                    <p class="pull-right readMore">
                        <a href="post/43/index.html">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>
        {{end}}



        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3><a href="post/42/index.html">Mac 下修改 iTunes 备份文件夹路径</a></h3>
                    <span>
                        <span class="post-meta">
      <time datetime="2017-11-30T11:08:02.000Z" itemprop="datePublished">
          2017-11-30
      </time>


    |
    <a href='tags/iPhone/index.html'>iPhone</a>,

    <a href='tags/备份/index.html'>备份</a>,

    <a href='tags/iTunes/index.html'>iTunes</a>


</span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>

                    <p><img src="http://kieran-hexo.qiniudn.com/hexo_42_1.png" alt="iTunes"></p>
                    <h3 id="前言"><a href="index.html#前言" class="headerlink" title="前言"></a>前言</h3><p>手机电池又不行了，淘宝入了块电池自己换<br>换之前自然先备份下 iPhone<br>但 256GB 的 Mac 内存吃紧，想着备份到移动硬盘上<br>结果找了半天没有在 iTunes 里发现更改备份文件夹的设置<br>无奈搜索一番，在此记录一下<br>

                </p>

                    <p class="pull-right readMore">
                        <a href="post/42/index.html">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>

        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3><a href="post/41/index.html">HTTP 2.0 杂谈</a></h3>
                    <span>
                        <span class="post-meta">
      <time datetime="2017-08-11T09:01:19.000Z" itemprop="datePublished">
          2017-08-11
      </time>


    |
    <a href='tags/HTTP-2-0/index.html'>HTTP 2.0</a>


</span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>


                        正经的开场当前互联网大部分的内容都使用 HTTP1.1 作为通信协议。大家在这项协议上投入了非常多的“精力”，比如从 HTTP1.0 规范的 60 页暴涨到 HTTP1.1 规范的 176 页。相信大部分人都没有兴趣完整的阅读完这份规范，这样导致的结果就是 HTTP1.1 中的很多细节都没有被很好的实现，即使一开始实现的功能也很少被使用。而随着时间推移…使用者才发现，原来这货有这么多毛病。
                        HTTP1.1 的功与过但在挑毛病之前，我们

                    </p>

                    <p class="pull-right readMore">
                        <a href="post/41/index.html">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>

        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3><a href="post/40/index.html">都是 Edge(IE) 缓存惹的祸</a></h3>
                    <span>
                        <span class="post-meta">
      <time datetime="2017-01-15T14:58:19.000Z" itemprop="datePublished">
          2017-01-15
      </time>


    |
    <a href='tags/Edge/index.html'>Edge</a>,

    <a href='tags/缓存/index.html'>缓存</a>


</span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>

                        上周某天上班的时候，突然后端同志私戳我，当时我正研究着另一个项目的PSD，内心认真规划着整个项目的代码结构，指尖流淌出迷人的键盘音……好吧，其实内心想着中午吃什么菜，喝什么汤仔细一看，第一反应是，怎么可能。于是爸爸我连忙打开虚拟机中滴 Win10，等待 Win10 爷爷慢悠悠的进入待机状态。迅速点开 Edge 。打开网页和 network 控制台，迅速进入调试状态。一点退出登录，“啪”一个退出请求出现在 network 里，账号也随之退

                    </p>

                    <p class="pull-right readMore">
                        <a href="post/40/index.html">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>

        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3><a href="post/39/index.html">挑食，不挑食</a></h3>
                    <span>
                        <span class="post-meta">
      <time datetime="2016-11-25T11:11:47.000Z" itemprop="datePublished">
          2016-11-25
      </time>

</span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>

                        “长大以后终于知道大人为什么不挑食。因为他们买菜的时候就只买自己吃的。”
                        最近，我发现我似乎“被”变挑食了。从来没有觉得自己挑食，以往在学校也是自己点自己喜欢吃的，所以也不会点了这个菜却不吃。但是在公司就不一样，菜是搭配好的，平均每份2~3个菜，也许其中只有1~2个菜是我的口味。只吃我爱吃的菜就能满足我的这顿餐，那个我不爱的菜自然被剩下了。好吧，可能这就是挑食。西兰花，花菜，豆芽菜，莴苣…这些以往我都是吃的，虽然谈不上多喜爱，但也没

                    </p>

                    <p class="pull-right readMore">
                        <a href="post/39/index.html">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>

        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3><a href="post/38/index.html">怪诞人生</a></h3>
                    <span>
                        <span class="post-meta">
      <time datetime="2016-11-14T16:48:08.000Z" itemprop="datePublished">
          2016-11-15
      </time>


    |
    <a href='tags/Misc/index.html'>Misc</a>,

    <a href='tags/关于我/index.html'>关于我</a>,

    <a href='tags/电影/index.html'>电影</a>,

    <a href='tags/战争/index.html'>战争</a>


</span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>

                        关于电影“每个人都有自己对战争的理解，对电影也一样。”
                        上面这句话来自李安《比利·林恩的中场战事》的结尾，比利对着镜头，像是对这部电影前面的总结，也像是李安的喃喃自语。有一种高处不胜寒的孤寂感。影片一开始，就是一个超大光圈的镜头，全片多处用了大光圈来虚化背景，这在电影里很不常见，会降低电影的感觉，但在这部倾向于个人叙事视角的电影里，多了几分真实。纵观整部影片，你会发现全片的光线都是很亮的那种，但是又不会显得突兀，包括舞台上，房屋内，

                    </p>

                    <p class="pull-right readMore">
                        <a href="post/38/index.html">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>

        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3><a href="post/37/index.html">一张 CD 的动画</a></h3>
                    <span>
                        <span class="post-meta">
      <time datetime="2016-05-10T08:01:05.000Z" itemprop="datePublished">
          2016-05-10
      </time>


    |
    <a href='tags/CSS3/index.html'>CSS3</a>


</span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>

                        两个 Div 实现的动画12&lt;div class="cd"&gt;&lt;/div&gt;&lt;div class="music"&gt;&lt;/div&gt;
                        为什么要写这个呢？答案是 玩没错 就是玩…可能脑抽吧其实我发现 CSS 动画真的很好玩！可以实现一些炫酷的效果接下来我会继续写几个好玩的 Demo补充下我并不充实的 CSS3 动画知识…玩并快乐的学习~代码请戳下面http://codepen.io/_kieran/p

                    </p>

                    <p class="pull-right readMore">
                        <a href="post/37/index.html">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>

        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3><a href="post/36/index.html">Mongoose 学习笔记</a></h3>
                    <span>
                        <span class="post-meta">
      <time datetime="2016-05-08T08:01:05.000Z" itemprop="datePublished">
          2016-05-08
      </time>


    |
    <a href='tags/Mogoose/index.html'>Mogoose</a>


</span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>

                        老是忘记，既然这样，就写篇文章记录下自己常用的用法吧碰到新的用法，随时保持更新
                        Schema123456789101112131415161718192021222324252627282930313233var mongoose = require('mongoose');var Schema = mongoose.Schema; var blogSchema = new Schema(&#123;  title:  String

                    </p>

                    <p class="pull-right readMore">
                        <a href="post/36/index.html">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>

        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3><a href="post/35/index.html">在 Coding 使用 Webhook 实现自动部署</a></h3>
                    <span>
                        <span class="post-meta">
      <time datetime="2016-04-05T15:49:39.000Z" itemprop="datePublished">
          2016-04-05
      </time>


    |
    <a href='tags/Webhook/index.html'>Webhook</a>


</span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>

                        最近碰到这样的一个需求，就是本地提交代码到 git 仓库以后需要网站远程自动同步代码。因为项目不便公开，所以 Coding 的免费的私有仓库成了首选看了下 Coding 的 webhook 正好能满足这个需求。既然如此，那我们先来看下 Webhook 的概念吧
                        WebHookWebhook 允许第三方应用监听 Coding.net 上的特定事件，在这些事件发生时通过 HTTP POST 方式通知( 超时5秒) 到第三方应用指定的 W

                    </p>

                    <p class="pull-right readMore">
                        <a href="post/35/index.html">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>

        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3><a href="post/34/index.html">网站速度优化（一）——Express中间件摆放顺序之坑</a></h3>
                    <span>
                        <span class="post-meta">
      <time datetime="2016-03-15T16:23:07.000Z" itemprop="datePublished">
          2016-03-16
      </time>


    |
    <a href='tags/Node-js/index.html'>Node.js</a>


</span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>

                    <p>为什么我写的Express服务静态资源访问速度这么慢？？？<br>直接把这个问题放进百度当然是没有答案的<br>准确的说是没有准确的答案<br>谷歌之 依然无果<br>翻译成英文再搜<br>最上面一条是StackoverFlow的<br>上面的人提了和我一样的问题<br>

                    </p>

                    <p class="pull-right readMore">
                        <a href="post/34/index.html">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>

        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3><a href="post/33/index.html">React Native Android在Windows下的踩坑之旅</a></h3>
                    <span>
                        <span class="post-meta">
      <time datetime="2015-12-23T16:32:43.000Z" itemprop="datePublished">
          2015-12-24
      </time>


    |
    <a href='tags/React/index.html'>React</a>,

    <a href='tags/React-Native/index.html'>React Native</a>,

    <a href='tags/Android/index.html'>Android</a>


</span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>


                        直接开始环境搭建在网上有很多教程，不过还是推荐去看官方文档 https://facebook.github.io/react-native/docs/android-setup.html#content  ， 当然，照着文档来还是会有很多问题，就我搭建的过程而言，简直就是一步三个坑，这里官方文档上有的我就一笔带过了，详细讲讲我遇到的一些坑，相信这些大家更有用一些。
                        坑一：Android SDK之前安装过Android Studio

                    </p>

                    <p class="pull-right readMore">
                        <a href="post/33/index.html">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>

        <div class="row">
            <div class="col-md-8 col-md-offset-2 opening-statement">
                <div class="col-md-4">
                    <h3><a href="post/32/index.html">让 WebStorm 支持 ECMAScript 6 和 JSX</a></h3>
                    <span>
                        <span class="post-meta">
      <time datetime="2015-12-05T14:05:56.000Z" itemprop="datePublished">
          2015-12-05
      </time>


    |
    <a href='tags/JavaScript/index.html'>JavaScript</a>,

    <a href='tags/WebStrom/index.html'>WebStrom</a>


</span>
                    </span>
                </div>
                <div class="col-md-8">
                    <p>

                        Oh shit! WebStrom竟然不支持ES6？！别闹！怎么可能！！！当然不可能了经过我一番捣鼓终于发现了在设置里的小秘密…看下图默认是 ECMAScript 5.1 的语法支持要想 WebStrom 的代码提示和静态语法检测支持 ES6 和 JSX 得在 setting 里把 JavaScript language version 改成 JSX harmony如果你不写 JSX 可以改成 ECMAScript 6，JSX是JS的超

                    </p>

                    <p class="pull-right readMore">
                        <a href="post/32/index.html">Read More...</a>
                    </p>

                </div>
                <div class="clearfix"></div>
                <hr class="nogutter">
            </div>
        </div>

        <nav class="pagination" role="pagination">

            <a class="pull-right" href="page/2/index.html">Older Posts →</a>
        </nav>
    </div>
</section>