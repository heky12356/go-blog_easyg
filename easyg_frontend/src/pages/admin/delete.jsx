import React from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import { useState, useEffect } from "react";
import Container from "react-bootstrap/esm/Container";
import Row from "react-bootstrap/esm/Row";
import Col from "react-bootstrap/esm/Col";
export default function Delete() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  useEffect(() => {
    axios
      .get("/api/test/getposts")
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

  const handleClose = (e) => {
    e.preventDefault();
    console.log(1);
  }

  return (
    <div>
      <br />
      {data.map(post => (
        <div key={post.uid}>
          <Container>
            <Row>
              <Col md={5}>
                <Link to={`/post/${post.uid}`} className="text-decoration-none">{post.title}</Link>
                
              </Col>
              <Col md={3}>
                <span onClick={handleClose} style={{cursor:"pointer"}}>x</span>
              </Col>
            </Row>
          </Container>
          

          <hr className="w-50"/>
          </div>
      ))}
    </div> 
  );
}
