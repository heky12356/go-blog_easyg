import React from "react";
import Container from "react-bootstrap/Container";
import { useParams } from "react-router-dom";
import axios from "axios";
import { useState, useEffect } from "react";
import Markdown from "react-markdown";
export default function Post() {
  const { uid } = useParams();
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  useEffect(() => {
    if (uid) {
      axios
        .get("/api/api/post/post/" + uid)
        .then((response) => {
          setData(response.data.post);
          setLoading(false);
        })
        .catch((err) => {
          setError(err);
          setLoading(false);
        });
    } else {
      setLoading(false);
    }
  }, []);

  if (loading) return <div>加载中...</div>;
  if (error) return <div>发生错误...</div>;
  //console.log(uid);
  //console.log(data);

  if (data) {
    return (
      <div style={{ height: "100%" }}>
        <div className="h-75">
          <h2 className="text-center pb-3 mb-3">{data.title}</h2>
          <Container className="" style={{whiteSpace: "pre-wrap" }}>
            <Markdown>{data.content}</Markdown>
          </Container>
        </div>
        <div>
          <hr />
          <p>
            tag: <br />
            {data.tags.map((tag) => (
              <span key={tag}>{tag} </span>
            ))}
          </p>
        </div>
      </div>
    );
  }
  return (
    <div>
      <h2 className="text-center">Give you a happy day</h2>
    </div>
  );
}
