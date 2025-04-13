import React from "react";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import Container from "react-bootstrap/esm/Container";
import Row from "react-bootstrap/esm/Row";
import Col from "react-bootstrap/esm/Col";
import { useState } from "react";
import axios from "axios";
export default function Register() {
    const [username, setUsername] = useState(null);
    const [email, setEmail] = useState(null);
    const [password, setPassword] = useState(null);
    const [confirmPassword, setConfirmPassword] = useState(null);

    function handleSubmit(e) {
        e.preventDefault();
        axios
        .post("/api/api/user/register", {
          username: username,
          email: email,
          password: password,
          confirmPassword: confirmPassword,
        })
        .then((response) => {
          console.log(response.data);
        })
        .catch((error) => {
          console.log(error);
        });
    }

  return (
    <Container style={{ height: "70vh" }} className="pt-5">
      <Row>
        <Col md={{ span: 4, offset: 4 }}>
          <Form>
            <Form.Group className="mb-3">
              <Form.Label>
                {username ? "" : <span className="text-danger">* </span>}
                Username
              </Form.Label>
              <Form.Control type="text" placeholder="Username" onChange={(e) => setUsername(e.target.value)} />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Label>Email</Form.Label>
              <Form.Control type="text" placeholder="Email" onChange={(e) => setEmail(e.target.value)} />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Label>
                {password ? "" : <span className="text-danger">* </span>}
                Password
              </Form.Label>
              <Form.Control type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)} />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Label>
                {confirmPassword ? (
                  ""
                ) : (
                  <span className="text-danger">* </span>
                )}
                Confirm-Password
              </Form.Label>
              <Form.Control type="password" placeholder="Confirm-Password" onChange={(e) => setConfirmPassword(e.target.value)} />
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
