package handler

import (
	"fmt"

	"goblogeasyg/internal/response"
	"goblogeasyg/internal/service"

	"github.com/gin-gonic/gin"
)

type PostHandlerInterface interface {
	CreatePost(c *gin.Context)
	GetPost(c *gin.Context)
	GetAllPost(c *gin.Context)
	DeletePost(c *gin.Context)
	CreateCategory(c *gin.Context)
	GetAllCategories(c *gin.Context)
}

type PostHandler struct {
	PostService service.PostServiceInterface
}

// GetAllCategories implements PostHandlerInterface.
func (p *PostHandler) GetAllCategories(c *gin.Context) {
	categories, err := p.PostService.GetAllCategories()
	if err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, fmt.Sprintf("get all categories failed: %v", err))
		return
	}
	response.SuccessResponse(c, categories, "get all categories success")
}

// CreateCategory implements PostHandlerInterface.
func (p *PostHandler) CreateCategory(c *gin.Context) {
	var category CreateCategoryRequest
	if err := c.ShouldBindJSON(&category); err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, fmt.Sprintf("create category failed: %v", err))
		return
	}
	if err := p.PostService.CreateCategory(category.Category); err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, fmt.Sprintf("create category failed: %v", err))
		return
	}
	response.SuccessResponse(c, nil, "create category success")
}

// CreatePost implements PostHandlerInterface.
func (p *PostHandler) CreatePost(c *gin.Context) {
	var post CreatePostRequest
	if err := c.ShouldBindJSON(&post); err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, fmt.Sprintf("create post failed: %v", err))
		return
	}
	fmt.Printf("%+v", post)
	if err := p.PostService.CreatePost(post.Title, post.Content, post.Tags, post.Categories); err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, fmt.Sprintf("create post failed: %v", err))
		return
	}
	response.SuccessResponse(c, nil, "create post success")
}

// DeletePost implements PostHandlerInterface.
func (p *PostHandler) DeletePost(c *gin.Context) {
	uid := c.Param("uid")
	if err := p.PostService.DeletePost(uid); err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, fmt.Sprintf("delete post failed: %v", err))
		return
	}
	response.SuccessResponse(c, nil, "delete post success")
}

// GetAllPost implements PostHandlerInterface.
func (p *PostHandler) GetAllPost(c *gin.Context) {
	posts, err := p.PostService.GetPosts()
	if err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, fmt.Sprintf("get all post failed: %v", err))
		return
	}
	response.SuccessResponse(c, posts, "get all post success")
}

// GetPost implements PostHandlerInterface.
func (p *PostHandler) GetPost(c *gin.Context) {
	uid := c.Param("uid")
	post, err := p.PostService.GetPost(uid)
	if err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, fmt.Sprintf("get post failed: %v", err))
		return
	}
	response.SuccessResponse(c, post, "get post success")
}

func NewPostHandler(postService service.PostServiceInterface) PostHandlerInterface {
	return &PostHandler{
		PostService: postService,
	}
}
