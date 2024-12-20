// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package css

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func CSS_Global() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n        :root {\n        --primary-color: #004c99;\n        --secondary-color: #0088cc; \n        --background-color: #f5f8fa; \n        --text-color: #333333; \n        --white: #ffffff;\n        --hover-color: #005bb5;\n        }\n\n        body {\n            margin: 0;\n            font-family: Arial, sans-serif;\n            background-color: var(--background-color);\n            color: var(--text-color);\n        }\n\n        .container {\n            max-width: 1200px;\n            margin: 0 auto;\n            padding: 0 20px;\n        }\n\n        /* Ensure footer sticks to the bottom */\n        body.body-layout {\n            display: flex;\n            flex-direction: column;\n            min-height: 100vh; /* Full viewport height */\n            margin: 0;\n        }\n\n        header.site-header {\n            flex: 0 0 auto; /* Header does not grow or shrink */\n        }\n\n        main.content {\n            flex: 1 0 auto; /* Main content takes up remaining space */\n        }\n\n        footer.site-footer {\n            flex: 0 0 auto; /* Footer does not grow or shrink */\n            background-color: var(--primary-color);\n            color: var(--white);\n            padding: 20px 0;\n        }\n\n        .flex {\n            display: flex;\n        }\n\n        .mw-80 {\n            max-width: 80px;\n        }\n\n        .mr-8 {\n            margin-rigth: 8px;\n        }\n\t</style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
