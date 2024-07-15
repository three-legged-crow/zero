#!/bin/bash

##############################################################
#                                                            #
#  @(#)generate-docs.sh  February 15th, 2022                 #
#  Copyright(c) 2022, Leyton Goth. All rights reserved.      #
#                                                            #
##############################################################

# 生成当前项目的 go doc, md 格式, 放到 docs 目录中

# 使用当前脚本目录作为工作目录
cd "$(dirname "${BASH_SOURCE[0]}")"

#######################################
# 使用 gomarkdoc 生成当前项目的 go doc, md 格式, 放到 docs 目录中
# Globals:
#   None
# Arguments:
#   None
# Outputs:
#   STDOUT
# Returns:
#   None
#######################################
function main() {
  # 安装 gomarkdoc
  if ! command -v gomarkdoc &> /dev/null; then
    go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
  fi

  # 切换到上一级目录
  cd ..

  if [ ! -d docs ]; then
    mkdir docs
  fi
  # 生成 go doc, 总入口
  gomarkdoc --output docs/readme.md ./...
  # 每个目录生成
  gomarkdoc --output '{{.Dir}}/readme.md' ./...
  echo 'OK'
}

#
main
