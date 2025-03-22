package app

import (
	"chatter/app/types"
)

var AppRouter types.Router

func init() {
	AppRouter.Register(
		types.Route{
			Path:    "GET /chatrooms",
			Handler: ListRooms,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "POST /chatroom/new",
			Handler: NewChatRoom,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /chatroom/formadd",
			Handler: RenderFormAdd,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /chatroom/{roomid}/members",
			Handler: ListChatMembers,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /chatroom/{roomid}/ws",
			Handler: JoinChatRoom,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /chatroom/{roomid}",
			Handler: DisplayChatRoom,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "DELETE /chatroom/{roomid}",
			Handler: DeleteChatRoom,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "DELETE /chatroom/member",
			Handler: ChatRemoveMember,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "POST /chatroom/addmember",
			Handler: ChatAddMember,
		},
	)

	//auth
	AppRouter.Register(
		types.Route{
			Path:    "GET /login",
			Handler: HandleGetLogin,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /loginanon",
			Handler: HandleLoginAnon,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "POST /login",
			Handler: HandlePostLogin,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /logout",
			Handler: HandleGetLogout,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /register",
			Handler: HandleGetRegister,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "POST /register",
			Handler: HandlePostRegister,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "GET /user",
			Handler: HandleGetProfile,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "POST /upload",
			Handler: HandlePostProfilePic,
		},
	)
	AppRouter.Register(
		types.Route{
			Path:    "POST /search/users",
			Handler: UserSearchEmail,
		},
	)

	AppRouter.Register(
		types.Route{
			Path:    "/",
			Handler: RenderNotFoundPage,
		},
	)
}
