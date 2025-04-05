export default async function Page() {
    const data = await fetch('http://localhost:8080/test/getposts')
    const jsondata = await data.json()
    const posts = jsondata.posts
    return (
      <ul>
        {posts.map((post: any) => (
          <div key={post.uid}>
            <h1>{post.title}</h1>
            <p>{post.content}</p>
            <hr />  
          </div>
        ))}
      </ul>
    )
  }