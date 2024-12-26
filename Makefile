# commit 最新hash值 代码版本号
ImageCommitId = $(shell git rev-parse --short HEAD)
# commit上一次hash值
ImageLastCommitId = $(shell git log --pretty="%H" | sed -n "2p" | cut -c1-7)
# 镜像版本名称
ProjectImgTag = wmt1030/seeyou-server:$(ImageCommitId)
# 镜像名称
ProjectImg = wmt1030/seeyou-server
# 容器名称
ProjectName = seeyou-server
# 端口
Port = 9102
# 上一次镜像名称
LastImage = wmt1030/seeyou-server:$(ImageLastCommitId)

project: project-delete project-build project-run

project-delete:
    # 删除旧容器
	if docker ps -a | awk '{print $NF}' | grep -q "${ProjectName}"; then \
	docker ps -a | grep ${ProjectName} | awk '{print $1}' | xargs -I docker stop {} | xargs -I docker rm {}; \
	fi
	# 删除旧镜像
	if docker images | grep ${ProjectImg} | grep ${ImageLastCommitId}; then \
	docker rmi ${LastImage}; \
	fi
# --build-arg 设置构建时的变量在dockerfile中接收 --build-arg version=$(VERSION) 可不加
# --platform linux/amd64 指定镜像环境，否则从linux服务器拉取运行失败
# -t 设置镜像名称
# -f dockerfile名称（路径）注意后面有个点
project-build:
	docker build --platform linux/amd64 -t ${ProjectImgTag} -f Dockerfile . --no-cache

# project-push:
# 	docker push ${ProjectImgTag}

#  --restart=always 重启docker之后容器也会自动重启
#  在启动时如果没有添加这个参数 使用命令 ' docker container update --restart=always 容器名字 ' 后续添加
# -d: 后台运行容器，并返回容器ID
# -p  8080:8080 指定容器暴露的端口
# --name 指定容器名称
project-run:
	docker run -d --name ${ProjectName} -p ${Port}:80 --restart=always ${ProjectImgTag}

