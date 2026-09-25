  占位符	            说明
    %v	    默认格式输出，万能占位符，日常首选
    %+v	    打印结构体，额外显示字段名
    %#v	    Go 语法格式输出，打印值类型 + 内容，调试神器
    %T	    打印变量类型
    %%	    输出一个百分号 %	
    %d	    十进制	
    %b	    二进制	
    %o	    八进制	
    %x	    十六进制小写	
    %X	    十六进制大写	
    %c	    字符，输出对应 
    %U	    Unicode 码点

    GMP


  ```函数中默认传递副本 除非传递指针才可以在函数内部修改值```
   ``` `切片slice`和`map`是引用类型 是引用传递```



方法	            方向	            作用
conn.Write(b)	    服务器 → 客户端	  把 b 发出去
conn.Read(b)	    客户端 → 服务器	  把对方发来的读进 b





	            net.Listen	                        net.Dial
角色	         服务端（被动）	                       客户端（主动）
作用	      监听本地端口，等待连接	                 主动发起连接
返回值	          net.Listener	                     net.Conn
阻塞行为	  阻塞等待 Accept，不阻塞在 Listen 本身	    阻塞直到连接建立或失败
典型场景	       服务器、守护进程	                      客户端、调用方

场景	                                   触发 io.EOF
客户端调用conn.Close()	                 ✅ 会
客户端进程正常退出	                      ✅ 会（OS 关闭 socket）
客户端调用 CloseWrite()（半关闭）	        ✅ 会





Git用法

git status            判断在哪个分支On branch main

git branch function1  创建function1分支
git switch function1  切换到function1分支

在main分支下git merge function1       把function1分支并入main分支
git branch -d function1               删除function1分支
git branch function2 97fd86bd58364377763ab8b7888f70574cd4792c     //在这个分支（用哈希值）上开发
git log --oneline                     提交记录
git remote -v                         检查绑定的github仓库
git push -u origin main               把你的本地 main 分支推送到远程仓库，并建立关联

git push                              推送到 GitHub
git pull                              从 GitHub 拉取	

git commit                            把改动记录到本地仓库
git push                              把本地提交上传到 GitHub