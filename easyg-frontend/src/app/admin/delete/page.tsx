"use client";
import React from "react";
import Link from "next/link";
import axios from "axios";
import { useState, useEffect } from "react";
import Container from "react-bootstrap/esm/Container";
import Row from "react-bootstrap/esm/Row";
import Col from "react-bootstrap/esm/Col";
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import query from "../../utils/query";
import CloseButton from 'react-bootstrap/CloseButton';
export default function Delete() {
  const [data, setData] = useState<{ uid: string , title: string}[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [showModal, setShowModal] = useState(false);
  const [modalMessage, setModalMessage] = useState("");
  const [iserror, setIserror] = useState(false);
  const [uid, setUid] = useState<string>("");
  useEffect(() => {
    query
      .get("/api/post/getposts")
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

  const deletepost = async () => {
    //console.log(uid);
    try {
        await query.delete(`/api/post/delete/${uid}`);
        if (data) {
            setData(data.filter((post: { uid: string }) => post.uid !== uid));
          } else {
            setData([]); // 或者根据业务逻辑设置一个默认值
          }
        setIserror(true)
        setModalMessage("删除成功!");
        setUid("");
      } catch (error : any) {
        //console.error("Error deleting post:", error);
        setIserror(true);
        setShowModal(true);
        setModalMessage(error.response.data.error);
      }    
  };

  const handleDelete = (post: { title: string; uid: string; }) => {
    //console.log(post);
    setShowModal(true);
    setModalMessage("确定删除" + post.title + "吗?");
    setUid(post.uid);
  }

  const handleClose = () => {
    setShowModal(false);
    setModalMessage("");
    setIserror(false);
  }

  return (
    <div>
      <br />
      {data?.map((post: { uid: string; title: string} ) => (
        <div key={post.uid}>
          <Container>
            <Row>
              <Col md={5}>
                <Link href={`/post/${post.uid}`} className="text-decoration-none">{post.title}</Link>
                
              </Col>
              <Col md={3}>
                <CloseButton 
                  onClick={() => handleDelete(post)}
                />
              </Col>
            </Row>
          </Container>
          

          <hr className="w-50"/>
          </div>
      ))}
      
      <Modal show={showModal} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Message</Modal.Title>
        </Modal.Header>
        <Modal.Body>{modalMessage}</Modal.Body>
        <Modal.Footer>
          <Button variant="success" onClick={handleClose} className={iserror ? "d-none" : ""}>
            取消
          </Button>
          <Button variant="danger" onClick={iserror ? handleClose : deletepost}>
            确定
          </Button>
        </Modal.Footer>
      </Modal>
    </div> 
  );
}
