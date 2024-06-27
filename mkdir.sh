#!/bin/bash

# 定义目标目录
TARGET_DIR="./data"

# 创建目标目录
mkdir -p "$TARGET_DIR"

# 在目标目录下创建子目录
mkdir -p "$TARGET_DIR/etcd"
mkdir -p "$TARGET_DIR/mysql"
mkdir -p "$TARGET_DIR/redis"
chmod -R 777 "$TARGET_DIR/etcd"
chmod -R 777 "$TARGET_DIR/mysql"
chmod -R 777 "$TARGET_DIR/redis"

echo "Directories created successfully:"
echo "$TARGET_DIR/etcd"
echo "$TARGET_DIR/mysql"
echo "$TARGET_DIR/redis"