local_url=127.0.0.1:9000
url=api.lingdei.doyi.online

mkdir -p ./build

rm -rf ./build/*
cp -r * ./build

cd ./build

sed -i "s/'${local_url}'/'${url}'/g" main.go
swag fmt
swag init

docker login --username=东南dnf registry.cn-shanghai.aliyuncs.com
docker build -t registry.cn-shanghai.aliyuncs.com/zjy2414/lingdei-middleware:latest .
docker push registry.cn-shanghai.aliyuncs.com/zjy2414/lingdei-middleware:latest

cd ../
rm -rf ./build