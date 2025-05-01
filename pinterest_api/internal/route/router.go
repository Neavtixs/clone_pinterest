package route

import (
	picture "pinterest_api/internal/app/Picture"
	"pinterest_api/internal/app/comment"
	"pinterest_api/internal/app/follow"
	likePost "pinterest_api/internal/app/like_post"
	"pinterest_api/internal/app/person"
	"pinterest_api/internal/app/post"
	"pinterest_api/internal/app/save"
	"pinterest_api/internal/app/user"
	"pinterest_api/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

type RouteConfig struct {
	App                *fiber.App
	UserController     *user.UserController
	PostController     *post.PostController
	FollowController   *follow.FollowController
	PictureController  *picture.PictureController
	SaveController     *save.SaveController
	LikePostController *likePost.LikePostController
	CommentController  *comment.CommentController
	PersonController   *person.PersonController
}

func (c *RouteConfig) Setup() {
	c.App.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:3000, http://127.0.0.1:4000, https://lawrient.xyz, https://api.lawrient.xyz",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization, Origin, Accept",
		AllowCredentials: true,
	}))

	c.SetupGuestRoute()
	c.SetupAuthRoute()
	c.SetupWebSocketRoutes()
	c.App.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(model.WebResponse[string]{
			StatusCode: fiber.StatusNotFound,
			Errors:     "Route is not found",
		})
	})

}

func (c *RouteConfig) SetupGuestRoute() {
	api := c.App.Group("/api")

	api.Post("/auth/register", c.UserController.HandleRegisterByEmail)
	api.Post("/auth/login", c.UserController.HandleLoginByEmail)
	api.Get("/auth/google", c.UserController.HandleGoogleRedirect)
	api.Get("/auth/google/callback", c.UserController.HandleGoogleCallback)
	api.Get("/bye", c.UserController.Logout)

	api.Get("/user", c.UserController.HandleGetUser)
	api.Put("/user/update_birth/:birth_date", c.UserController.HandleUpdateBirthDate)
	api.Put("/user/update", c.UserController.HandleUpdateUser)

	api.Get("/users/:name", c.UserController.HandleGetListByName)

	api.Get("/img/:filename", c.PictureController.GetPicture)
	api.Post("/img", c.PictureController.UploadPicture)

	api.Post("/post", c.PostController.HandleUpload)
	api.Get("/post/:postId", c.PostController.HandleShowDetail)
	api.Get("/posts/", c.PostController.HandleShowRandomList)
	api.Get("/posts/:username", c.PostController.HandleListByUsername)
	api.Delete("/post/:postId", c.PostController.HandleDelete)

	api.Post("/follow/:username", c.FollowController.HandleFollowUser)
	api.Delete("/unfollow/:username", c.FollowController.HandleUnFollowUser)

	api.Get("/user/:username", c.FollowController.HandleShowFollowByUsername)

	api.Post("/save_post/:postid", c.SaveController.HandleSavePost)
	api.Delete("/unsave_post/:postid", c.SaveController.HandleUnSavePost)

	api.Post("/like_post/:postid", c.LikePostController.HandleLikeaPost)
	api.Delete("/unlike_post/:postid", c.LikePostController.HandleUnLikeaPost)
}

func (c *RouteConfig) SetupWebSocketRoutes() {
	c.App.Use("/ws", func(ctx *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(ctx) {
			return ctx.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	c.App.Get("/ws/comment/:postid", websocket.New(c.CommentController.Handle))
}

func (c *RouteConfig) SetupAuthRoute() {
}
