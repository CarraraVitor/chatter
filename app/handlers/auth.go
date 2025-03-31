package handlers

import (
	"bufio"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

    "chatter/app/components"
    "chatter/app/database"
    "chatter/app/services"
    "chatter/app/types"
)


func HandleGetLogin(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()
    err := components.LoginPage().Render(ctx, w)
    if err != nil {
        log.Printf("[ERROR] :%s\n", err.Error())
        msg := "<h1> SERVER ERROR </h1>"
        temp, _ := template.New("Resp").Parse(msg)
        temp.Execute(w, nil)
        return
    }
}

func HandlePostLogin(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        log.Printf("[ERROR] :%s\n", err.Error())
        msg := " Preencha todos os campos "
        fmt.Fprintf(w, msg)
        return
    }

    email := r.PostForm.Get("Email")
    password := r.PostForm.Get("Password")

    user, err := database.GetUserByEmail(email)
    if err != nil {
        log.Printf("[ERROR] :%s\n", err.Error())
        msg := " Este usuário não existe "
        fmt.Fprintf(w, msg)
        return
    }

    if !services.VerifyPassword(password, user.Password) {
        msg := " Senha Incorreta "
        fmt.Fprintf(w, msg)
        return
    }
    
    token := services.GenerateSessionToken()
    session := services.CreateSession(token, user.Id)
    services.SetSessionTokenCookie(w, token, session)
    
    w.Header().Add("HX-Redirect", "/chatrooms")
}

func HandleLoginAnon(w http.ResponseWriter, r *http.Request) {
    user := services.AnonUser()
    _, err := services.SaveUser(user)
    if err != nil {
        fmt.Fprintf(w, "<h1> O servidor não conseguiu criar o usuário anônimo. Tente novamente </h1>")
        return
    }
    token := services.GenerateSessionToken()
    session := services.CreateSession(token, user.Id)
    services.SetSessionTokenCookie(w, token, session)
    components.ReloadPage().Render(context.Background(), w)
}

func HandleGetLogout(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("SessionToken")
    if err != nil {
        log.Printf("[ERROR] Failed Getting SessionTokenCookie: %s\n", err.Error())
        msg := "<h1> The Server Failed While LogingOut. Please, Try Again Later.</h1>"
        temp, _ := template.New("Resp").Parse(msg)
        temp.Execute(w, nil)
        return
    }
    token := cookie.Value
    validation, err := services.ValidateSessionToken(token)
    if err != nil {
        log.Printf("[ERROR] Failed Validanting SessionTokenCookie: %s\n", err.Error())
        msg := "<h1> Invalid Session Token. </h1>"
        temp, _ := template.New("Resp").Parse(msg)
        temp.Execute(w, nil)
        return
    }
    session := validation.Session
    services.DeleteSessionTokenCookie(w)
    services.InvalidateSession(session.Id)
    w.Header().Add("HX-Redirect", "/login")
}

func HandleGetRegister(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()
    err := components.RegisterPage().Render(ctx, w)
    if err != nil {
        log.Printf("[ERROR] :%s\n", err.Error())
        msg := "<h1> SERVER ERROR </h1>"
        temp, _ := template.New("Resp").Parse(msg)
        temp.Execute(w, nil)
        return
    }
}


func HandlePostRegister(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        log.Printf("[ERROR] POST Register:%s\n", err.Error())
        msg := "Dados Inválidos"
        fmt.Fprintf(w, msg)
        return
    }

    name := r.PostForm.Get("Name")
    email := r.PostForm.Get("Email")
    password := r.PostForm.Get("Password")
    repeat := r.PostForm.Get("RepeatPassword")

    if password != repeat {
        msg := "The Passwords are different"
        fmt.Fprintf(w, msg)
        return
    }

    
    _, err = database.GetUserByEmail(email)
    if err != nil && err != database.UserNotFound {
        log.Printf("[ERROR] GetUserByEmail: %s\n", err.Error())
        msg := "Ocorreu um erro no servidor. Tente novamente."
        fmt.Fprintf(w, msg)
        return
    }

    if err == nil {
        msg := "Este email já está registrado."
        fmt.Fprintf(w, msg)
        return
    }

    user, err := services.NewUser(name, email, password)
    if err != nil {
        msg := "Senha Inválida."
        fmt.Fprintf(w, msg)
        return
    }

    _, err = services.SaveUser(user)
    if err != nil {
        log.Printf("[ERROR] AddUser: %s\n", err.Error())
        msg := "Ocorreu um erro no servidor ao criar o usuário. Tente novamente."
        fmt.Fprintf(w, msg)
        return
    }

    w.Header().Add("HX-Redirect", "/login")
    return
}

func HandleGetProfile(w http.ResponseWriter, r *http.Request) {
    ctx := services.UserContext(r)
    err := components.UserProfile("Meu Perfil").Render(ctx, w)
    if err != nil {
        log.Printf("[ERROR] GetProfile: %s\n", err.Error())
        msg := "Ocorreu um erro no servidor ao carregar esta página. Tente novamente."
        fmt.Fprintf(w, msg)
        return
    }
}

func HandlePostProfilePic(w http.ResponseWriter, r *http.Request) {
    max_mem := int64(2000000)
    r.ParseMultipartForm(max_mem)

    header := *r.MultipartForm.File["file"][0]
    ftype := strings.Split(header.Header.Get("Content-Type"), "/")
    if (ftype[0] != "image") {
        fmt.Fprint(w, "Invalid file format.")
        return
    }

    file, err := header.Open()
    defer file.Close()
    if err != nil {
        log.Printf("ERROR: post profile pic: failed opening file: %s\n", err)
        fmt.Fprintf(w, "The server failed proccessing the file")
        return
    }

    buf := bufio.NewReader(file)
    content := make([]byte, 2000000)
    n, err := buf.Read(content)
    if err != nil {
        log.Printf("ERROR: post profile pic: failed reading file content: %s\n", err)
        return
    }
    content = content[:n]

    user := services.UserContext(r).Value("user").(types.User)
    uid := user.Id
    img_name := fmt.Sprintf("user_%s.%s", uid, ftype[1])
    img_path := fmt.Sprintf("./static/%s", img_name)
    err  = os.WriteFile(img_path, content, 0755)
    if err != nil {
        log.Printf("ERROR: post profile pic: failed saving image: %s\n", err)
        fmt.Fprintf(w, "Method 2 failed creating file: %s\n", err)
        return
    }

    user.ProfilePic = img_name
    database.UpdateUser(user)
     
    res := fmt.Sprintf(" <figure class=\"image is-128x128\"> <img class=\"is-rounded\" src=\"%s\"/> </figure>", img_path)
    fmt.Fprintf(w, res)
}

func UserSearchEmail(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        log.Printf("user search :%s\n", err)
        return
    }

    search := r.PostForm.Get("Search")
    if search == "" {
        return
    }
    chat_id := r.PostForm.Get("ChatId")
    if chat_id == "" {
        return
    }
    users := services.UserSearchEmail(search)
    ctx := context.Background()
    components.UserSearchResp(chat_id, users).Render(ctx, w)
}
