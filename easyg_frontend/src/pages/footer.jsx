import React from "react";
import { Link } from "react-router-dom";
export default function Footer() {
  return (
    <div>
      <div className="mb-4">
        <Link to="/about">about</Link>
        <span> | </span>
        <Link to="/admin">admin</Link>
        <span> | </span>
        <Link to="/login">login</Link>
        <span> | </span>
        <Link to="/register">register</Link>
      </div>
    </div>
  );
}
