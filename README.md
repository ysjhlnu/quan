
## 项目介绍
> 运维平台
> server 为后端
> web 为前端

## 功能


## QQ群

> 620176501

## 其他

> 本项目基于 gin-vue-admin

## 部署

```shell


cd /data/quan
git pull
cd web
npm run build



ps -ef | grep main | grep -v grep  | awk '{print $2}'|xargs kill -9

cd /data/quan/server 
nohup  go run  main.go   >>  /tmp/quan.log   2>&1  & 
tail -f /tmp/quan.log


server {
    listen       8080;
    server_name localhost;

    charset koi8-r;
    access_log  logs/host.access.log  main;

    location / {
        root /data/quan/web/dist;
        add_header Cache-Control 'no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0';
        try_files $uri $uri/ /index.html;
    }

    location /api {
        proxy_set_header Host $http_host;
        proxy_set_header  X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        rewrite ^/api/(.*)$ /$1 break;  #重写
        proxy_pass http://127.0.0.1:8888; # 设置代理服务器的协议和地址
     }

    location /api/swagger/index.html {
        proxy_pass http://127.0.0.1:8888/swagger/index.html;
     }
 }

```