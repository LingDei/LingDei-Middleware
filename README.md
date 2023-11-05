# 灵嘚中间件（LingDei-Middleware）

灵嘚中间件（LingDei-Middleware）是一个基于 Fiber 的 RESTful API 服务，用于提供灵嘚（LingDei）的后端服务。

## 配置

config/config.yaml

```yaml
# 数据库
mysql:
  DSN: 数据库连接字符串

QINIU:
  access_key: 七牛云access_key
  secret_key: 七牛云secret_key
  bucket: 七牛云存储空间名称
```

## 运行

```bash
go mod tidy
go run .
```

## 接口文档

### Security

**ApiKeyAuth**

| apiKey | _API Key_     |
| ------ | ------------- |
| In     | header        |
| Name   | Authorization |

---

### /barrage/add

#### POST

##### Summary

添加 Barrage

##### Description

添加 Barrage

##### Parameters

| Name       | Located in | Description | Required | Schema  |
| ---------- | ---------- | ----------- | -------- | ------- |
| video_uuid | query      | 视频 UUID   | Yes      | string  |
| content    | query      | 弹幕内容    | Yes      | string  |
| second     | query      | 弹幕秒数    | Yes      | integer |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /barrage/count

#### GET

##### Summary

获取 Barrage 数量

##### Description

获取 Barrage 数量

##### Parameters

| Name       | Located in | Description | Required | Schema |
| ---------- | ---------- | ----------- | -------- | ------ |
| video_uuid | query      | 视频 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                           |
| ---- | ----------- | ------------------------------------------------ |
| 200  | OK          | [model.BarrageCountResp](#modelbarragecountresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp)       |

### /barrage/delete

#### DELETE

##### Summary

删除 Barrage

##### Description

删除 Barrage

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| uuid | query      | 弹幕 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /barrage/get

#### GET

##### Summary

获取 Barrage

##### Description

获取 Barrage

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| uuid | query      | 弹幕 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.BarrageResp](#modelbarrageresp)     |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

### /barrage/list

#### GET

##### Summary

获取 Barrage 列表

##### Description

获取 Barrage 列表

##### Parameters

| Name       | Located in | Description | Required | Schema  |
| ---------- | ---------- | ----------- | -------- | ------- |
| video_uuid | query      | 视频 UUID   | Yes      | string  |
| page       | query      | 页数        | No       | integer |
| page_size  | query      | 每页数量    | No       | integer |

##### Responses

| Code | Description | Schema                                         |
| ---- | ----------- | ---------------------------------------------- |
| 200  | OK          | [model.BarrageListResp](#modelbarragelistresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp)     |

### /barrage/recent_list

#### GET

##### Summary

获取最近的 Barrage 列表

##### Description

获取最近的 Barrage 列表

##### Parameters

| Name       | Located in | Description | Required | Schema  |
| ---------- | ---------- | ----------- | -------- | ------- |
| video_uuid | query      | 视频 UUID   | Yes      | string  |
| second     | query      | 弹幕秒数    | Yes      | integer |
| limit      | query      | 几秒内的    | No       | integer |

##### Responses

| Code | Description | Schema                                         |
| ---- | ----------- | ---------------------------------------------- |
| 200  | OK          | [model.BarrageListResp](#modelbarragelistresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp)     |

---

### /category/add

#### POST

##### Summary

添加分类 🦸

##### Description

添加分类

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | query      | 分类名称    | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /category/delete

#### DELETE

##### Summary

删除分类 🦸

##### Description

删除分类

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| uuid | query      | 分类 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /category/get

#### GET

##### Summary

获取分类

##### Description

获取分类

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| uuid | query      | 分类 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.CategoryResp](#modelcategoryresp)   |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

### /category/list

#### GET

##### Summary

获取分类列表

##### Description

获取分类列表

##### Responses

| Code | Description | Schema                                           |
| ---- | ----------- | ------------------------------------------------ |
| 200  | OK          | [model.CategoryListResp](#modelcategorylistresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp)       |

### /category/update

#### POST

##### Summary

更新分类 🦸

##### Description

更新分类

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| uuid | query      | 分类 UUID   | Yes      | string |
| name | query      | 分类名称    | No       | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

---

### /collect/add

#### POST

##### Summary

收藏视频

##### Description

收藏视频

##### Parameters

| Name       | Located in | Description | Required | Schema |
| ---------- | ---------- | ----------- | -------- | ------ |
| video_uuid | query      | 视频 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /collect/check

#### GET

##### Summary

检查是否收藏过

##### Description

检查是否收藏过

##### Parameters

| Name       | Located in | Description | Required | Schema |
| ---------- | ---------- | ----------- | -------- | ------ |
| video_uuid | query      | 视频 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                             |
| ---- | ----------- | -------------------------------------------------- |
| 200  | OK          | [model.CollectStatusResp](#modelcollectstatusresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp)         |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /collect/delete

#### DELETE

##### Summary

取消收藏

##### Description

取消收藏

##### Parameters

| Name       | Located in | Description | Required | Schema |
| ---------- | ---------- | ----------- | -------- | ------ |
| video_uuid | query      | 视频 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /collect/list

#### GET

##### Summary

获取某位用户自己的收藏列表

##### Description

获取某位用户自己的收藏列表

##### Parameters

| Name      | Located in | Description | Required | Schema  |
| --------- | ---------- | ----------- | -------- | ------- |
| page      | query      | 页数        | No       | integer |
| page_size | query      | 每页数量    | No       | integer |

##### Responses

| Code | Description | Schema                                         |
| ---- | ----------- | ---------------------------------------------- |
| 200  | OK          | [model.CollectListResp](#modelcollectlistresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp)     |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

---

### /comment/add

#### POST

##### Summary

添加 Comment

##### Description

添加 Comment

##### Parameters

| Name       | Located in | Description | Required | Schema |
| ---------- | ---------- | ----------- | -------- | ------ |
| video_uuid | query      | 视频 UUID   | Yes      | string |
| content    | query      | 评论内容    | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /comment/count

#### GET

##### Summary

获取 Comment 数量

##### Description

获取 Comment 数量

##### Parameters

| Name       | Located in | Description | Required | Schema |
| ---------- | ---------- | ----------- | -------- | ------ |
| video_uuid | query      | 视频 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                           |
| ---- | ----------- | ------------------------------------------------ |
| 200  | OK          | [model.CommentCountResp](#modelcommentcountresp) |
| 400  | Bad Request | [model.CommentCountResp](#modelcommentcountresp) |

### /comment/delete

#### DELETE

##### Summary

删除 Comment

##### Description

删除 Comment

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| uuid | query      | 评论 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /comment/get

#### GET

##### Summary

获取 Comment

##### Description

获取 Comment

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| uuid | query      | 评论 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.CommentResp](#modelcommentresp)     |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

### /comment/list

#### GET

##### Summary

获取 Comment 列表

##### Description

获取 Comment 列表

##### Parameters

| Name       | Located in | Description | Required | Schema  |
| ---------- | ---------- | ----------- | -------- | ------- |
| video_uuid | query      | 视频 UUID   | Yes      | string  |
| page       | query      | 页数        | No       | integer |
| page_size  | query      | 每页数量    | No       | integer |

##### Responses

| Code | Description | Schema                                         |
| ---- | ----------- | ---------------------------------------------- |
| 200  | OK          | [model.CommentListResp](#modelcommentlistresp) |
| 400  | Bad Request | [model.CommentListResp](#modelcommentlistresp) |

---

### /fan/count

#### GET

##### Summary

获取粉丝数量

##### Description

获取粉丝数量

##### Parameters

| Name        | Located in | Description | Required | Schema |
| ----------- | ---------- | ----------- | -------- | ------ |
| follow_uuid | query      | 用户 UUID   | No       | string |

##### Responses

| Code | Description | Schema                                         |
| ---- | ----------- | ---------------------------------------------- |
| 200  | OK          | [model.FollowCountResp](#modelfollowcountresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp)     |

### /fan/list

#### GET

##### Summary

获取我的粉丝列表

##### Description

获取我的粉丝列表

##### Parameters

| Name      | Located in | Description | Required | Schema  |
| --------- | ---------- | ----------- | -------- | ------- |
| page      | query      | 页数        | No       | integer |
| page_size | query      | 每页数量    | No       | integer |

##### Responses

| Code | Description | Schema                                       |
| ---- | ----------- | -------------------------------------------- |
| 200  | OK          | [model.FollowListResp](#modelfollowlistresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp)   |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

---

### /follow/add

#### POST

##### Summary

关注用户

##### Description

关注用户

##### Parameters

| Name        | Located in | Description     | Required | Schema |
| ----------- | ---------- | --------------- | -------- | ------ |
| follow_uuid | query      | 被关注用户 UUID | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /follow/check

#### GET

##### Summary

获取关注状态

##### Description

获取关注状态

##### Parameters

| Name        | Located in | Description     | Required | Schema |
| ----------- | ---------- | --------------- | -------- | ------ |
| follow_uuid | query      | 被关注用户 UUID | Yes      | string |

##### Responses

| Code | Description | Schema                                           |
| ---- | ----------- | ------------------------------------------------ |
| 200  | OK          | [model.FollowStatusResp](#modelfollowstatusresp) |
| 400  | Bad Request | [model.FollowStatusResp](#modelfollowstatusresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /follow/count

#### GET

##### Summary

获取已关注 up 主的数量

##### Description

获取已关注 up 主的数量

##### Responses

| Code | Description | Schema                                         |
| ---- | ----------- | ---------------------------------------------- |
| 200  | OK          | [model.FollowCountResp](#modelfollowcountresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp)     |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /follow/delete

#### DELETE

##### Summary

取消关注用户

##### Description

取消关注用户

##### Parameters

| Name        | Located in | Description     | Required | Schema |
| ----------- | ---------- | --------------- | -------- | ------ |
| follow_uuid | query      | 被关注用户 UUID | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /follow/list

#### GET

##### Summary

获取我的关注列表

##### Description

获取我的关注列表

##### Parameters

| Name      | Located in | Description | Required | Schema  |
| --------- | ---------- | ----------- | -------- | ------- |
| page      | query      | 页数        | No       | integer |
| page_size | query      | 每页数量    | No       | integer |

##### Responses

| Code | Description | Schema                                       |
| ---- | ----------- | -------------------------------------------- |
| 200  | OK          | [model.FollowListResp](#modelfollowlistresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp)   |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

---

### /like/add

#### POST

##### Summary

给视频点赞

##### Description

给视频点赞

##### Parameters

| Name       | Located in | Description | Required | Schema |
| ---------- | ---------- | ----------- | -------- | ------ |
| video_uuid | query      | 视频 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /like/check

#### GET

##### Summary

检查是否点赞过

##### Description

检查是否点赞过

##### Parameters

| Name       | Located in | Description | Required | Schema |
| ---------- | ---------- | ----------- | -------- | ------ |
| video_uuid | query      | 视频 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                       |
| ---- | ----------- | -------------------------------------------- |
| 200  | OK          | [model.LikeStatusResp](#modellikestatusresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp)   |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /like/count

#### GET

##### Summary

获取某个视频的点赞数量

##### Description

获取某个视频的点赞数量

##### Parameters

| Name       | Located in | Description | Required | Schema |
| ---------- | ---------- | ----------- | -------- | ------ |
| video_uuid | query      | 视频 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.LikeCountResp](#modellikecountresp) |
| 400  | Bad Request | [model.LikeCountResp](#modellikecountresp) |

### /like/delete

#### DELETE

##### Summary

取消点赞

##### Description

取消点赞

##### Parameters

| Name       | Located in | Description | Required | Schema |
| ---------- | ---------- | ----------- | -------- | ------ |
| video_uuid | query      | 视频 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /like/list

#### GET

##### Summary

获取某位用户自己的点赞列表

##### Description

获取某位用户自己的点赞列表

##### Parameters

| Name      | Located in | Description | Required | Schema  |
| --------- | ---------- | ----------- | -------- | ------- |
| page      | query      | 页数        | No       | integer |
| page_size | query      | 每页数量    | No       | integer |

##### Responses

| Code | Description | Schema                                   |
| ---- | ----------- | ---------------------------------------- |
| 200  | OK          | [model.LikeListResp](#modellikelistresp) |
| 400  | Bad Request | [model.LikeListResp](#modellikelistresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

---

### /video/add

#### POST

##### Summary

添加视频

##### Description

添加视频

##### Parameters

| Name          | Located in | Description   | Required | Schema |
| ------------- | ---------- | ------------- | -------- | ------ |
| uuid          | query      | 视频 UUID     | Yes      | string |
| name          | query      | 视频名称      | Yes      | string |
| category_uuid | query      | 视频分类 UUID | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /video/delete

#### DELETE

##### Summary

删除视频 🦸

##### Description

删除视频

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| uuid | query      | 视频 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /video/follow_list

#### GET

##### Summary

获取我关注的用户的视频

##### Description

获取我关注的用户的视频

##### Parameters

| Name      | Located in | Description | Required | Schema  |
| --------- | ---------- | ----------- | -------- | ------- |
| page      | query      | 页数        | No       | integer |
| page_size | query      | 每页数量    | No       | integer |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.VideoListResp](#modelvideolistresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /video/get

#### GET

##### Summary

获取视频

##### Description

获取视频

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| uuid | query      | 视频 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.VideoResp](#modelvideoresp)         |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

### /video/list

#### GET

##### Summary

获取视频列表

##### Description

获取视频列表

##### Parameters

| Name          | Located in | Description   | Required | Schema  |
| ------------- | ---------- | ------------- | -------- | ------- |
| user_uuid     | query      | 用户 UUID     | No       | string  |
| category_uuid | query      | 视频分类 UUID | No       | string  |
| page          | query      | 页数          | No       | integer |
| page_size     | query      | 每页数量      | No       | integer |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.VideoListResp](#modelvideolistresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

### /video/my_list

#### GET

##### Summary

获取我的视频列表

##### Description

获取我的视频列表

##### Parameters

| Name      | Located in | Description | Required | Schema  |
| --------- | ---------- | ----------- | -------- | ------- |
| page      | query      | 页数        | No       | integer |
| page_size | query      | 每页数量    | No       | integer |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.VideoListResp](#modelvideolistresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /video/recommend_list

#### GET

##### Summary

获取推荐的视频列表

##### Description

获取推荐的视频列表

##### Parameters

| Name      | Located in | Description | Required | Schema  |
| --------- | ---------- | ----------- | -------- | ------- |
| page      | query      | 页数        | No       | integer |
| page_size | query      | 每页数量    | No       | integer |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.VideoListResp](#modelvideolistresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

### /video/search

#### GET

##### Summary

搜索视频

##### Description

搜索视频

##### Parameters

| Name      | Located in | Description | Required | Schema  |
| --------- | ---------- | ----------- | -------- | ------- |
| keyword   | query      | 关键词      | No       | string  |
| page      | query      | 页数        | No       | integer |
| page_size | query      | 每页数量    | No       | integer |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.VideoListResp](#modelvideolistresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

### /video/update

#### POST

##### Summary

更新视频 🦸

##### Description

更新视频

##### Parameters

| Name          | Located in | Description   | Required | Schema |
| ------------- | ---------- | ------------- | -------- | ------ |
| uuid          | query      | 视频 UUID     | Yes      | string |
| name          | query      | 视频名称      | No       | string |
| category_uuid | query      | 视频分类 UUID | No       | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /video/upload_token

#### GET

##### Summary

获取上传凭证

##### Description

获取上传凭证

##### Responses

| Code | Description | Schema                                         |
| ---- | ----------- | ---------------------------------------------- |
| 200  | OK          | [model.UploadTokenResp](#modeluploadtokenresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp)     |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /video/views/add

#### POST

##### Summary

添加视频播放量

##### Description

添加视频播放量

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| uuid | query      | 视频 UUID   | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

---

### Models

#### model.Barrage

| Name       | Type    | Description | Required |
| ---------- | ------- | ----------- | -------- |
| content    | string  |             | Yes      |
| second     | integer |             | Yes      |
| timestamp  | integer |             | Yes      |
| user_uuid  | string  |             | Yes      |
| uuid       | string  |             | Yes      |
| video_uuid | string  |             | Yes      |

#### model.BarrageCountResp

| Name  | Type    | Description | Required |
| ----- | ------- | ----------- | -------- |
| code  | integer |             | No       |
| count | integer |             | No       |

#### model.BarrageListResp

| Name         | Type                               | Description | Required |
| ------------ | ---------------------------------- | ----------- | -------- |
| barrage_list | [ [model.Barrage](#modelbarrage) ] |             | No       |
| code         | integer                            |             | No       |

#### model.BarrageResp

| Name    | Type                           | Description | Required |
| ------- | ------------------------------ | ----------- | -------- |
| barrage | [model.Barrage](#modelbarrage) |             | No       |
| code    | integer                        |             | No       |

#### model.Category

| Name | Type   | Description | Required |
| ---- | ------ | ----------- | -------- |
| name | string |             | Yes      |
| uuid | string |             | Yes      |

#### model.CategoryListResp

| Name          | Type                                 | Description | Required |
| ------------- | ------------------------------------ | ----------- | -------- |
| category_list | [ [model.Category](#modelcategory) ] |             | No       |
| code          | integer                              |             | No       |

#### model.CategoryResp

| Name     | Type                             | Description | Required |
| -------- | -------------------------------- | ----------- | -------- |
| category | [model.Category](#modelcategory) |             | No       |
| code     | integer                          |             | No       |

#### model.Collect

| Name       | Type                       | Description | Required |
| ---------- | -------------------------- | ----------- | -------- |
| user_uuid  | string                     |             | Yes      |
| uuid       | string                     |             | Yes      |
| video      | [model.Video](#modelvideo) |             | No       |
| video_uuid | string                     |             | Yes      |

#### model.CollectListResp

| Name         | Type                               | Description | Required |
| ------------ | ---------------------------------- | ----------- | -------- |
| code         | integer                            |             | No       |
| collect_list | [ [model.Collect](#modelcollect) ] |             | No       |

#### model.CollectStatusResp

| Name   | Type    | Description | Required |
| ------ | ------- | ----------- | -------- |
| code   | integer |             | No       |
| status | boolean |             | No       |

#### model.Comment

| Name       | Type                           | Description | Required |
| ---------- | ------------------------------ | ----------- | -------- |
| content    | string                         |             | Yes      |
| profile    | [model.Profile](#modelprofile) |             | No       |
| timestamp  | integer                        |             | Yes      |
| user_uuid  | string                         |             | Yes      |
| uuid       | string                         |             | Yes      |
| video_uuid | string                         |             | Yes      |

#### model.CommentCountResp

| Name  | Type    | Description | Required |
| ----- | ------- | ----------- | -------- |
| code  | integer |             | No       |
| count | integer |             | No       |

#### model.CommentListResp

| Name         | Type                               | Description | Required |
| ------------ | ---------------------------------- | ----------- | -------- |
| code         | integer                            |             | No       |
| comment_list | [ [model.Comment](#modelcomment) ] |             | No       |

#### model.CommentResp

| Name    | Type                           | Description | Required |
| ------- | ------------------------------ | ----------- | -------- |
| code    | integer                        |             | No       |
| comment | [model.Comment](#modelcomment) |             | No       |

#### model.Follow

| Name        | Type                     | Description | Required |
| ----------- | ------------------------ | ----------- | -------- |
| follow_uuid | string                   |             | Yes      |
| user        | [model.User](#modeluser) |             | No       |
| user_uuid   | string                   |             | Yes      |
| uuid        | string                   |             | Yes      |

#### model.FollowCountResp

| Name  | Type    | Description | Required |
| ----- | ------- | ----------- | -------- |
| code  | integer |             | No       |
| count | integer |             | No       |

#### model.FollowListResp

| Name        | Type                             | Description | Required |
| ----------- | -------------------------------- | ----------- | -------- |
| code        | integer                          |             | No       |
| follow_list | [ [model.Follow](#modelfollow) ] |             | No       |

#### model.FollowStatusResp

| Name   | Type    | Description | Required |
| ------ | ------- | ----------- | -------- |
| code   | integer |             | No       |
| status | boolean |             | No       |

#### model.Like

| Name       | Type                       | Description | Required |
| ---------- | -------------------------- | ----------- | -------- |
| user_uuid  | string                     |             | Yes      |
| uuid       | string                     |             | Yes      |
| video      | [model.Video](#modelvideo) |             | No       |
| video_uuid | string                     |             | Yes      |

#### model.LikeCountResp

| Name  | Type    | Description | Required |
| ----- | ------- | ----------- | -------- |
| code  | integer |             | No       |
| count | integer |             | No       |

#### model.LikeListResp

| Name      | Type                         | Description | Required |
| --------- | ---------------------------- | ----------- | -------- |
| code      | integer                      |             | No       |
| like_list | [ [model.Like](#modellike) ] |             | No       |

#### model.LikeStatusResp

| Name   | Type    | Description | Required |
| ------ | ------- | ----------- | -------- |
| code   | integer |             | No       |
| status | boolean |             | No       |

#### model.OperationResp

| Name | Type    | Description | Required |
| ---- | ------- | ----------- | -------- |
| code | integer |             | No       |
| msg  | string  |             | No       |

#### model.Profile

| Name       | Type   | Description | Required |
| ---------- | ------ | ----------- | -------- |
| avatar_url | string |             | No       |
| email      | string |             | No       |
| id         | string |             | Yes      |
| nickname   | string |             | Yes      |

#### model.UploadTokenResp

| Name         | Type    | Description | Required |
| ------------ | ------- | ----------- | -------- |
| code         | integer |             | No       |
| upload_token | string  |             | No       |
| video_uuid   | string  |             | No       |

#### model.User

| Name     | Type                           | Description | Required |
| -------- | ------------------------------ | ----------- | -------- |
| id       | string                         |             | Yes      |
| profile  | [model.Profile](#modelprofile) | Profile     | No       |
| role     | integer                        |             | No       |
| username | string                         |             | Yes      |

#### model.Video

| Name          | Type    | Description | Required |
| ------------- | ------- | ----------- | -------- |
| author_uuid   | string  |             | Yes      |
| category_uuid | string  |             | Yes      |
| name          | string  |             | Yes      |
| thumbnail_url | string  |             | Yes      |
| timestamp     | integer |             | Yes      |
| url           | string  |             | Yes      |
| uuid          | string  |             | Yes      |
| views         | integer |             | No       |

#### model.VideoListResp

| Name       | Type                           | Description | Required |
| ---------- | ------------------------------ | ----------- | -------- |
| code       | integer                        |             | No       |
| video_list | [ [model.Video](#modelvideo) ] |             | No       |

#### model.VideoResp

| Name   | Type                       | Description | Required |
| ------ | -------------------------- | ----------- | -------- |
| base64 | string                     |             | No       |
| code   | integer                    |             | No       |
| video  | [model.Video](#modelvideo) |             | No       |
