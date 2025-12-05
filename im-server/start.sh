cd /Users/yangzhongyu/Desktop/code/go/Super-IM/im-server/im_user/user_rpc &&  go run userrpc.go &

cd /Users/yangzhongyu/Desktop/code/go/Super-IM/im-server/im_gateway && go run gateway.go &

cd /Users/yangzhongyu/Desktop/code/go/Super-IM/im-server/im_file/file_api && go run file.go &

cd /Users/yangzhongyu/Desktop/code/go/Super-IM/im-server/im_auth/auth_api && go auth.go &

cd /Users/yangzhongyu/Desktop/code/go/Super-IM/im-server/im_chat/chat_api && go chat.go &

cd /Users/yangzhongyu/Desktop/code/go/Super-IM/im-server/im_user/user_api && go users.go &

#cd /Users/yangzhongyu/Desktop/code/go/Super-IM/im-webapp && npm run dev &