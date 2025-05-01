import axios from "axios"
import { Container, Row, Col } from "react-bootstrap"

export default async function Page({
    params,
  }: {
    params: Promise<{ uid: string }>
  }) {
    const { uid } = await params
    var post = {
      title: "",
      content: "",
      tags: [],
    };
    if (uid) {
      const res = await fetch(`http://localhost:8080/api/post/post/${uid}`)
      const data = await res.json()
      post = data.post
      console.log(data)
    }  
    return (
      <div style={{ height: "100%" }}>
        <div className="h-75">
          <h2 className="text-center pb-3 mb-3">{post.title}</h2>
          <Container className="" style={{ overflowWrap: "break-word" }}>
            <Row className="justify-content-md-center">
              <Col md={{ span: 8 }}>
                {/* <Mark>{data.content}</Mark> */}
                {post.content}
              </Col>
            </Row>
          </Container>
        </div>
        <div>
          <hr />
          <p>
            tag: <br />
            {/* {data.tags.map((tag) => (
              <span key={tag}>{tag} </span>
            ))} */}
          </p>
        </div>
      </div>
    );
  }