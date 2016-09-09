# shorturl

# 设计思路
一般用于将长网址转换成较短的网址，该服务提供多个短网址服务商支持
gggg:目前唯一没有测试出问题(实现)
adfly:暂不支持(未实现)
bitly：需要bitly账号[crazyant2016(crazyant)]，BITLY_API_KEY 和 BITLY_LOGIN(R_98fbed256f0946168b0be5e0f08028b0,crazyant2016),不过有容量限制，初级是5000(未实现)
google:需要 	GOOGL_API_KEY，没有测试(未实现)
gitio:只能转换 github.com的地址(未实现)
shorl：用短网址访问时，是先跳到该网站，然后再进行跳转，跳转过程会出现 shorl网站，很蛋疼(未实现)
rddme：同上，而且访问速度很慢，真恶心(未实现)
tinyurl：不支持标签结尾的网址，比如 http://blog.sina.com.cn/s/blog_8882a6260101a4ox.html 这种网址(未实现)
catchy,cligs,isgd,moourl,pendekin,snipurl,vamu：都无法连接上(未实现)

## 使用
参考测试用例


## 安装
参考Dockerfile

# 环境变量
暂无
