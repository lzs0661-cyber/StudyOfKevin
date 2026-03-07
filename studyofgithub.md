一、Git安装：
  1、mac和linux一般自带git
  2、windows需要安装；

二、Git本地设置：
  1、git config --global user.name "kevin"
  2、git config --global user.email "your github regist in github@xxx.com"
  上述两个命令会在～/.gitconfig中增加相关的配置

三、github配置访问权限：
  1、在本地执行以下命令生成公私钥：
    ssh-keygen -t rsa -C "your github regist in github@xxx.com"
    cat ~/.ssh/id_rsa.pub
    将公钥拷贝
  2、在github的Setting-->Security-->Deploy keys-->add deploy key
  3、将拷贝的公钥；
  4、测试是否联通：本地执行以下命令：
    ssh -T git@github.com

四、clone仓库
    1、github 的code下找到Clone，拷贝url
    2、本地执行git clone url
    
五、git的常见名利
   git config --global -l
   git config --system -l
   git config --global user.name /user.email "XXXX"

   
