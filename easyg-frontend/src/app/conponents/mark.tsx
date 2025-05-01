import React from 'react';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import rehypeRaw from 'rehype-raw';
import SyntaxHighlighter  from 'react-syntax-highlighter';
import { docco } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import "github-markdown-css";
type Props = {
  content: string;
};

const MarkdownRenderer = ({ content }: Props) => {
  return (
    <div className="markdown-body"> 
      <ReactMarkdown
        children={content}
        remarkPlugins={[remarkGfm]}
        rehypePlugins={[rehypeRaw]}
        components={{
          code({ node, className, children, ...props }) {
            const match = /language-(\w+)/.exec(className || '');
            const isBlock = node?.tagName === 'code';
            return isBlock && match ? (
              <SyntaxHighlighter
                language={match[1]}
                PreTag="div"
                // @ts-ignore
                style={ docco }
                {...props}
              >
                {String(children).replace(/\n$/, '')}
              </SyntaxHighlighter>
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
};

export default MarkdownRenderer;
