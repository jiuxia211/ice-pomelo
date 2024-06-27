#!/bin/bash

# 启动微服务的函数
start_service() {
  service_name=$1
  start_command=$2

  echo "Starting $service_name..."
  $start_command > "${service_name}.log" 2>&1 &
  echo "$service_name started"
}

# 清理旧的 PID 文件
rm -f service_pids.txt

# 启动所有微服务
start_service "api" "sh ./cmd/api/output/bootstrap.sh"
start_service "tiny_id" "sh ./cmd/tiny_id/output/bootstrap.sh"
start_service "user" "sh ./cmd/user/output/bootstrap.sh"
start_service "video" "sh ./cmd/video/output/bootstrap.sh"

echo "All services started."

# 捕获退出信号以停止所有服务
trap 'echo "Stopping services..."; kill $(jobs -p); exit' INT TERM

# 保持脚本运行以保持微服务运行
wait
