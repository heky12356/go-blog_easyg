"use client";
import React, { useState, KeyboardEvent } from "react";
import Badge from "react-bootstrap/Badge";
import "./TagInput.css";

interface TagInputProps {
  tags: string[];
  onTagsChange: (tags: string[]) => void;
  placeholder?: string;
  maxTags?: number;
}

const TagInput: React.FC<TagInputProps> = ({
  tags,
  onTagsChange,
  placeholder = "输入标签后按回车添加",
  maxTags = 10,
}) => {
  const [inputValue, setInputValue] = useState("");

  const handleKeyDown = (e: KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      e.preventDefault();
      addTag();
    } else if (e.key === "Backspace" && inputValue === "" && tags.length > 0) {
      // 当输入框为空且按退格键时，删除最后一个标签
      removeTag(tags.length - 1);
    }
  };

  const addTag = () => {
    const trimmedValue = inputValue.trim();
    
    if (trimmedValue === "") return;
    
    // 检查是否重复
    if (tags.includes(trimmedValue)) {
      setInputValue("");
      return;
    }
    
    // 检查是否超过最大数量
    if (tags.length >= maxTags) {
      setInputValue("");
      return;
    }
    
    // 添加新标签
    onTagsChange([...tags, trimmedValue]);
    setInputValue("");
  };

  const removeTag = (indexToRemove: number) => {
    const newTags = tags.filter((_, index) => index !== indexToRemove);
    onTagsChange(newTags);
  };

  return (
    <div className="tag-input-container">
      <div className="tag-input-wrapper">
        {tags.map((tag, index) => (
          <Badge
            key={index}
            bg="primary"
            className="tag-badge me-1 mb-1"
            onClick={() => removeTag(index)}
          >
            {tag}
            <span className="tag-remove ms-1">×</span>
          </Badge>
        ))}
        <input
          type="text"
          className="tag-input"
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
          onKeyDown={handleKeyDown}
          placeholder={tags.length === 0 ? placeholder : ""}
          disabled={tags.length >= maxTags}
        />
      </div>
      {tags.length >= maxTags && (
        <small className="text-muted">已达到最大标签数量限制 ({maxTags})</small>
      )}
    </div>
  );
};

export default TagInput;