"use client";
import React, { useState, useEffect } from "react";
import Form from "react-bootstrap/Form";
import Spinner from "react-bootstrap/Spinner";
import Alert from "react-bootstrap/Alert";
import "./CategorySelector.css";
import query from "../app/utils/query";

interface Category {
  id: string;
  name: string;
}

interface CategorySelectorProps {
  selectedCategoryIds: string[];
  onCategoryChange: (categoryIds: string[]) => void;
}

const CategorySelector: React.FC<CategorySelectorProps> = ({
  selectedCategoryIds,
  onCategoryChange,
}) => {
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchCategories();
  }, []);

  const fetchCategories = async () => {
    try {
      setLoading(true);
      setError(null);
      
      // 调用后端API获取分类列表
      const response = await query.get("/api/post/getallcategories");
      setCategories(response.data.data || []);
    } catch (err: any) {
      setError(err.response?.data?.error || "获取分类失败");
      console.error("获取分类失败:", err);
    } finally {
      setLoading(false);
    }
  };

  const handleCategoryChange = (categoryId: string, checked: boolean) => {
    if (checked) {
      // 添加分类
      if (!selectedCategoryIds.includes(categoryId)) {
        onCategoryChange([...selectedCategoryIds, categoryId]);
      }
    } else {
      // 移除分类
      onCategoryChange(selectedCategoryIds.filter(id => id !== categoryId));
    }
  };

  if (loading) {
    return (
      <div className="category-selector-container">
        <div className="d-flex justify-content-center align-items-center" style={{ height: "100px" }}>
          <Spinner animation="border" size="sm" />
          <span className="ms-2">加载分类中...</span>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="category-selector-container">
        <Alert variant="danger" className="mb-0">
          {error}
        </Alert>
      </div>
    );
  }

  return (
    <div className="category-selector-container">
      <div className="category-list">
        {categories.map((category) => (
          <Form.Check
            key={category.id}
            type="checkbox"
            id={`category-${category.id}`}
            label={category.name}
            checked={selectedCategoryIds.includes(String(category.id))}
            onChange={(e) => handleCategoryChange(String(category.id), e.target.checked)}
            className="category-option"
          />
        ))}
      </div>
      {categories.length === 0 && (
        <div className="text-muted text-center py-3">
          暂无分类
        </div>
      )}
      {selectedCategoryIds.length > 0 && (
        <div className="mt-2 text-muted small">
          已选择 {selectedCategoryIds.length} 个分类
        </div>
      )}
    </div>
  );
};

export default CategorySelector;