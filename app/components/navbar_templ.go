// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"chatter/app/types"
	"fmt"
)

func Navbar() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<nav class=\"navbar\" role=\"navigation\" aria-label=\"main navigation\"><div class=\"navbar-brand\"><a class=\"navbar-item\" href=\"/chatrooms\"><!-- the following code was ai generated --><svg class=\"icon\" viewBox=\"0 0 64 64\" xmlns=\"http://www.w3.org/2000/svg\"><defs><linearGradient id=\"grad\" x1=\"0%\" y1=\"0%\" x2=\"100%\" y2=\"100%\"><stop offset=\"0%\" style=\"stop-color:#00d1b2;stop-opacity:1\"></stop> <stop offset=\"100%\" style=\"stop-color:#00d8c2;stop-opacity:1\"></stop></linearGradient> <filter id=\"shadow\" x=\"-20%\" y=\"-20%\" width=\"140%\" height=\"140%\"><feDropShadow dx=\"2\" dy=\"4\" stdDeviation=\"3\" flood-color=\"#000\" flood-opacity=\"0.3\"></feDropShadow></filter></defs> <circle cx=\"32\" cy=\"32\" r=\"30\" fill=\"url(#grad)\" filter=\"url(#shadow)\"></circle> <path d=\"M20 24h24v4H20zm0 10h24v4H20zm0 10h16v4H20z\" fill=\"#fff\"></path></svg> <svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 410 100\" height=\"60\"><text x=\"0\" y=\"70\" fill=\"#00d1b2\" font-family=\"&#39;Arial Black&#39;, sans-serif\" font-size=\"72\" letter-spacing=\"5\">CHATTER</text></svg><!-- end of ai generated code --></a> <a id=\"navBurger\" role=\"button\" class=\"navbar-burger\" aria-label=\"menu\" aria-expanded=\"false\" data-target=\"navMenu\" _=\"on click \n            toggle .is-active on #navMenu\n            toggle .is-active on #navBurger\n          \"><span aria-hidden=\"true\"></span> <span aria-hidden=\"true\"></span> <span aria-hidden=\"true\"></span> <span aria-hidden=\"true\"></span></a></div><div id=\"navMenu\" class=\"navbar-menu\"><div class=\"navbar-start\"><a class=\"navbar-item\" href=\"/chatrooms\">Lista de Salas</a></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !isUserCtxAuthenticated(ctx) {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<div class=\"navbar-end\"><div class=\"navbar-item\"><div class=\"buttons\"><a class=\"button is-primary\" href=\"/register\"><strong>Registrar</strong></a> <a class=\"button is-light\" href=\"/login\">Login</a></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<div class=\"navbar-end\"><div class=\"navbar-item\"><a class=\"button\" href=\"/user\"><span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(ctx.Value("user").(types.User).Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/navbar.templ`, Line: 74, Col: 61}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</span> <span class=\"icon is-small\"><p class=\"image\"><img class=\"is-rounded\" src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/static/%s", ctx.Value("user").(types.User).ProfilePic))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/navbar.templ`, Line: 77, Col: 120}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\"></p></span></a></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
