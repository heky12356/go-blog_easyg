import React from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import { useState, useEffect } from "react";
function Home() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  useEffect(() => {
    axios
      .get("api/test/getposts")
      .then((response) => {
        setData(response.data.posts);
        setLoading(false);
      })
      .catch((err) => {
        setError(err);
        setLoading(false);
      });
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
      <Link to="/about">about</Link>
      <span> | </span>
      <Link to="/admin">admin</Link>
    </div>
  );
}

export default Home;
