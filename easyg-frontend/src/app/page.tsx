import { Container } from "react-bootstrap";
import Link from "next/link";
import { Key, ReactElement, JSXElementConstructor, ReactNode, ReactPortal } from "react";

export default async function Home() {
  const res = await fetch('http://localhost:8080/api/post/getposts');
  const data = await res.json();
  const posts = data.posts;
  // console.log(data)
  return (
    <>  
    <div>
      <div>Home</div>
      <br />
      {posts.map((post: { uid: string; title: string; }) => (
        <div key={post.uid}>
          <Link href={`/post/${post.uid}`} className="text-decoration-none">{post.title}</Link>
          <hr />
        </div>
      ))}
    </div>
    </>
  );
}
