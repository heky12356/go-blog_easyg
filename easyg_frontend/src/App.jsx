import React from "react";
import './App.css'
import Container from "react-bootstrap/Container";
import { Routes, Route, Link, useLocation } from "react-router-dom";
import Test from "./pages/test";
import Home from "./pages/home";
import Post from "./pages/post";
import About from "./pages/about";
import Admin from "./pages/admin";
import Create from "./pages/admin/create";
import Delete from "./pages/admin/delete";
import Login from "./pages/user/login";
import Register from "./pages/user/register";
import Footer from "./pages/footer";
import Verify from "./utils/verify";
export default function App() {
  const location = useLocation(); // 获取当前路径
  let titleSuffix = ""; // 初始化动态后缀为空
  // 动态解析路径后缀，只取整个路径中的第一级
  if (location.pathname !== "/") {
    titleSuffix = location.pathname.split("/")[1];
  }
  return (
    <div>
      <Container>
        <div className="d-flex">
          <h1 className="pt-4">
            <Link to={"/"} className="text-decoration-none text-black">
              Blog
            </Link>
          </h1>
          <div className="align-self-end pb-2">
            {titleSuffix ? " | " + titleSuffix : ""}
          </div>
        </div>
        <hr className="mb-5" />
        <Container className="container-height">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/test" element={<Test />} />
            <Route path="/about" element={<About />} />
            <Route
              path="/admin"
              element={
                <Verify>
                  <Admin />
                </Verify>
              }
            >
              <Route path="create" element={<Create />} />
              <Route path="delete" element={<Delete />} />
            </Route>
            <Route path="/post" element={<Post />}>
              <Route path=":uid" element={<Post />} />
            </Route>
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="*" element={<div>404</div>} />
          </Routes>
        </Container>
        <Footer />
      </Container>
    </div>
  );
}
