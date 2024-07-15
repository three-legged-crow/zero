default:
	@echo "✫ 更新 stub 依赖: make stub"
	@echo "✫ 代码检查: make lint"
	@echo "✫ 生成 go doc: make gd"
	@echo "✫ 格式化代码: make fmt"

# golang format code
fmt:
	find . -name '*.go' -exec $(GOPATH)/bin/goimports -w -l {} \;
	@echo "OK"

# stub code and dependencies
stub:
	@echo "OK"

# staticcheck
# 官网: https://staticcheck.io/
# 安装: go install honnef.co/go/tools/cmd/staticcheck@latest
# golangci-lint
# 官网: https://golangci-lint.run/
# 安装: go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1
lint:
	staticcheck ./...
	golangci-lint run

# 创建标准 TAG
# make tag: 从 master 打 TAG
# make tag FROM_BRANCH=true: 从当前分支打 TAG
tag:
	@./build/git-tag.sh $(FROM_BRANCH)

# generate go doc
gd:
	@./build/generate.sh
