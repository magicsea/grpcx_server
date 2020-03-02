# grpcx_server
使用grpc作为c2s,s2s的统一协议的尝试项目
## 动机
- 统一c2s,s2s的协议，都通过proto定义，减少沟通成本，不易犯错
- 使用grpc解决了跨语言，跨平台问题
- 立足游戏领域的解决方案（无状态，有状态服务）


# TODO
- [ ] gate支持bidirection
- [ ] game使用bid实例
- [ ] room使用bid绑定房间实例
- [ ] 抽取框架到独立库
- [ ] 完整游戏实例
