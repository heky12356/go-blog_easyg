#!/bin/bash

# 配置
APP_NAME="easyg"
INSTALL_DIR="/opt/easyg/backend"
BINARY_NAME="easyg"  # 编译后生成的二进制名称

# 函数：构建项目
build_project() {
  cd "$INSTALL_DIR" || exit
  go mod tidy
  go build -o "$BINARY_NAME"
}

# 函数：启动服务
start_service() {
  cd "$INSTALL_DIR" || exit
  nohup ./"$BINARY_NAME" > app.log 2>&1 &
  echo $! > "$INSTALL_DIR/service.pid"
  echo "服务已启动，PID: $(cat "$INSTALL_DIR/service.pid")"
}

# 函数：停止服务
stop_service() {
  if [ -f "$INSTALL_DIR/service.pid" ]; then
    kill "$(cat "$INSTALL_DIR/service.pid")"
    rm "$INSTALL_DIR/service.pid"
    echo "服务已停止"
  else
    echo "未找到 PID 文件，服务可能未运行"
  fi
}

# 函数：服务状态
status_service() {
  if [ -f "$INSTALL_DIR/service.pid" ]; then
    PID=$(cat "$INSTALL_DIR/service.pid")
    if ps -p "$PID" > /dev/null; then
      echo "服务正在运行 (PID: $PID)"
    else
      echo "服务未运行，但存在 PID 文件"
    fi
  else
    echo "服务未运行"
  fi
}

# 显示帮助
show_help() {
  echo "用法: $0 {install|start|stop|status}"
}

# 主逻辑
case "$1" in
  install)
    clone_repo
    build_project
    ;;
  start)
    start_service
    ;;
  stop)
    stop_service
    ;;
  status)
    status_service
    ;;
  *)
    show_help
    ;;
esac
