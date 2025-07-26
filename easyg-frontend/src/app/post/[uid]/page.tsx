import { Container, Row, Col } from "react-bootstrap"
import Mark from "../../conponents/mark"
export default async function Page({
    params,
  }: {
    params: { uid: string }
  }) {
    const { uid } = await params
    var post = {
      title: "",
      content: "",
      tags: [],
    };
    if (uid) {
      const res = await fetch(`${process.env.API_URL}/api/post/post/${uid}`, {
        next: { revalidate: 60 }, // 启用ISR缓存
      });
      const data = await res.json();
      post = data.data;
      //console.log(data)
    }  
    return (
      <div style={{ height: "100%" }}>
        <div className="h-75">
          <h2 className="text-center pb-3 mb-3">{post.title}</h2>
          <Container className="" style={{ overflowWrap: "break-word" }}>
            <Row className="justify-content-md-center">
              <Col md={{ span: 8 }}>
              <Mark content={post.content} />
              </Col>
            </Row>
          </Container>
        </div>
        <div>
          <hr />
          <p>
            tag: <br />
            {post.tags.map((tag) => (
              <span key={tag}>{tag} </span>
            ))}
          </p>
        </div>
      </div>
    );
  }