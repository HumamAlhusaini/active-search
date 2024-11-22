// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package examples

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/HumamAlhusaini/active-search/internal/model"
	"github.com/HumamAlhusaini/active-search/internal/views/components"
)

// example
func TestimonialGridExample() templ.Component {
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
		templ_7745c5c3_Err = components.TestimonialGrid(
			"Read what our customers think",
			[]model.Testimonial{
				{Avatar: components.Avatar(model.Avatar{ContainerClass: "rounded h-20", Source: "/static/images/avatar.jpg"}), Name: "Jane Doe", Rating: 5, Content: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Iure impedit, placeat sed provident enim fuga possimus ducimus est iusto inventore earum aliquid officia minus, maiores quia dicta magni ex labore? Lorem ipsum dolor sit amet consectetur adipisicing elit. Libero maxime quos laboriosam natus illum similique id nam, rerum, sunt veritatis dolorum accusamus voluptas odio minus necessitatibus perspiciatis, aliquid repellat iste."},
				{Avatar: components.Avatar(model.Avatar{ContainerClass: "rounded h-20", Source: "/static/images/avatar.jpg"}), Name: "Jane Doe", Rating: 4, Content: "maiores quia dicta magni ex labore? Lorem ipsum dolor sit amet consectetur adipisicing elit. Libero maxime quos laboriosam natus illum similique id nam, rerum, sunt veritatis dolorum accusamus voluptas odio minus necessitatibus perspiciatis, aliquid repellat iste."},
				{Avatar: components.Avatar(model.Avatar{ContainerClass: "rounded h-20", Source: "/static/images/avatar.jpg"}), Name: "Jane Doe", Rating: 3, Content: "Iure impedit, placeat sed provident enim fuga possimus ducimus est iusto inventore earum aliquid officia minus, maiores quia dicta magni ex labore? Lorem ipsum dolor sit amet consectetur adipisicing elit. Libero maxime quos laboriosam natus illum similique id nam, rerum, sunt veritatis dolorum accusamus voluptas odio minus necessitatibus perspiciatis, aliquid repellat iste."},
				{Avatar: components.Avatar(model.Avatar{ContainerClass: "rounded h-20", Source: "/static/images/avatar.jpg"}), Name: "Jane Doe", Rating: 5, Content: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Iure impedit, placeat sed provident enim fuga possimus ducimus est iusto inventore earum aliquid officia minus, maiores quia dicta magni ex labore? Lorem ipsum dolor sit amet consectetur adipisicing elit. Libero maxime quos laboriosam natus illum similique id nam, rerum, sunt veritatis dolorum accusamus voluptas odio minus necessitatibus perspiciatis, aliquid repellat iste."},
				{Avatar: components.Avatar(model.Avatar{ContainerClass: "rounded h-20", Source: "/static/images/avatar.jpg"}), Name: "Jane Doe", Rating: 5, Content: "Lorem ipsum dolor sit amet consectetur adipisicing elit. "},
				{Avatar: components.Avatar(model.Avatar{ContainerClass: "rounded h-20", Source: "/static/images/avatar.jpg"}), Name: "Jane Doe", Rating: 1, Content: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Iure impedit, placeat sed provident enim fuga possimus ducimus est iusto inventore earum aliquid officia minus."},
				{Avatar: components.Avatar(model.Avatar{ContainerClass: "rounded h-20", Source: "/static/images/avatar.jpg"}), Name: "Jane Doe", Rating: 2, Content: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Iure impedit, placeat sed provident enim fuga possimus ducimus est iusto inventore earum aliquid officia minus, maiores quia dicta magni ex labore? Lorem ipsum dolor sit amet consectetur adipisicing elit."},
				{Avatar: components.Avatar(model.Avatar{ContainerClass: "rounded h-20", Source: "/static/images/avatar.jpg"}), Name: "Jane Doe", Rating: 5, Content: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Iure impedit, placeat sed provident enim fuga possimus ducimus est iusto inventore earum aliquid officia minus, maiores quia dicta magni ex labore? Lorem ipsum dolor sit amet consectetur adipisicing elit. Libero maxime quos laboriosam natus illum similique id nam, rerum, sunt veritatis."},
				{Avatar: components.Avatar(model.Avatar{ContainerClass: "rounded h-20", Source: "/static/images/avatar.jpg"}), Name: "Jane Doe", Rating: 4, Content: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Iure impedit, placeat sed provident enim fuga possimus ducimus est iusto inventore earum aliquid officia minus, maiores quia dicta magni ex labore? Lorem ipsum dolor sit amet consectetur adipisicing elit. Libero maxime quos laboriosam natus illum similique id nam, rerum."},
				{Avatar: components.Avatar(model.Avatar{ContainerClass: "rounded h-20", Source: "/static/images/avatar.jpg"}), Name: "Jane Doe", Rating: 4, Content: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Iure impedit, placeat sed provident enim fuga possimus ducimus est iusto inventore earum aliquid officia minus, maiores quia dicta magni ex labore? Lorem ipsum dolor sit amet consectetur adipisicing elit. Libero maxime quos laboriosam natus illum similique id nam, rerum, sunt veritatis dolorum accusamus voluptas."},
			},
		).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
