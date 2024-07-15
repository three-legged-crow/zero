#!/bin/bash

##############################################################
#                                                            #
#  @(#)git-tag.sh  Saturday, October 8th, 2022               #
#  Copyright(c) 2022, Leyton Goth. All rights reserved.      #
#                                                            #
##############################################################

# 从 master 生成最新的 TAG

# 进入到脚本的目录
cd "$(dirname "${BASH_SOURCE[0]}")"

# 记录当前系统名
WORKING_OS="$(uname -s)"
# CPU 架构
ARCH="$(uname -m)"

# Mac 系统名
OS_DARWIN="Darwin"
# linux
OS_LINUX="Linux"

# 兼容 Linux 和 Mac 命令
sed0="sed"

if [ "$OS_DARWIN" == "$WORKING_OS" ]; then
  sed0="gsed"
fi

# 大版本号取值
SEM_VER_X=0
SEM_VER_Y=0

#######################################
# 从 master 最新的提交创建新的 TAG
# TAG 命名遵从: https://semver.org/
# three-eyed-raven 的命名方式: vX.Y.Z-BRANCH+BUILD_NO
#   X, Y 从常量取值, 来自 $SEM_VER_X, $SEM_VER_Y
#   Z: 每次增 1
#   BRANCH: 当为 master 时, 会被替换为 woa
#   BUILD_NO: 由 DATE-TAG 个数决定
# Globals:
#   None
# Arguments:
#   [1] <string> 分支名字, TAG 来源
# Outputs:
#   STDOUT
# Returns:
#   None
#######################################
function create_tag() {
  current_branch="$(git branch | awk '/\*/ { print $2; }')"
  echo "当前分支: $current_branch => 目标分支: $1"
  # 更新最新代码
  git stash
  git checkout "$1"
  git pull

  echo "---------------------------------------------------"
  git tag -l | xargs git tag -d > /dev/null
  git fetch --tags -q
  local last_tag="$(git for-each-ref --format='%(*committerdate:raw)%(committerdate:raw) %(refname:short) %(objectname:short)' 'refs/tags' | sort -nr | head -n 1 | awk '{print $3}')"
  echo "Last Tag: $last_tag"
  # 检查 $1 的 hash
  target_hash=$(git log -n 200 --pretty='%h %D' | grep "origin/$1\(,\|)\|\$\)" | awk '{print $1}' | head -n 1)
  if [ -z "$target_hash" ]; then
    echo "请将近期提交的代码合并到 $1 再打 TAG."
    git checkout "$current_branch"
    git stash pop
    return
  fi
  echo "基线 SHA1 => [$target_hash]"
  local sem_ver_z=0
  # 检查当前 y 版本是否有变化
  local last_sem_ver_y=$(echo "$last_tag" | awk -F '-' '{print $1}' | awk -F '.' '{print $2}')
  if [ "$last_sem_ver_y" -eq "$SEM_VER_Y" ]; then
    sem_ver_z=$(($(echo "$last_tag" | awk -F '-' '{print $1}' | awk -F '.' '{print $NF}') + 1))
  fi
  local pre_name="$current_branch"
  if [ "$1" == "master" ]; then
    pre_name="woa"
  else
    pre_name=$(echo "$current_branch" | "$sed0" 's|/|.|g')
  fi
  # date
  date0=$(date '+%Y%m%d')
  # build no
  count=$(($(git tag | wc -l) + 1))
  build_no="$(printf '%04d' $count)"
  # full name
  tag_name=$(echo "v$SEM_VER_X.$SEM_VER_Y.$sem_ver_z-${pre_name}+${date0}-b${build_no}" | "$sed0" 's/_/-/g')
  echo "New Tag: $tag_name"
  echo "---------------------------------------------------"

  git tag "$tag_name" "$target_hash"
  git push origin "$tag_name"

  git checkout "$current_branch"
  git stash pop
  echo "---------------------------------------------------"
  echo "OK"
  echo "==================================================="
}



#######################################
# bash 执行入口函数
# Globals:
#   None
# Arguments:
#   [1] <string> 是否使用当前分支来打 TAG, 当为字符串 'true' 时, 使用当前分支打 TAG
# Outputs:
#   STDOUT
# Returns:
#   None
#######################################
function main() {
  if [ "$1" == "true" ]; then
    create_tag "$(git branch | awk '/\*/ { print $2; }')"
  else
    create_tag 'master'
  fi
}

#
main "$1"
