"use client";
import React from "react";
import { useState } from "react";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Modal from 'react-bootstrap/Modal';
import query from "../../utils/query";
import TagInput from "../../../components/TagInput";
export default function Create() {
  // 使用 useState 来保存 title、content 和 tags 的值
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [tags, setTags] = useState<string[]>([]);
  const [showModal, setShowModal] = useState(false);
  const [modalMessage, setModalMessage] = useState("");
  const handleClose = () => {
    setShowModal(false);
    setModalMessage("");
  }

  // 当点击按钮或提交表单时，会调用此函数
  const handleSubmit = (e: any) => {
    e.preventDefault(); // 阻止默认表单提交行为

    // 构造 JSON 对象
    const data = {
      title: title,
      content: content,
      tags: tags,
    };
    console.log(data);
    // 使用 axios 实例 query 发送 POST 请求，注意修改 URL 为你的接口地址
    query
      .post("/api/post/create", data)
      .then((response : any) => {
        setShowModal(true);
        setModalMessage(response.data.message);
      })
      .catch((error : any) => {
        setShowModal(true);
        setModalMessage(error.response.data.error);
        //console.log(error);
      });
  };

  return (
  <div>
      <Container className="border mt-4 p-2">
      <Row>
        <Col md={9}>
          <Form>
            <Form.Group className="mb-3" controlId="title">
              <Form.Label>标题</Form.Label>
              <Form.Control
                type="text"
                placeholder=""
                value={title}
                onChange={(e) => setTitle(e.target.value)}
              />
            </Form.Group>
            <Form.Group className="mb-3" controlId="content">
              <Form.Label>内容</Form.Label>
              <Form.Control
                as="textarea"
                rows={17}
                value={content}
                onChange={(e) => setContent(e.target.value)}
              />
            </Form.Group>
          </Form>
        </Col>
        <Col>
          <Form.Label>Tag</Form.Label>
          <TagInput 
            tags={tags} 
            onTagsChange={setTags} 
            placeholder="输入标签后按回车添加" 
            maxTags={10}
          />
        </Col>
      </Row>
      <Row>
        <Button
          className="ms-2 w-50"
          size="sm"
          variant="primary"
          type="submit"
          onClick={handleSubmit}
        >
          Submit
        </Button>
      </Row>
    </Container>

    <Modal show={showModal} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Message</Modal.Title>
        </Modal.Header>
        <Modal.Body>{modalMessage}</Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
  </div>
  );
}
