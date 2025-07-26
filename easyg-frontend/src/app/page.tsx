import { Container } from "react-bootstrap";
import Link from "next/link";

export default async function Home() {
  let posts = [];
  try {
    const res = await fetch(`${process.env.API_URL}/api/post/getposts`);
    const data = await res.json();
    posts = data.data;
  } catch (error) {
    console.error('获取文章列表失败:', error);
  }
  // console.log(data)
  return (
    <>  
    <div>
      <div>Home</div>
      <br />
      {(posts == null) || posts.map((post: { uid: string; title: string; }) => (
        <div key={post.uid}>
          <Link href={`/post/${post.uid}`} className="text-decoration-none">{post.title}</Link>
          <hr />
        </div>
      ))}
    </div>
    </>
  );
}
