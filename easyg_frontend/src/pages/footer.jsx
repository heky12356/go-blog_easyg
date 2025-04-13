import React from "react";
import { Link } from "react-router-dom";
import { useState, useEffect } from "react";
import Cookies from "js-cookie";
export default function Footer() {
  const [islogged, setIslogged] = React.useState(false);
  useEffect(() => {
    const tokne = Cookies.get("refreshToken");
    if (tokne) {
      setIslogged(true);
    }
  }, []);
  const handlelogout = () => {
    Cookies.remove("refreshToken");
    Cookies.remove("accessToken");
    setIslogged(false);
  };
  return (
    <div className="mt-5">
      <div className="mb-4">
        <Link to="/about">about</Link>
        <span> | </span>
        <Link to="/admin">admin</Link>
        {islogged && (
          <span> | </span>
        )}
        {islogged && (
          <Link to="/#" onClick={handlelogout}>logout</Link>
        )}
        {!islogged && (
          <span> | </span>
        )}
        {!islogged && (
          <Link to="/login">login</Link>
        )}
        {!islogged && (
          <span> | </span>
        )}
        {!islogged && (
          <Link to="/register">register</Link>
        )}
      </div>
    </div>
  );
}
