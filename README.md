# Super-IM
## 启动过程
1. 按照mysql、redis、etcd
2. 启动各个rpc服务
3. 启动各个api服务

user_models 【用户表】
user_conf_models 【用户配置表】
friend_models 【好友关系表】
friend_verify_models 【好友申请表】


chat_models 【单聊聊天记录】
group_chat_models 【群聊聊天记录】

group_models 【群组表】
group_member_models【群聊成员表】
group_verify_models【加群申请表】


查询最近的聊天记录：
select least(sender_user_id, recv_user_id) as s_u, greatest(sender_user_id, recv_user_id) as r_u, count(id),max(created_at), max(msg_preview)
from chat_models where sender_user_id = 1 or recv_user_id = 1
group by least(sender_user_id, recv_user_id), greatest(sender_user_id, recv_user_id);
