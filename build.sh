# 创建 tmp 目录
mkdir tmp

# 生成二进制文件
go build -o tmp/goblog

# 复制环境配置信息
cp .env tmp/

# 进入 tmp 目录下运行我们的程序
cd tmp

./goblog