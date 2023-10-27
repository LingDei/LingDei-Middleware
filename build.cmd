set local_url = "127.0.0.1:9000"
set url = "api.lingdei.doyi.online"

sed -i "s/127.0.0.1:9000/api.lingdei.doyi.online/g" main.go
swag fmt
swag init

docker build -t registry.cn-shanghai.aliyuncs.com/zjy2414/lingdei-middleware:latest .
docker push registry.cn-shanghai.aliyuncs.com/zjy2414/lingdei-middleware:latest

sed -i "s/api.lingdei.doyi.online/127.0.0.1:9000/g" main.go
swag fmt
swag init