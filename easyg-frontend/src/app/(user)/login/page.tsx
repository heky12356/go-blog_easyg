"use client";
import React from "react";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import Container from "react-bootstrap/esm/Container";
import Row from "react-bootstrap/esm/Row";
import Col from "react-bootstrap/esm/Col";
import Cookies from "js-cookie";
import { useState } from "react";
import axios from "axios";
import { Navigate } from "react-router-dom";
export default function Login() {
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [isloggedin, setLoggedIn] = useState(false);
  function handleSubmit(e : any) {
    e.preventDefault();
    axios
      .post("/api/api/user/login", {
        username: username,
        password: password,
      })
      .then((response) => {
        //console.log(response.data);
        Cookies.set("refreshToken", response.data.refreshToken, {
          expires: 3,
          secure: true,
          sameSite: "Strict",
        });
        Cookies.set("accessToken", response.data.accessToken, {
          secure: true,
          sameSite: "Strict",
        });
        //console.log("success");
        setLoggedIn(true);
      })
      .catch((error) => {
        console.log(error);
        console.log(error.response.data);
      });
  }

  if (isloggedin) {
    window.location.href = "/";
  }
  return (
    <Container className="pt-5">
      <Row>
        <Col md={{ span: 4, offset: 4 }}>
          <Form>
            <Form.Group className="mb-3">
              <Form.Label>
                {username ? "" : <span className="text-danger">* </span>}
                Username
              </Form.Label>
              <Form.Control
                type="text"
                placeholder="Username"
                onChange={(e) => setUsername(e.target.value)}
              />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Label>
                {password ? "" : <span className="text-danger">* </span>}
                Password
              </Form.Label>
              <Form.Control
                type="password"
                placeholder="Password"
                onChange={(e) => setPassword(e.target.value)}
              />
            </Form.Group>
            <Button variant="primary" type="submit" onClick={handleSubmit}>
              Submit
            </Button>
          </Form>
        </Col>
      </Row>
    </Container>
  );
}
