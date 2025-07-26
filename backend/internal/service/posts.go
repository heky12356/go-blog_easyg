package service

import (
	"fmt"

	"goblogeasyg/internal/sql"
	"goblogeasyg/internal/utils"
)

type Post struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Uid     string   `json:"uid"`
}

type PostServiceInterface interface {
	CreatePost(title string, content string, tags []string) error
	GetPosts() ([]Post, error)
	GetPost(uid string) (Post, error)
	DeletePost(uid string) error
}

type PostService struct{}

func NewPostService() PostServiceInterface {
	return &PostService{}
}

// CreatePost 创建文章
func (p *PostService) CreatePost(title string, content string, tagsdata []string) error {
	if title == "" || content == "" {
		return fmt.Errorf("title or content cannot be empty")
	}

	// 获取tag并构造sql.Tag类型结构体
	var tags []sql.Tag
	for _, t := range tagsdata {
		tags = append(tags, sql.Tag{Name: t})
	}

	// 创建uid
	uid, err := utils.CreateUid()
	if err != nil {
		return fmt.Errorf("create uid failed: %w", err)
	}

	// 使用sql.CreatePost插入数据
	err = sql.CreatePost(sql.Article{
		Content: content,
		Title:   title,
		Tags:    tags,
		Uid:     uid,
	})
	if err != nil {
		return fmt.Errorf("create post failed: %w", err)
	}
	return nil
}

// GetPosts 获取全部文章
func (p *PostService) GetPosts() ([]Post, error) {
	var postList []Post
	posts, err := sql.GetPostsBase()
	if err != nil {
		return nil, fmt.Errorf("get posts failed: %w", err)
	}

	for _, post := range posts {
		postList = append(postList, Post{
			Title: post["title"].(string),
			Tags:  post["tags"].([]string),
			Uid:   post["uid"].(string),
		})
	}
	return postList, nil
}

// GetPost 获取文章详情
func (p *PostService) GetPost(uid string) (Post, error) {
	post, err := sql.GetPostByUid(uid)
	if err != nil {
		return Post{}, fmt.Errorf("get post failed: %w", err)
	}
	return Post{
		Title:   post["title"].(string),
		Content: post["content"].(string),
		Tags:    post["tags"].([]string),
		Uid:     post["uid"].(string),
	}, nil
}

// DeletePost 删除文章
func (p *PostService) DeletePost(uid string) error {
	err := sql.DeletePost(uid)
	if err != nil {
		return fmt.Errorf("delete post failed: %w", err)
	}
	return nil
}
