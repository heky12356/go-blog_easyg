#!/bin/bash

#更新系统软件包
sudo apt update && sudo apt upgrade -y

#安装 Curl
sudo apt install curl -y

# 安装 Node.js
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs

# 安装 Go
curl -O https://storage.googleapis.com/golang/go1.20.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
cp -r /usr/local/go/bin/* /usr/bin

# 设置Go环境变量
export PATH=$PATH:/usr/local/go/bin

if [ grep -q 'export PATH=$PATH:/usr/local/go/bin' ~/.bashrc -eq 1 ]
then
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    source ~/.bashrc
elif [ grep -q 'export PATH=$PATH:/usr/local/go/bin' ~/.bash_profile -eq 1 ]
then
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile
    source ~/.bash_profile
elif [ grep -q 'export PATH=$PATH:/usr/local/go/bin' ~/.profile -eq 1 ]
then
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    source ~/.profile
fi

# 安装 Vite
# npm install -g vite

# 安装 Git
sudo apt install git -y


# 设置变量
REPO_URL="https://github.com/heky12356/go-blog_easyg.git"
CLONE_DIR="/opt/easyg"

# 确定web路径
read -p "请输入web路径:" DEPLOY_DIR

# 克隆仓库
git clone $REPO_URL $CLONE_DIR

# 进入目录
cd "$CLONE_DIR/easyg_frontend"     

# 安装依赖
npm install

# 打包
npm run build

# 移动文件
mv dist/* $DEPLOY_DIR

echo "安装完成"
