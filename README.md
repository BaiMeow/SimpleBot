# SimpleBot

力求简单简洁，为小项目的开发提供支持

## 事件

### 已经实现或将要实现或迫切需要的事件

- [x] message.group.normal 群聊
- [x] message.private.friend 私聊
- [ ] notice.group_decrease.leave 群员自主退群
- [ ] notice.group_decrease.kick 群员被踢
- [ ] notice.group_decrease.kick_me 自己被踢
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

### 延后实现的事件

- message.group.anonymous 群匿名聊天，建议所有群关闭匿名
- message.group.notice 群系统提示
- message.private.group 群临时会话，建议加机器人好友
- message.private.other 其他私聊，不清楚有什么别的私聊方式
- notice.group_upload 群文件上传提示，onebot没有文件下载api
- notice.group_admin.set 管理设置，管理的事情，管理自己清楚
- notice.group_admin.unset 管理解除，管理的事情，管理自己清楚
- notice.notify.lucky_king 运气王，不能自主收发红包的机器人要啥运气王
- notice.notify.honor 群荣誉，真的有人在乎吗

如果你能证明上述事件有用，issue&pr

## API

### 已经实现或将要实现或迫切需要的API

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
- [ ] get_login_info 获取登录号信息
- [ ] get_friend_list 获取好友列表
- [ ] get_group_info 获取群信息
- [ ] get_group_list 获取群列表
- [ ] get_group_member_info 获取群成员信息
- [ ] get_group_member_list 获取群成员列表
- [ ] get_image 获取图片

### 延后实现的API

- send_msg 发送消息，重复
- send_like 发送好友赞，不在乎
- set_group_anonymous_ban 群组匿名用户禁言，建议所有群关闭匿名
- set_group_admin 群组设置管理员，设置管理这种操作还是自己来吧
- set_group_anonymous 群组匿名，建议所有群关闭匿名
- set_group_name 设置群名，犯不着让机器人改名
- get_stranger_info 获取陌生人信息，获取的信息和没获取区别不大
- get_group_honor_info 获取群荣誉信息，不在乎
- get_cookies 获取 Cookies
- get_csrf_token 获取 CSRF Token
- get_credentials 获取 QQ 相关接口凭证
- get_record 获取语音，网络环境语音交流效率太低，发的爽，收的费事
- can_send_image 检查是否可以发送图片,当然可以
- can_send_record 检查是否可以发送语音，当然可以
- get_status 获取运行状态，一般用不到
- get_version_info 获取版本信息，一般用不到
- set_restart 重启 OneBot 实现，一般用不到
- clean_cache 清理缓存，一般用不到
