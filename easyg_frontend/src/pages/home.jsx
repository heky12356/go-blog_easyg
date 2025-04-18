import React from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import { useState, useEffect } from "react";
function Home() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await axios.get("api/api/post/getposts");
        if (response.data.posts != null) setData(response.data.posts);
        else setData([]);
        setLoading(false);
      } catch (err) {
        setError(err);
        setLoading(false);
      }
    };
    fetchPosts();    
  }, []);

  if (loading) return <div>加载中...</div>;
  if (error) return <div>发生错误...</div>;
  //console.log(data);

  return (
    <div>
      <div>Home</div>
      <br />
      {data.map(post => (
        <div key={post.uid}>
          <Link to={`/post/${post.uid}`} className="text-decoration-none">{post.title}</Link>
          <hr />
        </div>
      ))}
    </div>
  );
}

export default Home;
