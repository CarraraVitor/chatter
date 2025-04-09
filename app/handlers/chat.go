package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"chatter/app/components"
	"chatter/app/services"
)

func DisplayChatRoom(w http.ResponseWriter, r *http.Request) {
	rid := r.PathValue("roomid")
	id, err := uuid.Parse(rid)
	c, ok := services.Chats[id]
	if err != nil || !ok {
        RenderNotFoundPage(w, r)
		return
	}

    user, is_logged := services.UserFromSessionCookie(r)
    if !is_logged {
        err = components.LoginAsAnon().Render(context.Background(), w)
        if err != nil {
            fmt.Fprintf(w, "Ocorreu um erro: \n %s \n", err)
        }
        return
    }

    user_rooms, err := services.ChatIdsByUser(user.Id)
	if err != nil {
        RenderNotFoundPage(w, r)
		return
	}
    if _, ok := user_rooms[id]; !ok {
        RenderNotFoundPage(w, r)
		return
    }

    msgs := services.ListChatMessages(c)
    for i, msg := range msgs {
        msg.Receiver = user
        msgs[i] = msg
    }

    members := services.ChatMembers(c.Id)

	ctx := services.UserContext(r)
	err = components.RenderChat("Chatter", *c, msgs, members).Render(ctx, w)
	if err != nil {
		fmt.Fprintf(w, "Ocorreu um erro: \n %s \n", err)
	}
}


func ListRooms(w http.ResponseWriter, r *http.Request) {
    rooms, err := services.ListChatsWithUser(r)
    if err != nil {
        fmt.Fprintf(w, "Ocorreu um erro carregando a lista de Chats")
        return
    }
	ishx := r.Header.Get("HX-Request")
	ctx := services.UserContext(r)
	if ishx == "true" {
		err := components.ListTable(rooms).Render(ctx, w)
		if err != nil {
			fmt.Fprintf(w, "Ocorreu: \n %s \n", err)
		}
        return
	} else {
		err := components.ListChatPage("Chatter - Listagem", rooms).Render(ctx, w)
		if err != nil {
			fmt.Fprintf(w, "Ocorreu: \n %s \n", err)
		}
        return
	}
}

func JoinChatRoom(w http.ResponseWriter, r *http.Request) {
	user, logged := services.UserFromSessionCookie(r)
    if !logged {
        //TODO: redirect to home page?
        fmt.Fprintf(w, "<h1> Você deve estar logado para acessar essa página </h1")
        return
    }

	rid := r.PathValue("roomid")
	chat_id, err := uuid.Parse(rid)
    if err != nil {
        RenderNotFoundPage(w, r)
		return
    }
    services.ConnectNewClient(w, r, chat_id, user)
}

func NewChatRoom(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("ChatRoomName")
	user, _ := services.UserFromSessionCookie(r)
	err := services.AddNewChat(name, user)
	ctx := context.Background()
	components.FormAddChatResp(err == nil).Render(ctx, w)
}


func DeleteChatRoom(w http.ResponseWriter, r *http.Request) {
	rid := r.PathValue("roomid")
	id, err := uuid.Parse(rid)
	c, ok := services.Chats[id]
	if err != nil || !ok {
        RenderNotFoundPage(w, r)
        return
	}
	services.DeleteChat(c)
}

func ListChatMembers(w http.ResponseWriter, r *http.Request) {
    rid := r.PathValue("roomid")
    id, err := uuid.Parse(rid)
    if err != nil {
        return
    }
    chat, ok := services.Chats[id]
    if !ok {
        return    
    }
    members := services.ChatMembers(id)
    ctx := context.Background()
    components.ChatMembersList(*chat, members).Render(ctx, w)
}

func ChatAddMember(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        renderAddMemberResponse(w, r, false)        
    }

    user_id := r.PostForm.Get("UserId")
    if user_id == "" {
        renderAddMemberResponse(w, r, false)        
    }
    chat_id := r.PostForm.Get("ChatId")
    if chat_id == "" {
        renderAddMemberResponse(w, r, false)        
    }

    err = services.ChatAddMember(user_id, chat_id)
    renderAddMemberResponse(w, r, err == nil)
}

func ChatRemoveMember(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    chat_id := params.Get("ChatId")
    if chat_id == "" {
        return
    }
    user_id := params.Get("UserId")
    if user_id == "" {
        return
    }
    err := services.ChatRemoveMember(user_id, chat_id)
    renderRemoveMemberResponse(w, r, err == nil)

}

func RenderFormAdd(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	components.FormAddChat().Render(ctx, w)
}

func RenderNotFoundPage(w http.ResponseWriter, r *http.Request) {
    ctx := services.UserContext(r)
    components.NotFoundPage("Sala Não Encontrada").Render(ctx, w)
}

func renderAddMemberResponse(w http.ResponseWriter, r *http.Request, s bool) {
    ctx := context.Background()
    components.ChatAddMemberResp(s).Render(ctx, w)
}

func renderRemoveMemberResponse(w http.ResponseWriter, r *http.Request, s bool) {
    ctx := context.Background()
    components.ChatRemoveMemberResp(s).Render(ctx, w)
}
