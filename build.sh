#!/bin/bash

# 定义变量
PROJECT_ROOT=$(pwd) # 假设脚本在项目根目录下运行
BUILD_DIR="build"
EXECUTABLE_NAME="ffmpegforBililiveRecorder" # 替换为你的应用程序名称
PLATFORMS=("linux/amd64" "linux/arm64" "windows/amd64" "windows/arm64" "darwin/amd64" "darwin/arm64")

# 创建构建目录
mkdir -p "${BUILD_DIR}"

# 检查 conf.json 文件是否存在
if [ ! -f "${PROJECT_ROOT}/config.toml" ]; then
    echo "警告: 根目录下没有找到 config.toml 文件."
    exit 1
fi

# 编译并打包每个平台
for PLATFORM in "${PLATFORMS[@]}"; do
    OS=${PLATFORM%/*}
    ARCH=${PLATFORM#*/}
    ZIP_FILE="${EXECUTABLE_NAME}_${OS}_${ARCH}-$(date +%Y%m%d).zip"
    TMP_DIR="${BUILD_DIR}/${OS}_${ARCH}"
    
    mkdir -p "${TMP_DIR}"

    echo "正在编译 ${OS}/${ARCH}..."
    GOOS=${OS} GOARCH=${ARCH} go build -o "${TMP_DIR}/${EXECUTABLE_NAME}" || { echo "编译失败"; continue; }
    
    # 对于Windows可执行文件添加.exe后缀
    if [ "${OS}" == "windows" ]; then
        mv "${TMP_DIR}/${EXECUTABLE_NAME}" "${TMP_DIR}/${EXECUTABLE_NAME}.exe"
        EXECUTABLE_NAME="${EXECUTABLE_NAME}.exe"
    fi
    
    # 将 config.toml 复制到临时目录
    cp "${PROJECT_ROOT}/config.toml" "${TMP_DIR}/"

    echo "正在打包 ${ZIP_FILE}..."
    cd "${TMP_DIR}"
    zip -r "../../${ZIP_FILE}" ./*
    cd ..
    
    # 如果是Windows，还原EXECUTABLE_NAME以便后续循环
    if [ "${OS}" == "windows" ]; then
        EXECUTABLE_NAME="ffmpegforBililiveRecorder"
    fi
    cd ..
    echo "打包完成: ${ZIP_FILE}"
done

echo "所有操作完成."
