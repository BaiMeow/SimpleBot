# SimpleBot

[![Go Report Card](https://goreportcard.com/badge/github.com/BaiMeow/SimpleBot)](https://goreportcard.com/report/github.com/BaiMeow/SimpleBot)
[![Go Reference](https://pkg.go.dev/badge/github.com/BaiMeow/SimpleBot.svg)](https://pkg.go.dev/github.com/BaiMeow/SimpleBot)

力求简单简洁，为小项目的开发提供支持的Golang onebot SDK

目前仍然在开发中，最基础的收发文本图片已经实现，这一部分没有意外应该不会有太多的变化，欢迎使用

推荐使用数组格式的消息，但也对cq码字符串消息做了兼容

## 事件

### 常用事件

- [x] message.group.normal 群聊
- [x] message.private.friend 私聊
- [x] notice.group_decrease.leave 群员自主退群*
- [x] notice.group_decrease.kick 群员被踢*
- [x] notice.group_decrease.kick_me 自己被踢*
- [ ] notice.group_increase.approve 群员被同意进群
- [ ] notice.group_increase.invite 群员被邀请进群
- [ ] notice.group_ban.ban 群禁言
- [ ] notice.group_ban.lift_ban 群禁言解除
- [ ] notice.friend_add 加好友被同意
- [ ] notice.group_recall 群撤回
- [ ] notice.friend_recall 好友撤回
- [ ] notice.notify.poke 戳一戳
- [ ] request.friend 加好友请求
- [x] request.group.add 他人加群请求
- [x] request.group.invite 收到加群邀请

> *群员退群的三个事件一起处理

### 以后再实现的事件

- message.group.anonymous 群匿名聊天
- message.group.notice 群系统提示
- message.private.group 群临时会话
- message.private.other 其他私聊
- notice.group_upload 群文件上传提示
- notice.group_admin.set 管理设置
- notice.group_admin.unset 管理解除
- notice.notify.lucky_king 运气王
- notice.notify.honor 群荣誉

## API

### 常用API

> 没有意外的话sdk中api名称是这里的驼峰命名形式

- [x] send_group_msg 发送群消息
- [x] send_private_msg 发送私聊消息
- [ ] delete_msg 撤回消息
- [ ] get_msg 获取消息
- [ ] get_forward_msg 获取合并转发消息
- [ ] set_group_kick 群组踢人
- [ ] set_group_ban 群组单人禁言
- [ ] set_group_whole_ban 群组全员禁言
- [ ] set_group_card 设置群名片（群备注）
- [ ] set_group_leave 退出群组
- [ ] set_group_special_title 设置群组专属头衔
- [ ] set_friend_add_request 处理加好友请求
- [x] set_group_add_request 处理加群请求／邀请
- [x] get_login_info 获取登录号信息*
- [ ] get_friend_list 获取好友列表
- [ ] get_group_info 获取群信息
- [ ] get_group_list 获取群列表
- [ ] get_group_member_info 获取群成员信息
- [ ] get_group_member_list 获取群成员列表
- [ ] get_image 获取图片

> *获取登录号信息：登陆时自动获取，仅保存qq号

### 以后再实现的API

- send_msg 发送消息
- send_like 发送好友赞
- set_group_anonymous_ban 群组匿名用户禁言
- set_group_admin 群组设置管理员
- set_group_anonymous 群组匿名
- set_group_name 设置群名
- get_stranger_info 获取陌生人信息
- get_group_honor_info 获取群荣誉信息
- get_cookies 获取 Cookies
- get_csrf_token 获取 CSRF Token
- get_credentials 获取 QQ 相关接口凭证
- get_record 获取语音
- can_send_image 检查是否可以发送图片
- can_send_record 检查是否可以发送语音
- get_status 获取运行状态
- get_version_info 获取版本信息
- set_restart 重启 OneBot 实现
- clean_cache 清理缓存

## 消息

### 常用消息

- [x] 纯文本
- [x] 图片
- [x] @
- [x] QQ表情
- [ ] 链接分享
- [ ] 推荐好友
- [ ] 推荐群
- [ ] 位置
- [ ] 音乐分享
- [ ] 音乐自定义分享
- [ ] 回复
- [ ] 合并转发
- [ ] 合并转发节点
- [ ] 合并转发自定义节点
- [ ] XML 消息
- [ ] JSON 消息

### 以后再实现的消息

- 语音
- 短视频
- 猜拳魔法消息
- 掷骰子魔法表情
- 窗口抖动（戳一戳）
- 戳一戳
- 匿名发消息

## 连接方式

目前采用正向websocket，今后可能会添加其他连接方式

## 示例

另见[example/](https://github.com/BaiMeow/SimpleBot/tree/main/example)

## 应用案例

- mc服务器综合管理SiS https://github.com/miaoscraft/SiS/tree/SimpleBot_Base