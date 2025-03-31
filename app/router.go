package app

import (
	"chatter/app/types"
	"chatter/app/handlers"
)

var AppRouter types.Router

func init() {
	AppRouter.Register(
		types.Route{
			Path:    "GET /chatrooms",
			Handler: handlers.ListRooms,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "POST /chatroom/new",
			Handler: handlers.NewChatRoom,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /chatroom/formadd",
			Handler: handlers.RenderFormAdd,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /chatroom/{roomid}/members",
			Handler: handlers.ListChatMembers,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /chatroom/{roomid}/ws",
			Handler: handlers.JoinChatRoom,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /chatroom/{roomid}",
			Handler: handlers.DisplayChatRoom,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "DELETE /chatroom/{roomid}",
			Handler: handlers.DeleteChatRoom,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "DELETE /chatroom/member",
			Handler: handlers.ChatRemoveMember,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "POST /chatroom/addmember",
			Handler: handlers.ChatAddMember,
		},
	)

	//auth
	AppRouter.Register(
		types.Route{
			Path:    "GET /login",
			Handler: handlers.HandleGetLogin,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /loginanon",
			Handler: handlers.HandleLoginAnon,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "POST /login",
			Handler: handlers.HandlePostLogin,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /logout",
			Handler: handlers.HandleGetLogout,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /register",
			Handler: handlers.HandleGetRegister,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "POST /register",
			Handler: handlers.HandlePostRegister,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /user",
			Handler: handlers.HandleGetProfile,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "POST /upload",
			Handler: handlers.HandlePostProfilePic,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "POST /search/users",
			Handler: handlers.UserSearchEmail,
		},
	)

	AppRouter.Register(
		types.Route{
			Path:    "/",
			Handler: handlers.RenderNotFoundPage,
		},
	)
}
