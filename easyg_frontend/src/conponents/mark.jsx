import React from "react";
import ReactMarkdown from "react-markdown"; //引入
import remarkGfm from "remark-gfm"; // 划线、表、任务列表和直接url等的语法扩展
import rehypeRaw from "rehype-raw"; // 解析标签，支持html语法
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter"; // 代码高亮
import { tomorrow } from "react-syntax-highlighter/dist/esm/styles/prism";
import "github-markdown-css";
export default function Mark({ children }) {
  return (
    <div className="markdown-body">
      <ReactMarkdown
        children={children}
        remarkPlugins={[remarkGfm]} // 仅传递 remarkGfm 插件
        rehypePlugins={[rehypeRaw]} // 仅传递 rehypeRaw 插件
        components={{
          code({ node, inline, className, children, ...props }) {
            const match = /language-(\w+)/.exec(className || "");
            return !inline && match ? (
              <SyntaxHighlighter
                children={String(children).replace(/\n$/, "")}
                style={tomorrow}
                language={match[1]}
                PreTag="div"
                {...props}
              />
            ) : (
              <code className={className} {...props}>
                {children}
              </code>
            );
          },
        }}
      />
    </div>
  );
}
