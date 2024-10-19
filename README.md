- [x] ETCD to register more service 
- [x] ETCD lease
- [x] kafka for user
- [x] double token for the refresh
- [x] Rewrite docker compose
- [ ] rabbit mq for shop
- [ ] websocket for chat and use mongo to save the chat message
- [ ] use neo4j to get the friendship
- [ ] logger and trace for error

[//]: # (All for the location not use cloud)
[//]: # (for cloud)
[//]: # (aws ec2 and use nginx and kubernetes to CICD and use serverless to deploy)

[//]: # (for the front)
[//]: # (Use react and redux OR nuxt and pinpa OR nextjs )

[//]: # (for the block chain i want to use own block chain)

[//]: # (use web3 to do transaction)


# Flande

## 簡介

### 一款快速配對的APP，登入以後可以查看到他人的限時分享
  - 上滑快速配對，下滑則再也不見！
  - 若是反悔，也可以使用每日的工具來讓他重新出現(但會出現在你意想不到的地方！)
#### 拍出一段現實30's的影片或照片盡可能的吸引他人吧！

#### 使用AR的方式讓你創造專屬於你的角色，透過服裝讓你在活動中成為最閃耀的一顆星

#### 每日的任務讓你們聊天永不尷尬，完成的任務越多，雙方就可以有更多的接觸機會

#### 公共的討論區讓你們用著自己的角色模型來相互聊天，或許會發生天雷勾動地火的事情

#### 遇到喜歡的人也可以將他放入我們專屬的膠囊區
  - 在限定的時間內等待你的來訪

---

## 目前為實驗性專案

## Backend port

1. 8082 主入口
2. 5050 登入
3. 5051 註冊
4. 5052 所有使用者

## Docker
```shell
docker compose up -d 
```

## ISSUE規範
- 前端：(front)_change:"pkg_line:msg"
- 後端: (backend)_change:"pkg_line_msg"
- App: (app)_change:"pkg_line_msg"
- UI/UX: (ui/ux)_change:"pkg_line_msg"
- Other: (other)_change:"pkg_line_msg"