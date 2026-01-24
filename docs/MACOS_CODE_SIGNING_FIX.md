# macOS 代码签名问题修复 - 开发文档

## 问题描述

macOS 用户通过 npm 安装 BrowserWing 后，运行时会立即被系统终止（killed）。系统日志显示：

```
proc 63530: load code signature error 2 for file "browserwing"
(AppleSystemPolicy) ASP: Security policy would not allow process
```

这是 macOS Gatekeeper 安全策略导致的，因为二进制文件没有经过 Apple 代码签名。

## 根本原因

1. BrowserWing 的二进制文件在 Linux CI 环境中构建
2. 这些二进制文件没有使用 Apple Developer 证书签名
3. npm 安装时从 GitHub 下载的二进制文件被标记 `quarantine` 属性
4. macOS Gatekeeper 默认阻止运行未签名的可执行文件

## 实施的解决方案

### 1. 短期解决方案 - 自动修复脚本

#### A. npm 安装脚本修复 (`npm/install.js`)

在 postinstall 脚本中添加了 `fixMacOSCodeSignature()` 函数：

```javascript
function fixMacOSCodeSignature(binaryPath) {
  // 1. 移除 quarantine 属性
  execSync(`xattr -d com.apple.quarantine "${binaryPath}" 2>/dev/null`);
  
  // 2. 应用 ad-hoc 签名
  execSync(`codesign -s - "${binaryPath}" 2>/dev/null`);
}
```

安装流程：
1. 下载并解压二进制文件
2. 设置可执行权限（chmod +x）
3. **自动移除 quarantine 属性**
4. **自动应用 ad-hoc 代码签名**
5. 显示 macOS 特定的说明信息

#### B. 一键安装脚本修复 (`install.sh`)

在 `install_binary()` 函数中添加了 `fix_macos_signature()` 调用：

```bash
fix_macos_signature() {
  xattr -d com.apple.quarantine "$BINARY_PATH" 2>/dev/null
  codesign -s - "$BINARY_PATH" 2>/dev/null
}
```

### 2. 用户文档

创建了详细的用户指南 (`docs/MACOS_INSTALLATION_FIX.md`)，包含：

- 3种手动修复方法
  1. 移除 quarantine 属性（推荐）
  2. Ad-hoc 签名
  3. 系统偏好设置手动允许
  
- nvm 用户的特殊说明
- 问题原因解释
- 验证步骤

### 3. 文档更新

更新了所有相关文档：

- `npm/README.md` - npm 包说明，添加 macOS 注意事项
- `README.md` - 英文主 README
- `README.zh-CN.md` - 中文主 README

在安装说明中添加了醒目的 macOS 警告：

```markdown
⚠️ macOS Users:
If you encounter a "killed" error when running, fix it with:
```bash
xattr -d com.apple.quarantine $(which browserwing)
```
```

### 4. 安装提示

在两个安装脚本中都添加了 macOS 特定的成功提示：

```
⚠️  macOS Users:
  If the app fails to start, run this command:
  xattr -d com.apple.quarantine $(which browserwing)
  
  See: https://github.com/browserwing/browserwing/blob/main/docs/MACOS_INSTALLATION_FIX.md
```

## 技术细节

### quarantine 属性

macOS 会为从互联网下载的文件添加 `com.apple.quarantine` 扩展属性：

```bash
# 查看属性
xattr -l /path/to/browserwing

# 移除属性
xattr -d com.apple.quarantine /path/to/browserwing
```

### Ad-hoc 签名

使用本地证书对二进制文件进行临时签名：

```bash
# Ad-hoc 签名（使用 - 表示本地签名）
codesign -s - /path/to/browserwing

# 验证签名
codesign -v /path/to/browserwing
```

Ad-hoc 签名的特点：
- 不需要 Apple Developer 证书
- 只在本地机器有效
- 可以绕过 Gatekeeper 的基本检查
- 不能用于 App Store 分发

## 长期解决方案（待实施）

为了彻底解决这个问题，需要：

### 1. 获取 Apple Developer 证书

- 注册 Apple Developer Program ($99/年)
- 创建 Developer ID Application 证书
- 导出证书到 CI 环境

### 2. 在 CI 中签名

修改 GitHub Actions 或发布流程：

```bash
# 导入证书
security import certificate.p12 -P $CERT_PASSWORD

# 签名二进制
codesign --force --options runtime \
  --sign "Developer ID Application: Your Name" \
  browserwing-darwin-arm64
  
# 验证
codesign -v browserwing-darwin-arm64
```

### 3. 公证（Notarization）

提交到 Apple 进行公证：

```bash
# 创建 zip
zip browserwing.zip browserwing-darwin-arm64

# 提交公证
xcrun notarytool submit browserwing.zip \
  --apple-id your@email.com \
  --team-id TEAMID \
  --password app-specific-password \
  --wait

# 装订公证票据
xcrun stapler staple browserwing-darwin-arm64
```

### 4. 更新构建流程

修改 `Makefile` 的 `build-mac` target：

```makefile
build-mac-arm64: copy-frontend
	@echo "Building macOS arm64..."
	cd $(BACKEND_DIR) && GOOS=darwin GOARCH=arm64 go build $(BUILD_TAGS) $(LDFLAGS) \
		-o ../$(BUILD_DIR)/$(APP_NAME)-darwin-arm64 .
	@if [ -n "$(APPLE_CERT)" ]; then \
		codesign --force --options runtime --sign "$(APPLE_CERT)" \
			$(BUILD_DIR)/$(APP_NAME)-darwin-arm64; \
	fi
```

## 测试验证

### 测试场景

1. **npm 全局安装**
   ```bash
   npm install -g browserwing
   browserwing --version
   ```

2. **nvm 环境**
   ```bash
   # 使用 nvm
   nvm use v24.13.0
   npm install -g browserwing
   browserwing --version
   ```

3. **一键安装脚本**
   ```bash
   curl -fsSL https://raw.githubusercontent.com/browserwing/browserwing/main/install.sh | bash
   browserwing --version
   ```

4. **手动下载**
   ```bash
   # 下载二进制
   curl -L -o browserwing https://github.com/browserwing/browserwing/releases/download/v0.0.2/browserwing-darwin-arm64
   chmod +x browserwing
   ./browserwing --version  # 应该失败
   
   # 修复
   xattr -d com.apple.quarantine browserwing
   ./browserwing --version  # 应该成功
   ```

### 验证步骤

```bash
# 1. 检查 quarantine 属性
xattr -l $(which browserwing)

# 2. 检查签名
codesign -v $(which browserwing)

# 3. 运行测试
browserwing --version
browserwing --help

# 4. 检查系统日志（如果失败）
log show --predicate 'eventMessage CONTAINS "browserwing"' --last 5m
```

## 影响的文件

```
修改的文件：
- npm/install.js           (添加 macOS 修复逻辑)
- npm/README.md            (添加 macOS 说明)
- install.sh               (添加 macOS 修复逻辑)
- README.md                (添加 macOS 警告)
- README.zh-CN.md          (添加 macOS 警告)

新增的文件：
- docs/MACOS_INSTALLATION_FIX.md       (用户修复指南)
- docs/MACOS_CODE_SIGNING_FIX.md       (本文档)
```

## 参考资料

- [Apple Gatekeeper Documentation](https://support.apple.com/en-us/HT202491)
- [Code Signing Guide](https://developer.apple.com/library/archive/documentation/Security/Conceptual/CodeSigningGuide/)
- [Notarizing macOS Software](https://developer.apple.com/documentation/security/notarizing_macos_software_before_distribution)
- [codesign man page](https://www.unix.com/man-page/osx/1/codesign/)
- [xattr man page](https://www.unix.com/man-page/osx/1/xattr/)

## 下一步

1. **立即发布** - 发布包含自动修复的新版本
2. **收集反馈** - 监控 GitHub Issues 和用户反馈
3. **考虑证书** - 评估购买 Apple Developer 证书的必要性
4. **自动化测试** - 在 macOS CI 环境中添加测试
