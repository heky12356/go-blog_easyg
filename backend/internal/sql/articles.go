package sql

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

// 常量定义
const (
	MaxCategoryLevel = 3 // 最大分类层级深度
)

// 自定义错误
var (
	ErrCategoryNotFound     = errors.New("分类不存在")
	ErrCategoryLevelTooDeep = errors.New("分类层级过深")
	ErrCategoryNameExists   = errors.New("同级分类名称重复")
	ErrCategoryHasArticles  = errors.New("分类下还有文章，无法删除")
	ErrCategoryHasChildren  = errors.New("分类下还有子分类，无法删除")
)

type Tag struct {
	gorm.Model
	Name    string    `gorm:"uniqueIndex"`
	Aticles []Article `gorm:"many2many:article_tags;"`
}

type Category struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Description string
	ParentID    *uint      `gorm:"index"` // 父分类ID，根分类为nil
	Parent      *Category  `gorm:"foreignKey:ParentID"`
	Children    []Category `gorm:"foreignKey:ParentID"`
	Articles    []Article  `gorm:"many2many:article_categories;"`
}

type Article struct {
	BaceModel
	Content    string
	Title      string
	Uid        string     `gorm:"primarykey"`
	Tags       []Tag      `gorm:"many2many:article_tags;"`
	Categories []Category `gorm:"many2many:article_categories;"`
}

func AutoMigrateArticle() (err error) {
	err = db.AutoMigrate(&Article{}, &Tag{}, &Category{})
	return
}

func CreatePost(artical Article) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 处理每一个tag
		for i, tag := range artical.Tags {
			var exitTag Tag
			if err := tx.Where("name = ?", tag.Name).First(&exitTag).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					if err := tx.Unscoped().Where("name = ?", tag.Name).First(&exitTag).Error; err == nil {
						// 如果软删除的tag存在，恢复它
						// log.Print(exitTag)
						if err = tx.Unscoped().Model(&exitTag).Update("deleted_at", nil).Error; err != nil {
							return err
						}
						artical.Tags[i] = exitTag
					} else {
						// 如果没有找到软删除的tag，则创建新tag
						err = tx.Create(&tag).Error
						if err != nil {
							return err
						}
						artical.Tags[i] = tag
					}
				}
			} else {
				artical.Tags[i] = exitTag
			}
		}

		// 处理每一个category
		for i, category := range artical.Categories {
			var exitCategory Category
			if err := tx.Where("name = ?", category.Name).First(&exitCategory).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					if err := tx.Unscoped().Where("name = ?", category.Name).First(&exitCategory).Error; err == nil {
						// 如果软删除的category存在，恢复它
						if err = tx.Unscoped().Model(&exitCategory).Update("deleted_at", nil).Error; err != nil {
							return err
						}
						artical.Categories[i] = exitCategory
					} else {
						// 如果没有找到软删除的category，则创建新category
						err = tx.Create(&category).Error
						if err != nil {
							return err
						}
						artical.Categories[i] = category
					}
				}
			} else {
				artical.Categories[i] = exitCategory
			}
		}

		// 创建文章
		err := tx.Create(&artical).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func GetPostsBase() (posts []map[string]interface{}, err error) {
	var post []Article
	// 使用 Preload 加载关联的 Tags 和 Categories
	err = db.Preload("Tags").Preload("Categories").Find(&post).Error
	if err != nil {
		return nil, err
	}
	for _, t := range post {
		tags := t.Tags
		tagsreturn := []string{}
		for _, tag := range tags {
			tagsreturn = append(tagsreturn, tag.Name)
		}

		categories := t.Categories
		categoriesreturn := []string{}
		for _, category := range categories {
			categoriesreturn = append(categoriesreturn, category.Name)
		}

		posts = append(posts, map[string]interface{}{
			"title":      t.Title,
			"uid":        t.Uid,
			"tags":       tagsreturn,
			"categories": categoriesreturn,
		})
	}
	return
}

func GetPostByUid(uid string) (post map[string]interface{}, err error) {
	var article Article
	err = db.Preload("Tags").Preload("Categories").Where("uid = ?", uid).First(&article).Error
	if err != nil {
		return nil, err
	}
	tags := article.Tags
	tagsreturn := []string{}
	for _, tag := range tags {
		tagsreturn = append(tagsreturn, tag.Name)
	}

	categories := article.Categories
	categoriesreturn := []string{}
	for _, category := range categories {
		categoriesreturn = append(categoriesreturn, category.Name)
	}

	post = map[string]interface{}{
		"title":      article.Title,
		"content":    article.Content,
		"uid":        article.Uid,
		"tags":       tagsreturn,
		"categories": categoriesreturn,
	}
	return
}

func DeletePost(uid string) (err error) {
	var article Article
	err = db.Where("uid = ?", uid).First(&article).Error
	if err != nil {
		return err
	}
	if err = db.Model(&article).Preload("Tags").Preload("Categories").First(&article).Error; err != nil {
		return err
	}
	log.Default().Print(article.Tags)

	// 处理标签删除
	for _, tag := range article.Tags {
		count := db.Model(&tag).Association("Aticles").Count()
		if count == 1 {
			err = db.Delete(&tag).Error
			if err != nil {
				return err
			}
		}
	}

	// 处理分类删除
	for _, category := range article.Categories {
		count := db.Model(&category).Association("Articles").Count()
		if count == 1 {
			err = db.Delete(&category).Error
			if err != nil {
				return err
			}
		}
	}

	// 清除文章与标签的关联
	err = db.Model(&article).Association("Tags").Clear()
	if err != nil {
		return err
	}

	// 清除文章与分类的关联
	err = db.Model(&article).Association("Categories").Clear()
	if err != nil {
		return err
	}

	// 删除文章
	err = db.Model(&article).Delete(&article).Error
	if err != nil {
		return err
	}

	log.Default().Print(article.Tags)

	return
}

// CreateCategory 通用创建分类（保留原有功能）
func CreateCategory(category Category) error {
	return db.Create(&category).Error
}

// CreateRootCategory 创建根分类
func CreateRootCategory(name, description string) (*Category, error) {
	// 检查根分类名称重复
	var count int64
	db.Model(&Category{}).Where("parent_id IS NULL AND name = ?", name).Count(&count)
	if count > 0 {
		return nil, ErrCategoryNameExists
	}
	
	category := Category{
		Name:        name,
		Description: description,
		ParentID:    nil, // 根分类
	}
	
	if err := db.Create(&category).Error; err != nil {
		return nil, err
	}
	
	return &category, nil
}

// CreateSubCategory 创建子分类
func CreateSubCategory(name, description string, parentID uint) (*Category, error) {
	// 1. 验证父分类是否存在
	var parent Category
	if err := db.First(&parent, parentID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}
	
	// 2. 检查层级深度
	level, err := getCategoryLevel(parentID)
	if err != nil {
		return nil, err
	}
	if level >= MaxCategoryLevel {
		return nil, ErrCategoryLevelTooDeep
	}
	
	// 3. 检查同级分类名称重复
	var count int64
	db.Model(&Category{}).Where("parent_id = ? AND name = ?", parentID, name).Count(&count)
	if count > 0 {
		return nil, ErrCategoryNameExists
	}
	
	// 4. 创建分类
	category := Category{
		Name:        name,
		Description: description,
		ParentID:    &parentID,
	}
	
	if err := db.Create(&category).Error; err != nil {
		return nil, err
	}
	
	return &category, nil
}

// getCategoryLevel 获取分类的层级深度
func getCategoryLevel(categoryID uint) (int, error) {
	var category Category
	if err := db.First(&category, categoryID).Error; err != nil {
		return 0, err
	}
	
	level := 1
	currentID := categoryID
	
	for {
		var current Category
		if err := db.First(&current, currentID).Error; err != nil {
			return 0, err
		}
		
		if current.ParentID == nil {
			break
		}
		
		level++
		currentID = *current.ParentID
		
		// 防止无限循环
		if level > MaxCategoryLevel+1 {
			return level, nil
		}
	}
	
	return level, nil
}

// GetAllCategories 获取所有分类（包含层级关系）
func GetAllCategories() (categories []Category, err error) {
	err = db.Preload("Children").Where("parent_id IS NULL").Find(&categories).Error
	return
}

// GetCategoryTree 获取完整的分类树
func GetCategoryTree() ([]map[string]interface{}, error) {
	var rootCategories []Category
	err := db.Preload("Children.Children").Where("parent_id IS NULL").Find(&rootCategories).Error
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, category := range rootCategories {
		categoryMap := buildCategoryTree(category)
		result = append(result, categoryMap)
	}

	return result, nil
}

// buildCategoryTree 递归构建分类树
func buildCategoryTree(category Category) map[string]interface{} {
	categoryMap := map[string]interface{}{
		"id":          category.ID,
		"name":        category.Name,
		"description": category.Description,
		"parent_id":   category.ParentID,
	}

	if len(category.Children) > 0 {
		var children []map[string]interface{}
		for _, child := range category.Children {
			children = append(children, buildCategoryTree(child))
		}
		categoryMap["children"] = children
	}

	return categoryMap
}

// GetCategoryByID 根据ID获取分类
func GetCategoryByID(id uint) (category Category, err error) {
	err = db.Preload("Parent").Preload("Children").First(&category, id).Error
	return
}

// UpdateCategory 更新分类
func UpdateCategory(category Category) error {
	return db.Save(&category).Error
}

// DeleteCategory 删除分类（如果没有文章关联）
func DeleteCategory(id uint) error {
	var category Category
	if err := db.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrCategoryNotFound
		}
		return err
	}
	
	// 检查是否有文章关联
	count := db.Model(&category).Association("Articles").Count()
	if count > 0 {
		return ErrCategoryHasArticles
	}
	
	// 检查是否有子分类
	var childCount int64
	db.Model(&Category{}).Where("parent_id = ?", id).Count(&childCount)
	if childCount > 0 {
		return ErrCategoryHasChildren
	}
	
	return db.Delete(&category).Error
}

// GetArticlesByCategory 根据分类获取文章
func GetArticlesByCategory(categoryID uint) (articles []Article, err error) {
	var category Category
	err = db.Preload("Articles.Tags").First(&category, categoryID).Error
	if err != nil {
		return nil, err
	}
	return category.Articles, nil
}

// GetCategoryPath 获取分类的完整路径
func GetCategoryPath(categoryID uint) ([]Category, error) {
	var path []Category
	currentID := categoryID
	
	for {
		var category Category
		if err := db.First(&category, currentID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, ErrCategoryNotFound
			}
			return nil, err
		}
		
		path = append([]Category{category}, path...) // 前置插入
		
		if category.ParentID == nil {
			break
		}
		
		currentID = *category.ParentID
	}
	
	return path, nil
}

// GetCategoryPathString 获取分类的路径字符串
func GetCategoryPathString(categoryID uint, separator string) (string, error) {
	path, err := GetCategoryPath(categoryID)
	if err != nil {
		return "", err
	}
	
	var pathNames []string
	for _, category := range path {
		pathNames = append(pathNames, category.Name)
	}
	
	return joinStrings(pathNames, separator), nil
}

// joinStrings 连接字符串数组
func joinStrings(strs []string, separator string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += separator + strs[i]
	}
	return result
}

// MoveCategoryToParent 移动分类到新的父分类下
func MoveCategoryToParent(categoryID uint, newParentID *uint) error {
	var category Category
	if err := db.First(&category, categoryID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrCategoryNotFound
		}
		return err
	}
	
	// 如果有新父分类，验证其存在性和层级
	if newParentID != nil {
		var parent Category
		if err := db.First(&parent, *newParentID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return ErrCategoryNotFound
			}
			return err
		}
		
		// 检查是否会造成循环引用
		if *newParentID == categoryID {
			return errors.New("不能将分类移动到自己下面")
		}
		
		// 检查新层级深度
		level, err := getCategoryLevel(*newParentID)
		if err != nil {
			return err
		}
		if level >= MaxCategoryLevel {
			return ErrCategoryLevelTooDeep
		}
		
		// 检查同级名称重复
		var count int64
		db.Model(&Category{}).Where("parent_id = ? AND name = ? AND id != ?", 
			*newParentID, category.Name, categoryID).Count(&count)
		if count > 0 {
			return ErrCategoryNameExists
		}
	} else {
		// 移动到根级别，检查根级别名称重复
		var count int64
		db.Model(&Category{}).Where("parent_id IS NULL AND name = ? AND id != ?", 
			category.Name, categoryID).Count(&count)
		if count > 0 {
			return ErrCategoryNameExists
		}
	}
	
	// 执行移动
	category.ParentID = newParentID
	return db.Save(&category).Error
}

// GetCategoryStats 获取分类统计信息
func GetCategoryStats(categoryID uint) (map[string]interface{}, error) {
	var result struct {
		ID             uint    `json:"id"`
		Name           string  `json:"name"`
		Description    string  `json:"description"`
		ParentID       *uint   `json:"parent_id"`
		DirectChildren int64   `json:"direct_children"`
		DirectArticles int64   `json:"direct_articles"`
	}
	
	// 一次查询获取基本信息和统计数据
	err := db.Raw(`
		SELECT 
			c.id,
			c.name,
			c.description,
			c.parent_id,
			COALESCE(child_count.cnt, 0) as direct_children,
			COALESCE(article_count.cnt, 0) as direct_articles
		FROM categories c
		LEFT JOIN (
			SELECT parent_id, COUNT(*) as cnt
			FROM categories 
			WHERE deleted_at IS NULL 
			GROUP BY parent_id
		) child_count ON c.id = child_count.parent_id
		LEFT JOIN (
			SELECT category_id, COUNT(*) as cnt
			FROM article_categories
			GROUP BY category_id
		) article_count ON c.id = article_count.category_id
		WHERE c.id = ? AND c.deleted_at IS NULL
	`, categoryID).Scan(&result).Error
	
	if err != nil {
		return nil, err
	}
	
	// 检查分类是否存在
	if result.ID == 0 {
		return nil, ErrCategoryNotFound
	}
	
	// 简单统计总子分类数量（递归）
	var totalChildren int64
	db.Raw(`
		WITH RECURSIVE category_tree AS (
			SELECT id FROM categories WHERE parent_id = ? AND deleted_at IS NULL
			UNION ALL
			SELECT c.id FROM categories c
			JOIN category_tree ct ON c.parent_id = ct.id
			WHERE c.deleted_at IS NULL
		)
		SELECT COUNT(*) FROM category_tree
	`, categoryID).Scan(&totalChildren)
	
	stats := map[string]interface{}{
		"id":              result.ID,
		"name":            result.Name,
		"description":     result.Description,
		"parent_id":       result.ParentID,
		"direct_children": result.DirectChildren,
		"total_children":  totalChildren,
		"direct_articles": result.DirectArticles,
	}
	
	return stats, nil
}
